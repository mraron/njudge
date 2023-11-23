package profile

import (
	"net/http"
	"net/url"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/mraron/njudge/internal/njudge"
	"github.com/mraron/njudge/internal/web/models"
	. "github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func SetProfileMiddleware(DB *sqlx.DB) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			name, err := url.QueryUnescape(c.Param("name"))
			if err != nil {
				return err
			}

			user, err := models.Users(Where("name = ?", name)).One(c.Request().Context(), DB)
			if err != nil {
				return err
			}

			c.Set("profile", user)

			return next(c)
		}
	}
}

func PrivateMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			p := c.Get("profile").(*njudge.User)
			u := c.Get("user").(*njudge.User)
			if p.Name != u.Name {
				return c.Redirect(http.StatusFound, "/")
			}

			return next(c)
		}
	}
}
