package problem_yaml

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"github.com/mraron/njudge/pkg/language"
	"github.com/mraron/njudge/pkg/language/langs/cpp"
	"github.com/mraron/njudge/pkg/language/langs/zip"
	"github.com/mraron/njudge/pkg/language/memory"
	"github.com/mraron/njudge/pkg/language/sandbox"
	"github.com/mraron/njudge/pkg/problems"
	"github.com/mraron/njudge/pkg/problems/evaluation"
	"github.com/mraron/njudge/pkg/problems/evaluation/batch"
	"github.com/mraron/njudge/pkg/problems/evaluation/communication"
	"github.com/mraron/njudge/pkg/problems/evaluation/output_only"
	"github.com/mraron/njudge/pkg/problems/executable/checker"
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
	TestCount          int       `yaml:"test_count"`
	InputList          []string  `yaml:"input_list"`
	OutputList         []string  `yaml:"output_list"`
	InputPattern       string    `yaml:"input_pattern"`
	OutputPattern      string    `yaml:"output_pattern"`
	ZeroIndexedPattern bool      `yaml:"zero_indexed_pattern"`
	Scoring            string    `yaml:"scoring"`
	MaxScore           float64   `yaml:"max_score"`
	Scores             []float64 `yaml:"scores"`
}

type Checker struct {
	Type string `yaml:"type"`
	Path string `yaml:"path"`
}

type Stub struct {
	Lang string `yaml:"lang"`
	Path string `yaml:"path"`
}

type Tests struct {
	TaskType           string   `yaml:"task_type"`
	TaskTypeArgs       []string `yaml:"task_type_args"`
	interactorBinary   []byte
	Stubs              []Stub    `yaml:"stubs"`
	InputFile          string    `yaml:"input_file"`
	OutputFile         string    `yaml:"output_file"`
	MemoryLimit        int       `yaml:"memory_limit"`
	TimeLimit          float64   `yaml:"time_limit"`
	TestCount          int       `yaml:"test_count"`
	InputPattern       string    `yaml:"input_pattern"`
	OutputPattern      string    `yaml:"output_pattern"`
	InputList          []string  `yaml:"input_list"`
	OutputList         []string  `yaml:"output_list"`
	ZeroIndexedPattern bool      `yaml:"zero_indexed_pattern"`
	Checker            Checker   `yaml:"checker"`
	FeedbackType       string    `yaml:"feedback_type"`
	Scores             []float64 `yaml:"scores"`
	Subtasks           []Subtask `yaml:"subtasks"`
}

type Problem struct {
	Path                   string               `yaml:"-"`
	GeneratedStatementList problems.Contents    `yaml:"-"`
	AttachmentList         problems.Attachments `yaml:"-"`

	ShortName      string       `yaml:"shortname"`
	TitleList      []Title      `yaml:"titles"`
	StatementList  []Statement  `yaml:"statements"`
	AttachmentInfo []Attachment `yaml:"attachments"`

	Tests Tests `yaml:"tests"`
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

func (p Problem) MemoryLimit() memory.Amount {
	return memory.Amount(p.Tests.MemoryLimit * 1024 * 1024)
}

func (p Problem) TimeLimit() int {
	return int(1000 * p.Tests.TimeLimit)
}

func (p Problem) InputOutputFiles() (string, string) {
	return p.Tests.InputFile, p.Tests.OutputFile
}

func (p Problem) Languages() []language.Language {
	if p.Tests.TaskType == "outputonly" {
		return []language.Language{zip.Zip{}}
	}
	if len(p.Tests.Stubs) > 0 {
		var res []language.Language
		for _, lang := range language.ListExcept(language.DefaultStore, []string{"zip"}) {
			for _, stub := range p.Tests.Stubs {
				file := problems.EvaluationFile{
					Name: filepath.Base(stub.Path),
					Role: "stub_" + stub.Lang,
					Path: filepath.Join(p.Path, stub.Path),
				}
				if file.StubOf(lang) {
					res = append(res, lang)
					break
				}
			}
		}
		return res
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
	ans := problems.Status{Compiled: false, CompilerOutput: "status skeleton", FeedbackType: problems.FeedbackTypeFromShortString(p.Tests.FeedbackType), Feedback: make([]problems.Testset, 0)}
	ans.Feedback = append(ans.Feedback, problems.Testset{Name: "tests"})

	getIthIO := func(typ string, index int, pattern string, gindex int, gpattern string, list []string, tests *Tests, subtask *Subtask) (string, error) {
		if pattern != "" {
			ind := index
			if subtask == nil || !subtask.ZeroIndexedPattern {
				ind += 1
			}
			return fmt.Sprintf(filepath.Join(p.Path, "tests", pattern), ind), nil
		} else if gpattern != "" {
			ind := gindex
			if !tests.ZeroIndexedPattern {
				ind += 1
			}
			return fmt.Sprintf(filepath.Join(p.Path, "tests", gpattern), ind), nil
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

			if tc.InputPath, err = getIthIO("input", i, p.Tests.InputPattern, -1, "", p.Tests.InputList, nil, nil); err != nil {
				return nil, err
			}
			if tc.AnswerPath, err = getIthIO("output", i, p.Tests.OutputPattern, -1, "", p.Tests.OutputList, nil, nil); err != nil {
				return nil, err
			}
			if p.Tests.TaskType == "outputonly" {
				tc.OutputPath = filepath.Base(tc.AnswerPath)
			}
			if len(p.Tests.Scores) > 0 {
				if i >= len(p.Tests.Scores) {
					return nil, errors.New("too few scores")
				}
				tc.MaxScore = p.Tests.Scores[i]
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

				if tc.InputPath, err = getIthIO(
					"input",
					i,
					p.Tests.Subtasks[s].InputPattern,
					globalIdx,
					p.Tests.InputPattern,
					p.Tests.Subtasks[s].InputList,
					&p.Tests,
					&p.Tests.Subtasks[s],
				); err != nil {
					return nil, err
				}
				if tc.AnswerPath, err = getIthIO(
					"output",
					i,
					p.Tests.Subtasks[s].OutputPattern,
					globalIdx,
					p.Tests.OutputPattern,
					p.Tests.Subtasks[s].OutputList,
					&p.Tests,
					&p.Tests.Subtasks[s],
				); err != nil {
					return nil, err
				}
				if p.Tests.TaskType == "outputonly" {
					tc.OutputPath = filepath.Base(tc.AnswerPath)
				}

				if len(p.Tests.Subtasks[s].Scores) > 0 {
					if i >= len(p.Tests.Subtasks[s].Scores) {
						return nil, errors.New("too few scores")
					}
					tc.MaxScore = p.Tests.Subtasks[s].Scores[i]
				} else if ans.Feedback[0].Groups[s].Scoring == problems.ScoringSum {
					tc.MaxScore = p.Tests.Subtasks[s].MaxScore / float64(p.Tests.Subtasks[s].TestCount)
				} else if i == 0 && ans.Feedback[0].Groups[s].Scoring == problems.ScoringGroup {
					tc.MaxScore = p.Tests.Subtasks[s].MaxScore
				} else if ans.Feedback[0].Groups[s].Scoring == problems.ScoringMin {
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
		return checker.NewWhitediff()
	} else if p.Tests.Checker.Type == "testlib" {
		return checker.NewTestlib(filepath.Join(p.Path, p.Tests.Checker.Path))
	} else if p.Tests.Checker.Type == "taskyaml" {
		return checker.NewTaskYAML(filepath.Join(p.Path, p.Tests.Checker.Path))
	}

	return checker.Noop{}
}

func (p Problem) EvaluationFiles() []problems.EvaluationFile {
	return make([]problems.EvaluationFile, 0)
}

func (p Problem) GetTaskType() problems.TaskType {
	if p.Tests.TaskType == "outputonly" {
		return output_only.New(p.Checker())
	}

	var compiler problems.Compiler = evaluation.Compile{}
	if len(p.Tests.Stubs) > 0 {
		stubCompiler := evaluation.NewCompilerWithStubs()
		for _, stub := range p.Tests.Stubs {
			file := problems.EvaluationFile{Name: filepath.Base(stub.Path), Role: "stub_" + stub.Lang, Path: filepath.Join(p.Path, stub.Path)}
			for _, lang := range p.Languages() {
				if file.StubOf(lang) {
					stubCompiler.AddStub(lang, file)
				}
			}
		}
		compiler = stubCompiler
	}

	if p.Tests.TaskType == "communication" {
		switch p.Tests.TaskTypeArgs[0] { // interactor type
		case "task_yaml":
			eval := &evaluation.TaskYAMLUserInteractorExecute{}
			return communication.New(compiler, p.Tests.interactorBinary, eval, evaluation.InteractiveRunnerWithExecutor(eval))
		case "polygon":
			return communication.New(evaluation.CompileCheckSupported{
				List:         p.Languages(),
				NextCompiler: evaluation.Compile{},
			}, p.Tests.interactorBinary, p.Checker())
		}
	}

	opts := []evaluation.BasicRunnerOption{evaluation.BasicRunnerWithChecker(p.Checker())}
	if i, o := p.InputOutputFiles(); i != "" || o != "" {
		opts = append(opts, evaluation.BasicRunnerWithFiles(i, o))
	}
	return batch.New(evaluation.CompileCheckSupported{
		List:         p.Languages(),
		NextCompiler: compiler,
	}, opts...)
}

type config struct {
	compileBinaries bool //TODO respect this
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

		if p.Tests.TaskType == "communication" {
			if len(p.Tests.TaskTypeArgs) < 2 {
				return nil, errors.New("too few task_type_arguments for communication task")
			}
			interactorPath := &p.Tests.TaskTypeArgs[1]
			if strings.HasSuffix(*interactorPath, ".cpp") {
				binaryName := strings.TrimSuffix(*interactorPath, filepath.Ext(*interactorPath))
				s, _ := sandbox.NewDummy()
				if err := cpp.AutoCompile(context.TODO(), fs, s, path, filepath.Join(path, *interactorPath), filepath.Join(path, binaryName)); err != nil {
					return nil, err
				}

				*interactorPath = binaryName
			}
			p.Tests.interactorBinary, err = os.ReadFile(filepath.Join(p.Path, *interactorPath))
			if err != nil {
				return nil, err
			}
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
	_ = problems.RegisterConfigType("problem_yaml", parser, identifier)
}
