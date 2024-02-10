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

func (pypy3) Id() string {
	return "pypy3"
}

func (pypy3) DisplayName() string {
	return "PyPy 3"
}

func (pypy3) DefaultFilename() string {
	return "main.py"
}

func (pypy3) Compile(s sandbox.Sandbox, r language.File, w io.Writer, e io.Writer, extras []language.File) error {
	_, err := io.Copy(w, r.Source)
	return err
}

func (pypy3) Run(s sandbox.Sandbox, binary io.Reader, stdin io.Reader, stdout io.Writer, tl time.Duration, ml memory.Amount) (*sandbox.Status, error) {
	stat := sandbox.Status{}
	stat.Verdict = sandbox.VerdictXX

	if err := sandbox.CreateFileFromSource(s, "a.out", binary); err != nil {
		return nil, err
	}

	rc := sandbox.RunConfig{
		InheritEnv:       true,
		TimeLimit:        tl,
		MemoryLimit:      memory.Amount(ml) * memory.KiB,
		Stdin:            stdin,
		Stdout:           stdout,
		WorkingDirectory: s.Pwd(),
	}
	return s.Run(context.TODO(), rc, "/usr/bin/pypy3", "a.out")
}

func init() {
	language.DefaultStore.Register("pypy3", pypy3{})
}
