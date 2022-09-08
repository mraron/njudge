package problem_json

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/mraron/njudge/pkg/language"
	"github.com/mraron/njudge/pkg/problems"
	"github.com/mraron/njudge/pkg/problems/checker"
)

type Title struct {
	Language string
	Title    string
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

type Subtask struct {
	TestCount  int      `json:"test_count"`
	InputList  []string `json:"input_list"`
	OutputList []string `json:"output_list"`
	Scoring    string   `json:"scoring"`
	MaxScore   float64  `json:"max_score"`
}

type Problem struct {
	Path                   string               `json:"-"`
	GeneratedStatementList problems.Contents    `json:"-"`
	AttachmentList         problems.Attachments `json:"-"`

	ShortName      string       `json:"shortname"`
	TitleList      []Title      `json:"titles"`
	StatementList  []Statement  `json:"statements"`
	AttachmentInfo []Attachment `json:"attachments"`

	Tests struct {
		InputFile     string   `json:"input_file"`
		OutputFile    string   `json:"output_file"`
		MemoryLimit   int      `json:"memory_limit"`
		TimeLimit     float64  `json:"time_limit"`
		TestCount     int      `json:"test_count"`
		InputPattern  string   `json:"input_pattern"`
		OutputPattern string   `json:"output_pattern"`
		InputList     []string `json:"input_list"`
		OutputList    []string `json:"output_list"`

		FeedbackType string    `json:"feedback_type"`
		Subtasks     []Subtask `json:"subtasks"`
	} `json:"tests"`
}

func (p Problem) Name() string {
	return p.ShortName
}

func (p Problem) Titles() problems.Contents {
	ans := make(problems.Contents, len(p.TitleList))
	for ind, val := range p.TitleList {
		ans[ind] = problems.BytesData{Loc: val.Language, Val: []byte(val.Title), Typ: "text"}
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
	return p.Tests.MemoryLimit * 1024 * 1024
}

func (p Problem) TimeLimit() int {
	return int(1000 * p.Tests.TimeLimit)
}

func (p Problem) InputOutputFiles() (string, string) {
	return "", ""
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
	ans := problems.Status{Compiled: false, CompilerOutput: "status skeleton", FeedbackType: problems.FeedbackIOI, Feedback: make([]problems.Testset, 0)}
	ans.Feedback = append(ans.Feedback, problems.Testset{Name: "tests"})
	ans.Feedback[0].Groups = make([]problems.Group, 1)
	ans.Feedback[0].Groups[0].Name = "base"
	ans.Feedback[0].Groups[0].Scoring = problems.ScoringSum

	for i := 0; i < p.Tests.TestCount; i++ {
		tc := problems.Testcase{}
		tc.InputPath, tc.AnswerPath = fmt.Sprintf(filepath.Join(p.Path, p.Tests.InputPattern), i+1), fmt.Sprintf(filepath.Join(p.Path, p.Tests.OutputPattern), i+1)
		tc.Index = i
		tc.Group = "base"

		ans.Feedback[0].Groups[0].Testcases = append(ans.Feedback[0].Groups[0].Testcases, tc)
	}

	return &ans, nil
}

func (p Problem) Checker() problems.Checker {
	return checker.Noop{}
}

func (p Problem) Files() []problems.File {
	return make([]problems.File, 0)
}

func (p Problem) GetTaskType() problems.TaskType {
	tt, err := problems.GetTaskType("outputonly")
	if err != nil {
		panic(err)
	}

	return tt
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
