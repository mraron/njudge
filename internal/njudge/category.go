package njudge

import (
	"context"
	"errors"
	"github.com/mraron/njudge/pkg/problems"
	"sort"

	"github.com/volatiletech/null/v8"
)

type Category struct {
	ID       int
	Name     string
	Visible  bool
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

type TaskArchiveNodeType string

const (
	TaskArchiveNodeRoot     TaskArchiveNodeType = "root"
	TaskArchiveNodeCategory TaskArchiveNodeType = "category"
	TaskArchiveNodeProblem  TaskArchiveNodeType = "problem"
)

type TaskArchiveNode struct {
	ID           int
	Type         TaskArchiveNodeType
	Category     *Category
	Problem      *Problem
	SolvedStatus SolvedStatus
	Children     []TaskArchiveNode
	Visible      bool
}

func (t TaskArchiveNode) Search(f func(node TaskArchiveNode) bool) {
	if f(t) {
		return
	}
	for _, child := range t.Children {
		child.Search(f)
	}
}

type TaskArchiveService struct {
	Categories        Categories
	Problems          Problems
	SolvedStatusQuery SolvedStatusQuery
	ProblemQuery      ProblemQuery
	ProblemStore      problems.Store
}

func (tas TaskArchiveService) problemNode(ctx context.Context, p Problem, u *User) (*TaskArchiveNode, error) {
	if !p.Visible && (u == nil || u.Role != "admin") {
		return nil, nil
	}

	curr := &TaskArchiveNode{
		ID:      p.ID,
		Type:    TaskArchiveNodeProblem,
		Problem: &p,
		Visible: p.Visible,
	}

	if u != nil {
		solvedStatus, err := tas.SolvedStatusQuery.GetSolvedStatus(ctx, p.ID, u.ID)
		if err != nil {
			return nil, err
		}
		curr.SolvedStatus = solvedStatus
	}

	return curr, nil
}

func (tas TaskArchiveService) categoryNode(ctx context.Context, c Category, u *User) (*TaskArchiveNode, error) {
	if !c.Visible && (u == nil || u.Role != "admin") {
		return nil, nil
	}
	curr := &TaskArchiveNode{
		ID:       c.ID,
		Type:     TaskArchiveNodeCategory,
		Category: &c,
		Visible:  c.Visible,
	}

	problemList, err := tas.ProblemQuery.GetProblemsWithCategory(ctx, NewCategoryIDFilter(curr.ID))
	if err != nil {
		return nil, err
	}
	sort.Slice(problemList, func(i, j int) bool {
		return problemList[i].ID < problemList[j].ID
	})
	for _, problem := range problemList {
		if p, err := tas.problemNode(ctx, problem, u); err != nil {
			return nil, err
		} else {
			curr.Children = append(curr.Children, *p)
		}
	}

	subCategories, err := tas.Categories.GetAllWithParent(ctx, curr.ID)
	if err != nil {
		return nil, err
	}
	sort.Slice(subCategories, func(i, j int) bool {
		//TODO add order attribute
		return subCategories[i].Name < subCategories[j].Name
	})
	for _, category := range subCategories {
		if cat, err := tas.categoryNode(ctx, category, u); err != nil {
			return nil, err
		} else {
			curr.Children = append(curr.Children, *cat)
		}
	}
	return curr, nil
}

func (tas TaskArchiveService) Create(ctx context.Context, categories []Category, u *User) (*TaskArchiveNode, error) {
	root := &TaskArchiveNode{
		ID:      0,
		Type:    TaskArchiveNodeRoot,
		Visible: true,
	}

	for _, cat := range categories {
		if curr, err := tas.categoryNode(ctx, cat, u); err != nil {
			return nil, err
		} else {
			root.Children = append(root.Children, *curr)
		}
	}

	return root, nil
}

func (tas TaskArchiveService) CreateFromAllTopLevel(ctx context.Context, u *User) (*TaskArchiveNode, error) {
	categories, err := tas.Categories.GetAllWithParent(ctx, 0)
	if err != nil {
		return nil, err
	}
	return tas.Create(ctx, categories, u)
}
