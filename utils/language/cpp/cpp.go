package cpp

import (
	"bytes"
	"github.com/mraron/njudge/utils/language"
	"io"
	"os/exec"
	"strings"
	"time"
)

type cpp struct {
	id   string
	name string
	ver  string
}

func (c cpp) Id() string {
	return c.id
}

func (c cpp) Name() string {
	return c.name
}

func (c cpp) DefaultFileName() string {
	return "main.cpp"
}

func (c cpp) InsecureCompile(wd string, r io.Reader, w io.Writer, e io.Writer) error {
	cmd := exec.Command("g++", "-std="+c.ver, "-x", "c++", "-O2", "-o", "/proc/self/fd/1", "-")

	cmd.Stdin = r
	cmd.Stdout = w
	cmd.Stderr = e

	cmd.Dir = wd

	return cmd.Run()
}

func (c cpp) Compile(s language.Sandbox, r language.File, w io.Writer, e io.Writer, extras []language.File) error {
	err := s.CreateFile("main.cpp", r.Source)
	if err != nil {
		return err
	}

	params := "main.cpp"
	for _, f := range extras {
		err := s.CreateFile(f.Name, f.Source)
		if err != nil {
			return err
		}

		if !strings.HasSuffix(f.Name, ".h") {
			params += " "
			params += f.Name
		}
	}

	errorStream := &bytes.Buffer{}
	if _, err := s.SetMaxProcesses(200).Env().TimeLimit(10*time.Second).MemoryLimit(2560000).Stdout(errorStream).Stderr(e).WorkingDirectory(s.Pwd()).Run("/usr/bin/g++ -std="+c.ver+" -O2 -static -DONLINE_JUDGE "+params, false); err != nil {
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

func (cpp) Run(s language.Sandbox, binary, stdin io.Reader, stdout io.Writer, tl time.Duration, ml int) (language.Status, error) {
	stat := language.Status{}
	stat.Verdict = language.VERDICT_XX

	if err := s.CreateFile("a.out", binary); err != nil {
		return stat, err
	}

	if err := s.MakeExecutable("a.out"); err != nil {
		return stat, err
	}

	return s.Stdin(stdin).Stdout(stdout).TimeLimit(tl).MemoryLimit(ml/1024).Run("a.out", true)
}

func New(id, name, ver string) language.Language {
	return cpp{id, name, ver}
}
