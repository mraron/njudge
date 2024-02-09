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

	"github.com/mraron/njudge/pkg/language"
	"github.com/mraron/njudge/pkg/problems"
	_ "github.com/mraron/njudge/pkg/problems/config/feladat_txt"
	_ "github.com/mraron/njudge/pkg/problems/config/polygon"
	_ "github.com/mraron/njudge/pkg/problems/evaluation/batch"
	_ "github.com/mraron/njudge/pkg/problems/evaluation/communication"
	_ "github.com/mraron/njudge/pkg/problems/evaluation/stub"

	"encoding/json"

	"github.com/labstack/echo/v4/middleware"
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
)

type ServerConfig struct {
	HTTPConfig  `mapstructure:",squash"`
	SandboxIds  string `json:"sandbox_ids" mapstructure:"sandbox_ids"`
	WorkerCount int    `json:"worker_count" mapstructure:"worker_count"`
	ProblemsDir string `json:"problems_dir" mapstructure:"problems_dir"`
	Mode        string `json:"mode" mapstructure:"mode"`
}

type Server struct {
	ServerConfig

	problemStore problems.Store
	httpServer   *HTTPServer

	queue  *Queue
	logger *zap.Logger
}

func NewServer(cfg ServerConfig) (*Server, error) {
	s := &Server{ServerConfig: cfg}

	var err error
	if s.Mode == "development" {
		s.logger, err = zap.NewDevelopment()
	} else {
		s.logger, err = zap.NewProduction()
	}

	if err != nil {
		return nil, err
	}

	minSandboxId, maxSandboxId := -1, -1
	if s.SandboxIds == "" {
		minSandboxId = 100
		maxSandboxId = 999
	} else {
		splitted := strings.Split(s.SandboxIds, "-")
		if len(splitted) != 2 {
			return nil, fmt.Errorf("sandbox_ids wrong format")
		}

		var err1, err2 error
		minSandboxId, err1 = strconv.Atoi(splitted[0])
		maxSandboxId, err2 = strconv.Atoi(splitted[1])
		if err1 != nil || err2 != nil {
			return nil, multierr.Combine(err1, err2)
		}
	}

	s.logger.Info("initializing workers")
	wp, err := NewIsolateWorkerProvider(minSandboxId, maxSandboxId, cfg.WorkerCount)
	if err != nil {
		return nil, err
	}

	s.problemStore = problems.NewFsStore(cfg.ProblemsDir)
	if err = s.problemStore.Update(); err != nil {
		s.logger.Info("failed to initialize problems", zap.Error(err))
	}

	ls := language.DefaultStore

	s.logger.Info("initializing the queue")
	s.queue, err = NewQueue(s.logger, s.problemStore, ls, wp)
	if err != nil {
		return nil, err
	}

	s.logger.Info("initializing the http server")
	s.httpServer = NewHTTPServer(s.HTTPConfig, s.queue, s.logger)

	return s, nil
}

func (s *Server) Run() {
	go func() {
		for {
			if err := s.problemStore.Update(); err != nil {
				s.logger.Error("updating problems", zap.Error(err))
			}

			time.Sleep(20 * time.Second)
		}
	}()

	s.logger.Info("starting the queue")
	go s.queue.Run()

	s.logger.Info("starting the http server")
	s.httpServer.Run()
}

type HTTPConfig struct {
	Host string `json:"host" mapstructure:"host"`
	Port string `json:"port" mapstructure:"port"`
}

type HTTPServer struct {
	HTTPConfig
	Enqueuer

	status      ServerStatus
	statusMutex sync.RWMutex

	start  time.Time
	logger *zap.Logger
}

func NewHTTPServer(cfg HTTPConfig, j Enqueuer, logger *zap.Logger) *HTTPServer {
	s := HTTPServer{HTTPConfig: cfg, Enqueuer: j, logger: logger}
	s.start = time.Now()

	return &s
}

func (s *HTTPServer) Run() error {
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
	e.POST("/judge", s.postJudge)

	go s.runUpdate()

	return e.Start(":" + s.Port)
}

func (s *HTTPServer) init() error {
	s.status.Host = s.Host
	s.status.Port = s.Port
	s.status.Url = "http://" + s.Host + ":" + s.Port

	return nil
}

func (s *HTTPServer) runUpdate() {
	go func() {
		for {
			l, err := load.Avg()

			if err != nil {
				log.Print("Error while getting load: ", err)
			} else {
				s.statusMutex.Lock()
				s.status.Load = l.Load1
				s.statusMutex.Unlock()
			}

			time.Sleep(60 * time.Second)
		}
	}()

	go func() {
		for {
			s.statusMutex.Lock()
			s.status.LanguageList, _ = s.Enqueuer.SupportedLanguages()
			s.status.ProblemList, _ = s.Enqueuer.SupportedProblems()
			s.statusMutex.Unlock()

			time.Sleep(20 * time.Second)
		}
	}()

	for {
		s.statusMutex.Lock()
		s.status.Uptime = time.Since(s.start)
		s.statusMutex.Unlock()
		time.Sleep(1 * time.Second)
	}
}

func (s *HTTPServer) getStatus(c echo.Context) error {
	s.statusMutex.RLock()
	defer s.statusMutex.RUnlock()
	return c.JSON(http.StatusOK, s.status)
}

func (s *HTTPServer) postJudge(c echo.Context) error {
	sub := Submission{}
	if err := c.Bind(&sub); err != nil {
		log.Print("getJudge error binding:", err)
		return c.String(http.StatusBadRequest, "Parse error")
	}

	if sub.Stream {
		c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		c.Response().WriteHeader(http.StatusOK)

		callback := NewWriterCallback(c.Response(), func() {
			c.Response().Flush()
		})

		res, err := s.Enqueue(c.Request().Context(), sub)
		if err != nil {
			return err
		}
		for resp := range res {
			if err := callback.Callback(resp); err != nil {
				return err
			}
		}

		return nil
	} else {
		callback := NewHTTPCallback(sub.CallbackUrl)
		res, err := s.Enqueue(context.Background(), sub)
		if err != nil {
			return err
		}
		go func() {
			for resp := range res {
				err := callback.Callback(resp)
				if err != nil {
					s.logger.Error("error calling back", zap.Error(err))
				}
			}
		}()

		return c.String(http.StatusOK, "queued")
	}
}

func (s *HTTPServer) ToString() (string, error) {
	val, err := json.Marshal(s)
	return string(val), err
}
