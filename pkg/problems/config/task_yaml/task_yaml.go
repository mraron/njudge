package task_yaml

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"github.com/mraron/njudge/pkg/language/memory"
	"github.com/mraron/njudge/pkg/problems/evaluation"
	checker2 "github.com/mraron/njudge/pkg/problems/executable/checker"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/mraron/njudge/pkg/language/langs/cpp"
	"github.com/mraron/njudge/pkg/language/sandbox"
	"github.com/spf13/afero"

	"github.com/mraron/njudge/pkg/language"
	"github.com/mraron/njudge/pkg/problems"
	"gopkg.in/yaml.v3"
)

type TaskYAML struct {
	Name                string `yaml:"name"`
	Title               string
	TimeLimit           float64 `yaml:"time_limit"`
	MemoryLimit         int     `yaml:"memory_limit"`
	InputCount          int     `yaml:"n_input"`
	Infile              string
	Outfile             string
	PrimaryLanguage     string           `yaml:"primary_language"`
	TokenMode           string           `yaml:"token_mode"`
	MaxSubmissionCount  int              `yaml:"max_submission_number"`
	PublicTestcases     string           `yaml:"public_testcases"`
	FeedbackLevel       string           `yaml:"feedback_level"`
	ScoreType           string           `yaml:"score_type"`
	ScoreTypeParameters [][2]interface{} `yaml:"score_type_parameters"`
	ScorePrecision      int              `yaml:"score_precision"`
	ScoreMode           string           `yaml:"score_mode"`
	TotalValue          int              `yaml:"total_value"`
	OutputOnly          bool             `yaml:"output_only"`
	TaskTypeParameters  []string         `yaml:"task_type_parameters"`
}

type Problem struct {
	TaskYAML

	StatementList  problems.Contents
	AttachmentList problems.Attachments

	InputPathPattern  string
	AnswerPathPattern string

	Path string

	files            []problems.File
	tasktype         string
	whiteDiffChecker bool
}

func (p Problem) Name() string {
	return p.TaskYAML.Name
}

func (p Problem) Titles() problems.Contents {
	return problems.Contents{problems.BytesData{Loc: "hungarian", Val: []byte(p.Title), Typ: "text"}}
}

func (p Problem) Statements() problems.Contents {
	return p.StatementList
}

func (p Problem) HTMLStatements() problems.Contents {
	return p.StatementList.FilterByType(problems.DataTypeHTML)
}

func (p Problem) PDFStatements() problems.Contents {
	return p.StatementList.FilterByType(problems.DataTypePDF)
}

func (p Problem) MemoryLimit() memory.Amount {
	return memory.Amount(1024 * 1024 * p.TaskYAML.MemoryLimit)
}

func (p Problem) TimeLimit() int {
	return int(p.TaskYAML.TimeLimit * float64(1000))
}

func (p Problem) InputOutputFiles() (string, string) {
	return "", ""
}

func (p Problem) Interactive() bool {
	return false
}

func (p Problem) Languages() []language.Language {
	if p.OutputOnly {
		return []language.Language{language.DefaultStore.Get("zip")}
	}

	lst1 := language.DefaultStore.List()

	lst2 := make([]language.Language, 0, len(lst1))
	for _, val := range lst1 {
		if p.tasktype == "stub" {
			hasStub := false
			for _, f := range p.files {
				if f.Role == "stub_"+val.ID() || (f.Role == "stub_cpp" && strings.HasPrefix(val.ID(), "cpp")) {
					hasStub = true
					break
				}
			}

			if hasStub {
				lst2 = append(lst2, val)
			}
		} else {
			if val.ID() != "zip" {
				lst2 = append(lst2, val)
			}
		}
	}

	return lst2
}

func (p Problem) Attachments() problems.Attachments {
	return p.AttachmentList
}

func (p Problem) Tags() []string {
	return make([]string, 0)
}

func (p Problem) StatusSkeleton(name string) (*problems.Status, error) {
	ans := problems.Status{Compiled: false, CompilerOutput: "status skeleton", FeedbackType: problems.FeedbackIOI, Feedback: make([]problems.Testset, 0)}
	ans.Feedback = append(ans.Feedback, problems.Testset{Name: "tests"})
	testset := &ans.Feedback[len(ans.Feedback)-1]

	tcByGroup := make(map[string][]problems.Testcase)
	ind, subtask := 0, 0

	testsLeft := make([][2]string, 0)
	testIndices := make([]int, 0)
	advanceTests := func() {
		if p.ScoreTypeParameters == nil {
			for i := 0; i < p.InputCount; i++ {
				testsLeft = append(testsLeft, [2]string{fmt.Sprintf(p.InputPathPattern, ind), fmt.Sprintf(p.AnswerPathPattern, ind)})
				testIndices = append(testIndices, ind)
				ind++
			}
		} else if val, ok := p.ScoreTypeParameters[subtask][1].(int); ok {
			for i := 0; i < val; i++ {
				testsLeft = append(testsLeft, [2]string{fmt.Sprintf(p.InputPathPattern, ind), fmt.Sprintf(p.AnswerPathPattern, ind)})
				testIndices = append(testIndices, ind)
				ind++
			}
		} else if val, ok := p.ScoreTypeParameters[subtask][1].(string); ok {
			indices := strings.Split(val, "|")
			for i := range indices {
				num, _ := strconv.Atoi(indices[i])
				testsLeft = append(testsLeft, [2]string{fmt.Sprintf(p.InputPathPattern, num), fmt.Sprintf(p.AnswerPathPattern, num)})
				testIndices = append(testIndices, num)
			}
		} else {
			panic("wrong format")
		}
	}
	advanceTests()

	subtaskCount := len(p.ScoreTypeParameters)
	isSum := false
	if subtaskCount == 0 {
		subtaskCount = 1
		isSum = true
	}

	subtasks := make([]string, subtaskCount)
	idx := 0
	for len(testsLeft) > 0 {
		tc := problems.Testcase{}
		tc.InputPath, tc.AnswerPath = testsLeft[0][0], testsLeft[0][1]
		if p.OutputOnly {
			// default cms loader behaviour
			tc.OutputPath = fmt.Sprintf("output_%03d.txt", testIndices[0])
		}
		tc.Index = idx + 1
		idx += 1
		tc.MaxScore = 0
		tc.VerdictName = problems.VerdictDR
		tc.MemoryLimit = p.MemoryLimit()
		tc.TimeLimit = time.Duration(p.TimeLimit()) * time.Millisecond

		subtasks[subtask] = "subtask" + strconv.Itoa(subtask+1)
		tc.Group = "subtask" + strconv.Itoa(subtask+1)

		if isSum {
			tc.MaxScore = 100.0 / float64(p.InputCount)
		}

		if len(testsLeft) == 1 {
			if !isSum {
				tc.MaxScore = float64(p.ScoreTypeParameters[subtask][0].(int))
			}
			subtask++
			testsLeft = testsLeft[1:]
			testIndices = testIndices[1:]

			if subtask < len(p.ScoreTypeParameters) {
				advanceTests()
			}
		} else {
			testsLeft = testsLeft[1:]
			testIndices = testIndices[1:]
		}

		if len(tcByGroup[tc.Group]) == 0 {
			tcByGroup[tc.Group] = make([]problems.Testcase, 0)
		}

		tcByGroup[tc.Group] = append(tcByGroup[tc.Group], tc)
	}

	for _, subtask := range subtasks {
		testset.Groups = append(testset.Groups, problems.Group{})
		group := &testset.Groups[len(testset.Groups)-1]

		group.Name = subtask
		if isSum {
			group.Scoring = problems.ScoringSum
		} else {
			group.Scoring = problems.ScoringMin
			for ind := range tcByGroup[subtask] {
				tcByGroup[subtask][ind].MaxScore = tcByGroup[subtask][len(tcByGroup[subtask])-1].MaxScore
			}
		}

		group.Testcases = append(group.Testcases, tcByGroup[subtask]...)
	}

	return &ans, nil
}

func (p Problem) Checker() problems.Checker {
	if p.tasktype == "communication" { // manager already printed the result
		return checker2.Noop{}
	}

	if p.whiteDiffChecker {
		return checker2.Whitediff{}
	}

	return checker2.NewTaskYAML(filepath.Join(p.Path, "check", "checker"))
}

func (p Problem) Files() []problems.File {
	return p.files
}

func (p Problem) GetTaskType() problems.TaskType {
	return problems.NewTaskType(
		"batch",
		evaluation.CompileCheckSupported{},
		evaluation.NewLinearEvaluator(
			evaluation.NewBasicRunner(evaluation.BasicRunnerWithChecker(p.Checker())),
		),
	)
	// TODO outputonly, stub, communication
	/*
		var (
			tt  problems.TaskType
			err error
		)

		if p.tasktype == "outputonly" {
			tt, err = problems.GetTaskType("outputonly")
		} else if p.tasktype == "batch" {
			tt, err = problems.GetTaskType("batch")
		} else if p.tasktype == "stub" {
			tt, err = problems.GetTaskType("stub")
		} else if p.tasktype == "communication" {
			res := communication.New()
			res.RunInteractorF = func(rc *batch.RunContext, utoi, itou *os.File, g *problems.Group, tc *problems.Testcase) (language.Status, error) {
				input, err := os.Open(tc.InputPath)
				if err != nil {
					return language.Status{}, errors.Join(err, input.Close())
				}
				defer input.Close()

				sbox := rc.Store["interactorSandbox"].(language.Sandbox).Stdin(input).Stdout(rc.Stdout).TimeLimit(2 * tc.TimeLimit).MemoryLimit(1024 * 1024)
				sbox.MapDir("/fifo", filepath.Dir(itou.Name()), []string{"rw"}, false)

				st, err := sbox.Run(fmt.Sprintf("interactor %s %s", filepath.Join("/fifo", filepath.Base(utoi.Name())), filepath.Join("/fifo", filepath.Base(itou.Name()))), true)
				if err != nil {
					return st, err
				}
				itou.Close()

				fmt.Fscanf(rc.Stdout, "%f", &tc.Score)
				if tc.Score == 0 {
					tc.VerdictName = problems.VerdictWA
				} else if tc.Score < 1.0 {
					tc.VerdictName = problems.VerdictPC
				} else {
					tc.VerdictName = problems.VerdictAC
				}

				tc.Score *= tc.MaxScore

				// For compatibility create a file named out
				return st, errors.Join(err, rc.Store["interactorSandbox"].(language.Sandbox).CreateFile("out", bytes.NewBuffer([]byte{})))
			}

			res.RunUserF = func(rc *batch.RunContext, utoi, itou *os.File, g *problems.Group, tc *problems.Testcase) (language.Status, error) {
				res, err := rc.Lang.Run(rc.Sandbox, bytes.NewReader(rc.Binary), itou, utoi, tc.TimeLimit, tc.MemoryLimit)
				utoi.Close()
				return res, err
			}

			tt = res
		}

		if err != nil {
			panic(err)
		}

		return tt*/
}

func parseGen(r io.Reader) (int, [][2]interface{}, error) {
	var err error
	subtasks, testcases, points := make([][2]interface{}, 0), 0, -1
	inputCount := 0

	sc := bufio.NewScanner(r)
	for sc.Scan() {
		line := strings.TrimSpace(sc.Text())
		splitted := strings.SplitN(line, "#", 2)

		if len(splitted) == 1 {
			if splitted[0] != "" {
				testcases++
				inputCount++
			}
		} else {
			testcase, comment := splitted[0], splitted[1]
			testcase, comment = strings.TrimSpace(testcase), strings.TrimSpace(comment)

			testcaseDetected := len(testcase) > 0
			copyTestcaseDetected := strings.HasPrefix(comment, "COPY:")
			subtaskDetected := strings.HasPrefix(comment, "ST:")

			cnt := 0
			if testcaseDetected {
				cnt++
			}
			if copyTestcaseDetected {
				cnt++
			}
			if subtaskDetected {
				cnt++
			}

			if cnt > 1 {
				return -1, nil, errors.New("multiple commands on one line of GEN file")
			}

			if testcaseDetected || copyTestcaseDetected {
				testcases++
				inputCount++
			}

			if subtaskDetected {
				if points == -1 {
					if testcases > 0 {
						return -1, nil, errors.New("trailing testcases")
					}
				} else {
					subtasks = append(subtasks, [2]interface{}{points, testcases})
				}

				testcases = 0
				if points, err = strconv.Atoi(strings.TrimSpace(comment[3:])); err != nil {
					return -1, nil, err
				}
			}
		}
	}

	if err = sc.Err(); err != nil {
		return -1, nil, err
	}

	subtasks = append(subtasks, [2]interface{}{points, testcases})
	return inputCount, subtasks, nil
}

func primaryLanguageToLocale(primaryLanguage string) string {
	switch primaryLanguage {
	case "hu":
		return "hungarian"
	case "en":
		return "english"
	}
	return "hungarian"
}

func Parser(fs afero.Fs, path string) (problems.Problem, error) {
	p := Problem{
		Path:              path,
		InputPathPattern:  filepath.Join(path, "input", "input%d.txt"),
		AnswerPathPattern: filepath.Join(path, "output", "output%d.txt"),
		AttachmentList:    make(problems.Attachments, 0),
		files:             make([]problems.File, 0),
	}

	YAMLFile, err := fs.Open(filepath.Join(path, "task.yaml"))
	if err != nil {
		return nil, err
	}

	dec := yaml.NewDecoder(YAMLFile)
	if err = dec.Decode(&p.TaskYAML); err != nil {
		return nil, err
	}

	statementPDF, err := afero.ReadFile(fs, filepath.Join(path, "statement", "statement.pdf"))
	if err != nil {
		return nil, err
	}

	genPath := filepath.Join(p.Path, "gen", "GEN")
	if _, err = fs.Stat(genPath); err == nil && len(p.ScoreTypeParameters) == 0 {
		gen, err := fs.Open(genPath)
		if err != nil {
			return nil, err
		}

		inputCount, subtasks, err := parseGen(gen)
		if err != nil {
			return nil, err
		}

		p.InputCount = inputCount
		p.ScoreTypeParameters = subtasks
	}

	p.StatementList = make(problems.Contents, 0)
	p.StatementList = append(p.StatementList, problems.BytesData{Loc: primaryLanguageToLocale(p.PrimaryLanguage), Val: statementPDF, Typ: "application/pdf"})

	exists := func(path string) bool {
		if _, err := fs.Stat(path); errors.Is(err, os.ErrNotExist) {
			return false
		} else if err == nil {
			return true
		} else { // could be both
			return false
		}
	}

	checkPath := filepath.Join(p.Path, "check")
	solPath := filepath.Join(p.Path, "sol")

	if !exists(checkPath) {
		p.whiteDiffChecker = true
	} else {
		checkerPath := filepath.Join(checkPath, "checker")
		checkerCppPath := filepath.Join(checkPath, "checker.cpp")

		managerCppPath := filepath.Join(checkPath, "manager.cpp")
		managerPath := filepath.Join(checkPath, "manager")

		if exists(checkerCppPath) {
			s, _ := sandbox.NewDummy()
			if err := cpp.AutoCompile(context.TODO(), fs, s, checkPath, checkerCppPath, checkerPath); err != nil {
				return nil, err
			}
		} else if exists(managerCppPath) {
			p.tasktype = "communication"
			s, _ := sandbox.NewDummy()
			if err := cpp.AutoCompile(context.TODO(), fs, s, checkPath, managerCppPath, managerPath); err != nil {
				return nil, err
			}

			p.files = append(p.files, problems.File{Name: "manager.cpp", Role: "interactor", Path: managerPath})
		} else {
			p.whiteDiffChecker = true
		}
	}

	if exists(solPath) {
		if exists(filepath.Join(solPath, "grader.cpp")) {
			p.tasktype = "batch"
			p.files = append(p.files, problems.File{Name: "grader.cpp", Role: "stub_cpp", Path: filepath.Join(solPath, "grader.cpp")})
		} else if exists(filepath.Join(solPath, "stub.cpp")) {
			p.tasktype = "communication"
			p.files = append(p.files, problems.File{Name: "stub.cpp", Role: "stub_cpp", Path: filepath.Join(solPath, "stub.cpp")})
		}

		var files []os.FileInfo
		files, err = afero.ReadDir(fs, solPath)
		if err != nil {
			return nil, err
		}

		for _, file := range files {
			if !file.IsDir() && filepath.Ext(file.Name()) == ".h" && file.Name() != "testlib.h" {
				p.files = append(p.files, problems.File{Name: file.Name(), Role: "stub_cpp", Path: filepath.Join(solPath, file.Name())})
			}
		}
	}

	hasStub := false
	for _, f := range p.files {
		if strings.Contains(f.Role, "stub") {
			hasStub = true
		}
	}

	if hasStub && p.tasktype != "communication" {
		p.tasktype = "stub"
	}

	if p.tasktype == "" {
		p.tasktype = "batch"
	}

	if p.OutputOnly {
		p.tasktype = "outputonly"
	}

	attPath := filepath.Join(path, "att")
	if exists(attPath) {
		files, err := afero.ReadDir(fs, filepath.Join(path, "att"))
		if err != nil {
			return nil, err
		}

		for _, file := range files {
			if !file.IsDir() {
				cont, err := afero.ReadFile(fs, filepath.Join(path, "att", file.Name()))
				if err != nil {
					return nil, err
				}

				p.AttachmentList = append(p.AttachmentList, problems.BytesData{Nam: file.Name(), Val: cont})
			}
		}
	}

	return p, nil
}

func Identifier(fs afero.Fs, path string) bool {
	_, err := fs.Stat(filepath.Join(path, "task.yaml"))
	return !os.IsNotExist(err)
}

func init() {
	problems.RegisterConfigType("task_yaml", Parser, Identifier)
}
