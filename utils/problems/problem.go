package problems

import (
	"github.com/mraron/njudge/utils/language"
	"github.com/spf13/afero"
	"io"
)

type Content struct {
	Locale   string
	Contents []byte
	Type     string
}

func (s Content) IsText() bool {
	return s.Type == "text"
}

func (s Content) IsHTML() bool {
	return s.Type == "text/html"
}

func (s Content) IsPDF() bool {
	return s.Type == "application/pdf"
}

func (s Content) String() string {
	return string(s.Contents)
}

type Contents []Content

func (cs Contents) Locales() []string {
	lst := make(map[string]bool)
	for _, val := range cs {
		lst[val.Locale] = true
	}

	ans := make([]string, len(lst))

	ind := 0
	for key := range lst {
		ans[ind] = key
		ind++
	}

	return ans
}

func (cs Contents) FilterByType(mime string) Contents {
	res := make(Contents, 0)
	for _, val := range cs {
		if mime == val.Type {
			res = append(res, val)
		}
	}

	return res
}

func (cs Contents) FilterByLocale(locale string) Contents {
	res := make(Contents, 0)
	for _, val := range cs {
		if locale == val.Locale {
			res = append(res, val)
		}
	}

	return res
}

type Attachment struct {
	Name     string
	Contents []byte
}

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

	Attachments() []Attachment
	Tags() []string

	Judgeable
}

type Judgeable interface {
	MemoryLimit() int
	TimeLimit() int
	Check(*Testcase) error
	InputOutputFiles() (string, string)
	Languages() []language.Language
	StatusSkeleton() Status
	Files() []File
	TaskTypeName() string
}

type Validatable interface {
	Validate(*Testcase) (bool, error)
}

type Exportable interface {
	Export(afero.Fs, string) error
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
