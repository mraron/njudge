package problemset

import (
	"context"
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/mraron/njudge/utils/problems"
	"github.com/mraron/njudge/web/helpers"
	"github.com/mraron/njudge/web/helpers/config"
	"github.com/mraron/njudge/web/models"
	"github.com/volatiletech/sqlboiler/v4/queries"
	. "github.com/volatiletech/sqlboiler/v4/queries/qm"
	"io/ioutil"
	"net/http"
	"strconv"
)

func GetList(DB *sqlx.DB, problemStore problems.Store) echo.HandlerFunc {
	return func(c echo.Context) error {
		u := c.Get("user").(*models.User)

		problemSet := c.Param("name")
		problemLst, err := models.ProblemRels(Where("problemset=?", problemSet), OrderBy("id DESC")).All(DB)

		if err != nil {
			return helpers.InternalError(c, err, "Belső hiba.")
		}

		if len(problemLst) == 0 {
			return c.Render(http.StatusNotFound, "404.gohtml", "Nem található.")
		}

		lst := make([]Problem, len(problemLst))

		for i := 0; i < len(problemLst); i ++ {
			cnt := struct {
				Count int64
			}{0}

			err := queries.Raw("SELECT COUNT(DISTINCT user_id) FROM submissions WHERE problemset=$1 and problem=$2 and verdict=0", problemSet, problemLst[i].Problem).Bind(context.TODO(), DB, &cnt)
			if err != nil && !errors.Is(err, sql.ErrNoRows) {
				return helpers.InternalError(c, err, "Belső hiba.")
			}

			solvedStatus, err := helpers.HasUserSolved(DB, u, problemSet, problemLst[i].Problem)
			if err != nil {
				return helpers.InternalError(c, err, "Belső hiba.")
			}

			lst[i] = Problem{Problem: problemStore.MustGet(problemLst[i].Problem), SolverCount: int(cnt.Count), SolvedStatus: solvedStatus}
		}

		return c.Render(http.StatusOK, "problemset_list.gohtml", struct {
			Lst []Problem
		}{lst})
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

		query := []QueryMod{}
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
			return helpers.InternalError(c, err, "Belső hiba")
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
			return c.Render(http.StatusForbidden, "error.gohtml", "Előbb lépj be.")
		}

		problemName := c.FormValue("problem")
		if has, _ := problemStore.Has(problemName); !has {
			return c.Render(http.StatusOK, "error.gohtml", "Hibás feladatazonosító.")
		}else {
			p, _ = problemStore.Get(problemName)
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
			return helpers.InternalError(c, err, "Belső hiba #0")
		}

		f, err := fileHeader.Open()
		if err != nil {
			return helpers.InternalError(c, err, "Belső hiba #1")
		}

		contents, err := ioutil.ReadAll(f)
		if err != nil {
			return helpers.InternalError(c, err, "Belső hiba #2")
		}

		if id, err = helpers.Submit(cfg, DB, problemStore, u.ID, c.Get("problemset").(string), problemStore.MustGet(c.FormValue("problem")).Name(), languageName, contents); err != nil {
			return helpers.InternalError(c, err, "Belső hiba #4")
		}

		return c.Redirect(http.StatusFound, "/problemset/status#submission"+strconv.Itoa(id))
	}
}

