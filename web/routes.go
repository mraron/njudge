package web

import (
	"github.com/labstack/echo/v4"
	"github.com/mraron/njudge/web/extmodels"
	"github.com/mraron/njudge/web/handlers/api"
	"github.com/mraron/njudge/web/handlers/problemset"
	"github.com/mraron/njudge/web/handlers/submission"
	"github.com/mraron/njudge/web/handlers/taskarchive"
	"github.com/mraron/njudge/web/handlers/user"
	"github.com/mraron/njudge/web/models"
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
	ps.GET("/:name/:problem/", problemset.GetProblem(s.DB, s.ProblemStore), problemset.RenameMiddleware(s.ProblemStore))
	ps.GET("/:name/:problem/problem", problemset.GetProblem(s.DB, s.ProblemStore))
	ps.GET("/:name/:problem/status", problemset.GetProblemStatus(s.DB, s.ProblemStore))
	ps.GET("/:name/:problem/submit", problemset.GetProblemSubmit(s.DB, s.ProblemStore))
	ps.GET("/:name/:problem/ranklist", problemset.GetProblemRanklist(s.DB, s.ProblemStore))
	ps.POST("/:name/:problem/tags", problemset.PostProblemTags(s.DB, s.ProblemStore))
	ps.GET("/:name/:problem/delete_tag/:id", problemset.DeleteProblemTags(s.DB, s.ProblemStore))
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

	problemRelDataProvider := api.ProblemRelDataProvider{DB: s.DB.DB}
	v1.GET("/problem_rels", api.GetList[models.ProblemRel](problemRelDataProvider))
	v1.POST("/problem_rels", api.Post[models.ProblemRel](problemRelDataProvider))
	v1.GET("/problem_rels/:id", api.Get[models.ProblemRel](problemRelDataProvider))
	v1.PUT("/problem_rels/:id", api.Put[models.ProblemRel](problemRelDataProvider))
	v1.DELETE("/problem_rels/:id", api.Delete[models.ProblemRel](problemRelDataProvider))

	partialDataProvider := api.PartialDataProvider{DB: s.DB.DB}
	v1.GET("/partials", api.GetList[models.Partial](partialDataProvider))
	v1.POST("/partials", api.Post[models.Partial](partialDataProvider))
	v1.GET("/partials/:name", api.Get[models.Partial](partialDataProvider))
	v1.PUT("/partials/:name", api.Put[models.Partial](partialDataProvider))
	v1.DELETE("/partials/:name", api.Delete[models.Partial](partialDataProvider))

	judgeDataProvider := api.JudgeDataProvider{DB: s.DB.DB}
	v1.GET("/judges", api.GetList[extmodels.Judge](judgeDataProvider))
	v1.POST("/judges", api.Post[extmodels.Judge](judgeDataProvider))
	v1.GET("/judges/:id", api.Get[extmodels.Judge](judgeDataProvider))
	v1.PUT("/judges/:id", api.Put[extmodels.Judge](judgeDataProvider))
	v1.DELETE("/judges/:id", api.Delete[extmodels.Judge](judgeDataProvider))

	userDataProvider := api.UserDataProvider{DB: s.DB.DB}
	v1.GET("/users", api.GetList[models.User](userDataProvider))
	v1.POST("/users", api.Post[models.User](userDataProvider))
	v1.GET("/users/:id", api.Get[models.User](userDataProvider))
	v1.PUT("/users/:id", api.Put[models.User](userDataProvider))
	v1.DELETE("/users/:id", api.Delete[models.User](userDataProvider))

	submissionDataProvider := api.SubmissionDataProvider{DB: s.DB.DB}
	v1.GET("/submissions", api.GetList[models.Submission](submissionDataProvider))
	v1.POST("/submissions", api.Post[models.Submission](submissionDataProvider))
	v1.GET("/submissions/:id", api.Get[models.Submission](submissionDataProvider))
	v1.PUT("/submissions/:id", api.Put[models.Submission](submissionDataProvider))
	v1.DELETE("/submissions/:id", api.Delete[models.Submission](submissionDataProvider))

	e.GET("/admin", s.getAdmin)
}
