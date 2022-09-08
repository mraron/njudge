package task_yaml

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"github.com/mraron/njudge/pkg/language/langs/cpp"
	"io"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/mraron/njudge/pkg/language"
	"github.com/mraron/njudge/pkg/problems"
	"github.com/mraron/njudge/pkg/problems/checker"
	"github.com/mraron/njudge/pkg/problems/tasktype/batch"
	"github.com/mraron/njudge/pkg/problems/tasktype/communication"
	"go.uber.org/multierr"
	"gopkg.in/yaml.v2"
)

type TaskYAML struct {
	Name                string `yaml:"name"`
	Title               string
	TimeLimit           float64 `yaml:"time_limit"`
	MemoryLimit         int     `yaml:"memory_limit"`
	InputCount          int     `yaml:"n_input"`
	Infile              string
	Outfile             string
	PrimaryLanguage     string   `yaml:"primary_language"`
	TokenMode           string   `yaml:"token_mode"`
	MaxSubmissionCount  int      `yaml:"max_submission_number"`
	PublicTestcases     string   `yaml:"public_testcases"`
	FeedbackLevel       string   `yaml:"feedback_level"`
	ScoreType           string   `yaml:"score_type"`
	ScoreTypeParameters [][2]int `yaml:"score_type_parameters"`
	ScorePrecision      int      `yaml:"score_precision"`
	ScoreMode           string   `yaml:"score_mode"`
	TotalValue          int      `yaml:"total_value"`
	OutputOnly          bool     `yaml:"output_only"`
	TaskTypeParameters  []string `yaml:"task_type_parameters"`
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
	return p.StatementList.FilterByType("text/html")
}

func (p Problem) PDFStatements() problems.Contents {
	return p.StatementList.FilterByType("application/pdf")
}

func (p Problem) MemoryLimit() int {
	return 1024 * 1024 * p.TaskYAML.MemoryLimit
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
		return []language.Language{language.Get("zip")}
	}

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
	return make([]string, 0)
}

func (p Problem) StatusSkeleton(name string) (*problems.Status, error) {
	ans := problems.Status{Compiled: false, CompilerOutput: "status skeleton", FeedbackType: problems.FeedbackIOI, Feedback: make([]problems.Testset, 0)}
	ans.Feedback = append(ans.Feedback, problems.Testset{Name: "tests"})
	testset := &ans.Feedback[len(ans.Feedback)-1]

	tcByGroup := make(map[string][]problems.Testcase)
	subtask := 0
	testsLeft := p.ScoreTypeParameters[subtask][1]
	subtasks := make([]string, len(p.ScoreTypeParameters))
	for ind := 0; ind < p.InputCount; ind++ {
		tc := problems.Testcase{}
		tc.InputPath, tc.AnswerPath = fmt.Sprintf(p.InputPathPattern, ind), fmt.Sprintf(p.AnswerPathPattern, ind)
		if p.OutputOnly {
			// default cms loader behaviour
			tc.OutputPath = fmt.Sprintf("output_%03d.txt", ind)
		}

		tc.Index = ind + 1
		tc.MaxScore = 0
		tc.VerdictName = problems.VerdictDR
		tc.MemoryLimit = p.MemoryLimit()
		tc.TimeLimit = time.Duration(p.TimeLimit()) * time.Millisecond

		if testsLeft == 0 {
			subtask++
			if subtask < len(p.ScoreTypeParameters) {
				testsLeft = p.ScoreTypeParameters[subtask][1]
			}
		}

		subtasks[subtask] = "subtask" + strconv.Itoa(subtask+1)
		tc.Group = "subtask" + strconv.Itoa(subtask+1)

		if len(tcByGroup[tc.Group]) == 0 {
			tcByGroup[tc.Group] = make([]problems.Testcase, 0)
			tc.MaxScore = float64(p.ScoreTypeParameters[subtask][0])
		}

		testsLeft--
		tcByGroup[tc.Group] = append(tcByGroup[tc.Group], tc)
	}

	for _, subtask := range subtasks {
		testset.Groups = append(testset.Groups, problems.Group{})
		group := &testset.Groups[len(testset.Groups)-1]

		group.Name = subtask
		group.Scoring = problems.ScoringGroup
		for _, tc := range tcByGroup[subtask] {
			group.Testcases = append(group.Testcases, tc)
		}
	}

	return &ans, nil
}

func (p Problem) Checker() problems.Checker {
	if p.tasktype == "communication" { // manager already printed the result
		return checker.Noop{}
	}

	if p.whiteDiffChecker {
		return checker.Whitediff{}
	}

	return checker.NewTaskYAML(filepath.Join(p.Path, "check", "checker"))
}

func (p Problem) Files() []problems.File {
	return p.files
}

func (p Problem) GetTaskType() problems.TaskType {
	var (
		tt  problems.TaskType
		err error
	)

	if p.tasktype == "outputonly" {
		tt, err = problems.GetTaskType("outputonly")
	} else if p.tasktype == "batch" {
		tt, err = problems.GetTaskType("batch")
	} else if p.tasktype == "communication" {
		res := communication.New()
		res.RunInteractorF = func(rc *batch.RunContext, utoi, itou *os.File, g *problems.Group, tc *problems.Testcase) (language.Status, error) {
			input, err := os.Open(tc.InputPath)
			if err != nil {
				return language.Status{}, multierr.Combine(err, input.Close())
			}
			defer input.Close()

			sbox := rc.Store["interactorSandbox"].(language.Sandbox).Stdin(input).Stdout(rc.Stdout).TimeLimit(2 * tc.TimeLimit).MemoryLimit(1024 * 1024)
			sbox.MapDir("/fifo", filepath.Dir(itou.Name()), []string{"rw"}, false)

			st, err := sbox.Run(fmt.Sprintf("interactor %s %s", filepath.Join("/fifo", filepath.Base(utoi.Name())), filepath.Join("/fifo", filepath.Base(itou.Name()))), true)
			if err != nil {
				return st, err
			}

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
			return st, multierr.Combine(err, rc.Store["interactorSandbox"].(language.Sandbox).CreateFile("out", bytes.NewBuffer([]byte{})))
		}

		res.RunUserF = func(rc *batch.RunContext, utoi, itou *os.File, g *problems.Group, tc *problems.Testcase) (language.Status, error) {
			time.Sleep(25 * time.Millisecond) // TODO very hacky
			return rc.Lang.Run(rc.Sandbox, bytes.NewReader(rc.Binary), itou, utoi, tc.TimeLimit, tc.MemoryLimit)
		}

		tt = res
	}

	if err != nil {
		panic(err)
	}

	return tt
}

func parseGen(r io.Reader) (int, [][2]int, error) {
	var err error
	subtasks, testcases, points := make([][2]int, 0), 0, -1
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
					subtasks = append(subtasks, [2]int{points, testcases})
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

	subtasks = append(subtasks, [2]int{points, testcases})
	return inputCount, subtasks, nil
}

func parser(path string) (problems.Problem, error) {
	p := Problem{
		Path:              path,
		InputPathPattern:  filepath.Join(path, "input", "input%d.txt"),
		AnswerPathPattern: filepath.Join(path, "output", "output%d.txt"),
		AttachmentList:    make(problems.Attachments, 0),
		files:             make([]problems.File, 0),
	}

	YAMLFile, err := os.Open(filepath.Join(path, "task.yaml"))
	if err != nil {
		return nil, err
	}

	dec := yaml.NewDecoder(YAMLFile)
	if err = dec.Decode(&p.TaskYAML); err != nil {
		return nil, err
	}

	statementPDF, err := ioutil.ReadFile(filepath.Join(path, "statement", "statement.pdf"))
	if err != nil {
		return nil, err
	}

	genPath := filepath.Join(p.Path, "gen", "GEN")
	if _, err = os.Stat(genPath); err == nil {
		gen, err := os.Open(genPath)
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
	p.StatementList = append(p.StatementList, problems.BytesData{Loc: "hungarian", Val: statementPDF, Typ: "application/pdf"})

	exists := func(path string) bool {
		if _, err := os.Stat(path); errors.Is(err, fs.ErrNotExist) {
			return false
		} else if err == nil {
			return true
		} else { // could be both
			return false
		}
	}

	compile := func(src string, to string) error {
		if bin, err := os.Create(to); err == nil {
			defer bin.Close()

			if file, err := os.Open(src); err == nil {
				defer file.Close()

				if err := cpp.Std14.InsecureCompile(filepath.Dir(src), file, bin, os.Stderr); err != nil {
					return err
				}

				if err := os.Chmod(to, os.ModePerm); err != nil {
					return err
				}
			}

			return nil
		} else {
			return err
		}
	}

	chmodX := func(path string) error {
		if stat, _ := os.Stat(path); stat.Mode().Perm()&fs.ModePerm != fs.ModePerm {
			return os.Chmod(path, os.ModePerm)
		}

		return nil
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
			if !exists(checkerPath) {
				if err := compile(checkerCppPath, checkerPath); err != nil {
					return nil, err
				}
			}

			chmodX(checkerPath)
		} else if exists(managerCppPath) {
			p.tasktype = "communication"
			if err := compile(managerCppPath, managerPath); err != nil {
				return nil, err
			}
			p.files = append(p.files, problems.File{Name: "manager.cpp", Role: "interactor", Path: managerPath})

			chmodX(managerPath)
		}

		if _, err := os.Stat(filepath.Join(solPath, "grader.cpp")); err == nil {
			p.tasktype = "batch"
			p.files = append(p.files, problems.File{Name: "grader.cpp", Role: "stub_cpp", Path: filepath.Join(solPath, "grader.cpp")})
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
		files, err = ioutil.ReadDir(solPath)
		if err != nil {
			return nil, err
		}

		for _, file := range files {
			if !file.IsDir() && filepath.Ext(file.Name()) == ".h" {
				p.files = append(p.files, problems.File{Name: filepath.Base(file.Name()), Role: "stub_cpp", Path: file.Name()})
			}
		}
	}

	if p.tasktype == "" {
		p.tasktype = "batch"
	}

	if p.OutputOnly {
		p.tasktype = "outputonly"
	}

	attPath := filepath.Join(path, "att")
	if exists(attPath) {
		files, err := ioutil.ReadDir(filepath.Join(path, "att"))
		if err != nil {
			return nil, err
		}

		for _, file := range files {
			if !file.IsDir() {
				cont, err := ioutil.ReadFile(filepath.Join(path, "att", file.Name()))
				if err != nil {
					return nil, err
				}

				p.AttachmentList = append(p.AttachmentList, problems.BytesData{Nam: file.Name(), Val: cont})
			}
		}
	}

	return p, nil
}

func identifier(path string) bool {
	_, err := os.Stat(filepath.Join(path, "task.yaml"))
	return !os.IsNotExist(err)
}

func init() {
	problems.RegisterConfigType("task_yaml", parser, identifier)
}
