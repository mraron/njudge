package csharp

import (
	"context"
	"github.com/mraron/njudge/pkg/language/memory"
	"github.com/mraron/njudge/pkg/language/sandbox"
	"io"
	"io/fs"
	"time"

	"github.com/mraron/njudge/pkg/language"
)

type csharp struct{}

func (csharp) Id() string {
	return "csharp"
}

func (csharp) DisplayName() string {
	return "C# (mono)"
}

func (csharp) DefaultFilename() string {
	return "main.cs"
}

func (csharp) Compile(s sandbox.Sandbox, r language.File, w io.Writer, e io.Writer, extras []language.File) error {
	err := sandbox.CreateFileFromSource(s, "main.cs", r.Source)
	if err != nil {
		return err
	}

	rc := sandbox.RunConfig{
		InheritEnv:    true,
		DirectoryMaps: []sandbox.DirectoryMap{{"/etc", "/etc", nil}},
		MaxProcesses:  -1,
		TimeLimit:     10 * time.Second,
		MemoryLimit:   1 * memory.GiB,
		Stdout:        e,
		Stderr:        e,
	}
	if _, err := s.Run(context.TODO(), rc, "/usr/bin/mcs", sandbox.SplitArgs("-out:main.exe -optimize+ main.cs")...); err != nil {
		return err
	}

	bin, err := s.Open("main.exe")
	if err != nil {
		return err
	}
	defer func(bin fs.File) {
		_ = bin.Close()
	}(bin)

	_, err = io.Copy(w, bin)

	return err
}

func (csharp) Run(s sandbox.Sandbox, binary io.Reader, stdin io.Reader, stdout io.Writer, tl time.Duration, ml memory.Amount) (*sandbox.Status, error) {
	stat := sandbox.Status{}
	stat.Verdict = sandbox.VerdictXX

	if err := sandbox.CreateFileFromSource(s, "main.exe", binary); err != nil {
		return nil, err
	}

	rc := sandbox.RunConfig{
		MaxProcesses:     -1,
		InheritEnv:       true,
		Stdin:            stdin,
		Stdout:           stdout,
		TimeLimit:        tl,
		MemoryLimit:      memory.Amount(ml) * memory.KiB,
		WorkingDirectory: s.Pwd(),
	}
	return s.Run(context.TODO(), rc, "/usr/bin/mono", "main.exe")
}

func init() {
	language.DefaultStore.Register("csharp", csharp{})
}
