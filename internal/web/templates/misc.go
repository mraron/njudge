package templates

import (
	"github.com/a-h/templ"
	"github.com/mraron/njudge/internal/njudge"
)

type Link struct {
	Text string
	Href templ.SafeURL
}

type SortColumn struct {
	Order njudge.SortDirection
	Href  string
}
