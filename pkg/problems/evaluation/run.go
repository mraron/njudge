package evaluation

import (
	"bytes"
	"github.com/mraron/njudge/pkg/language"
	"github.com/mraron/njudge/pkg/language/sandbox"
	"github.com/mraron/njudge/pkg/problems"
	"golang.org/x/net/context"
	"io"
	"os"
)

type ACRunner struct{}

func (a ACRunner) SetSolution(ctx context.Context, solution problems.Solution) error {
	return nil
}

func (a ACRunner) Run(ctx context.Context, sandboxProvider sandbox.Provider, testcase *problems.Testcase) error {
	testcase.VerdictName = problems.VerdictAC
	return nil
}

type BasicRunner struct {
	maxSizeInTestcase int

	inputFile  string
	outputFile string
	checker    problems.Checker

	lang language.Language
	bin  []byte
}

type BasicRunnerOption func(r *BasicRunner)

func BasicRunnerWithFiles(inputFile, outputFile string) BasicRunnerOption {
	return func(r *BasicRunner) {
		r.inputFile = inputFile
		r.outputFile = outputFile
	}
}

func BasicRunnerWithChecker(c problems.Checker) BasicRunnerOption {
	return func(r *BasicRunner) {
		r.checker = c
	}
}

func NewBasicRunner(options ...BasicRunnerOption) *BasicRunner {
	res := &BasicRunner{
		maxSizeInTestcase: 1 << 6,
	}
	for _, opt := range options {
		opt(res)
	}
	return res
}

func (r *BasicRunner) SetSolution(ctx context.Context, solution problems.Solution) error {
	r.lang = solution.GetLanguage()

	rc, err := solution.GetFile(ctx)
	if err != nil {
		return err
	}
	defer func() {
		_ = rc.Close()
	}()

	r.bin, err = io.ReadAll(rc)
	return err
}

func (r *BasicRunner) prepareIO(s sandbox.Sandbox, testcase *problems.Testcase) (io.ReadCloser, io.WriteCloser, error) {
	inputFile, err := os.Open(testcase.InputPath)
	if err != nil {
		return nil, nil, err
	}
	defer func(inputFile *os.File) {
		_ = inputFile.Close()
	}(inputFile)

	var (
		sandboxInput  io.ReadCloser  = inputFile
		sandboxOutput io.WriteCloser = nil
	)
	if r.inputFile != "" {
		if err = sandbox.CreateFileFromSource(s, r.inputFile, inputFile); err != nil {
			return nil, nil, err
		}
		if sandboxInput, err = s.Open(r.inputFile); err != nil {
			return nil, nil, err
		}
	}

	if r.outputFile == "" {
		r.outputFile = "output"
	}
	if sandboxOutput, err = s.Create(r.outputFile); err != nil {
		_ = sandboxInput.Close()
		return nil, nil, err
	}

	return sandboxInput, sandboxOutput, nil
}

func (r *BasicRunner) getFilePrefix(name string) ([]byte, error) {
	f, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer func(f *os.File) {
		_ = f.Close()
	}(f)

	buf := make([]byte, r.maxSizeInTestcase)
	n, err := f.Read(buf)
	if err != nil {
		return nil, err
	}

	buf = buf[:n]
	if n == r.maxSizeInTestcase {
		buf = append(buf, []byte("...")...)
	}
	return buf, nil
}

func (r *BasicRunner) Run(ctx context.Context, sandboxProvider sandbox.Provider, testcase *problems.Testcase) error {
	s, err := sandboxProvider.Get()
	if err != nil {
		return err
	}
	defer sandboxProvider.Put(s)

	sandboxInput, sandboxOutput, err := r.prepareIO(s, testcase)
	defer func(sandboxInput io.ReadCloser, sandboxOutput io.WriteCloser) {
		_ = sandboxInput.Close()
		_ = sandboxOutput.Close()
	}(sandboxInput, sandboxOutput)

	status, err := r.lang.Run(s, bytes.NewBuffer(r.bin), sandboxInput, sandboxOutput, testcase.TimeLimit, testcase.MemoryLimit)

	testcase.OutputPath = (sandboxOutput.(*os.File)).Name()
	testcase.TimeSpent = status.Time
	testcase.MemoryLimit = status.Memory

	var (
		expectedOutput []byte
		output         []byte
	)
	if expectedOutput, err = r.getFilePrefix(testcase.AnswerPath); err != nil {
		return err
	}
	if output, err = r.getFilePrefix(testcase.OutputPath); err != nil {
		return err
	}

	testcase.ExpectedOutput = string(expectedOutput)
	testcase.Output = string(output)

	switch status.Verdict {
	case sandbox.VerdictOK:
		testcase.VerdictName = problems.VerdictAC // checker can overwrite it
	case sandbox.VerdictTL:
		testcase.VerdictName = problems.VerdictTL
		return err
	case sandbox.VerdictML:
		testcase.VerdictName = problems.VerdictML
		return err
	case sandbox.VerdictRE:
		testcase.VerdictName = problems.VerdictRE
		return err
	case sandbox.VerdictXX:
		testcase.VerdictName = problems.VerdictXX
		return err
	case sandbox.VerdictCE:
		panic("solution should've been already compiled")
	}

	if r.checker == nil {
		return nil
	}
	return r.checker.Check(ctx, testcase)
}
