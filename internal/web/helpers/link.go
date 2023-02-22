package helpers

import (
	"github.com/jmoiron/sqlx"
	"github.com/mraron/njudge/internal/web/models"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"strconv"
)

type Link struct {
	Text string
	Href string
}

func TopCategoryLink(cat int, DB *sqlx.DB) (Link, error) {
	var (
		category *models.ProblemCategory
		err      error
	)

	orig := cat

	for {
		category, err = models.ProblemCategories(qm.Where("id = ?", cat)).One(DB)
		if err != nil {
			return Link{}, err
		}

		if !category.ParentID.Valid {
			break
		}
		cat = category.ParentID.Int
	}

	return Link{
		Text: category.Name,
		Href: "/task_archive#category" + strconv.Itoa(orig),
	}, nil
}
