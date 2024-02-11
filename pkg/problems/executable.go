package problems

import (
	"github.com/mraron/njudge/pkg/language/sandbox"
	"io"
)

type Executable interface {
	Execute(stdin io.Reader, stdout io.Writer, stderr io.Writer, args ...string) (int, error)
}

type SafelyExecutable interface {
	SafelyExecute(s sandbox.Sandbox, stdin io.Reader, stdout io.Writer, stderr io.Writer, args ...string) (int, error)
}
