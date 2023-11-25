package services

import (
	"context"

	"github.com/mraron/njudge/internal/njudge/db/models"
	"github.com/mraron/njudge/internal/web/domain/submission"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type RejudgeService interface {
	Rejudge(ctx context.Context, ID int) error
}

func (s SQLSubmission) Rejudge(ctx context.Context, ID int) error {
	sub := models.Submission{ID: ID, Judged: null.Time{Valid: false}, Started: false}
	rowsAff, err := sub.Update(ctx, s.db, boil.Whitelist(models.SubmissionColumns.ID, models.SubmissionColumns.Judged, models.SubmissionColumns.Started))
	if err != nil {
		return err
	}
	if rowsAff == 0 {
		return submission.ErrorSubmissionNotFound
	}

	return nil
}
