package problem_json

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/mraron/njudge/utils/language"
	"github.com/mraron/njudge/utils/problems"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
)

type Name struct {
	Language string
	Value    string
}

type Statement struct {
	Language string
	Path     string
	Type     string
}

type Attachment struct {
	Name string
	Path string
}

type Problem struct {
	Path                   string
	GeneratedStatementList problems.Contents
	AttachmentList         problems.Attachments

	ShortName      string `json:"shortname"`
	Names          []Name
	StatementList  []Statement  `json:"statements"`
	AttachmentInfo []Attachment `json:"attachments"`

	TestCount     int    `json:"testcount"`
	InputPattern  string `json:"inputpattern"`
	OutputPattern string `json:"outputpattern"`
}

func (p Problem) Name() string {
	return p.ShortName
}

func (p Problem) Titles() problems.Contents {
	ans := make(problems.Contents, len(p.Names))
	for ind, val := range p.Names {
		ans[ind] = problems.BytesData{Loc: val.Language, Val: []byte(val.Value), Typ: "text"}
	}

	return ans
}

func (p Problem) Statements() problems.Contents {
	return p.GeneratedStatementList
}

func (p Problem) HTMLStatements() problems.Contents {
	return p.GeneratedStatementList.FilterByType("text/html")
}

func (p Problem) PDFStatements() problems.Contents {
	return p.GeneratedStatementList.FilterByType("application/pdf")
}

func (p Problem) MemoryLimit() int {
	return 0
}

func (p Problem) TimeLimit() int {
	return 0
}

func (p Problem) InputOutputFiles() (string, string) {
	return "", ""
}

func (p Problem) Interactive() bool {
	return false
}

func (p Problem) Languages() []language.Language {
	return []language.Language{language.Get("zip")}
}

func (p Problem) Attachments() problems.Attachments {
	return p.AttachmentList
}

func (p Problem) Tags() []string {
	return []string{}
}

func (p Problem) StatusSkeleton(name string) (*problems.Status, error) {
	ans := problems.Status{false, "status skeleton", problems.FeedbackIOI, make([]problems.Testset, 0)}
	ans.Feedback = append(ans.Feedback, problems.Testset{Name: "tests"})
	ans.Feedback[0].Groups = make([]problems.Group, 1)
	ans.Feedback[0].Groups[0].Name = "base"
	ans.Feedback[0].Groups[0].Scoring = problems.ScoringSum

	for i := 0; i < p.TestCount; i++ {
		tc := problems.Testcase{}
		tc.InputPath, tc.AnswerPath = fmt.Sprintf(filepath.Join(p.Path, p.InputPattern), i+1), fmt.Sprintf(filepath.Join(p.Path, p.OutputPattern), i+1)
		tc.Index = i
		tc.Group = "base"

		ans.Feedback[0].Groups[0].Testcases = append(ans.Feedback[0].Groups[0].Testcases, tc)
	}

	return &ans, nil
}

func (p Problem) Check(tc *problems.Testcase) error {
	output := bytes.Buffer{}

	cmd := exec.Command(filepath.Join(p.Path, "check"), tc.InputPath, tc.OutputPath, tc.AnswerPath)
	cmd.Stdout = &output
	cmd.Stderr = &output

	err := cmd.Run()

	tc.CheckerOutput = output.String()

	if err == nil && cmd.ProcessState.Success() {
		return nil
	} else if err != nil {
		return err
	}

	return errors.New("proccess state is not success")
}

func (p Problem) Files() []problems.File {
	return make([]problems.File, 0)
}

func (p Problem) TaskTypeName() string {
	return "output_only"
}

func parser(path string) (problems.Problem, error) {
	problemJSON, err := os.Open(filepath.Join(path, "problem.json"))
	if err != nil {
		return nil, err
	}

	defer problemJSON.Close()

	p := Problem{}

	if err := json.NewDecoder(problemJSON).Decode(&p); err != nil {
		return nil, err
	}

	p.Path = path

	p.AttachmentList = make(problems.Attachments, len(p.AttachmentInfo))
	for ind, val := range p.AttachmentInfo {
		contents, err := ioutil.ReadFile(filepath.Join(path, val.Path))

		if err != nil {
			return nil, err
		}

		p.AttachmentList[ind] = problems.BytesData{Nam: val.Name, Val: contents}
	}

	p.GeneratedStatementList = make(problems.Contents, len(p.StatementList))
	for ind, val := range p.StatementList {
		contents, err := ioutil.ReadFile(filepath.Join(path, val.Path))

		if err != nil {
			return nil, err
		}

		p.GeneratedStatementList[ind] = problems.BytesData{Loc: val.Language, Val: contents, Typ: val.Type}
	}

	return p, nil
}

func identifier(path string) bool {
	_, err := os.Stat(filepath.Join(path, "problem.json"))
	return !os.IsNotExist(err)
}

func init() {
	problems.RegisterConfigType("problem_json", parser, identifier)
}
