package problemset

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/mraron/njudge/internal/web/domain/submission"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/mraron/njudge/internal/web/domain/problem"
	"github.com/mraron/njudge/internal/web/helpers/i18n"
	"github.com/mraron/njudge/internal/web/helpers/pagination"
	"github.com/mraron/njudge/internal/web/helpers/ui"
	"github.com/mraron/njudge/internal/web/models"
	"github.com/mraron/njudge/internal/web/services"
	"github.com/mraron/njudge/pkg/problems"
)

type ProblemsetRow struct {
	ID           int                  `json:"id"`
	Problem      string               `json:"problem"`
	Title        string               `json:"title"`
	Category     ui.Link              `json:"category"`
	Tags         []string             `json:"tags"`
	SolverCount  int                  `json:"solverCount"`
	SolvedStatus problem.SolvedStatus `json:"solvedStatus"`
}

type ProblemList struct {
	PaginationData pagination.Data `json:"paginationData"`
	Problems       []ProblemsetRow `json:"problems"`
	SolverSorter   ui.SortColumn

	Filtered bool

	TitleFilter string
	TagsFilter  string
}

func GetProblemList(DB *sqlx.DB, problemListService services.ProblemListService, problemRepo problem.Repository, problemStatsService services.ProblemStatsService) echo.HandlerFunc {
	type request struct {
		Page  int `query:"page"`
		Order string
		By    string

		TitleFilter    string `query:"title"`
		CategoryFilter int    `query:"category"`
		TagFilter      string `query:"tags"`

		Problemset string `param:"name"`
	}
	return func(c echo.Context) error {
		tr := c.Get(i18n.TranslatorContextKey).(i18n.Translator)

		data := request{}
		if err := c.Bind(&data); err != nil {
			return err
		}
		if data.Page <= 0 {
			data.Page = 1
		}

		data.Order, data.By = "DESC", "id"
		if c.QueryParam("by") == "solver_count" {
			data.By = "solver_count"
		}
		if c.QueryParam("order") == "ASC" {
			data.Order = "ASC"
		}

		paginationData := pagination.Data{
			Page:      data.Page,
			PerPage:   20,
			SortDir:   data.Order,
			SortField: data.By,
		}
		listRequest := services.ProblemListRequest{
			Problemset:  data.Problemset,
			Pagination:  paginationData,
			TitleFilter: data.TitleFilter,
			GETData:     c.Request().URL.Query(),
		}

		if data.TagFilter != "" {
			listRequest.TagFilter = strings.Split(data.TagFilter, ",")
		}

		if data.CategoryFilter != 0 {
			if data.CategoryFilter == -1 {
				listRequest.CategoryFilter = problem.NewCategoryEmptyFilter()
			} else {
				listRequest.CategoryFilter = problem.NewCategoryIDFilter(data.CategoryFilter)
			}
		}

		problemList, err := problemListService.GetProblemList(c.Request().Context(), listRequest)
		if err != nil {
			return err
		}
		paginationData.LastPage = len(problemList.Pages) - 2

		result := ProblemList{
			PaginationData: paginationData,
			Problems:       []ProblemsetRow{},
		}
		for ind := range problemList.Problems {
			p, err := problemRepo.Get(c.Request().Context(), problemList.Problems[ind].ID)
			if err != nil {
				return err
			}
			stat, err := problemStatsService.GetStatsData(c.Request().Context(), *p, c.Get("userID").(int))
			if err != nil {
				return err
			}

			result.Problems = append(result.Problems, ProblemsetRow{
				ID:           p.ProblemRel.ID,
				Problem:      p.ProblemRel.Problem,
				Title:        tr.TranslateContent(p.Titles()).String(),
				Category:     stat.CategoryLink,
				Tags:         stat.Tags.ToStringSlice(),
				SolverCount:  stat.SolverCount,
				SolvedStatus: stat.SolvedStatus,
			})
		}

		sortOrder, qu := "", c.Request().URL.Query()
		if qu.Get("by") == "solver_count" {
			sortOrder = qu.Get("order")
			if qu.Get("order") == "DESC" {
				qu.Set("order", "ASC")
			} else {
				qu.Set("order", "")
				qu.Set("by", "")
			}
		} else {
			qu.Set("by", "solver_count")
			qu.Set("order", "DESC")
		}
		result.SolverSorter = ui.SortColumn{
			Order: sortOrder,
			Href:  "?" + qu.Encode(),
		}

		result.Filtered = listRequest.IsFiltered()
		result.TitleFilter = data.TitleFilter
		result.TagsFilter = data.TagFilter

		return c.JSON(http.StatusOK, result)
	}
}

type StatusRow struct {
	Id          int                 `json:"id"`
	Date        string              `json:"date"`
	User        string              `json:"user"`
	Problem     ui.Link             `json:"problem"`
	Lang        string              `json:"language"`
	VerdictName string              `json:"verdictName"`
	VerdictType problem.VerdictType `json:"verdictType"`
	Time        int                 `json:"time"`
	Memory      int                 `json:"memory"`
	Score       float64             `json:"score"`
	MaxScore    float64             `json:"maxScore"`
}

func StatusRowFromSubmission(ctx context.Context, DB *sql.DB, problemStore problems.Store, tr i18n.Translator, sub *submission.Submission) (*StatusRow, error) {
	user, err := sub.User().One(ctx, DB)
	if err != nil {
		return nil, err
	}

	verdict := problem.Verdict(sub.Verdict)

	status := &problems.Status{}
	if err := status.Scan(sub.Status); err != nil {
		return nil, err
	}

	prob, err := problemStore.Get(sub.Problem)
	res := &StatusRow{
		Id:   sub.ID,
		Date: sub.Submitted.String(),
		User: user.Name,
		Problem: ui.Link{
			Text: tr.TranslateContent(prob.Titles()).String(),
			Href: fmt.Sprintf("/problemset/%s/%s/",
				sub.Problemset,
				sub.Problem),
		},
		Lang:        sub.Language,
		VerdictName: verdict.Translate(tr),
		VerdictType: verdict.VerdictType(),
	}

	if len(status.Feedback) > 0 {
		res.Time = int(status.Feedback[0].MaxTimeSpent() / time.Millisecond)
		res.Memory = status.Feedback[0].MaxMemoryUsage()
		res.Score = status.Feedback[0].Score()
		res.MaxScore = status.Feedback[0].MaxScore()
	}

	return res, nil
}

func GetStatus(DB *sqlx.DB, problemStore problems.Store, statusPageService services.StatusPageService) echo.HandlerFunc {
	type request struct {
		AC         string `query:"ac"`
		UserID     int    `query:"user_id"`
		Problemset string `query:"problem_set"`
		Problem    string `query:"problem"`
		Page       int    `query:"page"`
	}
	return func(c echo.Context) error {
		tr := c.Get(i18n.TranslatorContextKey).(i18n.Translator)

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
			Problemset: data.Problemset,
			Problem:    data.Problem,
			UserID:     data.UserID,
		}

		if data.AC == "1" {
			ac := problems.VerdictAC
			statusReq.Verdict = &ac
		}

		statusPage, err := statusPageService.GetStatusPage(c.Request().Context(), statusReq)
		if err != nil {
			return err
		}

		statusRows := make([]*StatusRow, len(statusPage.Submissions))
		for i := range statusPage.Submissions {
			sub := &statusPage.Submissions[i]
			if statusRows[i], err = StatusRowFromSubmission(c.Request().Context(), DB.DB,
				problemStore, tr, sub); err != nil {
				return err
			}
		}

		return c.JSON(http.StatusOK, struct {
			PaginationData pagination.Data `json:"paginationData"`
			Submissions    []*StatusRow    `json:"submissions"`
		}{statusPage.PaginationData, statusRows})
	}
}

func PostSubmit(subService services.SubmitService) echo.HandlerFunc {
	type request struct {
		Problemset     string `param:"name"`
		ProblemName    string `form:"problem"`
		LanguageName   string `form:"language"`
		SubmissionCode string `form:"submissionCode"`
	}
	return func(c echo.Context) error {
		u := c.Get("user").(*models.User)

		data := request{}
		if err := c.Bind(&data); err != nil {
			return err
		}

		code := data.SubmissionCode
		if len(code) == 0 {
			fileHeader, err := c.FormFile("file")
			if err != nil {
				return err
			}

			f, err := fileHeader.Open()
			if err != nil {
				return err
			}

			contents, err := io.ReadAll(f)
			if err != nil {
				return err
			}

			code = string(contents)
			if err := f.Close(); err != nil {
				return err
			}
		}

		_, err := subService.Submit(c.Request().Context(), services.SubmitRequest{
			UserID:     u.ID,
			Problemset: data.Problemset,
			Problem:    data.ProblemName,
			Language:   data.LanguageName,
			Source:     []byte(code),
		})
		if err != nil {
			return err
		}

		return c.String(http.StatusOK, "")
	}
}
