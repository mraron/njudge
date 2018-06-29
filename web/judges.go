package web

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/mraron/njudge/web/models"
	"github.com/mraron/njudge/web/roles"
	"net/http"
	"strconv"
	"time"
)

type Judge struct {
	Id          int64
	Name        string
	Host        string
	Port        string
	Load        float64
	ProblemsDir string
	ProblemList []string
	Uptime      time.Duration
	Ping        int
	Online      bool
}

func NewJudgeFromModelsJudge(j *models.Judge) (res Judge) {
	res.Id = j.Id
	res.Host = j.Host
	res.Port = j.Port
	res.Ping = j.Ping
	res.Online = j.Online

	if j.State != nil {
		res.Name = j.State.Id
		res.Load = j.State.Load
		res.ProblemsDir = j.State.ProblemsDir
		res.ProblemList = j.State.ProblemList
		res.Uptime = j.State.Uptime
	}

	return
}

func (s *Server) getAPIJudges(c echo.Context) error {
	u := c.Get("user").(*models.User)
	if !roles.Can(u.Role, roles.ActionView, "api/v1/judges") {
		return s.unauthorizedError(c)
	}

	data, err := parsePaginationData(c)
	if err != nil {
		return s.internalError(c, err, "error")
	}

	lst, err := models.JudgesAPIGet(s.db, data._page, data._perPage, data._sortDir, data._sortField)
	if err != nil {
		return s.internalError(c, err, "error")
	}

	local := make([]Judge, len(lst))
	for ind, _ := range lst {
		local[ind] = NewJudgeFromModelsJudge(lst[ind])
	}

	return c.JSON(http.StatusOK, local)
}

func (s *Server) postAPIJudges(c echo.Context) error {
	u := c.Get("user").(*models.User)
	if !roles.Can(u.Role, roles.ActionCreate, "api/v1/judges") {
		return s.unauthorizedError(c)
	}
	fmt.Println("itt")
	j := new(models.Judge)
	if err := c.Bind(j); err != nil {
		fmt.Println(err)
		return s.internalError(c, err, "error")
	}

	err := j.Insert(s.db)
	if err != nil {
		fmt.Println("itt")
		return s.internalError(c, err, "error")
	}

	return c.String(http.StatusOK, "inserted")
}

func (s *Server) getAPIJudge(c echo.Context) error {
	u := c.Get("user").(*models.User)
	if !roles.Can(u.Role, roles.ActionView, "api/v1/judges") {
		return s.unauthorizedError(c)
	}

	id_ := c.Param("id")

	id, err := strconv.Atoi(id_)
	if err != nil {
		return s.internalError(c, err, "error")
	}

	j, err := models.JudgeFromId(s.db, id)
	if err != nil {
		return s.internalError(c, err, "error")
	}

	return c.JSON(http.StatusOK, NewJudgeFromModelsJudge(j))
}

func (s *Server) deleteAPIJudge(c echo.Context) error {
	u := c.Get("user").(*models.User)
	if !roles.Can(u.Role, roles.ActionDelete, "api/v1/judges") {
		return s.unauthorizedError(c)
	}

	id_ := c.Param("id")

	id, err := strconv.Atoi(id_)
	if err != nil {
		return s.internalError(c, err, "error")
	}
	fmt.Println("wut")
	j, err := models.JudgeFromId(s.db, id)
	if err != nil {
		return s.internalError(c, err, "error")
	}

	err = j.Delete(s.db)
	if err != nil {
		return s.internalError(c, err, "error")
	}

	return c.String(http.StatusOK, "deleted")
}

func (s *Server) putAPIJudge(c echo.Context) error {
	u := c.Get("user").(*models.User)
	if !roles.Can(u.Role, roles.ActionEdit, "api/v1/judges") {
		return s.unauthorizedError(c)
	}

	id_ := c.Param("id")

	id, err := strconv.Atoi(id_)
	if err != nil {
		return s.internalError(c, err, "error")
	}

	j := new(Judge)
	if err = c.Bind(j); err != nil {
		return s.internalError(c, err, "error")
	}

	model, err := models.JudgeFromId(s.db, id)
	if err != nil {
		return s.internalError(c, err, "error")
	}

	fmt.Println(j.Host, j.Port)
	model.Host = j.Host
	model.Port = j.Port

	err = model.Update(s.db)
	fmt.Println(err)
	if err != nil {
		return s.internalError(c, err, "error")
	}

	return c.String(http.StatusOK, "updated")
}
