package web

import (
	"github.com/labstack/echo/v4"
	"github.com/mraron/njudge/internal/web/extmodels"
	"github.com/mraron/njudge/internal/web/handlers/api"
	"github.com/mraron/njudge/internal/web/handlers/problemset"
	"github.com/mraron/njudge/internal/web/handlers/problemset/problem"
	"github.com/mraron/njudge/internal/web/handlers/submission"
	"github.com/mraron/njudge/internal/web/handlers/taskarchive"
	"github.com/mraron/njudge/internal/web/handlers/user"
	"github.com/mraron/njudge/internal/web/handlers/user/profile"
	"github.com/mraron/njudge/internal/web/models"
)

func (s *Server) prepareRoutes(e *echo.Echo) {
	e.Use(user.SetUserMiddleware(s.DB))

	e.GET("/", s.getHome)

	e.Static("/static", "static")

	e.GET("/submission/:id", submission.Get(s.DB))
	e.GET("/submission/rejudge/:id", submission.Rejudge(s.DB))
	e.GET("/task_archive", taskarchive.Get(s.DB, s.ProblemStore))

	ps := e.Group("/problemset", problemset.SetNameMiddleware())
	ps.GET("/:name/", problemset.GetProblemList(s.DB, s.ProblemStore))
	ps.POST("/:name/submit", problemset.PostSubmit(s.Server, s.DB, s.ProblemStore))
	ps.GET("/status/", problemset.GetStatus(s.DB))

	psProb := ps.Group("/:name/:problem", problem.RenameMiddleware(s.ProblemStore), problem.SetProblemMiddleware(s.DB, s.ProblemStore))
	psProb.GET("/", problem.Get(s.DB))
	psProb.GET("/problem", problem.Get(s.DB))
	psProb.GET("/status", problem.GetStatus(s.DB))
	psProb.GET("/submit", problem.GetSubmit(s.DB))
	psProb.GET("/ranklist", problem.GetRanklist(s.DB))
	psProb.POST("/tags", problem.PostTag(s.DB, s.ProblemStore))
	psProb.GET("/delete_tag/:id", problem.DeleteTag(s.DB, s.ProblemStore))
	psProb.GET("/pdf/:language/", problem.GetPDF())
	psProb.GET("/attachment/:attachment/", problem.GetAttachment())
	psProb.GET("/:file", problem.GetFile())

	u := e.Group("/user")

	u.GET("/auth/callback", user.AuthCallback(s.DB))
	u.GET("/auth", user.Auth())

	u.GET("/login", user.GetLogin())
	u.POST("/login", user.PostLogin(s.DB))
	u.GET("/logout", user.Logout())
	u.GET("/register", user.GetRegister())
	u.POST("/register", user.Register(s.Server, s.DB))
	u.GET("/activate", user.GetActivateInfo())
	u.GET("/activate/:name/:key", user.Activate(s.DB))

	pr := u.Group("/profile", profile.SetProfileMiddleware(s.DB))
	pr.GET("/:name/", profile.GetProfile(s.DB))
	pr.GET("/:name/submissions/", profile.GetSubmissions(s.DB))

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
