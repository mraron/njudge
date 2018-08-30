package web

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/mraron/njudge/web/models"
	"github.com/mraron/njudge/web/roles"
	"net/http"
	"strconv"
)

func (s *Server) getAPIProblemRels(c echo.Context) error {
	u := c.Get("user").(*models.User)
	if !roles.Can(u.Role, roles.ActionView, "api/v1/problem_rels") {
		return s.unauthorizedError(c)
	}

	data, err := parsePaginationData(c)
	if err != nil {
		fmt.Println(err)
		return err
	}

	lst, err := models.ProblemRelAPIGet(s.db, data._page, data._perPage, data._sortDir, data._sortField)
	if err != nil {
		fmt.Println(err, "models")
		return err
	}

	return c.JSON(http.StatusOK, lst)
}

func (s *Server) postAPIProblemRel(c echo.Context) error {
	u := c.Get("user").(*models.User)
	if !roles.Can(u.Role, roles.ActionCreate, "api/v1/problem_rels") {
		return s.unauthorizedError(c)
	}

	pr := new(models.ProblemRel)
	if err := c.Bind(pr); err != nil {
		return s.internalError(c, err, err.Error())
	}

	err := s.db.Create(pr).Error
	if err != nil {
		return s.internalError(c, err, err.Error())
	}

	return c.String(http.StatusOK, "ok")
}

func (s *Server) getAPIProblemRel(c echo.Context) error {
	u := c.Get("user").(*models.User)
	if !roles.Can(u.Role, roles.ActionView, "api/v1/problem_rels") {
		return s.unauthorizedError(c)
	}

	id_ := c.Param("id")

	id, err := strconv.Atoi(id_)
	if err != nil {
		fmt.Println(err)
		return err
	}

	pr, err := models.ProblemRelFromId(s.db, id)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return c.JSON(http.StatusOK, pr)
}

func (s *Server) deleteAPIProblemRel(c echo.Context) error {
	u := c.Get("user").(*models.User)
	if !roles.Can(u.Role, roles.ActionDelete, "api/v1/problem_rels") {
		return s.unauthorizedError(c)
	}

	id_ := c.Param("id")

	id, err := strconv.Atoi(id_)
	if err != nil {
		return s.internalError(c, err, "error")
	}

	pr, err := models.ProblemRelFromId(s.db, id)
	if err != nil {
		return s.internalError(c, err, "error")
	}

	err = s.db.Delete(pr).Error
	if err != nil {
		return s.internalError(c, err, "error")
	}

	return c.String(http.StatusOK, "ok")
}

func (s *Server) putAPIProblemRel(c echo.Context) error {
	u := c.Get("user").(*models.User)
	if !roles.Can(u.Role, roles.ActionEdit, "api/v1/problem_rels") {
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

	pr.ID = uint(id)
	err = s.db.Save(pr).Error
	if err != nil {
		return s.internalError(c, err, "error")
	}

	return c.String(http.StatusOK, "ok")
}
