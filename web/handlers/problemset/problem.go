package problemset

import (
	"bytes"
	"context"
	"errors"
	"io"
	"mime"
	"net/http"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/mraron/njudge/utils/problems"
	"github.com/mraron/njudge/utils/problems/config/polygon"
	"github.com/mraron/njudge/web/helpers"
	"github.com/mraron/njudge/web/helpers/config"
	"github.com/mraron/njudge/web/helpers/i18n"
	"github.com/mraron/njudge/web/models"
	"github.com/volatiletech/sqlboiler/v4/queries"
	. "github.com/volatiletech/sqlboiler/v4/queries/qm"
)

var (
	ErrorFileNotFound = errors.New("file not found")
)

type Problem struct {
	problems.Problem
	SolverCount  int
	SolvedStatus helpers.SolvedStatus
	LastLanguage string
	CategoryLink helpers.Link
	CategoryId   int
}

func GetProblem(DB *sqlx.DB, problemStore problems.Store) echo.HandlerFunc {
	return func(c echo.Context) error {
		name, problem := c.Param("name"), c.Param("problem")

		lst, err := models.ProblemRels(Where("problemset=?", name)).All(DB)
		if err != nil {
			return err
		}

		ok := false
		for _, val := range lst {
			if val.Problem == problem {
				ok = true
			}

			if ok {
				break
			}
		}

		if !ok {
			return echo.NewHTTPError(http.StatusNotFound, errors.New("no such problem in problemset"))
		}

		lastLanguage := ""
		if u := c.Get("user").(*models.User); u != nil {
			sub, err := models.Submissions(Select("language"), Where("user_id = ?", u.ID), OrderBy("id DESC"), Limit(1)).One(DB)
			if err == nil {
				lastLanguage = sub.Language
			}
		}

		return c.Render(http.StatusOK, "problemset/problem/problem", Problem{Problem: problemStore.MustGet(problem), LastLanguage: lastLanguage})
	}
}

func GetProblemPDF(problemStore problems.Store) echo.HandlerFunc {
	return func(c echo.Context) error {
		p, err := problemStore.Get(c.Param("problem"))
		lang := c.Param("language")

		if err != nil {
			if errors.Is(err, problems.ErrorProblemNotFound) {
				return echo.NewHTTPError(http.StatusNotFound, err)
			}

			return err
		}

		if len(p.PDFStatements()) == 0 {
			return echo.NewHTTPError(http.StatusNotFound, ErrorFileNotFound)
		}

		dat, err := i18n.TranslateContent(lang, p.PDFStatements()).Value()
		if err != nil {
			return err
		}

		return c.Blob(http.StatusOK, "application/pdf", dat)
	}
}

func GetProblemFile(cfg config.Server, problemStore problems.Store) echo.HandlerFunc {
	return func(c echo.Context) error {
		p, err := problemStore.Get(c.Param("problem"))

		if err != nil {
			if errors.Is(err, problems.ErrorProblemNotFound) {
				return echo.NewHTTPError(http.StatusNotFound, err)
			}

			return err
		}

		fileLoc := ""

		switch p.(type) {
		case polygon.Problem:
			if len(p.HTMLStatements()) == 0 {
				return echo.NewHTTPError(http.StatusNotFound, ErrorFileNotFound)
			}

			//@TODO what the fuck is this? how does polygon do it ATM
			if strings.HasSuffix(c.Param("file"), ".css") {
				fileLoc = filepath.Join(cfg.ProblemsDir, p.Name(), "statements", ".html", p.HTMLStatements()[0].Locale(), c.Param("file"))
			} else {
				fileLoc = filepath.Join(cfg.ProblemsDir, p.Name(), "statements", p.HTMLStatements()[0].Locale(), c.Param("file"))
			}

		default:
			return echo.NewHTTPError(http.StatusNotFound, ErrorFileNotFound)
		}

		return c.Attachment(fileLoc, c.Param("file"))
	}
}

func GetProblemAttachment(problemStore problems.Store) echo.HandlerFunc {
	return func(c echo.Context) error {
		p, err := problemStore.Get(c.Param("problem"))
		attachment := c.Param("attachment")

		if err != nil {
			if errors.Is(err, problems.ErrorProblemNotFound) {
				return echo.NewHTTPError(http.StatusNotFound, err)
			}

			return err
		}

		for _, val := range p.Attachments() {
			if val.Name() == attachment {
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

		return echo.NewHTTPError(http.StatusNotFound, ErrorFileNotFound)
	}
}

func GetProblemRanklist(DB *sqlx.DB, problemStore problems.Store) echo.HandlerFunc {
	return func(c echo.Context) error {
		problemSet := c.Param("name")
		problem := c.Param("problem")
		prob, err := problemStore.Get(problem)
		if err != nil {
			if errors.Is(err, problems.ErrorProblemNotFound) {
				return echo.NewHTTPError(http.StatusNotFound, err)
			}

			return err
		}

		sbs := make([]*models.Submission, 0)

		//@TODO something better?
		if err := queries.Raw("SELECT DISTINCT ON (s1.user_id) s1.* FROM (SELECT s1.user_id, MAX(s1.score) as score FROM submissions s1 WHERE problemset=$1 AND problem=$2 GROUP BY s1.user_id) s2 INNER JOIN submissions s1 ON s1.user_id=s2.user_id AND s1.score=s2.score AND s1.problemset=$1 AND s1.problem=$2", problemSet, problem).Bind(context.TODO(), DB, &sbs); err != nil {
			return err
		}

		sort.Slice(sbs, func(i, j int) bool {
			return sbs[i].Score.Float32 > sbs[j].Score.Float32
		})

		return c.Render(http.StatusOK, "problemset/problem/ranklist", struct {
			Problem     problems.Problem
			Submissions []*models.Submission
		}{prob, sbs})
	}
}

func GetProblemStatus(DB *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		ac := c.QueryParam("ac")
		problemset := c.Param("name")
		problem := c.Param("problem")
		page, err := strconv.Atoi(c.QueryParam("page"))
		if err != nil || page <= 0 {
			page = 1
		}

		var query []QueryMod
		if ac == "1" {
			query = []QueryMod{Where("verdict = 0"), Where("problem = ?", problem), Where("problemset = ?", problemset)}
		} else {
			query = []QueryMod{Where("problem = ?", problem), Where("problemset = ?", problemset)}
		}

		statusPage, err := helpers.GetStatusPage(DB, page, 20, OrderBy("id DESC"), query, c.Request().URL.Query())
		if err != nil {
			return err
		}

		return c.Render(http.StatusOK, "problemset/problem/status", statusPage)
	}
}
