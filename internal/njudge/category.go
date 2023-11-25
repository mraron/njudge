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

func NewCategory(name string, parent *Category) *Category {
	if parent == nil {
		return &Category{
			Name: name,
			ParentID: null.Int{
				Valid: false,
			},
		}
	}
	return &Category{
		Name: name,
		ParentID: null.Int{
			Int:   parent.ID,
			Valid: true,
		},
	}
}

var (
	ErrorCategoryNotFound = errors.New("njudge: category not found")
)

type Categories interface {
	GetAll(ctx context.Context) ([]Category, error)
	GetAllWithParent(ctx context.Context, parentID int) ([]Category, error)
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
