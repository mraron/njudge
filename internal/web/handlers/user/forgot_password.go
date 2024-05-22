package user

import (
	"bytes"
	"errors"
	"github.com/mraron/njudge/internal/web/templates"
	"github.com/mraron/njudge/internal/web/templates/i18n"
	"github.com/mraron/njudge/internal/web/templates/mail"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/mraron/njudge/internal/njudge"
	"github.com/mraron/njudge/internal/njudge/email"
)

func GetForgotPassword() echo.HandlerFunc {
	return func(c echo.Context) error {

		tr := c.Get(i18n.TranslatorContextKey).(i18n.Translator)

		if u := c.Get("user").(*njudge.User); u != nil {
			return templates.Render(c, http.StatusOK, templates.Error(tr.Translate(alreadyLoggedInMessage)))
		}

		templates.DeleteFlash(c, templates.ForgotPasswordEmailMessageContextKey)

		return templates.Render(c, http.StatusOK, templates.ForgotPasswordEmail())
	}
}

func PostForgotPassword(url string, users njudge.Users, mailService email.Service) echo.HandlerFunc {
	type request struct {
		Email string `form:"email"`
	}
	return func(c echo.Context) error {
		tr := c.Get(i18n.TranslatorContextKey).(i18n.Translator)

		if u := c.Get("user").(*njudge.User); u != nil {
			return templates.Render(c, http.StatusOK, templates.Error(tr.Translate(alreadyLoggedInMessage)))
		}

		data := request{}
		if err := c.Bind(&data); err != nil {
			return err
		}

		u, err := users.GetByEmail(c.Request().Context(), data.Email)
		if err != nil && !errors.Is(err, njudge.ErrorUserNotFound) {
			return err
		} else if err == nil { // We must not expose that an email is not registered
			if u.ForgottenPasswordKey == nil || !u.ForgottenPasswordKey.IsValid() {
				u.SetForgottenPasswordKey(njudge.NewForgottenPasswordKey(1 * time.Hour))
				if err := users.Update(c.Request().Context(), u, njudge.Fields(njudge.UserFields.ForgottenPasswordKey)); err != nil {
					return err
				}

				m := email.Mail{}
				m.Recipients = []string{u.Email}
				m.Subject = tr.Translate("Password reset")

				message := &bytes.Buffer{}
				vm := mail.ForgotPasswordViewModel{
					Name: u.Name,
					URL:  url,
					Key:  u.ForgottenPasswordKey.Key,
				}
				if err = vm.Execute(message); err != nil {
					return err
				}
				m.Message = message.String()

				if err := mailService.Send(c.Request().Context(), m); err != nil {
					return err
				}

			}

		}

		templates.SetFlash(c, templates.ForgotPasswordEmailMessageContextKey, tr.Translate("An email with further instructions was sent to the given address (if it's registered in our system)."))

		return c.Redirect(http.StatusFound, c.Echo().Reverse("GetForgotPassword"))
	}
}

type GetForgotPasswordFormRequest struct {
	Name string `param:"name"`
	Key  string `param:"key"`
}

func GetForgotPasswordForm() echo.HandlerFunc {
	return func(c echo.Context) error {
		tr := c.Get(i18n.TranslatorContextKey).(i18n.Translator)

		if u := c.Get("user").(*njudge.User); u != nil {
			return templates.Render(c, http.StatusOK, templates.Error(tr.Translate(alreadyLoggedInMessage)))
		}

		data := GetForgotPasswordFormRequest{}
		if err := c.Bind(&data); err != nil {
			return err
		}

		templates.DeleteFlash(c, templates.ForgotPasswordFormMessageContextKey)

		vm := templates.ForgotPasswordFormViewModel{
			Name: data.Name,
			Key:  data.Key,
		}
		return templates.Render(c, http.StatusOK, templates.ForgotPasswordForm(vm))
	}
}

type PostForgotPasswordFormRequest struct {
	Password1 string `form:"password1"`
	Password2 string `form:"password1"`

	Name string `form:"name"`
	Key  string `form:"key"`
}

func PostForgotPasswordForm(users njudge.Users) echo.HandlerFunc {
	return func(c echo.Context) error {
		tr := c.Get(i18n.TranslatorContextKey).(i18n.Translator)

		if u := c.Get("user").(*njudge.User); u != nil {
			return templates.Render(c, http.StatusOK, templates.Error(tr.Translate(alreadyLoggedInMessage)))
		}

		data := PostForgotPasswordFormRequest{}
		if err := c.Bind(&data); err != nil {
			return err
		}

		u, err := users.GetByName(c.Request().Context(), data.Name)
		if err != nil {
			return err
		}

		if u.ForgottenPasswordKey == nil || u.ForgottenPasswordKey.Key != data.Key || !u.ForgottenPasswordKey.IsValid() {
			templates.SetFlash(c, templates.ForgotPasswordFormMessageContextKey, tr.Translate("Invalid key provided."))
		} else {
			if data.Password1 != data.Password2 {
				templates.SetFlash(c, templates.ForgotPasswordFormMessageContextKey, tr.Translate("The two passwords don't match."))
			} else {
				u.ForgottenPasswordKey = nil
				u.SetPassword(data.Password1)

				if err := users.Update(c.Request().Context(), u, njudge.Fields(njudge.UserFields.ForgottenPasswordKey, njudge.UserFields.Password)); err == nil {
					templates.SetFlash(c, templates.ForgotPasswordFormMessageContextKey, tr.Translate("Password changed succesfully! You can login with your new password."))
				} else {
					return err
				}
			}
		}

		return c.Redirect(http.StatusFound, c.Echo().Reverse("GetForgotPasswordForm", data.Name, data.Key))
	}
}
