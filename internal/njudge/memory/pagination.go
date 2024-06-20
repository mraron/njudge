package memory

import "github.com/mraron/njudge/internal/njudge"

func Paginate[T any](data []T, page, perPage int) ([]T, njudge.PaginationData) {
	minPage, maxPage := 1, (len(data)+perPage-1)/perPage
	if page > maxPage {
		page = maxPage
	}
	if page < minPage {
		page = minPage
	}
	start, end := (page-1)*perPage, page*perPage
	if end >= len(data) {
		end = len(data)
	}

	res := make([]T, end-start)
	for i := start; i < end; i++ {
		res[i-start] = data[i]
	}

	return res, njudge.PaginationData{
		Page:    page,
		PerPage: perPage,
		Count:   len(data),
		Pages:   maxPage,
	}
}
