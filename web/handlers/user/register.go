package user

import (
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/mraron/njudge/web/helpers"
	"github.com/mraron/njudge/web/helpers/config"
	"github.com/mraron/njudge/web/helpers/mail"
	"github.com/mraron/njudge/web/models"
	. "github.com/volatiletech/sqlboiler/v4/queries/qm"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"unicode"
)

func GetRegister() echo.HandlerFunc {
	return func(c echo.Context) error {
		if u := c.Get("user").(*models.User); u != nil {
			return c.Render(http.StatusOK, "error", "Már be vagy lépve...")
		}

		return c.Render(http.StatusOK, "user/register", nil)
	}
}

func Register(cfg config.Server, DB *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			errStrings = make([]string, 0)
			key             = helpers.GenerateActivationKey(255)
			err    error
			tx     *sql.Tx
		)

		if u := c.Get("user").(*models.User); u != nil {
			return c.Render(http.StatusOK, "error.gohtml", "Már be vagy lépve...")
		}

		used := func(col, value, msg string) {
			u := ""
			if DB.Get(&u, "SELECT name FROM users WHERE "+col+"=$1", value); u != "" {
				errStrings = append(errStrings, msg)
			}
		}

		required := func(value, msg string) {
			if c.FormValue(value) == "" {
				errStrings = append(errStrings, msg)
			}
		}

		alphaNumeric := func(value, msg string) {
			for _, r := range value {
				if !unicode.IsLetter(r) && !unicode.IsDigit(r) {
					errStrings = append(errStrings, msg)
					return
				}
			}
		}

		used("name", c.FormValue("name"), "A név foglalt")
		used("email", c.FormValue("email"), "Az email cím foglalt")

		required("name", "A név mező szükséges")
		required("password", "A jelszó mező szükséges")
		required("password2", "A jelszó ellenörző mező szükséges")
		required("email", "Az email mező szükséges")

		alphaNumeric(c.FormValue("name"), "A név csak alfanumerikus karakterekből állhat")

		if c.FormValue("password") != c.FormValue("password2") {
			errStrings = append(errStrings, "A két jelszó nem egyezik meg")
		}

		if len(errStrings) > 0 {
			return c.Render(http.StatusOK, "user/register", errStrings)
		}

		mustPanic := func(err error) {
			if err != nil {
				panic(err)
			}
		}

		transaction := func() {
			defer func() {
				if p := recover(); p != nil {
					tx.Rollback()

					var ok bool
					if err, ok = p.(error); !ok {
						err = fmt.Errorf("can't cast to error: %v", err)
					}
				}
			}()

			tx, err := DB.Begin()
			mustPanic(err)

			hashed, err := bcrypt.GenerateFromPassword([]byte(c.FormValue("password")), bcrypt.DefaultCost)
			mustPanic(err)

			_, err = tx.Exec("INSERT INTO users (name,password,email,activation_key,role) VALUES ($1,$2,$3,$4,$5)", c.FormValue("name"), hashed, c.FormValue("email"), key, "user")
			mustPanic(err)

			m := mail.Mail{}
			m.Recipients = []string{c.FormValue("email")}
			m.Message = fmt.Sprintf(`Kedves %s!<br> Köszönjük regisztrációd. Aktiváló link: <a href="http://`+cfg.Url+`/user/activate/%s/%s">http://`+cfg.Url+`/user/activate/%s/%s</a>`, c.FormValue("name"), c.FormValue("name"), key, c.FormValue("name"), key)
			m.Subject = "Regisztráció aktiválása"
			mustPanic(m.Send(cfg))

			mustPanic(tx.Commit())
		}

		if transaction(); err != nil {
			return err
		}

		return c.Redirect(http.StatusFound, "/user/activate")
	}
}

func GetActivateInfo() echo.HandlerFunc {
	return func(c echo.Context) error {
		if u := c.Get("user").(*models.User); u != nil {
			return c.Render(http.StatusOK, "error.gohtml", "Már be vagy lépve...")
		}

		return c.Render(http.StatusOK, "activate.gohtml", nil)
	}
}

func Activate(DB *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			user *models.User
			err  error
			tx   *sql.Tx
		)

		if u := c.Get("user").(*models.User); u != nil {
			return c.Render(http.StatusOK, "error.gohtml", "Már be vagy lépve...")
		}

		if user, err = models.Users(Where("name=?", c.Param("name"))).One(DB); err != nil {
			return err
		}

		if !user.ActivationKey.Valid {
			return c.Render(http.StatusOK, "error.gohtml", "Ez a regisztráció már aktív!")
		}

		if user.ActivationKey.String != c.Param("key") {
			return c.Render(http.StatusOK, "error.gohtml", "Hibás aktiválási kulcs. Biztos jó linkre kattintottál?")
		}

		if tx, err = DB.Begin(); err != nil {
			return err
		}

		if _, err = tx.Exec("UPDATE users SET activation_key=NULL WHERE name=$1", c.Param("name")); err != nil {
			return err
		}

		if err = tx.Commit(); err != nil {
			return err
		}

		return c.Render(http.StatusOK, "message.gohtml", "Sikeres aktiválás, mostmár beléphetsz.")
	}
}