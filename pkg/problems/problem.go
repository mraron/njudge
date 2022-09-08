package problems

import (
	"github.com/mraron/njudge/pkg/language"
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
	HTMLStatements() Contents
	PDFStatements() Contents

	Attachments() Attachments
	Tags() []string

	Judgeable
}

type Judgeable interface {
	MemoryLimit() int
	TimeLimit() int
	Checker() Checker
	InputOutputFiles() (string, string)
	Languages() []language.Language
	StatusSkeleton(testset string) (*Status, error)
	Files() []File
	GetTaskType() TaskType
}

type Checker interface {
	Name() string
	Check(testcase *Testcase) error
}

func Truncate(s string) string {
	if len(s) < 256 {
		return s
	}

	return s[:255] + "..."
}
