// Package problems contains utilities useful for parsing, displaying, judging (a submission for) problems
package problems

import (
	"context"
	"github.com/mraron/njudge/pkg/language"
	"github.com/mraron/njudge/pkg/language/memory"
	"strings"
)

type Problem interface {
	Name() string
	Titles() Contents
	Statements() Contents
	MemoryLimit() memory.Amount
	TimeLimit() int

	Attachments() Attachments
	Tags() []string

	EvaluationInfo
}

type EvaluationFile struct {
	Name string
	Role string
	Path string
}

func (f EvaluationFile) StubOf(lang language.Language) bool {
	return f.Role == "stub_"+lang.ID() || (f.Role == "stub_cpp" && strings.HasPrefix(lang.ID(), "cpp"))
}

type EvaluationInfo interface {
	InputOutputFiles() (string, string)
	Languages() []language.Language
	StatusSkeleton(testset string) (*Status, error)
	EvaluationFiles() []EvaluationFile
	GetTaskType() TaskType
}

type Checker interface {
	Check(ctx context.Context, testcase *Testcase) error
}

func Truncate(s string) string {
	if len(s) < 256 {
		return s
	}

	return s[:255] + "..."
}
