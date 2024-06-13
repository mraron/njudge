package checker

import (
	"bytes"
	"context"
	"fmt"
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
	testIndex := strconv.Itoa(tc.Index)

	dir, err := os.MkdirTemp("/tmp", "feladat_txt_checker")
	if err != nil {
		return err
	}
	defer func(path string) {
		_ = os.RemoveAll(path)
	}(dir)

	requiredName := filepath.Base(tc.AnswerPath) // for example in.5
	participantOutput := filepath.Join(dir, requiredName)

	err = os.Symlink(tc.OutputPath, participantOutput)
	if err != nil {
		return err
	}

	stdout := &bytes.Buffer{}
	stderr := &bytes.Buffer{}

	args := []string{f.ellenPath, f.testcaseDir, dir, testIndex}
	cmd := exec.Command("/bin/sh", "-c", "ulimit -s unlimited && "+strings.Join(args, " "))
	cmd.Stdout = stdout
	cmd.Stderr = stderr

	err = cmd.Run()

	str := stdout.String()
	tc.CheckerOutput = problems.Base64String(problems.Truncate(str))
	if err == nil || strings.HasPrefix(err.Error(), "exit status") {
		var splitted []string
		if strings.Contains(str, ":") {
			splitted = strings.Split(strings.TrimSpace(str), ":")
		} else {
			splitted = strings.Split(strings.TrimSpace(str), "\n")
		}

		score := 0.0
		allOk := true
		for i := 0; i < len(splitted); i++ {
			splitted[i] = strings.TrimSpace(splitted[i])
			if len(splitted[i]) == 0 {
				continue
			}

			curr := strings.Split(splitted[i], ";")
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
	}

	tc.VerdictName = problems.VerdictXX
	return err
}
