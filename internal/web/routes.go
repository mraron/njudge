package web

import (
	"strings"

	"github.com/labstack/echo/v4/middleware"
	"github.com/mraron/njudge/internal/web/extmodels"
	"github.com/mraron/njudge/internal/web/helpers/i18n"
	"github.com/mraron/njudge/internal/web/models"

	"github.com/labstack/echo/v4"
	"github.com/mraron/njudge/internal/web/handlers"
	"github.com/mraron/njudge/internal/web/handlers/api"
	"github.com/mraron/njudge/internal/web/handlers/problemset"
	"github.com/mraron/njudge/internal/web/handlers/taskarchive"
	"github.com/mraron/njudge/internal/web/handlers/user"
	"github.com/mraron/njudge/internal/web/handlers/user/profile"
	"github.com/mraron/njudge/internal/web/helpers"
)

func (s *Server) prepareRoutes(e *echo.Echo) {
	e.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		TokenLookup: "form:_csrf",
		Skipper: func(c echo.Context) bool {
			return strings.HasPrefix(c.Request().URL.Path, "/api") || strings.HasPrefix(c.Request().URL.Path, "/admin")
		},
		CookiePath: "/",
	}))
	e.Use(i18n.SetTranslatorMiddleware())
	e.Use(user.SetUserMiddleware(s.Users))
	e.Use(helpers.ClearTemporaryFlashes())

	e.GET("/", handlers.GetHome())
	e.GET("/page/:page", handlers.GetPage(s.PartialsStore))

	e.Static("/static", "static")

	e.GET("/submission/:id", handlers.GetSubmission(s.Submissions)).Name = "getSubmission"
	e.GET("/submission/rejudge/:id", handlers.RejudgeSubmission(s.Submissions), user.RequireLoginMiddleware()).Name = "rejudgeSubmission"
	e.GET("/task_archive", taskarchive.Get(s.Categories, s.ProblemQuery, s.ProblemInfoQuery, s.ProblemStore))

	ps := e.Group("/problemset", problemset.SetNameMiddleware())
	ps.GET("/:name/", problemset.GetProblemList(s.ProblemStore, s.Problems, s.Categories, s.ProblemListQuery, s.ProblemInfoQuery))
	ps.POST("/:name/submit", problemset.PostSubmit(s.SubmitService), user.RequireLoginMiddleware())
	ps.GET("/status/", problemset.GetStatus(s.SubmissionListQuery)).Name = "getProblemsetStatus"

	psProb := ps.Group("/:name/:problem", problemset.RenameProblemMiddleware(s.ProblemStore),
		problemset.SetProblemMiddleware(s.ProblemStore, s.ProblemQuery, s.ProblemInfoQuery))
	psProb.GET("/", problemset.GetProblem()).Name = "getProblemMain"
	psProb.GET("/problem", problemset.GetProblem())
	psProb.GET("/status", problemset.GetProblemStatus(s.SubmissionListQuery, s.ProblemStore))
	psProb.GET("/submit", problemset.GetProblemSubmit())
	psProb.GET("/ranklist", problemset.GetProblemRanklist(s.SubmissionListQuery))

	psProb.POST("/tags", problemset.PostProblemTag(s.TagsService))
	psProb.GET("/delete_tag/:id", problemset.DeleteProblemTag(s.TagsService))

	psProb.GET("/pdf/:language/", problemset.GetProblemPDF())
	psProb.GET("/attachment/:attachment/", problemset.GetProblemAttachment())
	psProb.GET("/:file", problemset.GetProblemFile())

	u := e.Group("/user")

	u.GET("/auth/callback", user.OAuthCallback(s.Users))
	u.GET("/auth", user.BeginOAuth())

	u.GET("/login", user.GetLogin()).Name = "getUserLogin"
	u.POST("/login", user.PostLogin(s.Users))
	u.GET("/logout", user.Logout())
	u.GET("/register", user.GetRegister())
	u.POST("/register", user.Register(s.Server, s.RegisterService, s.MailService))
	u.GET("/activate", user.GetActivateInfo())
	u.GET("/activate/:name/:key", user.Activate(s.Users))

	u.GET("/forgotten_password", user.GetForgottenPassword()).Name = "GetForgottenPassword"
	u.POST("/forgotten_password", user.PostForgottenPassword(s.Server, s.Users, s.MailService))
	u.GET("/forgotten_password_form/:name/:key", user.GetForgottenPasswordForm()).Name = "GetForgottenPasswordForm"
	u.POST("/forgotten_password_form", user.PostForgottenPasswordForm(s.Users)).Name = "PostForgottenPasswordForm"

	pr := u.Group("/profile", profile.SetProfileMiddleware(s.Users))
	pr.GET("/:name/", profile.GetProfile(s.SubmissionListQuery))
	pr.GET("/:name/submissions/", profile.GetSubmissions(s.SubmissionListQuery))

	prs := pr.Group("/:name/settings", user.RequireLoginMiddleware(), profile.PrivateMiddleware())
	prs.GET("/", profile.GetSettings())
	prs.POST("/change_password/", profile.PostSettingsChangePassword(s.Users))
	prs.POST("/misc/", profile.PostSettingsMisc(s.Users))

	if s.DB != nil {
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
}
