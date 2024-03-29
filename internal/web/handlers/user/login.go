package user

import (
	"errors"
	"net/http"

	"github.com/markbates/goth/gothic"
	"github.com/mraron/njudge/internal/njudge"
	"github.com/mraron/njudge/internal/web/helpers/i18n"

	"github.com/mraron/njudge/internal/web/helpers"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

var ErrorLogin = errors.New("login error")

type LoginErrorWithMessage struct {
	TranslatedMessage string
}

func (LoginErrorWithMessage) Error() string {
	return ErrorLogin.Error()
}

func (LoginErrorWithMessage) Is(target error) bool {
	return target == ErrorLogin
}

var alreadyLoggedInMessage = "You're already logged in..."

type Authenticator func(c echo.Context) (*njudge.User, error)

func loginUserHandler(auth Authenticator) echo.HandlerFunc {
	return func(c echo.Context) error {
		tr := c.Get(i18n.TranslatorContextKey).(i18n.Translator)

		if u := c.Get("user").(*njudge.User); u != nil {
			return c.Render(http.StatusOK, "error.gohtml", tr.Translate(alreadyLoggedInMessage))
		}

		user, err := auth(c)
		if err != nil {
			if errors.Is(err, ErrorLogin) {
				helpers.SetFlash(c, "LoginMessage", err.(LoginErrorWithMessage).TranslatedMessage)
				return c.Redirect(http.StatusFound, c.Echo().Reverse("getUserLogin"))
			} else {
				return err
			}
		}
		defer helpers.DeleteFlash(c, "LoginRedirect")

		if !user.ActivationInfo.Activated {
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

		if u := c.Get("user").(*njudge.User); u != nil {
			return c.Render(http.StatusOK, "error.gohtml", tr.Translate(alreadyLoggedInMessage))
		}

		gothic.BeginAuthHandler(c.Response(), c.Request())
		return nil
	}
}

func GetLogin() echo.HandlerFunc {
	return func(c echo.Context) error {
		tr := c.Get(i18n.TranslatorContextKey).(i18n.Translator)

		if u := c.Get("user").(*njudge.User); u != nil {
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

func PostLogin(us njudge.Users) echo.HandlerFunc {
	type request struct {
		Name     string `form:"name"`
		Password string `form:"password"`
	}

	return loginUserHandler(func(c echo.Context) (*njudge.User, error) {
		tr := c.Get(i18n.TranslatorContextKey).(i18n.Translator)

		data := request{}
		if err := c.Bind(&data); err != nil {
			return nil, err
		}

		wrongCredentialsErr := LoginErrorWithMessage{TranslatedMessage: tr.Translate("Wrong credentials.")}

		u, err := us.GetByName(c.Request().Context(), data.Name)
		if err != nil {
			if errors.Is(err, njudge.ErrorUserNotFound) {
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

func OAuthCallback(us njudge.Users) echo.HandlerFunc {
	return loginUserHandler(func(c echo.Context) (*njudge.User, error) {
		tr := c.Get(i18n.TranslatorContextKey).(i18n.Translator)

		user, err := gothic.CompleteUserAuth(c.Response(), c.Request())
		if err != nil {
			return nil, LoginErrorWithMessage{tr.Translate("Invalid authentication token.")}
		}

		return us.GetByEmail(c.Request().Context(), user.Email)
	})
}

func Logout() echo.HandlerFunc {
	return func(c echo.Context) error {
		tr := c.Get(i18n.TranslatorContextKey).(i18n.Translator)
		if u := c.Get("user").(*njudge.User); u == nil {
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
