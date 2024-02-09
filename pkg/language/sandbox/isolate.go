package sandbox

import (
	"bufio"
	"fmt"
	"github.com/mraron/njudge/pkg/language/memory"
	"golang.org/x/net/context"
	"io"
	"log/slog"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

var IsolateRoot = getEnv("ISOLATE_ROOT", "/var/local/lib/isolate/")
var IsolateMetafilePattern = "isolate_metafile"

type Isolate struct {
	ID int

	OsFS

	Logger *slog.Logger
}

type IsolateOption func(*Isolate) error

func IsolateOptionUseLogger(logger *slog.Logger) IsolateOption {
	return func(isolate *Isolate) error {
		isolate.Logger = logger
		return nil
	}
}

func NewIsolate(ID int, opts ...IsolateOption) (*Isolate, error) {
	res := &Isolate{
		ID: ID,
		OsFS: OsFS{
			base: filepath.Join(IsolateRoot, strconv.Itoa(ID), "box"),
		},
	}
	for _, opt := range opts {
		if err := opt(res); err != nil {
			return nil, err
		}
	}
	if res.Logger == nil {
		res.Logger = slog.New(slog.NewJSONHandler(io.Discard, nil))
	}
	res.Logger = res.Logger.With(slog.String("sandbox", res.Id()))
	return res, nil
}

func (i Isolate) Id() string {
	return "isolate" + strconv.Itoa(i.ID)
}

func (i Isolate) Init(ctx context.Context) error {
	// cleanup because the previous invocation might not have cleaned up
	if err := i.Cleanup(ctx); err != nil {
		return err
	}

	cmd := []string{"isolate", "--cg", "-b", strconv.Itoa(i.ID), "--init"}
	i.Logger.Info("running init", "cmd", cmd)
	return exec.CommandContext(ctx, cmd[0], cmd[1:]...).Run()
}

func (i Isolate) buildArgs(config RunConfig) ([]string, error) {
	args := []string{"isolate", "--cg", "-b", strconv.Itoa(i.ID)}
	if config.MaxProcesses > 0 {
		args = append(args, fmt.Sprintf("--processes=%d", config.MaxProcesses))
	} else {
		args = append(args, "--processes=100")
	}
	if config.InheritEnv {
		args = append(args, "--full-env")
	}
	for ind := range config.Env {
		args = append(args, fmt.Sprintf("--env=%s", config.Env[ind]))
	}
	for _, rule := range config.DirectoryMaps {
		arg := fmt.Sprintf("--dir=%s=%s", rule.Inside, rule.Outside)
		for _, opt := range rule.Options {
			arg += ":" + string(opt)
		}
		args = append(args, arg)
	}
	if config.TimeLimit > 0 {
		ms := config.TimeLimit / time.Millisecond
		args = append(args, fmt.Sprintf("--time=%d.%d", ms/1000, ms%1000))
		args = append(args, fmt.Sprintf("--wall-time=%d.%d", (2*ms+1000)/1000, (2*ms+1000)%1000)) // TODO?
	}
	if config.MemoryLimit > 0 {
		args = append(args, "--cg-mem="+strconv.Itoa(int(config.MemoryLimit/memory.KiB)))
	}
	for _, arg := range config.Args {
		args = append(args, arg)
	}

	return args, nil
}

func (i Isolate) Run(ctx context.Context, config RunConfig, toRun string, toRunArgs ...string) (*Status, error) {
	logger := i.Logger
	if config.RunID != "" {
		logger = i.Logger.With("run_id", config.RunID)
	}

	args, err := i.buildArgs(config)
	if err != nil {
		return nil, fmt.Errorf("failed to build isolate command: %w", err)
	}
	metafile, err := os.CreateTemp(os.TempDir(), IsolateMetafilePattern)
	if err != nil {
		return nil, fmt.Errorf("failed to create metafile: %w", err)
	}
	defer func(metafile *os.File) {
		_ = metafile.Close()
		_ = os.Remove(filepath.Join(os.TempDir(), metafile.Name()))
	}(metafile)

	args = append(args, fmt.Sprintf("--meta=%s", metafile.Name()))

	args = append(args, "--run", "-s", "--", toRun)
	args = append(args, toRunArgs...)

	logger.Info("built args", "args", args)

	cmd := exec.CommandContext(ctx, args[0], args[1:]...)
	cmd.Stdin = config.Stdin
	cmd.Stdout = config.Stdout
	cmd.Stderr = config.Stderr
	cmd.Dir = config.WorkingDirectory
	_ = cmd.Run()

	st := Status{
		Verdict: VerdictOK,
	}
	sc := bufio.NewScanner(metafile)
	for sc.Scan() {
		lst := strings.Split(sc.Text(), ":")
		switch lst[0] {
		case "max-rss":
			fallthrough
		case "cg-mem":
			mem, _ := strconv.Atoi(lst[1])
			st.Memory += mem
		case "time":
			tmp, _ := strconv.ParseFloat(lst[1], 32)
			st.Time = time.Duration(tmp*1000) * time.Millisecond
		case "status":
			switch lst[1] {
			case "TO":
				st.Verdict = VerdictTL
			case "RE":
				st.Verdict = VerdictRE
			case "SG":
				st.Verdict = VerdictRE
			case "XX":
				st.Verdict = VerdictXX
			}
		case "exitcode":
			st.ExitCode, _ = strconv.Atoi(lst[1])
		}
	}
	if err = sc.Err(); err != nil {
		return nil, fmt.Errorf("failed to scan metafile: %w", err)
	}

	return &st, nil
}

func (i Isolate) Cleanup(ctx context.Context) error {
	cmd := []string{"isolate", "--cg", "-b", strconv.Itoa(i.ID), "--cleanup"}

	i.Logger.Info("running cleanup ", "cmd", cmd)
	return exec.CommandContext(ctx, cmd[0], cmd[1:]...).Run()
}
