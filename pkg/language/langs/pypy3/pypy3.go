package pypy3

import (
	"io"
	"time"

	"github.com/mraron/njudge/pkg/language"
)

type pypy3 struct{}

func (pypy3) Id() string {
	return "pypy3"
}

func (pypy3) Name() string {
	return "PyPy 3"
}

func (pypy3) DefaultFileName() string {
	return "main.py"
}

func (pypy3) InsecureCompile(wd string, r io.Reader, w io.Writer, e io.Writer) error {
	return nil
}

func (pypy3) Compile(s language.Sandbox, r language.File, w io.Writer, e io.Writer, extras []language.File) error {
	_, err := io.Copy(w, r.Source)
	return err
}

func (pypy3) Run(s language.Sandbox, binary, stdin io.Reader, stdout io.Writer, tl time.Duration, ml int) (language.Status, error) {
	stat := language.Status{}
	stat.Verdict = language.VerdictXX

	if err := s.CreateFile("a.out", binary); err != nil {
		return stat, err
	}

	if st, err := s.Env().TimeLimit(tl).MemoryLimit(ml/1024).Stdin(stdin).Stdout(stdout).WorkingDirectory(s.Pwd()).Run("/usr/bin/pypy3 a.out", true); err != nil {
		return st, err
	} else {
		stat = st
	}

	return stat, nil
}

func init() {
	language.DefaultStore.Register("pypy3", pypy3{})
}
