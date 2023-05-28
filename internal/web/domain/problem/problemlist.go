package problem

import "github.com/mraron/njudge/internal/web/helpers/pagination"

type PaginationList struct {
	Pages    []pagination.Link
	Problems []Problem
}
