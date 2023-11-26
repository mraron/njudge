package db

import (
	"context"
	"database/sql"
	"errors"
	"sort"

	"github.com/mraron/njudge/internal/njudge"
	"github.com/mraron/njudge/internal/njudge/db/models"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type SubmissionListQuery struct {
	db *sql.DB
}

func NewSubmissionListQuery(db *sql.DB) *SubmissionListQuery {
	return &SubmissionListQuery{
		db: db,
	}
}

func (s *SubmissionListQuery) getSubmissionList(ctx context.Context, req njudge.SubmissionListRequest, mods ...qm.QueryMod) (*njudge.SubmissionList, int64, error) {
	mods = append(mods, qm.OrderBy(string(req.SortField)+" "+string(req.SortDir)))

	if req.UserID > 0 {
		mods = append(mods, models.SubmissionWhere.UserID.EQ(req.UserID))
	}

	if req.Verdict != nil {
		mods = append(mods, models.SubmissionWhere.Verdict.EQ(int(*req.Verdict)))
	}

	if req.Problemset != "" || req.Problem != "" {
		mods = append(
			mods,
			qm.InnerJoin("problem_rels pr on pr.id = submissions.problem_id"),
		)
	}
	if req.Problemset != "" {
		mods = append(
			mods,
			qm.Where("pr.problemset=?", req.Problemset),
		)
	}

	if req.Problem != "" {
		mods = append(
			mods,
			qm.Where("pr.problem=?", req.Problem),
		)
	}

	res, err := NewSubmissions(s.db).getAll(ctx, mods...)
	if err != nil {
		return nil, -1, err
	}

	mods = append(mods, qm.GroupBy("submissions.id"))
	cnt, err := models.Submissions(mods...).Count(ctx, s.db)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, -1, err
	}

	return &njudge.SubmissionList{
		Submissions: res,
	}, cnt, nil
}

func (s *SubmissionListQuery) GetSubmissionList(ctx context.Context, req njudge.SubmissionListRequest) (*njudge.SubmissionList, error) {
	res, _, err := s.getSubmissionList(ctx, req)
	return res, err
}

func (s *SubmissionListQuery) GetPagedSubmissionList(ctx context.Context, req njudge.SubmissionListRequest) (*njudge.PagedSubmissionList, error) {
	submissions, cnt, err := s.getSubmissionList(
		ctx,
		req,
		qm.Limit(req.PerPage),
		qm.Offset((req.Page-1)*req.PerPage),
	)

	if err != nil {
		return nil, err
	}

	return &njudge.PagedSubmissionList{
		PaginationData: njudge.PaginationData{
			Page:    req.Page,
			PerPage: req.PerPage,
			Pages:   (int(cnt) + req.PerPage - 1) / req.PerPage,
			Count:   int(cnt),
		},
		Submissions: submissions.Submissions,
	}, nil
}

func (s *SubmissionListQuery) getSubmissionListUserPredicate(ctx context.Context, userID int, pred func(njudge.Submission) bool, predExclude func(njudge.Submission) bool) (*njudge.SubmissionList, error) {
	allSubmissions, _, err := s.getSubmissionList(ctx, njudge.SubmissionListRequest{
		SortDir:   njudge.SortDESC,
		SortField: njudge.SubmissionSortFieldID,
		UserID:    userID,
	})

	if err != nil {
		return nil, err
	}

	last := make(map[int]njudge.Submission)
	exclude := make(map[int]bool)
	for ind := range allSubmissions.Submissions {
		if pred(allSubmissions.Submissions[ind]) {
			last[allSubmissions.Submissions[ind].ProblemID] = allSubmissions.Submissions[ind]
		}

		if predExclude(allSubmissions.Submissions[ind]) {
			exclude[allSubmissions.Submissions[ind].ProblemID] = true
		}

	}

	submissions := make([]njudge.Submission, 0)
	for ind, sub := range last {
		if _, ok := exclude[ind]; !ok {
			submissions = append(submissions, sub)
		}
	}

	sort.Slice(submissions, func(i, j int) bool {
		return submissions[i].Submitted.Before(submissions[j].Submitted)
	})

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
