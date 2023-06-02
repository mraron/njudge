package glue

import (
	"github.com/mraron/njudge/internal/judge"
	"github.com/mraron/njudge/internal/web/models"
)

type JudgeFinder interface {
	FindJudge([]*models.Judge, *models.Submission) (*models.Judge, error)
}

type FindJudgerNaive struct{}

func (FindJudgerNaive) FindJudge(judges []*models.Judge, sub *models.Submission) (*models.Judge, error) {
	for _, j := range judges {
		st, err := judge.ParseServerStatus(j.State)
		if err != nil {
			return nil, err
		}

		if j.Online && st.SupportsProblem(sub.Problem) {
			return j, nil
		}
	}

	return nil, nil
}
