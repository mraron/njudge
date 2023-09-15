package problem

import "github.com/mraron/njudge/internal/web/models"

type ProblemTag struct {
	models.ProblemTag
}

type Tag struct {
	models.Tag
}

type Tags []Tag

func (t Tags) ToStringSlice() []string {
	res := make([]string, len(t))
	for ind := range t {
		res[ind] = t[ind].Name
	}

	return res
}
