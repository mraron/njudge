package problemset

import (
	"errors"
	"fmt"
	"net/http"
	"slices"
	"strings"

	"github.com/mraron/njudge/internal/web/handlers/user"
	"github.com/mraron/njudge/internal/web/templates"

	"github.com/labstack/echo/v4"
	"github.com/mraron/njudge/internal/njudge"
	"github.com/mraron/njudge/pkg/problems"
)

const (
	ContextKey                  = "problemset"
	ProblemContextKey           = "problem"
	ProblemInfoContextKey       = "problemInfo"
	ProblemStoredDataContextKey = "problemStoredData"
)

func SetMiddleware(ps njudge.Problemsets) func(echo.HandlerFunc) echo.HandlerFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if r, err := ps.GetByName(c.Request().Context(), c.Param("name")); err != nil {
				if errors.Is(err,njudge.ErrorProblemsetNotFound) {
					return echo.NewHTTPError(http.StatusNotFound, err.Error() )
				}
				return err
			} else {
				c.Set(ContextKey, r)
			}
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
			c.Set(ProblemContextKey, *p)

			info, err := pinfo.GetProblemData(c.Request().Context(), p.ID, c.Get(user.IDContextKey).(int))
			if err != nil {
				return err
			}
			c.Set(ProblemInfoContextKey, *info)

			storedData, err := p.WithStoredData(store)
			if err != nil {
				return err
			}
			c.Set(ProblemStoredDataContextKey, storedData)

			return next(c)
		}
	}
}

func VisibilityMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			p := c.Get(ProblemContextKey).(njudge.Problem)
			if !p.Visible {
				u := c.Get(templates.UserContextKey).(*njudge.User)
				if u == nil || u.Role != "admin" {
					return c.JSON(http.StatusNotFound, njudge.ErrorProblemNotFound.Error())
				}
			}

			return next(c)
		}

	}
}
