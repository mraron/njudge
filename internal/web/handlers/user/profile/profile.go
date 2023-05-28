package profile

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/mraron/njudge/internal/web/domain/submission"
	"github.com/mraron/njudge/internal/web/helpers/pagination"
	"github.com/mraron/njudge/internal/web/models"
	"github.com/mraron/njudge/internal/web/services"
	. "github.com/volatiletech/sqlboiler/v4/queries/qm"
	"net/http"
)

func GetProfile(DB *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		u := c.Get("profile").(*models.User)
		var (
			solved, attempted models.SubmissionSlice
			err               error
		)

		solved, err = models.Submissions(Select("max(submissions.id) as id, problemset, problem"),
			Where("user_id = ?", u.ID), Where("verdict = ?", 0),
			GroupBy("submissions.problemset, submissions.problem"),
		).All(c.Request().Context(), DB)
		if err != nil {
			return err
		}

		attempted, err = models.Submissions(Select("max(submissions.id) as id, problemset, problem"),
			Where("user_id = ?", u.ID), Where("verdict <> ?", 0),
			Where("not exists(select id from submissions as other where other.user_id = ? and "+
				"verdict = 0 and other.problem = submissions.problem and other.problemset=submissions.problemset)", u.ID),
			GroupBy("submissions.problemset, submissions.problem"),
		).All(c.Request().Context(), DB)
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

func GetSubmissions(statusPageService services.StatusPageService) echo.HandlerFunc {
	type request struct {
		Page int `query:"page"`
	}
	return func(c echo.Context) error {
		u := c.Get("profile").(*models.User)

		data := request{}
		if err := c.Bind(&data); err != nil {
			return err
		}

		if data.Page <= 0 {
			data.Page = 1
		}

		statusReq := services.StatusPageRequest{
			Pagination: pagination.Data{
				Page:      data.Page,
				PerPage:   20,
				SortDir:   "DESC",
				SortField: "id",
			},
			UserID:    u.ID,
			GETValues: c.Request().URL.Query(),
		}

		statusPage, err := statusPageService.GetStatusPage(c.Request().Context(), statusReq)
		if err != nil {
			return err
		}

		c.Set("title", fmt.Sprintf("%s beküldései", u.Name))
		return c.Render(http.StatusOK, "user/profile/submissions", struct {
			User       *models.User
			StatusPage *submission.StatusPage
		}{u, statusPage})
	}
}
