package cython3

import (
	"time"

	"github.com/mraron/njudge/pkg/language"
)

const (
	PYTHON3_aplusb = `inp = input().split(' ')
a,b = int(inp[0]), int(inp[1]) 
print(a+b)`
	PYTHON3_ce    = `inp = input(()`
	PYTHON3_print = `print("Hello world")`
	PYTHON3_tl    = `x = 0
while True:
	x = x+1`
	PYTHON3_re = `x = [1,2,3]
print(x[4])`
	PYTHON3_rediv0 = `print(1/0)`
)

func (c cython3) Test(s language.Sandbox) error {
	for _, test := range []language.LanguageTest{
		{c, PYTHON3_aplusb, language.VERDICT_OK, "1 2\n", "3\n", 1500 * time.Millisecond, 128 * 1024 * 1024},
		{c, PYTHON3_ce, language.VERDICT_CE, "", "", 1 * time.Second, 128 * 1024 * 1024},
		{c, PYTHON3_print, language.VERDICT_OK, "", "Hello world\n", 1 * time.Second, 128 * 1024 * 1024},
		{c, PYTHON3_tl, language.VERDICT_TL, "", "", 100 * time.Millisecond, 128 * 1024 * 1024},
		{c, PYTHON3_re, language.VERDICT_RE, "", "", 1000 * time.Millisecond, 128 * 1024 * 1024},
		{c, PYTHON3_rediv0, language.VERDICT_RE, "", "", 1000 * time.Millisecond, 128 * 1024 * 1024},
	} {
		if err := test.Run(s); err != nil {
			return err
		}
	}

	return nil
}
