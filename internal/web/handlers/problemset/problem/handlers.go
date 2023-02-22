package problem

import (
	"bytes"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/mraron/njudge/internal/web/helpers"
	"github.com/mraron/njudge/internal/web/helpers/i18n"
	"github.com/mraron/njudge/internal/web/models"
	"github.com/mraron/njudge/pkg/problems"
	"github.com/volatiletech/sqlboiler/v4/queries"
	. "github.com/volatiletech/sqlboiler/v4/queries/qm"
	"io"
	"mime"
	"net/http"
	"path/filepath"
	"sort"
	"strconv"
)

func Get(DB *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		prob := c.Get("problem").(problems.Problem)

		c.Set("title", fmt.Sprintf("Leírás - %s (%s)", i18n.TranslateContent("hungarian", prob.Titles()).String(), prob.Name()))

		p := New(c)
		if err := p.FillFields(c, DB); err != nil {
			return err
		}
		return c.Render(http.StatusOK, "problemset/problem/problem", p)
	}
}

func GetPDF() echo.HandlerFunc {
	return func(c echo.Context) error {
		lang := c.Param("language")

		p := New(c)
		dat, err := p.GetPDF(lang)
		if err != nil {
			return err
		}
		return c.Blob(http.StatusOK, "application/pdf", dat)
	}
}

func GetFile() echo.HandlerFunc {
	return func(c echo.Context) error {
		p := New(c)
		fileLoc, err := p.GetFile(c.Param("file"))
		if err == ErrorFileNotFound {
			return echo.NewHTTPError(http.StatusNotFound, err)
		} else if err != nil {
			return err
		}

		return c.File(fileLoc)
	}
}

func GetAttachment() echo.HandlerFunc {
	return func(c echo.Context) error {
		p := New(c)
		attachment := c.Param("attachment")

		val, err := p.GetAttachment(attachment)
		if err == ErrorFileNotFound {
			return echo.NewHTTPError(http.StatusNotFound, ErrorFileNotFound)
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

func GetRanklist(DB *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		problemset, problem := c.Param("name"), c.Param("problem")
		prob := c.Get("problem").(problems.Problem)

		sbs := make([]*models.Submission, 0)

		//@TODO something better?
		if err := queries.Raw("SELECT DISTINCT ON (s1.user_id) s1.* FROM (SELECT s1.user_id, MAX(s1.score) as score FROM submissions s1 WHERE problemset=$1 AND problem=$2 GROUP BY s1.user_id) s2 INNER JOIN submissions s1 ON s1.user_id=s2.user_id AND s1.score=s2.score AND s1.problemset=$1 AND s1.problem=$2", problemset, problem).Bind(c.Request().Context(), DB, &sbs); err != nil {
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

func GetSubmit(DB *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		prob := c.Get("problem").(problems.Problem)

		c.Set("title", fmt.Sprintf("Beküldés - %s (%s)", i18n.TranslateContent("hungarian", prob.Titles()).String(), prob.Name()))
		return c.Render(http.StatusOK, "problemset/problem/submit", Problem{
			Problem:      prob,
			LastLanguage: helpers.GetUserLastLanguage(c, DB),
		})
	}
}

func GetStatus(DB *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		ac := c.QueryParam("ac")
		problemset, problem := c.Param("name"), c.Param("problem")

		prob := c.Get("problem").(problems.Problem)

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

func PostTag(DB *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.FormValue("tagID"))
		if err != nil {
			return err
		}

		u := c.Get("user").(*models.User)
		if u == nil {
			return helpers.UnauthorizedError(c)
		}

		p := New(c)
		tg := TagManager{p}
		if err = tg.CreateTag(DB, id, u); err != nil {
			return err
		}

		return c.Redirect(http.StatusFound, filepath.Dir(c.Request().URL.Path)+"/")
	}
}

func DeleteTag(DB *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return err
		}

		u := c.Get("user").(*models.User)
		if u == nil {
			return helpers.UnauthorizedError(c)
		}

		p := New(c)
		tg := TagManager{p}
		if err = tg.DeleteTag(DB, id, u); err != nil {
			return err
		}

		return c.Redirect(http.StatusFound, filepath.Dir(filepath.Dir(c.Request().URL.Path))+"/")
	}
}
