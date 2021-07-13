package web

import (
	"github.com/labstack/echo/v4"
	"github.com/mraron/njudge/web/helpers"
	. "github.com/volatiletech/sqlboiler/v4/queries/qm"
	"net/http"
	"strconv"
)

func (s *Server) getProblemsetProblemStatus(c echo.Context) error {
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

	statusPage, err := helpers.GetStatusPage(s.DB, page, 20, OrderBy("id DESC"), query, c.Request().URL.Query())
	if err != nil {
		return helpers.InternalError(c, err, "Belső hiba")
	}

	return c.Render(http.StatusOK, "problemset_problem_status.gohtml", statusPage)
}

func (s *Server) getProblemsetStatus(c echo.Context) error {
	ac := c.QueryParam("ac")
	userID := c.QueryParam("user_id")
	problemset := c.QueryParam("problem_set")
	problem := c.QueryParam("problem")
	page, err := strconv.Atoi(c.QueryParam("page"))
	if  err != nil || page<=0 {
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

	statusPage, err := helpers.GetStatusPage(s.DB, page, 20, OrderBy("id DESC"), query, c.Request().URL.Query())
	if err != nil {
		return helpers.InternalError(c, err, "Belső hiba")
	}

	return c.Render(http.StatusOK, "status.gohtml", statusPage)
}
