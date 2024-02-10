package csharp

import (
	"context"
	"github.com/mraron/njudge/pkg/language/memory"
	"github.com/mraron/njudge/pkg/language/sandbox"
	"io"
	"time"

	"github.com/mraron/njudge/pkg/language"
)

type CSharp struct{}

func (CSharp) ID() string {
	return "csharp"
}

func (CSharp) DisplayName() string {
	return "C# (mono)"
}

func (CSharp) DefaultFilename() string {
	return "main.cs"
}

func (CSharp) Compile(s sandbox.Sandbox, f sandbox.File, stderr io.Writer, extras []sandbox.File) (*sandbox.File, error) {
	err := sandbox.CreateFileFromSource(s, f.Name, f.Source)
	if err != nil {
		return nil, err
	}

	rc := sandbox.RunConfig{
		InheritEnv:    true,
		DirectoryMaps: []sandbox.DirectoryMap{{"/etc", "/etc", nil}},
		MaxProcesses:  -1,
		TimeLimit:     10 * time.Second,
		MemoryLimit:   1 * memory.GiB,
		Stdout:        stderr,
		Stderr:        stderr,
	}
	if _, err := s.Run(context.TODO(), rc, "/usr/bin/mcs", sandbox.SplitArgs("-out:main.exe -optimize+ "+f.Name)...); err != nil {
		return nil, err
	}

	return sandbox.ExtractFile(s, "main.exe")
}

func (CSharp) Run(s sandbox.Sandbox, binary sandbox.File, stdin io.Reader, stdout io.Writer, tl time.Duration, ml memory.Amount) (*sandbox.Status, error) {
	stat := sandbox.Status{}
	stat.Verdict = sandbox.VerdictXX

	if err := sandbox.CreateFileFromSource(s, binary.Name, binary.Source); err != nil {
		return nil, err
	}

	rc := sandbox.RunConfig{
		MaxProcesses:     -1,
		InheritEnv:       true,
		Stdin:            stdin,
		Stdout:           stdout,
		TimeLimit:        tl,
		MemoryLimit:      ml,
		WorkingDirectory: s.Pwd(),
	}
	return s.Run(context.TODO(), rc, "/usr/bin/mono", "main.exe")
}

func init() {
	language.DefaultStore.Register("csharp", CSharp{})
}
