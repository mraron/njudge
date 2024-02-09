package python3

import (
	"github.com/mraron/njudge/pkg/language/memory"
	"github.com/mraron/njudge/pkg/language/sandbox"
	"golang.org/x/net/context"
	"io"
	"time"

	"github.com/mraron/njudge/pkg/language"
)

type python3 struct{}

func (python3) Id() string {
	return "python3"
}

func (python3) Name() string {
	return "Python 3"
}

func (python3) DefaultFileName() string {
	return "main.py"
}

func (python3) Compile(s sandbox.Sandbox, r language.File, w io.Writer, e io.Writer, extras []language.File) error {
	_, err := io.Copy(w, r.Source)
	return err
}

func (python3) Run(s sandbox.Sandbox, binary io.Reader, stdin io.Reader, stdout io.Writer, tl time.Duration, ml int) (*sandbox.Status, error) {
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
	return s.Run(context.TODO(), rc, "/usr/bin/python3", "a.out")
}

func init() {
	language.DefaultStore.Register("python3", python3{})
}
