package problemset

import (
	"bytes"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/mraron/njudge/internal/web/domain/problem"
	"github.com/mraron/njudge/internal/web/helpers"
	"github.com/mraron/njudge/internal/web/helpers/i18n"
	"github.com/mraron/njudge/internal/web/helpers/pagination"
	"github.com/mraron/njudge/internal/web/models"
	"github.com/mraron/njudge/internal/web/services"
	"github.com/mraron/njudge/pkg/problems"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"io"
	"mime"
	"net/http"
	"path/filepath"
	"strconv"
)

type ProblemInfo struct {
	ID          string   `json:"id"`
	Title       string   `json:"title"`
	TimeLimit   string   `json:"timeLimit"`
	MemoryLimit string   `json:"memoryLimit"`
	Tags        []string `json:"tags"`
	Type        string   `json:"type"`
}

func ProblemInfoFromProblem(ctx context.Context, DB *sql.DB, tr i18n.Translator, prob problem.Problem) (*ProblemInfo, error) {
	tags, err := prob.ProblemRel.ProblemProblemTags().All(ctx, DB)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}

	lst := make([]string, 0)
	for _, tag := range tags {
		if t, err := tag.Tag().One(ctx, DB); err != nil {
			return nil, err
		} else {
			lst = append(lst, t.Name)
		}
	}

	return &ProblemInfo{
		ID:          prob.ProblemRel.Problem,
		Title:       tr.TranslateContent(prob.Titles()).String(),
		TimeLimit:   fmt.Sprintf("%d", prob.TimeLimit()),
		MemoryLimit: fmt.Sprintf("%d", prob.MemoryLimit()/1024/1024),
		Tags:        lst,
		Type:        prob.GetTaskType().Name(),
	}, nil
}

type ProblemAttachment struct {
	Type string `json:"type"`
	Name string `json:"name"`
	Href string `json:"href"`
}

type ProblemAttachmentData struct {
	Statements []ProblemAttachment `json:"statements"`
	Files      []ProblemAttachment `json:"files"`
}

func ProblemAttachmentDataFromProblem(tr i18n.Translator, ps, probName string, prob problem.Problem) ProblemAttachmentData {
	res := ProblemAttachmentData{
		Statements: make([]ProblemAttachment, 0),
		Files:      make([]ProblemAttachment, 0),
	}
	for _, pdf := range prob.Problem.Statements().FilterByType(problems.DataTypePDF) {
		res.Statements = append(res.Statements, ProblemAttachment{
			Name: pdf.Locale(),
			Type: "pdf",
			Href: fmt.Sprintf("/problemset/%s/%s/pdf/%s/", ps, probName, pdf.Locale()),
		})
	}
	for _, elem := range prob.Attachments() {
		res.Files = append(res.Files, ProblemAttachment{
			Name: elem.Name(),
			Href: fmt.Sprintf("/problemset/%s/%s/attachment/%s/", ps, probName, elem.Name()),
		})
	}
	return res
}

func GetProblem(DB *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		tr := c.Get(i18n.TranslatorContextKey).(i18n.Translator)
		prob := c.Get("problem").(problem.Problem)
		//stats := c.Get("problemStats").(problem.StatsData)

		pinfo, err := ProblemInfoFromProblem(c.Request().Context(), DB.DB, tr, prob)
		if err != nil {
			return err
		}

		pdata := ProblemAttachmentDataFromProblem(tr, c.Param("name"), c.Param("problem"), prob)

		return c.JSON(http.StatusOK, struct {
			Info        ProblemInfo           `json:"info"`
			Attachments ProblemAttachmentData `json:"attachments"`
		}{
			*pinfo,
			pdata,
		})
	}
}

func GetProblemPDF() echo.HandlerFunc {
	return func(c echo.Context) error {
		p := c.Get("problem").(problem.Problem)

		lang := c.Param("language")

		dat, err := p.GetPDF(lang)
		if err != nil {
			return err
		}

		return c.Blob(http.StatusOK, "application/pdf", dat)
	}
}

func GetProblemFile() echo.HandlerFunc {
	return func(c echo.Context) error {
		p := c.Get("problem").(problem.Problem)

		fileLoc, err := p.GetFile(c.Param("file"))
		if err == problem.ErrorFileNotFound {
			return echo.NewHTTPError(http.StatusNotFound, err)
		} else if err != nil {
			return err
		}

		return c.File(fileLoc)
	}
}

func GetProblemAttachment() echo.HandlerFunc {
	return func(c echo.Context) error {
		p := c.Get("problem").(problem.Problem)
		attachment := c.Param("attachment")

		val, err := p.GetAttachment(attachment)
		if err == problem.ErrorFileNotFound {
			return echo.NewHTTPError(http.StatusNotFound, err)
		} else if err != nil {
			return err
		}

		dat, err := val.Value()
		if err != nil {
			return err
		}

		c.Response().Header().Set("Content-Disposition", "attachment; filename="+val.Name())
		c.Response().Header().Set("Content-Type", mime.TypeByExtension(filepath.Ext(val.Name())))
		c.Response().Header().Set("Content-Length", strconv.Itoa(len(dat)))

		if _, err := io.Copy(c.Response(), bytes.NewReader(dat)); err != nil {
			return err
		}

		return c.NoContent(http.StatusOK)

	}
}

type ProblemRanklistResult struct {
	Username     string  `json:"username"`
	Score        float64 `json:"score"`
	SubmissionID int     `json:"submissionID"`
	Accepted     bool    `json:"accepted"`
}

type ProblemRanklist struct {
	MaxScore float64                 `json:"maxScore"`
	Results  []ProblemRanklistResult `json:"results"`
}

func ProblemRanklistFromSubmissions(ctx context.Context, DB *sql.DB, prob problem.Problem, subs []*models.Submission) (*ProblemRanklist, error) {
	testset, err := prob.StatusSkeleton("")
	if err != nil {
		return nil, err
	}

	results := make([]ProblemRanklistResult, len(subs))
	for i := range subs {
		user, err := subs[i].User().One(ctx, DB)
		if err != nil {
			return nil, err
		}

		results[i] = ProblemRanklistResult{
			Username:     user.Name,
			Score:        float64(subs[i].Score.Float32),
			SubmissionID: subs[i].ID,
			Accepted:     subs[i].Verdict == 0,
		}
	}

	res := &ProblemRanklist{
		MaxScore: testset.Feedback[0].MaxScore(),
		Results:  results,
	}

	return res, nil
}

func GetProblemRanklist(DB *sqlx.DB) echo.HandlerFunc {
	var (
		queryTemplate = `SELECT DISTINCT ON (s1.user_id) s1.* FROM
                                        (SELECT s1.user_id, MAX(s1.score) as score FROM submissions s1 WHERE problemset=$1 AND problem=$2 GROUP BY s1.user_id) s2
                                        INNER JOIN submissions s1 ON s1.user_id=s2.user_id AND s1.score=s2.score AND s1.problemset=$1 AND s1.problem=$2`
		queryAll   = fmt.Sprintf("SELECT * FROM (%s) t ORDER BY t.score DESC LIMIT $3 OFFSET $4", queryTemplate)
		queryCount = fmt.Sprintf("SELECT COUNT(*) FROM (%s) t", queryTemplate)
	)

	type request struct {
		Page int `query:"page"`
	}

	return func(c echo.Context) error {
		data := request{}
		if err := c.Bind(&data); err != nil {
			return err
		}

		problemset, problemName := c.Param("name"), c.Param("problem")
		prob := c.Get("problem").(problem.Problem)

		if data.Page <= 0 {
			data.Page = 1
		}

		cnt := 0
		row := queries.Raw(queryCount, problemset, problemName).QueryRowContext(c.Request().Context(), DB)
		if err := row.Scan(&cnt); err != nil {
			return err
		}

		perPage := 20
		pages := (cnt + perPage - 1) / perPage
		if data.Page > pages && pages > 0 {
			data.Page = pages
		}

		sbs := make([]*models.Submission, 0)
		if err := queries.Raw(queryAll, problemset, problemName, perPage, (data.Page-1)*perPage).Bind(c.Request().Context(), DB, &sbs); err != nil {
			return err
		}

		ranklist, err := ProblemRanklistFromSubmissions(c.Request().Context(), DB.DB, prob, sbs)
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, struct {
			Ranklist       ProblemRanklist `json:"ranklist"`
			PaginationData pagination.Data `json:"paginationData"`
		}{
			Ranklist: *ranklist,
			PaginationData: pagination.Data{
				Page:     data.Page,
				PerPage:  perPage,
				LastPage: pages,
			},
		})
	}
}

func GetProblemSubmit() echo.HandlerFunc {
	return func(c echo.Context) error {
		tr := c.Get(i18n.TranslatorContextKey).(i18n.Translator)

		prob := c.Get("problem").(problem.Problem)
		stats := c.Get("problemStats").(problem.StatsData)

		c.Set("title", tr.Translate("Submit - %s (%s)", tr.TranslateContent(prob.Titles()).String(), prob.Name()))
		return c.Render(http.StatusOK, "problemset/problem/submit", struct {
			problem.Problem
			problem.StatsData
		}{Problem: prob, StatsData: stats})
	}
}

func GetProblemStatus(DB *sqlx.DB, problemStore problems.Store, statusPageService services.StatusPageService) echo.HandlerFunc {
	type request struct {
		AC   string `query:"ac"`
		Own  bool   `query:"own"`
		Page int    `query:"page"`

		Problemset string `param:"name"`
		Problem    string `param:"problem"`
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
		}

		if user := c.Get("user").(*models.User); data.Own && user != nil {
			statusReq.UserID = user.ID
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

func PostProblemTag(tgs services.TagsService) echo.HandlerFunc {
	type request struct {
		TagID int `form:"tagID"`
	}
	return func(c echo.Context) error {
		data := request{}
		if err := c.Bind(&data); err != nil {
			return err
		}

		u := c.Get("user").(*models.User)
		if u == nil {
			return helpers.UnauthorizedError(c)
		}

		pr := c.Get("problem").(problem.Problem)
		if err := tgs.Add(c.Request().Context(), data.TagID, pr.ID, u.ID); err != nil {
			return err
		}

		return c.Redirect(http.StatusFound, c.Echo().Reverse("getProblemMain", pr.Problemset, pr.ProblemRel.Problem))
	}
}

func DeleteProblemTag(tgs services.TagsService) echo.HandlerFunc {
	type request struct {
		TagID int `param:"id"`
	}
	return func(c echo.Context) error {
		data := request{}
		if err := c.Bind(&data); err != nil {
			return err
		}

		u := c.Get("user").(*models.User)
		if u == nil {
			return helpers.UnauthorizedError(c)
		}

		pr := c.Get("problem").(problem.Problem)
		if err := tgs.Delete(c.Request().Context(), data.TagID, pr.ID, u.ID); err != nil {
			return err
		}

		return c.Redirect(http.StatusFound, c.Echo().Reverse("getProblemMain", pr.Problemset, pr.ProblemRel.Problem))
	}
}
