package evaluation

import (
	"archive/zip"
	"bytes"
	"context"
	"errors"
	"fmt"
	"github.com/mraron/njudge/pkg/language"
	"github.com/mraron/njudge/pkg/language/memory"
	"github.com/mraron/njudge/pkg/language/sandbox"
	"github.com/mraron/njudge/pkg/problems"
	"github.com/spf13/afero"
	"io"
	"os"
	"path/filepath"
	"sync"
	"syscall"
	"time"
)

func setSolution(ctx context.Context, bin *string, dst *[]byte, solution problems.Solution) error {
	file, err := solution.GetFile(ctx)
	if err != nil {
		return err
	}

	if bin != nil {
		*bin = file.Name
	}
	*dst, err = io.ReadAll(file.Source)
	return errors.Join(err, file.Source.Close())
}

type ACRunner struct{}

func (a ACRunner) SetSolution(_ context.Context, _ problems.Solution) error {
	return nil
}

func (a ACRunner) Run(_ context.Context, _ sandbox.Provider, testcase *problems.Testcase) error {
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

var defaultOutputFile = "output"

func (r *BasicRunner) getOutputFile() string {
	if r.outputFile == "" {
		return defaultOutputFile
	}
	return r.outputFile
}

func (r *BasicRunner) SetSolution(ctx context.Context, solution problems.Solution) error {
	r.lang = solution.GetLanguage()
	return setSolution(ctx, &r.binName, &r.bin, solution)
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
		if sandboxOutput, err = s.Create(r.getOutputFile()); err != nil {
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

func mapVerdict(status *sandbox.Status, testcase *problems.Testcase) bool {
	switch status.Verdict {
	case sandbox.VerdictOK:
		testcase.VerdictName = problems.VerdictAC // checker can overwrite it
		return true
	case sandbox.VerdictTL:
		testcase.VerdictName = problems.VerdictTL
	case sandbox.VerdictML:
		testcase.VerdictName = problems.VerdictML
	case sandbox.VerdictRE:
		testcase.VerdictName = problems.VerdictRE
	case sandbox.VerdictXX:
		testcase.VerdictName = problems.VerdictXX
	case sandbox.VerdictCE:
		panic("solution should've been already compiled")
	}
	return false
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
	if err != nil {
		return err
	}

	if sandboxOutput != nil {
		testcase.OutputPath = (sandboxOutput.(*os.File)).Name()
	} else {
		testcase.OutputPath = filepath.Join(s.Pwd(), r.getOutputFile())
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

	if !mapVerdict(status, testcase) || r.checker == nil {
		return nil
	}
	return r.checker.Check(ctx, testcase)
}

type ZipRunner struct {
	checker problems.Checker
	bin     []byte
}

func NewZipRunner(checker problems.Checker) *ZipRunner {
	return &ZipRunner{checker: checker, bin: nil}
}

func (z *ZipRunner) SetSolution(ctx context.Context, solution problems.Solution) error {
	return setSolution(ctx, nil, &z.bin, solution)
}

func (z *ZipRunner) Run(ctx context.Context, _ sandbox.Provider, testcase *problems.Testcase) error {
	archive, err := zip.NewReader(bytes.NewReader(z.bin), int64(len(z.bin)))
	if err != nil {
		return err
	}
	for _, f := range archive.File {
		if f.Name == testcase.OutputPath {
			fileHandle, err := f.Open()
			if err != nil {
				return err
			}
			defer func(fileHandle io.ReadCloser) {
				_ = fileHandle.Close()
			}(fileHandle)

			tempFile, err := os.CreateTemp("", "njudge_zip_input")
			if err != nil {
				return err
			}
			defer func(tempFile *os.File, name string) {
				_ = tempFile.Close()
				_ = os.Remove(name)
			}(tempFile, tempFile.Name())

			_, err = io.CopyN(tempFile, fileHandle, 32*1024*1024) // 32MiB limit
			if err != nil && !errors.Is(err, io.EOF) {
				return err
			}

			testcase.OutputPath = tempFile.Name()
			break
		}
	}

	if z.checker == nil {
		return nil
	}
	return z.checker.Check(ctx, testcase)
}

type InteractiveRunner struct {
	lang        language.Language
	userBinName string
	userBin     []byte

	interactorBin []byte

	checker problems.Checker

	executor UserInteractorExecutor

	fs afero.Fs
}

type InteractiveRunnerOption func(runner *InteractiveRunner)

func InteractiveRunnerWithFs(fs afero.Fs) InteractiveRunnerOption {
	return func(runner *InteractiveRunner) {
		runner.fs = fs
	}
}

func InteractiveRunnerWithExecutor(executor UserInteractorExecutor) InteractiveRunnerOption {
	return func(runner *InteractiveRunner) {
		runner.executor = executor
	}
}

func NewInteractiveRunner(interactorBinary []byte, checker problems.Checker, opts ...InteractiveRunnerOption) *InteractiveRunner {
	res := &InteractiveRunner{
		userBin:       nil,
		interactorBin: interactorBinary,
		checker:       checker,
		fs:            afero.NewOsFs(),
		executor:      PolygonUserInteractorExecute{},
	}
	for _, opt := range opts {
		opt(res)
	}
	return res
}

func (r *InteractiveRunner) SetSolution(ctx context.Context, solution problems.Solution) error {
	r.lang = solution.GetLanguage()
	return setSolution(ctx, &r.userBinName, &r.userBin, solution)
}

func (r *InteractiveRunner) getSandboxes(provider sandbox.Provider) (userSandbox, interactorSandbox sandbox.Sandbox, err error) {
	userSandbox, err = provider.Get()
	if err != nil {
		return
	}
	interactorSandbox, err = provider.Get()
	return
}

func (r *InteractiveRunner) prepareFIFO(dir string, name string) (*os.File, error) {
	if err := syscall.Mkfifo(filepath.Join(dir, name), 0666); err != nil {
		return nil, err
	}
	return os.OpenFile(filepath.Join(dir, name), os.O_RDWR, 0666)
}

type UserInteractorExecutor interface {
	ExecuteUser(ctx context.Context, userSandbox sandbox.Sandbox, language language.Language, userBin sandbox.File, userStdin, userStdout *os.File, timeLimit time.Duration, memoryLimit memory.Amount) (*sandbox.Status, error)
	ExecuteInteractor(ctx context.Context, interactorSandbox sandbox.Sandbox, userStdin, userStdout *os.File, testcase *problems.Testcase) (*sandbox.Status, error)
}

type PolygonUserInteractorExecute struct{}

func (p PolygonUserInteractorExecute) ExecuteUser(ctx context.Context, userSandbox sandbox.Sandbox, language language.Language, userBin sandbox.File, userStdin, userStdout *os.File, timeLimit time.Duration, memoryLimit memory.Amount) (*sandbox.Status, error) {
	return language.Run(ctx, userSandbox, userBin, userStdin, userStdout, timeLimit, memoryLimit)
}

func (p PolygonUserInteractorExecute) ExecuteInteractor(ctx context.Context, interactorSandbox sandbox.Sandbox, userStdin, userStdout *os.File, testcase *problems.Testcase) (*sandbox.Status, error) {
	return interactorSandbox.Run(ctx, sandbox.RunConfig{
		RunID:            "interactor",
		TimeLimit:        2 * testcase.TimeLimit,
		MemoryLimit:      1 * memory.GiB,
		Stdin:            userStdout,
		Stdout:           userStdin,
		Stderr:           io.Discard,
		InheritEnv:       true,
		WorkingDirectory: interactorSandbox.Pwd(),
	}, "interactor", "input", "output")
}

type interactorOutput struct {
	checkerMessage *bytes.Buffer
	scoreMul       *bytes.Buffer
}

type TaskYAMLUserInteractorExecute struct {
	forChecker sync.Map
}

func (t *TaskYAMLUserInteractorExecute) Check(ctx context.Context, testcase *problems.Testcase) error {
	if testcase.VerdictName != problems.VerdictAC {
		return nil
	}

	val, ok := t.forChecker.Load(testcase.Index)
	if !ok {
		return errors.New("index not found in forChecker")
	}
	res := val.(*interactorOutput)
	testcase.CheckerOutput = res.checkerMessage.String()

	mul := 0.0
	n, err := fmt.Fscanf(res.scoreMul, "%f", &mul)
	if err != nil {
		return err
	}
	if n < 1 {
		return errors.New("can't parse score multiplier")
	}
	testcase.Score = testcase.MaxScore * mul
	if mul == 0.0 {
		testcase.VerdictName = problems.VerdictWA
	} else if mul < 1.0 {
		testcase.VerdictName = problems.VerdictPC
	}
	return nil
}

func (t *TaskYAMLUserInteractorExecute) ExecuteUser(ctx context.Context, userSandbox sandbox.Sandbox, language language.Language, userBin sandbox.File, userStdin, userStdout *os.File, timeLimit time.Duration, memoryLimit memory.Amount) (*sandbox.Status, error) {
	return language.Run(ctx, userSandbox, userBin, userStdin, userStdout, timeLimit, memoryLimit)
}

func (t *TaskYAMLUserInteractorExecute) ExecuteInteractor(ctx context.Context, interactorSandbox sandbox.Sandbox, userStdin, userStdout *os.File, testcase *problems.Testcase) (*sandbox.Status, error) {
	inputFile, err := os.Open(filepath.Join(interactorSandbox.Pwd(), "input"))
	if err != nil {
		return nil, err
	}

	res := &interactorOutput{}
	res.checkerMessage = &bytes.Buffer{}
	res.scoreMul = &bytes.Buffer{}
	t.forChecker.Store(testcase.Index, res)

	return interactorSandbox.Run(ctx, sandbox.RunConfig{
		RunID:            "interactor",
		TimeLimit:        2 * testcase.TimeLimit,
		MemoryLimit:      1 * memory.GiB,
		Stdin:            inputFile,
		Stdout:           res.scoreMul,
		Stderr:           res.checkerMessage,
		InheritEnv:       true,
		WorkingDirectory: interactorSandbox.Pwd(),
		DirectoryMaps: []sandbox.DirectoryMap{
			{
				Inside:  filepath.Dir(userStdin.Name()),
				Outside: filepath.Dir(userStdin.Name()),
				Options: []sandbox.DirectoryMapOption{
					sandbox.AllowReadWrite,
					sandbox.NoExec,
					sandbox.Maybe,
				},
			},
		},
	}, "interactor", userStdin.Name(), userStdout.Name())
}

func (r *InteractiveRunner) Run(ctx context.Context, sandboxProvider sandbox.Provider, testcase *problems.Testcase) error {
	userSandbox, interactorSandbox, err := r.getSandboxes(sandboxProvider)
	if err != nil {
		return err
	}
	defer func() {
		_ = userSandbox.Cleanup(context.Background())
		_ = interactorSandbox.Cleanup(context.Background())
		sandboxProvider.Put(userSandbox)
		sandboxProvider.Put(interactorSandbox)
	}()
	if err = userSandbox.Init(ctx); err != nil {
		return err
	}
	if err = interactorSandbox.Init(ctx); err != nil {
		return err
	}

	if err = sandbox.CreateFile(interactorSandbox, sandbox.File{
		Name:   "interactor",
		Source: io.NopCloser(bytes.NewBuffer(r.interactorBin)),
	}); err != nil {
		return err
	}
	if err = interactorSandbox.MakeExecutable("interactor"); err != nil {
		return err
	}

	inputFile, err := r.fs.Open(testcase.InputPath)
	if err != nil {
		return err
	}
	if err = sandbox.CreateFile(interactorSandbox, sandbox.File{
		Name:   "input",
		Source: inputFile,
	}); err != nil {
		return err
	}

	dir, err := os.MkdirTemp("", "njudge_interactive_runner")
	if err != nil {
		return err
	}
	defer func(path string) {
		_ = os.RemoveAll(path)
	}(dir)

	userStdin, err := r.prepareFIFO(dir, "fifo1")
	if err != nil {
		return err
	}
	defer func(userStdin *os.File) {
		_ = userStdin.Close()
	}(userStdin)
	userStdout, err := r.prepareFIFO(dir, "fifo2")
	if err != nil {
		return err
	}
	defer func(userStdout *os.File) {
		_ = userStdout.Close()
	}(userStdout)

	var (
		userStatus, interactorStatus *sandbox.Status
		userError, interactorError   error
		done                         = make(chan struct{})
	)

	go func() {
		userStatus, userError = r.executor.ExecuteUser(ctx, userSandbox, r.lang, sandbox.File{
			Name:   r.userBinName,
			Source: io.NopCloser(bytes.NewBuffer(r.userBin)),
		}, userStdin, userStdout, testcase.TimeLimit, testcase.MemoryLimit)
		done <- struct{}{}
	}()

	interactorStatus, interactorError = r.executor.ExecuteInteractor(ctx, interactorSandbox, userStdin, userStdout, testcase)
	<-done

	testcase.OutputPath = filepath.Join(interactorSandbox.Pwd(), "output")
	testcase.TimeSpent = userStatus.Time
	testcase.MemoryUsed = userStatus.Memory

	if userError != nil || interactorError != nil {
		return errors.Join(userError, interactorError)
	}

	if interactorStatus.Verdict != sandbox.VerdictOK {
		testcase.VerdictName = problems.VerdictXX
		return fmt.Errorf("interactor didn't return ok: %v", interactorStatus)
	}

	if !mapVerdict(userStatus, testcase) || r.checker == nil {
		return nil
	}
	return r.checker.Check(ctx, testcase)
}
