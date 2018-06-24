package web

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/mraron/njudge/web/models"
	"net/http"
	"strconv"
)

func (s *Server) getAPIProblemRels(c echo.Context) error {
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
	pr := new(models.ProblemRel)
	if err := c.Bind(pr); err != nil {
		return err
	}

	return pr.Insert(s.db)
}

func (s *Server) getAPIProblemRel(c echo.Context) error {
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

	return pr.Delete(s.db)
}

func (s *Server) putAPIProblemRel(c echo.Context) error {
	id_ := c.Param("id")

	id, err := strconv.Atoi(id_)
	if err != nil {
		return err
	}

	pr := new(models.ProblemRel)
	if err = c.Bind(pr); err != nil {
		return err
	}

	pr.Id = int64(id)
	return pr.Update(s.db)
}
