package njudge

import "context"

type SubmissionSortField string

var (
	SubmissionSortFieldID    SubmissionSortField = "id"
	SubmissionSortFieldScore SubmissionSortField = "score"
)

type SubmissionListRequest struct {
	Problemset string
	Problem    string
	SortDir    SortDirection
	SortField  SubmissionSortField
	Page       int
	PerPage    int
	Verdict    *Verdict
	UserID     int
}

type PagedSubmissionList struct {
	PaginationData PaginationData
	Submissions    []Submission
}

type SubmissionList struct {
	Submissions []Submission
}

type SubmissionListQuery interface {
	GetPagedSubmissionList(ctx context.Context, req SubmissionListRequest) (*PagedSubmissionList, error)

	GetSubmissionList(ctx context.Context, req SubmissionListRequest) (*SubmissionList, error)
	GetAttemptedSubmissionList(ctx context.Context, userID int) (*SubmissionList, error)
	GetSolvedSubmissionList(ctx context.Context, userID int) (*SubmissionList, error)
}
