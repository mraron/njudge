package nim

import (
	"context"
	"github.com/mraron/njudge/pkg/language/memory"
	"github.com/mraron/njudge/pkg/language/sandbox"
	"io"
	"time"

	"github.com/mraron/njudge/pkg/language"
)

type nim struct{}

func (nim) ID() string {
	return "nim"
}

func (nim) DisplayName() string {
	return "Nim"
}

func (nim) DefaultFilename() string {
	return "main.nim"
}

func (nim) Compile(ctx context.Context, s sandbox.Sandbox, f sandbox.File, stderr io.Writer, extras []sandbox.File) (*sandbox.File, error) {
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

	if _, err := s.Run(ctx, rc, "/usr/bin/nim", sandbox.SplitArgs("compile -d:release --nimcache=. "+f.Name)...); err != nil {
		return nil, err
	}

	return sandbox.ExtractFile(s, "main")
}

func (nim) Run(ctx context.Context, s sandbox.Sandbox, binary sandbox.File, stdin io.Reader, stdout io.Writer, tl time.Duration, ml memory.Amount) (*sandbox.Status, error) {
	return sandbox.RunBinary(ctx, s, binary, stdin, stdout, tl, ml)
}

func init() {
	language.DefaultStore.Register("nim", nim{})
}
