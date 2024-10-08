package cython3

import (
	"context"
	"github.com/mraron/njudge/pkg/language/memory"
	"github.com/mraron/njudge/pkg/language/sandbox"
	"io"
	"time"

	"github.com/mraron/njudge/pkg/language"
)

type Cython3 struct {
}

func (c Cython3) ID() string {
	return "cython3"
}

func (c Cython3) DisplayName() string {
	return "Cython3"
}

func (c Cython3) DefaultFilename() string {
	return "main.py"
}

func (c Cython3) Compile(ctx context.Context, s sandbox.Sandbox, f sandbox.File, stderr io.Writer, extras []sandbox.File) (*sandbox.File, error) {
	err := sandbox.CreateFile(s, f)
	if err != nil {
		return nil, err
	}

	rc := sandbox.RunConfig{
		MaxProcesses:     200,
		InheritEnv:       true,
		TimeLimit:        10 * time.Second,
		MemoryLimit:      256 * memory.MiB,
		Stdout:           stderr,
		Stderr:           stderr,
		WorkingDirectory: s.Pwd(),
	}
	if _, err := s.Run(ctx, rc, "/usr/bin/cython3", sandbox.SplitArgs("-3 --embed -o main.c "+f.Name)...); err != nil {
		return nil, err
	}

	if _, err := s.Run(ctx, rc, "/usr/bin/gcc", sandbox.SplitArgs("-O2 -I/usr/include/python3.8 main.c -lpython3.8 -lpthread -lm -lutil -ldl")...); err != nil {
		return nil, err
	}

	return sandbox.ExtractFile(s, "a.out")
}

func (Cython3) Run(ctx context.Context, s sandbox.Sandbox, binary sandbox.File, stdin io.Reader, stdout io.Writer, tl time.Duration, ml memory.Amount) (*sandbox.Status, error) {
	return sandbox.RunBinary(ctx, s, binary, stdin, stdout, tl, ml)

}

func init() {
	language.DefaultStore.Register("cython3", Cython3{})
}
