package problem

import (
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/mraron/njudge/internal/web/models"
	"github.com/mraron/njudge/pkg/problems"
	. "github.com/volatiletech/sqlboiler/v4/queries/qm"
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

func SetProblemMiddleware(DB *sqlx.DB, problemStore problems.Store) func(echo.HandlerFunc) echo.HandlerFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			name, problem := c.Param("name"), c.Param("problem")

			rel, err := models.ProblemRels(Where("problemset=?", name), Where("problem=?", problem)).One(DB)
			if err != nil {
				return err
			}

			if rel == nil {
				return echo.NewHTTPError(http.StatusNotFound, errors.New("no such problem in problemset"))
			}

			prob, err := problemStore.Get(problem)
			if err != nil {
				return err
			}

			c.Set("problem", prob)
			c.Set("problemRel", rel)
			return next(c)
		}
	}
}
