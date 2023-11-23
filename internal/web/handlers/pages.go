package handlers

import (
	"html/template"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mraron/njudge/internal/njudge"
	"github.com/mraron/njudge/internal/web/helpers/config"
	"github.com/mraron/njudge/internal/web/helpers/roles"
	"github.com/mraron/njudge/internal/web/helpers/templates/partials"
)

func GetHome() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.Render(http.StatusOK, "home.gohtml", nil)
	}
}

func GetAdmin(cfg config.Server) echo.HandlerFunc {
	return func(c echo.Context) error {
		u := c.Get("user").(*njudge.User)
		if !roles.Can(roles.Role(u.Role), roles.ActionView, "admin_panel") {
			return c.Render(http.StatusUnauthorized, "error.gohtml", "Engedély megtagadva.")
		}

		return c.Render(http.StatusOK, "admin.gohtml", struct {
			Url string
		}{cfg.Url})
	}
}

func GetPage(store partials.Store) echo.HandlerFunc {
	return func(c echo.Context) error {
		contents, err := store.Get("page_" + c.Param("page"))
		if err != nil {
			return err
		}

		return c.Render(http.StatusOK, "page.gohtml", struct {
			Contents template.HTML
		}{template.HTML(contents)})
	}
}
