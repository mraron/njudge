package cpp

import (
	"bytes"
	"io"
	"strings"
	"time"

	"github.com/mraron/njudge/pkg/language"
)

type Cpp struct {
	id   string
	name string
	ver  string
}

func (c Cpp) Id() string {
	return c.id
}

func (c Cpp) Name() string {
	return c.name
}

func (c Cpp) DefaultFileName() string {
	return "main.cpp"
}

func (c Cpp) Compile(s language.Sandbox, r language.File, w io.Writer, e io.Writer, extras []language.File) error {
	err := s.CreateFile("main.cpp", r.Source)
	if err != nil {
		return err
	}

	params := "main.cpp"
	for _, f := range extras {
		err := s.CreateFile(f.Name, f.Source)
		if err != nil {
			return err
		}

		if !strings.HasSuffix(f.Name, ".h") {
			params += " "
			params += f.Name
		}
	}

	errorStream := &bytes.Buffer{}
	if _, err := s.SetMaxProcesses(200).Env().TimeLimit(10*time.Second).MemoryLimit(2560000).Stdout(errorStream).Stderr(e).WorkingDirectory(s.Pwd()).Run("/usr/bin/g++ -std="+c.ver+" -O2 -static -DONLINE_JUDGE "+params, false); err != nil {
		e.Write(errorStream.Bytes())
		return err
	}

	bin, err := s.GetFile("a.out")
	if err != nil {
		return err
	}

	_, err = io.Copy(w, bin)
	return err
}

func (Cpp) Run(s language.Sandbox, binary, stdin io.Reader, stdout io.Writer, tl time.Duration, ml int) (language.Status, error) {
	stat := language.Status{}
	stat.Verdict = language.VerdictXX

	if err := s.CreateFile("a.out", binary); err != nil {
		return stat, err
	}

	if err := s.MakeExecutable("a.out"); err != nil {
		return stat, err
	}

	return s.Stdin(stdin).Stdout(stdout).TimeLimit(tl).MemoryLimit(ml/1024).Run("a.out", true)
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
