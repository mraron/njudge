package problems

import (
	"github.com/mraron/njudge/utils/language"
	"io"
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
	Check(testcase *Testcase) error
	InputOutputFiles() (string, string)
	Languages() []language.Language
	StatusSkeleton(testset string) (*Status, error)
	Files() []File
	TaskTypeName() string
}

type TaskType interface {
	Name() string
	Compile(Judgeable, language.Sandbox, language.Language, io.Reader, io.Writer) (io.Reader, error)
	Run(Judgeable, *language.SandboxProvider, language.Language, io.Reader, chan string, chan Status) (Status, error)
}

func Truncate(s string) string {
	if len(s) < 256 {
		return s
	}

	return s[:255] + "..."
}
