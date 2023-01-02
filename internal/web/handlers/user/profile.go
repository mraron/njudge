package user

import (
	"fmt"
	"github.com/mraron/njudge/internal/web/helpers"
	models2 "github.com/mraron/njudge/internal/web/models"
	"net/http"
	"net/url"
	"strconv"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	. "github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func ProfileMiddleware(DB *sqlx.DB) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			name, err := url.QueryUnescape(c.Param("name"))
			if err != nil {
				return err
			}

			user, err := models2.Users(Where("name = ?", name)).One(DB)
			if err != nil {
				return err
			}

			c.Set("profile", user)

			return next(c)
		}
	}
}

type problem struct {
	Problemset string
	Problem    string
}

func Profile(DB *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		u := c.Get("profile").(*models2.User)
		var (
			solved, attempted models2.SubmissionSlice
			err               error
		)

		solved, err = models2.Submissions(Select("max(submissions.id) as id, problemset, problem"), Where("user_id = ?", u.ID), Where("verdict = ?", 0), GroupBy("submissions.problemset, submissions.problem")).All(DB)
		if err != nil {
			return err
		}

		attempted, err = models2.Submissions(Select("max(submissions.id) as id, problemset, problem"), Where("user_id = ?", u.ID), Where("verdict <> ?", 0), Where("not exists(select id from submissions as other where other.user_id = ? and verdict = 0 and other.problem = submissions.problem and other.problemset=submissions.problemset)", u.ID), GroupBy("submissions.problemset, submissions.problem")).All(DB)
		if err != nil {
			return err
		}

		c.Set("title", fmt.Sprintf("%s", u.Name))
		return c.Render(http.StatusOK, "user/profile/main", struct {
			User                       *models2.User
			SolvedProblems             models2.SubmissionSlice
			AttemptedNotSolvedProblems models2.SubmissionSlice
		}{u, solved, attempted})
	}
}

func Submissions(DB *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		u := c.Get("profile").(*models2.User)

		page, err := strconv.Atoi(c.QueryParam("page"))
		if err != nil || page <= 0 {
			page = 1
		}

		statusPage, err := helpers.GetStatusPage(DB, page, 20, OrderBy("id DESC"), []QueryMod{Where("user_id = ?", u.ID)}, c.Request().URL.Query())
		if err != nil {
			return err
		}

		c.Set("title", fmt.Sprintf("%s beküldései", u.Name))
		return c.Render(http.StatusOK, "user/profile/submissions", struct {
			User       *models2.User
			StatusPage *helpers.StatusPage
		}{u, statusPage})
	}
}
