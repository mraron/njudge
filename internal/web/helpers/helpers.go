package helpers

import (
	"github.com/labstack/echo/v4"
	"github.com/mraron/njudge/internal/njudge/db/models"
	"github.com/mraron/njudge/internal/web/templates"
	"net/http"
)

func CensorUserPassword(user *models.User) {
	user.Password = "***CENSORED***"
}

func LoginRequired(c echo.Context) error {
	templates.SetFlash(c, "LoginMessage", "A kért oldal megtekintéséhez belépés szükséges!")
	to := ""
	if c.Request().Method == "GET" {
		to = "?next=" + c.Request().URL.Path
	}
	return c.Redirect(http.StatusFound, "/user/login"+to)
}
