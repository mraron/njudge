package julia

import (
	"time"

	"github.com/mraron/njudge/pkg/language"
)

const print = `println("Hello world")`

func (j julia) Test(s language.Sandbox) error {
	for _, test := range []language.LanguageTest{
		{j, print, language.VerdictOK, "", "Hello world\n", 1 * time.Second, 128 * 1024 * 1024},
	} {
		if err := test.Run(s); err != nil {
			return err
		}
	}

	return nil
}
