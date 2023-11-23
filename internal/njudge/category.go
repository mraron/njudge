package njudge

import (
	"context"
	"errors"

	"github.com/volatiletech/null/v8"
)

type Category struct {
	ID       int
	Name     string
	ParentID null.Int
}

var (
	ErrorCategoryNotFound = errors.New("njudge: category not found")
)

type Categories interface {
	GetAll(ctx context.Context) ([]Category, error)
	Insert(ctx context.Context, c Category) (*Category, error)
}

type CategoryFilterType int

const (
	CategoryFilterNone CategoryFilterType = iota
	CategoryFilterEmpty
	CategoryFilterID
)

type CategoryFilter struct {
	Type  CategoryFilterType
	Value interface{}
}

func NewCategoryIDFilter(ID int) CategoryFilter {
	return CategoryFilter{CategoryFilterID, ID}
}

func NewCategoryEmptyFilter() CategoryFilter {
	return CategoryFilter{CategoryFilterEmpty, ""}
}
