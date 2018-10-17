package cpp11

import (
	"bytes"
	"fmt"
	"github.com/mraron/njudge/utils/language"
	"io"
	"os/exec"
	"strings"
	"time"
)

type cpp11 struct{}

func (cpp11) Id() string {
	return "cpp11"
}

func (cpp11) Name() string {
	return "C++ 11"
}

func (cpp11) InsecureCompile(wd string, r io.Reader, w io.Writer, e io.Writer) error {
	cmd := exec.Command("g++", "-std=c++11", "-x", "c++", "-o", "/proc/self/fd/1", "-")

	cmd.Stdin = r
	cmd.Stdout = w
	cmd.Stderr = e

	cmd.Dir = wd

	return cmd.Run()
}

func (c cpp11) Compile(s language.Sandbox, r language.File, w io.Writer, e io.Writer, extras []language.File) error {
	if err := s.Init(); err != nil {
		fmt.Println("hmmm2")
		return err
	}
	fmt.Println("hmmm")

	defer s.Cleanup()

	err := s.CreateFile("main.cpp", r.Source)
	if err != nil {
		panic(err)
	}

	fmt.Println(extras, "???")

	params := "main.cpp"
	for _, f := range extras {
		err := s.CreateFile(f.Name, f.Source)
		if err != nil {
			panic(err)
		}

		if !strings.HasSuffix(f.Name, ".h") {
			params += " "
			params += f.Name
		}
	}

	errorStream := &bytes.Buffer{}
	if err := s.SetMaxProcesses(-1).Env().TimeLimit(10*time.Second).MemoryLimit(2560000).Stdout(errorStream).Stderr(e).WorkingDirectory(s.Pwd()).Run("/usr/bin/g++ -std=c++11 -static -DONLINE_JUDGE "+params, false); err != nil {
		e.Write(errorStream.Bytes())
		return err
	}

	bin, err := s.GetFile("a.out")
	if err != nil {
		return err
	}

	_, err = io.Copy(w, bin)
	return err
}

func (cpp11) Run(s language.Sandbox, binary, stdin io.Reader, stdout io.Writer, tl time.Duration, ml int) (language.Status, error) {
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
