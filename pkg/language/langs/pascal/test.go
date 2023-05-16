package pascal

import (
	"time"

	"github.com/mraron/njudge/pkg/language"
)

const print = `begin
    writeln('Hello world');
end.
`

func (p pascal) Test(s language.Sandbox) error {
	for _, test := range []language.LanguageTest{
		{p, print, language.VerdictOK, "", "Hello world\n", 1 * time.Second, 128 * 1024 * 1024},
	} {
		if err := test.Run(s); err != nil {
			return err
		}
	}

	return nil
}
