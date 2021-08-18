package feladat_txt

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"github.com/mraron/njudge/utils/language"
	"github.com/mraron/njudge/utils/problems"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

type Problem struct {
	Path           string
	ShortName      string
	Title          string
	StatementList  problems.Contents
	AttachmentList []problems.Attachment
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
	return problems.Contents{problems.Content{"hungarian", []byte(p.Title), "text"}}
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
	group.Scoring = problems.SCORING_SUM

	for _, tc := range tcbygroup[""] {
		testcase := problems.Testcase{idx, tc.InputPath, "", tc.AnswerPath, "tests", "base", problems.VERDICT_DR, float64(0.0), float64(tc.MaxScore), "-", "-", "-", 0 * time.Millisecond, 0, time.Duration(p.TimeLimit()) * time.Millisecond, p.MemoryLimit()}
		group.Testcases = append(group.Testcases, testcase)
		testset.Testcases = append(testset.Testcases, testcase)

		idx++
	}

	return &ans, nil
}

func (p Problem) Check(tc *problems.Testcase) error {
	testind := strconv.Itoa(tc.Index)

	dir, err := ioutil.TempDir("/tmp", "feladat_txt_checker")
	if err != nil {
		return err
	}

	pout_tmp := filepath.Join(dir, filepath.Base(tc.AnswerPath))

	err = os.Symlink(tc.OutputPath, pout_tmp)
	if err != nil {
		return err
	}

	stdout := &bytes.Buffer{}
	stderr := &bytes.Buffer{}

	cmd := exec.Command("/bin/sh", "-c", "ulimit -s unlimited && "+strings.Join([]string{filepath.Join(p.Path, "ellen"), p.Path, dir, testind}, " "))
	cmd.Stdout = stdout
	cmd.Stderr = stderr

	err = cmd.Run()

	str := stdout.String()
	tc.CheckerOutput = problems.Truncate(str)
	if err == nil || strings.HasPrefix(err.Error(), "exit status") {
		var spltd []string
		if strings.Contains(str, ":") {
			spltd = strings.Split(strings.TrimSpace(str), ":")
		}else {
			spltd = strings.Split(strings.TrimSpace(str), "\n")
		}

		score := 0.0
		allOk := true
		for i := 0; i < len(spltd); i++ {
			spltd[i] = strings.TrimSpace(spltd[i])
			curr := strings.Split(spltd[i], ";")

			if strings.TrimSpace(curr[len(curr)-2]) == "1" {
				score = score + float64(p.Points[i*p.TestCount+tc.Index-1])
			}else {
				allOk = false
			}
		}

		tc.Score = score
		if score == tc.MaxScore && allOk {
			tc.VerdictName = problems.VERDICT_AC
		} else if score != 0.0 {
			tc.VerdictName = problems.VERDICT_PC
		} else {
			tc.VerdictName = problems.VERDICT_WA
		}

		return nil
	} else if err != nil {
		tc.VerdictName = problems.VERDICT_XX
		return err
	}

	tc.VerdictName = problems.VERDICT_XX
	return errors.New("proccess state is not success")
}

func (p Problem) Files() []problems.File {
	return make([]problems.File, 0)
}

func (p Problem) TaskTypeName() string {
	return "batch"
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
	p.StatementList = append(p.StatementList, problems.Content{"hungarian", cont, "application/pdf"})

	p.AttachmentList = make([]problems.Attachment, 0)

	if _, err := os.Stat(filepath.Join(path, "ellen")); os.IsNotExist(err) {
		if checkerBinary, err := os.Create(filepath.Join(path, "ellen")); err == nil {
			defer checkerBinary.Close()

			if lang := language.Get("cpp11"); lang != nil {
				if checkerFile, err := os.Open(filepath.Join(path, "ellen.cpp")); err == nil {
					defer checkerFile.Close()

					if err := lang.InsecureCompile(path, checkerFile, checkerBinary, os.Stderr); err != nil {
						return nil, err
					}

					if err := os.Chmod(filepath.Join(path, "ellen"), os.ModePerm); err != nil {
						return nil, err
					}
				} else {
					return nil, err
				}
			} else {
				return nil, errors.New("error while parsing feladat_txt problem can't compile feladat_txt checker because there's no cpp11 compiler")
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
		p.AttachmentList = append(p.AttachmentList, problems.Attachment{"minta.zip", cont})
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
