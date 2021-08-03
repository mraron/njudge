package helpers

import (
	"errors"
	"github.com/jmoiron/sqlx"
	"github.com/mraron/njudge/web/models"
	. "github.com/volatiletech/sqlboiler/v4/queries/qm"
	"net/url"
	"strconv"
)

type StatusPageLink struct {
	Name string
	Active bool
	Disabled bool
	Url string
}

type StatusPage struct {
	CurrentPage int
	Pages []StatusPageLink
	Submissions []*models.Submission
}

func GetStatusPage(DB *sqlx.DB, page, perPage int, order QueryMod, query []QueryMod, qu url.Values) (*StatusPage, error) {
	pagination := []QueryMod{Limit(perPage), Offset((page-1)*perPage)}
	sbs, err := models.Submissions(append(append(pagination, query...), order)...).All(DB)
	if err != nil {
		return nil, err
	}

	cnt, err := models.Submissions(query...).Count(DB)
	if err != nil {
		return nil, err
	}

	pageCnt := (int(cnt)+perPage-1)/perPage
	pages := make([]StatusPageLink, pageCnt+2)
	pages[0] = StatusPageLink{"Előző", false, true, "#"}
	if page>1 {
		qu.Set("page", strconv.Itoa(page-1))

		pages[0].Disabled = false
		pages[0].Url = "?"+qu.Encode()
	}
	for i := 1; i < len(pages)-1; i++ {
		qu.Set("page", strconv.Itoa(i))
		pages[i] = StatusPageLink{strconv.Itoa(i), false, false, "?"+qu.Encode()}
		if i==page {
			pages[i].Active = true
		}
	}
	pages[len(pages)-1] = StatusPageLink{"Következő", false, true, "#"}
	if page<pageCnt {
		qu.Set("page", strconv.Itoa(page+1))

		pages[len(pages)-1].Disabled = false
		pages[len(pages)-1].Url = "?"+qu.Encode()
	}

	if page>len(pages) {
		return nil, errors.New("no such page")
	}

	return &StatusPage{page, pages, sbs}, nil
}