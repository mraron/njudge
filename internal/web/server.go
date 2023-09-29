package web

import (
	"context"
	"github.com/mraron/njudge/internal/web/services"
	_ "mime"
	"net/http"
	"time"

	"github.com/antonlindstrom/pgstore"
	"github.com/mraron/njudge/internal/web/helpers/config"
	"github.com/mraron/njudge/internal/web/helpers/templates"
	"github.com/mraron/njudge/internal/web/helpers/templates/partials"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
	"github.com/mraron/njudge/pkg/problems"
	_ "github.com/mraron/njudge/pkg/problems/config/feladat_txt"
	_ "github.com/mraron/njudge/pkg/problems/config/polygon"
	_ "github.com/mraron/njudge/pkg/problems/config/problem_yaml"
	_ "github.com/mraron/njudge/pkg/problems/config/task_yaml"
	_ "github.com/mraron/njudge/pkg/problems/tasktype/batch"
	_ "github.com/mraron/njudge/pkg/problems/tasktype/communication"
	_ "github.com/mraron/njudge/pkg/problems/tasktype/output_only"
	_ "github.com/mraron/njudge/pkg/problems/tasktype/stub"
)

type Server struct {
	config.Server
	DB *sqlx.DB

	ProblemStore problems.Store
	MailService  services.MailService
}

func (s *Server) Run() {
	s.SetupEnvironment()
	s.StartBackgroundJobs()

	e := echo.New()
	if s.Mode == "development" {
		e.Debug = true
	}
	e.HTTPErrorHandler = func(err error, c echo.Context) {
		code := http.StatusInternalServerError
		if he, ok := err.(*echo.HTTPError); ok {
			code = he.Code
		}

		msg := "Hiba történt"
		if s.Mode == "development" {
			msg = err.Error()
		}

		if err := c.JSON(code, struct {
			Message string `json:"message"`
		}{
			Message: msg,
		}); err != nil {
			c.Logger().Error(err)
		}

		c.Logger().Error(err)
	}
	e.Renderer = templates.New(s.Server, s.ProblemStore, s.DB.DB, partials.NewCached(s.DB.DB, 30*time.Second))

	store, err := pgstore.NewPGStoreFromPool(s.DB.DB, []byte(s.CookieSecret))
	if err != nil {
		panic(err)
	}

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(session.MiddlewareWithConfig(session.Config{
		Skipper: nil,
		Store:   store,
	}))

	s.prepareRoutes(e)

	panic(e.Start(":" + s.Port))
}

func (s *Server) Submit(uid int, problemset, problem, language string, source []byte) (int, error) {
	subService := services.NewSQLSubmitService(s.DB.DB, s.ProblemStore)
	sub, err := subService.Submit(context.Background(), services.SubmitRequest{
		UserID:     uid,
		Problemset: problemset,
		Problem:    problem,
		Language:   language,
		Source:     source,
	})

	if err != nil {
		return -1, err
	}
	return sub.ID, nil
}
