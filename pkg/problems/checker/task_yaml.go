package checker

import (
	"bytes"
	"fmt"
	"os/exec"

	"github.com/mraron/njudge/pkg/problems"
)

type TaskYAML struct {
	path string
}

func NewTaskYAML(path string) TaskYAML {
	return TaskYAML{path: path}
}

func (TaskYAML) Name() string {
	return "taskyaml"
}

func (t TaskYAML) Check(tc *problems.Testcase) error {
	checkerPath := t.path

	stdout, stderr := bytes.Buffer{}, bytes.Buffer{}

	cmd := exec.Command(checkerPath, tc.InputPath, tc.AnswerPath, tc.OutputPath)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("can't check task_yaml task: %w", err)
	}
	fmt.Fscanf(&stdout, "%f", &tc.Score)

	if tc.Score == 1.0 {
		tc.VerdictName = problems.VerdictAC
	} else if tc.Score > 0 {
		tc.VerdictName = problems.VerdictPC
	} else {
		tc.VerdictName = problems.VerdictWA
	}

	tc.Score *= tc.MaxScore

	tc.CheckerOutput = stderr.String()
	return nil
}
