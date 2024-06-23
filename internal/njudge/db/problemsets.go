package db

import (
	"database/sql"
	"github.com/mraron/njudge/internal/njudge"
	"github.com/mraron/njudge/internal/njudge/db/models"
	"golang.org/x/net/context"
)

type Problemsets struct {
	db *sql.DB
}

func NewProblemsets(db *sql.DB) *Problemsets {
	return &Problemsets{db: db}
}

func (p Problemsets) toNjudge(ps *models.Problemset) *njudge.Problemset {
	return &njudge.Problemset{
		Name:           ps.Name,
		CodeVisibility: njudge.CodeVisibility(ps.CodeVisibility),
	}
}

func (p Problemsets) GetByName(ctx context.Context, problemsetName string) (*njudge.Problemset, error) {
	res, err := models.Problemsets(models.ProblemsetWhere.Name.EQ(problemsetName)).One(ctx, p.db)
	if err != nil {
		return nil, err
	}
	return p.toNjudge(res), nil
}

func (p Problemsets) GetAll(ctx context.Context) ([]njudge.Problemset, error) {

	panic("implement me")
}

func (p Problemsets) Insert(ctx context.Context, problemset njudge.Problemset) error {
	//TODO implement me
	panic("implement me")
}
