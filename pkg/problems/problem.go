// Package problems contains utilities useful for parsing, displaying, judging (submission for) problems
package problems

import (
	"context"
	"github.com/mraron/njudge/pkg/language"
	"github.com/mraron/njudge/pkg/language/memory"
)

type File struct {
	Name string
	Role string
	Path string
}

type Problem interface {
	Name() string
	Titles() Contents
	Statements() Contents
	MemoryLimit() memory.Amount
	TimeLimit() int

	Attachments() Attachments
	Tags() []string

	Judgeable
}

type Judgeable interface {
	Checker() Checker
	InputOutputFiles() (string, string)
	Languages() []language.Language
	StatusSkeleton(testset string) (*Status, error)
	Files() []File
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
