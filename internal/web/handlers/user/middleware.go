package user

import (
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/mraron/njudge/internal/web/helpers"
	"github.com/mraron/njudge/internal/web/models"
	. "github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func currentUser(c echo.Context, DB *sqlx.DB) (*models.User, error) {
	var (
		u   *models.User
		err error
	)

	storage, err := session.Get("user", c)
	if err != nil {
		return nil, err
	}

	if _, ok := storage.Values["id"]; !ok {
		return nil, nil
	}
	u, err = models.Users(Where("id=?", storage.Values["id"])).One(DB)
	return u, err
}

func SetUserMiddleware(DB *sqlx.DB) func(echo.HandlerFunc) echo.HandlerFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			user, err := currentUser(c, DB)
			c.Set("user", user)

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
			if c.Get("user").(*models.User) == nil {
				return helpers.LoginRequired(c)
			}

			return next(c)
		}
	}
}
