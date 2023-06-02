package helpers

import (
	"errors"
	"math/rand"
	"net/http"
	"time"

	"github.com/mraron/njudge/internal/web/helpers/config"
	"github.com/mraron/njudge/internal/web/models"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
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

func LoginRequired(c echo.Context) error {
	SetFlash(c, "LoginMessage", "A kért oldal megtekintéséhez belépés szükséges!")
	to := ""
	if c.Request().Method == "GET" {
		to = "?next=" + c.Request().URL.Path
	}
	return c.Redirect(http.StatusFound, "/user/login"+to)
}

func UnauthorizedError(c echo.Context) error {
	return echo.NewHTTPError(http.StatusUnauthorized, errors.New("unauthorized"))
}

func GetJWT(cfg config.Keys) (string, error) {
	if cfg.PrivateKey == nil {
		return "", nil
	}

	claims := &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(10 * time.Minute).Unix(),
		NotBefore: time.Now().Unix(),
		Issuer:    "njudge web",
		IssuedAt:  time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS512, claims)
	jwt, err := token.SignedString(cfg.PrivateKey)
	if err != nil {
		return "", err
	}

	return jwt, nil
}
