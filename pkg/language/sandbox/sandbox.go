package sandbox

import (
	"errors"
	"github.com/mraron/njudge/pkg/language/memory"
	"golang.org/x/net/context"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"syscall"
	"time"
)

type FS interface {
	Pwd() string
	Create(name string) (io.WriteCloser, error)
	MakeExecutable(name string) error

	fs.FS
}

func CreateFileFromSource(fs FS, name string, source io.Reader) error {
	if err := syscall.Unlink(filepath.Join(fs.Pwd(), name)); err != nil && !errors.Is(err, os.ErrNotExist) {
		return err
	}

	f, err := fs.Create(name)
	if err != nil {
		return err
	}
	defer func(f io.WriteCloser) {
		_ = f.Close()
	}(f)

	if _, err = io.Copy(f, source); err != nil {
		return err
	}
	return nil
}

func SplitArgs(s string) []string {
	return strings.Split(s, " ")
}

type DirectoryMapOption string

var (
	AllowSpecial   DirectoryMapOption = "dev"
	MountFS        DirectoryMapOption = "fs"
	Maybe          DirectoryMapOption = "maybe"
	NoExec         DirectoryMapOption = "noexec"
	AllowReadWrite DirectoryMapOption = "rw"
	Temporary      DirectoryMapOption = "tmp"
)

type DirectoryMap struct {
	Inside  string
	Outside string
	Options []DirectoryMapOption
}

type RunConfig struct {
	RunID string

	MaxProcesses     int
	InheritEnv       bool
	Env              []string
	WorkingDirectory string
	DirectoryMaps    []DirectoryMap

	TimeLimit   time.Duration
	MemoryLimit memory.Amount

	Stdin  io.Reader
	Stderr io.Writer
	Stdout io.Writer

	Args []string
}

func (rc *RunConfig) SetMaxProcesses(i int) *RunConfig {
	rc.MaxProcesses = i
	return rc
}

func (rc *RunConfig) SetInheritEnv() *RunConfig {
	rc.InheritEnv = true
	return rc
}

func (rc *RunConfig) SetEnv(env string) *RunConfig {
	rc.Env = append(rc.Env, env+"="+os.Getenv(env))
	return rc
}

func (rc *RunConfig) SetTimeLimit(tl time.Duration) *RunConfig {
	rc.TimeLimit = tl
	return rc
}

func (rc *RunConfig) SetMemoryLimit(amount memory.Amount) *RunConfig {
	rc.MemoryLimit = amount
	return rc
}

func (rc *RunConfig) SetStdin(reader io.Reader) *RunConfig {
	rc.Stdin = reader
	return rc
}

func (rc *RunConfig) SetStderr(writer io.Writer) *RunConfig {
	rc.Stderr = writer
	return rc
}

func (rc *RunConfig) SetStdout(writer io.Writer) *RunConfig {
	rc.Stdout = writer
	return rc
}

func (rc *RunConfig) MapDir(directoryMap DirectoryMap) *RunConfig {
	rc.DirectoryMaps = append(rc.DirectoryMaps, directoryMap)
	return rc
}

func (rc *RunConfig) SetWorkingDirectory(dir string) *RunConfig {
	rc.WorkingDirectory = dir
	return rc
}

type Sandbox interface {
	Id() string

	Init(ctx context.Context) error
	FS
	Run(ctx context.Context, config RunConfig, command string, commandArgs ...string) (*Status, error)
	Cleanup(ctx context.Context) error
}

type ChanProvider struct {
	sandboxes chan Sandbox
}

func NewSandboxProvider() *ChanProvider {
	return &ChanProvider{make(chan Sandbox, 100)}
}

func (sp *ChanProvider) Get() (Sandbox, error) {
	return <-sp.sandboxes, nil
}

func (sp *ChanProvider) MustGet() Sandbox {
	if s, err := sp.Get(); err != nil {
		panic(err)
	} else {
		return s
	}
}

func (sp *ChanProvider) Put(s Sandbox) {
	sp.sandboxes <- s
}

type Provider interface {
	Get() (Sandbox, error)
	Put(s Sandbox)
}
