package checker

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
	"syscall"

	"github.com/mraron/njudge/pkg/problems"
)

// Testlib is [polygon]'s default checker format
//
// [polygon]: https://polygon.codeforces.com
type Testlib struct {
	path string
}

func NewTestlib(path string) Testlib {
	return Testlib{path: path}
}

func (Testlib) Name() string {
	return "testlib"
}

func (t Testlib) Check(tc *problems.Testcase) error {
	output := &bytes.Buffer{}

	args := []string{t.path, tc.InputPath, tc.OutputPath, tc.AnswerPath}
	cmd := exec.Command("/bin/sh", "-c", "ulimit -s unlimited && "+strings.Join(args, " "))
	cmd.Stdout = output
	cmd.Stderr = output

	err := cmd.Run()

	tc.CheckerOutput = problems.Truncate(output.String())
	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			if status, ok := exitErr.Sys().(syscall.WaitStatus); ok {
				if status.ExitStatus() == 1 {
					tc.VerdictName = problems.VerdictWA
				} else if status.ExitStatus() == 2 {
					tc.VerdictName = problems.VerdictPE
				} else if status.ExitStatus() == 7 { //only support quitp
					tc.VerdictName = problems.VerdictPC

					rel := 0
					fmt.Sscanf(output.String(), "points %d", &rel)

					tc.Score = float64(rel) / 100.0 * tc.MaxScore
				} else { //3 -> fail
					tc.VerdictName = problems.VerdictXX
				}
			}
		} else {
			tc.VerdictName = problems.VerdictXX
			return err
		}
	} else {
		tc.Score = tc.MaxScore
		tc.VerdictName = problems.VerdictAC
	}

	return nil
}
