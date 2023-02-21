package problem

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/mraron/njudge/pkg/problems"
	"golang.org/x/exp/slices"
	"net/http"
	"strings"
)

func RenameMiddleware(problemStore problems.Store) func(echo.HandlerFunc) echo.HandlerFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			lst, err := problemStore.List()
			if err != nil {
				return err
			}

			if !slices.Contains(lst, c.Param("problem")) {
				for _, elem := range lst {
					if strings.HasSuffix(elem, "_"+c.Param("problem")) {
						return c.Redirect(http.StatusTemporaryRedirect, fmt.Sprintf("/problemset/%s/%s/", c.Param("name"), elem))
					}
				}
			}

			return next(c)
		}
	}
}
