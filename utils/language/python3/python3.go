package python3

import (
	"github.com/mraron/njudge/utils/language"
	"io"
	"os"
	"time"
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

func (python3) InsecureCompile(wd string, r io.Reader, w io.Writer, e io.Writer) error {
	return nil
}

func (python3) Compile(s language.Sandbox, r language.File, w io.Writer, e io.Writer, extras []language.File) error {
	_, err := io.Copy(w, r.Source)
	return err
}

func (python3) Run(s language.Sandbox, binary, stdin io.Reader, stdout io.Writer, tl time.Duration, ml int) (language.Status, error) {
	stat := language.Status{}
	stat.Verdict = language.VERDICT_XX

	if err := s.CreateFile("a.out", binary); err != nil {
		return stat, err
	}

	if st, err := s.Env().TimeLimit(tl).MemoryLimit(ml/1024).Stdin(stdin).Stdout(stdout).Stderr(os.Stderr).WorkingDirectory(s.Pwd()).Run("/usr/bin/python3 a.out", true); err != nil {
		return st, err
	} else {
		stat = st
	}

	return stat, nil
}

func init() {
	language.Register("python3", python3{})
}
