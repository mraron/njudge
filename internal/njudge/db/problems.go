package db

import (
	"context"

	"github.com/mraron/njudge/internal/njudge"
)

type Problems struct{}

func NewProblems() *Problems {
	return &Problems{}
}

func (ps *Problems) Get(ctx context.Context, ID int) (*njudge.Problem, error) {
	panic("not implemented") // TODO: Implement
}

func (ps *Problems) GetAll(ctx context.Context) ([]njudge.Problem, error) {
	panic("not implemented") // TODO: Implement
}

func (ps *Problems) Insert(ctx context.Context, p njudge.Problem) (*njudge.Problem, error) {
	panic("not implemented") // TODO: Implement
}

func (ps *Problems) Delete(ctx context.Context, ID int) error {
	panic("not implemented") // TODO: Implement
}

func (ps *Problems) Update(ctx context.Context, p njudge.Problem, fields []string) error {
	panic("not implemented") // TODO: Implement
}
