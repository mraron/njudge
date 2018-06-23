package web

import (
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/mraron/njudge/utils/problems"
	_ "github.com/mraron/njudge/web/models"
	"html/template"
	"io/ioutil"
	"net/http"
	"path/filepath"

	_ "github.com/lib/pq"
	_ "github.com/mraron/njudge/utils/problems/polygon"
	_ "mime"
)

type Server struct {
	Port         string
	ProblemsDir  string
	TemplatesDir string

	problems map[string]problems.Problem
	db       *sqlx.DB
}

func New(port string, problemsDir string, templatesDir string) *Server {
	return &Server{port, problemsDir, templatesDir, make(map[string]problems.Problem), nil}
}

func (s *Server) updateProblems() {
	if s.problems == nil {
		s.problems = make(map[string]problems.Problem)
	}

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
}

func (s *Server) connectToDB() {
	var err error
	s.db, err = sqlx.Open("postgres", "postgres://mraron:***REMOVED***@localhost/mraron")

	if err != nil {
		panic(err)
	}
}

func (s *Server) Run() error {
	e := echo.New()
	e.Use(middleware.Logger())

	t := &Template{
		templates: template.Must(template.New("templater").Funcs(s.templatefuncs()).ParseGlob(filepath.Join(s.TemplatesDir, "*.html"))),
	}

	e.Renderer = t

	e.GET("/", s.getHome)

	ps := e.Group("/problemset")

	ps.GET("/:name/", s.getProblemsetMain)
	ps.GET("/:name/:problem/", s.getProblemsetProblem)
	ps.GET("/:name/:problem/pdf/:language/", s.getProblemsetProblemPDFLanguage)
	ps.GET("/:name/:problem/attachment/:attachment/", s.GetProblemsetProblemAttachment)
	ps.GET("/:name/:problem/:file", s.getProblemsetProblemFile)

	s.updateProblems()
	s.connectToDB()

	return e.Start(":" + s.Port)
}

func (s *Server) getHome(c echo.Context) error {
	return c.Render(http.StatusOK, "home.html", s.problems)
}
