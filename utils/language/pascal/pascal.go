package pascal

import (
	"errors"
	"github.com/mraron/njudge/utils/language"
	"io"
	"time"
)

type pascal struct{}

func (pascal) Id() string {
	return "pascal"
}

func (pascal) Name() string {
	return "Pascal"
}

func (pascal) DefaultFileName() string {
	return "main.pas"
}

func (pascal) InsecureCompile(wd string, r io.Reader, w io.Writer, e io.Writer) error {
	return errors.New("unsupported operation")
}

func (pascal) Compile(s language.Sandbox, r language.File, w io.Writer, e io.Writer, extras []language.File) error {
	err := s.CreateFile("main.pas", r.Source)
	if err != nil {
		return err
	}

	if _, err := s.SetMaxProcesses(-1).Env().TimeLimit(10*time.Second).MemoryLimit(256000).Stdout(e).Stderr(e).WorkingDirectory(s.Pwd()).MapDir("/etc", "/etc", []string{"noexec"}, false).Run("/usr/bin/fpc -Mobjfpc -O2 -Xss main.pas", false); err != nil {
		return err
	}

	bin, err := s.GetFile("main")
	if err != nil {
		return err
	}

	_, err = io.Copy(w, bin)

	return err
}

func (pascal) Run(s language.Sandbox, binary, stdin io.Reader, stdout io.Writer, tl time.Duration, ml int) (language.Status, error) {
	stat := language.Status{}
	stat.Verdict = language.VERDICT_XX

	if err := s.CreateFile("a.out", binary); err != nil {
		return stat, err
	}

	if err := s.MakeExecutable("a.out"); err != nil {
		return stat, err
	}

	return s.SetMaxProcesses(-1).Stdin(stdin).Stdout(stdout).TimeLimit(tl).MemoryLimit(ml/1024).Run("a.out", true)
}

func init() {
	language.Register("pascal", pascal{})
}
