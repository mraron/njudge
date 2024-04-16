package checker

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/mraron/njudge/pkg/problems"
)

// Ellen (an abbreviation of checker in hungarian) checker is used by [mester], [bíró] systems.
//
// Arguments of the checker are:
//
//  1. Absolute path to testdata
//  2. Absolute path to output data
//  3. Index of testcase to check
//
// Testcases are named in.{x} and out.{x} where x is an 1 indexed integer.
//
// It produces its output to stdout in the following format:
//
// {test_index};{subtest_number};{point_multiplier};{verdict_message}
//   - test_index: the index received via the argument
//   - subtest_number: refer to the [github.com/mraron/njudge/pkg/problems/config/feladat_txt] config's docs
//   - point_multiplier: 0 or 1, depending on the correctness
//   - verdict_message: A message displayed to the user (it must not contain ";" or ":" characters)
//
// If there are multiple subtests the same format should be used delimited by a single ":"
//
// [mester]: https://mester.inf.elte.hu
// [bíró]: https://biro.inf.elte.hu
type Ellen struct {
	testcaseDir string
	ellenPath   string
	testCount   int
	points      []int
}

func NewEllen(ellenPath, testcaseDir string, testCount int, points []int) Ellen {
	return Ellen{
		testcaseDir: testcaseDir,
		ellenPath:   ellenPath,
		testCount:   testCount,
		points:      points,
	}
}

func (Ellen) Name() string {
	return "feladattxt"
}

func (f Ellen) Check(ctx context.Context, testcase *problems.Testcase) error {
	tc := testcase
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
			if len(spltd[i]) == 0 {
				continue
			}

			curr := strings.Split(spltd[i], ";")
			if len(curr) < 2 {
				return fmt.Errorf("wrong format for ellen output: %q %v", str, curr)
			}

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
