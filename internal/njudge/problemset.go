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

type ProblemsetRanklistRequest struct {
	Name    string
	Page    int
	PerPage int

	FilterAdmin bool
}

type ProblemsetRanklistRow struct {
	Place int
	ID    int
	Name  string
	Score float64
}

type ProblemsetRanklist struct {
	Rows           []ProblemsetRanklistRow
	PaginationData PaginationData
}

type ProblemsetRanklistService interface {
	GetRanklist(ctx context.Context, req ProblemsetRanklistRequest) (*ProblemsetRanklist, error)
}

func GetUserPlaceInProblemsetRanklist(ctx context.Context, name string, userID int, service ProblemsetRanklistService) (int, error) {
	res, err := service.GetRanklist(ctx, ProblemsetRanklistRequest{
		Name:        name,
		FilterAdmin: true,
	})
	if err != nil {
		return 0, err
	}
	for _, row := range res.Rows {
		if row.ID == userID {
			return row.Place, nil
		}
	}
	return 0, nil
}
