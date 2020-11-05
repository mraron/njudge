package web

import (
	"github.com/labstack/echo/v4"
	"github.com/mraron/njudge/judge"
	"github.com/mraron/njudge/web/models"
	"github.com/mraron/njudge/web/roles"
	"github.com/volatiletech/sqlboiler/v4/boil"
	. "github.com/volatiletech/sqlboiler/v4/queries/qm"
	"net/http"
	"strconv"
	"time"
)

type Judge struct {
	Id          int64         `json:"id"`
	Name        string        `json:"name"`
	Host        string        `json:"host"`
	Port        string        `json:"port"`
	Load        float64       `json:"load"`
	ProblemsDir string        `json:"problems_dir"`
	ProblemList []string      `json:"problems_list"`
	Uptime      time.Duration `json:"uptime"`
	Ping        int           `json:"ping"`
	Online      bool          `json:"online"`
}

func NewJudgeFromModelsJudge(j *models.Judge) (res Judge) {
	res.Id = int64(j.ID)
	res.Host = j.Host
	res.Port = j.Port
	res.Ping = j.Ping
	res.Online = j.Online

	server := &judge.Server{}
	err := server.FromString(j.State)

	if err == nil {
		res.Name = server.Id
		res.Load = server.Load
		res.ProblemsDir = server.ProblemsDir
		res.ProblemList = server.ProblemList
		res.Uptime = server.Uptime
	}

	return
}

func (s *Server) getAPIJudges(c echo.Context) error {
	u := c.Get("user").(*models.User)
	if !roles.Can(roles.Role(u.Role), roles.ActionView, "api/v1/judges") {
		return s.unauthorizedError(c)
	}

	data, err := parsePaginationData(c)
	if err != nil {
		return s.internalError(c, err, "error")
	}

	lst, err := models.Judges(OrderBy(data._sortField+" "+data._sortDir), Limit(data._perPage), Offset(data._perPage*(data._page-1))).All(s.db)
	if err != nil {
		return s.internalError(c, err, "error")
	}

	local := make([]Judge, len(lst))
	for ind, _ := range lst {
		local[ind] = NewJudgeFromModelsJudge(lst[ind])
	}

	return c.JSON(http.StatusOK, local)
}

func (s *Server) postAPIJudge(c echo.Context) error {
	u := c.Get("user").(*models.User)
	if !roles.Can(roles.Role(u.Role), roles.ActionCreate, "api/v1/judges") {
		return s.unauthorizedError(c)
	}

	j := new(models.Judge)
	if err := c.Bind(j); err != nil {
		return s.internalError(c, err, "error")
	}

	err := j.Insert(s.db, boil.Infer())
	if err != nil {
		return s.internalError(c, err, "error")
	}

	return c.JSON(http.StatusOK, NewJudgeFromModelsJudge(j))
}

func (s *Server) getAPIJudge(c echo.Context) error {
	u := c.Get("user").(*models.User)
	if !roles.Can(roles.Role(u.Role), roles.ActionView, "api/v1/judges") {
		return s.unauthorizedError(c)
	}

	id_ := c.Param("id")

	id, err := strconv.Atoi(id_)
	if err != nil {
		return s.internalError(c, err, "error")
	}

	j, err := models.Judges(Where("id=?", id)).One(s.db)
	if err != nil {
		return s.internalError(c, err, "error")
	}

	return c.JSON(http.StatusOK, NewJudgeFromModelsJudge(j))
}

func (s *Server) deleteAPIJudge(c echo.Context) error {
	u := c.Get("user").(*models.User)
	if !roles.Can(roles.Role(u.Role), roles.ActionDelete, "api/v1/judges") {
		return s.unauthorizedError(c)
	}

	id_ := c.Param("id")

	id, err := strconv.Atoi(id_)
	if err != nil {
		return s.internalError(c, err, "error")
	}

	j, err := models.Judges(Where("id=?", id)).One(s.db)
	if err != nil {
		return s.internalError(c, err, "error")
	}

	_, err = j.Delete(s.db)
	if err != nil {
		return s.internalError(c, err, "error")
	}

	return c.String(http.StatusOK, "deleted")
}

func (s *Server) putAPIJudge(c echo.Context) error {
	u := c.Get("user").(*models.User)
	if !roles.Can(roles.Role(u.Role), roles.ActionEdit, "api/v1/judges") {
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

	model, err := models.Judges(Where("id=?", id)).One(s.db)
	if err != nil {
		return s.internalError(c, err, "error")
	}

	model.Host = j.Host
	model.Port = j.Port

	_, err = model.Update(s.db, boil.Infer())
	if err != nil {
		return s.internalError(c, err, "error")
	}

	return c.JSON(http.StatusOK, struct {
		Message string `json:"message"`
	}{"updated"})
}
