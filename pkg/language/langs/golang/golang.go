package golang

import (
	"context"
	"github.com/mraron/njudge/pkg/language/memory"
	"github.com/mraron/njudge/pkg/language/sandbox"
	"io"
	"time"

	"github.com/mraron/njudge/pkg/language"
)

type Golang struct{}

func (Golang) ID() string {
	return "golang"
}

func (Golang) DisplayName() string {
	return "Go"
}

func (Golang) DefaultFilename() string {
	return "main.go"
}

func (Golang) Compile(ctx context.Context, s sandbox.Sandbox, f sandbox.File, stderr io.Writer, extras []sandbox.File) (*sandbox.File, error) {
	err := sandbox.CreateFile(s, f)
	if err != nil {
		return nil, err
	}

	rc := sandbox.RunConfig{
		MaxProcesses:     -1,
		InheritEnv:       true,
		Env:              []string{"GOCACHE=/tmp"},
		TimeLimit:        10 * time.Second,
		MemoryLimit:      1 * memory.GiB,
		Stdout:           stderr,
		Stderr:           stderr,
		WorkingDirectory: s.Pwd(),
		Args:             []string{"--open-files=2048"},
	}
	if _, err := s.Run(ctx, rc, "/usr/bin/gccgo", f.Name); err != nil {
		return nil, err
	}

	return sandbox.ExtractFile(s, "a.out")
}

func (Golang) Run(ctx context.Context, s sandbox.Sandbox, binary sandbox.File, stdin io.Reader, stdout io.Writer, tl time.Duration, ml memory.Amount) (*sandbox.Status, error) {
	return sandbox.RunBinary(ctx, s, binary, stdin, stdout, tl, ml)
}

func init() {
	language.DefaultStore.Register("golang", Golang{})
}
