package njudge

import (
	"context"
	"errors"
	"github.com/mraron/njudge/internal/web/templates/i18n"
	"github.com/mraron/njudge/pkg/language"
	"io"
	"time"

	"github.com/mraron/njudge/pkg/problems"
)

var (
	ErrorProblemNotFound    = errors.New("njudge: problem not found")
	ErrorFileNotFound       = errors.New("njudge: file not found")
	ErrorStatementNotFound  = errors.New("njudge: statement not found")
	ErrorProblemTagNotFound = errors.New("njudge: problem tag not found")
)

func Fields(s ...string) []string {
	return s
}

var ProblemFields = struct {
	ID          string
	Problemset  string
	Problem     string
	Category    string
	SolverCount string
	Visible     string

	Tags string
}{
	ID:          "id",
	Problemset:  "problemset",
	Problem:     "problem",
	Category:    "category",
	SolverCount: "solver_count",
	Visible:     "visible",

	Tags: "tags",
}

type Problem struct {
	ID          int
	Problemset  string
	Problem     string
	Category    *Category
	SolverCount int
	Visible     bool

	Tags ProblemTags
}

func NewProblem(problemset, problem string) Problem {
	return Problem{
		ID:          0,
		Problemset:  problemset,
		Problem:     problem,
		Category:    nil,
		SolverCount: 0,

		Tags: nil,
	}
}

func (p *Problem) SetCategory(c Category) {
	p.Category = &c
}

func (p *Problem) WithStoredData(store problems.Store) (ProblemStoredData, error) {
	pp, err := store.GetProblem(p.Problem)
	if err != nil {
		return ProblemStoredData{}, err
	}

	return ProblemStoredData{pp}, nil
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
	Update(ctx context.Context, p Problem, fields []string) error
}

type ProblemStoredData struct {
	problems.Problem
}

func (p *ProblemStoredData) GetLanguage(id string) language.Language {
	for _, lang := range p.Languages() {
		if lang.ID() == id {
			return lang
		}
	}
	return nil
}

func (p *ProblemStoredData) GetPDF(lang Language) (io.ReadCloser, error) {
	if len(p.Statements().FilterByType(problems.DataTypePDF)) == 0 {
		return nil, ErrorStatementNotFound
	}

	return i18n.TranslateContent(string(lang), p.Statements().FilterByType(problems.DataTypePDF)).ValueReader()
}

func (p *ProblemStoredData) GetAttachment(attachment string) (problems.NamedData, error) {
	for _, val := range p.Attachments() {
		if val.Name() == attachment {
			return val, nil
		}
	}

	return nil, ErrorFileNotFound
}

type ProblemUserInfo struct {
	SolvedStatus SolvedStatus
	LastLanguage string
	Submissions  []Submission
}

type ProblemInfo struct {
	UserInfo *ProblemUserInfo
}

type ProblemInfoQuery interface {
	GetProblemData(ctx context.Context, problemID, userID int) (*ProblemInfo, error)
}

type SolvedStatusQuery interface {
	GetSolvedStatus(ctx context.Context, problemID, userID int) (SolvedStatus, error)
}

type ProblemQuery interface {
	GetProblem(ctx context.Context, problemset, problem string) (*Problem, error)
	GetProblemsWithCategory(ctx context.Context, f CategoryFilter) ([]Problem, error)
}
