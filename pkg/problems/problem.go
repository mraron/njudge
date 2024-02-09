// Package problems contains utilities useful for parsing, displaying, judging (submission for) problems
package problems

import (
	"github.com/mraron/njudge/pkg/language"
	"golang.org/x/net/context"
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
	MemoryLimit() int
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
