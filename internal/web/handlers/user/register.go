package user

import (
	"bytes"
	"fmt"
	"github.com/mraron/njudge/internal/web/domain/email"
	"github.com/mraron/njudge/internal/web/helpers"
	"github.com/mraron/njudge/internal/web/helpers/config"
	"github.com/mraron/njudge/internal/web/helpers/i18n"
	"github.com/mraron/njudge/internal/web/models"
	"github.com/mraron/njudge/internal/web/services"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"net/http"
	"unicode"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
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
		tr := c.Get("translator").(i18n.Translator)

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

		used("name", data.Name, tr.Translate("The nickname is already registered."))
		used("email", data.Email, tr.Translate("The email is already registered."))

		required("name", tr.Translate("The nickname field is required."))
		required("password", tr.Translate("The password field is required."))
		required("password2", tr.Translate("The password confirmation field is required."))
		required("email", tr.Translate("The email field is required."))

		alphaNumeric(data.Name, tr.Translate("The nickname can only consist of alphanumeric characters: letters (including non-latin characters such as 'á' or 'ű') and digits."))

		if data.Password != data.Password2 {
			errStrings = append(errStrings, tr.Translate("The two passwords don't match."))
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
			m.Subject = tr.Translate("Activate your account")

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
		tr := c.Get("translator").(i18n.Translator)

		if u := c.Get("user").(*models.User); u != nil {
			return c.Render(http.StatusOK, "error.gohtml", tr.Translate(alreadyLoggedInMessage))
		}

		return c.Render(http.StatusOK, "user/activate.gohtml", nil)
	}
}

func Activate(DB *sqlx.DB) echo.HandlerFunc {
	type request struct {
		Name string `param:"name"`
		Key  string `param:"key"`
	}
	return func(c echo.Context) error {
		data := request{}
		if err := c.Bind(&data); err != nil {
			return err
		}

		tr := c.Get("translator").(i18n.Translator)

		if u := c.Get("user").(*models.User); u != nil {
			return c.Render(http.StatusOK, "error.gohtml", tr.Translate(alreadyLoggedInMessage))
		}

		user, err := models.Users(models.UserWhere.Name.EQ(data.Name)).One(c.Request().Context(), DB)
		if err != nil {
			return err
		}

		if !user.ActivationKey.Valid {
			return c.Render(http.StatusOK, "error.gohtml", tr.Translate("This account has already been activated."))
		}

		if user.ActivationKey.String != data.Key {
			return c.Render(http.StatusOK, "error.gohtml", tr.Translate("Wrong activation key. Are you sure you've clicked on the right link?"))
		}

		user.ActivationKey.Valid = false
		if _, err := user.Update(c.Request().Context(), DB, boil.Whitelist(models.UserColumns.ActivationKey)); err != nil {
			return err
		}

		return c.Render(http.StatusOK, "message.gohtml", tr.Translate("Successful activation. You can login now!"))
	}
}
