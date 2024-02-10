package language

import (
	"bytes"
	"context"
	"fmt"
	"github.com/mraron/njudge/pkg/language/memory"
	"github.com/mraron/njudge/pkg/language/sandbox"
	"testing"
	"time"
)

// Testable is an interface that can be used to test Languages by providing a sandbox.Sandbox
// This will probably run a bunch of tests (most likely via defining multiple Test objects and calling Test.Run).
type Testable interface {
	Test(*testing.T, sandbox.Sandbox) error
}

// Test is a struct that holds a test for a language.
type Test struct {
	Name            string
	Language        Language
	Source          string
	Input           string
	ExpectedOutput  string
	ExpectedVerdict sandbox.Verdict
	TimeLimit       time.Duration
	MemoryLimit     memory.Amount
}

// Run the test given a sandbox.
func (test Test) Run(s sandbox.Sandbox) error {
	err := s.Init(context.Background())
	if err != nil {
		return err
	}

	src := bytes.NewBufferString(test.Source)
	stderr := &bytes.Buffer{}

	compiledBinary, err := test.Language.Compile(context.TODO(), s, sandbox.File{test.Language.DefaultFilename(), src}, stderr, nil)
	stderrContent := stderr.String()

	if (test.ExpectedVerdict&sandbox.VerdictCE == 0 && err != nil) || (test.ExpectedVerdict&sandbox.VerdictCE != 0 && err == nil && stderrContent == "") {
		return fmt.Errorf("error: %s stderr: %s", err, stderrContent)
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
		status, err := test.Language.Run(context.TODO(), s, *compiledBinary, bytes.NewBufferString(test.Input), output, test.TimeLimit, test.MemoryLimit)

		outputContent := output.String()
		if status.Verdict&test.ExpectedVerdict == 0 || err != nil || outputContent != test.ExpectedOutput {
			return fmt.Errorf("EXPECTED %s got %s, source %q\n error: %s status: %v output: %q expected output: %q", test.ExpectedVerdict, status.Verdict, test.Source, err, status, outputContent, test.ExpectedOutput)
		}

		err = s.Cleanup(context.Background())
		if err != nil {
			return fmt.Errorf("cleanup err: %w", err)
		}
	}

	return nil
}
