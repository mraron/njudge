package python3

import (
	"time"

	"github.com/mraron/njudge/pkg/language"
)

const (
	aplusb = `inp = input().split(' ')
a,b = int(inp[0]), int(inp[1]) 
print(a+b)`
	ce    = `inp = input(()`
	print = `print("Hello world")`
	tl    = `x = 0
while True:
	x = x+1`
	re = `x = [1,2,3]
print(x[4])`
	rediv0 = `print(1/0)`
)

func (p python3) Test(s language.Sandbox) error {
	for _, test := range []language.LanguageTest{
		{p, aplusb, language.VERDICT_OK, "1 2", "3\n", 1 * time.Second, 128 * 1024 * 1024},
		{p, ce, language.VERDICT_RE, "", "", 1 * time.Second, 128 * 1024 * 1024},
		{p, print, language.VERDICT_OK, "", "Hello world\n", 1 * time.Second, 128 * 1024 * 1024},
		{p, tl, language.VERDICT_TL, "", "", 100 * time.Millisecond, 128 * 1024 * 1024},
		{p, re, language.VERDICT_RE, "", "", 1000 * time.Millisecond, 128 * 1024 * 1024},
		{p, rediv0, language.VERDICT_RE, "", "", 1000 * time.Millisecond, 128 * 1024 * 1024},
	} {
		if err := test.Run(s); err != nil {
			return err
		}
	}

	return nil
}
