package runner

import (
	"fmt"
	"io"
	"os/exec"
	"strings"
	"syscall"
)

type Stdlib struct {
	path string

	stdin          io.Reader
	stdout, stderr io.Writer

	cmd *exec.Cmd
	err error
}

func NewStdlib(path string) *Stdlib {
	return &Stdlib{path: path}
}

func (s *Stdlib) Stdin(r io.Reader) {
	s.stdin = r
}

func (s *Stdlib) Stdout(w io.Writer) {
	s.stdout = w
}

func (s *Stdlib) Stderr(w io.Writer) {
	s.stderr = w
}

func (s *Stdlib) Run(args []string) error {
	s.cmd = exec.Command("/bin/sh", "-c", "ulimit -s unlimited && "+strings.Join(append([]string{s.path}, args...), " "))
	s.cmd.Stdin = s.stdin
	s.cmd.Stdout = s.stdout
	s.cmd.Stderr = s.stderr

	s.err = s.cmd.Run()
	return s.err
}

func (s *Stdlib) ReturnCode() (int, error) {
	if exitErr, ok := s.err.(*exec.ExitError); ok {
		if status, ok := exitErr.Sys().(syscall.WaitStatus); ok {
			return status.ExitStatus(), nil
		} else {
			return -1, fmt.Errorf("can't convert to syscall.WaitStatus")
		}
	}

	return -1, fmt.Errorf("not an *exec.ExitError")
}
