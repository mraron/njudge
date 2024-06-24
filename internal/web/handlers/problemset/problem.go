package problemset

import (
	"bytes"
	"cmp"
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/mraron/njudge/internal/njudge"
	"github.com/mraron/njudge/internal/web/templates"
	"github.com/mraron/njudge/internal/web/templates/i18n"
	"github.com/mraron/njudge/pkg/language/memory"
	"github.com/mraron/njudge/pkg/problems"
	"github.com/mraron/njudge/pkg/problems/evaluation/output_only"
	"golang.org/x/exp/slices"
	"io"
	"mime"
	"net/http"
	"path/filepath"
	"strconv"
	"time"
)

func GetProblem(tags njudge.Tags) echo.HandlerFunc {
	dataType := func(t string) string {
		switch t {
		case "pdf":
			return problems.DataTypePDF
		case "html":
			return problems.DataTypeHTML
		}
		return ""
	}
	return func(c echo.Context) error {
		tr := c.Get(i18n.TranslatorContextKey).(i18n.Translator)
		prob := c.Get(ProblemContextKey).(njudge.Problem)
		info := c.Get(ProblemInfoContextKey).(njudge.ProblemInfo)
		storedData := c.Get(ProblemStoredDataContextKey).(njudge.ProblemStoredData)

		title := tr.TranslateContent(storedData.Titles()).String()
		name := storedData.Name()
		c.Set(templates.TitleContextKey, tr.Translate("Statement - %s (%s)", title, name))
		vm := templates.ProblemViewModel{
			Title:        title,
			Problemset:   prob.Problemset,
			Name:         name,
			ShowTags:     false,
			Tags:         prob.Tags.ToTags(),
			TaskTypeName: storedData.GetTaskType().Name(),
			Languages:    storedData.Languages(),
			UserInfo:     info.UserInfo,
			Attachments:  storedData.Attachments(),
			Statements:   nil,
			TagsToAdd:    nil,
			Author:       prob.Author,
		}
		for _, lang := range storedData.Languages() {
			res, err := extractCompileAndRunCommand(c.Request().Context(), storedData.GetTaskType(), lang)
			if err != nil {
				return err
			}
			vm.LanguageCompileCommands = append(vm.LanguageCompileCommands, templates.LanguageCompileCommand{
				Name:    lang.DisplayName(),
				Command: res,
			})
		}
		vm.InputFile, vm.OutputFile = storedData.InputOutputFiles()
		for _, st := range storedData.Statements() {
			if st.IsHTML() || st.IsPDF() {
				vm.Statements = append(vm.Statements, st)
			}
		}
		statementType := c.QueryParam("type")
		locale := c.QueryParam("locale")
		poss := vm.Statements
		if statementType != "" {
			poss = poss.FilterByType(dataType(statementType))
		}
		if locale != "" {
			poss = poss.FilterByLocale(locale)
		}
		if len(poss) > 0 {
			s := tr.TranslateContent(poss)
			vm.Statement = &s
		} else { // prefer html then pdf
			if HTMLs := vm.Statements.FilterByLocale(problems.DataTypeHTML); len(HTMLs) > 0 {
				s := tr.TranslateContent(HTMLs)
				vm.Statement = &s
			} else if PDFs := vm.Statements.FilterByLocale(problems.DataTypePDF); len(PDFs) > 0 {
				s := tr.TranslateContent(PDFs)
				vm.Statement = &s
			}
		}

		if storedData.GetTaskType().Name() != output_only.Name {
			vm.DisplayLimits = true
			vm.TimeLimit = storedData.TimeLimit()
			vm.MemoryLimit = storedData.MemoryLimit()
		}
		if info.UserInfo != nil {
			if info.UserInfo.SolvedStatus == njudge.Solved {
				vm.CanAddTags = true
				vm.ShowTags = true
			} else {
				if u := c.Get(templates.UserContextKey).(*njudge.User); u != nil && u.Settings.ShowUnsolvedTags {
					vm.ShowTags = true
				}
			}
			if u := c.Get(templates.UserContextKey).(*njudge.User); u != nil && u.Role == "admin" {
				vm.CanEdit = true
			}
		}
		var err error
		if vm.TagsToAdd, err = tags.GetAll(c.Request().Context()); err != nil {
			return err
		}

		return templates.Render(c, http.StatusOK, templates.Problem(vm))
	}
}

func GetProblemEdit(users njudge.Users, cs njudge.Categories) echo.HandlerFunc {
	return func(c echo.Context) error {
		tr := c.Get(i18n.TranslatorContextKey).(i18n.Translator)
		if u := c.Get(templates.UserContextKey).(*njudge.User); u.Role != "admin" {
			return echo.NotFoundHandler(c)
		}
		prob := c.Get(ProblemContextKey).(njudge.Problem)

		categories, err := cs.GetAll(c.Request().Context())
		if err != nil {
			return err
		}

		par := make(map[int]int)
		for ind := range categories {
			if categories[ind].ParentID.Valid {
				par[categories[ind].ID] = categories[ind].ParentID.Int
			}
		}
		categoryNameByID := make(map[int]string)
		for ind := range categories {
			categoryNameByID[categories[ind].ID] = categories[ind].Name
		}
		cat := -1
		if prob.Category != nil {
			cat = prob.Category.ID
		}
		categoriesList := makeCategoryFilterOptions(tr, categories, cat, categoryNameByID, par)

		tags := make([]templates.ProblemTag, len(prob.Tags))
		for i := range prob.Tags {
			tags[i] = templates.ProblemTag{
				Name:  prob.Tags[i].Tag.Name,
				Added: prob.Tags[i].Added,
			}
			user, err := users.Get(c.Request().Context(), prob.Tags[i].UserID)
			if err != nil {
				return err
			}
			tags[i].User = user.Name
		}

		return templates.Render(c, http.StatusOK, templates.ProblemEdit(templates.ProblemEditViewModel{
			Categories: categoriesList,
			Visible:    prob.Visible,
			Tags:       tags,
			Author:     prob.Author,
		}))
	}
}

func PostProblemEdit(ps njudge.Problems, cs njudge.Categories) echo.HandlerFunc {
	type Request struct {
		Category int    `form:"category"`
		Visible  string `form:"visible"`
		Author   string `form:"author"`
	}
	return func(c echo.Context) error {
		if u := c.Get(templates.UserContextKey).(*njudge.User); u.Role != "admin" {
			return echo.NotFoundHandler(c)
		}
		var (
			err  error
			data Request
		)
		if err := c.Bind(&data); err != nil {
			return err
		}
		prob := c.Get(ProblemContextKey).(njudge.Problem)
		if data.Category == -1 {
			prob.Category = nil
		} else {
			prob.Category, err = cs.Get(c.Request().Context(), data.Category)
			if err != nil {
				return err
			}
		}
		prob.Visible = data.Visible == "on"
		prob.Author = data.Author
		err = ps.Update(c.Request().Context(), prob, njudge.Fields(
			njudge.ProblemFields.Category, njudge.ProblemFields.Visible,
			njudge.ProblemFields.Author,
		))
		if err != nil {
			return err
		}

		return c.Redirect(http.StatusFound, "./edit")
	}
}

func GetProblemPDF() echo.HandlerFunc {
	return func(c echo.Context) error {
		storedData := c.Get(ProblemStoredDataContextKey).(njudge.ProblemStoredData)

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

func GetProblemAttachment() echo.HandlerFunc {
	return func(c echo.Context) error {
		storedData := c.Get(ProblemStoredDataContextKey).(njudge.ProblemStoredData)
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
		storedData := c.Get(ProblemStoredDataContextKey).(njudge.ProblemStoredData)

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

		userCache := make(map[int]*njudge.User)
		maxScore := ss.Feedback[0].MaxScore()
		isScored := ss.FeedbackType == problems.FeedbackIOI || ss.FeedbackType == problems.FeedbackLazyIOI
		vm := templates.ProblemRanklistViewModel{
			ScoredProblem: isScored,
			ScoresRows:    nil,
		}
		slices.SortFunc(submissions.Submissions, func(a, b njudge.Submission) int {
			if a.Verdict == njudge.VerdictAC && b.Verdict == njudge.VerdictAC {
				return cmp.Compare(a.ID, b.ID)
			}
			if a.Verdict == njudge.VerdictAC {
				return -1
			}
			if b.Verdict == njudge.VerdictAC {
				return 1
			}
			return -cmp.Compare(a.Score, b.Score)
		})
		for ind := range submissions.Submissions {
			if _, ok := userCache[submissions.Submissions[ind].UserID]; !ok {
				u, err := users.Get(c.Request().Context(), submissions.Submissions[ind].UserID)
				if err != nil {
					return err
				}
				userCache[submissions.Submissions[ind].UserID] = u

				vm.ScoresRows = append(vm.ScoresRows, templates.ProblemRanklistRow{
					SubmissionID: submissions.Submissions[ind].ID,
					Name:         u.Name,
					Text: fmt.Sprintf("%s/%s",
						strconv.FormatFloat(float64(submissions.Submissions[ind].Score), 'f', -1, 64),
						strconv.FormatFloat(maxScore, 'f', -1, 64),
					),
					Solved:  submissions.Submissions[ind].Verdict == njudge.VerdictAC,
					SortKey: -int64(submissions.Submissions[ind].Score * 1000000),
				})
			}

			u := userCache[submissions.Submissions[ind].UserID]
			if submissions.Submissions[ind].Verdict == njudge.VerdictAC {
				vm.TimeRows = append(vm.TimeRows, templates.ProblemRanklistRow{
					SubmissionID: submissions.Submissions[ind].ID,
					Name:         u.Name,
					Text:         fmt.Sprintf("%d ms", submissions.Submissions[ind].Status.Feedback[0].MaxTimeSpent()/time.Millisecond),
					SortKey:      int64(submissions.Submissions[ind].Status.Feedback[0].MaxTimeSpent()),
				})
				vm.SizeRows = append(vm.SizeRows, templates.ProblemRanklistRow{
					SubmissionID: submissions.Submissions[ind].ID,
					Name:         u.Name,
					Text:         fmt.Sprintf("%d B", len(submissions.Submissions[ind].Source)),
					SortKey:      int64(len(submissions.Submissions[ind].Source)),
				})
				maxMemory, memFormat := submissions.Submissions[ind].Status.Feedback[0].MaxMemoryUsage(), ""
				if maxMemory < 16*memory.MiB {
					memFormat = fmt.Sprintf("%.02f KiB", float64(maxMemory)/float64(memory.KiB))
				} else {
					memFormat = fmt.Sprintf("%.02f MiB", float64(maxMemory)/float64(memory.MiB))
				}
				vm.MemRows = append(vm.MemRows, templates.ProblemRanklistRow{
					SubmissionID: submissions.Submissions[ind].ID,
					Name:         u.Name,
					Text:         memFormat,
					SortKey:      int64(maxMemory),
				})
			}
		}

		compareFunc := func(a, b templates.ProblemRanklistRow) int {
			if cmp.Compare(a.SortKey, b.SortKey) == 0 {
				return cmp.Compare(a.SubmissionID, b.SubmissionID)
			}
			return cmp.Compare(a.SortKey, b.SortKey)
		}
		trimUsers := func(lst []templates.ProblemRanklistRow) []templates.ProblemRanklistRow {
			res := make([]templates.ProblemRanklistRow, 0, 5)
			hadUser := make(map[string]struct{})
			for i := range lst {
				if _, ok := hadUser[lst[i].Name]; ok || len(res) == 5 {
					continue
				}
				res = append(res, lst[i])
				hadUser[lst[i].Name] = struct{}{}
			}
			return res
		}

		slices.SortFunc(vm.ScoresRows, compareFunc)
		slices.SortFunc(vm.TimeRows, compareFunc)
		vm.TimeRows = trimUsers(vm.TimeRows)
		slices.SortFunc(vm.MemRows, compareFunc)
		vm.MemRows = trimUsers(vm.MemRows)
		slices.SortFunc(vm.SizeRows, compareFunc)
		vm.SizeRows = trimUsers(vm.SizeRows)

		c.Set(templates.TitleContextKey, tr.Translate("Results - %s (%s)", tr.TranslateContent(storedData.Titles()).String(), storedData.Name()))
		return templates.Render(c, http.StatusOK, templates.ProblemRanklist(vm))
	}
}

func GetProblemSubmit() echo.HandlerFunc {
	return func(c echo.Context) error {
		tr := c.Get(i18n.TranslatorContextKey).(i18n.Translator)

		p := c.Get(ProblemContextKey).(njudge.Problem)
		storedData := c.Get(ProblemStoredDataContextKey).(njudge.ProblemStoredData)
		info := c.Get(ProblemInfoContextKey).(njudge.ProblemInfo)

		title := tr.TranslateContent(storedData.Titles()).String()
		vm := templates.ProblemSubmitViewModel{
			Problemset: p.Problemset,
			Name:       p.Problem,
			Title:      title,
			UserInfo:   info.UserInfo,
			Languages:  storedData.Languages(),
		}

		c.Set(templates.TitleContextKey, tr.Translate("Submit - %s (%s)", title, storedData.Name()))

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

		prob := c.Get(ProblemContextKey).(njudge.Problem)
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

		c.Set(templates.TitleContextKey, tr.Translate("Submissions - %s (%s)", tr.TranslateContent(storedData.Titles()).String(), storedData.Name()))
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

		u := c.Get(templates.UserContextKey).(*njudge.User)
		if u == nil {
			return c.NoContent(http.StatusUnauthorized)
		}

		pr := c.Get(ProblemContextKey).(njudge.Problem)
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

		u := c.Get(templates.UserContextKey).(*njudge.User)
		if u == nil {
			return c.NoContent(http.StatusUnauthorized)
		}

		pr := c.Get(ProblemContextKey).(njudge.Problem)
		if err := tgs.Delete(c.Request().Context(), data.TagID, pr.ID, u.ID); err != nil {
			return err
		}

		return c.Redirect(http.StatusFound, c.Echo().Reverse("getProblemMain", pr.Problemset, pr.Problem))
	}
}
