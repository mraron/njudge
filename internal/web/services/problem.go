package services

import (
	"context"
	"database/sql"
	"github.com/mraron/njudge/internal/web/domain/problem"
	"github.com/mraron/njudge/internal/web/domain/submission"
	"github.com/mraron/njudge/internal/web/helpers"
	"github.com/mraron/njudge/internal/web/helpers/ui"
	"github.com/mraron/njudge/internal/web/models"
	"github.com/mraron/njudge/pkg/problems"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type SQLProblem struct {
	db *sql.DB
	ps problems.Store
}

func NewSQLProblem(db *sql.DB, ps problems.Store) *SQLProblem {
	return &SQLProblem{db, ps}
}

func (s SQLProblem) GetStatsData(ctx context.Context, p problem.Problem, userID int) (*problem.StatsData, error) {
	res := problem.StatsData{}

	var err error
	res.SolverCount = p.ProblemRel.SolverCount
	if userID != 0 {
		if res.LastLanguage, err = helpers.GetUserLastLanguage(ctx, s.db, userID); err != nil {
			return nil, err
		}
		if res.SolvedStatus, err = helpers.HasUserSolved(s.db, userID, p.Problemset, p.ProblemRel.Problem); err != nil {
			return nil, err
		}

		if submissions, err := models.Submissions(models.SubmissionWhere.Problemset.EQ(p.Problemset),
			models.SubmissionWhere.Problem.EQ(p.ProblemRel.Problem), models.SubmissionWhere.UserID.EQ(userID),
			qm.OrderBy("id DESC"), qm.Limit(5)).All(ctx, s.db); err != nil {
			return nil, err
		} else {
			res.Submissions = make([]submission.Submission, len(submissions))
			for ind := range submissions {
				res.Submissions[ind] = submission.Submission{Submission: *submissions[ind]}
			}
		}
	}

	if p.ProblemRel.CategoryID.Valid {
		res.CategoryID = p.ProblemRel.CategoryID.Int
		res.CategoryLink, err = ui.TopCategoryLink(ctx, s.db, res.CategoryID)
		if err != nil {
			return nil, err
		}
	}

	tags, err := models.Tags(qm.InnerJoin("problem_tags pt on pt.tag_id = tags.id"), qm.Where("pt.problem_id = ?", p.ID)).All(ctx, s.db)
	if err != nil {
		return nil, err
	}

	res.Tags = make([]problem.Tag, len(tags))
	for ind := range tags {
		res.Tags[ind] = problem.Tag{Tag: *tags[ind]}
	}
	return &res, nil
}

func (s SQLProblem) Get(ctx context.Context, ID int) (*problem.Problem, error) {
	pr, err := models.ProblemRels(models.ProblemRelWhere.ID.EQ(ID)).One(ctx, s.db)
	if err != nil {
		return nil, err
	}

	p, err := s.ps.Get(pr.Problem)
	if err != nil {
		return nil, err
	}

	return &problem.Problem{Problem: p, ProblemRel: *pr}, nil
}

func (s SQLProblem) GetAll(ctx context.Context) ([]problem.Problem, error) {
	prs, err := models.ProblemRels().All(ctx, s.db)
	if err != nil {
		return nil, err
	}

	var res []problem.Problem
	for ind := range prs {
		p, err := s.ps.Get(prs[ind].Problem)
		if err != nil {
			return nil, err
		}
		res = append(res, problem.Problem{
			Problem:    p,
			ProblemRel: *prs[ind],
		})
	}

	return res, nil
}

func (s SQLProblem) GetByNames(ctx context.Context, problemset string, problemName string) (*problem.Problem, error) {
	pr, err := models.ProblemRels(models.ProblemRelWhere.Problemset.EQ(problemset), models.ProblemRelWhere.Problem.EQ(problemName)).One(ctx, s.db)
	if err != nil {
		return nil, err
	}

	p, err := s.ps.Get(pr.Problem)
	if err != nil {
		return nil, err
	}

	return &problem.Problem{Problem: p, ProblemRel: *pr}, nil
}
