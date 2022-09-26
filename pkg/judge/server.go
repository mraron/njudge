package judge

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"github.com/mraron/njudge/pkg/language/sandbox"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/shirou/gopsutil/load"
	"go.uber.org/multierr"

	"github.com/mraron/njudge/pkg/problems"
	_ "github.com/mraron/njudge/pkg/problems/config/feladat_txt"
	_ "github.com/mraron/njudge/pkg/problems/config/polygon"
	_ "github.com/mraron/njudge/pkg/problems/tasktype/batch"
	_ "github.com/mraron/njudge/pkg/problems/tasktype/communication"
	_ "github.com/mraron/njudge/pkg/problems/tasktype/stub"

	"io/ioutil"
	"path/filepath"

	"encoding/json"
	"fmt"

	"github.com/kataras/go-errors"
	"github.com/labstack/echo/v4/middleware"
	"github.com/mraron/njudge/pkg/language"
	_ "github.com/mraron/njudge/pkg/language/langs/cpp"
	_ "github.com/mraron/njudge/pkg/language/langs/golang"
	_ "github.com/mraron/njudge/pkg/language/langs/julia"
	_ "github.com/mraron/njudge/pkg/language/langs/nim"
	_ "github.com/mraron/njudge/pkg/language/langs/octave"
	_ "github.com/mraron/njudge/pkg/language/langs/pascal"
	_ "github.com/mraron/njudge/pkg/language/langs/python3"
)

//@TODO use contexts and afero => tests

type ServerStatus struct {
	Id          string        `json:"id" mapstructure:"id"`
	Host        string        `json:"host" mapstructure:"host"`
	Port        string        `json:"port" mapstructure:"port"`
	Url         string        `json:"url"`
	Load        float64       `json:"load"`
	Uptime      time.Duration `json:"uptime"`
	ProblemList []string      `json:"problem_list"`
}

func ParseServerStatus(s string) (res ServerStatus, err error) {
	err = json.Unmarshal([]byte(s), &res)
	return
}

func (s ServerStatus) SupportsProblem(want string) bool {
	for _, key := range s.ProblemList {
		if key == want {
			return true
		}
	}

	return false
}

func (s ServerStatus) String() string {
	res, _ := json.Marshal(s)
	return string(res)
}

type Server struct {
	serverStatusMutex sync.RWMutex
	ServerStatus      `mapstructure:",squash"`

	ProblemsDir string `json:"problems_dir" mapstructure:"problems_dir"`
	LogDir      string `json:"log_dir" mapstructure:"log_dir"`
	SandboxIds  []int  `json:"sandbox_ids" mapstructure:"sandbox_ids"`

	PublicKeyLocation string `json:"public_key" mapstructure:"public_key"`

	ProblemStore problems.Store

	sandboxProvider *language.SandboxProvider
	publicKey       *rsa.PublicKey
	start           time.Time
	queue           chan submission
}

func NewServer() *Server {
	s := Server{}
	s.start = time.Now()
	s.queue = make(chan submission, 100)

	return &s
}

func (s *Server) Run() error {
	s.ProblemStore = problems.NewFsStore(s.ProblemsDir)
	s.updateProblems()

	s.sandboxProvider = language.NewSandboxProvider()
	for _, id := range s.SandboxIds {
		s.sandboxProvider.Put(sandbox.NewIsolate(id))
	}

	if s.PublicKeyLocation != "" {
		publicKeyContents, err := ioutil.ReadFile(s.PublicKeyLocation)
		if err != nil {
			return err
		}

		block, _ := pem.Decode(publicKeyContents)
		if block == nil {
			return fmt.Errorf("can't parse pem public key file: %s", s.PublicKeyLocation)
		}

		s.publicKey, err = x509.ParsePKCS1PublicKey(block.Bytes)
		if err != nil {
			return fmt.Errorf("can't decode publickey: %v", err)
		}
	}

	s.Url = "http://" + s.Host + ":" + s.Port

	e := echo.New()
	e.Use(middleware.Logger())

	if s.PublicKeyLocation != "" {
		e.Use(middleware.JWTWithConfig(middleware.JWTConfig{
			SigningMethod: "RS512",
			ParseTokenFunc: func(auth string, c echo.Context) (interface{}, error) {
				keyFunc := func(t *jwt.Token) (interface{}, error) {
					if t.Method.Alg() != "RS512" {
						return nil, fmt.Errorf("unexpected jwt signing method=%v", t.Header["alg"])
					}
					return s.publicKey, nil
				}

				// claims are of type `jwt.MapClaims` when token is created with `jwt.Parse`
				token, err := jwt.Parse(auth, keyFunc)
				if err != nil {
					return nil, err
				}
				if !token.Valid {
					return nil, errors.New("invalid token")
				}
				return token, nil
			},
		}))
	}

	e.GET("/status", s.getStatus)
	e.POST("/update", s.postUpdateProblems)
	e.POST("/judge", s.postJudge)

	go s.runUpdateLoad()
	go s.runUpdateUptime()
	go s.runUpdateProblems()
	go s.runJudger()

	return e.Start(":" + s.Port)
}

func (s *Server) updateProblems() error {
	var (
		err  error
		err2 error
	)

	s.serverStatusMutex.Lock()
	defer s.serverStatusMutex.Unlock()

	err = s.ProblemStore.Update()
	s.ProblemList, err2 = s.ProblemStore.List()

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
			s.serverStatusMutex.Lock()
			s.ServerStatus.Load = l.Load1
			s.serverStatusMutex.Unlock()
		}

		time.Sleep(60 * time.Second)
	}
}

func (s *Server) runUpdateUptime() {
	for {
		s.serverStatusMutex.Lock()
		s.ServerStatus.Uptime = time.Since(s.start)
		s.serverStatusMutex.Unlock()
		time.Sleep(1 * time.Second)
	}
}

func (s *Server) runJudger() {
	judge := func() error {
		sub := <-s.queue

		defer func() {
			sub.done <- true
		}()

		if ok, err := s.ProblemStore.Has(sub.Problem); !ok {
			return err
		}

		f, err := os.Create(filepath.Join(s.LogDir, fmt.Sprintf("judger.%s", sub.Id)))
		if err != nil {
			return multierr.Append(err, f.Close())
		}

		logger := log.New(f, "[judging]", log.Lshortfile)

		p, _ := s.ProblemStore.Get(sub.Problem)
		st, err := Judge(logger, p, sub.Source, language.Get(sub.Language), s.sandboxProvider, sub.c)
		if err != nil {
			st.Compiled = false
			st.CompilerOutput = "internal error"
			return multierr.Combine(sub.c.Callback("", st, true), err, f.Close())
		} else {
			return multierr.Append(sub.c.Callback("", st, true), f.Close())
		}
	}

	for {
		if err := judge(); err != nil {
			log.Print("judger: ", err)
		}
	}
}

func (s *Server) getStatus(c echo.Context) error {
	s.serverStatusMutex.RLock()
	defer s.serverStatusMutex.RUnlock()
	return c.JSON(http.StatusOK, s.ServerStatus)
}

func (s *Server) postUpdateProblems(c echo.Context) error {
	return s.updateProblems()
}

func (s *Server) postJudge(c echo.Context) error {
	sub := submission{}
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
