package web

import (
	"github.com/antonlindstrom/pgstore"
	"github.com/mraron/njudge/internal/web/helpers"
	"github.com/mraron/njudge/internal/web/helpers/config"
	"github.com/mraron/njudge/internal/web/helpers/roles"
	"github.com/mraron/njudge/internal/web/helpers/templates"
	"github.com/mraron/njudge/internal/web/models"
	_ "mime"
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
	"github.com/mraron/njudge/pkg/problems"
	_ "github.com/mraron/njudge/pkg/problems/config/feladat_txt"
	_ "github.com/mraron/njudge/pkg/problems/config/polygon"
	_ "github.com/mraron/njudge/pkg/problems/config/problem_json"
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
}

func (s *Server) Run() {
	s.SetupEnvironment()
	s.StartBackgroundJobs()

	e := echo.New()
	if s.Mode == "development" {
		e.Debug = true
	} else {
		e.HTTPErrorHandler = func(err error, c echo.Context) {
			code := http.StatusInternalServerError
			if he, ok := err.(*echo.HTTPError); ok {
				code = he.Code
			}

			if err := c.Render(code, "error.gohtml", "Hiba történt"); err != nil {
				c.Logger().Error(err)
			}

			c.Logger().Error(err)
		}
	}
	e.Renderer = templates.New(s.Server, s.ProblemStore, s.DB.DB)

	store, err := pgstore.NewPGStoreFromPool(s.DB.DB, []byte(s.CookieSecret))
	if err != nil {
		panic(err)
	}

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(session.Middleware(store))

	s.prepareRoutes(e)

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
		Url string
	}{s.Url})
}

func (s *Server) Submit(uid int, problemset, problem, language string, source []byte) (int, error) {
	return helpers.Submit(s.Server, s.DB, s.ProblemStore, uid, problemset, problem, language, source)
}
