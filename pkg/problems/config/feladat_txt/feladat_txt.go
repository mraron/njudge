package feladat_txt

import (
	"bufio"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/mraron/njudge/pkg/language"
	"github.com/mraron/njudge/pkg/language/cpp"
	"github.com/mraron/njudge/pkg/problems"
	"github.com/mraron/njudge/pkg/problems/checker"
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

func (p Problem) HTMLStatements() problems.Contents {
	return p.StatementList.FilterByType("text/html")
}

func (p Problem) PDFStatements() problems.Contents {
	return p.StatementList.FilterByType("application/pdf")
}

func (p Problem) MemoryLimit() int {
	return 1024 * p.MemoryLimitKB
}

func (p Problem) TimeLimit() int {
	return int(p.TimeLimitS * float64(1000))
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

	tcbygroup := make(map[string][]problems.Testcase)
	for ind := 0; ind < p.TestCount; ind++ {
		tc := problems.Testcase{}
		tc.InputPath, tc.AnswerPath = fmt.Sprintf(p.InputPathPattern, ind+1), fmt.Sprintf(p.AnswerPathPattern, ind+1)
		tc.Index = ind + 1

		points_sum := 0.0
		for x := 0; x < p.SubtaskCount; x++ {
			points_sum = points_sum + float64(p.Points[x*p.TestCount+ind])
		}

		tc.MaxScore = points_sum

		if len(tcbygroup[tc.Group]) == 0 {
			tcbygroup[tc.Group] = make([]problems.Testcase, 0)
		}

		tcbygroup[tc.Group] = append(tcbygroup[tc.Group], tc)
	}

	idx := 1

	testset.Groups = append(testset.Groups, problems.Group{})
	group := &testset.Groups[len(testset.Groups)-1]

	group.Name = "base"
	group.Scoring = problems.ScoringSum

	for _, tc := range tcbygroup[""] {
		testcase := problems.Testcase{idx, tc.InputPath, "", tc.AnswerPath, "tests", "base", problems.VerdictDR, float64(0.0), float64(tc.MaxScore), "-", "-", "-", 0 * time.Millisecond, 0, time.Duration(p.TimeLimit()) * time.Millisecond, p.MemoryLimit()}
		group.Testcases = append(group.Testcases, testcase)

		idx++
	}

	return &ans, nil
}

func (p Problem) Checker() problems.Checker {
	return checker.NewFeladatTXT(filepath.Join(p.Path, "ellen"), p.Path, p.TestCount, p.Points)
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

func parser(path string) (problems.Problem, error) {
	f, err := os.Open(filepath.Join(path, "feladat.txt"))
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

	feladat_pdf, err := os.Open(filepath.Join(path, "feladat.pdf"))
	if err != nil {
		return nil, err
	}
	defer feladat_pdf.Close()

	cont, err := ioutil.ReadAll(feladat_pdf)
	if err != nil {
		return nil, err
	}

	p.StatementList = make(problems.Contents, 0)
	p.StatementList = append(p.StatementList, problems.BytesData{Loc: "hungarian", Val: cont, Typ: "application/pdf"})

	p.AttachmentList = make(problems.Attachments, 0)

	if _, err := os.Stat(filepath.Join(path, "ellen")); os.IsNotExist(err) {
		if checkerBinary, err := os.Create(filepath.Join(path, "ellen")); err == nil {
			defer checkerBinary.Close()
			if checkerFile, err := os.Open(filepath.Join(path, "ellen.cpp")); err == nil {
				defer checkerFile.Close()

				if err := cpp.Std14.InsecureCompile(path, checkerFile, checkerBinary, os.Stderr); err != nil {
					return nil, err
				}

				if err := os.Chmod(filepath.Join(path, "ellen"), os.ModePerm); err != nil {
					return nil, err
				}
			} else {
				return nil, err
			}

		} else {
			return nil, err
		}
	}

	if _, err = os.Stat(filepath.Join(path, "minta.zip")); err == nil {
		cont, err := ioutil.ReadFile(filepath.Join(path, "minta.zip"))
		if err != nil {
			return nil, err
		}
		p.AttachmentList = append(p.AttachmentList, problems.BytesData{Nam: "minta.zip", Val: cont})
	}

	p.InputPathPattern = filepath.Join(p.Path, "in.%d")
	p.AnswerPathPattern = filepath.Join(p.Path, "out.%d")

	return p, nil
}

func identifier(path string) bool {
	_, err := os.Stat(filepath.Join(path, "feladat.txt"))
	return !os.IsNotExist(err)
}

func init() {
	problems.RegisterConfigType("feladat_txt", parser, identifier)
}
