package user

import (
	"bytes"
	"database/sql"
	"errors"
	"net/http"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/mraron/njudge/internal/njudge"
	"github.com/mraron/njudge/internal/njudge/email"
	"github.com/mraron/njudge/internal/web/helpers"
	"github.com/mraron/njudge/internal/web/helpers/config"
	"github.com/mraron/njudge/internal/web/helpers/i18n"
	"github.com/mraron/njudge/internal/web/models"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"go.uber.org/multierr"
	"golang.org/x/crypto/bcrypt"
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

func PostForgottenPassword(cfg config.Server, DB *sqlx.DB, mailService email.Service) echo.HandlerFunc {
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

		u, err := models.Users(models.UserWhere.Email.EQ(data.Email)).One(c.Request().Context(), DB)
		if err != nil && !errors.Is(err, sql.ErrNoRows) {
			return err
		}
		if err == nil {
			fpkey, err := models.ForgottenPasswordKeys(models.ForgottenPasswordKeyWhere.UserID.EQ(u.ID)).One(c.Request().Context(), DB)
			if err != nil && !errors.Is(err, sql.ErrNoRows) {
				return err
			}
			if err != nil || (err == nil && time.Now().After(fpkey.Valid)) {
				if err == nil {
					if _, err := fpkey.Delete(c.Request().Context(), DB); err != nil {
						return err
					}
				}

				key := models.ForgottenPasswordKey{
					UserID: u.ID,
					Key:    helpers.GenerateActivationKey(32),
					Valid:  time.Now().Add(1 * time.Hour),
				}

				tx, err := DB.BeginTx(c.Request().Context(), nil)
				if err != nil {
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
					key.Key,
				}, nil); err != nil {
					return multierr.Combine(tx.Rollback(), err)
				}
				m.Message = message.String()

				if err := mailService.Send(c.Request().Context(), m); err != nil {
					return multierr.Combine(tx.Rollback(), err)
				}

				if err := key.Insert(c.Request().Context(), tx, boil.Infer()); err != nil {
					return multierr.Combine(tx.Rollback(), err)
				}
				if err := tx.Commit(); err != nil {
					return err
				}
			}
		}

		helpers.SetFlash(c, "ForgottenPasswordMessage", tr.Translate("An email with further instructions was sent to the given address (if it's registered in our system)."))

		return c.Redirect(http.StatusFound, c.Echo().Reverse("GetForgottenPassword"))
	}
}

func GetForgottenPasswordForm(DB *sqlx.DB) echo.HandlerFunc {
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

func PostForgottenPasswordForm(DB *sqlx.DB) echo.HandlerFunc {
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

		u, err := models.Users(models.UserWhere.Name.EQ(data.Name)).One(c.Request().Context(), DB)
		if err != nil {
			return err
		}

		key, err := models.ForgottenPasswordKeys(models.ForgottenPasswordKeyWhere.UserID.EQ(u.ID)).One(c.Request().Context(), DB)
		if err != nil || key.Key != data.Key || key.Valid.Before(time.Now()) {
			helpers.SetFlash(c, "ForgottenPasswordFormMessage", tr.Translate("Invalid key provided."))
		} else {
			if data.Password1 != data.Password2 {
				helpers.SetFlash(c, "ForgottenPasswordFormMessage", tr.Translate("The two passwords don't match."))
			} else {
				password, err := bcrypt.GenerateFromPassword([]byte(data.Password2), bcrypt.DefaultCost)
				if err != nil {
					return err
				}

				tx := func() (ret error) {
					var tx *sql.Tx
					defer func() {
						if res := recover(); res != nil {
							ret = multierr.Combine(err, tx.Rollback())
						} else {
							ret = tx.Commit()
						}
					}()

					tx, err := DB.BeginTx(c.Request().Context(), nil)
					if err != nil {
						panic(err)
					}

					u.Password = string(password)
					if _, err = u.Update(c.Request().Context(), tx, boil.Whitelist(models.UserColumns.Password)); err != nil {
						panic(err)
					}

					if _, err = key.Delete(c.Request().Context(), tx); err != nil {
						panic(err)
					}

					return
				}

				if err := tx(); err == nil {
					helpers.SetFlash(c, "ForgottenPasswordFormMessage", tr.Translate("Password changed succesfully! You can login with your new password."))
				} else {
					return err
				}
			}
		}

		return c.Redirect(http.StatusFound, c.Echo().Reverse("GetForgottenPasswordForm", data.Name, data.Key))
	}
}
