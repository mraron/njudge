package njudge

import (
	"context"
	"errors"
	"sync"
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

type MemorySubmissions struct {
	sync.Mutex
	nextId int
	data   []Submission
}

func NewMemorySubmissions() *MemorySubmissions {
	return &MemorySubmissions{
		nextId: 1,
		data:   make([]Submission, 0),
	}
}

func (m *MemorySubmissions) Get(ctx context.Context, ID int) (*Submission, error) {
	m.Lock()
	defer m.Unlock()
	for ind := range m.data {
		if m.data[ind].ID == ID {
			res := m.data[ind]
			return &res, nil
		}
	}

	return nil, ErrorSubmissionNotFound
}
func (m *MemorySubmissions) GetAll(ctx context.Context) ([]Submission, error) {
	m.Lock()
	defer m.Unlock()
	res := make([]Submission, len(m.data))
	copy(res, m.data)

	return res, nil
}

func (m *MemorySubmissions) Insert(ctx context.Context, s Submission) (*Submission, error) {
	m.Lock()
	defer m.Unlock()
	s.ID = m.nextId
	m.nextId++

	m.data = append(m.data, s)

	res := m.data[len(m.data)-1]
	return &res, nil
}

func (m *MemorySubmissions) Delete(ctx context.Context, ID int) error {
	m.Lock()
	defer m.Unlock()
	for ind := range m.data {
		if m.data[ind].ID == ID {
			m.data[ind] = m.data[len(m.data)-1]
			m.data = m.data[:len(m.data)-1]
			return nil
		}
	}

	return ErrorSubmissionNotFound
}

func (m *MemorySubmissions) Update(ctx context.Context, s Submission) error {
	m.Lock()
	defer m.Unlock()
	for ind := range m.data {
		if m.data[ind].ID == s.ID {
			m.data[ind] = s
			return nil
		}
	}
	return ErrorSubmissionNotFound
}
