package pagination

import (
	"errors"
	"github.com/labstack/echo/v4"
	"net/url"
	"strconv"
)

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

func Links(page, perPage int, cnt int64, qu url.Values) ([]Link, error){
	pageCnt := (int(cnt)+perPage-1)/perPage
	pages := make([]Link, pageCnt+2)
	pages[0] = Link{"&laquo;", false, true, "#"}
	if page>1 {
		qu.Set("page", strconv.Itoa(page-1))

		pages[0].Disabled = false
		pages[0].Url = "?"+qu.Encode()
	}
	for i := 1; i < len(pages)-1; i++ {
		qu.Set("page", strconv.Itoa(i))
		pages[i] = Link{strconv.Itoa(i), false, false, "?"+qu.Encode()}
		if i==page {
			pages[i].Active = true
			pages[i].Disabled = true
		}
	}
	pages[len(pages)-1] = Link{"&raquo;", false, true, "#"}
	if page<pageCnt {
		qu.Set("page", strconv.Itoa(page+1))

		pages[len(pages)-1].Disabled = false
		pages[len(pages)-1].Url = "?"+qu.Encode()
	}

	if page>len(pages) {
		return nil, errors.New("no such page")
	}

	return pages, nil
}

func abs(x int) int {
	if x<0 {
		return -x
	}

	return x
}

func LinksWithCountLimit(page, perPage int, cnt int64, qu url.Values, pageLimit int) ([]Link, error) {
	links, err := Links(page, perPage, cnt, qu)
	if err != nil {
		return links, err
	}

	empty := Link{"...", false, true, "#"}

	ans := make([]Link, 0)
	ans = append(ans, links[0])

	for i:=1;i<len(links)-1;i++ {
		if i==1 {
			ans = append(ans, links[i])
			if abs(i-page)>pageLimit+1 {
				ans = append(ans, empty)
			}
			continue
		}

		if i==len(links)-2 {
			if abs(i-page)>pageLimit+1 {
				ans = append(ans, empty)

			}
			ans = append(ans, links[i])
			continue
		}

		if abs(i-page)>pageLimit {
			continue
		}

		ans = append(ans, links[i])
	}

	ans = append(ans, links[len(links)-1])
	return ans, nil
}