package problemset

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/mraron/njudge/internal/njudge"
	"github.com/mraron/njudge/pkg/problems"
	"golang.org/x/exp/slices"
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

func SetProblemMiddleware(store problems.Store, ps njudge.ProblemQuery, pinfo njudge.ProblemInfoQuery) func(echo.HandlerFunc) echo.HandlerFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			problemset, problemName := c.Param("name"), c.Param("problem")
			p, err := ps.GetProblem(c.Request().Context(), problemset, problemName)
			if err != nil {
				return err
			}
			c.Set("problem", *p)

			info, err := pinfo.GetProblemData(c.Request().Context(), p.ID, c.Get("userID").(int))
			if err != nil {
				return err
			}
			c.Set("problemInfo", *info)

			sdata, err := p.WithStoredData(store)
			if err != nil {
				return err
			}
			c.Set("problemStoredData", sdata)

			return next(c)
		}
	}
}
