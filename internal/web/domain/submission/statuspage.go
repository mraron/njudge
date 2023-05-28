package submission

import (
	"github.com/mraron/njudge/internal/web/helpers/pagination"
)

type StatusPage struct {
	CurrentPage int
	Pages       []pagination.Link
	Submissions []Submission
}
