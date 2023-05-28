package problemset

import (
	"bytes"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/mraron/njudge/internal/web/domain/problem"
	"github.com/mraron/njudge/internal/web/helpers"
	"github.com/mraron/njudge/internal/web/helpers/i18n"
	"github.com/mraron/njudge/internal/web/helpers/pagination"
	"github.com/mraron/njudge/internal/web/models"
	"github.com/mraron/njudge/internal/web/services"
	"github.com/mraron/njudge/pkg/problems"
	"github.com/volatiletech/sqlboiler/v4/queries"
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

func GetProblemStatus(statusPageService services.StatusPageService) echo.HandlerFunc {
	type request struct {
		AC     string `query:"ac"`
		UserID int    `query:"user_id"`
		Page   int    `query:"page"`

		Problemset string `param:"name"`
		Problem    string `param:"problem"`
	}
	return func(c echo.Context) error {
		prob := c.Get("problem").(problem.Problem)

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

			UserID:    0,
			GETValues: c.Request().URL.Query(),
		}

		if data.AC == "1" {
			ac := problems.VerdictAC
			statusReq.Verdict = &ac
		}

		statusPage, err := statusPageService.GetStatusPage(c.Request().Context(), statusReq)
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

		return c.Redirect(http.StatusFound, c.Echo().Reverse("getProblemMain", pr.Problemset, pr.ProblemRel.Problem))
	}
}

func DeleteProblemTag(tgs services.TagsService) echo.HandlerFunc {
	type request struct {
		TagID int `param:"id"`
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
		if err := tgs.Delete(c.Request().Context(), data.TagID, pr.ID, u.ID); err != nil {
			return err
		}

		return c.Redirect(http.StatusFound, c.Echo().Reverse("getProblemMain", pr.Problemset, pr.ProblemRel.Problem))
	}
}
