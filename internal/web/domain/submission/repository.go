package submission

import (
	"context"
	"errors"
)

var ErrorSubmissionNotFound = errors.New("submission not found")

type Repository interface {
	Get(ctx context.Context, ID int) (*Submission, error)
}
