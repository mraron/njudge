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
	"github.com/mraron/njudge/web/models"
	_ "mime"
	"strconv"

	"fmt"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/mraron/njudge/web/roles"
)

type Server struct {
	Port         string
	ProblemsDir  string
	TemplatesDir string

	MailAccount         string
	MailServerHost      string
	MailServerPort      string
	MailAccountPassword string

	problems map[string]problems.Problem
	db       *sqlx.DB
}

func New(port string, problemsDir string, templatesDir string, mailServerAccount, mailServerHost, mailServerPort, mailAccountPassword string) *Server {
	return &Server{port, problemsDir, templatesDir, mailServerAccount, mailServerHost, mailServerPort, mailAccountPassword, make(map[string]problems.Problem), nil}
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

func (s *Server) internalError(c echo.Context, err error, msg string) error {
	c.Logger().Print("internal error:", err)
	return c.Render(http.StatusInternalServerError, "error.html", msg)
}

func (s *Server) unauthorizedError(c echo.Context) error {
	return c.String(http.StatusUnauthorized, "unauthorized")
}

func (s *Server) Run() error {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(session.Middleware(sessions.NewCookieStore([]byte("titkosdolog"))))
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			user, err := s.currentUser(c)
			c.Set("user", user)

			if err != nil {
				return s.internalError(c, err, "belső hiba")
			}

			return next(c)
		}
	})

	t := &Template{
		templates: template.Must(template.New("templater").Funcs(s.templatefuncs()).ParseGlob(filepath.Join(s.TemplatesDir, "*.html"))),
	}

	e.Renderer = t

	e.GET("/", s.getHome)

	e.Static("/static", "public")

	ps := e.Group("/problemset")

	ps.GET("/:name/", s.getProblemsetMain)
	ps.GET("/:name/:problem/", s.getProblemsetProblem)
	ps.GET("/:name/:problem/pdf/:language/", s.getProblemsetProblemPDFLanguage)
	ps.GET("/:name/:problem/attachment/:attachment/", s.getProblemsetProblemAttachment)
	ps.GET("/:name/:problem/:file", s.getProblemsetProblemFile)

	u := e.Group("/user")

	u.GET("/login", s.getUserLogin)
	u.POST("/login", s.postUserLogin)
	u.GET("/logout", s.getUserLogout)
	u.GET("/register", s.getUserRegister)
	u.POST("/register", s.postUserRegister)
	u.GET("/activate", s.getUserActivate)
	u.GET("/activate/:name/:key", s.getActivateUser)

	v1 := e.Group("/api/v1")

	v1.GET("/problem_rels", s.getAPIProblemRels)
	v1.POST("/problem_rels", s.postAPIProblemRel)
	v1.GET("/problem_rels/:id", s.getAPIProblemRel)
	v1.PUT("/problem_rels/:id", s.putAPIProblemRel)
	v1.DELETE("/problem_rels/:id", s.deleteAPIProblemRel)

	e.GET("/admin", s.getAdmin)

	s.updateProblems()
	s.connectToDB()
	models.SetDatabase(s.db)

	return e.Start(":" + s.Port)
}

func (s *Server) getHome(c echo.Context) error {
	fmt.Println("főoldal")
	return c.Render(http.StatusOK, "home.html", s.problems)
}

func (s *Server) getAdmin(c echo.Context) error {
	u := c.Get("user").(*models.User)
	if !roles.Can(u.Role, roles.ActionView, "admin_panel") {
		return c.Render(http.StatusUnauthorized, "error.html", "Engedély megtagadva.")
	}

	return c.Render(http.StatusOK, "admin.html", nil)
}

type paginationData struct {
	_page      int
	_perPage   int
	_sortDir   string
	_sortField string
}

func parsePaginationData(c echo.Context) (*paginationData, error) {
	res := &paginationData{}
	var err error

	_page := c.QueryParam("_page")
	_perPage := c.QueryParam("_perPage")

	res._sortDir = c.QueryParam("_sortDir")
	res._sortField = c.QueryParam("_sortField")

	res._page, err = strconv.Atoi(_page)
	if err != nil {
		return nil, err
	}

	res._perPage, err = strconv.Atoi(_perPage)
	if err != nil {
		return nil, err
	}

	return res, nil
}
