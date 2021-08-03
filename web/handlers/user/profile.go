package user

import (
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/mraron/njudge/web/models"
	. "github.com/volatiletech/sqlboiler/v4/queries/qm"
	"net/http"
	"net/url"
)

func Profile(DB *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		name, err := url.QueryUnescape(c.Param("name"))
		if err != nil {
			return err
		}

		user, err := models.Users(Where("name = ?", name)).One(DB)
		if err != nil {
			return err
		}

		return c.Render(http.StatusOK, "user/profile/main", struct {
			User *models.User
		}{user})
	}
}
