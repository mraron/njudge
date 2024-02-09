package cython3

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

type cython3 struct {
}

func (c cython3) Id() string {
	return "cython3"
}

func (c cython3) Name() string {
	return "Cython3"
}

func (c cython3) DefaultFileName() string {
	return "main.py"
}

func (c cython3) Compile(s sandbox.Sandbox, r language.File, w io.Writer, e io.Writer, extras []language.File) error {
	err := sandbox.CreateFileFromSource(s, "main.py", r.Source)
	if err != nil {
		return err
	}

	rc := sandbox.RunConfig{
		MaxProcesses:     200,
		InheritEnv:       true,
		TimeLimit:        10 * time.Second,
		MemoryLimit:      256 * memory.MiB,
		Stdout:           e,
		Stderr:           e,
		WorkingDirectory: s.Pwd(),
	}
	if _, err := s.Run(context.TODO(), rc, "/usr/bin/cython3", sandbox.SplitArgs("-3 --embed -o main.c main.py")...); err != nil {
		return err
	}

	if _, err := s.Run(context.TODO(), rc, "/usr/bin/gcc", sandbox.SplitArgs("-O2 -I/usr/include/python3.8 -o main main.c -lpython3.8 -lpthread -lm -lutil -ldl")...); err != nil {
		return err
	}

	bin, err := s.Open("main")
	if err != nil {
		return err
	}
	defer func(bin fs.File) {
		_ = bin.Close()
	}(bin)

	_, err = io.Copy(w, bin)
	return err
}

func (cython3) Run(s sandbox.Sandbox, binary io.Reader, stdin io.Reader, stdout io.Writer, tl time.Duration, ml int) (*sandbox.Status, error) {
	return cpp.RunBinary("a.out")(s, binary, stdin, stdout, tl, ml)

}

func init() {
	language.DefaultStore.Register("cython3", cython3{})
}
