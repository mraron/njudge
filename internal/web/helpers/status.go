package helpers

import (
	"github.com/mraron/njudge/internal/web/helpers/pagination"
	"github.com/mraron/njudge/internal/web/models"
	"net/url"

	"github.com/jmoiron/sqlx"
	. "github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type StatusPage struct {
	CurrentPage int
	Pages       []pagination.Link
	Submissions []*models.Submission
}

func GetStatusPage(DB *sqlx.DB, page, perPage int, order QueryMod, query []QueryMod, qu url.Values) (*StatusPage, error) {
	sbs, err := models.Submissions(append(append([]QueryMod{Limit(perPage), Offset((page - 1) * perPage)}, query...), order)...).All(DB)
	if err != nil {
		return nil, err
	}

	cnt, err := models.Submissions(query...).Count(DB)
	if err != nil {
		return nil, err
	}

	pages, err := pagination.LinksWithCountLimit(page, perPage, cnt, qu, 6)
	if err != nil {
		return nil, err
	}

	return &StatusPage{page, pages, sbs}, nil
}
