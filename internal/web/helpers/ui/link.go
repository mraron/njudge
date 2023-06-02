package ui

import (
	"context"
	"database/sql"
	"github.com/mraron/njudge/internal/web/models"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"strconv"
)

type Link struct {
	Text string
	Href string
}

func TopCategoryLink(ctx context.Context, db *sql.DB, categoryID int) (Link, error) {
	var (
		category *models.ProblemCategory
		err      error
	)

	orig := categoryID

	for {
		category, err = models.ProblemCategories(qm.Where("id = ?", categoryID)).One(ctx, db)
		if err != nil {
			return Link{}, err
		}

		if !category.ParentID.Valid {
			break
		}
		categoryID = category.ParentID.Int
	}

	return Link{
		Text: category.Name,
		Href: "/task_archive#category" + strconv.Itoa(orig),
	}, nil
}
