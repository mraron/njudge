// Package sandbox declares the Sandbox interface which can be used to run untrusted code. Hopefully in a secure way.
// It has two implementations built-in:
//   - Isolate: which uses isolate for sandboxing.
//   - Dummy: which is used for testing and not really secure.
package sandbox

import (
	"bytes"
	"context"
	"errors"
	"github.com/mraron/njudge/pkg/language/memory"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"syscall"
	"time"
)

// ErrorSandboxNotInitialized should returned when a Run is called on a Sandbox without a prior Init call.
var ErrorSandboxNotInitialized = errors.New("initialize the sandbox first")

// FS is a file system abstraction for sandboxes.
type FS interface {
	Pwd() string
	Create(name string) (io.WriteCloser, error)
	MakeExecutable(name string) error

	fs.FS
}

// CreateFileFromSource is a convenience method for creating a file inside a sandbox with the given content.
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

// ExtractFile is a convenience method for getting the content of a file from inside the sandbox.
func ExtractFile(s FS, name string) (*File, error) {
	bin, err := s.Open(name)
	if err != nil {
		return nil, err
	}
	defer func(bin fs.File) {
		_ = bin.Close()
	}(bin)

	binaryContent, err := io.ReadAll(bin)
	if err != nil {
		return nil, err
	}

	return &File{
		Name:   name,
		Source: bytes.NewBuffer(binaryContent),
	}, nil
}

// RunBinary is a convenience method for running a binary (which needs no special configuration) inside the given Sandbox.
func RunBinary(ctx context.Context, s Sandbox, binary File, stdin io.Reader, stdout io.Writer, tl time.Duration, ml memory.Amount) (*Status, error) {
	stat := Status{}
	stat.Verdict = VerdictXX

	if err := CreateFileFromSource(s, binary.Name, binary.Source); err != nil {
		return nil, err
	}

	if err := s.MakeExecutable(binary.Name); err != nil {
		return nil, err
	}

	rc := RunConfig{
		Stdin:       stdin,
		Stdout:      stdout,
		TimeLimit:   tl,
		MemoryLimit: ml,
	}
	return s.Run(ctx, rc, binary.Name)
}

// SplitArgs is used to split a program's arguments for Sandbox.Run
// It just calls strings.Split, but it's useful for context.
func SplitArgs(s string) []string {
	return strings.Split(s, " ")
}

// DirectoryMap is used to map a directory from outside to the inside of a Sandbox.
type DirectoryMap struct {
	Inside  string
	Outside string
	Options []DirectoryMapOption
}

// DirectoryMapOption is an option for a DirectoryMap.
// The values are defined as in isolate. You may refer to isolate's help for more information on them.
type DirectoryMapOption string

const (
	AllowSpecial   DirectoryMapOption = "dev"
	MountFS        DirectoryMapOption = "fs"
	Maybe          DirectoryMapOption = "maybe"
	NoExec         DirectoryMapOption = "noexec"
	AllowReadWrite DirectoryMapOption = "rw"
	Temporary      DirectoryMapOption = "tmp"
)

// RunConfig is used to configure the sandbox, setting limits and streams (stdin, stdout, stderr).
type RunConfig struct {
	RunID string // RunID is some kind of ID of the run, doesn't need to be set but if set is used to give additional context.

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

	Args []string // Args are given to the underlying sandbox implementation as-is. This should only be used for things that are currently not supported in RunConfig.
}

// Sandbox is used to Run a command inside a secure sandbox.
type Sandbox interface {
	Id() string // Id should return an unique ID for sandboxes of the same type

	Init(ctx context.Context) error
	FS
	Run(ctx context.Context, config RunConfig, command string, commandArgs ...string) (*Status, error)
	Cleanup(ctx context.Context) error
}

// Provider can be used to store Sandboxes
type Provider interface {
	Get() (Sandbox, error)
	Put(s Sandbox)
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

func (sp *ChanProvider) Put(s Sandbox) {
	sp.sandboxes <- s
}

// File is a named io.Reader which emulates a file.
type File struct {
	Name   string
	Source io.Reader
}
