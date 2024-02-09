package golang

import (
	"github.com/mraron/njudge/pkg/language/langs/cpp"
	"github.com/mraron/njudge/pkg/language/memory"
	"github.com/mraron/njudge/pkg/language/sandbox"
	"golang.org/x/net/context"
	"io"
	"io/fs"
	"time"

	"github.com/mraron/njudge/pkg/language"
)

type golang struct{}

func (golang) Id() string {
	return "golang"
}

func (golang) Name() string {
	return "Go"
}

func (golang) DefaultFileName() string {
	return "main.go"
}

func (golang) Compile(s sandbox.Sandbox, r language.File, w io.Writer, e io.Writer, extras []language.File) error {
	err := sandbox.CreateFileFromSource(s, "main.go", r.Source)
	if err != nil {
		return err
	}

	rc := sandbox.RunConfig{
		MaxProcesses:     -1,
		InheritEnv:       true,
		Env:              []string{"GOCACHE=/tmp"},
		TimeLimit:        10 * time.Second,
		MemoryLimit:      1 * memory.GiB,
		Stdout:           e,
		Stderr:           e,
		WorkingDirectory: s.Pwd(),
		Args:             []string{"--open-files=2048"},
	}
	if _, err := s.Run(context.TODO(), rc, "/usr/bin/gccgo", "main.go"); err != nil {
		return err
	}

	bin, err := s.Open("a.out")
	if err != nil {
		return err
	}
	defer func(bin fs.File) {
		_ = bin.Close()
	}(bin)

	_, err = io.Copy(w, bin)

	return err
}

func (golang) Run(s sandbox.Sandbox, binary io.Reader, stdin io.Reader, stdout io.Writer, tl time.Duration, ml int) (*sandbox.Status, error) {
	return cpp.RunBinary("a.out")(s, binary, stdin, stdout, tl, ml)

}

func init() {
	language.DefaultStore.Register("golang", golang{})
}
