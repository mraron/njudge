package templates

import (
	"encoding/base64"
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
)

const (
	TopMessageContextKey                 = "_top_message"
	LoginMessageContextKey               = "_login_message"
	LoginRedirectContextKey              = "_login_redirect"
	ForgotPasswordEmailMessageContextKey = "_forgot_password_email_message"
	ForgotPasswordFormMessageContextKey  = "_forgot_password_form_message"
	ChangePasswordContextKey             = "_change_password"
)

func SetFlash(c echo.Context, name string, value interface{}) {
	val, _ := json.Marshal(value)
	c.SetCookie(&http.Cookie{Name: "flash" + name, Value: base64.URLEncoding.EncodeToString(val), Path: "/"})
}

func DeleteFlash(c echo.Context, name string) {
	c.SetCookie(&http.Cookie{Name: "flash" + name, Path: "/", MaxAge: 0, Expires: time.Unix(1, 0)})
}

func GetFlash(c echo.Context, name string) interface{} {
	val, err := c.Cookie("flash" + name)
	if err != nil {
		return nil
	}

	encodedJson, err := base64.URLEncoding.DecodeString(val.Value)
	if err != nil {
		return nil
	}

	var res interface{}
	switch encodedJson[0] {
	case '[':
		res = []string{}
	default:
		res = string("")
	}

	if err := json.Unmarshal(encodedJson, &res); err != nil {
		return nil
	}

	return res
}

func MoveFlashesToContextMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			for _, cookie := range c.Cookies() {
				if strings.HasPrefix(cookie.Name, "flash") {
					without, _ := strings.CutPrefix(cookie.Name, "flash")
					c.Set(without, GetFlash(c, without))
				}
			}
			return next(c)
		}
	}
}

func ClearTemporaryFlashesMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			DeleteFlash(c, TopMessageContextKey)
			return next(c)
		}
	}
}
