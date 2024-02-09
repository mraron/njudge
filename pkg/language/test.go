package language

import (
	"bytes"
	"fmt"
	"github.com/mraron/njudge/pkg/language/memory"
	"github.com/mraron/njudge/pkg/language/sandbox"
	"golang.org/x/net/context"
	"testing"
	"time"
)

type Testable interface {
	Test(*testing.T, sandbox.Sandbox) error
}

type Test struct {
	Name            string
	Language        Language
	Source          string
	ExpectedVerdict sandbox.Verdict
	Input           string
	ExpectedOutput  string
	TimeLimit       time.Duration
	MemoryLimit     memory.Amount
}

func (test Test) Run(s sandbox.Sandbox) error {
	err := s.Init(context.Background())
	if err != nil {
		return err
	}

	src := bytes.NewBufferString(test.Source)
	bin := &bytes.Buffer{}
	stderr := &bytes.Buffer{}

	err = test.Language.Compile(s, File{test.Language.DefaultFileName(), src}, bin, stderr, []File{})
	stderrContent := stderr.String()

	if (test.ExpectedVerdict&sandbox.VerdictCE == 0 && err != nil) || (test.ExpectedVerdict&sandbox.VerdictCE != 0 && err == nil && stderrContent == "") {
		return fmt.Errorf("error: %v stderr: %s", err, stderrContent)
	}

	err = s.Cleanup(context.Background())
	if err != nil {
		return fmt.Errorf("cleanup err: %w", err)
	}

	if test.ExpectedVerdict&sandbox.VerdictCE == 0 {
		err := s.Init(context.Background())
		if err != nil {
			return err
		}

		output := &bytes.Buffer{}
		status, err := test.Language.Run(s, bin, bytes.NewBufferString(test.Input), output, test.TimeLimit, int(test.MemoryLimit))

		outputContent := output.String()
		if status.Verdict&test.ExpectedVerdict == 0 || err != nil || outputContent != test.ExpectedOutput {
			return fmt.Errorf("EXPECTED %s got %s, source %q\n error: %v status: %v output: %q expected output: %q", test.ExpectedVerdict, status.Verdict, test.Source, err, status, outputContent, test.ExpectedOutput)
		}

		err = s.Cleanup(context.Background())
		if err != nil {
			return fmt.Errorf("cleanup err: %w", err)
		}
	}

	return nil
}
