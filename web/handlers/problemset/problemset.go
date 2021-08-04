package problemset

import (
	"errors"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/mraron/njudge/utils/problems"
	"github.com/mraron/njudge/web/helpers"
	"github.com/mraron/njudge/web/helpers/config"
	"github.com/mraron/njudge/web/helpers/pagination"
	"github.com/mraron/njudge/web/models"
	. "github.com/volatiletech/sqlboiler/v4/queries/qm"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

type ProblemList struct {
	Pages []pagination.Link
	Problems []Problem
	SolverSorter helpers.SortColumn
}

func GetProblemList(DB *sqlx.DB, problemStore problems.Store, u *models.User, page, perPage int, order QueryMod, query []QueryMod, qu url.Values) (*ProblemList, error) {
	ps, err := models.ProblemRels(append(append([]QueryMod{Limit(perPage), Offset((page-1)*perPage)}, query...), order)...).All(DB)
	if err != nil {
		return nil, err
	}

	cnt, err := models.ProblemRels(query...).Count(DB)
	if err != nil {
		return nil, err
	}

	pageCnt := (int(cnt)+perPage-1)/perPage
	pages := make([]pagination.Link, pageCnt+2)
	pages[0] = pagination.Link{"&laquo;", false, true, "#"}
	if page>1 {
		qu.Set("page", strconv.Itoa(page-1))

		pages[0].Disabled = false
		pages[0].Url = "?"+qu.Encode()
	}
	for i := 1; i < len(pages)-1; i++ {
		qu.Set("page", strconv.Itoa(i))
		pages[i] = pagination.Link{strconv.Itoa(i), false, false, "?"+qu.Encode()}
		if i==page {
			pages[i].Active = true
		}
	}
	pages[len(pages)-1] = pagination.Link{"&raquo;", false, true, "#"}
	if page<pageCnt {
		qu.Set("page", strconv.Itoa(page+1))

		pages[len(pages)-1].Disabled = false
		pages[len(pages)-1].Url = "?"+qu.Encode()
	}

	if page>len(pages) {
		return nil, errors.New("no such page")
	}

	problems := make([]Problem, len(ps))
	for i, p := range ps {
		problems[i].Problem, err = problemStore.Get(p.Problem)
		if err != nil {
			return nil, err
		}

		if u != nil {
			problems[i].SolvedStatus, err = helpers.HasUserSolved(DB, u, p.Problemset, p.Problem)
			if err != nil {
				return nil, err
			}
		}

		problems[i].SolverCount = p.SolverCount
		if p.CategoryID.Valid {
			cat := p.CategoryID.Int
			var category *models.ProblemCategory
			for {
				category, err = models.ProblemCategories(Where("id = ?", cat)).One(DB)
				if err != nil {
					return nil, err
				}

				if !category.ParentID.Valid {
					break
				}
				cat = category.ParentID.Int
			}

			problems[i].CategoryLink = helpers.Link{
				Text: category.Name,
				Href: "/task_archive#category"+strconv.Itoa(p.CategoryID.Int),
			}
		}
	}

	sortOrder := ""
	qu.Set("page", strconv.Itoa(page))
	if qu.Get("by") == "solver_count" {
		sortOrder = qu.Get("order")
		if qu.Get("order") == "DESC" {
			qu.Set("order", "ASC")
		}else {
			qu.Set("order", "")
			qu.Set("by", "")
		}
	}else {
		qu.Set("by", "solver_count")
		qu.Set("order", "DESC")
	}

	return &ProblemList{Pages: pages, Problems: problems, SolverSorter: helpers.SortColumn{sortOrder, "?"+qu.Encode()}}, nil
}

func GetList(DB *sqlx.DB, problemStore problems.Store) echo.HandlerFunc {
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

		problemList, err := GetProblemList(DB, problemStore, u, page, 20, OrderBy(by+" "+order), []QueryMod{Where("problemset=?", problemSet)}, c.Request().URL.Query())
		if err != nil {
			return err
		}

		return c.Render(http.StatusOK, "problemset/list", problemList)
	}
}

func GetStatus(DB* sqlx.DB) echo.HandlerFunc {
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

		return c.Render(http.StatusOK, "status.gohtml", statusPage)
	}
}

func PostSubmit(cfg config.Server, DB* sqlx.DB, problemStore problems.Store) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			u   *models.User
			err error
			id  int
			p   problems.Problem
		)

		if u = c.Get("user").(*models.User); u == nil {
			return c.Render(http.StatusForbidden, "message.html", "Előbb lépj be.")
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

		fileHeader, err := c.FormFile("source")
		if err != nil {
			return err
		}

		f, err := fileHeader.Open()
		if err != nil {
			return err
		}

		contents, err := ioutil.ReadAll(f)
		if err != nil {
			return err
		}

		if id, err = helpers.Submit(cfg, DB, problemStore, u.ID, c.Get("problemset").(string), problemStore.MustGet(c.FormValue("problem")).Name(), languageName, contents); err != nil {
			return err
		}

		return c.Redirect(http.StatusFound, "/problemset/status/#submission"+strconv.Itoa(id))
	}
}

