package services

import (
	"context"
	"database/sql"
	"errors"

	"github.com/mraron/njudge/internal/njudge/db/models"
	"github.com/mraron/njudge/internal/web/domain/submission"
)

type SQLSubmission struct {
	db *sql.DB
}

func NewSQLSubmission(db *sql.DB) SQLSubmission {
	return SQLSubmission{db}
}

func (s SQLSubmission) Get(ctx context.Context, ID int) (*submission.Submission, error) {
	res, err := models.Submissions(models.SubmissionWhere.ID.EQ(ID)).One(ctx, s.db)
	if errors.Is(sql.ErrNoRows, err) {
		return nil, submission.ErrorSubmissionNotFound
	}

	return &submission.Submission{Submission: *res}, nil
}
