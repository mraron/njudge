package problemset

import (
	"bytes"
	"context"
	"fmt"
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
	"io"
	"mime"
	"net/http"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

type Problem struct {
	problems.Problem
	SolverCount int
	SolvedStatus int
	LastLanguage string
}

func GetProblem(DB *sqlx.DB, problemStore problems.Store) echo.HandlerFunc {
	return func(c echo.Context) error {
		name, problem := c.Param("name"), c.Param("problem")

		lst, err := models.ProblemRels(Where("problemset=?", name)).All(DB)
		if err != nil {
			return helpers.InternalError(c, err, "Belső hiba.")
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
			return c.JSON(http.StatusNotFound, nil)
		}

		lastLanguage := ""
		if u := c.Get("user").(*models.User); u != nil {
			fmt.Println(u)
			sub, err := models.Submissions(Select("language"), Where("user_id = ?", u.ID), OrderBy("id DESC"), Limit(1)).One(DB)
			if err == nil {
				lastLanguage = sub.Language
			}
		}

		return c.Render(http.StatusOK, "problemset_problem.gohtml", Problem{Problem:problemStore.MustGet(problem), LastLanguage: lastLanguage})
	}
}

func GetProblemPDF(problemStore problems.Store) echo.HandlerFunc {
	return func(c echo.Context) error {
		p, lang := problemStore.MustGet(c.Param("problem")), c.Param("language")

		if p == nil {
			return c.String(http.StatusNotFound, "no such problem")
		}

		if len(p.PDFStatements()) == 0 {
			return c.String(http.StatusNotFound, "no pdf statement")
		}

		return c.Blob(http.StatusOK, "application/pdf", i18n.TranslateContent(lang, p.PDFStatements()).Contents)
	}
}

func GetProblemFile(cfg config.Server, problemStore problems.Store) echo.HandlerFunc {
	return func(c echo.Context) error {
		p := problemStore.MustGet(c.Param("problem"))

		if p == nil {
			return c.String(http.StatusNotFound, "not found")
		}

		fileLoc := ""

		switch p.(type) {
		case polygon.Problem:
			if len(p.HTMLStatements()) == 0 {
				return c.String(http.StatusNotFound, "file not found")
			}

			if strings.HasSuffix(c.Param("file"), ".css") {
				fileLoc = filepath.Join(cfg.ProblemsDir, p.Name(), "statements", ".html", p.HTMLStatements()[0].Locale, c.Param("file"))
			} else {
				fileLoc = filepath.Join(cfg.ProblemsDir, p.Name(), "statements", p.HTMLStatements()[0].Locale, c.Param("file"))
			}

		default:
			return c.String(http.StatusNotFound, "not found")
		}

		return c.Attachment(fileLoc, c.Param("file"))
	}
}

func GetProblemAttachment(problemStore problems.Store) echo.HandlerFunc {
	return func(c echo.Context) error {
		p, attachment := problemStore.MustGet(c.Param("problem")), c.Param("attachment")
		if p == nil {
			return c.String(http.StatusNotFound, "no such problem")
		}

		for _, val := range p.Attachments() {
			if val.Name == attachment {
				c.Response().Header().Set("Content-Disposition", "attachment; filename="+val.Name)
				c.Response().Header().Set("Content-Type", mime.TypeByExtension(filepath.Ext(val.Name)))
				c.Response().Header().Set("Content-Length", strconv.Itoa(len(val.Contents)))

				io.Copy(c.Response(), bytes.NewReader(val.Contents))

				return c.NoContent(http.StatusOK)
			}
		}

		return c.String(http.StatusNotFound, "no such attachment")
	}
}

func GetProblemRanklist(DB *sqlx.DB, problemStore problems.Store) echo.HandlerFunc {
	return func(c echo.Context) error {
		problemSet := c.Param("name")
		problem := c.Param("problem")
		prob := problemStore.MustGet(problem)

		sbs := make([]*models.Submission, 0)

		//@TODO
		if err := queries.Raw("SELECT DISTINCT ON (s1.user_id) s1.* FROM (SELECT s1.user_id, MAX(s1.score) as score FROM submissions s1 WHERE problemset=$1 AND problem=$2 GROUP BY s1.user_id) s2 INNER JOIN submissions s1 ON s1.user_id=s2.user_id AND s1.score=s2.score AND s1.problemset=$1 AND s1.problem=$2", problemSet, problem).Bind(context.TODO(), DB, &sbs); err != nil {
			return helpers.InternalError(c, err, "hiba")
		}

		sort.Slice(sbs, func(i, j int) bool {
			return sbs[i].Score.Float32 > sbs[j].Score.Float32
		})

		return c.Render(http.StatusOK, "problemset_problem_ranklist.gohtml", struct {
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
		if  err != nil || page<=0 {
			page = 1
		}

		query := []QueryMod{}
		if ac == "1" {
			query = []QueryMod{Where("verdict = 0"), Where("problem = ?", problem), Where("problemset = ?", problemset)}
		} else {
			query = []QueryMod{Where("problem = ?", problem), Where("problemset = ?", problemset)}
		}

		statusPage, err := helpers.GetStatusPage(DB, page, 20, OrderBy("id DESC"), query, c.Request().URL.Query())
		if err != nil {
			return helpers.InternalError(c, err, "Belső hiba")
		}

		return c.Render(http.StatusOK, "problemset_problem_status.gohtml", statusPage)
	}
}