package db

import (
	"context"
	"database/sql"

	"github.com/mraron/njudge/internal/njudge"
	"github.com/mraron/njudge/internal/njudge/db/models"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type Tags struct {
	db *sql.DB
}

func NewTags(db *sql.DB) *Tags {
	return &Tags{
		db: db,
	}
}

func (ts *Tags) toNjudge(t *models.Tag) njudge.Tag {
	return njudge.Tag{
		ID:   t.ID,
		Name: t.Name,
	}
}

func (ts *Tags) toModel(t njudge.Tag) *models.Tag {
	return &models.Tag{
		ID:   t.ID,
		Name: t.Name,
	}
}

func (ts *Tags) Get(ctx context.Context, ID int) (*njudge.Tag, error) {
	dbobj, err := models.Tags(models.TagWhere.ID.EQ(ID)).One(ctx, ts.db)
	if err != nil {
		return nil, MaskNotFoundError(err, njudge.ErrorTagNotFound)
	}

	res := ts.toNjudge(dbobj)
	return &res, nil
}

func (ts *Tags) GetByName(ctx context.Context, name string) (*njudge.Tag, error) {
	dbobj, err := models.Tags(models.TagWhere.Name.EQ(name)).One(ctx, ts.db)
	if err != nil {
		return nil, err
	}

	res := ts.toNjudge(dbobj)
	return &res, nil
}

func (ts *Tags) GetAll(ctx context.Context) ([]njudge.Tag, error) {
	dbobjs, err := models.Tags().All(ctx, ts.db)
	if err != nil {
		return nil, err
	}

	res := make([]njudge.Tag, len(dbobjs))
	for ind := range dbobjs {
		res[ind] = ts.toNjudge(dbobjs[ind])
	}

	return res, nil
}

func (ts *Tags) Insert(ctx context.Context, t njudge.Tag) (*njudge.Tag, error) {
	dbobj := ts.toModel(t)
	err := dbobj.Insert(ctx, ts.db, boil.Infer())
	if err != nil {
		return nil, err
	}

	res := ts.toNjudge(dbobj)
	return &res, nil
}

func (ts *Tags) Delete(ctx context.Context, ID int) error {
	_, err := models.Tags(models.TagWhere.ID.EQ(ID)).DeleteAll(ctx, ts.db)
	return err
}

func (ts *Tags) Update(ctx context.Context, t njudge.Tag) error {
	dbobj := ts.toModel(t)
	_, err := dbobj.Update(ctx, ts.db, boil.Infer())
	return err
}
