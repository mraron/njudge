package problem_yaml

import (
	"bytes"
	"context"
	"fmt"
	"github.com/mraron/njudge/pkg/language"
	"github.com/mraron/njudge/pkg/language/langs/cpp"
	"github.com/mraron/njudge/pkg/language/sandbox"
	"github.com/mraron/njudge/pkg/problems"
	"github.com/mraron/njudge/pkg/problems/evaluation"
	"github.com/mraron/njudge/pkg/problems/evaluation/checker"
	"github.com/spf13/afero"
	"github.com/yuin/goldmark"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
	"strings"
	"time"
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
	TestCount     int      `yaml:"test_count"`
	InputList     []string `yaml:"input_list"`
	OutputList    []string `yaml:"output_list"`
	InputPattern  string   `yaml:"input_pattern"`
	OutputPattern string   `yaml:"output_pattern"`
	Scoring       string   `yaml:"scoring"`
	MaxScore      float64  `yaml:"max_score"`
}

type Checker struct {
	Type string `yaml:"type"`
	Path string `yaml:"path"`
}

type Problem struct {
	Path                   string               `yaml:"-"`
	GeneratedStatementList problems.Contents    `yaml:"-"`
	AttachmentList         problems.Attachments `yaml:"-"`

	ShortName      string       `yaml:"shortname"`
	TitleList      []Title      `yaml:"titles"`
	StatementList  []Statement  `yaml:"statements"`
	AttachmentInfo []Attachment `yaml:"attachments"`

	Tests struct {
		TaskType      string    `yaml:"task_type"`
		InputFile     string    `yaml:"input_file"`
		OutputFile    string    `yaml:"output_file"`
		MemoryLimit   int       `yaml:"memory_limit"`
		TimeLimit     float64   `yaml:"time_limit"`
		TestCount     int       `yaml:"test_count"`
		InputPattern  string    `yaml:"input_pattern"`
		OutputPattern string    `yaml:"output_pattern"`
		InputList     []string  `yaml:"input_list"`
		OutputList    []string  `yaml:"output_list"`
		Checker       Checker   `yaml:"checker"`
		FeedbackType  string    `yaml:"feedback_type"`
		Subtasks      []Subtask `yaml:"subtasks"`
	} `yaml:"tests"`
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
	if p.Tests.TaskType == "outputonly" {
		return []language.Language{language.DefaultStore.Get("zip")}
	}

	return language.ListExcept(language.DefaultStore, []string{"zip"})
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

	getIthIO := func(typ string, index int, pattern string, gindex int, gpattern string, list []string) (string, error) {
		if pattern != "" {
			return fmt.Sprintf(filepath.Join(p.Path, "tests", pattern), index+1), nil
		} else if gpattern != "" {
			return fmt.Sprintf(filepath.Join(p.Path, "tests", gpattern), gindex+1), nil
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
			tc.VerdictName = problems.VerdictDR

			if tc.InputPath, err = getIthIO("input", i, p.Tests.InputPattern, -1, "", p.Tests.InputList); err != nil {
				return nil, err
			}
			if tc.AnswerPath, err = getIthIO("output", i, p.Tests.OutputPattern, -1, "", p.Tests.OutputList); err != nil {
				return nil, err
			}
			if p.Tests.TaskType == "outputonly" {
				tc.OutputPath = filepath.Base(tc.AnswerPath)
			}

			tc.Index = i + 1
			tc.Group = "base"

			ans.Feedback[0].Groups[0].Testcases = append(ans.Feedback[0].Groups[0].Testcases, tc)
		}

		ans.Feedback[0].SetTimeLimit(time.Duration(p.TimeLimit()) * time.Millisecond)
		ans.Feedback[0].SetMemoryLimit(p.MemoryLimit())

	} else {
		ans.Feedback[0].Groups = make([]problems.Group, len(p.Tests.Subtasks))
		globalIdx := 0
		for s := 0; s < len(p.Tests.Subtasks); s++ {
			ans.Feedback[0].Groups[s].Name = fmt.Sprintf("subtask%d", s+1)
			ans.Feedback[0].Groups[s].Scoring = problems.ScoringFromString(p.Tests.Subtasks[s].Scoring)

			if p.Tests.Subtasks[s].InputList != nil {
				p.Tests.Subtasks[s].TestCount = len(p.Tests.Subtasks[s].InputList)
			}

			for i := 0; i < p.Tests.Subtasks[s].TestCount; i++ {
				var err error

				tc := problems.Testcase{}
				tc.VerdictName = problems.VerdictDR

				if tc.InputPath, err = getIthIO("input", i, p.Tests.Subtasks[s].InputPattern, globalIdx, p.Tests.InputPattern, p.Tests.Subtasks[s].InputList); err != nil {
					return nil, err
				}
				if tc.AnswerPath, err = getIthIO("output", i, p.Tests.Subtasks[s].OutputPattern, globalIdx, p.Tests.OutputPattern, p.Tests.Subtasks[s].OutputList); err != nil {
					return nil, err
				}
				if p.Tests.TaskType == "outputonly" {
					tc.OutputPath = filepath.Base(tc.AnswerPath)
				}

				if ans.Feedback[0].Groups[s].Scoring == problems.ScoringSum {
					tc.MaxScore = p.Tests.Subtasks[s].MaxScore / float64(p.Tests.Subtasks[s].TestCount)
				} else if i == 0 {
					tc.MaxScore = p.Tests.Subtasks[s].MaxScore
				}

				tc.Index = i + 1
				tc.Group = ans.Feedback[0].Groups[s].Name

				ans.Feedback[0].Groups[s].Testcases = append(ans.Feedback[0].Groups[s].Testcases, tc)
				globalIdx++
			}
		}

		ans.Feedback[0].SetTimeLimit(time.Duration(p.TimeLimit()) * time.Millisecond)
		ans.Feedback[0].SetMemoryLimit(p.MemoryLimit())
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
	return problems.NewTaskType("batch", evaluation.CompileCopyFile{}, evaluation.NewLinearEvaluator(evaluation.ACRunner{}))
	/*
		tasktype := "batch"
		if p.Tests.TaskType != "" {
			tasktype = p.Tests.TaskType
		}

		tt, err := problems.GetTaskType(tasktype)
		if err != nil {
			panic(err)
		}

		return tt*/
}

type config struct {
	compileBinaries bool
}

func newConfig() *config {
	return &config{compileBinaries: true}
}

type Option func(*config)

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

	parser := func(fs afero.Fs, path string) (problems.Problem, error) {
		problemYAML, err := fs.Open(filepath.Join(path, "problem.yaml"))
		if err != nil {
			return nil, err
		}

		defer problemYAML.Close()

		p := Problem{}

		if err := yaml.NewDecoder(problemYAML).Decode(&p); err != nil {
			return nil, err
		}

		p.Path = path

		p.AttachmentList = make(problems.Attachments, len(p.AttachmentInfo))
		for ind, val := range p.AttachmentInfo {
			contents, err := afero.ReadFile(fs, filepath.Join(path, val.Path))

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

			contents, err = afero.ReadFile(fs, filepath.Join(path, val.Path))
			if err != nil {
				return nil, err
			}

			if val.Type == "text/markdown" {
				res := &bytes.Buffer{}
				if err := goldmark.Convert(contents, res); err != nil {
					return nil, err
				}
				contents = res.Bytes()
				val.Type = problems.DataTypeHTML
			}

			p.GeneratedStatementList[ind] = problems.BytesData{Loc: val.Language, Val: contents, Typ: val.Type}
		}

		if strings.HasSuffix(p.Tests.Checker.Path, ".cpp") {
			binaryName := strings.TrimSuffix(p.Tests.Checker.Path, filepath.Ext(p.Tests.Checker.Path))
			s, _ := sandbox.NewDummy()
			if err := cpp.AutoCompile(context.TODO(), fs, s, path, filepath.Join(path, p.Tests.Checker.Path), filepath.Join(path, binaryName)); err != nil {
				return nil, err
			}

			p.Tests.Checker.Path = binaryName
		}

		return p, nil
	}

	identifier := func(fs afero.Fs, path string) bool {
		_, err := fs.Stat(filepath.Join(path, "problem.yaml"))
		return !os.IsNotExist(err)
	}

	return parser, identifier
}

func init() {
	parser, identifier := ParserAndIdentifier()
	problems.RegisterConfigType("problem_yaml", parser, identifier)
}
