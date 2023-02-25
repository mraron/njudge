package profile

import (
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/mraron/njudge/internal/web/models"
	. "github.com/volatiletech/sqlboiler/v4/queries/qm"
	"net/url"
)

func SetProfileMiddleware(DB *sqlx.DB) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			name, err := url.QueryUnescape(c.Param("name"))
			if err != nil {
				return err
			}

			user, err := models.Users(Where("name = ?", name)).One(DB)
			if err != nil {
				return err
			}

			c.Set("profile", user)

			return next(c)
		}
	}
}