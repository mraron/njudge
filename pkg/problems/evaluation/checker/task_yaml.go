package checker

import (
	"bytes"
	"context"
	"fmt"
	"github.com/mraron/njudge/pkg/language/runner"
	"github.com/mraron/njudge/pkg/problems"
)

// TaskYAML checker format is used by CMS as described in
// the CMS documentation's [Checker] and [Standard manager output] sections.
//
// [Checker]: https://cms.readthedocs.io/en/v1.4/Task%20types.html#checker
// [Standard manager output]: https://cms.readthedocs.io/en/v1.4/Task%20types.html#standard-manager-output
type TaskYAML struct {
	path string

	executable runner.Executable
}

func NewTaskYAML(path string) *TaskYAML {
	return &TaskYAML{
		path:       path,
		executable: runner.NewStdlib(path),
	}
}

func (t *TaskYAML) Name() string {
	return "taskyaml"
}

func (t *TaskYAML) Check(ctx context.Context, testcase *problems.Testcase) error {
	tc := testcase
	stdout, stderr := bytes.Buffer{}, bytes.Buffer{}

	t.executable.Stdout(&stdout)
	t.executable.Stderr(&stderr)

	if err := t.executable.Run([]string{tc.InputPath, tc.AnswerPath, tc.OutputPath}); err != nil {
		return fmt.Errorf("can't check task_yaml task: %w", err)
	}
	if _, err := fmt.Fscanf(&stdout, "%f", &tc.Score); err != nil {
		return err
	}

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
