package problem

import "context"

type Repository interface {
	Get(ctx context.Context, ID int) (*Problem, error)
	GetAll(ctx context.Context) ([]Problem, error)
	GetByNames(ctx context.Context, problemset string, problemName string) (*Problem, error)
}

type ProblemTagRepository interface {
	Get(ctx context.Context, ID int) (*ProblemTag, error)
	Add(ctx context.Context, pt ProblemTag) error
	Delete(ctx context.Context, ID int) error
}
