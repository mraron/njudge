package web

import (
	"context"
	"errors"
	"github.com/antonlindstrom/pgstore"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/markbates/goth/gothic"
	"github.com/mraron/njudge/internal/njudge/db/models"
	"github.com/mraron/njudge/internal/web/handlers"
	"github.com/mraron/njudge/internal/web/handlers/api"
	"github.com/mraron/njudge/internal/web/handlers/problemset"
	"github.com/mraron/njudge/internal/web/handlers/user/profile"
	"github.com/mraron/njudge/internal/web/templates"
	"github.com/mraron/njudge/internal/web/templates/i18n"
	"github.com/quasoft/memstore"
	slogecho "github.com/samber/slog-echo"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/mraron/njudge/internal/web/handlers/user"
)

func (s *Server) routes(e *echo.Echo) {
	e.GET("/", handlers.GetHome(s.PartialsStore))
	e.GET("/page/:page", handlers.GetPage(s.PartialsStore))
	e.GET("/submissionRowUpdate/:id", func(c echo.Context) error {
		type request struct {
			ID int `param:"id"`
		}
		data := &request{}
		if err := c.Bind(data); err != nil {
			return err
		}
		sub, err := s.Submissions.Get(c.Request().Context(), data.ID)
		if err != nil {
			return err
		}
		return templates.Render(c, http.StatusOK, templates.SubmissionRowUpdate(*sub))
	})
	e.GET("/submissionFeedbackUpdate/:id", func(c echo.Context) error {
		type request struct {
			ID int `param:"id"`
		}
		data := &request{}
		if err := c.Bind(data); err != nil {
			return err
		}
		sub, err := s.Submissions.Get(c.Request().Context(), data.ID)
		if err != nil {
			return err
		}
		return templates.Render(c, http.StatusOK, templates.SubmissionFeedbackUpdate(*sub))
	})

	e.Static("/static", "static")

	e.GET("/submission/:id", handlers.GetSubmission(s.Submissions, s.Problems, s.Problemsets, s.SolvedStatusQuery)).Name = "getSubmission"
	e.GET("/submission/rejudge/:id", handlers.RejudgeSubmission(s.Submissions), user.RequireLoginMiddleware()).Name = "rejudgeSubmission"
	e.GET("/task_archive", handlers.GetTaskArchive(s.TaskArchiveService))

	ps := e.Group("/problemset", problemset.SetMiddleware(s.Problemsets))
	ps.GET("/:name/", problemset.GetProblemList(s.ProblemStore, s.Problems, s.Categories, s.ProblemListQuery, s.ProblemInfoQuery, s.Tags))
	ps.GET("/:name/ranklist/", problemset.GetRanklist(s.ProblemsetRanklistService))
	ps.POST("/:name/submit", problemset.PostSubmit(s.Submissions, s.SubmitService), user.RequireLoginMiddleware())
	e.GET("/problemset/status/", problemset.GetStatus(s.SubmissionListQuery)).Name = "getProblemsetStatus"

	psProb := ps.Group("/:name/:problem", problemset.RenameProblemMiddleware(s.ProblemStore),
		problemset.SetProblemMiddleware(s.ProblemStore, s.ProblemQuery, s.ProblemInfoQuery), problemset.VisibilityMiddleware())
	psProb.GET("/", problemset.GetProblem(s.Tags)).Name = "getProblemMain"
	psProb.GET("/problem", problemset.GetProblem(s.Tags))
	psProb.GET("/edit", problemset.GetProblemEdit(s.Users, s.Categories), user.RequireLoginMiddleware())
	psProb.POST("/edit", problemset.PostProblemEdit(s.Problems, s.Categories), user.RequireLoginMiddleware())
	psProb.GET("/status", problemset.GetProblemStatus(s.SubmissionListQuery, s.ProblemStore))
	psProb.GET("/submit", problemset.GetProblemSubmit())
	psProb.GET("/ranklist", problemset.GetProblemRanklist(s.SubmissionListQuery, s.Users))

	psProb.POST("/tags", problemset.PostProblemTag(s.TagsService))
	psProb.GET("/delete_tag/:id", problemset.DeleteProblemTag(s.TagsService))

	psProb.GET("/pdf/:language/", problemset.GetProblemPDF())
	psProb.GET("/attachment/:attachment/", problemset.GetProblemAttachment())

	u := e.Group("/user")

	u.GET("/auth/callback", user.OAuthCallback(s.Users))
	u.GET("/auth", user.BeginOAuth())

	u.GET("/login", user.GetLogin(s.GoogleAuth.Enabled)).Name = "getUserLogin"
	u.POST("/login", user.PostLogin(s.Users))
	u.GET("/logout", user.Logout())
	u.GET("/register", user.GetRegister())
	u.POST("/register", user.PostRegister(s.Config.Url, s.Users, s.MailService))
	u.GET("/activate/:name/:key", user.Activate(s.Users))

	u.GET("/forgot_password", user.GetForgotPassword()).Name = "GetForgotPassword"
	u.POST("/forgot_password", user.PostForgotPassword(s.Config.Url, s.Users, s.MailService))
	u.GET("/forgot_password_form/:name/:key", user.GetForgotPasswordForm()).Name = "GetForgotPasswordForm"
	u.POST("/forgot_password_form", user.PostForgotPasswordForm(s.Users)).Name = "PostForgoTPasswordForm"

	pr := u.Group("/profile", profile.SetProfileMiddleware(s.Users))
	pr.GET("/:name/", profile.GetProfile(s.SubmissionListQuery, s.Problems, s.ProblemsetRanklistService))
	pr.GET("/:name/submissions/", profile.GetSubmissions(s.SubmissionListQuery))

	prs := pr.Group("/:name/settings", user.RequireLoginMiddleware(), profile.PrivateMiddleware())
	prs.GET("/", profile.GetSettings())
	prs.POST("/change_password/", profile.PostSettingsChangePassword(s.Users))
	prs.POST("/misc/", profile.PostSettingsMisc(s.Users))

	if s.DB != nil {
		v1 := e.Group("/api/v1")

		problemRelDataProvider := api.ProblemRelDataProvider{DB: s.DB}
		v1.GET("/problem_rels", api.GetList[models.ProblemRel](problemRelDataProvider))
		v1.POST("/problem_rels", api.Post[models.ProblemRel](problemRelDataProvider))
		v1.GET("/problem_rels/:id", api.Get[models.ProblemRel](problemRelDataProvider))
		v1.PUT("/problem_rels/:id", api.Put[models.ProblemRel](problemRelDataProvider))
		v1.DELETE("/problem_rels/:id", api.Delete[models.ProblemRel](problemRelDataProvider))

		partialDataProvider := api.PartialDataProvider{DB: s.DB}
		v1.GET("/partials", api.GetList[models.Partial](partialDataProvider))
		v1.POST("/partials", api.Post[models.Partial](partialDataProvider))
		v1.GET("/partials/:name", api.Get[models.Partial](partialDataProvider))
		v1.PUT("/partials/:name", api.Put[models.Partial](partialDataProvider))
		v1.DELETE("/partials/:name", api.Delete[models.Partial](partialDataProvider))

		judgeDataProvider := api.JudgeDataProvider{DB: s.DB}
		v1.GET("/judges", api.GetList[models.Judge](judgeDataProvider))
		v1.POST("/judges", api.Post[models.Judge](judgeDataProvider))
		v1.GET("/judges/:id", api.Get[models.Judge](judgeDataProvider))
		v1.PUT("/judges/:id", api.Put[models.Judge](judgeDataProvider))
		v1.DELETE("/judges/:id", api.Delete[models.Judge](judgeDataProvider))

		userDataProvider := api.UserDataProvider{DB: s.DB}
		v1.GET("/users", api.GetList[models.User](userDataProvider))
		v1.POST("/users", api.Post[models.User](userDataProvider))
		v1.GET("/users/:id", api.Get[models.User](userDataProvider))
		v1.PUT("/users/:id", api.Put[models.User](userDataProvider))
		v1.DELETE("/users/:id", api.Delete[models.User](userDataProvider))

		submissionDataProvider := api.SubmissionDataProvider{DB: s.DB}
		v1.GET("/submissions", api.GetList[models.Submission](submissionDataProvider))
		v1.POST("/submissions", api.Post[models.Submission](submissionDataProvider))
		v1.GET("/submissions/:id", api.Get[models.Submission](submissionDataProvider))
		v1.PUT("/submissions/:id", api.Put[models.Submission](submissionDataProvider))
		v1.DELETE("/submissions/:id", api.Delete[models.Submission](submissionDataProvider))

		e.GET("/admin", handlers.GetAdmin(), user.RequireLoginMiddleware())
	}
}

func (s *Server) SetupEcho(ctx context.Context, e *echo.Echo) {
	if s.Mode == ModeDemo || s.Mode == ModeDebug || s.Mode == ModeDevelopment {
		e.Debug = true
	} else {
		e.HTTPErrorHandler = func(err error, c echo.Context) {
			if c.Response().Committed {
				return
			}
			code := http.StatusInternalServerError
			var he *echo.HTTPError
			if errors.As(err, &he) {
				code = he.Code
			}

			_ = templates.Render(c, code, templates.Error("Hiba történt."))
			c.Logger().Error(err)
		}
	}

	var (
		store sessions.Store
		err   error
	)

	if s.Mode.UsesDB() {
		store, err = pgstore.NewPGStoreFromPool(s.DB, []byte(s.CookieSecret))
		if err != nil {
			panic(err)
		}
	} else {
		store = memstore.NewMemStore(
			[]byte("authkey123"),
			[]byte("enckey12341234567890123456789012"),
		)
	}
	gothic.Store = store

	e.Use(slogecho.New(s.Logger))
	e.Use(i18n.SetTranslatorMiddleware()) // because we need translation on the error page
	e.Use(middleware.Recover())
	e.Use(session.Middleware(store))
	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 5,
	}))
	e.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		ContextKey:  templates.CSRFTokenContextKey,
		TokenLookup: templates.CSRFTokenLookup,
		Skipper: func(c echo.Context) bool {
			return strings.HasPrefix(c.Request().URL.Path, "/api") || strings.HasPrefix(c.Request().URL.Path, "/admin")
		},
		CookiePath: "/",
	}))
	e.Use(user.SetUserMiddleware(s.Users))
	e.Use(templates.MoveFlashesToContextMiddleware())
	e.Use(templates.ClearTemporaryFlashesMiddleware())
	e.Use(templates.Middleware(s.Users, s.Problems, s.ProblemStore, s.PartialsStore))

	s.routes(e)
}
