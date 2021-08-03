package web

import (
	"github.com/labstack/echo/v4"
	"github.com/mraron/njudge/web/extmodels"
	"github.com/mraron/njudge/web/helpers"
	"github.com/mraron/njudge/web/helpers/pagination"
	"github.com/mraron/njudge/web/helpers/roles"
	"github.com/mraron/njudge/web/models"
	"github.com/volatiletech/sqlboiler/v4/boil"
	. "github.com/volatiletech/sqlboiler/v4/queries/qm"
	"net/http"
	"strconv"
)

func (s *Server) getAPIJudges(c echo.Context) error {
	u := c.Get("user").(*models.User)
	if !roles.Can(roles.Role(u.Role), roles.ActionView, "api/v1/judges") {
		return helpers.UnauthorizedError(c)
	}

	data, err := pagination.Parse(c)
	if err != nil {
		return err
	}

	lst, err := models.Judges(OrderBy(data.SortField+" "+data.SortDir), Limit(data.PerPage), Offset(data.PerPage*(data.Page-1))).All(s.DB)
	if err != nil {
		return err
	}

	local := make([]extmodels.Judge, len(lst))
	for ind, _ := range lst {
		local[ind] = extmodels.NewJudgeFromModelsJudge(lst[ind])
	}

	return c.JSON(http.StatusOK, local)
}

func (s *Server) postAPIJudge(c echo.Context) error {
	u := c.Get("user").(*models.User)
	if !roles.Can(roles.Role(u.Role), roles.ActionCreate, "api/v1/judges") {
		return helpers.UnauthorizedError(c)
	}

	j := new(models.Judge)
	if err := c.Bind(j); err != nil {
		return err
	}

	err := j.Insert(s.DB, boil.Infer())
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, extmodels.NewJudgeFromModelsJudge(j))
}

func (s *Server) getAPIJudge(c echo.Context) error {
	u := c.Get("user").(*models.User)
	if !roles.Can(roles.Role(u.Role), roles.ActionView, "api/v1/judges") {
		return helpers.UnauthorizedError(c)
	}

	id_ := c.Param("id")

	id, err := strconv.Atoi(id_)
	if err != nil {
		return err
	}

	j, err := models.Judges(Where("id=?", id)).One(s.DB)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, extmodels.NewJudgeFromModelsJudge(j))
}

func (s *Server) deleteAPIJudge(c echo.Context) error {
	u := c.Get("user").(*models.User)
	if !roles.Can(roles.Role(u.Role), roles.ActionDelete, "api/v1/judges") {
		return helpers.UnauthorizedError(c)
	}

	id_ := c.Param("id")

	id, err := strconv.Atoi(id_)
	if err != nil {
		return err
	}

	j, err := models.Judges(Where("id=?", id)).One(s.DB)
	if err != nil {
		return err
	}

	_, err = j.Delete(s.DB)
	if err != nil {
		return err
	}

	return c.String(http.StatusOK, "deleted")
}

func (s *Server) putAPIJudge(c echo.Context) error {
	u := c.Get("user").(*models.User)
	if !roles.Can(roles.Role(u.Role), roles.ActionEdit, "api/v1/judges") {
		return helpers.UnauthorizedError(c)
	}

	id_ := c.Param("id")

	id, err := strconv.Atoi(id_)
	if err != nil {
		return err
	}

	j := new(extmodels.Judge)
	if err = c.Bind(j); err != nil {
		return err
	}

	model, err := models.Judges(Where("id=?", id)).One(s.DB)
	if err != nil {
		return err
	}

	model.Host = j.Host
	model.Port = j.Port

	_, err = model.Update(s.DB, boil.Infer())
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, struct {
		Message string `json:"message"`
	}{"updated"})
}
