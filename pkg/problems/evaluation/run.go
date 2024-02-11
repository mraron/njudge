package evaluation

import (
	"bytes"
	"context"
	"errors"
	"github.com/mraron/njudge/pkg/language"
	"github.com/mraron/njudge/pkg/language/sandbox"
	"github.com/mraron/njudge/pkg/problems"
	"github.com/spf13/afero"
	"io"
	"os"
	"path/filepath"
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

	lang    language.Language
	binName string
	bin     []byte

	fs afero.Fs
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

func BasicRunnerWithFs(fs afero.Fs) BasicRunnerOption {
	return func(r *BasicRunner) {
		r.fs = fs
	}
}

func NewBasicRunner(options ...BasicRunnerOption) *BasicRunner {
	res := &BasicRunner{
		maxSizeInTestcase: 1 << 6,
		fs:                afero.NewOsFs(),
	}
	for _, opt := range options {
		opt(res)
	}
	return res
}

func (r *BasicRunner) SetSolution(ctx context.Context, solution problems.Solution) error {
	r.lang = solution.GetLanguage()

	file, err := solution.GetFile(ctx)
	if err != nil {
		return err
	}

	r.binName = file.Name
	r.bin, err = io.ReadAll(file.Source)
	return errors.Join(err, file.Source.Close())
}

func (r *BasicRunner) prepareIO(s sandbox.Sandbox, testcase *problems.Testcase) (io.ReadCloser, io.WriteCloser, error) {
	inputFile, err := r.fs.Open(testcase.InputPath)
	if err != nil {
		return nil, nil, err
	}

	var (
		sandboxInput  io.ReadCloser  = inputFile
		sandboxOutput io.WriteCloser = nil
	)
	if r.inputFile != "" {
		//TODO maybe link it with restricted permissions?
		if err = sandbox.CreateFile(s, sandbox.File{Name: r.inputFile, Source: inputFile}); err != nil {
			return nil, nil, err
		}
		sandboxInput = nil
	}

	if r.outputFile == "" {
		r.outputFile = "output"
		if sandboxOutput, err = s.Create(r.outputFile); err != nil {
			return sandboxInput, nil, err
		}
	}

	return sandboxInput, sandboxOutput, nil
}

func (r *BasicRunner) getReadCloserPrefix(closer io.ReadCloser) ([]byte, error) {
	buf := make([]byte, r.maxSizeInTestcase)
	n, err := closer.Read(buf)
	defer func(closer io.ReadCloser) {
		_ = closer.Close()
	}(closer)
	if err != nil && !errors.Is(err, io.EOF) {
		return nil, err
	}

	buf = buf[:n]
	if n == r.maxSizeInTestcase {
		buf = append(buf, []byte("...")...)
	}
	return buf, nil
}

func (r *BasicRunner) setOutputExpectedOutput(s sandbox.Sandbox, testcase *problems.Testcase) error {
	var (
		expectedOutput []byte
		output         []byte

		err error
	)

	answerFile, err := r.fs.Open(testcase.AnswerPath)
	if err != nil {
		return err
	}
	if expectedOutput, err = r.getReadCloserPrefix(answerFile); err != nil {
		return err
	}

	outputFile, err := s.Open(filepath.Base(testcase.OutputPath))
	if output, err = r.getReadCloserPrefix(outputFile); err != nil {
		return err
	}

	testcase.ExpectedOutput = string(expectedOutput)
	testcase.Output = string(output)
	return nil
}

func (r *BasicRunner) Run(ctx context.Context, sandboxProvider sandbox.Provider, testcase *problems.Testcase) error {
	s, err := sandboxProvider.Get()
	if err != nil {
		return err
	}
	defer sandboxProvider.Put(s)
	if err = s.Init(ctx); err != nil {
		return err
	}
	defer func(s sandbox.Sandbox, ctx context.Context) {
		_ = s.Cleanup(ctx)
	}(s, ctx)

	sandboxInput, sandboxOutput, err := r.prepareIO(s, testcase)
	if err != nil {
		return err
	}
	defer func(sandboxInput io.ReadCloser, sandboxOutput io.WriteCloser) {
		if sandboxInput != nil {
			_ = sandboxInput.Close()
		}
		if sandboxOutput != nil {
			_ = sandboxOutput.Close()
		}
	}(sandboxInput, sandboxOutput)

	status, err := r.lang.Run(ctx, s, sandbox.File{
		Name:   r.binName,
		Source: io.NopCloser(bytes.NewBuffer(r.bin)),
	}, sandboxInput, sandboxOutput, testcase.TimeLimit, testcase.MemoryLimit)

	if sandboxOutput != nil {
		testcase.OutputPath = (sandboxOutput.(*os.File)).Name()
	} else {
		testcase.OutputPath = filepath.Join(s.Pwd(), r.outputFile)
		if _, err := os.Stat(testcase.OutputPath); errors.Is(err, os.ErrNotExist) {
			if f, err := os.Create(testcase.OutputPath); err != nil {
				return err
			} else {
				_ = f.Close()
			}
		}
	}
	testcase.TimeSpent = status.Time
	testcase.MemoryUsed = status.Memory

	if err = r.setOutputExpectedOutput(s, testcase); err != nil {
		return err
	}

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
