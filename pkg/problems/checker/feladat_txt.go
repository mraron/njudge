package checker

import (
	"bytes"
	"errors"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/mraron/njudge/pkg/problems"
)

type FeladatTXT struct {
	testcaseDir string
	ellenPath   string
	testCount   int
	points      []int
}

func NewFeladatTXT(ellenPath, testcaseDir string, testCount int, points []int) FeladatTXT {
	return FeladatTXT{
		testcaseDir: testcaseDir,
		ellenPath:   ellenPath,
		testCount:   testCount,
		points:      points,
	}
}

func (FeladatTXT) Name() string {
	return "feladattxt"
}

func (f FeladatTXT) Check(tc *problems.Testcase) error {
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

	args := []string{f.ellenPath, f.testcaseDir, dir, testind}
	cmd := exec.Command("/bin/sh", "-c", "ulimit -s unlimited && "+strings.Join(args, " "))
	cmd.Stdout = stdout
	cmd.Stderr = stderr

	err = cmd.Run()

	str := stdout.String()
	tc.CheckerOutput = problems.Truncate(str)
	if err == nil || strings.HasPrefix(err.Error(), "exit status") {
		var spltd []string
		if strings.Contains(str, ":") {
			spltd = strings.Split(strings.TrimSpace(str), ":")
		} else {
			spltd = strings.Split(strings.TrimSpace(str), "\n")
		}

		score := 0.0
		allOk := true
		for i := 0; i < len(spltd); i++ {
			spltd[i] = strings.TrimSpace(spltd[i])
			curr := strings.Split(spltd[i], ";")

			if strings.TrimSpace(curr[len(curr)-2]) == "1" {
				score = score + float64(f.points[i*f.testCount+tc.Index-1])
			} else {
				allOk = false
			}
		}

		tc.Score = score
		if score == tc.MaxScore && allOk {
			tc.VerdictName = problems.VerdictAC
		} else if score != 0.0 {
			tc.VerdictName = problems.VerdictPC
		} else {
			tc.VerdictName = problems.VerdictWA
		}

		return nil
	} else if err != nil {
		tc.VerdictName = problems.VerdictXX
		return err
	}

	tc.VerdictName = problems.VerdictXX
	return errors.New("process state is not success")
}
