package cpp11

import (
	"bytes"
	"fmt"
	"github.com/mraron/njudge/utils/language"
	"io"
	"os/exec"
	"time"
)

type cpp11 struct {}

func (cpp11) Id() string {
	return "cpp11"
}

func (cpp11) Name() string {
	return "C++ 11"
}

func (cpp11) InsecureCompile(wd string, r io.Reader, w io.Writer, e io.Writer) (error) {
	cmd := exec.Command("g++","-std=c++11", "-x", "c++", "-o", "/proc/self/fd/1", "-")

	cmd.Stdin = r
	cmd.Stdout = w
	cmd.Stderr = e

	cmd.Dir = wd

	return cmd.Run()
}

func (c cpp11) Compile(s language.Sandbox, r io.Reader, w io.Writer, e io.Writer) error {
	if err := s.Init(); err != nil {
		fmt.Println("hmmm2")
		return err
	}
	fmt.Println("hmmm")

	defer s.Cleanup()

	bin := &bytes.Buffer{}
	if err := s.SetMaxProcesses(-1).Env().TimeLimit(10 * time.Second).MemoryLimit(2560000).Stdin(r).Stdout(bin).Stderr(e).WorkingDirectory("/tmp").Run("/usr/bin/g++ -std=c++11 -static -DONLINE_JUDGE -x c++ -o /proc/self/fd/1 -", false); err != nil {
		e.Write(bin.Bytes())
		return err
	}

	_, err := w.Write(bin.Bytes())

	return err
}

func (cpp11) Run(s language.Sandbox, binary, stdin io.Reader, stdout io.Writer, tl time.Duration, ml int) (language.Status, error)  {
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

	err := s.Stdin(stdin).Stdout(stdout).TimeLimit(tl).MemoryLimit(ml/1024).Run("a.out", true)

	stat = s.Status()

	return stat, err
}



func init() {
	language.Register("cpp11", cpp11{})
}