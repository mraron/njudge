package helpers

import (
	"errors"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/mraron/njudge/web/helpers/config"
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

func UnauthorizedError(c echo.Context) error {
	return echo.NewHTTPError(http.StatusUnauthorized, errors.New("unauthorized"))
}

func GetJWT(cfg config.Keys) (string, error) {
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

