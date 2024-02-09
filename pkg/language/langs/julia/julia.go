package julia

import (
	"github.com/mraron/njudge/pkg/language/memory"
	"github.com/mraron/njudge/pkg/language/sandbox"
	"golang.org/x/net/context"
	"io"
	"time"

	"github.com/mraron/njudge/pkg/language"
)

type julia struct{}

func (julia) Id() string {
	return "julia"
}

func (julia) Name() string {
	return "Julia"
}
func (julia) DefaultFileName() string {
	return "main.jl"
}

func (julia) Compile(s sandbox.Sandbox, r language.File, w io.Writer, e io.Writer, extras []language.File) error {
	_, err := io.Copy(w, r.Source)
	return err
}

func (julia) Run(s sandbox.Sandbox, binary io.Reader, stdin io.Reader, stdout io.Writer, tl time.Duration, ml int) (*sandbox.Status, error) {
	stat := sandbox.Status{}
	stat.Verdict = sandbox.VerdictXX

	if err := sandbox.CreateFileFromSource(s, "a.out", binary); err != nil {
		return nil, err
	}

	rc := sandbox.RunConfig{
		InheritEnv:       true,
		MaxProcesses:     100,
		TimeLimit:        tl,
		MemoryLimit:      memory.Amount(ml) * memory.KiB,
		Stdin:            stdin,
		Stdout:           stdout,
		WorkingDirectory: s.Pwd(),
	}

	return s.Run(context.TODO(), rc, "/usr/local/bin/julia", "a.out")
}

func init() {
	language.DefaultStore.Register("julia", julia{})
}
