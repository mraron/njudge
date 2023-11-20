package njudge

import (
	"context"
	"errors"
	"time"

	"github.com/volatiletech/null/v8"
)

type Submission struct {
	ID int

	UserID    int
	ProblemID int
	Language  string
	Source    []byte
	Private   bool

	Started   bool
	Verdict   Verdict
	Ontest    null.String
	Submitted time.Time
	Status    string
	Judged    null.Time
}

var ErrorSubmissionNotFound = errors.New("submission not found")

type SubmissionRepository interface {
	Get(ctx context.Context, ID int) (*Submission, error)
}
