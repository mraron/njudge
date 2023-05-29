package user

import (
	"bytes"
	"context"
	"database/sql"
	"fmt"
	"github.com/mraron/njudge/internal/web/domain/email"
	"github.com/mraron/njudge/internal/web/helpers"
	"github.com/mraron/njudge/internal/web/helpers/config"
	"github.com/mraron/njudge/internal/web/models"
	"github.com/mraron/njudge/internal/web/services"
	"net/http"
	"unicode"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	. "github.com/volatiletech/sqlboiler/v4/queries/qm"
	"golang.org/x/crypto/bcrypt"
)

type RegistrationPageData struct {
	ErrorStrings []string
	Name         string
	Email        string
}

func GetRegister() echo.HandlerFunc {
	return func(c echo.Context) error {
		if u := c.Get("user").(*models.User); u != nil {
			return c.Render(http.StatusOK, "error", "Már be vagy lépve...")
		}

		return c.Render(http.StatusOK, "user/register", RegistrationPageData{})
	}
}

func Register(cfg config.Server, DB *sqlx.DB, mailService services.MailService) echo.HandlerFunc {
	type request struct {
		Name      string `form:"name"`
		Email     string `form:"email"`
		Password  string `form:"password"`
		Password2 string `form:"password2"`
	}
	return func(c echo.Context) error {
		data := request{}
		if err := c.Bind(&data); err != nil {
			return err
		}

		var (
			errStrings = make([]string, 0)
			key        = helpers.GenerateActivationKey(32)
			err        error
		)

		if u := c.Get("user").(*models.User); u != nil {
			return c.Render(http.StatusOK, "error.gohtml", "Már be vagy lépve...")
		}

		used := func(col, value, msg string) {
			if err != nil {
				u := ""
				err = DB.Get(&u, "SELECT name FROM users WHERE "+col+"=$1", value)
				if u != "" {
					errStrings = append(errStrings, msg)
				}
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

		used("name", data.Name, "A név foglalt")
		used("email", data.Email, "Az email cím foglalt")

		required("name", "A név mező szükséges")
		required("password", "A jelszó mező szükséges")
		required("password2", "A jelszó ellenörző mező szükséges")
		required("email", "Az email mező szükséges")

		alphaNumeric(data.Name, "A név csak alfanumerikus karakterekből állhat")

		if data.Password != data.Password2 {
			errStrings = append(errStrings, "A két jelszó nem egyezik meg")
		}

		if err != nil {
			return err
		}

		if len(errStrings) > 0 {
			return c.Render(http.StatusOK, "user/register", RegistrationPageData{
				ErrorStrings: errStrings,
				Name:         data.Name,
				Email:        data.Email,
			})
		}

		mustPanic := func(err error) {
			if err != nil {
				panic(err)
			}
		}

		transaction := func() {
			tx, err := DB.Begin()
			defer func() {
				if p := recover(); p != nil {
					tx.Rollback()

					var ok bool
					if err, ok = p.(error); !ok {
						err = fmt.Errorf("can't cast to error: %v", err)
					}
				}
			}()

			mustPanic(err)

			hashed, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
			mustPanic(err)

			_, err = tx.Exec("INSERT INTO users (name,password,email,activation_key,role) VALUES ($1,$2,$3,$4,$5)", data.Name, hashed, data.Email, key, "user")
			mustPanic(err)

			m := email.Mail{}
			m.Recipients = []string{c.FormValue("email")}
			m.Subject = "Regisztráció aktiválása"

			message := &bytes.Buffer{}
			mustPanic(c.Echo().Renderer.Render(message, "mail/activation", struct {
				Name          string
				URL           string
				ActivationKey string
			}{
				c.FormValue("name"),
				cfg.Url,
				key,
			}, nil))
			m.Message = message.String()

			mustPanic(mailService.Send(c.Request().Context(), m))

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

		return c.Render(http.StatusOK, "user/activate.gohtml", nil)
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

		if user, err = models.Users(Where("name=?", c.Param("name"))).One(context.TODO(), DB); err != nil {
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
