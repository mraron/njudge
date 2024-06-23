package njudge

import (
	"context"
)

type SortDirection string

const (
	SortASC  SortDirection = "ASC"
	SortDESC SortDirection = "DESC"
)

type ProblemSortField string

var (
	ProblemSortFieldID          ProblemSortField = "id"
	ProblemSortFieldSolverCount ProblemSortField = "solver_count"
)

type ProblemListRequest struct {
	Problemset     string
	SortDir        SortDirection
	SortField      ProblemSortField
	Page           int
	PerPage        int
	TitleFilter    string
	TagFilter      []string
	AuthorFilter   *string
	CategoryFilter CategoryFilter
	User           *User
}

func (r ProblemListRequest) IsFiltered() bool {
	return r.TitleFilter != "" || len(r.TagFilter) > 0 || r.CategoryFilter.Type != CategoryFilterNone || r.AuthorFilter != nil
}

type ProblemList struct {
	PaginationData PaginationData
	Problems       []Problem
}

type ProblemListQuery interface {
	GetProblemList(ctx context.Context, req ProblemListRequest) (*ProblemList, error)
}
