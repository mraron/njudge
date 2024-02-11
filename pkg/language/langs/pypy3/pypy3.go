package pypy3

import (
	"context"
	"github.com/mraron/njudge/pkg/language/memory"
	"github.com/mraron/njudge/pkg/language/sandbox"
	"io"
	"time"

	"github.com/mraron/njudge/pkg/language"
)

type PyPy3 struct{}

func (PyPy3) ID() string {
	return "pypy3"
}

func (PyPy3) DisplayName() string {
	return "PyPy 3"
}

func (PyPy3) DefaultFilename() string {
	return "main.py"
}

func (PyPy3) Compile(ctx context.Context, s sandbox.Sandbox, f sandbox.File, stderr io.Writer, extras []sandbox.File) (*sandbox.File, error) {
	return &f, nil
}

func (PyPy3) Run(ctx context.Context, s sandbox.Sandbox, binary sandbox.File, stdin io.Reader, stdout io.Writer, tl time.Duration, ml memory.Amount) (*sandbox.Status, error) {
	stat := sandbox.Status{}
	stat.Verdict = sandbox.VerdictXX

	if err := sandbox.CreateFile(s, binary); err != nil {
		return nil, err
	}

	rc := sandbox.RunConfig{
		InheritEnv:       true,
		TimeLimit:        tl,
		MemoryLimit:      ml,
		Stdin:            stdin,
		Stdout:           stdout,
		WorkingDirectory: s.Pwd(),
	}
	return s.Run(ctx, rc, "/usr/bin/pypy3", binary.Name)
}

func init() {
	language.DefaultStore.Register("pypy3", PyPy3{})
}
