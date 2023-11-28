package user

import (
	"bytes"
	"errors"
	"net/http"
	"unicode"

	"github.com/mraron/njudge/internal/njudge"
	"github.com/mraron/njudge/internal/njudge/email"
	"github.com/mraron/njudge/internal/web/helpers/config"
	"github.com/mraron/njudge/internal/web/helpers/i18n"

	"github.com/labstack/echo/v4"
)

type RegistrationPageData struct {
	ErrorStrings []string
	Name         string
	Email        string
}

func GetRegister() echo.HandlerFunc {
	return func(c echo.Context) error {
		if u := c.Get("user").(*njudge.User); u != nil {
			return c.Render(http.StatusOK, "error", "Már be vagy lépve...")
		}

		return c.Render(http.StatusOK, "user/register", RegistrationPageData{})
	}
}

func Register(cfg config.Server, registerService njudge.RegisterService, mailService email.Service) echo.HandlerFunc {
	type request struct {
		Name      string `form:"name"`
		Email     string `form:"email"`
		Password  string `form:"password"`
		Password2 string `form:"password2"`
	}
	return func(c echo.Context) error {
		tr := c.Get(i18n.TranslatorContextKey).(i18n.Translator)

		data := request{}
		if err := c.Bind(&data); err != nil {
			return err
		}

		if u := c.Get("user").(*njudge.User); u != nil {
			return c.Render(http.StatusOK, "error.gohtml", "Már be vagy lépve...")
		}

		register := func() ([]string, error) {
			var (
				errMessages = make([]string, 0)
				err         error
			)

			required := func(value, msg string) {
				if c.FormValue(value) == "" {
					errMessages = append(errMessages, msg)
				}
			}

			alphaNumeric := func(value, msg string) {
				for _, r := range value {
					if !unicode.IsLetter(r) && !unicode.IsDigit(r) {
						errMessages = append(errMessages, msg)
						return
					}
				}
			}

			required("name", tr.Translate("The nickname field is required."))
			required("password", tr.Translate("The password field is required."))
			required("password2", tr.Translate("The password confirmation field is required."))
			required("email", tr.Translate("The email field is required."))

			alphaNumeric(data.Name, tr.Translate("The nickname can only consist of alphanumeric characters: letters (including non-latin characters such as 'á' or 'ű') and digits."))

			if data.Password != data.Password2 {
				errMessages = append(errMessages, tr.Translate("The two passwords don't match."))
			}

			if len(errMessages) > 0 {
				return errMessages, nil
			}

			u, err := registerService.Register(c.Request().Context(), njudge.RegisterRequest{
				Name:     data.Name,
				Email:    data.Email,
				Password: data.Password,
			})

			if errors.Is(err, njudge.ErrorSameName) {
				errMessages = append(errMessages, tr.Translate("The nickname is already registered."))
			}
			if errors.Is(err, njudge.ErrorSameEmail) {
				errMessages = append(errMessages, tr.Translate("The email is already registered."))
			}
			if len(errMessages) > 0 {
				return errMessages, nil
			}
			if err != nil {
				return nil, err
			}

			m := email.Mail{}
			m.Recipients = []string{c.FormValue("email")}
			m.Subject = tr.Translate("Activate your account")

			message := &bytes.Buffer{}
			err = c.Echo().Renderer.Render(message, "mail/activation", struct {
				Name          string
				URL           string
				ActivationKey string
			}{
				c.FormValue("name"),
				cfg.Url,
				u.ActivationInfo.Key,
			}, nil)
			if err != nil {
				return errMessages, err
			}
			m.Message = message.String()

			if err = mailService.Send(c.Request().Context(), m); err != nil {
				return errMessages, err
			}

			return nil, nil
		}

		if errMessages, err := register(); err == nil && len(errMessages) > 0 {
			return c.Render(http.StatusOK, "user/register", RegistrationPageData{
				ErrorStrings: errMessages,
				Name:         data.Name,
				Email:        data.Email,
			})
		} else if err != nil {
			return err
		}

		return c.Redirect(http.StatusFound, "/user/activate")
	}
}

func GetActivateInfo() echo.HandlerFunc {
	return func(c echo.Context) error {
		tr := c.Get(i18n.TranslatorContextKey).(i18n.Translator)

		if u := c.Get("user").(*njudge.User); u != nil {
			return c.Render(http.StatusOK, "error.gohtml", tr.Translate(alreadyLoggedInMessage))
		}

		return c.Render(http.StatusOK, "user/activate.gohtml", nil)
	}
}

func Activate(users njudge.Users) echo.HandlerFunc {
	type request struct {
		Name string `param:"name"`
		Key  string `param:"key"`
	}
	return func(c echo.Context) error {
		data := request{}
		if err := c.Bind(&data); err != nil {
			return err
		}

		tr := c.Get(i18n.TranslatorContextKey).(i18n.Translator)

		if u := c.Get("user").(*njudge.User); u != nil {
			return c.Render(http.StatusOK, "error.gohtml", tr.Translate(alreadyLoggedInMessage))
		}

		user, err := users.GetByName(c.Request().Context(), data.Name)
		if err != nil {
			return err
		}

		if user.ActivationInfo.Activated {
			return c.Render(http.StatusOK, "error.gohtml", tr.Translate("This account has already been activated."))
		}

		if user.ActivationInfo.Key != data.Key {
			return c.Render(http.StatusOK, "error.gohtml", tr.Translate("Wrong activation key. Are you sure you've clicked on the right link?"))
		}

		user.Activate()
		if err := users.Update(c.Request().Context(), user, njudge.Fields(njudge.UserFields.ActivationInfo)); err != nil {
			return err
		}

		return c.Render(http.StatusOK, "message.gohtml", tr.Translate("Successful activation. You can login now!"))
	}
}
