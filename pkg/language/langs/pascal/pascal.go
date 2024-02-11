package pascal

import (
	"context"
	"github.com/mraron/njudge/pkg/language/memory"
	"github.com/mraron/njudge/pkg/language/sandbox"
	"io"
	"time"

	"github.com/mraron/njudge/pkg/language"
)

type Pascal struct{}

func (Pascal) ID() string {
	return "pascal"
}

func (Pascal) DisplayName() string {
	return "Pascal"
}

func (Pascal) DefaultFilename() string {
	return "main.pas"
}

func (Pascal) Compile(ctx context.Context, s sandbox.Sandbox, f sandbox.File, stderr io.Writer, extras []sandbox.File) (*sandbox.File, error) {
	err := sandbox.CreateFile(s, f)
	if err != nil {
		return nil, err
	}
	rc := sandbox.RunConfig{
		MaxProcesses:     -1,
		InheritEnv:       true,
		TimeLimit:        10 * time.Second,
		MemoryLimit:      256 * memory.MiB,
		Stdout:           stderr,
		Stderr:           stderr,
		WorkingDirectory: s.Pwd(),
		DirectoryMaps: []sandbox.DirectoryMap{
			{
				Inside:  "/etc",
				Outside: "/etc",
				Options: []sandbox.DirectoryMapOption{sandbox.NoExec},
			},
		},
	}
	if _, err := s.Run(ctx, rc, "/usr/bin/fpc", sandbox.SplitArgs("-Mobjfpc -O2 -Xss "+f.Name)...); err != nil {
		return nil, err
	}

	return sandbox.ExtractFile(s, "main")
}

func (Pascal) Run(ctx context.Context, s sandbox.Sandbox, binary sandbox.File, stdin io.Reader, stdout io.Writer, tl time.Duration, ml memory.Amount) (*sandbox.Status, error) {
	return sandbox.RunBinary(ctx, s, binary, stdin, stdout, tl, ml)
}

func init() {
	language.DefaultStore.Register("pascal", Pascal{})
}
