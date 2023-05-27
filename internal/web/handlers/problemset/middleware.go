package problemset

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/mraron/njudge/internal/web/domain/problem"
	"github.com/mraron/njudge/internal/web/services"
	"github.com/mraron/njudge/pkg/problems"
	"golang.org/x/exp/slices"
	"net/http"
	"strings"
)

func SetNameMiddleware() func(echo.HandlerFunc) echo.HandlerFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("problemset", c.Param("name"))
			return next(c)
		}
	}
}

func RenameProblemMiddleware(problemStore problems.Store) func(echo.HandlerFunc) echo.HandlerFunc {
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

func SetProblemMiddleware(pr problem.Repository, problemStatsService services.ProblemStatsService) func(echo.HandlerFunc) echo.HandlerFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			problemset, problemName := c.Param("name"), c.Param("problem")
			p, err := pr.GetByNames(c.Request().Context(), problemset, problemName)
			if err != nil {
				return err
			}

			c.Set("problem", *p)

			stats, err := problemStatsService.GetStatsData(c.Request().Context(), *p, c.Get("userID").(int))
			if err != nil {
				return err
			}
			c.Set("problemStats", *stats)

			return next(c)
		}
	}
}
