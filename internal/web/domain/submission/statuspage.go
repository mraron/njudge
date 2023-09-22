package submission

import "github.com/mraron/njudge/internal/web/helpers/pagination"

type StatusPage struct {
	PaginationData pagination.Data
	Submissions    []Submission
}
