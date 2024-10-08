package polygon

import (
	"errors"
	"github.com/mraron/njudge/pkg/language/langs/zip"
	"github.com/mraron/njudge/pkg/language/memory"
	"github.com/mraron/njudge/pkg/problems/evaluation"
	"github.com/mraron/njudge/pkg/problems/evaluation/batch"
	"github.com/mraron/njudge/pkg/problems/evaluation/communication"
	"github.com/mraron/njudge/pkg/problems/evaluation/output_only"
	"github.com/mraron/njudge/pkg/problems/evaluation/stub"
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
	binary []byte
}

type Assets struct {
	Attachments []Attachment `xml:"attachments>attachment"`
	Stubs       []Stub       `xml:"stubs>stub"`
	Checker     Checker      `xml:"checker"`
	Interactor  Interactor   `xml:"interactor"`
}

type Problem struct {
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

func (p Problem) Valid() error {
	switch p.TaskType {
	case batch.Name:
	case communication.Name:
	case output_only.Name:
	case stub.Name:
	case "":
	default:
		return errors.New("unknown task type")
	}
	if problems.FeedbackTypeFromShortString(p.FeedbackType) == problems.FeedbackUnknown {
		return errors.New("unknown feedback type")
	}
	return nil
}

func (p Problem) Name() string {
	return p.ShortName
}

func (p Problem) Titles() problems.Contents {
	ans := make(problems.Contents, len(p.Names))
	for i := 0; i < len(p.Names); i++ {
		ans[i] = problems.BytesData{Loc: p.Names[i].Language, Val: []byte(p.Names[i].Value), Typ: problems.DataTypeText}
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

func (p Problem) MemoryLimit() memory.Amount {
	return memory.Amount(p.Judging.Testsets[0].MemoryLimit)
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

func (p Problem) Languages() []language.Language {
	if p.TaskType == output_only.Name {
		return []language.Language{zip.Zip{}}
	}

	return language.ListExcept(language.DefaultStore, []string{"zip"})
}

func (p Problem) EvaluationFiles() []problems.EvaluationFile {
	res := make([]problems.EvaluationFile, 0)
	for _, stub := range p.Assets.Stubs {
		res = append(res, problems.EvaluationFile{Name: stub.Name, Role: "stub_" + stub.Language, Path: filepath.Join(p.Path, stub.Path)})
	}

	if p.Assets.Interactor.Source.Path != "" {
		res = append(res, problems.EvaluationFile{Name: "interactor", Role: "interactor", Path: filepath.Join(p.Path, "files", "interactor")})
	}

	return res
}

func (p Problem) GetTaskType() problems.TaskType {
	if p.Assets.Interactor.Source.Path != "" {
		return communication.New(evaluation.CompileCheckSupported{
			List:         p.Languages(),
			NextCompiler: evaluation.Compile{},
		}, p.Assets.Interactor.binary, p.Checker())
	}
	if p.TaskType == output_only.Name {
		return output_only.New(p.Checker())
	}
	if p.TaskType == stub.Name {
		compiler := evaluation.NewCompilerWithStubs()
		for _, lang := range p.Languages() {
			for _, file := range p.EvaluationFiles() {
				if file.StubOf(lang) {
					compiler.AddStub(lang, file)
				}
			}
		}
		return stub.New(compiler, evaluation.BasicRunnerWithChecker(p.Checker()))
	}
	return batch.New(evaluation.CompileCheckSupported{
		List:         p.Languages(),
		NextCompiler: evaluation.Compile{},
	}, evaluation.BasicRunnerWithChecker(p.Checker()))
}
