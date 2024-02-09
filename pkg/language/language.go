package language

import (
	"github.com/mraron/njudge/pkg/language/sandbox"
	"io"
	"time"
)

type File struct {
	Name   string
	Source io.Reader
}

type Language interface {
	Id() string
	Name() string //TODO remove this
	DefaultFileName() string
	Compile(s sandbox.Sandbox, f File, binary io.Writer, stderr io.Writer, extras []File) error
	Run(s sandbox.Sandbox, binary io.Reader, stdin io.Reader, stdout io.Writer, tl time.Duration, ml int) (*sandbox.Status, error) //TODO change ml to memory.Amount
}
