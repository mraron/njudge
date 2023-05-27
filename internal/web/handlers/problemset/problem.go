package problemset

import (
	"bytes"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/mraron/njudge/internal/web/domain/problem"
	"github.com/mraron/njudge/internal/web/helpers"
	"github.com/mraron/njudge/internal/web/helpers/i18n"
	"github.com/mraron/njudge/internal/web/models"
	"github.com/mraron/njudge/internal/web/services"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"io"
	"mime"
	"net/http"
	"path/filepath"
	"sort"
	"strconv"
)

func GetProblem() echo.HandlerFunc {
	return func(c echo.Context) error {
		prob := c.Get("problem").(problem.Problem)
		stats := c.Get("problemStats").(problem.StatsData)

		c.Set("title", fmt.Sprintf("Leírás - %s (%s)", i18n.TranslateContent("hungarian", prob.Titles()).String(), prob.Name()))

		return c.Render(http.StatusOK, "problemset/problem/problem", struct {
			problem.Problem
			problem.StatsData
		}{Problem: prob, StatsData: stats})
	}
}

func GetProblemPDF() echo.HandlerFunc {
	return func(c echo.Context) error {
		p := c.Get("problem").(problem.Problem)

		lang := c.Param("language")

		dat, err := p.GetPDF(lang)
		if err != nil {
			return err
		}

		return c.Blob(http.StatusOK, "application/pdf", dat)
	}
}

func GetProblemFile() echo.HandlerFunc {
	return func(c echo.Context) error {
		p := c.Get("problem").(problem.Problem)

		fileLoc, err := p.GetFile(c.Param("file"))
		if err == problem.ErrorFileNotFound {
			return echo.NewHTTPError(http.StatusNotFound, err)
		} else if err != nil {
			return err
		}

		return c.File(fileLoc)
	}
}

func GetProblemAttachment() echo.HandlerFunc {
	return func(c echo.Context) error {
		p := c.Get("problem").(problem.Problem)
		attachment := c.Param("attachment")

		val, err := p.GetAttachment(attachment)
		if err == problem.ErrorFileNotFound {
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

func GetProblemRanklist(DB *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		problemset, problemName := c.Param("name"), c.Param("problem")
		prob := c.Get("problem").(problem.Problem)

		sbs := make([]*models.Submission, 0)

		//@TODO something better?
		if err := queries.Raw("SELECT DISTINCT ON (s1.user_id) s1.* FROM (SELECT s1.user_id, MAX(s1.score) as score FROM submissions s1 WHERE problemset=$1 AND problem=$2 GROUP BY s1.user_id) s2 INNER JOIN submissions s1 ON s1.user_id=s2.user_id AND s1.score=s2.score AND s1.problemset=$1 AND s1.problem=$2", problemset, problemName).Bind(c.Request().Context(), DB, &sbs); err != nil {
			return err
		}

		sort.Slice(sbs, func(i, j int) bool {
			return sbs[i].Score.Float32 > sbs[j].Score.Float32
		})

		c.Set("title", fmt.Sprintf("Eredmények - %s (%s)", i18n.TranslateContent("hungarian", prob.Titles()).String(), prob.Name()))
		return c.Render(http.StatusOK, "problemset/problem/ranklist", struct {
			Problem     problem.Problem
			Submissions []*models.Submission
		}{prob, sbs})
	}
}

func GetProblemSubmit() echo.HandlerFunc {
	return func(c echo.Context) error {
		prob := c.Get("problem").(problem.Problem)
		stats := c.Get("problemStats").(problem.StatsData)

		c.Set("title", fmt.Sprintf("Beküldés - %s (%s)", i18n.TranslateContent("hungarian", prob.Titles()).String(), prob.Name()))
		return c.Render(http.StatusOK, "problemset/problem/submit", struct {
			problem.Problem
			problem.StatsData
		}{Problem: prob, StatsData: stats})
	}
}

func GetProblemStatus(DB *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		ac := c.QueryParam("ac")
		problemset, problemName := c.Param("name"), c.Param("problem")

		prob := c.Get("problem").(problem.Problem)

		page, err := strconv.Atoi(c.QueryParam("page"))
		if err != nil || page <= 0 {
			page = 1
		}

		var query []qm.QueryMod
		if ac == "1" {
			query = []qm.QueryMod{qm.Where("verdict = 0"), qm.Where("problem = ?", problemName), qm.Where("problemset = ?", problemset)}
		} else {
			query = []qm.QueryMod{qm.Where("problem = ?", problemName), qm.Where("problemset = ?", problemset)}
		}

		statusPage, err := helpers.GetStatusPage(DB.DB, page, 20, qm.OrderBy("id DESC"), query, c.Request().URL.Query())
		if err != nil {
			return err
		}

		c.Set("title", fmt.Sprintf("Beküldések - %s (%s)", i18n.TranslateContent("hungarian", prob.Titles()).String(), prob.Name()))
		return c.Render(http.StatusOK, "problemset/problem/status", statusPage)
	}
}

func PostProblemTag(tgs services.TagsService) echo.HandlerFunc {
	type request struct {
		TagID int `form:"tagID"`
	}
	return func(c echo.Context) error {
		data := request{}
		if err := c.Bind(&data); err != nil {
			return err
		}

		u := c.Get("user").(*models.User)
		if u == nil {
			return helpers.UnauthorizedError(c)
		}

		pr := c.Get("problem").(problem.Problem)
		if err := tgs.Add(c.Request().Context(), data.TagID, pr.ID, u.ID); err != nil {
			return err
		}

		return c.Redirect(http.StatusFound, c.Echo().Reverse("getProblemMain", pr.Problemset, pr.Problem))
	}
}

func DeleteProblemTag(tgs services.TagsService) echo.HandlerFunc {
	type request struct {
		TagID int `form:"tagID"`
	}
	return func(c echo.Context) error {
		data := request{}
		if err := c.Bind(&data); err != nil {
			return err
		}

		u := c.Get("user").(*models.User)
		if u == nil {
			return helpers.UnauthorizedError(c)
		}

		pr := c.Get("problemRel").(problem.Problem)
		if err := tgs.Delete(c.Request().Context(), data.TagID, u.ID); err != nil {
			return err
		}

		return c.Redirect(http.StatusFound, c.Echo().Reverse("getProblemMain", pr.Problemset, pr.Problem))
	}
}
