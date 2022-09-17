package julia

import (
	"io"
	"os"
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
func (julia) InsecureCompile(wd string, r io.Reader, w io.Writer, e io.Writer) error {
	return nil
}

func (julia) Compile(s language.Sandbox, r language.File, w io.Writer, e io.Writer, extras []language.File) error {
	_, err := io.Copy(w, r.Source)
	return err
}

func (julia) Run(s language.Sandbox, binary, stdin io.Reader, stdout io.Writer, tl time.Duration, ml int) (language.Status, error) {
	stat := language.Status{}
	stat.Verdict = language.VERDICT_XX

	if err := s.CreateFile("a.out", binary); err != nil {
		return stat, err
	}

	if st, err := s.Env().SetMaxProcesses(100).TimeLimit(tl).MemoryLimit(ml/1024).Stdin(stdin).Stdout(stdout).Stderr(os.Stderr).WorkingDirectory(s.Pwd()).Run("/usr/local/bin/julia a.out", true); err != nil {
		return st, err
	} else {
		stat = st
	}

	return stat, nil
}

func init() {
	language.Register("julia", julia{})
}
