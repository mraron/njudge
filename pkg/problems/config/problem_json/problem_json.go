package problem_json

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/gomarkdown/markdown"
	"github.com/mraron/njudge/pkg/language"
	"github.com/mraron/njudge/pkg/problems"
	"github.com/mraron/njudge/pkg/problems/checker"
	"github.com/spf13/afero"
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
	TestCount     int      `json:"test_count"`
	InputList     []string `json:"input_list"`
	OutputList    []string `json:"output_list"`
	InputPattern  string   `json:"input_pattern"`
	OutputPattern string   `json:"output_pattern"`
	Scoring       string   `json:"scoring"`
	MaxScore      float64  `json:"max_score"`
}

type Checker struct {
	Type string `json:"type"`
	Path string `json:"location"`
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
		InputFile     string    `json:"input_file"`
		OutputFile    string    `json:"output_file"`
		MemoryLimit   int       `json:"memory_limit"`
		TimeLimit     float64   `json:"time_limit"`
		TestCount     int       `json:"test_count"`
		InputPattern  string    `json:"input_pattern"`
		OutputPattern string    `json:"output_pattern"`
		InputList     []string  `json:"input_list"`
		OutputList    []string  `json:"output_list"`
		Checker       Checker   `json:"checker"`
		FeedbackType  string    `json:"feedback_type"`
		Subtasks      []Subtask `json:"subtasks"`
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
	return p.Tests.InputFile, p.Tests.OutputFile
}

func (p Problem) Languages() []language.Language {
	lst1 := language.List()

	lst2 := make([]language.Language, 0, len(lst1))
	for _, val := range lst1 {
		if val.Id() != "zip" {
			lst2 = append(lst2, val)
		}
	}

	return lst2
}

func (p Problem) Attachments() problems.Attachments {
	return p.AttachmentList
}

func (p Problem) Tags() []string {
	return []string{}
}

func (p Problem) StatusSkeleton(name string) (*problems.Status, error) {
	ans := problems.Status{Compiled: false, CompilerOutput: "status skeleton", FeedbackType: problems.FeedbackFromString(p.Tests.FeedbackType), Feedback: make([]problems.Testset, 0)}
	ans.Feedback = append(ans.Feedback, problems.Testset{Name: "tests"})

	getIthIO := func(typ string, index int, pattern string, list []string) (string, error) {
		if pattern != "" {
			return fmt.Sprintf(filepath.Join(p.Path, "tests", pattern), index+1), nil
		} else {
			if index < len(list) {
				return filepath.Join(p.Path, "tests", list[index]), nil
			} else {
				return "", fmt.Errorf("too few %s %d >= %d", typ, index, len(list))
			}
		}
	}

	if p.Tests.Subtasks == nil {
		ans.Feedback[0].Groups = make([]problems.Group, 1)
		ans.Feedback[0].Groups[0].Name = "base"
		ans.Feedback[0].Groups[0].Scoring = problems.ScoringSum

		if p.Tests.InputList != nil {
			p.Tests.TestCount = len(p.Tests.InputList)
		}

		for i := 0; i < p.Tests.TestCount; i++ {
			var err error

			tc := problems.Testcase{}
			if tc.InputPath, err = getIthIO("input", i, p.Tests.InputPattern, p.Tests.InputList); err != nil {
				return nil, err
			}
			if tc.AnswerPath, err = getIthIO("output", i, p.Tests.OutputPattern, p.Tests.OutputList); err != nil {
				return nil, err
			}

			tc.Index = i
			tc.Group = "base"

			ans.Feedback[0].Groups[0].Testcases = append(ans.Feedback[0].Groups[0].Testcases, tc)
		}
	} else {
		ans.Feedback[0].Groups = make([]problems.Group, len(p.Tests.Subtasks))
		for s := 0; s < len(p.Tests.Subtasks); s++ {
			ans.Feedback[0].Groups[s].Name = fmt.Sprintf("subtask%d", s+1)
			ans.Feedback[0].Groups[s].Scoring = problems.ScoringFromString(p.Tests.Subtasks[s].Scoring)

			if p.Tests.Subtasks[s].InputList != nil {
				p.Tests.Subtasks[s].TestCount = len(p.Tests.Subtasks[s].InputList)
			}

			for i := 0; i < p.Tests.Subtasks[s].TestCount; i++ {
				var err error

				tc := problems.Testcase{}
				if tc.InputPath, err = getIthIO("input", i, p.Tests.Subtasks[s].InputPattern, p.Tests.Subtasks[s].InputList); err != nil {
					return nil, err
				}
				if tc.AnswerPath, err = getIthIO("output", i, p.Tests.Subtasks[s].OutputPattern, p.Tests.Subtasks[s].OutputList); err != nil {
					return nil, err
				}

				if ans.Feedback[0].Groups[s].Scoring == problems.ScoringSum {
					tc.MaxScore = p.Tests.Subtasks[s].MaxScore / float64(p.Tests.Subtasks[s].TestCount)
				} else if i == 0 {
					tc.MaxScore = p.Tests.Subtasks[s].MaxScore
				}

				tc.Index = i
				tc.Group = ans.Feedback[0].Groups[s].Name

				ans.Feedback[0].Groups[s].Testcases = append(ans.Feedback[0].Groups[s].Testcases, tc)
			}
		}
	}

	return &ans, nil
}

func (p Problem) Checker() problems.Checker {
	if p.Tests.Checker.Type == "" || p.Tests.Checker.Type == "whitediff" {
		return checker.Whitediff{}
	} else if p.Tests.Checker.Type == "testlib" {
		return checker.NewTestlib(filepath.Join(p.Path, p.Tests.Checker.Path))
	} else if p.Tests.Checker.Type == "taskyaml" {
		return checker.NewTaskYAML(filepath.Join(p.Path, p.Tests.Checker.Path))
	}

	return checker.Noop{}
}

func (p Problem) Files() []problems.File {
	return make([]problems.File, 0)
}

func (p Problem) GetTaskType() problems.TaskType {
	tt, err := problems.GetTaskType("batch")
	if err != nil {
		panic(err)
	}

	return tt
}

type config struct {
	fs              afero.Fs
	compileBinaries bool
}

func newConfig() *config {
	return &config{fs: afero.NewOsFs(), compileBinaries: true}
}

type Option func(*config)

func UseFS(fs afero.Fs) Option {
	return func(c *config) {
		c.fs = fs
	}
}

func CompileBinaries(compile bool) Option {
	return func(c *config) {
		c.compileBinaries = compile
	}
}

func ParserAndIdentifier(opts ...Option) (problems.ConfigParser, problems.ConfigIdentifier) {
	cfg := newConfig()
	for _, opt := range opts {
		opt(cfg)
	}

	parser := func(path string) (problems.Problem, error) {
		problemJSON, err := cfg.fs.Open(filepath.Join(path, "problem.json"))
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
			contents, err := afero.ReadFile(cfg.fs, filepath.Join(path, val.Path))

			if err != nil {
				return nil, err
			}

			p.AttachmentList[ind] = problems.BytesData{Nam: val.Name, Val: contents}
		}

		p.GeneratedStatementList = make(problems.Contents, len(p.StatementList))
		for ind, val := range p.StatementList {
			var (
				contents []byte
				err      error
			)

			contents, err = afero.ReadFile(cfg.fs, filepath.Join(path, val.Path))
			if err != nil {
				return nil, err
			}

			if val.Type == "text/markdown" {
				contents = markdown.ToHTML(contents, nil, nil)
				val.Type = "text/html"
			}

			p.GeneratedStatementList[ind] = problems.BytesData{Loc: val.Language, Val: contents, Typ: val.Type}
		}

		return p, nil
	}

	identifier := func(path string) bool {
		_, err := cfg.fs.Stat(filepath.Join(path, "problem.json"))
		return !os.IsNotExist(err)
	}

	return parser, identifier
}

func init() {
	parser, identifier := ParserAndIdentifier()
	problems.RegisterConfigType("problem_json", parser, identifier)
}
