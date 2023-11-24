package problemset

import (
	"bytes"
	"errors"
	"io"
	"mime"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/mraron/njudge/internal/njudge"
	"github.com/mraron/njudge/internal/web/helpers"
	"github.com/mraron/njudge/internal/web/helpers/i18n"
	"github.com/mraron/njudge/internal/web/helpers/pagination"
	"github.com/mraron/njudge/pkg/problems"
)

func GetProblem() echo.HandlerFunc {
	return func(c echo.Context) error {
		tr := c.Get(i18n.TranslatorContextKey).(i18n.Translator)
		prob := c.Get("problem").(njudge.Problem)
		info := c.Get("problemInfo").(njudge.ProblemInfo)
		sdata := c.Get("problemStoredData").(njudge.ProblemStoredData)

		c.Set("title", tr.Translate("Statement - %s (%s)",
			tr.TranslateContent(sdata.Titles()).String(), sdata.Name()))

		return c.Render(http.StatusOK, "problemset/problem/problem", struct {
			njudge.Problem
			njudge.ProblemStoredData
			njudge.ProblemInfo
		}{Problem: prob, ProblemStoredData: sdata, ProblemInfo: info})
	}
}

func GetProblemPDF() echo.HandlerFunc {
	return func(c echo.Context) error {
		sdata := c.Get("problemStoredData").(njudge.ProblemStoredData)

		lang := c.Param("language")

		r, err := sdata.GetPDF(njudge.Language(lang))
		if err != nil {
			return err
		}

		data, err := io.ReadAll(r)
		if err != nil {
			return err
		}

		return c.Blob(http.StatusOK, "application/pdf", data)
	}
}

func GetProblemFile() echo.HandlerFunc {
	return func(c echo.Context) error {
		sdata := c.Get("problemStoredData").(njudge.ProblemStoredData)

		fileLoc, err := sdata.GetFile(c.Param("file"))
		if errors.Is(err, njudge.ErrorFileNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, err)
		} else if err != nil {
			return err
		}

		return c.File(fileLoc)
	}
}

func GetProblemAttachment() echo.HandlerFunc {
	return func(c echo.Context) error {
		sdata := c.Get("problemStoredData").(njudge.ProblemStoredData)
		attachment := c.Param("attachment")

		val, err := sdata.GetAttachment(attachment)
		if errors.Is(err, njudge.ErrorFileNotFound) {
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

func GetProblemRanklist(slist njudge.SubmissionListQuery) echo.HandlerFunc {
	return func(c echo.Context) error {
		tr := c.Get(i18n.TranslatorContextKey).(i18n.Translator)

		problemset, problemName := c.Param("name"), c.Param("problem")
		prob := c.Get("problem").(njudge.Problem)
		sdata := c.Get("problemStoredData").(njudge.ProblemStoredData)

		submissions, err := slist.GetSubmissionList(c.Request().Context(), njudge.SubmissionListRequest{
			Problemset: problemset,
			Problem:    problemName,
			SortDir:    njudge.SortDESC,
			SortField:  njudge.SubmissionSortFieldScore,
		})
		if err != nil {
			return err
		}

		hadUser := make(map[int]bool)
		res := make([]njudge.Submission, 0)
		for ind := range submissions.Submissions {
			if _, ok := hadUser[submissions.Submissions[ind].UserID]; ok {
				continue
			}

			res = append(res, submissions.Submissions[ind])
			hadUser[submissions.Submissions[ind].UserID] = true
		}

		c.Set("title", tr.Translate("Results - %s (%s)", tr.TranslateContent(sdata.Titles()).String(), sdata.Name()))
		return c.Render(http.StatusOK, "problemset/problem/ranklist", struct {
			Problem           njudge.Problem
			ProblemStoredData njudge.ProblemStoredData
			Submissions       []njudge.Submission
		}{prob, sdata, res})
	}
}

func GetProblemSubmit() echo.HandlerFunc {
	return func(c echo.Context) error {
		tr := c.Get(i18n.TranslatorContextKey).(i18n.Translator)

		p := c.Get("problem").(njudge.Problem)
		prob := c.Get("problemStoredData").(njudge.ProblemStoredData)
		info := c.Get("problemInfo").(njudge.ProblemInfo)

		c.Set("title", tr.Translate("Submit - %s (%s)", tr.TranslateContent(prob.Titles()).String(), prob.Name()))
		return c.Render(http.StatusOK, "problemset/problem/submit", struct {
			njudge.Problem
			njudge.ProblemStoredData
			njudge.ProblemInfo
		}{Problem: p, ProblemStoredData: prob, ProblemInfo: info})
	}
}

func GetProblemStatus(slist njudge.SubmissionListQuery, pstore problems.Store) echo.HandlerFunc {
	type request struct {
		AC     string `query:"ac"`
		UserID int    `query:"user_id"`
		Page   int    `query:"page"`

		Problemset string `param:"name"`
		Problem    string `param:"problem"`
	}
	return func(c echo.Context) error {
		tr := c.Get(i18n.TranslatorContextKey).(i18n.Translator)

		prob := c.Get("problem").(njudge.Problem)
		sdata, err := prob.WithStoredData(pstore)
		if err != nil {
			return err
		}

		data := request{}
		if err := c.Bind(&data); err != nil {
			return err
		}

		if data.Page <= 0 {
			data.Page = 1
		}

		statusReq := njudge.SubmissionListRequest{

			Page:      data.Page,
			PerPage:   20,
			SortDir:   njudge.SortDESC,
			SortField: njudge.SubmissionSortFieldID,

			Problemset: data.Problemset,
			Problem:    data.Problem,

			UserID: 0,
		}

		if data.AC == "1" {
			ac := njudge.VerdictAC
			statusReq.Verdict = &ac
		}

		submissionList, err := slist.GetPagedSubmissionList(c.Request().Context(), statusReq)
		if err != nil {
			return err
		}

		qu := (*c.Request().URL).Query()
		links, err := pagination.Links(submissionList.PaginationData.Page, submissionList.PaginationData.PerPage, int64(submissionList.PaginationData.Count), qu)
		if err != nil {
			return err
		}

		result := StatusPage{
			Submissions: submissionList.Submissions,
			Pages:       links,
		}

		c.Set("title", tr.Translate("Submissions - %s (%s)", tr.TranslateContent(sdata.Titles()).String(), sdata.Name()))
		return c.Render(http.StatusOK, "problemset/problem/status", result)
	}
}

func PostProblemTag(tags njudge.Tags, problems njudge.Problems) echo.HandlerFunc {
	type request struct {
		TagID int `form:"tagID"`
	}
	return func(c echo.Context) error {
		data := request{}
		if err := c.Bind(&data); err != nil {
			return err
		}

		u := c.Get("user").(*njudge.User)
		if u == nil {
			return helpers.UnauthorizedError(c)
		}

		tg, err := tags.Get(c.Request().Context(), data.TagID)
		if err != nil {
			return err
		}

		pr := c.Get("problem").(njudge.Problem)
		if err := pr.AddTag(*tg, u.ID); err != nil {
			return err
		}

		if err := problems.Update(c.Request().Context(), pr); err != nil {
			return err
		}

		return c.Redirect(http.StatusFound, c.Echo().Reverse("getProblemMain", pr.Problemset, pr.Problem))
	}
}

func DeleteProblemTag(tags njudge.Tags, problems njudge.Problems) echo.HandlerFunc {
	type request struct {
		TagID int `param:"id"`
	}
	return func(c echo.Context) error {
		data := request{}
		if err := c.Bind(&data); err != nil {
			return err
		}

		u := c.Get("user").(*njudge.User)
		if u == nil {
			return helpers.UnauthorizedError(c)
		}

		tg, err := tags.Get(c.Request().Context(), data.TagID)
		if err != nil {
			return err
		}

		pr := c.Get("problem").(njudge.Problem)
		if err := pr.DeleteTag(*tg); err != nil {
			return err
		}

		if err := problems.Update(c.Request().Context(), pr); err != nil {
			return err
		}

		return c.Redirect(http.StatusFound, c.Echo().Reverse("getProblemMain", pr.Problemset, pr.Problem))
	}
}
