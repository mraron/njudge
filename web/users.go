package web

import (
	"github.com/labstack/echo/v4"
	"github.com/mraron/njudge/web/helpers"
	"github.com/mraron/njudge/web/helpers/pagination"
	"github.com/mraron/njudge/web/helpers/roles"
	"github.com/mraron/njudge/web/models"
	"github.com/volatiletech/sqlboiler/v4/boil"
	. "github.com/volatiletech/sqlboiler/v4/queries/qm"
	"net/http"
	"strconv"
)

func (s *Server) getAPIUsers(c echo.Context) error {
	u := c.Get("user").(*models.User)

	if !roles.Can(roles.Role(u.Role), roles.ActionView, "api/v1/users") {
		return helpers.UnauthorizedError(c)
	}

	data, err := pagination.Parse(c)
	if err != nil {
		return helpers.InternalError(c, err, "error")
	}

	lst, err := models.Users(OrderBy(data.SortField+" "+data.SortDir), Limit(data.PerPage), Offset(data.PerPage*(data.Page-1))).All(s.DB)
	if err != nil {
		return helpers.InternalError(c, err, "error")
	}

	for i := 0; i < len(lst); i++ {
		helpers.CensorUserPassword(lst[i])
	}

	return c.JSON(http.StatusOK, lst)
}

func (s *Server) postAPIUser(c echo.Context) error {
	u := c.Get("user").(*models.User)
	if !roles.Can(roles.Role(u.Role), roles.ActionCreate, "api/v1/users") {
		return helpers.UnauthorizedError(c)
	}

	pr := new(models.User)
	if err := c.Bind(pr); err != nil {
		return helpers.InternalError(c, err, "error")
	}

	return pr.Insert(s.DB, boil.Infer())
}

func (s *Server) getAPIUser(c echo.Context) error {
	u := c.Get("user").(*models.User)
	if !roles.Can(roles.Role(u.Role), roles.ActionView, "api/v1/users") {
		return helpers.UnauthorizedError(c)
	}

	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return helpers.InternalError(c, err, "error")
	}

	pr, err := models.Users(Where("id=?", id)).One(s.DB)
	if err != nil {
		return helpers.InternalError(c, err, "error")
	}

	helpers.CensorUserPassword(pr)

	return c.JSON(http.StatusOK, pr)
}

func (s *Server) deleteAPIUser(c echo.Context) error {
	u := c.Get("user").(*models.User)
	if !roles.Can(roles.Role(u.Role), roles.ActionDelete, "api/v1/users") {
		return helpers.UnauthorizedError(c)
	}

	id_ := c.Param("id")

	id, err := strconv.Atoi(id_)
	if err != nil {
		return helpers.InternalError(c, err, "error")
	}

	pr, err := models.Users(Where("id=?", id)).One(s.DB)
	if err != nil {
		return helpers.InternalError(c, err, "error")
	}

	_, err = pr.Delete(s.DB)
	if err != nil {
		return helpers.InternalError(c, err, "error")
	}

	return c.String(http.StatusOK, "ok")
}

func (s *Server) putAPIUser(c echo.Context) error {
	u := c.Get("user").(*models.User)
	if !roles.Can(roles.Role(u.Role), roles.ActionEdit, "api/v1/users") {
		return helpers.UnauthorizedError(c)
	}

	id_ := c.Param("id")

	id, err := strconv.Atoi(id_)
	if err != nil {
		return helpers.InternalError(c, err, "error")
	}

	pr := new(models.User)
	if err = c.Bind(pr); err != nil {
		return helpers.InternalError(c, err, "error")
	}

	pr.ID = id
	_, err = pr.Update(s.DB, boil.Infer())

	if err != nil {
		return helpers.InternalError(c, err, "error")
	}

	return c.JSON(http.StatusOK, struct {
		Message string `json:"message"`
	}{"updated"})
}

