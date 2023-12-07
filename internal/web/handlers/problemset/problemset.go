package problemset

import (
	"fmt"
	"io"
	"net/http"
	"sort"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/mraron/njudge/internal/njudge"
	"github.com/mraron/njudge/internal/web/helpers/i18n"
	"github.com/mraron/njudge/internal/web/helpers/pagination"
	"github.com/mraron/njudge/internal/web/helpers/ui"
	"github.com/mraron/njudge/pkg/problems"
)

type CategoryFilterOption struct {
	Name     string
	Value    string
	Selected bool
}

type Problem struct {
	njudge.Problem
	njudge.ProblemStoredData
	njudge.ProblemInfo

	CategoryLink ui.Link
}

type ProblemList struct {
	Pages        []pagination.Link
	Problems     []Problem
	SolverSorter ui.SortColumn

	Filtered bool

	TitleFilter           string
	TagsFilter            string
	CategoryFilterOptions []CategoryFilterOption
}

func GetProblemList(store problems.Store, ps njudge.Problems, cs njudge.Categories, problemListQuery njudge.ProblemListQuery, pinfo njudge.ProblemInfoQuery) echo.HandlerFunc {
	type request struct {
		Page  int `query:"page"`
		Order njudge.SortDirection
		By    njudge.ProblemSortField

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

		data.Order, data.By = njudge.SortDESC, njudge.ProblemSortFieldID
		if c.QueryParam("by") == "solver_count" {
			data.By = njudge.ProblemSortFieldSolverCount
		}
		if c.QueryParam("order") == "ASC" {
			data.Order = njudge.SortASC
		}

		listRequest := njudge.ProblemListRequest{
			Problemset:  data.Problemset,
			Page:        data.Page,
			PerPage:     20,
			SortDir:     data.Order,
			SortField:   data.By,
			TitleFilter: data.TitleFilter,
		}

		if data.TagFilter != "" {
			listRequest.TagFilter = strings.Split(data.TagFilter, ",")
		}

		if data.CategoryFilter != 0 {
			if data.CategoryFilter == -1 {
				listRequest.CategoryFilter = njudge.NewCategoryEmptyFilter()
			} else {
				listRequest.CategoryFilter = njudge.NewCategoryIDFilter(data.CategoryFilter)
			}
		}

		problemList, err := problemListQuery.GetProblemList(c.Request().Context(), listRequest)
		if err != nil {
			return err
		}

		u := *c.Request().URL
		links, err := pagination.Links(problemList.PaginationData.Page, problemList.PaginationData.PerPage, int64(problemList.PaginationData.Count), u.Query())
		if err != nil {
			return err
		}
		result := ProblemList{
			Pages: links,
		}

		for ind := range problemList.Problems {
			p, err := ps.Get(c.Request().Context(), problemList.Problems[ind].ID)
			if err != nil {
				return err
			}
			info, err := pinfo.GetProblemData(c.Request().Context(), p.ID, c.Get("userID").(int))
			if err != nil {
				return err
			}
			data, err := p.WithStoredData(store)
			if err != nil {
				return err
			}

			result.Problems = append(result.Problems, Problem{
				Problem:           *p,
				ProblemInfo:       *info,
				ProblemStoredData: data,
			})
		}

		sortOrder, u := "", *c.Request().URL
		qu := u.Query()
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
		result.CategoryFilterOptions = []CategoryFilterOption{
			{Name: "-"},
		}

		emptySelected := false
		if data.CategoryFilter == -1 {
			emptySelected = true
		}
		result.CategoryFilterOptions = append(result.CategoryFilterOptions, CategoryFilterOption{
			Name:     tr.Translate("No category"),
			Value:    "-1",
			Selected: emptySelected,
		})

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

		var getCategoryNameRec func(int) string
		getCategoryNameRec = func(id int) string {
			if _, ok := par[id]; !ok {
				return categoryNameByID[id]
			} else {
				return getCategoryNameRec(par[id]) + " -- " + categoryNameByID[id]
			}
		}

		for ind := range categories {
			curr := CategoryFilterOption{
				Name:     getCategoryNameRec(categories[ind].ID),
				Value:    strconv.Itoa(categories[ind].ID),
				Selected: false,
			}

			if strconv.Itoa(categories[ind].ID) == c.QueryParam("category") {
				curr.Selected = true
			}

			result.CategoryFilterOptions = append(result.CategoryFilterOptions, curr)
		}

		sort.Slice(result.CategoryFilterOptions, func(i, j int) bool {
			return result.CategoryFilterOptions[i].Name < result.CategoryFilterOptions[j].Name
		})

		for ind := range result.Problems {
			if result.Problems[ind].Category != nil {
				cid := result.Problems[ind].Category.ID
				for {
					if _, ok := par[cid]; ok {
						cid = par[cid]
					} else {
						break
					}
				}

				result.Problems[ind].CategoryLink = ui.Link{
					Text: categoryNameByID[cid],
					Href: fmt.Sprintf("/task_archive#category%d", result.Problems[ind].Category.ID),
				}
			}
		}

		c.Set("title", tr.Translate("Problems"))
		return c.Render(http.StatusOK, "problemset/list", result)
	}
}

type StatusPage struct {
	Pages       []pagination.Link
	Submissions []njudge.Submission
}

func GetStatus(slist njudge.SubmissionListQuery) echo.HandlerFunc {
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

		statusReq := njudge.SubmissionListRequest{
			Page:       data.Page,
			PerPage:    20,
			SortDir:    njudge.SortDESC,
			SortField:  njudge.SubmissionSortFieldID,
			Problemset: data.Problemset,
			Problem:    data.Problem,
			UserID:     data.UserID,
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
		links, err := pagination.LinksWithCountLimit(submissionList.PaginationData.Page, submissionList.PaginationData.PerPage, int64(submissionList.PaginationData.Count), qu, 5)
		if err != nil {
			return err
		}

		result := StatusPage{
			Submissions: submissionList.Submissions,
			Pages:       links,
		}

		c.Set("title", tr.Translate("Submissions"))
		return c.Render(http.StatusOK, "status.gohtml", result)
	}
}

func PostSubmit(subService njudge.SubmitService) echo.HandlerFunc {
	type request struct {
		Problemset     string `param:"name"`
		ProblemName    string `form:"problem"`
		LanguageName   string `form:"language"`
		SubmissionCode string `form:"submissionCode"`
	}
	return func(c echo.Context) error {
		u := c.Get("user").(*njudge.User)

		data := request{}
		if err := c.Bind(&data); err != nil {
			return err
		}

		code := data.SubmissionCode
		if len(code) == 0 {
			fileHeader, err := c.FormFile("source")
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

		sub, err := subService.Submit(c.Request().Context(), njudge.SubmitRequest{
			UserID:     u.ID,
			Problemset: data.Problemset,
			Problem:    data.ProblemName,
			Language:   data.LanguageName,
			Source:     []byte(code),
		})
		if err != nil {
			return err
		}

		return c.Redirect(http.StatusFound, c.Echo().Reverse("getProblemsetStatus")+"#submission"+strconv.Itoa(sub.ID))
	}
}
