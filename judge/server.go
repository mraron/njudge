package judge

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/shirou/gopsutil/load"
	"io"
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

	"bytes"
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

type Submission struct {
	Id          string `json:"id"`
	Problem     string `json:"problem"`
	Language    string `json:"language"`
	Source      []byte `json:"source"`
	Stream      bool   `json:"stream"`
	CallbackUrl string `json:"callback_url"`

	w io.Writer
	done chan bool
}

//@TODO use contexts and afero => tests

type Server struct {
	Id          string
	Host        string
	Port        string
	url string


	ProblemsDir string `json:"problems_dir"`
	LogDir      string `json:"log_dir"`
	ProblemList []string `json:"problem_list"`
	SandboxIds  []int `json:"sandbox_ids"`

	Load        float64
	Uptime      time.Duration

	PublicKeyLocation       string `json:"public_key"`

	sandboxProvider *language.SandboxProvider
	publicKey       *rsa.PublicKey
	problems        map[string]problems.Problem
	start           time.Time
	queue           chan Submission
}

func New() *Server {
	s := Server{}
	s.problems = make(map[string]problems.Problem)
	s.start = time.Now()
	s.queue = make(chan Submission, 100)

	return &s
}

func NewFromUrl(url, token string) (*Server, error) {
	dst := url + "/status"

	client := &http.Client{}

	req, err := http.NewRequest("GET", dst, nil)
	if err != nil {
		return nil, err
	}

	if token != "" {
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	ans := Server{}

	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&ans)
	if err != nil {
		return nil, err
	}

	ans.url = url
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("judger returned: "+resp.Status)
	}

	return &ans, nil
}

func NewFromString(str string) (*Server, error) {
	s := New()
	if err := json.Unmarshal([]byte(str), s); err != nil {
		return nil, err
	}

	s.url = fmt.Sprintf("http://%s:%s", s.Host, s.Port)
	return s, nil
}

//func NewFromCloning(s *Server) *Server {
//	return New(s.Id, s.Host, s.Port, s.ProblemsDir, s.LogDir, s.PublicKeyLocation)
//}

func (s *Server) SupportsProblem(name string) bool {
	for _, p := range s.ProblemList {
		if p == name {
			return true
		}
	}

	return true
}

func (s *Server) Submit(sub Submission, token string) error {
	dst := s.url + "/judge"

	buf := bytes.Buffer{}

	enc := json.NewEncoder(&buf)
	err := enc.Encode(sub)
	if err != nil {
		return err
	}

	client := &http.Client{}

	req, err := http.NewRequest("POST", dst, &buf)
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json")
	if token != "" {
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	//@TODO ? maybe use json
	if string(data) != "queued" {
		return errors.New("Submit: server says: " + string(data))
	}

	return nil
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

			f, err := os.Create(filepath.Join(s.LogDir, fmt.Sprintf("judger.%d", sub.Id)))
			if err != nil {
				log.Print("judger: can't create logfile", err)
				f.Close()
				return
			}

			logger := log.New(f, "[judging]", log.Lshortfile)

			callbacker := NewCombineCallback()
			if sub.CallbackUrl != "" {
				callbacker.Append(NewHTTPCallback(sub.CallbackUrl))
			}

			if sub.Stream {
				callbacker.Append(NewWriterCallback(sub.w))
			}

			err = Judge(logger, s.problems[sub.Problem], sub.Source, language.Get(sub.Language), s.sandboxProvider, callbacker)
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
	return c.JSON(http.StatusOK, s)
}

func (s *Server) postUpdateProblems(c echo.Context) error {
	s.runUpdateProblems()
	return nil
}

func (s *Server) postJudge(c echo.Context) error {
	sub := Submission{}
	if err := c.Bind(&sub); err != nil {
		log.Print("getJudge error binding:", err)
		return c.String(http.StatusBadRequest, "Parse error")
	}

	sub.done = make(chan bool, 1)
	if sub.Stream {
		sub.w = c.Response()
		s.queue <- sub
		<-sub.done
		return c.String(http.StatusOK, "")
	}else {
		s.queue <- sub
		return c.String(http.StatusOK, "queued")
	}
}

func (s Server) ToString() (string, error) {
	val, err := json.Marshal(s)
	return string(val), err
}

