package pypy3

import (
	"context"
	"github.com/mraron/njudge/pkg/language/memory"
	"github.com/mraron/njudge/pkg/language/sandbox"
	"io"
	"time"

	"github.com/mraron/njudge/pkg/language"
)

type pypy3 struct{}

func (pypy3) ID() string {
	return "pypy3"
}

func (pypy3) DisplayName() string {
	return "PyPy 3"
}

func (pypy3) DefaultFilename() string {
	return "main.py"
}

func (pypy3) Compile(s sandbox.Sandbox, f sandbox.File, stderr io.Writer, extras []sandbox.File) (*sandbox.File, error) {
	return &f, nil
}

func (pypy3) Run(s sandbox.Sandbox, binary sandbox.File, stdin io.Reader, stdout io.Writer, tl time.Duration, ml memory.Amount) (*sandbox.Status, error) {
	stat := sandbox.Status{}
	stat.Verdict = sandbox.VerdictXX

	if err := sandbox.CreateFileFromSource(s, binary.Name, binary.Source); err != nil {
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
	return s.Run(context.TODO(), rc, "/usr/bin/pypy3", binary.Name)
}

func init() {
	language.DefaultStore.Register("pypy3", pypy3{})
}
