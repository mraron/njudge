package nim

import (
	"errors"
	"io"
	"time"

	"github.com/mraron/njudge/pkg/language"
)

type nim struct{}

func (nim) Id() string {
	return "nim"
}

func (nim) Name() string {
	return "Nim"
}

func (nim) DefaultFileName() string {
	return "main.nim"
}

func (nim) InsecureCompile(wd string, r io.Reader, w io.Writer, e io.Writer) error {
	return errors.New("not supported")
}

func (nim) Compile(s language.Sandbox, r language.File, w io.Writer, e io.Writer, extras []language.File) error {
	err := s.CreateFile("main.nim", r.Source)
	if err != nil {
		return err
	}

	s.MapDir("/etc", "/etc", []string{}, true)

	if _, err := s.SetMaxProcesses(-1).Env().TimeLimit(10*time.Second).MemoryLimit(256000).Stdout(e).Stderr(e).WorkingDirectory(s.Pwd()).Run("/usr/bin/nim compile -d:release --nimcache=. main.nim", false); err != nil {
		return err
	}

	bin, err := s.GetFile("main")
	if err != nil {
		return err
	}

	_, err = io.Copy(w, bin)

	return err
}

func (nim) Run(s language.Sandbox, binary, stdin io.Reader, stdout io.Writer, tl time.Duration, ml int) (language.Status, error) {
	stat := language.Status{}
	stat.Verdict = language.VerdictXX

	if err := s.CreateFile("a.out", binary); err != nil {
		return stat, err
	}

	if err := s.MakeExecutable("a.out"); err != nil {
		return stat, err
	}

	return s.SetMaxProcesses(-1).Stdin(stdin).Stdout(stdout).TimeLimit(tl).MemoryLimit(ml/1024).Run("a.out", true)
}

func init() {
	language.DefaultStore.Register("nim", nim{})
}
