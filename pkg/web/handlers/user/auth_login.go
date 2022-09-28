package user

import (
	"github.com/mraron/njudge/pkg/web/models"
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/markbates/goth/gothic"
	. "github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func AuthCallback(DB *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		if u := c.Get("user").(*models.User); u != nil {
			return c.Render(http.StatusOK, "error.gohtml", "Már be vagy lépve...")
		}

		user, err := gothic.CompleteUserAuth(c.Response(), c.Request())
		if err != nil {
			return c.Render(http.StatusOK, "login.gohtml", []string{"Hiba: érvénytelen token."})
		}

		lst, err := models.Users(Where("email = ?", user.Email)).All(DB)
		if len(lst) == 0 {
			return c.Render(http.StatusOK, "login.gohtml", []string{"Hiba: a felhasználó nincs regisztrálva."})
		}

		if lst[0].ActivationKey.Valid {
			return c.Render(http.StatusOK, "login.gohtml", []string{"Hiba: az account nincs aktiválva."})
		}

		storage, _ := session.Get("user", c)
		storage.Values["id"] = lst[0].ID

		if err = storage.Save(c.Request(), c.Response()); err != nil {
			return err
		}

		c.Set("user", lst[0])

		return c.Render(http.StatusOK, "message.gohtml", "Sikeres belépés.")
	}
}

func Auth() echo.HandlerFunc {
	return func(c echo.Context) error {
		if u := c.Get("user").(*models.User); u != nil {
			return c.Render(http.StatusOK, "error.gohtml", "Már be vagy lépve...")
		}

		gothic.BeginAuthHandler(c.Response(), c.Request())
		return nil
	}
}
