package judge

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/shirou/gopsutil/load"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/mraron/njudge/utils/problems"
	_ "github.com/mraron/njudge/utils/problems/config/feladat_txt"
	_ "github.com/mraron/njudge/utils/problems/config/polygon"
	_ "github.com/mraron/njudge/utils/problems/tasktype/batch"
	_ "github.com/mraron/njudge/utils/problems/tasktype/communication"
	_ "github.com/mraron/njudge/utils/problems/tasktype/stub"

	"io/ioutil"
	"path/filepath"

	"encoding/json"
	"fmt"
	"github.com/kataras/go-errors"
	"github.com/labstack/echo/v4/middleware"
	"github.com/mraron/njudge/utils/language"
	_ "github.com/mraron/njudge/utils/language/cpp11"
	_ "github.com/mraron/njudge/utils/language/cpp14"
	_ "github.com/mraron/njudge/utils/language/golang"
	_ "github.com/mraron/njudge/utils/language/julia"
	_ "github.com/mraron/njudge/utils/language/nim"
	_ "github.com/mraron/njudge/utils/language/octave"
	_ "github.com/mraron/njudge/utils/language/pascal"
	_ "github.com/mraron/njudge/utils/language/python3"
)


//@TODO use contexts and afero => tests

type ServerStatus struct {
	Id string `json:"id"`
	Host string `json:"host"`
	Port string `json:"port"`
	Url string `json:"url"`
	Load float64 `json:"load"`
	Uptime      time.Duration `json:"uptime"`
	ProblemList []string `json:"problem_list"`
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
	ServerStatus

	ProblemsDir string `json:"problems_dir"`
	LogDir      string `json:"log_dir"`
	SandboxIds  []int `json:"sandbox_ids"`

	PublicKeyLocation       string `json:"public_key"`

	sandboxProvider *language.SandboxProvider
	publicKey       *rsa.PublicKey
	problems        map[string]problems.Problem
	start           time.Time
	queue           chan submission
}

func NewServer() *Server {
	s := Server{}
	s.problems = make(map[string]problems.Problem)
	s.start = time.Now()
	s.queue = make(chan submission, 100)

	return &s
}

func (s *Server) Run() error {
	s.sandboxProvider = language.NewSandboxProvider()
	for _, id := range s.SandboxIds {
		s.sandboxProvider.Put(language.NewIsolateSandbox(id))
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

	s.Url = "http://"+s.Host+":"+s.Port

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

func (s *Server) runUpdateProblems() {
	files, err := ioutil.ReadDir(s.ProblemsDir)
	if err != nil {
		panic(err)
	}

	pList := make([]string, 0)

	for _, f := range files {
		if f.IsDir() {
			path := filepath.Join(s.ProblemsDir, f.Name())
			p, err := problems.Parse(path)
			if err == nil {
				s.problems[p.Name()] = p
				pList = append(pList, p.Name())
			}
		}
	}

	s.ProblemList = pList
}

func (s *Server) runUpdateLoad() {
	for {
		l, err := load.Avg()

		if err != nil {
			log.Print("Error while getting load: ", err)
		} else {
			s.Load = l.Load1
		}

		time.Sleep(60 * time.Second)
	}
}

func (s *Server) runUpdateUptime() {
	for {
		s.Uptime = time.Since(s.start)
		time.Sleep(1 * time.Second)
	}
}

func (s *Server) runJudger() {
	for {
		func() {
			sub := <-s.queue

			defer func() {
				sub.done <- true
			}()

			if _, ok := s.problems[sub.Problem]; !ok {
				log.Print("judger: can't find problem", sub.Problem)
				return
			}

			f, err := os.Create(filepath.Join(s.LogDir, fmt.Sprintf("judger.%s", sub.Id)))
			if err != nil {
				log.Print("judger: can't create logfile", err)
				f.Close()
				return
			}

			logger := log.New(f, "[judging]", log.Lshortfile)

			err = Judge(logger, s.problems[sub.Problem], sub.Source, language.Get(sub.Language), s.sandboxProvider, sub.c)
			if err != nil {
				log.Print("judger: error while running Judge", err)
				f.Close()
				return
			}

			f.Close()
		}()
	}
}

func (s *Server) getStatus(c echo.Context) error {
	return c.JSON(http.StatusOK, s.ServerStatus)
}

func (s *Server) postUpdateProblems(c echo.Context) error {
	s.runUpdateProblems()
	return nil
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
	}else {
		sub.c = NewHTTPCallback(sub.CallbackUrl)
		s.queue <- sub
		return c.String(http.StatusOK, "queued")
	}
}

func (s Server) ToString() (string, error) {
	val, err := json.Marshal(s)
	return string(val), err
}

