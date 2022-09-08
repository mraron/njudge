package runner

import (
	"io"
)

type Executable interface {
	Stdin(io.Reader)
	Stdout(io.Writer)
	Stderr(io.Writer)

	Run(args []string) error
	ReturnCode() (int, error)
}
