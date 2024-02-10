package cpp

import (
	"context"
	"github.com/mraron/njudge/pkg/language/memory"
	"github.com/mraron/njudge/pkg/language/sandbox"
	"io"
	"io/fs"
	"strings"
	"time"

	"github.com/mraron/njudge/pkg/language"
)

type Cpp struct {
	ID   string
	name string
	ver  string
}

func (c Cpp) Id() string {
	return c.ID
}

func (c Cpp) DisplayName() string {
	return c.name
}

func (c Cpp) DefaultFilename() string {
	return "main.cpp"
}

func (c Cpp) Compile(s sandbox.Sandbox, r language.File, w io.Writer, e io.Writer, extras []language.File) error {
	err := sandbox.CreateFileFromSource(s, "main.cpp", r.Source)
	if err != nil {
		return err
	}

	params := "main.cpp"
	for _, f := range extras {
		err := sandbox.CreateFileFromSource(s, f.Name, f.Source)
		if err != nil {
			return err
		}

		if !strings.HasSuffix(f.Name, ".h") {
			params += " "
			params += f.Name
		}
	}

	rc := sandbox.RunConfig{
		MaxProcesses:     200,
		InheritEnv:       true,
		TimeLimit:        10 * time.Second,
		MemoryLimit:      256 * memory.MiB,
		Stdout:           e,
		Stderr:           e,
		WorkingDirectory: s.Pwd(),
	}
	if _, err := s.Run(context.TODO(), rc, "/usr/bin/g++", sandbox.SplitArgs("-std="+c.ver+" -O2 -static -DONLINE_JUDGE "+params)...); err != nil {
		return err
	}

	bin, err := s.Open("a.out")
	if err != nil {
		return err
	}
	defer func(bin fs.File) {
		_ = bin.Close()
	}(bin)

	_, err = io.Copy(w, bin)
	return err
}

func (Cpp) Run(s sandbox.Sandbox, binary io.Reader, stdin io.Reader, stdout io.Writer, tl time.Duration, ml memory.Amount) (*sandbox.Status, error) {
	return RunBinary("a.out")(s, binary, stdin, stdout, tl, ml)
}

func RunBinary(binaryName string) func(sandbox.Sandbox, io.Reader, io.Reader, io.Writer, time.Duration, memory.Amount) (*sandbox.Status, error) {
	return func(s sandbox.Sandbox, binary io.Reader, stdin io.Reader, stdout io.Writer, tl time.Duration, ml memory.Amount) (*sandbox.Status, error) {
		stat := sandbox.Status{}
		stat.Verdict = sandbox.VerdictXX

		if err := sandbox.CreateFileFromSource(s, "a.out", binary); err != nil {
			return nil, err
		}

		if err := s.MakeExecutable("a.out"); err != nil {
			return nil, err
		}

		rc := sandbox.RunConfig{
			Stdin:       stdin,
			Stdout:      stdout,
			TimeLimit:   tl,
			MemoryLimit: memory.Amount(ml) * memory.KiB,
		}
		return s.Run(context.TODO(), rc, "a.out")
	}
}

func New(id, name, ver string) language.Language {
	return Cpp{id, name, ver}
}

var Std11 = New("cpp11", "C++ 11", "c++11").(Cpp)
var Std14 = New("cpp14", "C++ 14", "c++14").(Cpp)
var Std17 = New("cpp17", "C++ 17", "c++17").(Cpp)

var latest = Std17

func init() {
	language.DefaultStore.Register("cpp11", Std11)
	language.DefaultStore.Register("cpp14", Std14)
	language.DefaultStore.Register("cpp17", Std17)
}
