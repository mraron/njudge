package user

import (
	"bytes"
	"errors"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/mraron/njudge/internal/njudge"
	"github.com/mraron/njudge/internal/njudge/email"
	"github.com/mraron/njudge/internal/web/helpers"
	"github.com/mraron/njudge/internal/web/helpers/config"
	"github.com/mraron/njudge/internal/web/helpers/i18n"
)

func GetForgottenPassword() echo.HandlerFunc {
	return func(c echo.Context) error {

		tr := c.Get(i18n.TranslatorContextKey).(i18n.Translator)

		if u := c.Get("user").(*njudge.User); u != nil {
			return c.Render(http.StatusOK, "error.gohtml", tr.Translate(alreadyLoggedInMessage))
		}

		helpers.DeleteFlash(c, "ForgottenPasswordMessage")

		return c.Render(http.StatusOK, "user/forgotten_password", nil)
	}
}

func PostForgottenPassword(cfg config.Server, users njudge.Users, mailService email.Service) echo.HandlerFunc {
	type request struct {
		Email string `form:"email"`
	}
	return func(c echo.Context) error {
		tr := c.Get(i18n.TranslatorContextKey).(i18n.Translator)

		if u := c.Get("user").(*njudge.User); u != nil {
			return c.Render(http.StatusOK, "error.gohtml", tr.Translate(alreadyLoggedInMessage))
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
				if err := users.Update(c.Request().Context(), *u, njudge.Fields(njudge.UserFields.ForgottenPasswordKey)); err != nil {
					return err
				}

				m := email.Mail{}
				m.Recipients = []string{u.Email}
				m.Subject = tr.Translate("Password reset")

				message := &bytes.Buffer{}
				if err := c.Echo().Renderer.Render(message, "mail/forgotten_password", struct {
					Name string
					URL  string
					Key  string
				}{
					u.Name,
					cfg.Url,
					u.ForgottenPasswordKey.Key,
				}, nil); err != nil {
					return err
				}
				m.Message = message.String()

				if err := mailService.Send(c.Request().Context(), m); err != nil {
					return err
				}

			}

		}

		helpers.SetFlash(c, "ForgottenPasswordMessage", tr.Translate("An email with further instructions was sent to the given address (if it's registered in our system)."))

		return c.Redirect(http.StatusFound, c.Echo().Reverse("GetForgottenPassword"))
	}
}

func GetForgottenPasswordForm() echo.HandlerFunc {
	type request struct {
		Name string `param:"name"`
		Key  string `param:"key"`
	}

	return func(c echo.Context) error {
		tr := c.Get(i18n.TranslatorContextKey).(i18n.Translator)

		if u := c.Get("user").(*njudge.User); u != nil {
			return c.Render(http.StatusOK, "error.gohtml", tr.Translate(alreadyLoggedInMessage))
		}

		data := request{}
		if err := c.Bind(&data); err != nil {
			return err
		}

		helpers.DeleteFlash(c, "ForgottenPasswordFormMessage")

		return c.Render(http.StatusOK, "user/forgotten_password_form", struct {
			Name string
			Key  string
		}{data.Name, data.Key})
	}
}

func PostForgottenPasswordForm(users njudge.Users) echo.HandlerFunc {
	type request struct {
		Password1 string `form:"password1"`
		Password2 string `form:"password1"`

		Name string `form:"name"`
		Key  string `form:"key"`
	}
	return func(c echo.Context) error {
		tr := c.Get(i18n.TranslatorContextKey).(i18n.Translator)

		if u := c.Get("user").(*njudge.User); u != nil {
			return c.Render(http.StatusOK, "error.gohtml", tr.Translate(alreadyLoggedInMessage))
		}

		data := request{}
		if err := c.Bind(&data); err != nil {
			return err
		}

		u, err := users.GetByName(c.Request().Context(), data.Name)
		if err != nil {
			return err
		}

		if u.ForgottenPasswordKey == nil || u.ForgottenPasswordKey.Key != data.Key || !u.ForgottenPasswordKey.IsValid() {
			helpers.SetFlash(c, "ForgottenPasswordFormMessage", tr.Translate("Invalid key provided."))
		} else {
			if data.Password1 != data.Password2 {
				helpers.SetFlash(c, "ForgottenPasswordFormMessage", tr.Translate("The two passwords don't match."))
			} else {
				u.ForgottenPasswordKey = nil
				u.SetPassword(data.Password1)

				if err := users.Update(c.Request().Context(), *u, njudge.Fields(njudge.UserFields.ForgottenPasswordKey, njudge.UserFields.Password)); err == nil {
					helpers.SetFlash(c, "ForgottenPasswordFormMessage", tr.Translate("Password changed succesfully! You can login with your new password."))
				} else {
					return err
				}
			}
		}

		return c.Redirect(http.StatusFound, c.Echo().Reverse("GetForgottenPasswordForm", data.Name, data.Key))
	}
}
