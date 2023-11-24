package memory

import (
	"context"
	"sort"

	"github.com/mraron/njudge/internal/njudge"
)

type SubmissionListQuery struct {
	subs njudge.Submissions
	ps   njudge.Problems
}

func NewSubmissionListQuery(subs njudge.Submissions, ps njudge.Problems) *SubmissionListQuery {
	return &SubmissionListQuery{
		subs: subs,
		ps:   ps,
	}
}

func (s *SubmissionListQuery) filterProblemsetProblem(ctx context.Context, req njudge.SubmissionListRequest, sub njudge.Submission) (bool, error) {
	if req.Problemset != "" || req.Problem != "" {
		ok := true
		p, err := s.ps.Get(ctx, sub.ProblemID)
		if err != nil {
			return false, err
		}

		if req.Problemset != "" && req.Problemset != p.Problemset {
			ok = false
		} else if req.Problem != "" && req.Problem != p.Problem {
			ok = false
		}

		return ok, nil
	}

	return true, nil
}

func (s *SubmissionListQuery) filterVerdict(ctx context.Context, req njudge.SubmissionListRequest, sub njudge.Submission) (bool, error) {
	if req.Verdict != nil {
		return *req.Verdict == sub.Verdict, nil
	}

	return true, nil
}

func (s *SubmissionListQuery) filterUser(ctx context.Context, req njudge.SubmissionListRequest, sub njudge.Submission) (bool, error) {
	if req.UserID > 0 {
		return req.UserID == sub.UserID, nil
	}

	return true, nil
}

func (s *SubmissionListQuery) GetSubmissionList(ctx context.Context, req njudge.SubmissionListRequest) (*njudge.SubmissionList, error) {
	allSubmissions, err := s.subs.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	filters := []func(context.Context, njudge.SubmissionListRequest, njudge.Submission) (bool, error){
		s.filterProblemsetProblem,
		s.filterVerdict,
		s.filterUser,
	}

	submissions := make([]njudge.Submission, 0)
	for ind := range allSubmissions {
		ok := true
		for _, filter := range filters {
			currOk, err := filter(ctx, req, allSubmissions[ind])
			if err != nil {
				return nil, err
			}

			ok = ok && currOk
		}

		if ok {
			submissions = append(submissions, allSubmissions[ind])
		}
	}

	inv := req.SortDir == njudge.SortDESC
	sort.Slice(submissions, func(i, j int) bool {
		switch req.SortField {
		case njudge.SubmissionSortFieldScore:
			return inv != (submissions[i].Score < submissions[j].Score)
		default:
			return inv != (submissions[i].ID < submissions[j].ID)
		}
	})

	return &njudge.SubmissionList{
		Submissions: submissions,
	}, nil
}

func (s *SubmissionListQuery) GetPagedSubmissionList(ctx context.Context, req njudge.SubmissionListRequest) (*njudge.PagedSubmissionList, error) {
	ss, err := s.GetSubmissionList(ctx, req)
	if err != nil {
		return nil, err
	}

	submissions := ss.Submissions

	var pdata njudge.PaginationData
	submissions, pdata = Paginate(submissions, req.Page, req.PerPage)

	return &njudge.PagedSubmissionList{
		PaginationData: pdata,
		Submissions:    submissions,
	}, nil
}

func (s *SubmissionListQuery) getSubmissionListUserPredicate(ctx context.Context, userID int, pred func(njudge.Submission) bool, predExclude func(njudge.Submission) bool) (*njudge.SubmissionList, error) {
	allSubmissions, err := s.subs.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	last := make(map[int]njudge.Submission)
	exclude := make(map[int]bool)
	for ind := range allSubmissions {
		if allSubmissions[ind].UserID == userID {
			if pred(allSubmissions[ind]) {
				last[allSubmissions[ind].ProblemID] = allSubmissions[ind]
			}

			if predExclude(allSubmissions[ind]) {
				exclude[allSubmissions[ind].ProblemID] = true
			}
		}
	}

	submissions := make([]njudge.Submission, 0)
	for ind, sub := range last {
		if _, ok := exclude[ind]; !ok {
			submissions = append(submissions, sub)
		}
	}

	return &njudge.SubmissionList{
		Submissions: submissions,
	}, nil
}

func (s *SubmissionListQuery) GetAttemptedSubmissionList(ctx context.Context, userID int) (*njudge.SubmissionList, error) {
	return s.getSubmissionListUserPredicate(ctx, userID, func(s njudge.Submission) bool {
		return true
	}, func(s njudge.Submission) bool {
		return s.Verdict == njudge.VerdictAC
	})
}

func (s *SubmissionListQuery) GetSolvedSubmissionList(ctx context.Context, userID int) (*njudge.SubmissionList, error) {
	return s.getSubmissionListUserPredicate(ctx, userID, func(s njudge.Submission) bool {
		return s.Verdict == njudge.VerdictAC
	}, func(s njudge.Submission) bool {
		return false
	})
}
