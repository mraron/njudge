package feladat_txt

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"github.com/mraron/njudge/pkg/language/memory"
	"github.com/mraron/njudge/pkg/problems/evaluation"
	"github.com/mraron/njudge/pkg/problems/evaluation/batch"
	"github.com/mraron/njudge/pkg/problems/executable/checker"
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
)

type Problem struct {
	Path           string
	ShortName      string
	Title          string
	StatementList  problems.Contents
	AttachmentList problems.Attachments
	TestCount      int
	MemoryLimitKB  int
	TimeLimitS     float64
	SubtaskCount   int
	Points         []int

	InputPathPattern  string
	AnswerPathPattern string
}

func (p Problem) Name() string {
	return p.ShortName
}

func (p Problem) Titles() problems.Contents {
	return problems.Contents{problems.BytesData{Loc: "hungarian", Val: []byte(p.Title), Typ: "text"}}
}

func (p Problem) Statements() problems.Contents {
	return p.StatementList
}

func (p Problem) MemoryLimit() memory.Amount {
	return memory.Amount(p.MemoryLimitKB) * memory.KiB
}

func (p Problem) TimeLimit() int {
	return int(p.TimeLimitS * float64(1000))
}

func (p Problem) InputOutputFiles() (string, string) {
	return "", ""
}

func (p Problem) Languages() []language.Language {
	return language.ListExcept(language.DefaultStore, []string{"zip"})
}

func (p Problem) Attachments() problems.Attachments {
	return p.AttachmentList
}

func (p Problem) Tags() []string {
	return make([]string, 0)
}

func (p Problem) StatusSkeleton(name string) (*problems.Status, error) {
	ans := problems.Status{Compiled: false, CompilerOutput: "", FeedbackType: problems.FeedbackIOI, Feedback: make([]problems.Testset, 0)}
	ans.Feedback = append(ans.Feedback, problems.Testset{Name: "tests"})
	testset := &ans.Feedback[len(ans.Feedback)-1]

	testcaseByGroup := make(map[string][]problems.Testcase)
	for ind := 0; ind < p.TestCount; ind++ {
		tc := problems.Testcase{}
		tc.InputPath, tc.AnswerPath = fmt.Sprintf(p.InputPathPattern, ind+1), fmt.Sprintf(p.AnswerPathPattern, ind+1)
		tc.Index = ind + 1

		pointSum := 0.0
		for x := 0; x < p.SubtaskCount; x++ {
			pointSum = pointSum + float64(p.Points[x*p.TestCount+ind])
		}

		tc.MaxScore = pointSum

		if len(testcaseByGroup[tc.Group]) == 0 {
			testcaseByGroup[tc.Group] = make([]problems.Testcase, 0)
		}

		testcaseByGroup[tc.Group] = append(testcaseByGroup[tc.Group], tc)
	}

	idx := 1

	testset.Groups = append(testset.Groups, problems.Group{})
	group := &testset.Groups[len(testset.Groups)-1]

	group.Name = "base"
	group.Scoring = problems.ScoringSum

	for _, tc := range testcaseByGroup[""] {
		testcase := problems.Testcase{
			Index:          idx,
			InputPath:      tc.InputPath,
			OutputPath:     "",
			AnswerPath:     tc.AnswerPath,
			Testset:        "tests",
			Group:          "base",
			VerdictName:    problems.VerdictDR,
			Score:          float64(0.0),
			MaxScore:       float64(tc.MaxScore),
			Output:         "-",
			ExpectedOutput: "-",
			CheckerOutput:  "-",
			TimeSpent:      0 * time.Millisecond,
			MemoryUsed:     0,
			TimeLimit:      time.Duration(p.TimeLimit()) * time.Millisecond,
			MemoryLimit:    p.MemoryLimit(),
		}
		group.Testcases = append(group.Testcases, testcase)

		idx++
	}

	return &ans, nil
}

func (p Problem) Checker() problems.Checker {
	return checker.NewEllen(filepath.Join(p.Path, "ellen"), p.Path, p.TestCount, p.Points)
}

func (p Problem) EvaluationFiles() []problems.EvaluationFile {
	return make([]problems.EvaluationFile, 0)
}

func (p Problem) GetTaskType() problems.TaskType {
	return batch.New(evaluation.CompileCheckSupported{
		List:         p.Languages(),
		NextCompiler: evaluation.Compile{},
	}, evaluation.BasicRunnerWithChecker(p.Checker()))
}

func Parse(fs afero.Fs, path string) (problems.Problem, error) {
	f, err := fs.Open(filepath.Join(path, "feladat.txt"))
	if err != nil {
		return nil, err
	}
	defer f.Close()

	br := bufio.NewReader(f)

	var (
		s string
	)

	ind := 0
	lst := make([]string, 0)

	for err == nil {
		s, err = br.ReadString(';')

		if err == nil {
			if s != "\n" {
				str := strings.TrimSpace(s)
				lst = append(lst, str[:len(str)-1])
			}
		}
	}

	p := &Problem{}

	p.Path = path
	p.ShortName = filepath.Base(path)

	p.Title = lst[ind]
	ind++

	if lst[ind] != "NO" {
		return nil, errors.New("modules not supported")
	}
	ind++

	p.TestCount, err = strconv.Atoi(lst[ind])
	if err != nil {
		return nil, err
	}
	ind++

	p.MemoryLimitKB, err = strconv.Atoi(lst[ind])
	if err != nil {
		return nil, err
	}
	ind++

	p.TimeLimitS, err = strconv.ParseFloat(lst[ind], 64)
	if err != nil {
		return nil, err
	}
	ind++

	p.SubtaskCount, err = strconv.Atoi(lst[ind])
	if err != nil {
		return nil, err
	}
	ind++

	p.Points = make([]int, p.SubtaskCount*p.TestCount)

	for i := 0; i < p.SubtaskCount*p.TestCount; i++ {
		p.Points[i], err = strconv.Atoi(lst[ind])
		if err != nil {
			return nil, err
		}
		ind++
	}

	p.StatementList = make(problems.Contents, 0)
	feladat_pdf, err := afero.ReadFile(fs, filepath.Join(path, "feladat.pdf"))
	if err != nil {
		return nil, err
	}
	p.StatementList = append(p.StatementList, problems.BytesData{Loc: "hungarian", Val: feladat_pdf, Typ: "application/pdf"})

	box, _ := sandbox.NewDummy()
	if err := cpp.AutoCompile(context.TODO(), fs, box, path, filepath.Join(path, "ellen.cpp"), filepath.Join(path, "ellen")); err != nil {
		return nil, err
	}

	p.AttachmentList = make(problems.Attachments, 0)
	if _, err = fs.Stat(filepath.Join(path, "minta.zip")); err == nil {
		cont, err := afero.ReadFile(fs, filepath.Join(path, "minta.zip"))
		if err != nil {
			return nil, err
		}
		p.AttachmentList = append(p.AttachmentList, problems.BytesData{Nam: "minta.zip", Val: cont})
	}

	p.InputPathPattern = filepath.Join(p.Path, "in.%d")
	p.AnswerPathPattern = filepath.Join(p.Path, "out.%d")

	return p, nil
}

func Identify(fs afero.Fs, path string) bool {
	_, err := fs.Stat(filepath.Join(path, "feladat.txt"))
	return !os.IsNotExist(err)
}

func init() {
	_ = problems.RegisterConfigType("feladat_txt", Parse, Identify)
}
