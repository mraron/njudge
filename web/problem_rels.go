package web

import (
	"github.com/labstack/echo/v4"
	"github.com/mraron/njudge/web/helpers/roles"
	"github.com/mraron/njudge/web/models"
	"github.com/volatiletech/sqlboiler/v4/boil"
	. "github.com/volatiletech/sqlboiler/v4/queries/qm"
	"net/http"
	"strconv"
)

func (s *Server) getAPIProblemRels(c echo.Context) error {
	u := c.Get("user").(*models.User)

	if !roles.Can(roles.Role(u.Role), roles.ActionView, "api/v1/problem_rels") {
		return s.unauthorizedError(c)
	}

	data, err := parsePaginationData(c)
	if err != nil {
		return s.internalError(c, err, "error")
	}

	lst, err := models.ProblemRels(OrderBy(data._sortField+" "+data._sortDir), Limit(data._perPage), Offset(data._perPage*(data._page-1))).All(s.db)
	if err != nil {
		return s.internalError(c, err, "error")
	}

	return c.JSON(http.StatusOK, lst)
}

func (s *Server) postAPIProblemRel(c echo.Context) error {
	u := c.Get("user").(*models.User)
	if !roles.Can(roles.Role(u.Role), roles.ActionCreate, "api/v1/problem_rels") {
		return s.unauthorizedError(c)
	}

	pr := new(models.ProblemRel)
	if err := c.Bind(pr); err != nil {
		return s.internalError(c, err, "error")
	}

	return pr.Insert(s.db, boil.Infer())
}

func (s *Server) getAPIProblemRel(c echo.Context) error {
	u := c.Get("user").(*models.User)
	if !roles.Can(roles.Role(u.Role), roles.ActionView, "api/v1/problem_rels") {
		return s.unauthorizedError(c)
	}

	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return s.internalError(c, err, "error")
	}

	pr, err := models.ProblemRels(Where("id=?", id)).One(s.db)
	if err != nil {
		return s.internalError(c, err, "error")
	}

	return c.JSON(http.StatusOK, pr)
}

func (s *Server) deleteAPIProblemRel(c echo.Context) error {
	u := c.Get("user").(*models.User)
	if !roles.Can(roles.Role(u.Role), roles.ActionDelete, "api/v1/problem_rels") {
		return s.unauthorizedError(c)
	}

	id_ := c.Param("id")

	id, err := strconv.Atoi(id_)
	if err != nil {
		return s.internalError(c, err, "error")
	}

	pr, err := models.ProblemRels(Where("id=?", id)).One(s.db)
	if err != nil {
		return s.internalError(c, err, "error")
	}

	_, err = pr.Delete(s.db)
	if err != nil {
		return s.internalError(c, err, "error")
	}

	return c.String(http.StatusOK, "ok")
}

func (s *Server) putAPIProblemRel(c echo.Context) error {
	u := c.Get("user").(*models.User)
	if !roles.Can(roles.Role(u.Role), roles.ActionEdit, "api/v1/problem_rels") {
		return s.unauthorizedError(c)
	}

	id_ := c.Param("id")

	id, err := strconv.Atoi(id_)
	if err != nil {
		return s.internalError(c, err, "error")
	}

	pr := new(models.ProblemRel)
	if err = c.Bind(pr); err != nil {
		return s.internalError(c, err, "error")
	}

	pr.ID = id
	_, err = pr.Update(s.db, boil.Infer())

	if err != nil {
		return s.internalError(c, err, "error")
	}

	return c.JSON(http.StatusOK, struct {
		Message string `json:"message"`
	}{"updated"})
}
