package web

import (
	"github.com/labstack/echo"
	"strconv"
)

type paginationData struct {
	_page      int
	_perPage   int
	_sortDir   string
	_sortField string
}

func parsePaginationData(c echo.Context) (*paginationData, error) {
	res := &paginationData{}
	var err error

	_page := c.QueryParam("_page")
	_perPage := c.QueryParam("_perPage")

	res._sortDir = c.QueryParam("_sortDir")
	res._sortField = c.QueryParam("_sortField")

	res._page, err = strconv.Atoi(_page)
	if err != nil {
		return nil, err
	}

	res._perPage, err = strconv.Atoi(_perPage)
	if err != nil {
		return nil, err
	}

	return res, nil
}
