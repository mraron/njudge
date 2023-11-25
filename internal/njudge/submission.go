package njudge

import (
	"context"
	"errors"
	"time"

	"github.com/mraron/njudge/pkg/language"
	"github.com/mraron/njudge/pkg/problems"
	"github.com/volatiletech/null/v8"
)

var SubmissionFields = struct {
	ID string

	UserID    string
	ProblemID string
	Language  string
	Source    string
	Private   string

	Started   string
	Verdict   string
	Ontest    string
	Submitted string
	Status    string
	Judged    string
	Score     string
}{
	ID:        "id",
	UserID:    "user_id",
	ProblemID: "problem_id",
	Language:  "language",
	Source:    "source",
	Private:   "private",
	Started:   "started",
	Verdict:   "verdict",
	Ontest:    "ontest",
	Submitted: "submitted",
	Status:    "status",
	Judged:    "judged",
	Score:     "score",
}

var SubmissionRejudgeFields = []string{SubmissionFields.Judged, SubmissionFields.Started}

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
	Score     float32
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

func (s *Submission) MarkForRejudge() {
	s.Judged.Valid = false
	s.Started = false
}

func (s *Submission) GetUser(ctx context.Context, us Users) (*User, error) {
	return us.Get(ctx, s.UserID)
}

func (s *Submission) GetProblem(ctx context.Context, ps Problems) (*Problem, error) {
	return ps.Get(ctx, s.ProblemID)
}

var ErrorSubmissionNotFound = errors.New("njudge: submission not found")

type Submissions interface {
	Get(ctx context.Context, ID int) (*Submission, error)
	GetAll(ctx context.Context) ([]Submission, error)
	Insert(ctx context.Context, s Submission) (*Submission, error)
	Delete(ctx context.Context, ID int) error
	Update(ctx context.Context, s Submission, fields []string) error
}

var ErrorUnsupportedLanguage = errors.New("njudge: unsupported language")

type SubmitRequest struct {
	UserID     int
	Problemset string
	Problem    string
	Language   string
	Source     []byte
}

type SubmitService interface {
	Submit(ctx context.Context, subRequest SubmitRequest) (*Submission, error)
}
