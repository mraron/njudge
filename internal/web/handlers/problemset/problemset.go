package problemset

import (
	"context"
	"github.com/mraron/njudge/internal/web/domain/problem"
	"github.com/mraron/njudge/internal/web/helpers/config"
	"github.com/mraron/njudge/internal/web/helpers/i18n"
	"github.com/mraron/njudge/internal/web/helpers/pagination"
	"github.com/mraron/njudge/internal/web/helpers/ui"
	"github.com/mraron/njudge/internal/web/models"
	"github.com/mraron/njudge/internal/web/services"
	"github.com/mraron/njudge/pkg/problems"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"unicode"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	. "github.com/volatiletech/sqlboiler/v4/queries/qm"
	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

type CategoryFilter struct {
	Name     string
	Value    string
	Selected bool
}

type StatProblem struct {
	problem.Problem
	problem.StatsData
}

type ProblemList struct {
	Pages        []pagination.Link
	Problems     []StatProblem
	SolverSorter ui.SortColumn

	Filtered        bool
	TitleFilter     string
	TagsFilter      string
	CategoryFilters []CategoryFilter
}

func getProblemList(c echo.Context, DB *sqlx.DB, problemRepo problem.Repository, problemStatsService services.ProblemStatsService, page, perPage int, order QueryMod, query []QueryMod, qu url.Values) (*ProblemList, error) {
	ps, err := models.ProblemRels(append(append([]QueryMod{Limit(perPage), Offset((page - 1) * perPage)}, query...), order)...).All(context.TODO(), DB)
	if err != nil {
		return nil, err
	}

	cnt, err := models.ProblemRels(query...).Count(context.TODO(), DB)
	if err != nil {
		return nil, err
	}

	pages, err := pagination.Links(page, perPage, cnt, qu)
	if err != nil {
		return nil, err
	}

	problemsList := make([]StatProblem, len(ps))
	for i, p := range ps {
		res, err := problemRepo.Get(c.Request().Context(), p.ID)
		if err != nil {
			return nil, err
		}

		stats, err := problemStatsService.GetStatsData(c.Request().Context(), *res, c.Get("userID").(int))
		if err != nil {
			return nil, err
		}

		problemsList[i] = StatProblem{
			Problem:   *res,
			StatsData: *stats,
		}
	}

	sortOrder := ""
	qu.Set("page", strconv.Itoa(page))
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

	return &ProblemList{Pages: pages, Problems: problemsList, SolverSorter: ui.SortColumn{sortOrder, "?" + qu.Encode()}}, nil
}

func GetProblemList(DB *sqlx.DB, problemStore problems.Store, problemRepo problem.Repository, problemStatsService services.ProblemStatsService) echo.HandlerFunc {
	return func(c echo.Context) error {
		problemSet := c.Param("name")
		page, err := strconv.Atoi(c.QueryParam("page"))
		if err != nil || page <= 0 {
			page = 1
		}

		order, by := "DESC", "id"
		if c.QueryParam("by") == "solver_count" {
			by = "solver_count"
		}
		if c.QueryParam("order") == "ASC" {
			order = "ASC"
		}

		qmods := []QueryMod{Where("problemset=?", problemSet)}
		filtered := false

		if c.QueryParam("title") != "" {
			filtered = true

			rels, err := models.ProblemRels().All(context.TODO(), DB)
			if err != nil {
				return err
			}

			lst := make([]interface{}, 0, len(rels))
			for _, rel := range rels {
				p, err := problemStore.Get(rel.Problem)
				if err != nil {
					return err
				}

				curr := strings.ToLower(i18n.TranslateContent("hungarian", p.Titles()).String())
				want := strings.ToLower(c.QueryParam("title"))

				t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
				a, _, _ := transform.String(t, curr)
				b, _, _ := transform.String(t, want)

				if strings.Contains(a, b) {
					lst = append(lst, rel.Problem)
				}
			}

			qmods = append(qmods, WhereIn("problem in ?", lst...))
		}

		cats, err := models.ProblemCategories().All(context.TODO(), DB)
		if err != nil {
			return err
		}
		par := make(map[int]int)
		for _, cat := range cats {
			if cat.ParentID.Valid {
				par[cat.ID] = cat.ParentID.Int
			}
		}

		if c.QueryParam("category") != "" {
			filtered = true

			cid, err := strconv.Atoi(c.QueryParam("category"))
			if err != nil {
				return err
			}

			var pars []interface{}
			for _, cat := range cats {
				curr := cat.ID
				ok := false
				for {
					if curr == cid {
						ok = true
						break
					}

					if _, ok := par[curr]; ok {
						curr = par[curr]
					} else {
						break
					}
				}

				if ok {
					pars = append(pars, cat.ID)
				}
			}

			qmods = append(qmods, WhereIn("category_id in ?", pars...))
		}

		if c.QueryParam("tags") != "" {
			filtered = true

			tagNames := strings.Split(c.QueryParam("tags"), ",")
			lst := make([]interface{}, len(tagNames))
			for ind, val := range tagNames {
				lst[ind] = val
			}

			rels, err := models.ProblemRels(InnerJoin("problem_tags pt on pt.problem_id = problem_rels.id"),
				InnerJoin("tags t on pt.tag_id = t.id"), WhereIn("t.name in ?", lst...), GroupBy("problem_rels.id"), Having("COUNT(DISTINCT t.name) = ?", len(lst))).All(context.TODO(), DB)
			if err != nil {
				return nil
			}

			lst = make([]interface{}, len(rels))
			for ind, val := range rels {
				lst[ind] = val.ID
			}

			qmods = append(qmods, WhereIn("id IN ?", lst...))
		}

		problemList, err := getProblemList(c, DB, problemRepo, problemStatsService, page, 20, OrderBy(by+" "+order), qmods, c.Request().URL.Query())
		if err != nil {
			return err
		}

		problemList.Filtered = filtered
		problemList.TitleFilter = c.QueryParam("title")
		problemList.TagsFilter = c.QueryParam("tags")

		problemList.CategoryFilters = []CategoryFilter{{"-", "", false}}
		nameById := make(map[int]string)
		for _, cat := range cats {
			nameById[cat.ID] = cat.Name
		}

		var getName func(int) string
		getName = func(id int) string {
			if _, ok := par[id]; !ok {
				return nameById[id]
			} else {
				return getName(par[id]) + " -- " + nameById[id]
			}
		}

		for _, cat := range cats {
			curr := CategoryFilter{
				Name:     getName(cat.ID),
				Value:    strconv.Itoa(cat.ID),
				Selected: false,
			}

			if strconv.Itoa(cat.ID) == c.QueryParam("category") {
				curr.Selected = true
			}

			problemList.CategoryFilters = append(problemList.CategoryFilters, curr)
		}

		sort.Slice(problemList.CategoryFilters, func(i, j int) bool {
			return problemList.CategoryFilters[i].Name < problemList.CategoryFilters[j].Name
		})

		c.Set("title", "Feladatok")
		return c.Render(http.StatusOK, "problemset/list", problemList)
	}
}

func GetStatus(statusPageService services.StatusPageService) echo.HandlerFunc {
	type request struct {
		AC         string `query:"ac"`
		UserID     int    `query:"user_id"`
		Problemset string `query:"problem_set"`
		Problem    string `query:"problem"`
		Page       int    `query:"page"`
	}
	return func(c echo.Context) error {
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
			GETValues:  c.Request().URL.Query(),
		}

		if data.AC == "1" {
			ac := problems.VerdictAC
			statusReq.Verdict = &ac
		}

		statusPage, err := statusPageService.GetStatusPage(c.Request().Context(), statusReq)
		if err != nil {
			return err
		}

		c.Set("title", "Beküldések")
		return c.Render(http.StatusOK, "status.gohtml", statusPage)
	}
}

func PostSubmit(cfg config.Server, subService services.SubmitService) echo.HandlerFunc {
	type request struct {
		Problemset     string `param:"problemset"`
		ProblemName    string `form:"problem"`
		LanguageName   string `form:"language"`
		SubmissionCode []byte `form:"submissionCode"`
		SubmissionFile []byte `form:"source"`
	}
	return func(c echo.Context) error {
		u := c.Get("user").(*models.User)
		if u == nil {
			return c.Render(http.StatusForbidden, "message", "Előbb lépj be.")
		}

		data := request{}
		if err := c.Bind(&data); err != nil {
			return err
		}

		code := data.SubmissionCode
		if len(code) == 0 {
			code = data.SubmissionFile
		}

		sub, err := subService.Submit(c.Request().Context(), services.SubmitRequest{
			UserID:     u.ID,
			Problemset: data.Problemset,
			Problem:    data.ProblemName,
			Language:   data.LanguageName,
			Source:     code,
		})
		if err != nil {
			return err
		}

		return c.Redirect(http.StatusFound, "/problemset/status/#submission"+strconv.Itoa(sub.ID))
	}
}
