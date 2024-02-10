package python3

import (
	"context"
	"github.com/mraron/njudge/pkg/language/memory"
	"github.com/mraron/njudge/pkg/language/sandbox"
	"io"
	"time"

	"github.com/mraron/njudge/pkg/language"
)

type python3 struct{}

func (python3) ID() string {
	return "python3"
}

func (python3) DisplayName() string {
	return "Python 3"
}

func (python3) DefaultFilename() string {
	return "main.py"
}

func (python3) Compile(s sandbox.Sandbox, f sandbox.File, stderr io.Writer, extras []sandbox.File) (*sandbox.File, error) {
	return &f, nil
}

func (python3) Run(s sandbox.Sandbox, binary sandbox.File, stdin io.Reader, stdout io.Writer, tl time.Duration, ml memory.Amount) (*sandbox.Status, error) {
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
	return s.Run(context.TODO(), rc, "/usr/bin/python3", binary.Name)
}

func init() {
	language.DefaultStore.Register("python3", python3{})
}
