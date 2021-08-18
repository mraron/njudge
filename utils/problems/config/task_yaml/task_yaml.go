package task_yaml

import (
	"bytes"
	"fmt"
	"github.com/friendsofgo/errors"
	"github.com/mraron/njudge/utils/language"
	"github.com/mraron/njudge/utils/problems"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"time"
)

type TaskYAML struct {
	Name string `yaml:"name"`
	Title string
	TimeLimit float64 `yaml:"time_limit"`
	MemoryLimit int `yaml:"memory_limit"`
	InputCount int `yaml:"n_input"`
	Infile string
	Outfile string
	PrimaryLanguage string `yaml:"primary_language"`
	TokenMode string `yaml:"token_mode"`
	MaxSubmissionCount int `yaml:"max_submission_number"`
	PublicTestcases string `yaml:"public_testcases"`
	FeedbackLevel string `yaml:"feedback_level"`
	ScoreType string `yaml:"score_type"`
	ScoreTypeParameters [][2]int `yaml:"score_type_parameters"`
	ScoreMode string `yaml:"score_mode"`
	TotalValue int `yaml:"total_value"`
}

type Problem struct {
	TaskYAML

	StatementList problems.Contents
	AttachmentList []problems.Attachment

	InputPathPattern string
	AnswerPathPattern string

	Path string
}

func (p Problem) Name() string {
	return filepath.Base(p.Path)
}

func (p Problem) Titles() problems.Contents {
	return []problems.Content{{"hungarian", []byte(p.Title), "text"}}
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
	lst1 := language.List()

	lst2 := make([]language.Language, 0, len(lst1))
	for _, val := range lst1 {
		if val.Id() != "zip" {
			lst2 = append(lst2, val)
		}
	}

	return lst2
}

func (p Problem) Attachments() []problems.Attachment {
	return p.AttachmentList
}

func (p Problem) Tags() []string {
	return make([]string, 0)
}

func (p Problem) StatusSkeleton(name string) (*problems.Status, error) {
	ans := problems.Status{false, "status skeleton", problems.FEEDBACK_IOI, make([]problems.Testset, 0)}
	ans.Feedback = append(ans.Feedback, problems.Testset{Name: "tests"})
	testset := &ans.Feedback[len(ans.Feedback)-1]

	tcByGroup := make(map[string][]problems.Testcase)
	subtask := 0
	testsLeft := p.ScoreTypeParameters[subtask][1]
	subtasks := make([]string, len(p.ScoreTypeParameters))
	for ind := 0; ind < p.InputCount; ind++ {
		tc := problems.Testcase{}
		tc.InputPath, tc.AnswerPath = fmt.Sprintf(p.InputPathPattern, ind), fmt.Sprintf(p.AnswerPathPattern, ind)
		tc.Index = ind + 1
		tc.MaxScore = 0
		tc.VerdictName = problems.VERDICT_DR
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
		group.Scoring = problems.SCORING_GROUP
		for _, tc := range tcByGroup[subtask] {
			group.Testcases = append(group.Testcases, tc)
			testset.Testcases = append(testset.Testcases, tc)
		}
	}



	return &ans, nil
}


func (p Problem) Check(tc *problems.Testcase) error {
	checkerPath := filepath.Join(p.Path, "check", "checker")

	stdout, stderr := bytes.Buffer{}, bytes.Buffer{}

	cmd := exec.Command(checkerPath, tc.InputPath, tc.AnswerPath, tc.OutputPath)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		return errors.Wrap(err, "can't check task_yaml task")
	}

	fmt.Fscanf(&stdout, "%f", &tc.Score)

	if tc.Score == 1.0 {
		tc.VerdictName = problems.VERDICT_AC
	}else if tc.Score > 0 {
		tc.VerdictName = problems.VERDICT_PC
	}else {
		tc.VerdictName = problems.VERDICT_WA
	}

	tc.Score *= tc.MaxScore

	tc.CheckerOutput = stderr.String()
	return nil
}

func (p Problem) Files() []problems.File {
	return make([]problems.File, 0)
}

func (p Problem) TaskTypeName() string {
	return "batch"
}

func parser(path string) (problems.Problem, error) {
	fmt.Println(path)
	p := Problem{Path: path, InputPathPattern: filepath.Join(path, "input", "input%d.txt"), AnswerPathPattern: filepath.Join(path, "output", "output%d.txt"), AttachmentList: make([]problems.Attachment, 0)}

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

	p.StatementList = make([]problems.Content, 0)
	p.StatementList = append(p.StatementList, problems.Content{"hungarian", statementPDF, "application/pdf"})

	checkerPath := filepath.Join(p.Path, "check", "checker")
	if _, err := os.Stat(checkerPath); os.IsNotExist(err) {
		if checkerBinary, err := os.Create(checkerPath); err == nil {
			defer checkerBinary.Close()

			if lang := language.Get("cpp14"); lang != nil {
				if checkerFile, err := os.Open(checkerPath+".cpp"); err == nil {
					defer checkerFile.Close()

					if err := lang.InsecureCompile(filepath.Join(path, "check"), checkerFile, checkerBinary, os.Stderr); err != nil {
						return nil, err
					}

					if err := os.Chmod(checkerPath, os.ModePerm); err != nil {
						return nil, err
					}
				} else {
					return nil, err
				}
			} else {
				return nil, errors.New("error while parsing task_yaml problem can't compile task_yaml checker because there's no cpp14 compiler")
			}
		} else {
			return nil, err
		}
	}

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

			p.AttachmentList = append(p.AttachmentList, problems.Attachment{file.Name(),cont})
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
