package problemset

import (
	"bytes"
	"errors"
	"github.com/mraron/njudge/internal/web/templates"
	"github.com/mraron/njudge/internal/web/templates/i18n"
	"github.com/mraron/njudge/pkg/problems/evaluation/output_only"
	"io"
	"mime"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/mraron/njudge/internal/njudge"
	"github.com/mraron/njudge/pkg/problems"
)

func GetProblem(tags njudge.Tags) echo.HandlerFunc {
	return func(c echo.Context) error {
		tr := c.Get(i18n.TranslatorContextKey).(i18n.Translator)
		prob := c.Get("problem").(njudge.Problem)
		info := c.Get("problemInfo").(njudge.ProblemInfo)
		storedData := c.Get("problemStoredData").(njudge.ProblemStoredData)

		title := tr.TranslateContent(storedData.Titles()).String()
		name := storedData.Name()
		c.Set("title", tr.Translate("Statement - %s (%s)", title, name))
		vm := templates.ProblemViewModel{
			Title:        title,
			Name:         name,
			UserInfo:     info.UserInfo,
			ShowTags:     true,
			Tags:         prob.Tags.ToTags(),
			TaskTypeName: storedData.GetTaskType().Name(),
			Languages:    storedData.Languages(),
			Statements:   storedData.Statements(),
			Attachments:  storedData.Attachments(),
		}
		if storedData.GetTaskType().Name() != output_only.Name {
			vm.DisplayLimits = true
			vm.TimeLimit = storedData.TimeLimit()
			vm.MemoryLimit = storedData.MemoryLimit()
		}
		if info.UserInfo != nil {
			if info.UserInfo.SolvedStatus == njudge.Solved {
				vm.CanAddTags = true
			} else {
				if u := c.Get("user").(*njudge.User); u != nil && !u.Settings.ShowUnsolvedTags {
					vm.ShowTags = false
				}
			}
		}
		var err error
		if vm.TagsToAdd, err = tags.GetAll(c.Request().Context()); err != nil {
			return err
		}

		return templates.Render(c, http.StatusOK, templates.Problem(vm))
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

func GetProblemRanklist(subList njudge.SubmissionListQuery, users njudge.Users) echo.HandlerFunc {
	return func(c echo.Context) error {
		tr := c.Get(i18n.TranslatorContextKey).(i18n.Translator)

		problemset, problemName := c.Param("name"), c.Param("problem")
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

		ss, err := storedData.StatusSkeleton("")
		if err != nil {
			return err
		}

		hadUser := make(map[int]bool)
		vm := templates.ProblemRanklistViewModel{
			MaxScore: ss.Feedback[0].MaxScore(),
			Rows:     nil,
		}
		for ind := range submissions.Submissions {
			if _, ok := hadUser[submissions.Submissions[ind].UserID]; ok {
				continue
			}
			u, err := users.Get(c.Request().Context(), submissions.Submissions[ind].UserID)
			if err != nil {
				return err
			}
			vm.Rows = append(vm.Rows, templates.ProblemRanklistRow{
				ID:    submissions.Submissions[ind].ID,
				Name:  u.Name,
				Score: float64(submissions.Submissions[ind].Score),
			})
			hadUser[submissions.Submissions[ind].UserID] = true
		}

		c.Set("title", tr.Translate("Results - %s (%s)", tr.TranslateContent(storedData.Titles()).String(), storedData.Name()))
		return templates.Render(c, http.StatusOK, templates.ProblemRanklist(vm))
	}
}

func GetProblemSubmit() echo.HandlerFunc {
	return func(c echo.Context) error {
		tr := c.Get(i18n.TranslatorContextKey).(i18n.Translator)

		p := c.Get("problem").(njudge.Problem)
		storedData := c.Get("problemStoredData").(njudge.ProblemStoredData)
		info := c.Get("problemInfo").(njudge.ProblemInfo)

		title := tr.TranslateContent(storedData.Titles()).String()
		vm := templates.ProblemSubmitViewModel{
			Problemset: p.Problemset,
			Name:       p.Problem,
			Title:      title,
			UserInfo:   info.UserInfo,
			Languages:  storedData.Languages(),
		}

		c.Set("title", tr.Translate("Submit - %s (%s)", title, storedData.Name()))

		return templates.Render(c, http.StatusOK, templates.ProblemSubmit(vm))
	}
}

type GetProblemStatusRequest struct {
	AC     string `query:"ac"`
	UserID int    `query:"user_id"`
	Page   int    `query:"page"`

	Problemset string `param:"name"`
	Problem    string `param:"problem"`
}

func GetProblemStatus(subList njudge.SubmissionListQuery, probList problems.Store) echo.HandlerFunc {
	return func(c echo.Context) error {
		tr := c.Get(i18n.TranslatorContextKey).(i18n.Translator)

		prob := c.Get("problem").(njudge.Problem)
		storedData, err := prob.WithStoredData(probList)
		if err != nil {
			return err
		}

		data := GetProblemStatusRequest{}
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
		return templates.Render(c, http.StatusOK, templates.ProblemStatus(result))
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
