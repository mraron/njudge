package profile

import (
	"net/http"
	"net/url"

	"github.com/labstack/echo/v4"
	"github.com/mraron/njudge/internal/njudge"
)

func SetProfileMiddleware(u njudge.Users) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			name, err := url.QueryUnescape(c.Param("name"))
			if err != nil {
				return err
			}

			user, err := u.GetByName(c.Request().Context(), name)
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
