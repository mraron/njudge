package problem_json

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/mraron/njudge/utils/language"
	"github.com/mraron/njudge/utils/problems"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
)

//@TODO Make this work!!! :(

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
	GeneratedStatementList []problems.Content
	AttachmentList         []problems.Attachment

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

func (p Problem) Titles() []problems.Content {
	ans := make([]problems.Content, len(p.Names))
	for ind, val := range p.Names {
		ans[ind] = problems.Content{val.Language, []byte(val.Value), "text"}
	}

	return ans
}

func (p Problem) Statements() []problems.Content {
	return p.GeneratedStatementList
}

func (p Problem) HTMLStatements() []problems.Content {
	return problems.FilterContentArray(p.GeneratedStatementList, "text/html")
}

func (p Problem) PDFStatements() []problems.Content {
	return problems.FilterContentArray(p.GeneratedStatementList, "application/pdf")
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

func (p Problem) Attachments() []problems.Attachment {
	return p.AttachmentList
}

func (p Problem) Tags() []string {
	return []string{}
}

func (p Problem) StatusSkeleton() problems.Status {
	ans := problems.Status{false, "status skeleton", problems.FEEDBACK_IOI, make([]problems.Testset, 0)}
	ans.Feedback = append(ans.Feedback, problems.Testset{Name: "main"})
	ans.Feedback[0].Groups = make([]problems.Group, 1)
	ans.Feedback[0].Groups[0].Name = "base"
	ans.Feedback[0].Groups[0].Scoring = problems.SCORING_SUM

	for i := 0; i < p.TestCount; i++ {
		tc := problems.Testcase{}
		tc.InputPath, tc.AnswerPath = fmt.Sprintf(filepath.Join(p.Path, p.InputPattern), i+1), fmt.Sprintf(filepath.Join(p.Path, p.OutputPattern), i+1)
		tc.Index = i
		tc.Group = "base"

		ans.Feedback[0].Testcases = append(ans.Feedback[0].Testcases, tc)
		ans.Feedback[0].Groups[0].Testcases = append(ans.Feedback[0].Groups[0].Testcases, tc)
	}

	return ans
}

func (p Problem) Check(tc *problems.Testcase, stdout io.Writer, stderr io.Writer) error {
	cmd := exec.Command(filepath.Join(p.Path, "check"), tc.InputPath, tc.OutputPath, tc.AnswerPath)
	cmd.Stdout = stdout
	cmd.Stderr = stderr

	err := cmd.Run()

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

	p.AttachmentList = make([]problems.Attachment, len(p.AttachmentInfo))
	for ind, val := range p.AttachmentInfo {
		contents, err := ioutil.ReadFile(filepath.Join(path, val.Path))

		if err != nil {
			return nil, err
		}

		p.AttachmentList[ind] = problems.Attachment{val.Name, contents}
	}

	p.GeneratedStatementList = make([]problems.Content, len(p.StatementList))
	for ind, val := range p.StatementList {
		contents, err := ioutil.ReadFile(filepath.Join(path, val.Path))

		if err != nil {
			return nil, err
		}

		p.GeneratedStatementList[ind] = problems.Content{Locale: val.Language, Contents: contents, Type: val.Type}
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
