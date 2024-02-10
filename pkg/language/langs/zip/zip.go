package zip

import (
	"github.com/mraron/njudge/pkg/language/memory"
	"github.com/mraron/njudge/pkg/language/sandbox"
	"io"
	"time"

	"github.com/mraron/njudge/pkg/language"
)

type zip struct{}

func (zip) ID() string {
	return "zip"
}

func (zip) DisplayName() string {
	return "ZIP archívum"
}

func (zip) DefaultFilename() string {
	return "main.zip"
}

func (zip) Compile(s sandbox.Sandbox, f sandbox.File, stderr io.Writer, extras []sandbox.File) (*sandbox.File, error) {
	return nil, nil
}

func (zip) Run(s sandbox.Sandbox, binary sandbox.File, stdin io.Reader, stdout io.Writer, tl time.Duration, ml memory.Amount) (*sandbox.Status, error) {
	return &sandbox.Status{}, nil
}

func (zip) Test(sandbox.Sandbox) error {
	return nil
}

func init() {
	language.DefaultStore.Register("zip", zip{})
}
