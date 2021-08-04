package pagination

import "github.com/labstack/echo/v4"

type Data struct {
	Page      int `query:"_page"`
	PerPage   int `query:"_perPage"`
	SortDir   string `query:"_sortDir"`
	SortField string `query:"_sortField"`
}

func Parse(c echo.Context) (*Data, error) {
	data := &Data{}
	return data, c.Bind(data)
}

type Link struct {
	Name string
	Active bool
	Disabled bool
	Url string
}
