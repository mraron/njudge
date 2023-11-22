package njudge

import (
	"context"
	"errors"
	"time"

	"github.com/mraron/njudge/pkg/language"
	"github.com/mraron/njudge/pkg/problems"
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
	Status    problems.Status
	Judged    null.Time
}

func NewSubmission(u User, p Problem, language language.Language) (*Submission, error) {
	return &Submission{
		UserID:    u.ID,
		ProblemID: p.ID,
		Language:  language.Id(),
		Source:    []byte(""),
		Private:   false,

		Started: false,
		Verdict: VerdictUP,
		Ontest: null.String{
			Valid: false,
		},
		Submitted: time.Now(),
		Status:    problems.Status{},
		Judged: null.Time{
			Valid: false,
		},
	}, nil
}

func (s *Submission) SetSource(src []byte) {
	s.Source = src
}

var ErrorSubmissionNotFound = errors.New("njudge: submission not found")

type Submissions interface {
	Get(ctx context.Context, ID int) (*Submission, error)
	GetAll(ctx context.Context) ([]Submission, error)
	Insert(ctx context.Context, s Submission) (*Submission, error)
	Delete(ctx context.Context, ID int) error
	Update(ctx context.Context, s Submission) error
}
