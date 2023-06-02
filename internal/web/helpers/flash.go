package helpers

import (
	"encoding/base64"
	"encoding/json"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
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

func ClearTemporaryFlashes() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			DeleteFlash(c, "TopMessage")
			return next(c)
		}
	}
}
