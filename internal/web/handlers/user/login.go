package user

import (
	"net/http"

	"github.com/mraron/njudge/internal/web/helpers"
	"github.com/mraron/njudge/internal/web/models"

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

		helpers.DeleteFlash(c, "LoginMessage")

		to := "/"
		if val := c.QueryParams().Get("next"); val != "" {
			to = val
		}
		helpers.SetFlash(c, "LoginRedirect", to)

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

		defer helpers.DeleteFlash(c, "LoginRedirect")
		u, err = models.Users(Where("name=?", c.FormValue("name"))).One(c.Request().Context(), DB)

		if err != nil {
			helpers.SetFlash(c, "Login", "Hibás felhasználónév és jelszó páros.")
			return c.Redirect(http.StatusFound, "/user/login")
		}

		if err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(c.FormValue("password"))); err != nil {
			helpers.SetFlash(c, "Login", "Hibás felhasználónév és jelszó páros.")
			return c.Redirect(http.StatusFound, "/user/login")
		}

		if u.ActivationKey.Valid {
			helpers.SetFlash(c, "Login", "Hiba: az account nincs aktiválva.")
			return c.Redirect(http.StatusFound, "/user/login")
		}

		storage, _ := session.Get("user", c)
		storage.Values["id"] = u.ID

		if err = storage.Save(c.Request(), c.Response()); err != nil {
			return err
		}

		c.Set("user", u)

		to := "/"
		if val, ok := helpers.GetFlash(c, "LoginRedirect").(string); ok {
			to = val
		}

		helpers.SetFlash(c, "TopMessage", "Sikeres belépés!")
		return c.Redirect(http.StatusFound, to)
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
