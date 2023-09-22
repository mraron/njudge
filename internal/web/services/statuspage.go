package services

import (
	"context"
	"database/sql"
	"github.com/mraron/njudge/internal/web/domain/submission"
	"github.com/mraron/njudge/internal/web/helpers/pagination"
	"github.com/mraron/njudge/internal/web/models"
	"github.com/mraron/njudge/pkg/problems"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type StatusPageRequest struct {
	Pagination pagination.Data

	Problemset string
	Problem    string
	Verdict    *problems.VerdictName
	UserID     int
}

type StatusPageService interface {
	GetStatusPage(ctx context.Context, req StatusPageRequest) (*submission.StatusPage, error)
}

type SQLStatusPageService struct {
	db *sql.DB
}

func NewSQLStatusPageService(db *sql.DB) *SQLStatusPageService {
	return &SQLStatusPageService{db}
}

func (s SQLStatusPageService) GetStatusPage(ctx context.Context, req StatusPageRequest) (*submission.StatusPage, error) {
	order := qm.OrderBy(req.Pagination.SortField + " " + req.Pagination.SortDir)

	var query []qm.QueryMod
	if req.Problemset != "" {
		query = append(query, models.SubmissionWhere.Problemset.EQ(req.Problemset))
	}
	if req.Problem != "" {
		query = append(query, models.SubmissionWhere.Problem.EQ(req.Problem))
	}
	if req.Verdict != nil {
		query = append(query, models.SubmissionWhere.Verdict.EQ(int(*req.Verdict)))
	}
	if req.UserID > 0 {
		query = append(query, models.SubmissionWhere.UserID.EQ(req.UserID))
	}

	sbs, err := models.Submissions(append(append(
		[]qm.QueryMod{
			qm.Limit(req.Pagination.PerPage),
			qm.Offset((req.Pagination.Page - 1) * req.Pagination.PerPage),
		}, query...), order)...).All(ctx, s.db)
	if err != nil {
		return nil, err
	}

	cnt, err := models.Submissions(query...).Count(ctx, s.db)
	if err != nil {
		return nil, err
	}

	req.Pagination.LastPage = (int(cnt) + req.Pagination.PerPage - 1) / req.Pagination.PerPage

	var submissions []submission.Submission
	for ind := range sbs {
		submissions = append(submissions, submission.Submission{Submission: *sbs[ind]})
	}

	return &submission.StatusPage{
		PaginationData: req.Pagination,
		Submissions:    submissions,
	}, nil
}
