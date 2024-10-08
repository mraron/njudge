package handlers

import (
	"github.com/mraron/njudge/internal/web/templates"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mraron/njudge/internal/njudge"
	"github.com/mraron/njudge/internal/web/helpers/roles"
)

func GetHome(store templates.Store) echo.HandlerFunc {
	return func(c echo.Context) error {
		res, _ := store.Get("home")
		return templates.Render(c, http.StatusOK, templates.Home(res))
	}
}

func GetAdmin() echo.HandlerFunc {
	return func(c echo.Context) error {
		u := c.Get(templates.UserContextKey).(*njudge.User)
		if !roles.Can(roles.Role(u.Role), roles.ActionView, "admin_panel") {
			return echo.NotFoundHandler(c)
		}

		return templates.Render(c, http.StatusOK, templates.Admin())
	}
}

func GetPage(store templates.Store) echo.HandlerFunc {
	return func(c echo.Context) error {
		contents, err := store.Get("page_" + c.Param("page"))
		if err != nil {
			return err
		}

		return templates.Render(c, http.StatusOK, templates.PageWithContent(contents))
	}
}
