package octave

import (
	"github.com/mraron/njudge/utils/language"
	"io"
	"os"
	"time"
)

type octave struct{}

func (octave) Id() string {
	return "octave"
}

func (octave) Name() string {
	return "Octave"
}

func (octave) DefaultFileName() string {
	return "main.m"
}

func (octave) InsecureCompile(wd string, r io.Reader, w io.Writer, e io.Writer) error {
	return nil
}

func (octave) Compile(s language.Sandbox, r language.File, w io.Writer, e io.Writer, extras []language.File) error {
	_, err := io.Copy(w, r.Source)
	return err
}

func (octave) Run(s language.Sandbox, binary, stdin io.Reader, stdout io.Writer, tl time.Duration, ml int) (language.Status, error) {
	stat := language.Status{}
	stat.Verdict = language.VERDICT_XX

	if err := s.CreateFile("a.out", binary); err != nil {
		return stat, err
	}

	lapack, err := os.Open("/usr/lib/x86_64-linux-gnu/lapack/liblapack.so.3")
	if err != nil {
		return stat, err
	}
	defer lapack.Close()

	if err := s.CreateFile("liblapack.so.3", lapack); err != nil {
		return stat, err
	}

	if st, err := s.Verbose().Env().SetMaxProcesses(100).TimeLimit(tl).MemoryLimit(ml/1024).Stdin(stdin).Stdout(stdout).Stderr(os.Stderr).WorkingDirectory(s.Pwd()).Run("/usr/bin/octave-cli a.out", true); err != nil {
		return st, err
	} else {
		stat = st
	}

	return stat, nil
}

func init() {
	language.Register("octave", octave{})
}
