package golang

import (
	"time"
	"github.com/mraron/njudge/utils/language"
	"io"
	"os/exec"
	"bytes"
)

type golang struct{}

func (golang) Id() string {
	return "golang"
}

func (golang) Name() string {
	return "Go"
}


func (golang) InsecureCompile(wd string, r io.Reader, w io.Writer, e io.Writer) (error) {
	cmd := exec.Command("gccgo","-x", "go", "-o", "/dev/stdout", "-")

	cmd.Stdin = r
	cmd.Stdout = w
	cmd.Stderr = e

	cmd.Dir = wd

	return cmd.Run()
}

func (golang) Compile(s language.Sandbox, r io.Reader, w io.Writer, e io.Writer) error {
	if err := s.Init(); err != nil {
		return err
	}

	defer s.Cleanup()

	bin := &bytes.Buffer{}

	if err := s.SetMaxProcesses(-1).Env().TimeLimit(10 * time.Second).MemoryLimit(256000).Stdin(r).Stdout(bin).Stderr(e).WorkingDirectory("/tmp").Run("/usr/bin/gccgo -static -DONLINE_JUDGE -x go -o /dev/stdout -", false); err != nil {
		e.Write(bin.Bytes())
		return err
	}

	_, err := w.Write(bin.Bytes())

	return err
}

func (golang) Run(s language.Sandbox, binary, stdin io.Reader, stdout io.Writer, tl time.Duration, ml int) (language.Status, error)  {
	stat := language.Status{}
	stat.Verdict = language.VERDICT_XX

	if err := s.Init(); err != nil {
		return stat, err
	}

	defer s.Cleanup()

	if err := s.CreateFile("a.out", binary); err != nil {
		return stat, err
	}

	if err := s.MakeExecutable("a.out"); err != nil {
		return stat, err
	}

	err := s.SetMaxProcesses(-1).Stdin(stdin).Stdout(stdout).TimeLimit(tl).MemoryLimit(ml/1024).Run("a.out", true)

	stat = s.Status()

	return stat, err
}



func init() {
	language.Register("golang", golang{})
}