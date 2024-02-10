// Package language is used to compile and run untrusted user code securely.
// The heavy lifting is done by the great [isolate] library.
//
// [isolate]: https://github.com/ioi/isolate
package language

import (
	"github.com/mraron/njudge/pkg/language/memory"
	"github.com/mraron/njudge/pkg/language/sandbox"
	"io"
	"time"
)

// Language is @TODO
type Language interface {
	ID() string
	DisplayName() string
	DefaultFilename() string
	Compile(s sandbox.Sandbox, f sandbox.File, stderr io.Writer, extras []sandbox.File) (*sandbox.File, error) //TODO remove extras?
	Run(s sandbox.Sandbox, binary sandbox.File, stdin io.Reader, stdout io.Writer, tl time.Duration, ml memory.Amount) (*sandbox.Status, error)
}
