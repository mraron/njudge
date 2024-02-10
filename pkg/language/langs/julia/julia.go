package julia

import (
	"context"
	"github.com/mraron/njudge/pkg/language/memory"
	"github.com/mraron/njudge/pkg/language/sandbox"
	"io"
	"time"

	"github.com/mraron/njudge/pkg/language"
)

type julia struct{}

func (julia) ID() string {
	return "julia"
}

func (julia) DisplayName() string {
	return "Julia"
}
func (julia) DefaultFilename() string {
	return "main.jl"
}

func (julia) Compile(s sandbox.Sandbox, f sandbox.File, stderr io.Writer, extras []sandbox.File) (*sandbox.File, error) {
	return &f, nil
}

func (julia) Run(s sandbox.Sandbox, binary sandbox.File, stdin io.Reader, stdout io.Writer, tl time.Duration, ml memory.Amount) (*sandbox.Status, error) {
	stat := sandbox.Status{}
	stat.Verdict = sandbox.VerdictXX

	if err := sandbox.CreateFileFromSource(s, binary.Name, binary.Source); err != nil {
		return nil, err
	}

	rc := sandbox.RunConfig{
		InheritEnv:       true,
		MaxProcesses:     100,
		TimeLimit:        tl,
		MemoryLimit:      ml,
		Stdin:            stdin,
		Stdout:           stdout,
		WorkingDirectory: s.Pwd(),
	}

	return s.Run(context.TODO(), rc, "/usr/local/bin/julia", binary.Name)
}

func init() {
	language.DefaultStore.Register("julia", julia{})
}
