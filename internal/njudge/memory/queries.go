package memory

import (
	"context"
	"sort"

	"github.com/mraron/njudge/internal/njudge"
)

type ProblemInfoQuery struct {
	ss njudge.Submissions
}

func NewProblemInfoQuery(ss njudge.Submissions) *ProblemInfoQuery {
	return &ProblemInfoQuery{
		ss: ss,
	}
}

func (p *ProblemInfoQuery) GetProblemData(ctx context.Context, problemID int, userID int) (*njudge.ProblemInfo, error) {
	submissions, err := p.ss.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	res := &njudge.ProblemInfo{}

	userSolved := make(map[int]struct{})
	userTried := make(map[int]struct{})
	for ind := range submissions {
		if submissions[ind].ProblemID != problemID {
			continue
		}

		userTried[submissions[ind].UserID] = struct{}{}
		if submissions[ind].Verdict == njudge.VerdictAC {
			userSolved[submissions[ind].UserID] = struct{}{}
		}
	}

	if userID > 0 {
		res.UserInfo = &njudge.ProblemUserInfo{}
		res.UserInfo.SolvedStatus = njudge.Unattempted
		if _, ok := userTried[userID]; ok {
			res.UserInfo.SolvedStatus = njudge.Attempted
			if _, ok := userSolved[userID]; ok {
				res.UserInfo.SolvedStatus = njudge.Solved
			}
		}

		lastSubmissionInd := -1
		res.UserInfo.Submissions = make([]njudge.Submission, 0)
		for ind := range submissions {
			if submissions[ind].ProblemID != problemID || submissions[ind].UserID != userID {
				continue
			}

			if lastSubmissionInd == -1 || submissions[lastSubmissionInd].Submitted.Before(submissions[ind].Submitted) {
				lastSubmissionInd = ind
			}
			res.UserInfo.Submissions = append(res.UserInfo.Submissions, submissions[ind])
		}

		if lastSubmissionInd != -1 {
			res.UserInfo.LastLanguage = submissions[lastSubmissionInd].Language
		}
	}

	return res, nil
}

type ProblemQuery struct {
	pp njudge.Problems
}

func NewProblemQuery(pp njudge.Problems) *ProblemQuery {
	return &ProblemQuery{
		pp: pp,
	}
}

func (p *ProblemQuery) GetProblem(ctx context.Context, problemset string, problem string) (*njudge.Problem, error) {
	problems, err := p.pp.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	for ind := range problems {
		if problems[ind].Problemset == problemset && problems[ind].Problem == problem {
			return &problems[ind], nil
		}
	}

	return nil, njudge.ErrorProblemNotFound
}

func (p *ProblemQuery) GetProblemsWithCategory(ctx context.Context, f njudge.CategoryFilter) ([]njudge.Problem, error) {
	problems, err := p.pp.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	res := make([]njudge.Problem, 0)
	for ind := range problems {
		ok := true
		switch f.Type {
		case njudge.CategoryFilterEmpty:
			ok = problems[ind].Category == nil
		case njudge.CategoryFilterID:
			ok = problems[ind].Category != nil && problems[ind].Category.ID == f.Value.(int)
		default:
		}

		if ok {
			res = append(res, problems[ind])
		}
	}

	return res, nil
}

type SubmissionsQuery struct {
	submissions njudge.Submissions
}

func NewSubmissionsQuery(submissions njudge.Submissions) *SubmissionsQuery {
	return &SubmissionsQuery{
		submissions: submissions,
	}
}

func (s *SubmissionsQuery) GetUnstarted(ctx context.Context, limit int) ([]njudge.Submission, error) {
	submissions, err := s.submissions.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	unstarted := make([]njudge.Submission, 0)
	for ind := range submissions {
		if !submissions[ind].Started {
			unstarted = append(unstarted, submissions[ind])
		}
	}

	sort.Slice(unstarted, func(i, j int) bool {
		return unstarted[i].Submitted.Before(unstarted[j].Submitted)
	})

	if len(unstarted) < limit {
		limit = len(unstarted)
	}

	return unstarted[:limit], nil
}
