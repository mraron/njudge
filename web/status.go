package web

import (
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/mraron/njudge/web/helpers"
	"github.com/mraron/njudge/web/models"
	. "github.com/volatiletech/sqlboiler/v4/queries/qm"
	"net/http"
	"net/url"
	"strconv"
)

type StatusPageLink struct {
	Name string
	Active bool
	Disabled bool
	Url string
}

type StatusPage struct {
	CurrentPage int
	Pages []StatusPageLink
	Submissions []*models.Submission
}

func (s *Server) GetStatusPage(page, perPage int, order QueryMod, query []QueryMod, qu url.Values) (*StatusPage, error) {
	pagination := []QueryMod{Limit(perPage), Offset((page-1)*perPage)}
	sbs, err := models.Submissions(append(append(pagination, query...), order)...).All(s.db)
	if err != nil {
		return nil, err
	}

	cnt, err := models.Submissions(query...).Count(s.db)
	if err != nil {
		return nil, err
	}

	pageCnt := (int(cnt)+perPage-1)/perPage
	pages := make([]StatusPageLink, pageCnt+2)
	pages[0] = StatusPageLink{"Előző", false, true, "#"}
	if page>1 {
		qu.Set("page", strconv.Itoa(page-1))

		pages[0].Disabled = false
		pages[0].Url = "?"+qu.Encode()
	}
	for i := 1; i < len(pages)-1; i++ {
		qu.Set("page", strconv.Itoa(i))
		pages[i] = StatusPageLink{strconv.Itoa(i), false, false, "?"+qu.Encode()}
		if i==page {
			pages[i].Active = true
		}
	}
	pages[len(pages)-1] = StatusPageLink{"Következő", false, true, "#"}
	if page<pageCnt {
		qu.Set("page", strconv.Itoa(page+1))

		pages[len(pages)-1].Disabled = false
		pages[len(pages)-1].Url = "?"+qu.Encode()
	}

	if page>len(pages) {
		return nil, errors.New("no such page")
	}

	return &StatusPage{page, pages, sbs}, nil
}

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

	statusPage, err := s.GetStatusPage(page, 20, OrderBy("id DESC"), query, c.Request().URL.Query())
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

	statusPage, err := s.GetStatusPage(page, 20, OrderBy("id DESC"), query, c.Request().URL.Query())
	if err != nil {
		return helpers.InternalError(c, err, "Belső hiba")
	}

	return c.Render(http.StatusOK, "status.gohtml", statusPage)
}
