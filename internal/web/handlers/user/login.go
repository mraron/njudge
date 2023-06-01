package user

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/markbates/goth/gothic"
	"github.com/mraron/njudge/internal/web/helpers/i18n"

	"github.com/mraron/njudge/internal/web/helpers"
	"github.com/mraron/njudge/internal/web/models"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	. "github.com/volatiletech/sqlboiler/v4/queries/qm"
	"golang.org/x/crypto/bcrypt"
)

var LoginError = errors.New("login error")

type LoginErrorWithMessage struct {
	TranslatedMessage string
}

func (LoginErrorWithMessage) Error() string {
	return LoginError.Error()
}

func (LoginErrorWithMessage) Is(target error) bool {
	return target == LoginError
}

const alreadyLoggedInMessage = "You're already logged in..."

type Authenticator func(c echo.Context) (*models.User, error)

func loginUserHandler(auth Authenticator) echo.HandlerFunc {
	return func(c echo.Context) error {
		tr := c.Get(i18n.TranslatorContextKey).(i18n.Translator)

		if u := c.Get("user").(*models.User); u != nil {
			return c.Render(http.StatusOK, "error.gohtml", tr.Translate(alreadyLoggedInMessage))
		}

		user, err := auth(c)
		if err != nil {
			if errors.Is(err, LoginError) {
				helpers.SetFlash(c, "LoginMessage", err.(LoginErrorWithMessage).TranslatedMessage)
				return c.Redirect(http.StatusFound, c.Echo().Reverse("getUserLogin"))
			} else {
				return err
			}
		}
		defer helpers.DeleteFlash(c, "LoginRedirect")

		if user.ActivationKey.Valid {
			helpers.SetFlash(c, "LoginMessage", tr.Translate("The account is not activated. Check your emails!"))
			return c.Redirect(http.StatusFound, "/user/login")
		}

		storage, _ := session.Get("user", c)
		storage.Values["id"] = user.ID

		if err = storage.Save(c.Request(), c.Response()); err != nil {
			return err
		}

		c.Set("user", user)

		to := "/"
		if val, ok := helpers.GetFlash(c, "LoginRedirect").(string); ok {
			to = val
		}

		helpers.SetFlash(c, "TopMessage", tr.Translate("Successful login!"))
		return c.Redirect(http.StatusFound, to)
	}
}

func BeginOAuth() echo.HandlerFunc {
	return func(c echo.Context) error {
		tr := c.Get(i18n.TranslatorContextKey).(i18n.Translator)

		if u := c.Get("user").(*models.User); u != nil {
			return c.Render(http.StatusOK, "error.gohtml", tr.Translate(alreadyLoggedInMessage))
		}

		gothic.BeginAuthHandler(c.Response(), c.Request())
		return nil
	}
}

func GetLogin() echo.HandlerFunc {
	return func(c echo.Context) error {
		tr := c.Get(i18n.TranslatorContextKey).(i18n.Translator)

		if u := c.Get("user").(*models.User); u != nil {
			return c.Render(http.StatusOK, "error.gohtml", tr.Translate(alreadyLoggedInMessage))
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
	type request struct {
		Name     string `form:"name"`
		Password string `form:"password"`
	}

	return loginUserHandler(func(c echo.Context) (*models.User, error) {
		tr := c.Get(i18n.TranslatorContextKey).(i18n.Translator)

		data := request{}
		if err := c.Bind(&data); err != nil {
			return nil, err
		}

		wrongCredentialsErr := LoginErrorWithMessage{TranslatedMessage: tr.Translate("Wrong credentials.")}

		u, err := models.Users(Where("name=?", data.Name)).One(c.Request().Context(), DB)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return nil, wrongCredentialsErr
			}

			return nil, err
		}

		if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(data.Password)); err != nil {
			return nil, wrongCredentialsErr
		}

		return u, nil
	})
}

func OAuthCallback(DB *sqlx.DB) echo.HandlerFunc {
	return loginUserHandler(func(c echo.Context) (*models.User, error) {
		tr := c.Get(i18n.TranslatorContextKey).(i18n.Translator)

		user, err := gothic.CompleteUserAuth(c.Response(), c.Request())
		if err != nil {
			return nil, LoginErrorWithMessage{tr.Translate("Invalid authentication token.")}
		}

		lst, err := models.Users(Where("email = ?", user.Email)).All(c.Request().Context(), DB)
		if len(lst) == 0 {
			return nil, LoginErrorWithMessage{tr.Translate("Your email is not associated with any registered account.")}
		}

		return lst[0], nil
	})
}

func Logout() echo.HandlerFunc {
	return func(c echo.Context) error {
		tr := c.Get(i18n.TranslatorContextKey).(i18n.Translator)
		if u := c.Get("user").(*models.User); u == nil {
			return c.Render(http.StatusOK, "error.gohtml", tr.Translate("Can't logout if you've not logged in."))
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
