package web

import (
	"github.com/mraron/njudge/pkg/language"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/labstack/echo/v4/middleware"
	"github.com/mraron/njudge/internal/web/helpers/i18n"

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
	e.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		TokenLookup: "form:_csrf",
		Skipper: func(c echo.Context) bool {
			return strings.HasPrefix(c.Request().URL.Path, "/api") || strings.HasPrefix(c.Request().URL.Path, "/admin")
		},
		CookiePath: "/",
	}))
	e.Use(i18n.SetTranslatorMiddleware())
	e.Use(user.SetUserMiddleware(s.DB))
	e.Use(helpers.ClearTemporaryFlashes())

	e.GET("/", handlers.GetHome())

	e.GET("/page/:page", handlers.GetPage(partials.NewCached(s.DB.DB, 30*time.Second)))

	e.Static("/static", "static")

	e.GET("/submission/rejudge/:id", handlers.RejudgeSubmission(services.NewSQLSubmission(s.DB.DB)), user.RequireLoginMiddleware()).Name = "rejudgeSubmission"

	ps := e.Group("/problemset", problemset.SetNameMiddleware())
	ps.GET("/:name/", problemset.GetProblemList(s.DB, services.NewSQLProblemListService(s.DB.DB, s.ProblemStore, services.NewSQLProblem(s.DB.DB, s.ProblemStore)), services.NewSQLProblem(s.DB.DB, s.ProblemStore), services.NewSQLProblem(s.DB.DB, s.ProblemStore)))
	//ps.POST("/:name/submit", problemset.PostSubmit(services.NewSQLSubmitService(s.DB.DB, s.ProblemStore)), user.RequireLoginMiddleware())

	psProb := ps.Group("/:name/:problem", problemset.RenameProblemMiddleware(s.ProblemStore), problemset.SetProblemMiddleware(services.NewSQLProblem(s.DB.DB, s.ProblemStore), services.NewSQLProblem(s.DB.DB, s.ProblemStore)))
	//psProb.GET("/", problemset.GetProblem()).Name = "getProblemMain"
	////psProb.GET("/problem", problemset.GetProblem())
	//psProb.GET("/status", problemset.GetProblemStatus(services.NewSQLStatusPageService(s.DB.DB)))
	psProb.GET("/submit", problemset.GetProblemSubmit())
	//psProb.GET("/ranklist", problemset.GetProblemRanklist(s.DB))
	psProb.POST("/tags", problemset.PostProblemTag(services.NewSQLTagsService(s.DB.DB)))
	psProb.GET("/delete_tag/:id", problemset.DeleteProblemTag(services.NewSQLTagsService(s.DB.DB)))
	//psProb.GET("/pdf/:language/", problemset.GetProblemPDF())
	//psProb.GET("/attachment/:attachment/", problemset.GetProblemAttachment())
	psProb.GET("/:file", problemset.GetProblemFile())

	u := e.Group("/user")

	//u.GET("/auth/callback", user.OAuthCallback(s.DB))
	//u.GET("/auth", user.BeginOAuth())
	//u.POST("/login", user.PostLogin(s.DB))
	//u.GET("/logout", user.Logout())

	//u.GET("/login", user.GetLogin()).Name = "getUserLogin"
	u.GET("/register", user.GetRegister())
	u.POST("/register", user.Register(s.Server, s.DB, s.MailService))
	u.GET("/activate", user.GetActivateInfo())
	u.GET("/activate/:name/:key", user.Activate(s.DB))

	u.GET("/forgotten_password", user.GetForgottenPassword()).Name = "GetForgottenPassword"
	u.POST("/forgotten_password", user.PostForgottenPassword(s.Server, s.DB, s.MailService))
	u.GET("/forgotten_password_form/:name/:key", user.GetForgottenPasswordForm(s.DB)).Name = "GetForgottenPasswordForm"
	u.POST("/forgotten_password_form", user.PostForgottenPasswordForm(s.DB)).Name = "PostForgottenPasswordForm"

	pr := u.Group("/profile", profile.SetProfileMiddleware(s.DB))
	//pr.GET("/:name/", profile.GetProfile(s.DB))
	//pr.GET("/:name/submissions/", profile.GetSubmissions(services.NewSQLStatusPageService(s.DB.DB)))

	_ = pr.Group("/:name/settings", user.RequireLoginMiddleware(), profile.PrivateMiddleware())
	//prs.GET("/", profile.GetSettings(s.DB))
	//prs.POST("/change_password/", profile.PostSettingsChangePassword(s.DB))
	//prs.POST("/misc/", profile.PostSettingsMisc(s.DB))

	apiGroup := e.Group("/api")

	{
		v2 := apiGroup.Group("/v2")
		v2.GET("/archive/", taskarchive.Get(s.DB, s.ProblemStore))

		v2.GET("/", func(c echo.Context) error {
			type Post struct {
				Title   string `json:"title"`
				Content string `json:"content"`
				Date    string `json:"date"`
			}
			return c.JSON(http.StatusOK, struct {
				Posts []Post `json:"posts"`
			}{make([]Post, 0)})
		})

		u := v2.Group("/user")

		pr := u.Group("/profile", profile.SetProfileMiddleware(s.DB))
		pr.GET("/:name/", profile.GetProfile(s.DB))
		pr.GET("/:name/submissions/", profile.GetSubmissions(s.DB, s.ProblemStore, services.NewSQLStatusPageService(s.DB.DB)))

		prs := pr.Group("/:name/settings", user.RequireLoginMiddleware(), profile.PrivateMiddleware())
		prs.GET("/", profile.GetSettings(s.DB))
		prs.POST("/change_password/", profile.PostSettingsChangePassword(s.DB))
		prs.POST("/other/", profile.PostSettingsMisc(s.DB))

		u.GET("/auth/", func(c echo.Context) error {

			u := c.Get("user").(*models.User)
			var (
				ud  *profile.UserData
				err error
			)
			if u != nil {
				ud, err = profile.UserDataFromUser(u)
				if err != nil {
					return err
				}
			}

			return c.JSON(http.StatusOK, struct {
				UserData *profile.UserData `json:"userData"`
			}{ud})
		})

		u.GET("/auth/callback", user.OAuthCallback(s.DB))
		u.GET("/auth/google/", user.BeginOAuth())
		u.POST("/auth/login/", user.PostLogin(s.DB))
		u.GET("/auth/logout/", user.Logout())

		ps := v2.Group("/problemset", problemset.SetNameMiddleware())
		ps.GET("/:name/", problemset.GetProblemList(s.DB, services.NewSQLProblemListService(s.DB.DB, s.ProblemStore, services.NewSQLProblem(s.DB.DB, s.ProblemStore)), services.NewSQLProblem(s.DB.DB, s.ProblemStore), services.NewSQLProblem(s.DB.DB, s.ProblemStore)))
		ps.POST("/:name/submit/", problemset.PostSubmit(services.NewSQLSubmitService(s.DB.DB, s.ProblemStore)), user.RequireLoginMiddleware())
		ps.GET("/status/", problemset.GetStatus(s.DB, s.ProblemStore, services.NewSQLStatusPageService(s.DB.DB))).Name = "getProblemsetStatus"

		psProb := ps.Group("/:name/:problem", problemset.SetProblemMiddleware(services.NewSQLProblem(s.DB.DB, s.ProblemStore), services.NewSQLProblem(s.DB.DB, s.ProblemStore)))
		psProb.GET("/", problemset.GetProblem(s.DB))
		psProb.GET("/submissions/", problemset.GetProblemStatus(s.DB, s.ProblemStore, services.NewSQLStatusPageService(s.DB.DB)))
		psProb.GET("/ranklist/", problemset.GetProblemRanklist(s.DB))
		psProb.GET("/pdf/:language/", problemset.GetProblemPDF())
		psProb.GET("/attachment/:attachment/", problemset.GetProblemAttachment())

		v2.GET("/submission/:id/", handlers.GetSubmission(s.DB, s.ProblemStore, services.NewSQLSubmission(s.DB.DB))).Name = "getSubmission"

		data := v2.Group("/data")
		data.GET("/languages/", func(c echo.Context) error {
			type Language struct {
				ID    string `json:"id"`
				Label string `json:"label"`
			}

			var res []Language

			for _, lang := range language.DefaultStore.List() {
				res = append(res, Language{
					ID:    lang.Id(),
					Label: lang.Name(),
				})
			}

			return c.JSON(http.StatusOK, struct {
				Languages []Language `json:"languages"`
			}{res})
		})

		data.GET("/categories/", func(c echo.Context) error {
			type CategoryFilterOption struct {
				Label string `json:"label"`
				Value string `json:"value"`
			}

			var res []CategoryFilterOption

			categories, err := models.ProblemCategories().All(c.Request().Context(), s.DB)
			if err != nil {
				return err
			}

			par := make(map[int]int)
			for ind := range categories {
				if categories[ind].ParentID.Valid {
					par[categories[ind].ID] = categories[ind].ParentID.Int
				}
			}

			categoryNameByID := make(map[int]string)
			for ind := range categories {
				categoryNameByID[categories[ind].ID] = categories[ind].Name
			}

			var getCategoryNameRec func(int) string
			getCategoryNameRec = func(id int) string {
				if _, ok := par[id]; !ok {
					return categoryNameByID[id]
				} else {
					return getCategoryNameRec(par[id]) + " -- " + categoryNameByID[id]
				}
			}

			for ind := range categories {
				curr := CategoryFilterOption{
					Label: getCategoryNameRec(categories[ind].ID),
					Value: strconv.Itoa(categories[ind].ID),
				}

				res = append(res, curr)
			}

			sort.Slice(res, func(i, j int) bool {

				return res[i].Label < res[j].Label
			})

			return c.JSON(http.StatusOK, struct {
				Categories []CategoryFilterOption `json:"categories"`
			}{res})
		})

		data.GET("/tags/", func(c echo.Context) error {
			tags, err := models.Tags().All(c.Request().Context(), s.DB)
			if err != nil {
				return err
			}

			var res []string
			for _, tag := range tags {
				res = append(res, tag.Name)
			}

			return c.JSON(http.StatusOK, struct {
				Tags []string `json:"tags"`
			}{res})
		})
	}

	v1 := apiGroup.Group("/api/v1")

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
