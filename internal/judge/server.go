package judge

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/shirou/gopsutil/load"
	"go.uber.org/multierr"
	"go.uber.org/zap"

	"github.com/mraron/njudge/pkg/problems"
	_ "github.com/mraron/njudge/pkg/problems/config/feladat_txt"
	_ "github.com/mraron/njudge/pkg/problems/config/polygon"
	_ "github.com/mraron/njudge/pkg/problems/tasktype/batch"
	_ "github.com/mraron/njudge/pkg/problems/tasktype/communication"
	_ "github.com/mraron/njudge/pkg/problems/tasktype/stub"

	"encoding/json"

	"github.com/labstack/echo/v4/middleware"
	"github.com/mraron/njudge/pkg/language"
	_ "github.com/mraron/njudge/pkg/language/langs/cpp"
	_ "github.com/mraron/njudge/pkg/language/langs/csharp"
	_ "github.com/mraron/njudge/pkg/language/langs/golang"
	_ "github.com/mraron/njudge/pkg/language/langs/java"
	_ "github.com/mraron/njudge/pkg/language/langs/julia"
	_ "github.com/mraron/njudge/pkg/language/langs/nim"
	_ "github.com/mraron/njudge/pkg/language/langs/pascal"
	_ "github.com/mraron/njudge/pkg/language/langs/pypy3"
	_ "github.com/mraron/njudge/pkg/language/langs/python3"
	_ "github.com/mraron/njudge/pkg/language/langs/zip"
	"github.com/mraron/njudge/pkg/language/sandbox"
)

type Server struct {
	Status      `mapstructure:",squash"`
	statusMutex sync.RWMutex

	Mode        string `json:"mode" mapstructure:"mode"`
	ProblemsDir string `json:"problems_dir" mapstructure:"problems_dir"`
	LogDir      string `json:"log_dir" mapstructure:"log_dir"`
	SandboxIds  string `json:"sandbox_ids" mapstructure:"sandbox_ids"`
	WorkerCount int    `json:"worker_count" mapstructure:"worker_count"`

	minSandboxId, maxSandboxId int
	problemStore               problems.Store
	start                      time.Time
	queue                      chan Submission
	workers                    chan *Worker
	logger                     *zap.Logger
	sandboxIdUsed              map[int]struct{}
}

func NewServer() *Server {
	s := Server{}
	s.start = time.Now()
	s.queue = make(chan Submission, 128)
	s.sandboxIdUsed = make(map[int]struct{})

	return &s
}

func (s *Server) Run() error {
	if err := s.init(); err != nil {
		return err
	}

	e := echo.New()
	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogMethod:   true,
		LogURI:      true,
		LogStatus:   true,
		LogHost:     true,
		LogRemoteIP: true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			s.logger.Info("request",
				zap.String("method", v.Method),
				zap.String("URI", v.URI),
				zap.String("host", v.Host),
				zap.String("remoteip", v.RemoteIP),
				zap.Int("status", v.Status),
			)

			return nil
		},
	}))

	e.GET("/status", s.getStatus)
	e.POST("/update", s.postUpdateProblems)
	e.POST("/judge", s.postJudge)

	go s.runUpdateLoad()
	go s.runUpdateUptime()
	go s.runUpdateProblems()
	go s.runJudger()

	return e.Start(":" + s.Port)
}

func (s *Server) init() error {
	var err error
	if s.Mode == "development" {
		s.logger, err = zap.NewDevelopment()
	} else {
		s.logger, err = zap.NewProduction()
	}
	if err != nil {
		return err
	}

	s.Url = "http://" + s.Host + ":" + s.Port

	if s.SandboxIds == "" {
		s.minSandboxId = 100
		s.maxSandboxId = 999
	} else {
		splitted := strings.Split(s.SandboxIds, "-")
		if len(splitted) != 2 {
			return fmt.Errorf("sandbox_ids wrong format")
		}

		var err1, err2 error
		s.minSandboxId, err1 = strconv.Atoi(splitted[0])
		s.maxSandboxId, err2 = strconv.Atoi(splitted[1])
		if err1 != nil || err2 != nil {
			return multierr.Combine(err1, err2)
		}
	}

	s.logger.Info("initializing workers")
	s.workers = make(chan *Worker, s.WorkerCount)
	for i := 0; i < s.WorkerCount; i++ {
		provider := language.NewSandboxProvider()
		if err := s.populateProvider(provider, 2); err != nil {
			return err
		}

		s.workers <- NewWorker(i+1, provider)
	}

	s.logger.Info("updating languages")
	if err := s.updateLanguages(); err != nil {
		return err
	}
	s.logger.Sugar().Info("languages: ", s.LanguageList)

	s.logger.Info("parsing problems")
	s.problemStore = problems.NewFsStore(s.ProblemsDir)
	return s.updateProblems()
}

func (s *Server) populateProvider(provider *language.SandboxProvider, cnt int) error {
	for j := s.minSandboxId; j <= s.maxSandboxId; j++ {
		if _, ok := s.sandboxIdUsed[j]; !ok {
			provider.Put(sandbox.NewIsolate(j))
			cnt -= 1
			s.sandboxIdUsed[j] = struct{}{}
		}

		if cnt == 0 {
			break
		}
	}

	if cnt != 0 {
		return fmt.Errorf("not enough sandboxes")
	}

	return nil
}

func (s *Server) updateLanguages() error {
	if len(s.LanguageList) == 0 {
		for _, l := range language.List() {
			s.LanguageList = append(s.LanguageList, l.Id())
		}
	}

	return nil
}

func (s *Server) updateProblems() error {
	var (
		err  error
		err2 error
	)

	s.statusMutex.Lock()
	defer s.statusMutex.Unlock()

	err = s.problemStore.Update()
	s.ProblemList, err2 = s.problemStore.List()

	return multierr.Append(err, err2)
}

func (s *Server) runUpdateProblems() {
	for {
		if err := s.updateProblems(); err != nil {
			log.Print(err)
		}
		time.Sleep(20 * time.Second)
	}
}

func (s *Server) runUpdateLoad() {
	for {
		l, err := load.Avg()

		if err != nil {
			log.Print("Error while getting load: ", err)
		} else {
			s.statusMutex.Lock()
			s.Status.Load = l.Load1
			s.statusMutex.Unlock()
		}

		time.Sleep(60 * time.Second)
	}
}

func (s *Server) runUpdateUptime() {
	for {
		s.statusMutex.Lock()
		s.Status.Uptime = time.Since(s.start)
		s.statusMutex.Unlock()
		time.Sleep(1 * time.Second)
	}
}

func (s *Server) runJudger() {
	judge := func(worker *Worker, sub Submission) error {
		defer func() {
			sub.done <- true
		}()

		if ok, err := s.problemStore.Has(sub.Problem); !ok {
			return err
		}

		p, _ := s.problemStore.Get(sub.Problem)
		st, err := worker.Judge(context.Background(), s.logger, p, sub.Source, language.Get(sub.Language), sub.c)
		if err != nil {
			s.logger.Error("judge error", zap.Error(err))

			st.Compiled = false
			st.CompilerOutput = "internal error"
			return multierr.Combine(sub.c.Callback("", st, true), err)
		} else {
			return sub.c.Callback("", st, true)
		}
	}

	for {
		w := <-s.workers
		sub := <-s.queue

		if err := judge(w, sub); err != nil {
			s.logger.Error("judging error", zap.Error(err))
		}

		s.workers <- w
	}
}

func (s *Server) getStatus(c echo.Context) error {
	s.statusMutex.RLock()
	defer s.statusMutex.RUnlock()
	return c.JSON(http.StatusOK, s.Status)
}

func (s *Server) postUpdateProblems(c echo.Context) error {
	return s.updateProblems()
}

func (s *Server) postJudge(c echo.Context) error {
	sub := Submission{}
	if err := c.Bind(&sub); err != nil {
		log.Print("getJudge error binding:", err)
		return c.String(http.StatusBadRequest, "Parse error")
	}

	sub.done = make(chan bool, 1)
	if sub.Stream {
		c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		c.Response().WriteHeader(http.StatusOK)

		sub.c = NewWriterCallback(c.Response(), func() {
			c.Response().Flush()
		})

		s.queue <- sub
		<-sub.done
		return sub.c.(*WriterCallback).Error()
	} else {
		sub.c = NewHTTPCallback(sub.CallbackUrl)
		s.queue <- sub
		return c.String(http.StatusOK, "queued")
	}
}

func (s *Server) ToString() (string, error) {
	val, err := json.Marshal(s)
	return string(val), err
}
