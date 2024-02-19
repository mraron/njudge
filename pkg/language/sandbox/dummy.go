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

// DummyPattern is the pattern in which Dummy creates its temporary directory.
const DummyPattern = "dummy_sandbox*"

var (
	dummyID = 0
)

// Dummy is a very straightforward implementation of a Sandbox.
// It creates a temporary directory and executes the commands without many precautions.
type Dummy struct {
	ID  int
	Dir string

	OsFS

	Logger *slog.Logger

	inited bool
}

type DummyOption func(*Dummy) error

func DummyWithLogger(logger *slog.Logger) DummyOption {
	return func(dummy *Dummy) error {
		dummy.Logger = logger
		return nil
	}
}

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

func (d *Dummy) Init(_ context.Context) error {
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
	if config.TimeLimit > 0 {
		ctxWithTimeout, cancelFunc := context.WithTimeout(ctx, config.TimeLimit)
		defer cancelFunc()
		ctx = ctxWithTimeout
	}
	commandMerged := append([]string{command}, commandArgs...)
	cmd := exec.CommandContext(ctx, "bash", "-c", strings.Join(commandMerged, " "))
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
		if errors.Is(err, context.DeadlineExceeded) { //TODO validate
			st.Verdict = VerdictTL
		} else if strings.HasPrefix(err.Error(), "exit status") || strings.HasPrefix(err.Error(), "signal:") { // TODO
			st.Verdict = VerdictRE
			st.ExitCode = err.(*exec.ExitError).ExitCode()
		}
	}

	return &st, nil
}

func (d *Dummy) Cleanup(_ context.Context) error {
	d.Logger.Info("cleanup dummy sandbox", "dir", d.Dir)
	d.inited = false
	d.OsFS = OsFS{}
	return os.RemoveAll(d.Dir)

}
