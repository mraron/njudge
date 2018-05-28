package zip

import (
	"io"
	"github.com/mraron/njudge/utils/language"
	"time"
)

type zip struct {}

func (zip) Id() string {
	return "zip"
}

func (zip) Name() string {
	return "ZIP archívum"
}

func (zip) InsecureCompile(s string, r io.Reader, w1 io.Writer, w2 io.Writer) (error) {
	return nil
}

func (zip) Compile(s language.Sandbox, src io.Reader, bin io.Writer, cerr io.Writer) (error) {
	return nil
}

func (zip) Run(s language.Sandbox, binary io.Reader, stdin io.Reader, stdout io.Writer, tl time.Duration, mem int) (language.Status, error) {
	return language.Status{}, nil
}

func init() {
	language.Register("zip", zip{})
}