package web

import (
	"time"

	"github.com/labstack/echo/v4"
	"github.com/mraron/njudge/internal/web/extmodels"
	"github.com/mraron/njudge/internal/web/handlers"
	"github.com/mraron/njudge/internal/web/handlers/api"
	"github.com/mraron/njudge/internal/web/handlers/problemset"
	"github.com/mraron/njudge/internal/web/handlers/taskarchive"
	"github.com/mraron/njudge/internal/web/handlers/user"
	"github.com/mraron/njudge/internal/web/handlers/user/profile"
	"github.com/mraron/njudge/internal/web/helpers"
	"github.com/mraron/njudge/internal/web/helpers/templates/partials"
	"github.com/mraron/njudge/internal/web/models"
	"github.com/mraron/njudge/internal/web/services"
)

func (s *Server) prepareRoutes(e *echo.Echo) {
	e.Use(user.SetUserMiddleware(s.DB))
	e.Use(helpers.ClearTemporaryFlashes())

	e.GET("/", handlers.GetHome())
	e.GET("/page/:page", handlers.GetPage(partials.NewCached(s.DB.DB, 30*time.Second)))

	e.Static("/static", "static")

	e.GET("/submission/:id", handlers.GetSubmission(services.NewSQLSubmission(s.DB.DB))).Name = "getSubmission"
	e.GET("/submission/rejudge/:id", handlers.RejudgeSubmission(services.NewSQLSubmission(s.DB.DB)), user.RequireLoginMiddleware()).Name = "rejudgeSubmission"
	e.GET("/task_archive", taskarchive.Get(s.DB, s.ProblemStore))

	ps := e.Group("/problemset", problemset.SetNameMiddleware())
	ps.GET("/:name/", problemset.GetProblemList(s.DB, services.NewSQLProblemListService(s.DB.DB, s.ProblemStore, services.NewSQLProblem(s.DB.DB, s.ProblemStore)), services.NewSQLProblem(s.DB.DB, s.ProblemStore), services.NewSQLProblem(s.DB.DB, s.ProblemStore)))
	ps.POST("/:name/submit", problemset.PostSubmit(services.NewSQLSubmitService(s.DB.DB, s.ProblemStore)))
	ps.GET("/status/", problemset.GetStatus(services.NewSQLStatusPageService(s.DB.DB)))

	psProb := ps.Group("/:name/:problem", problemset.RenameProblemMiddleware(s.ProblemStore), problemset.SetProblemMiddleware(services.NewSQLProblem(s.DB.DB, s.ProblemStore), services.NewSQLProblem(s.DB.DB, s.ProblemStore)))
	psProb.GET("/", problemset.GetProblem()).Name = "getProblemMain"
	psProb.GET("/problem", problemset.GetProblem())
	psProb.GET("/status", problemset.GetProblemStatus(services.NewSQLStatusPageService(s.DB.DB)))
	psProb.GET("/submit", problemset.GetProblemSubmit())
	psProb.GET("/ranklist", problemset.GetProblemRanklist(s.DB))
	psProb.POST("/tags", problemset.PostProblemTag(services.NewSQLTagsService(s.DB.DB)))
	psProb.GET("/delete_tag/:id", problemset.DeleteProblemTag(services.NewSQLTagsService(s.DB.DB)))
	psProb.GET("/pdf/:language/", problemset.GetProblemPDF())
	psProb.GET("/attachment/:attachment/", problemset.GetProblemAttachment())
	psProb.GET("/:file", problemset.GetProblemFile())

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
	pr.GET("/:name/submissions/", profile.GetSubmissions(services.NewSQLStatusPageService(s.DB.DB)))

	prs := pr.Group("/:name/settings", user.RequireLoginMiddleware(), profile.PrivateMiddleware())
	prs.GET("/", profile.GetSettings(s.DB))
	prs.POST("/change_password/", profile.PostSettingsChangePassword(s.DB), user.RequireLoginMiddleware(), user.RequireLoginMiddleware())

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

	e.GET("/admin", handlers.GetAdmin(s.Server), user.RequireLoginMiddleware())
}
