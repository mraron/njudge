package web

import (
	"github.com/labstack/echo/v4"
	"github.com/mraron/njudge/web/handlers/problemset"
	"github.com/mraron/njudge/web/handlers/submission"
	"github.com/mraron/njudge/web/handlers/taskarchive"
	"github.com/mraron/njudge/web/handlers/user"
)

func (s *Server) prepareRoutes(e *echo.Echo) {
	e.GET("/", s.getHome)

	e.Static("/static", "static")

	e.GET("/submission/:id", submission.Get(s.DB))
	e.GET("/submission/rejudge/:id", submission.Rejudge(s.DB))
	e.GET("/task_archive", taskarchive.Get(s.DB, s.ProblemStore))

	ps := e.Group("/problemset", func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("problemset", c.Param("name"))
			return next(c)
		}
	})

	ps.GET("/:name/", problemset.GetList(s.DB, s.ProblemStore))
	ps.GET("/:name/:problem/", problemset.GetProblem(s.DB, s.ProblemStore))
	ps.GET("/:name/:problem/problem", problemset.GetProblem(s.DB, s.ProblemStore))
	ps.GET("/:name/:problem/status", problemset.GetProblemStatus(s.DB))
	ps.GET("/:name/:problem/ranklist", problemset.GetProblemRanklist(s.DB, s.ProblemStore))
	ps.GET("/:name/:problem/pdf/:language/", problemset.GetProblemPDF(s.ProblemStore))
	ps.GET("/:name/:problem/attachment/:attachment/", problemset.GetProblemAttachment(s.ProblemStore))
	ps.GET("/:name/:problem/:file", problemset.GetProblemFile(s.Server, s.ProblemStore))
	ps.POST("/:name/submit", problemset.PostSubmit(s.Server, s.DB, s.ProblemStore))
	ps.GET("/status/", problemset.GetStatus(s.DB))

	u := e.Group("/user")

	u.GET("/auth/callback", user.AuthCallback(s.DB))
	u.GET("/auth", user.Auth())

	u.GET("/login", user.GetLogin())
	u.POST("/login", user.Login(s.DB))
	u.GET("/logout", user.Logout())
	u.GET("/register", user.GetRegister())
	u.POST("/register", user.Register(s.Server, s.DB))
	u.GET("/activate", user.GetActivateInfo())
	u.GET("/activate/:name/:key", user.Activate(s.DB))

	profile := u.Group("/profile", user.ProfileMiddleware(s.DB))
	profile.GET("/:name/", user.Profile(s.DB))
	profile.GET("/:name/submissions/", user.Submissions(s.DB))

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
}