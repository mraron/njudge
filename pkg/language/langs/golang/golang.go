package golang

import (
	"github.com/mraron/njudge/pkg/language"
	"io"
	"os/exec"
	"time"
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

func (golang) InsecureCompile(wd string, r io.Reader, w io.Writer, e io.Writer) error {
	cmd := exec.Command("gccgo", "-x", "go", "-o", "/dev/stdout", "-")

	cmd.Stdin = r
	cmd.Stdout = w
	cmd.Stderr = e

	cmd.Dir = wd

	return cmd.Run()
}

func (golang) Compile(s language.Sandbox, r language.File, w io.Writer, e io.Writer, extras []language.File) error {
	err := s.CreateFile("main.go", r.Source)
	if err != nil {
		return err
	}

	if _, err := s.AddArg("--open-files=2048").SetMaxProcesses(-1).Env().SetEnv("GOCACHE=/tmp").TimeLimit(10*time.Second).MemoryLimit(4*256000).Stdout(e).Stderr(e).WorkingDirectory(s.Pwd()).Run("/usr/bin/gccgo main.go", false); err != nil {
		return err
	}

	bin, err := s.GetFile("a.out")
	if err != nil {
		return err
	}

	_, err = io.Copy(w, bin)

	return err
}

func (golang) Run(s language.Sandbox, binary, stdin io.Reader, stdout io.Writer, tl time.Duration, ml int) (language.Status, error) {
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
	language.Register("golang", golang{})
}
