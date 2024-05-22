package helpers

import (
	"github.com/mraron/njudge/internal/njudge/db/models"
)

func CensorUserPassword(user *models.User) {
	user.Password = "***CENSORED***"
}
