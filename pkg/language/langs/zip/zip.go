package zip

import (
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

func (zip) InsecureCompile(s string, r io.Reader, w1 io.Writer, w2 io.Writer) error {
	return nil
}

func (zip) Compile(s language.Sandbox, src language.File, bin io.Writer, cerr io.Writer, extras []language.File) error {
	return nil
}

func (zip) Run(s language.Sandbox, binary io.Reader, stdin io.Reader, stdout io.Writer, tl time.Duration, mem int) (language.Status, error) {
	return language.Status{}, nil
}

func (zip) Test(language.Sandbox) error {
	return nil
}

func init() {
	language.DefaultStore.Register("zip", zip{})
}
