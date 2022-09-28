package problemset

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"mime"
	"net/http"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"github.com/mraron/njudge/pkg/web/helpers"
	"github.com/mraron/njudge/pkg/web/helpers/config"
	"github.com/mraron/njudge/pkg/web/helpers/i18n"
	"github.com/mraron/njudge/pkg/web/models"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/mraron/njudge/pkg/problems"
	"github.com/mraron/njudge/pkg/problems/config/polygon"
	"github.com/volatiletech/sqlboiler/v4/queries"
	. "github.com/volatiletech/sqlboiler/v4/queries/qm"
	"golang.org/x/exp/slices"
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
	Tags         []*models.Tag
	Submissions  []*models.Submission
}

func topCategoryLink(cat int, DB *sqlx.DB) (helpers.Link, error) {
	var (
		category *models.ProblemCategory
		err      error
	)

	orig := cat

	for {
		category, err = models.ProblemCategories(Where("id = ?", cat)).One(DB)
		if err != nil {
			return helpers.Link{}, err
		}

		if !category.ParentID.Valid {
			break
		}
		cat = category.ParentID.Int
	}

	return helpers.Link{
		Text: category.Name,
		Href: "/task_archive#category" + strconv.Itoa(orig),
	}, nil
}

func lastLanguage(c echo.Context, DB *sqlx.DB) string {
	if res := c.Get("last_language"); res != nil {
		return c.Get("last_language").(string)
	}

	res := ""
	if u := c.Get("user").(*models.User); u != nil {
		sub, err := models.Submissions(Select("language"), Where("user_id = ?", u.ID), OrderBy("id DESC"), Limit(1)).One(DB)
		if err == nil {
			c.Set("last_language", sub.Language)
			res = sub.Language
		}
	}

	return res
}

func (p *Problem) FillFields(c echo.Context, DB *sqlx.DB, problemRel *models.ProblemRel) error {
	var err error
	p.SolverCount = problemRel.SolverCount
	if u := c.Get("user").(*models.User); u != nil {
		p.LastLanguage = lastLanguage(c, DB)
		p.SolvedStatus, err = helpers.HasUserSolved(DB, u, problemRel.Problemset, problemRel.Problem)
		if err != nil {
			return err
		}
		p.Submissions, err = models.Submissions(Where("problemset = ?", problemRel.Problemset), Where("problem = ?", problemRel.Problem), Where("user_id = ?", u.ID), OrderBy("id DESC"), Limit(5)).All(DB)
		if err != nil {
			return err
		}
	}

	if problemRel.CategoryID.Valid {
		p.CategoryId = problemRel.CategoryID.Int
		p.CategoryLink, err = topCategoryLink(p.CategoryId, DB)
		if err != nil {
			return err
		}
	}

	tags, err := models.Tags(InnerJoin("problem_tags pt on pt.tag_id = tags.id"), Where("pt.problem_id = ?", problemRel.ID)).All(DB)
	if err != nil {
		return err
	}
	p.Tags = tags

	return nil
}

func RenameMiddleware(problemStore problems.Store) func(echo.HandlerFunc) echo.HandlerFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			lst, err := problemStore.List()
			if err != nil {
				return err
			}

			if !slices.Contains(lst, c.Param("problem")) {
				for _, elem := range lst {
					if strings.HasSuffix(elem, "_"+c.Param("problem")) {
						return c.Redirect(http.StatusTemporaryRedirect, fmt.Sprintf("/problemset/%s/%s/", c.Param("name"), elem))
					}
				}
			}

			return next(c)
		}
	}
}

func GetProblem(DB *sqlx.DB, problemStore problems.Store) echo.HandlerFunc {
	return func(c echo.Context) error {
		name, problem := c.Param("name"), c.Param("problem")

		rel, err := models.ProblemRels(Where("problemset=?", name), Where("problem=?", problem)).One(DB)
		if err != nil {
			return err
		}

		if rel == nil {
			return echo.NewHTTPError(http.StatusNotFound, errors.New("no such problem in problemset"))
		}

		prob := problemStore.MustGet(problem)

		c.Set("title", fmt.Sprintf("Leírás - %s (%s)", i18n.TranslateContent("hungarian", prob.Titles()).String(), prob.Name()))

		p := Problem{Problem: prob}
		if err := p.FillFields(c, DB, rel); err != nil {
			return err
		}
		return c.Render(http.StatusOK, "problemset/problem/problem", p)
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

		switch p := p.(problems.ProblemWrapper).Problem.(type) {
		case polygon.Problem:
			if len(p.HTMLStatements()) == 0 || strings.HasSuffix(c.Param("file"), ".tex") || strings.HasSuffix(c.Param("file"), ".json") {
				return echo.NewHTTPError(http.StatusNotFound, ErrorFileNotFound)
			}

			if strings.HasSuffix(c.Param("file"), ".css") {
				fileLoc = filepath.Join(p.Path, "statements", ".html", p.HTMLStatements()[0].Locale(), c.Param("file"))
			} else {
				fileLoc = filepath.Join(p.Path, "statements", p.HTMLStatements()[0].Locale(), c.Param("file"))
			}

		default:
			return echo.NewHTTPError(http.StatusNotFound, ErrorFileNotFound)
		}

		return c.File(fileLoc)
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

		c.Set("title", fmt.Sprintf("Eredmények - %s (%s)", i18n.TranslateContent("hungarian", prob.Titles()).String(), prob.Name()))
		return c.Render(http.StatusOK, "problemset/problem/ranklist", struct {
			Problem     problems.Problem
			Submissions []*models.Submission
		}{prob, sbs})
	}
}

func GetProblemSubmit(DB *sqlx.DB, problemStore problems.Store) echo.HandlerFunc {
	return func(c echo.Context) error {
		problem := c.Param("problem")
		prob, err := problemStore.Get(problem)
		if err != nil {
			if errors.Is(err, problems.ErrorProblemNotFound) {
				return echo.NewHTTPError(http.StatusNotFound, err)
			}

			return err
		}

		c.Set("title", fmt.Sprintf("Beküldés - %s (%s)", i18n.TranslateContent("hungarian", prob.Titles()).String(), prob.Name()))
		return c.Render(http.StatusOK, "problemset/problem/submit", Problem{
			Problem:      prob,
			LastLanguage: lastLanguage(c, DB),
		})
	}
}

func GetProblemStatus(DB *sqlx.DB, problemStore problems.Store) echo.HandlerFunc {
	return func(c echo.Context) error {
		ac := c.QueryParam("ac")
		problemset := c.Param("name")
		problem := c.Param("problem")
		prob, err := problemStore.Get(problem)
		if err != nil {
			if errors.Is(err, problems.ErrorProblemNotFound) {
				return echo.NewHTTPError(http.StatusNotFound, err)
			}

			return err
		}

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

		c.Set("title", fmt.Sprintf("Beküldések - %s (%s)", i18n.TranslateContent("hungarian", prob.Titles()).String(), prob.Name()))
		return c.Render(http.StatusOK, "problemset/problem/status", statusPage)
	}
}