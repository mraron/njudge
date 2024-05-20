package problemset

import (
	"bytes"
	"errors"
	"github.com/mraron/njudge/internal/web/templates"
	"io"
	"mime"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/mraron/njudge/internal/njudge"
	"github.com/mraron/njudge/internal/web/helpers/i18n"
	"github.com/mraron/njudge/pkg/problems"
)

func GetProblem() echo.HandlerFunc {
	return func(c echo.Context) error {
		tr := c.Get(i18n.TranslatorContextKey).(i18n.Translator)
		prob := c.Get("problem").(njudge.Problem)
		info := c.Get("problemInfo").(njudge.ProblemInfo)
		storedData := c.Get("problemStoredData").(njudge.ProblemStoredData)

		c.Set("title", tr.Translate("Statement - %s (%s)",
			tr.TranslateContent(storedData.Titles()).String(), storedData.Name()))

		return c.Render(http.StatusOK, "problemset/problem/problem", struct {
			njudge.Problem
			njudge.ProblemStoredData
			njudge.ProblemInfo
		}{Problem: prob, ProblemStoredData: storedData, ProblemInfo: info})
	}
}

func GetProblemPDF() echo.HandlerFunc {
	return func(c echo.Context) error {
		storedData := c.Get("problemStoredData").(njudge.ProblemStoredData)

		lang := c.Param("language")

		r, err := storedData.GetPDF(njudge.Language(lang))
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
		storedData := c.Get("problemStoredData").(njudge.ProblemStoredData)

		fileLoc, err := storedData.GetFile(c.Param("file"))
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
		storedData := c.Get("problemStoredData").(njudge.ProblemStoredData)
		attachment := c.Param("attachment")

		val, err := storedData.GetAttachment(attachment)
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

func GetProblemRanklist(subList njudge.SubmissionListQuery) echo.HandlerFunc {
	return func(c echo.Context) error {
		tr := c.Get(i18n.TranslatorContextKey).(i18n.Translator)

		problemset, problemName := c.Param("name"), c.Param("problem")
		prob := c.Get("problem").(njudge.Problem)
		storedData := c.Get("problemStoredData").(njudge.ProblemStoredData)

		submissions, err := subList.GetSubmissionList(c.Request().Context(), njudge.SubmissionListRequest{
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

		c.Set("title", tr.Translate("Results - %s (%s)", tr.TranslateContent(storedData.Titles()).String(), storedData.Name()))
		return c.Render(http.StatusOK, "problemset/problem/ranklist", struct {
			Problem           njudge.Problem
			ProblemStoredData njudge.ProblemStoredData
			Submissions       []njudge.Submission
		}{prob, storedData, res})
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

func GetProblemStatus(subList njudge.SubmissionListQuery, probList problems.Store) echo.HandlerFunc {
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
		storedData, err := prob.WithStoredData(probList)
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

		submissionList, err := subList.GetPagedSubmissionList(c.Request().Context(), statusReq)
		if err != nil {
			return err
		}

		qu := (*c.Request().URL).Query()
		links, err := templates.LinksWithCountLimit(submissionList.PaginationData.Page, submissionList.PaginationData.PerPage, int64(submissionList.PaginationData.Count), qu, 5)
		if err != nil {
			return err
		}

		result := templates.SubmissionsViewModel{
			Submissions: submissionList.Submissions,
			Pages:       links,
		}

		c.Set("title", tr.Translate("Submissions - %s (%s)", tr.TranslateContent(storedData.Titles()).String(), storedData.Name()))
		return c.Render(http.StatusOK, "problemset/problem/status", result)
	}
}

func PostProblemTag(tgs njudge.TagsService) echo.HandlerFunc {
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
			return c.NoContent(http.StatusUnauthorized)
		}

		pr := c.Get("problem").(njudge.Problem)
		if err := tgs.Add(c.Request().Context(), data.TagID, pr.ID, u.ID); err != nil {
			return err
		}

		return c.Redirect(http.StatusFound, c.Echo().Reverse("getProblemMain", pr.Problemset, pr.Problem))
	}
}

func DeleteProblemTag(tgs njudge.TagsService) echo.HandlerFunc {
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
			return c.NoContent(http.StatusUnauthorized)
		}

		pr := c.Get("problem").(njudge.Problem)
		if err := tgs.Delete(c.Request().Context(), data.TagID, pr.ID, u.ID); err != nil {
			return err
		}

		return c.Redirect(http.StatusFound, c.Echo().Reverse("getProblemMain", pr.Problemset, pr.Problem))
	}
}
