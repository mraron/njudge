package profile

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/mraron/njudge/internal/web/helpers"
	"github.com/mraron/njudge/internal/web/models"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	. "github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type problem struct {
	Problemset string
	Problem    string
}

func GetProfile(DB *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		u := c.Get("profile").(*models.User)
		var (
			solved, attempted models.SubmissionSlice
			err               error
		)

		solved, err = models.Submissions(Select("max(submissions.id) as id, problemset, problem"), Where("user_id = ?", u.ID), Where("verdict = ?", 0), GroupBy("submissions.problemset, submissions.problem")).All(DB)
		if err != nil {
			return err
		}

		attempted, err = models.Submissions(Select("max(submissions.id) as id, problemset, problem"), Where("user_id = ?", u.ID), Where("verdict <> ?", 0), Where("not exists(select id from submissions as other where other.user_id = ? and verdict = 0 and other.problem = submissions.problem and other.problemset=submissions.problemset)", u.ID), GroupBy("submissions.problemset, submissions.problem")).All(DB)
		if err != nil {
			return err
		}

		c.Set("title", fmt.Sprintf("%s", u.Name))
		return c.Render(http.StatusOK, "user/profile/main", struct {
			User                       *models.User
			SolvedProblems             models.SubmissionSlice
			AttemptedNotSolvedProblems models.SubmissionSlice
		}{u, solved, attempted})
	}
}

func GetSubmissions(DB *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		u := c.Get("profile").(*models.User)

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
			User       *models.User
			StatusPage *helpers.StatusPage
		}{u, statusPage})
	}
}

func GetSettings(DB *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		u := c.Get("user").(*models.User)

		helpers.DeleteFlash(c, "ChangePassword")
		return c.Render(http.StatusOK, "user/profile/settings", struct {
			User *models.User
		}{u})
	}
}

func PostSettingsChangePassword(DB *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.Redirect(http.StatusFound, "../")
	}
}
