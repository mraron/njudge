package web

import (
	"github.com/gorilla/sessions"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
	"github.com/mraron/njudge/utils/problems"
	_ "github.com/mraron/njudge/utils/problems/config/feladat_txt"
	_ "github.com/mraron/njudge/utils/problems/config/polygon"
	_ "github.com/mraron/njudge/utils/problems/config/task_yaml"
	"github.com/mraron/njudge/web/helpers"
	"github.com/mraron/njudge/web/helpers/config"
	"github.com/mraron/njudge/web/helpers/roles"
	"github.com/mraron/njudge/web/helpers/templates"

	_ "github.com/mraron/njudge/utils/language/cpp11"
	_ "github.com/mraron/njudge/utils/language/cpp14"
	_ "github.com/mraron/njudge/utils/language/golang"
	_ "github.com/mraron/njudge/utils/language/julia"
	_ "github.com/mraron/njudge/utils/language/nim"
	_ "github.com/mraron/njudge/utils/language/octave"
	_ "github.com/mraron/njudge/utils/language/pascal"
	_ "github.com/mraron/njudge/utils/language/python3"
	_ "github.com/mraron/njudge/utils/language/zip"

	"github.com/mraron/njudge/web/models"
	_ "github.com/mraron/njudge/web/models"
	_ "mime"
	"net/http"
)

type Server struct {
	config.Server

	ProblemStore problems.Store

	judges []*models.Judge
	DB     *sqlx.DB
}

func (s *Server) UpdateProblem(pr string) error {
	return s.ProblemStore.UpdateProblem(pr)
}

func (s *Server) GetProblem(pr string) problems.Problem {
	p, _ := s.ProblemStore.Get(pr)
	return p
}

func (s *Server) Run() {
	s.SetupEnvironment()
	s.StartBackgroundJobs()

	e := echo.New()
	store := sessions.NewCookieStore([]byte(s.CookieSecret))

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(session.Middleware(store))
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			user, err := s.currentUser(c)
			if err != nil {
				return helpers.InternalError(c, err, "belső hiba")
			}
			c.Set("user", user)

			return next(c)
		}
	})

	e.Renderer = templates.New(s.TemplatesDir, s.ProblemStore)

	e.GET("/", s.getHome)

	e.Static("/static", "public")

	e.GET("/submission/:id", s.getSubmission)
	e.GET("/submission/rejudge/:id", s.getSubmissionRejudge)
	e.GET("/task_archive", s.getTaskArchive)

	ps := e.Group("/problemset", func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("problemset", c.Param("name"))
			return next(c)
		}
	})

	ps.GET("/:name/", s.getProblemsetList)
	ps.GET("/:name/:problem/", s.getProblemsetProblem)
	ps.GET("/:name/:problem/problem", s.getProblemsetProblem)
	ps.GET("/:name/:problem/status", s.getProblemsetProblemStatus)
	ps.GET("/:name/:problem/ranklist", s.getProblemsetProblemRanklist)
	ps.GET("/:name/:problem/pdf/:language/", s.getProblemsetProblemPDFLanguage)
	ps.GET("/:name/:problem/attachment/:attachment/", s.getProblemsetProblemAttachment)
	ps.GET("/:name/:problem/:file", s.getProblemsetProblemFile)
	ps.POST("/:name/submit", s.postProblemsetSubmit)
	ps.GET("/status", s.getProblemsetStatus)

	u := e.Group("/user")

	u.GET("/auth/callback", s.getUserAuthCallback)
	u.GET("/auth", s.getUserAuth)

	u.GET("/login", s.getUserLogin)
	u.POST("/login", s.postUserLogin)
	u.GET("/logout", s.getUserLogout)
	u.GET("/register", s.getUserRegister)
	u.POST("/register", s.postUserRegister)
	u.GET("/activate", s.getUserActivate)
	u.GET("/activate/:name/:key", s.getActivateUser)

	u.GET("/profile/:name/", s.getUserProfile)

	v1 := e.Group("/api/v1")

	v1.GET("/problem_rels", s.getAPIProblemRels)
	v1.POST("/problem_rels", s.postAPIProblemRel)
	v1.GET("/problem_rels/:id", s.getAPIProblemRel)
	v1.PUT("/problem_rels/:id", s.putAPIProblemRel)
	v1.DELETE("/problem_rels/:id", s.deleteAPIProblemRel)

	v1.GET("/judges", s.getAPIJudges)
	v1.POST("/judges", s.postAPIJudge)
	v1.GET("/judges/:id", s.getAPIJudge)
	v1.PUT("/judges/:id", s.putAPIJudge)
	v1.DELETE("/judges/:id", s.deleteAPIJudge)

	v1.GET("/users", s.getAPIUsers)
	v1.POST("/users", s.postAPIUser)
	v1.GET("/users/:id", s.getAPIUser)
	v1.PUT("/users/:id", s.putAPIUser)
	v1.DELETE("/users/:id", s.deleteAPIUser)

	v1.GET("/submissions", s.getAPISubmissions)
	v1.POST("/submissions", s.postAPISubmission)
	v1.GET("/submissions/:id", s.getAPISubmission)
	v1.PUT("/submissions/:id", s.putAPISubmission)
	v1.DELETE("/submissions/:id", s.deleteAPISubmission)

	e.GET("/admin", s.getAdmin)

	panic(e.Start(":" + s.Port))
}

func (s *Server) getHome(c echo.Context) error {
	return c.Render(http.StatusOK, "home.gohtml", nil)
}

func (s *Server) getAdmin(c echo.Context) error {
	u := c.Get("user").(*models.User)
	if !roles.Can(roles.Role(u.Role), roles.ActionView, "admin_panel") {
		return c.Render(http.StatusUnauthorized, "error.gohtml", "Engedély megtagadva.")
	}

	return c.Render(http.StatusOK, "admin.gohtml", struct {
		Host string
	}{s.Hostname + ":" + s.Port})
}
