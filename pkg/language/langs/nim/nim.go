package nim

import (
	"context"
	"github.com/mraron/njudge/pkg/language/langs/cpp"
	"github.com/mraron/njudge/pkg/language/memory"
	"github.com/mraron/njudge/pkg/language/sandbox"
	"io"
	"io/fs"
	"time"

	"github.com/mraron/njudge/pkg/language"
)

type nim struct{}

func (nim) Id() string {
	return "nim"
}

func (nim) DisplayName() string {
	return "Nim"
}

func (nim) DefaultFilename() string {
	return "main.nim"
}

func (nim) Compile(s sandbox.Sandbox, r language.File, w io.Writer, e io.Writer, extras []language.File) error {
	err := sandbox.CreateFileFromSource(s, "main.nim", r.Source)
	if err != nil {
		return err
	}

	rc := sandbox.RunConfig{
		MaxProcesses:     -1,
		InheritEnv:       true,
		TimeLimit:        10 * time.Second,
		MemoryLimit:      256 * memory.MiB,
		Stdout:           e,
		Stderr:           e,
		WorkingDirectory: s.Pwd(),
		DirectoryMaps: []sandbox.DirectoryMap{
			{
				Inside:  "/etc",
				Outside: "/etc",
				Options: []sandbox.DirectoryMapOption{sandbox.NoExec},
			},
		},
	}

	if _, err := s.Run(context.TODO(), rc, "/usr/bin/nim", sandbox.SplitArgs("compile -d:release --nimcache=. main.nim")...); err != nil {
		return err
	}

	bin, err := s.Open("main")
	if err != nil {
		return err
	}
	defer func(bin fs.File) {
		_ = bin.Close()
	}(bin)

	_, err = io.Copy(w, bin)

	return err
}

func (nim) Run(s sandbox.Sandbox, binary io.Reader, stdin io.Reader, stdout io.Writer, tl time.Duration, ml memory.Amount) (*sandbox.Status, error) {
	return cpp.RunBinary("a.out")(s, binary, stdin, stdout, tl, ml)
}

func init() {
	language.DefaultStore.Register("nim", nim{})
}
