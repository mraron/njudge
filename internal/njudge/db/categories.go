package db

import (
	"context"
	"database/sql"
	"errors"

	"github.com/mraron/njudge/internal/njudge"
	"github.com/mraron/njudge/internal/njudge/db/models"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func MaskNotFoundError(err, mask error) error {
	if errors.Is(err, sql.ErrNoRows) {
		return mask
	}

	return err
}

type Categories struct {
	db *sql.DB
}

func NewCategories(db *sql.DB) *Categories {
	return &Categories{
		db: db,
	}
}

func (cs *Categories) toNjudge(c *models.ProblemCategory) njudge.Category {
	return njudge.Category{
		ID:       c.ID,
		Name:     c.Name,
		Visible:  c.Visible,
		ParentID: c.ParentID,
	}
}

func (cs *Categories) toModel(c njudge.Category) *models.ProblemCategory {
	return &models.ProblemCategory{
		Name:     c.Name,
		Visible:  c.Visible,
		ParentID: c.ParentID,
	}
}

func (cs *Categories) Get(ctx context.Context, id int) (*njudge.Category, error) {
	dbObj, err := models.ProblemCategories(qm.Where("id=?", id)).One(ctx, cs.db)
	if err != nil {
		return nil, MaskNotFoundError(err, njudge.ErrorCategoryNotFound)
	}

	res := cs.toNjudge(dbObj)
	return &res, nil
}

func (cs *Categories) getAll(ctx context.Context, mods ...qm.QueryMod) ([]njudge.Category, error) {
	dbobjs, err := models.ProblemCategories(mods...).All(ctx, cs.db)
	if err != nil {
		return nil, MaskNotFoundError(err, njudge.ErrorCategoryNotFound)
	}

	res := make([]njudge.Category, len(dbobjs))
	for ind := range dbobjs {
		res[ind] = cs.toNjudge(dbobjs[ind])
	}

	return res, nil
}

func (cs *Categories) GetAll(ctx context.Context) ([]njudge.Category, error) {
	return cs.getAll(ctx)
}

func (cs *Categories) GetAllWithParent(ctx context.Context, parentID int) ([]njudge.Category, error) {
	if parentID == 0 {
		return cs.getAll(ctx, models.ProblemCategoryWhere.ParentID.IsNull())
	}

	return cs.getAll(ctx, models.ProblemCategoryWhere.ParentID.EQ(null.Int{
		Valid: true,
		Int:   parentID,
	}))
}

func (cs *Categories) Insert(ctx context.Context, c njudge.Category) (*njudge.Category, error) {
	dbobj := cs.toModel(c)
	err := dbobj.Insert(ctx, cs.db, boil.Infer())
	if err != nil {
		return nil, err
	}

	res := cs.toNjudge(dbobj)
	return &res, nil
}
