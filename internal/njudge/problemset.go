package njudge

import (
	"context"
	"errors"
)

type CodeVisibility string

const (
	CodeVisibilityPrivate CodeVisibility = "private"
	CodeVisibilitySolved  CodeVisibility = "solved"
	CodeVisibilityPublic  CodeVisibility = "public"
)

type Problemset struct {
	Name           string
	CodeVisibility CodeVisibility
}

var ErrorProblemsetNotFound = errors.New("njudge: problemset not found")

type Problemsets interface {
	GetByName(ctx context.Context, problemsetName string) (*Problemset, error)
	GetAll(ctx context.Context) ([]Problemset, error)
	Insert(ctx context.Context, problemset Problemset) error
}
