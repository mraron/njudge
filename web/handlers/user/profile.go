package user

import (
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/mraron/njudge/web/helpers"
	"github.com/mraron/njudge/web/models"
	. "github.com/volatiletech/sqlboiler/v4/queries/qm"
	"net/http"
	"net/url"
	"strconv"
)

func ProfileMiddleware(DB *sqlx.DB) echo.MiddlewareFunc {
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

func Profile(DB *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.Render(http.StatusOK, "user/profile/main", struct {
			User *models.User
		}{c.Get("profile").(*models.User)})
	}
}

func Submissions(DB *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		u := c.Get("profile").(*models.User)

		page, err := strconv.Atoi(c.QueryParam("page"))
		if err != nil || page <= 0 {
			page = 1
		}

		statusPage, err := helpers.GetStatusPage(DB, page, 20, OrderBy("id DESC"), []QueryMod{Where("user_id = ?", u.ID)}, c.Request().URL.Query())
		if err != nil {
			return err
		}

		return c.Render(http.StatusOK, "user/profile/submissions", struct {
			User *models.User
			StatusPage *helpers.StatusPage
		}{u, statusPage})
	}
}