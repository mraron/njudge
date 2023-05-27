package user

import (
	"context"
	"github.com/mraron/njudge/internal/web/models"
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	. "github.com/volatiletech/sqlboiler/v4/queries/qm"
	"golang.org/x/crypto/bcrypt"
)

func GetLogin() echo.HandlerFunc {
	return func(c echo.Context) error {
		if u := c.Get("user").(*models.User); u != nil {
			return c.Render(http.StatusOK, "error.gohtml", "Már be vagy lépve...")
		}

		return c.Render(http.StatusOK, "user/login", nil)
	}
}

func PostLogin(DB *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			u   = &models.User{}
			err error
		)

		if u := c.Get("user").(*models.User); u != nil {
			return c.Render(http.StatusOK, "error.gohtml", "Már be vagy lépve...")
		}

		u, err = models.Users(Where("name=?", c.FormValue("name"))).One(context.TODO(), DB)
		if err != nil {
			return c.Render(http.StatusOK, "user/login", []string{"Hibás felhasználónév és jelszó páros."})
		}

		if err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(c.FormValue("password"))); err != nil {
			return c.Render(http.StatusOK, "user/login", []string{"Hibás felhasználónév és jelszó páros."})
		}

		if u.ActivationKey.Valid {
			return c.Render(http.StatusOK, "user/login", []string{"Hiba: az account nincs aktiválva."})
		}

		storage, _ := session.Get("user", c)
		storage.Values["id"] = u.ID

		if err = storage.Save(c.Request(), c.Response()); err != nil {
			return err
		}

		c.Set("user", u)

		return c.Render(http.StatusOK, "message.gohtml", "Sikeres belépés.")
	}
}

func Logout() echo.HandlerFunc {
	return func(c echo.Context) error {
		if u := c.Get("user").(*models.User); u == nil {
			return c.Render(http.StatusOK, "error.gohtml", "Ahhoz hogy kijelentkezz előbb be kell hogy jelentkezz...")
		}

		storage, _ := session.Get("user", c)
		storage.Options.MaxAge = -1
		storage.Values["id"] = -1

		if err := storage.Save(c.Request(), c.Response()); err != nil {
			return err
		}

		return c.Redirect(http.StatusFound, "/")
	}
}
