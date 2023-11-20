package njudge

type Category struct {
	ID       int
	Name     string
	ParentID int
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
