package user

import (
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/mraron/njudge/internal/njudge"
	"github.com/mraron/njudge/internal/web/helpers"
)

func currentUser(c echo.Context, us njudge.Users) (*njudge.User, error) {
	var (
		u   *njudge.User
		err error
	)

	storage, err := session.Get("user", c)
	if err != nil {
		return nil, err
	}

	if _, ok := storage.Values["id"]; !ok {
		return nil, nil
	}

	id := storage.Values["id"].(int)
	u, err = us.Get(c.Request().Context(), id)
	return u, err
}

func SetUserMiddleware(us njudge.Users) func(echo.HandlerFunc) echo.HandlerFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			user, err := currentUser(c, us)
			c.Set("user", user)

			if user != nil {
				c.Set("userID", user.ID)
			} else {
				c.Set("userID", 0)
			}

			if err != nil {
				return next(c)
			}

			return next(c)
		}
	}
}

func RequireLoginMiddleware() func(echo.HandlerFunc) echo.HandlerFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if c.Get("user").(*njudge.User) == nil {
				return helpers.LoginRequired(c)
			}

			return next(c)
		}
	}
}
