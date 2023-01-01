package java

import (
	"errors"
	"io"
	"time"

	"github.com/mraron/njudge/pkg/language"
)

type java struct{}

func (java) Id() string {
	return "java"
}

func (java) Name() string {
	return "Java"
}

func (java) DefaultFileName() string {
	return "main.java"
}

func (java) InsecureCompile(wd string, r io.Reader, w io.Writer, e io.Writer) error {
	return errors.New("can't insecure compile java")
}

func (java) Compile(s language.Sandbox, r language.File, w io.Writer, e io.Writer, extras []language.File) error {
	err := s.CreateFile("main.java", r.Source)
	if err != nil {
		return err
	}

	if _, err := s.AddArg("--open-files=2048").SetMaxProcesses(-1).Env().TimeLimit(10*time.Second).MemoryLimit(4*256000).Stdout(e).Stderr(e).WorkingDirectory(s.Pwd()).Run("/usr/bin/javac main.java", false); err != nil {
		return err
	}

	bin, err := s.GetFile("main.class")
	if err != nil {
		return err
	}

	_, err = io.Copy(w, bin)

	return err
}

func (java) Run(s language.Sandbox, binary, stdin io.Reader, stdout io.Writer, tl time.Duration, ml int) (language.Status, error) {
	stat := language.Status{}
	stat.Verdict = language.VERDICT_XX

	if err := s.CreateFile("main.class", binary); err != nil {
		return stat, err
	}

	return s.SetMaxProcesses(-1).Env().Stdin(stdin).Stdout(stdout).TimeLimit(tl).MemoryLimit(ml/1024).WorkingDirectory(s.Pwd()).Run("/usr/bin/java main", true)
}

func init() {
	language.Register("java", java{})
}
