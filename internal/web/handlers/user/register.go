package user

import (
	"bytes"
	"errors"
	"github.com/mraron/njudge/internal/web/templates"
	"net/http"
	"unicode"

	"github.com/mraron/njudge/internal/njudge"
	"github.com/mraron/njudge/internal/njudge/email"
	"github.com/mraron/njudge/internal/web/helpers/config"
	"github.com/mraron/njudge/internal/web/helpers/i18n"

	"github.com/labstack/echo/v4"
)

func GetRegister() echo.HandlerFunc {
	return func(c echo.Context) error {
		tr := c.Get(i18n.TranslatorContextKey).(i18n.Translator)

		if u := c.Get("user").(*njudge.User); u != nil {
			return templates.Render(c, http.StatusOK, templates.Error(tr.Translate(alreadyLoggedInMessage)))
		}

		vm := templates.RegisterViewModel{}
		return templates.Render(c, http.StatusOK, templates.Register(vm))
	}
}

type PostRegisterRequest struct {
	Name      string `form:"name"`
	Email     string `form:"email"`
	Password  string `form:"password"`
	Password2 string `form:"password2"`
}

func PostRegister(cfg config.Server, registerService njudge.RegisterService, mailService email.Service) echo.HandlerFunc {
	return func(c echo.Context) error {
		tr := c.Get(i18n.TranslatorContextKey).(i18n.Translator)

		data := PostRegisterRequest{}
		if err := c.Bind(&data); err != nil {
			return err
		}

		if u := c.Get("user").(*njudge.User); u != nil {
			return templates.Render(c, http.StatusOK, templates.Error(tr.Translate(alreadyLoggedInMessage)))
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
			vm := templates.RegisterViewModel{
				ValidationMessages: errMessages,
				TempName:           data.Name,
				TempEmail:          data.Email,
			}
			return templates.Render(c, http.StatusOK, templates.Register(vm))
		} else if err != nil {
			return err
		}

		return templates.Render(c, http.StatusOK, templates.Info(tr.Translate("Thank you for registering! We've sent you an email with further instructions about finishing your registration.")))
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
			return templates.Render(c, http.StatusOK, templates.Error(tr.Translate(alreadyLoggedInMessage)))
		}

		user, err := users.GetByName(c.Request().Context(), data.Name)
		if err != nil {
			return err
		}

		if user.ActivationInfo.Activated {
			return templates.Render(c, http.StatusOK, templates.Error(tr.Translate("This account has already been activated.")))
		}

		if user.ActivationInfo.Key != data.Key {
			return templates.Render(c, http.StatusOK, templates.Error(tr.Translate("Wrong activation key. Are you sure you've clicked on the right link?")))
		}

		user.Activate()
		if err := users.Update(c.Request().Context(), user, njudge.Fields(njudge.UserFields.ActivationInfo)); err != nil {
			return err
		}

		return templates.Render(c, http.StatusOK, templates.Error(tr.Translate("Successful activation. You can login now!")))
	}
}
