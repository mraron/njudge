package judge

import (
	"github.com/labstack/echo"
	"github.com/shirou/gopsutil/load"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/mraron/njudge/utils/problems"
	_ "github.com/mraron/njudge/utils/problems/config/polygon"
	_ "github.com/mraron/njudge/utils/problems/config/feladat_txt"
	_ "github.com/mraron/njudge/utils/problems/tasktype/batch"
	_ "github.com/mraron/njudge/utils/problems/tasktype/stub"

	"io/ioutil"
	"path/filepath"

	"bytes"
	"encoding/json"
	"fmt"
	"github.com/kataras/go-errors"
	"github.com/labstack/echo/middleware"
	"github.com/mraron/njudge/utils/language"
	_ "github.com/mraron/njudge/utils/language/cpp11"
	_ "github.com/mraron/njudge/utils/language/cpp14"
	_ "github.com/mraron/njudge/utils/language/golang"
	_ "github.com/mraron/njudge/utils/language/julia"
	_ "github.com/mraron/njudge/utils/language/octave"
	_ "github.com/mraron/njudge/utils/language/python3"
)

type Submission struct {
	Id          int    `json:"id"`
	Problem     string `json:"problem"`
	Language    string `json:"language"`
	Source      []byte `json:"source"`
	CallbackUrl string `json:"callback_url"`
}

type Server struct {
	Id          string
	Host        string
	Port        string
	Load        float64
	ProblemsDir string
	LogDir      string
	ProblemList []string
	Uptime      time.Duration

	sandboxProvider   *language.SandboxProvider
	secretKey string
	problems  map[string]problems.Problem
	start     time.Time
	queue     chan Submission
}

func New(id, host, port, problemsDir, logDir, secretKey string) *Server {
	return &Server{Id: id, Host: host, Port: port, Load: 0.0, secretKey: secretKey, LogDir: logDir, ProblemsDir: problemsDir, problems: make(map[string]problems.Problem), ProblemList: make([]string, 0), Uptime: 0 * time.Second, start: time.Now(), queue: make(chan Submission, 100)}
}

func NewFromUrl(url string) (*Server, error) {
	dst := url + "/status"

	resp, err := http.Get(dst)
	if err != nil {
		return nil, err
	}

	ans := Server{}

	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&ans)
	if err != nil {
		return nil, err
	}

	return &ans, nil
}

func NewFromCloning(s *Server) *Server {
	return New(s.Id, s.Host, s.Port, s.ProblemsDir, s.LogDir, s.secretKey)
}

func (s *Server) SupportsProblem(name string) bool {
	for _, p := range s.ProblemList {
		if p == name {
			return true
		}
	}

	return true
}

func (s *Server) Submit(sub Submission) error {
	dst := "http://" + s.Host + ":" + s.Port + "/judge"

	buf := bytes.Buffer{}
	fmt.Println("postin ye", dst)
	enc := json.NewEncoder(&buf)
	err := enc.Encode(sub)
	if err != nil {
		return err
	}

	resp, err := http.Post(dst, "application/json", &buf)
	if err != nil {
		return err
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if string(data) != "Queued" {
		return errors.New("Submit: server says: " + string(data))
	}

	return nil
}

func (s *Server) Run() error {
	s.sandboxProvider = language.NewSandboxProvider()
	s.sandboxProvider.Put(language.NewIsolateSandbox(0))
	s.sandboxProvider.Put(language.NewIsolateSandbox(1))
	s.sandboxProvider.Put(language.NewIsolateSandbox(2))

	e := echo.New()
	e.Use(middleware.Logger())

	e.GET("/status", s.getStatus())
	e.POST("/update", s.postUpdateProblems())
	e.POST("/judge", s.postJudge())

	go s.runUpdateLoad()
	go s.runUpdateUptime()
	go s.runUpdateProblems()
	go s.runJudger()

	return e.Start(":" + s.Port)
}

//@TODO make it a background process
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
		sub := <-s.queue

		if _, ok := s.problems[sub.Problem]; !ok {
			log.Print("judger: can't find problem", sub.Problem)
			continue
		}

		f, err := os.Create(filepath.Join(s.LogDir, fmt.Sprintf("judger.%d", sub.Id)))
		if err != nil {
			log.Print("judger: can't create logfile", err)
			continue
		}

		logger := log.New(f, "[judging]", log.Lshortfile)

		err = Judge(logger, s.problems[sub.Problem], sub.Source, language.Get(sub.Language), s.sandboxProvider, NewHTTPCallback(sub.CallbackUrl))
		if err != nil {
			log.Print("judger: error while running Judge", err)
			continue
		}
	}
}

func (s *Server) getStatus() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, s)
	}
}

func (s *Server) postUpdateProblems() echo.HandlerFunc {
	return func(c echo.Context) error {
		s.runUpdateProblems()
		return nil
	}
}

func (s *Server) postJudge() echo.HandlerFunc {
	return func(c echo.Context) error {
		sub := Submission{}
		if err := c.Bind(&sub); err != nil {
			log.Print("getJudge error binding:", err)
			return c.String(http.StatusBadRequest, "Parse error")
		}

		s.queue <- sub
		fmt.Println("queued")

		return c.String(http.StatusOK, "Queued")
	}
}

func (s Server) ToString() (string, error) {
	val, err := json.Marshal(s)
	return string(val), err
}

func (s *Server) FromString(str string) error {
	return json.Unmarshal([]byte(str), s)
}
