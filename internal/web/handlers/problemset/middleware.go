package problemset

import (
	"errors"
	"fmt"
	"net/http"
	"slices"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/mraron/njudge/internal/njudge"
	"github.com/mraron/njudge/pkg/problems"
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
			lst, err := problemStore.ListProblems()
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
				if errors.Is(err, njudge.ErrorProblemNotFound) {
					return c.JSON(http.StatusNotFound, err.Error())
				}
				return err
			}
			c.Set("problem", *p)

			info, err := pinfo.GetProblemData(c.Request().Context(), p.ID, c.Get("userID").(int))
			if err != nil {
				return err
			}
			c.Set("problemInfo", *info)

			storedData, err := p.WithStoredData(store)
			if err != nil {
				return err
			}
			c.Set("problemStoredData", storedData)

			return next(c)
		}
	}
}

func VisibilityMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			p := c.Get("problem").(njudge.Problem)
			if !p.Visible {
				u := c.Get("user").(*njudge.User)
				if u == nil || u.Role != "admin" {
					return c.JSON(http.StatusNotFound, njudge.ErrorProblemNotFound.Error())
				}
			}

			return next(c)
		}

	}
}
