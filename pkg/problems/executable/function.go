package executable

import (
	"github.com/mraron/njudge/pkg/language/sandbox"
	"io"
)

type Func func(stdin io.Reader, stdout io.Writer, stderr io.Writer, args ...string) (int, error)

type Function struct {
	f Func
}

func NewFunction(f Func) *Function {
	return &Function{f: f}
}

func (s *Function) Execute(stdin io.Reader, stdout io.Writer, stderr io.Writer, args ...string) (int, error) {
	return s.f(stdin, stdout, stderr, args...)
}

func (s *Function) SafelyExecute(_ sandbox.Sandbox, stdin io.Reader, stdout io.Writer, stderr io.Writer, args ...string) (int, error) {
	return s.Execute(stdin, stdout, stderr, args...)
}
