package zip

import (
	"github.com/mraron/njudge/pkg/language/sandbox"
	"io"
	"time"

	"github.com/mraron/njudge/pkg/language"
)

type zip struct{}

func (zip) Id() string {
	return "zip"
}

func (zip) Name() string {
	return "ZIP archívum"
}

func (zip) DefaultFileName() string {
	return "main.zip"
}

func (zip) Compile(s sandbox.Sandbox, src language.File, bin io.Writer, cerr io.Writer, extras []language.File) error {
	return nil
}

func (zip) Run(s sandbox.Sandbox, binary io.Reader, stdin io.Reader, stdout io.Writer, tl time.Duration, ml int) (*sandbox.Status, error) {
	return &sandbox.Status{}, nil
}

func (zip) Test(sandbox.Sandbox) error {
	return nil
}

func init() {
	language.DefaultStore.Register("zip", zip{})
}
