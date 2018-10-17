package problems

import (
	"github.com/mraron/njudge/utils/language"
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

func FilterContentArray(arr []Content, mime string) (res []Content) {
	res = make([]Content, 0)
	for _, val := range arr {
		if mime == val.Type {
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

//@TODO add problem config type name

type Problem interface {
	Titles() []Content
	Statements() []Content
	HTMLStatements() []Content
	PDFStatements() []Content

	Attachments() []Attachment
	Tags() []string

	JudgingInformation
}

type JudgingInformation interface {
	Name() string
	MemoryLimit() int
	TimeLimit() int
	Check(string, string, string, io.Writer, io.Writer) error
	InputOutputFiles() (string, string)
	Languages() []language.Language
	StatusSkeleton() Status
	Files() []File
	TaskTypeName() string
}

type TaskType interface {
	Name() string
	Compile(JudgingInformation, language.Sandbox, language.Language, io.Reader, io.Writer) (io.Reader, error)
	Run(JudgingInformation, language.Sandbox, language.Language, io.Reader, chan string, chan Status) (Status, error)
}
