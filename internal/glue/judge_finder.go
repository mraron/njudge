package glue

import (
	"github.com/mraron/njudge/internal/judge"
	"github.com/mraron/njudge/internal/njudge/db/models"
)

type JudgeFinder interface {
	FindJudge(judges []*models.Judge, problem string) (*models.Judge, error)
}

type FindJudgerNaive struct{}

func (FindJudgerNaive) FindJudge(judges []*models.Judge, problem string) (*models.Judge, error) {
	for _, j := range judges {
		st, err := judge.ParseServerStatus(j.State)
		if err != nil {
			return nil, err
		}

		if j.Online && st.SupportsProblem(problem) {
			return j, nil
		}
	}

	return nil, nil
}
