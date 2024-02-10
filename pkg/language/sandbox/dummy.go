package sandbox

import (
	"context"
	"errors"
	"io"
	"log/slog"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

var (
	DummyPattern = "dummy_sandbox*"

	dummyID = 0
)

type Dummy struct {
	ID  int
	Dir string

	OsFS

	Logger *slog.Logger

	inited bool
}

type DummyOption func(*Dummy) error

func NewDummy(opts ...DummyOption) (*Dummy, error) {
	dummyID += 1
	res := &Dummy{
		ID: dummyID,
	}
	for _, opt := range opts {
		err := opt(res)
		if err != nil {
			return nil, err
		}
	}
	if res.Logger == nil {
		res.Logger = slog.New(slog.NewJSONHandler(io.Discard, nil))
	}
	return res, nil
}

func (d *Dummy) Id() string {
	return strconv.Itoa(d.ID)
}

func (d *Dummy) Init(ctx context.Context) error {
	var err error
	if d.Dir, err = os.MkdirTemp("", DummyPattern); err != nil {
		return err
	}

	d.Logger.Info("init dummy sandbox", "dir", d.Dir)
	d.inited = true
	d.OsFS = NewOsFS(d.Dir)
	return nil
}

func (d *Dummy) Run(ctx context.Context, config RunConfig, command string, commandArgs ...string) (*Status, error) {
	if !d.inited {
		return nil, ErrorSandboxNotInitialized
	}

	logger := d.Logger
	if config.RunID != "" {
		logger = d.Logger.With("run_id", config.RunID)
	}

	started := time.Now()
	ctxWithTimeout, cancelFunc := context.WithTimeout(ctx, config.TimeLimit)
	defer cancelFunc()
	commandMerged := append([]string{command}, commandArgs...)
	cmd := exec.CommandContext(ctxWithTimeout, "bash", "-c", strings.Join(commandMerged, " "))
	cmd.Stdin = config.Stdin
	cmd.Stdout = config.Stdout
	cmd.Stderr = config.Stderr
	cmd.Dir = config.WorkingDirectory
	cmd.Env = append(config.Env, "PATH="+os.Getenv("PATH")+":"+d.Dir)

	if config.MaxProcesses > 0 {
		logger.Warn("dummy doesn't support MaxProcesses")
	}
	if !config.InheritEnv {
		logger.Warn("dummy always inherits environment")
	}
	if config.MemoryLimit > 0 {
		logger.Warn("dummy doesn't support memory limit")
	}
	if len(config.DirectoryMaps) > 0 {
		logger.Warn("dummy doesn't support directory mapping")
	}

	st := Status{
		Verdict: VerdictOK,
	}

	err := cmd.Run()
	st.Time = time.Since(started)
	if err != nil {
		if errors.Is(err, context.Canceled) {
			st.Verdict = VerdictTL
		} else if strings.HasPrefix(err.Error(), "exit status") || strings.HasPrefix(err.Error(), "signal:") { // TODO
			st.Verdict = VerdictRE
		}
		st.ExitCode = err.(*exec.ExitError).ExitCode()
	}

	return &st, nil
}

func (d *Dummy) Cleanup(ctx context.Context) error {
	d.Logger.Info("cleanup dummy sandbox", "dir", d.Dir)
	d.inited = false
	d.OsFS = OsFS{}
	return os.RemoveAll(d.Dir)

}

/*
func (s Dummy) Id() string {
	return s.tmpdir
}

func (s *Dummy) Init(logger *log.Logger) error {
	var err error
	if s.tmpdir, err = os.MkdirTemp("", "dummysandbox"); err != nil {
		return err
	}

	s.workingDir = s.tmpdir
	s.logger = logger
	return nil
}

func (s Dummy) Pwd() string {
	return s.tmpdir
}

func (s *Dummy) CreateFilePopulated(name string, r io.Reader) error {
	filename := filepath.Join(s.tmpdir, name)
	s.logger.Print("Creating file ", filename)

	f, err := os.Create(filename)
	if err != nil {
		s.logger.Print("Error occurred while creating file ", err)
		return err
	}

	if _, err := io.Copy(f, r); err != nil {
		s.logger.Print("Error occurred while populating it with its content: ", err)
		f.Close()
		return err
	}

	return f.Close()
}

func (s *Dummy) Create(name string) (io.WriteCloser, error) {
	return os.Create(filepath.Join(s.Pwd(), name))
}

func (s Dummy) Open(name string) (fs.File, error) {
	return os.Open(filepath.Join(s.Pwd(), name))
}

func (s Dummy) MakeExecutable(name string) error {
	filename := filepath.Join(s.Pwd(), name)

	err := os.Chmod(filename, 0777)
	s.logger.Print("Making executable: ", filename, " error: ", err)

	return err
}

func (s *Dummy) SetMaxProcesses(i int) Sandbox {
	return s
}

func (s *Dummy) Env() Sandbox {
	s.env = os.Environ()
	return s
}

func (s *Dummy) SetEnv(env string) Sandbox {
	s.env = append(s.env, env+"="+os.Getenv(env))
	return s
}

func (s *Dummy) AddArg(string) Sandbox {
	return s
}

func (s *Dummy) TimeLimit(tl time.Duration) Sandbox {
	s.tl = tl
	return s
}

func (s *Dummy) MemoryLimit(int) Sandbox {
	return s
}

func (s *Dummy) Stdin(reader io.Reader) Sandbox {
	s.stdin = reader
	return s
}

func (s *Dummy) Stderr(writer io.Writer) Sandbox {
	s.stderr = writer
	return s
}

func (s *Dummy) Stdout(writer io.Writer) Sandbox {
	s.stdout = writer
	return s
}

func (s *Dummy) MapDir(x string, y string, i []string, b bool) Sandbox {
	return s
}

func (s *Dummy) WorkingDirectory(dir string) Sandbox {
	s.workingDir = dir
	return s
}

func (s *Dummy) Verbose() Sandbox {
	return s
}

func (s *Dummy) Run(prg string, needStatus bool) (Status, error) {
	cmd := exec.Command("bash", "-c", prg)
	cmd.Stdin = s.stdin
	cmd.Stdout = s.stdout
	cmd.Stderr = s.stderr
	cmd.Dir = s.workingDir
	cmd.Env = append(s.env, "PATH="+os.Getenv("PATH")+":"+s.tmpdir)

	var (
		st               Status
		errKill, errWait error
		finish           = make(chan bool, 1)
		wg               sync.WaitGroup
	)

	st.Verdict = VerdictOK

	start := time.NewTimer(s.tl)
	if err := cmd.Start(); err != nil {
		st.Verdict = VerdictXX
		return st, err
	}
	defer start.Stop()

	wg.Add(1)
	go func() {
		defer wg.Done()
		errWait = cmd.Wait()
		finish <- true
	}()

	select {
	case <-start.C:
		st.Verdict = VerdictTL
		if errKill = cmd.Process.Kill(); errKill != nil {
			st.Verdict = VerdictXX
		}
	case <-finish:
	}

	wg.Wait()

	if errWait != nil && (strings.HasPrefix(errWait.Error(), "exit status") || strings.HasPrefix(errWait.Error(), "signal:")) {
		if st.Verdict == VerdictOK {
			st.Verdict = VerdictRE
		}
		errWait = nil
	}

	if errWait != nil {
		return st, errWait
	}

	return st, errKill
}

func (s *Dummy) Cleanup() error {
	return os.RemoveAll(s.tmpdir)
}
*/
