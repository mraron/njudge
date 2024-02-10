package executable

import (
	"errors"
	"io"
	"os/exec"
	"strings"
	"syscall"
)

type Stdlib struct {
	Cmd  *exec.Cmd
	Path string
}

func NewStdlib(path string) *Stdlib {
	return &Stdlib{Path: path}
}

func (s *Stdlib) Execute(stdin io.Reader, stdout io.Writer, stderr io.Writer, args ...string) (int, error) {
	s.Cmd = exec.Command("/bin/sh", "-c", "ulimit -s unlimited && "+strings.Join(append([]string{s.Path}, args...), " "))
	s.Cmd.Stdin = stdin
	s.Cmd.Stdout = stdout
	s.Cmd.Stderr = stderr

	err := s.Cmd.Run()
	var exitErr *exec.ExitError
	if errors.As(err, &exitErr) {
		if status, ok := exitErr.Sys().(syscall.WaitStatus); ok {
			return status.ExitStatus(), nil
		} else {
			return -1, err
		}
	}

	return -1, err
}
