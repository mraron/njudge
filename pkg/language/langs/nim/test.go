package nim

import (
	"time"

	"github.com/mraron/njudge/pkg/language"
)

const print = `echo "Hello world"`

func (n nim) Test(s language.Sandbox) error {
	for _, test := range []language.LanguageTest{
		{Language: n, Source: print, ExpectedVerdict: language.VERDICT_OK, Input: "", ExpectedOutput: "Hello world\n", TimeLimit: 1 * time.Second, MemoryLimit: 128 * 1024 * 1024},
	} {
		if err := test.Run(s); err != nil {
			return err
		}
	}

	return nil
}
