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

type Problem interface {
	Name() string

	Titles() []Content
	Statements() []Content
	HTMLStatements() []Content
	PDFStatements() []Content

	MemoryLimit() int
	TimeLimit() int
	InputOutputFiles() (string, string)
	Interactive() bool
	Languages() []language.Language
	Attachments() []Attachment
	Tags() []string

	Compile(language.Sandbox, language.Language, io.Reader, io.Writer) (io.Reader, error)
	Run(language.Sandbox, language.Language, io.Reader, chan string, chan Status) (Status, error)
}
