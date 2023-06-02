package language

import (
	"errors"
	"io"
	"log"
	"time"
)

type Sandbox interface {
	Id() string
	Init(*log.Logger) error

	Pwd() string

	CreateFile(string, io.Reader) error
	GetFile(string) (io.Reader, error)
	MakeExecutable(string) error

	SetMaxProcesses(int) Sandbox
	Env() Sandbox
	SetEnv(string) Sandbox
	AddArg(string) Sandbox
	TimeLimit(time.Duration) Sandbox
	MemoryLimit(int) Sandbox
	Stdin(io.Reader) Sandbox
	Stderr(io.Writer) Sandbox
	Stdout(io.Writer) Sandbox
	MapDir(string, string, []string, bool) Sandbox
	WorkingDirectory(string) Sandbox
	Verbose() Sandbox
	Run(string, bool) (Status, error)

	Cleanup() error
}

type SandboxProvider struct {
	sandboxes chan Sandbox
}

func NewSandboxProvider() *SandboxProvider {
	return &SandboxProvider{make(chan Sandbox, 100)}
}

func (sp *SandboxProvider) Get() (Sandbox, error) {
	if len(sp.sandboxes) == 0 {
		return nil, errors.New("no sandbox available")
	}

	return <-sp.sandboxes, nil
}

func (sp *SandboxProvider) MustGet() Sandbox {
	if s, err := sp.Get(); err != nil {
		panic(err)
	} else {
		return s
	}
}

func (sp *SandboxProvider) Put(s Sandbox) {
	sp.sandboxes <- s
}
