package python3

import (
	"context"
	"github.com/mraron/njudge/pkg/language/memory"
	"github.com/mraron/njudge/pkg/language/sandbox"
	"io"
	"time"

	"github.com/mraron/njudge/pkg/language"
)

type Python3 struct{}

func (Python3) ID() string {
	return "python3"
}

func (Python3) DisplayName() string {
	return "Python 3"
}

func (Python3) DefaultFilename() string {
	return "main.py"
}

func (Python3) Compile(ctx context.Context, s sandbox.Sandbox, f sandbox.File, stderr io.Writer, extras []sandbox.File) (*sandbox.File, error) {
	return &f, nil
}

func (Python3) Run(ctx context.Context, s sandbox.Sandbox, binary sandbox.File, stdin io.Reader, stdout io.Writer, tl time.Duration, ml memory.Amount) (*sandbox.Status, error) {
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
	return s.Run(ctx, rc, "/usr/bin/python3", binary.Name)
}

func init() {
	language.DefaultStore.Register("python3", Python3{})
}
