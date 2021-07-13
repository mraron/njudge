package helpers

import (
	"github.com/labstack/echo/v4"
	"github.com/mraron/njudge/web/models"
	"math/rand"
	"net/http"
	"time"
)

func GenerateActivationKey(length int) string {
	var (
		alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ012345678901234567890123456789"
		ans      = make([]byte, length)
	)

	src := rand.NewSource(time.Now().UnixNano())

	for i := 0; i < length; i++ {
		ans[i] = alphabet[(int(src.Int63()))%len(alphabet)]
	}

	return string(ans)
}

func CensorUserPassword(user *models.User) {
	user.Password = "***CENSORED***"
}

func InternalError(c echo.Context, err error, msg string) error {
	c.Logger().Print("internal error:", err)
	return c.Render(http.StatusInternalServerError, "error.gohtml", msg)
}

func UnauthorizedError(c echo.Context) error {
	return c.String(http.StatusUnauthorized, "unauthorized")
}

