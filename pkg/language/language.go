// Package language is used to compile and run untrusted code securely.
// The heavy lifting is done by the great [isolate] library.
// This library comes with a lot of built-in languages, and it's easy to implement your own languages.
//
// [isolate]: https://github.com/ioi/isolate
package language

import (
	"context"
	"github.com/mraron/njudge/pkg/language/memory"
	"github.com/mraron/njudge/pkg/language/sandbox"
	"io"
	"time"
)

// Language is the main building block of this package, it's used to compile a source file
// and then run the resulting binary.
type Language interface {
	ID() string
	DisplayName() string
	DefaultFilename() string
	Compile(ctx context.Context, s sandbox.Sandbox, f sandbox.File, stderr io.Writer, extras []sandbox.File) (*sandbox.File, error) //TODO remove extras?
	Run(ctx context.Context, s sandbox.Sandbox, binary sandbox.File, stdin io.Reader, stdout io.Writer, tl time.Duration, ml memory.Amount) (*sandbox.Status, error)
}
