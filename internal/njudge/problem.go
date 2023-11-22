package njudge

import (
	"context"
	"errors"
	"io"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/mraron/njudge/internal/web/domain/submission"
	"github.com/mraron/njudge/internal/web/helpers/i18n"
	"github.com/mraron/njudge/pkg/problems"
	"github.com/mraron/njudge/pkg/problems/config/polygon"
)

var (
	ErrorProblemNotFound    = errors.New("njudge: problem not found")
	ErrorFileNotFound       = errors.New("njudge: file not found")
	ErrorStatementNotFound  = errors.New("njudge: statement not found")
	ErrorProblemTagNotFound = errors.New("njudge: problem tag not found")
)

type Problem struct {
	ID         int
	Problemset string
	Problem    string
	Category   *Category
	Tags       []ProblemTag
}

func NewProblem(problemset, problem string) Problem {
	return Problem{0, problemset, problem, nil, nil}
}

func (p *Problem) SetCategory(c Category) {
	p.Category = &c
}

func (p *Problem) WithStoredData(store problems.Store) (ProblemStoredData, error) {
	pp, err := store.Get(p.Problem)
	if err != nil {
		return nil, err
	}

	return &problemStoredData{pp}, nil
}

func (p *Problem) HasTag(t Tag) bool {
	for ind := range p.Tags {
		if p.Tags[ind].Tag.ID == t.ID {
			return true
		}
	}

	return false
}

func (p *Problem) AddTag(t Tag, userID int) error {
	if p.HasTag(t) {
		return nil
	}

	p.Tags = append(p.Tags, ProblemTag{
		ProblemID: p.ID,
		Tag:       t,
		UserID:    userID,
		Added:     time.Now(),
	})
	return nil
}

func (p *Problem) DeleteTag(t Tag) error {
	if !p.HasTag(t) {
		return ErrorProblemTagNotFound
	}

	for ind := range p.Tags {
		if p.Tags[ind].Tag.ID == t.ID {
			p.Tags[ind] = p.Tags[len(p.Tags)-1]
			p.Tags = p.Tags[:len(p.Tags)-1]
			return nil
		}
	}

	panic("")
}

type Problems interface {
	Get(ctx context.Context, ID int) (*Problem, error)
	GetAll(ctx context.Context) ([]Problem, error)
	Insert(ctx context.Context, p Problem) (*Problem, error)
	Delete(ctx context.Context, ID int) error
	Update(ctx context.Context, p Problem) error
}

type MemoryProblems struct {
	sync.Mutex
	nextId int
	data   []Problem

	nextProblemTagId int
}

func NewMemoryProblems() *MemoryProblems {
	return &MemoryProblems{
		nextId:           1,
		data:             make([]Problem, 0),
		nextProblemTagId: 1,
	}
}

func (m *MemoryProblems) Get(ctx context.Context, ID int) (*Problem, error) {
	m.Lock()
	defer m.Unlock()
	for ind := range m.data {
		if m.data[ind].ID == ID {
			res := m.data[ind]
			res.Tags = make([]ProblemTag, len(res.Tags))
			copy(res.Tags, m.data[ind].Tags)
			return &res, nil
		}
	}

	return nil, ErrorProblemNotFound
}

func (m *MemoryProblems) GetAll(ctx context.Context) ([]Problem, error) {
	m.Lock()
	defer m.Unlock()
	res := make([]Problem, len(m.data))
	copy(res, m.data)
	for ind := range res {
		res[ind].Tags = make([]ProblemTag, len(res[ind].Tags))
		copy(res[ind].Tags, m.data[ind].Tags)
	}

	return res, nil
}

func (m *MemoryProblems) Insert(ctx context.Context, p Problem) (*Problem, error) {
	m.Lock()
	defer m.Unlock()
	p.ID = m.nextId
	m.nextId++

	for ind := range p.Tags {
		p.Tags[ind].ID = m.nextProblemTagId
		m.nextProblemTagId++
	}

	m.data = append(m.data, p)

	res := m.data[len(m.data)-1]
	return &res, nil
}

func (m *MemoryProblems) Delete(ctx context.Context, id int) error {
	m.Lock()
	defer m.Unlock()
	for ind := range m.data {
		if m.data[ind].ID == id {
			m.data[ind] = m.data[len(m.data)-1]
			m.data = m.data[:len(m.data)-1]
			return nil
		}
	}

	return ErrorProblemNotFound
}

func (m *MemoryProblems) Update(ctx context.Context, p Problem) error {
	m.Lock()
	defer m.Unlock()
	for ind := range m.data {
		if m.data[ind].ID == p.ID {
			for tagInd := range p.Tags {
				if p.Tags[tagInd].ID == 0 {
					p.Tags[tagInd].ID = m.nextProblemTagId
					m.nextProblemTagId++
				}
			}

			m.data[ind] = p
			return nil
		}
	}
	return ErrorProblemNotFound
}

type ProblemStoredData interface {
	problems.Problem

	GetPDF(lang Language) (io.ReadCloser, error)
	GetFile(filename string) (string, error)
	GetAttachment(name string) (problems.NamedData, error)
}

type problemStoredData struct {
	problems.Problem
}

func (p *problemStoredData) GetPDF(lang Language) (io.ReadCloser, error) {
	if len(p.Statements().FilterByType(problems.DataTypePDF)) == 0 {
		return nil, ErrorStatementNotFound
	}

	return i18n.TranslateContent(string(lang), p.Statements().FilterByType(problems.DataTypePDF)).ValueReader()
}

func (p *problemStoredData) GetFile(file string) (fileLoc string, err error) {
	switch p := p.Problem.(problems.ProblemWrapper).Problem.(type) {
	case polygon.Problem:
		if len(p.Statements().FilterByType(problems.DataTypeHTML)) == 0 || strings.HasSuffix(file, ".tex") || strings.HasSuffix(file, ".json") {
			err = ErrorFileNotFound
		}

		fileLoc = filepath.Join(p.Path, "statements", ".html", p.HTMLStatements()[0].Locale(), file)
		if _, err := os.Stat(fileLoc); err != nil {
			fileLoc = filepath.Join(p.Path, "statements", p.HTMLStatements()[0].Locale(), file)
		}
	default:
		err = ErrorFileNotFound
	}

	return
}

func (p *problemStoredData) GetAttachment(attachment string) (problems.NamedData, error) {
	for _, val := range p.Attachments() {
		if val.Name() == attachment {
			return val, nil
		}
	}

	return nil, ErrorFileNotFound
}

type ProblemUserData struct {
	SolvedStatus SolvedStatus
	LastLanguage string
}

type ProblemData struct {
	SolverCount int
	UserData    *ProblemUserData
	Tags        []Tag
	Submissions []submission.Submission
}
