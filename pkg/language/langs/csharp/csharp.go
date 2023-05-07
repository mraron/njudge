package csharp

import (
	"errors"
	"io"
	"time"

	"github.com/mraron/njudge/pkg/language"
)

type csharp struct{}

func (csharp) Id() string {
	return "csharp"
}

func (csharp) Name() string {
	return "C# (mono)"
}

func (csharp) DefaultFileName() string {
	return "main.cs"
}

func (csharp) InsecureCompile(wd string, r io.Reader, w io.Writer, e io.Writer) error {
	return errors.New("can't insecure compile c#")
}

func (csharp) Compile(s language.Sandbox, r language.File, w io.Writer, e io.Writer, extras []language.File) error {
	err := s.CreateFile("main.cs", r.Source)
	if err != nil {
		return err
	}

	if _, err := s.Env().MapDir("/etc", "/etc", []string{}, false).SetMaxProcesses(-1).TimeLimit(10*time.Second).MemoryLimit(4*256000).Stdout(e).Stderr(e).Run("/usr/bin/mcs -out:main.exe -optimize+ main.cs", false); err != nil {
		return err
	}

	bin, err := s.GetFile("main.exe")
	if err != nil {
		return err
	}

	_, err = io.Copy(w, bin)

	return err
}

func (csharp) Run(s language.Sandbox, binary, stdin io.Reader, stdout io.Writer, tl time.Duration, ml int) (language.Status, error) {
	stat := language.Status{}
	stat.Verdict = language.VERDICT_XX

	if err := s.CreateFile("main.exe", binary); err != nil {
		return stat, err
	}

	return s.SetMaxProcesses(-1).Env().Stdin(stdin).Stdout(stdout).TimeLimit(tl).MemoryLimit(ml/1024).WorkingDirectory(s.Pwd()).Run("/usr/bin/mono main.exe", true)
}

func init() {
	language.DefaultStore.Register("csharp", csharp{})
}
