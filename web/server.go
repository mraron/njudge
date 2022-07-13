package web

import (
	"sync"

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

	_ "mime"
	"net/http"

	"github.com/mraron/njudge/web/models"
	_ "github.com/mraron/njudge/web/models"
	. "github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type Server struct {
	config.Server
	DB *sqlx.DB

	ProblemStore problems.Store

	judgesMutex sync.RWMutex
	judges      []*models.Judge
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

	store := sessions.NewCookieStore([]byte(s.CookieSecret))

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(session.Middleware(store))
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			currentUser := func(c echo.Context) (*models.User, error) {
				var (
					u   *models.User
					err error
				)

				storage, err := session.Get("user", c)
				if err != nil {
					panic(err)
				}

				if _, ok := storage.Values["id"]; !ok {
					return nil, nil
				}
				u, err = models.Users(Where("id=?", storage.Values["id"])).One(s.DB)
				return u, err
			}

			user, err := currentUser(c)
			if err != nil {
				return err
			}
			c.Set("user", user)

			return next(c)
		}
	})

	e.Renderer = templates.New(s.Server, s.ProblemStore)

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
