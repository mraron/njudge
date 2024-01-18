package helpers

import (
	"errors"
	"net/http"
	"time"

	"github.com/mraron/njudge/internal/njudge/db/models"
	"github.com/mraron/njudge/internal/web/helpers/config"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

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
