package runner

import "io"

type Func func(args []string, stdin io.Reader, stdout io.Writer, stderr io.Writer) (int, error)

type Function struct {
	f Func

	stdin          io.Reader
	stdout, stderr io.Writer

	err        error
	returnCode int
}

func NewFunction(f Func) *Function {
	return &Function{f: f}
}

func (s *Function) Stdin(r io.Reader) {
	s.stdin = r
}

func (s *Function) Stdout(w io.Writer) {
	s.stdout = w
}

func (s *Function) Stderr(w io.Writer) {
	s.stderr = w
}

func (s *Function) Run(args []string) error {
	s.returnCode, s.err = s.f(args, s.stdin, s.stdout, s.stderr)
	return s.err
}

func (s *Function) ReturnCode() (int, error) {
	return s.returnCode, nil
}
