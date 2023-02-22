package problemset

import (
	"github.com/mraron/njudge/internal/web/handlers/problemset/problem"
	"github.com/mraron/njudge/internal/web/helpers"
	"github.com/mraron/njudge/internal/web/helpers/config"
	"github.com/mraron/njudge/internal/web/helpers/i18n"
	"github.com/mraron/njudge/internal/web/helpers/pagination"
	"github.com/mraron/njudge/internal/web/models"
	"io"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"unicode"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/mraron/njudge/pkg/problems"
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

type ProblemList struct {
	Pages        []pagination.Link
	Problems     []problem.Problem
	SolverSorter helpers.SortColumn

	Filtered        bool
	TitleFilter     string
	TagsFilter      string
	CategoryFilters []CategoryFilter
}

func getProblemList(c echo.Context, DB *sqlx.DB, problemStore problems.Store, u *models.User, page, perPage int, order QueryMod, query []QueryMod, qu url.Values) (*ProblemList, error) {
	ps, err := models.ProblemRels(append(append([]QueryMod{Limit(perPage), Offset((page - 1) * perPage)}, query...), order)...).All(DB)
	if err != nil {
		return nil, err
	}

	cnt, err := models.ProblemRels(query...).Count(DB)
	if err != nil {
		return nil, err
	}

	pages, err := pagination.Links(page, perPage, cnt, qu)
	if err != nil {
		return nil, err
	}

	problemsList := make([]problem.Problem, len(ps))
	for i, p := range ps {
		problemsList[i].Problem, err = problemStore.Get(p.Problem)
		if err != nil {
			return nil, err
		}
		problemsList[i].ProblemRel = p

		if err := problemsList[i].FillFields(c, DB); err != nil {
			return nil, err
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

	return &ProblemList{Pages: pages, Problems: problemsList, SolverSorter: helpers.SortColumn{sortOrder, "?" + qu.Encode()}}, nil
}

func GetProblemList(DB *sqlx.DB, problemStore problems.Store) echo.HandlerFunc {
	return func(c echo.Context) error {
		u := c.Get("user").(*models.User)

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

			rels, err := models.ProblemRels().All(DB)
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

		cats, err := models.ProblemCategories().All(DB)
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

			pars := []interface{}{}
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
				InnerJoin("tags t on pt.tag_id = t.id"), WhereIn("t.name in ?", lst...), GroupBy("problem_rels.id"), Having("COUNT(DISTINCT t.name) = ?", len(lst))).All(DB)
			if err != nil {
				return nil
			}

			lst = make([]interface{}, len(rels))
			for ind, val := range rels {
				lst[ind] = val.ID
			}

			qmods = append(qmods, WhereIn("id IN ?", lst...))
		}

		problemList, err := getProblemList(c, DB, problemStore, u, page, 20, OrderBy(by+" "+order), qmods, c.Request().URL.Query())
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

func GetStatus(DB *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		ac := c.QueryParam("ac")
		userID := c.QueryParam("user_id")
		problemset := c.QueryParam("problem_set")
		problem := c.QueryParam("problem")
		page, err := strconv.Atoi(c.QueryParam("page"))
		if err != nil || page <= 0 {
			page = 1
		}

		query := make([]QueryMod, 0)
		if problem != "" {
			query = append(query, Where("problem = ?", problem), Where("problemset = ?", problemset))
		}
		if ac == "1" {
			query = append(query, Where("verdict = 0"))
		}
		if userID != "" {
			query = append(query, Where("user_id = ?", userID))
		}

		statusPage, err := helpers.GetStatusPage(DB, page, 20, OrderBy("id DESC"), query, c.Request().URL.Query())
		if err != nil {
			return err
		}

		c.Set("title", "Beküldések")
		return c.Render(http.StatusOK, "status.gohtml", statusPage)
	}
}

func PostSubmit(cfg config.Server, DB *sqlx.DB, problemStore problems.Store) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			u   *models.User
			err error
			id  int
			p   problems.Problem
		)

		if u = c.Get("user").(*models.User); u == nil {
			return c.Render(http.StatusForbidden, "message", "Előbb lépj be.")
		}

		problemName := c.FormValue("problem")
		if p, err = problemStore.Get(problemName); err != nil {
			return err
		}

		languageName := c.FormValue("language")

		found := false
		for _, lang := range p.Languages() {
			if lang.Id() == languageName {
				found = true
				break
			}
		}

		if !found {
			return c.Render(http.StatusOK, "error.gohtml", "Hibás nyelvazonosító.")
		}

		code := []byte(c.FormValue("submissionCode"))
		if string(code) == "" {
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

			code = contents
		}

		if id, err = helpers.Submit(cfg, DB, problemStore, u.ID, c.Get("problemset").(string), problemStore.MustGet(c.FormValue("problem")).Name(), languageName, code); err != nil {
			return err
		}

		return c.Redirect(http.StatusFound, "/problemset/status/#submission"+strconv.Itoa(id))
	}
}
