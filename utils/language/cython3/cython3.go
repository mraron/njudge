package cython3

import (
	"bytes"
	"errors"
	"github.com/mraron/njudge/utils/language"
	"io"
	"time"
)


type cython3 struct {
}

func (c cython3) Id() string {
	return "cython3"
}

func (c cython3) Name() string {
	return "Cython3"
}

func (c cython3) DefaultFileName() string {
	return "main.py"
}

func (c cython3) InsecureCompile(wd string, r io.Reader, w io.Writer, e io.Writer) error {
	return errors.New("not supported")
}

func (c cython3) Compile(s language.Sandbox, r language.File, w io.Writer, e io.Writer, extras []language.File) error {
	err := s.CreateFile("main.py", r.Source)
	if err != nil {
		return err
	}

	errorStream := &bytes.Buffer{}
	if _, err := s.SetMaxProcesses(200).Env().TimeLimit(10*time.Second).MemoryLimit(2560000).Stdout(errorStream).Stderr(e).WorkingDirectory(s.Pwd()).Run("/usr/bin/cython3 -3 --embed -o main.c main.py", false); err != nil {
		e.Write(errorStream.Bytes())
		return err
	}

	if _, err := s.SetMaxProcesses(200).Env().TimeLimit(10*time.Second).MemoryLimit(2560000).Stdout(errorStream).Stderr(e).WorkingDirectory(s.Pwd()).Run("/usr/bin/gcc -O2 -I/usr/include/python3.9 -o main main.c -lpython3.9 -lpthread -lm -lutil -ldl", false); err != nil {
		e.Write(errorStream.Bytes())
		return err
	}

	bin, err := s.GetFile("main")
	if err != nil {
		return err
	}

	_, err = io.Copy(w, bin)
	return err
}

func (cython3) Run(s language.Sandbox, binary, stdin io.Reader, stdout io.Writer, tl time.Duration, ml int) (language.Status, error) {
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

func init() {
	language.Register("cython3", cython3{})
}
