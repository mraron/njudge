package polygon

import (
	"path/filepath"

	"github.com/mraron/njudge/pkg/language"
	"github.com/mraron/njudge/pkg/problems"
)

type Name struct {
	Language string `xml:"language,attr"`
	Value    string `xml:"value,attr"`
}

type Statement struct {
	Language string `xml:"language,attr"`
	Path     string `xml:"path,attr"`
	Type     string `xml:"type,attr"`
}

type Attachment struct {
	Name     string `xml:"name,attr"`
	Location string `xml:"location,attr"`
}

type Stub struct {
	Name     string `xml:"name,attr"`
	Path     string `xml:"path,attr"`
	Language string `xml:"language,attr"`
}

type Source struct {
	Path string `xml:"path,attr"`
	Type string `xml:"type,attr"`
}

type Checker struct {
	Type   string `xml:"type,attr"`
	Source Source `xml:"source"`
}

type Interactor struct {
	Source Source `xml:"source"`
}

type Assets struct {
	Attachments []Attachment `xml:"attachments>attachment"`
	Stubs       []Stub       `xml:"stubs>stub"`
	Checker     Checker      `xml:"checker"`
	Interactor  Interactor   `xml:"interactor"`
}

type Problem struct {
	config *config

	Path                   string
	JSONStatementList      []JSONStatement
	AttachmentsList        problems.Attachments
	GeneratedStatementList problems.Contents

	TaskType      string      `xml:"njudge-task-type,attr"`
	FeedbackType  string      `xml:"njudge-feedback-type,attr"`
	Revision      int         `xml:"revision,attr"`
	ShortName     string      `xml:"short-name,attr"`
	Url           string      `xml:"url,attr"`
	Names         []Name      `xml:"names>name"`
	StatementList []Statement `xml:"statements>statement"`
	Judging       Judging     `xml:"judging"`
	Assets        Assets      `xml:"assets"`
	TagsList      []struct {
		Value string `xml:"value,attr"`
	} `xml:"tags>tag"`
}

func (p Problem) Name() string {
	return p.ShortName
}

func (p Problem) Titles() problems.Contents {
	ans := make(problems.Contents, len(p.Names))
	for i := 0; i < len(p.Names); i++ {
		ans[i] = problems.BytesData{Loc: p.Names[i].Language, Val: []byte(p.Names[i].Value), Typ: "text"}
	}

	return ans
}

func (p Problem) Statements() problems.Contents {
	return p.GeneratedStatementList
}

func (p Problem) HTMLStatements() problems.Contents {
	return p.GeneratedStatementList.FilterByType(problems.DataTypeHTML)
}

func (p Problem) PDFStatements() problems.Contents {
	return p.GeneratedStatementList.FilterByType(problems.DataTypePDF)
}

func (p Problem) MemoryLimit() int {
	return p.Judging.Testsets[0].MemoryLimit
}

func (p Problem) TimeLimit() int {
	return p.Judging.Testsets[0].TimeLimit
}

func (p Problem) InputOutputFiles() (string, string) {
	return p.Judging.InputFile, p.Judging.OutputFile
}

func (p Problem) Attachments() problems.Attachments {
	return p.AttachmentsList
}

func (p Problem) Tags() (lst []string) {
	lst = make([]string, len(p.TagsList))
	for ind, val := range p.TagsList {
		lst[ind] = val.Value
	}

	return
}

func (Problem) Languages() []language.Language {
	lst1 := language.DefaultStore.List()

	lst2 := make([]language.Language, 0, len(lst1))
	for _, val := range lst1 {
		if val.Id() != "zip" {
			lst2 = append(lst2, val)
		}
	}

	return lst2
}

func (p Problem) Files() []problems.File {
	res := make([]problems.File, 0)
	for _, stub := range p.Assets.Stubs {
		res = append(res, problems.File{Name: stub.Name, Role: "stub_" + stub.Language, Path: filepath.Join(p.Path, stub.Path)})
	}

	if p.Assets.Interactor.Source.Path != "" {
		res = append(res, problems.File{Name: "interactor", Role: "interactor", Path: filepath.Join(p.Path, "files", "interactor")})
	}

	return res
}

func (p Problem) GetTaskType() problems.TaskType {
	if p.Assets.Interactor.Source.Path != "" {
		p.TaskType = "communication"
	}

	if p.TaskType == "" {
		p.TaskType = "batch"
	}

	tt, err := problems.GetTaskType(p.TaskType)
	if err != nil {
		panic(err)
	}

	return tt
}
