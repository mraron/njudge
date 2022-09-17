package pypy3

import (
	"github.com/mraron/njudge/pkg/language"
	"github.com/mraron/njudge/pkg/language/sandbox"
	"testing"
	"time"
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

func TestCompileAndRun(t *testing.T) {
	for _, test := range []language.LanguageTest{
		{sandbox.NewDummy(), language.Get("pypy3"), PYTHON3_aplusb, language.VERDICT_OK, "1 2", "3\n", 1 * time.Second, 128 * 1024 * 1024},
		{sandbox.NewDummy(), language.Get("pypy3"), PYTHON3_ce, language.VERDICT_RE, "", "", 1 * time.Second, 128 * 1024 * 1024},
		{sandbox.NewDummy(), language.Get("pypy3"), PYTHON3_print, language.VERDICT_OK, "", "Hello world\n", 1 * time.Second, 128 * 1024 * 1024},
		{sandbox.NewDummy(), language.Get("pypy3"), PYTHON3_tl, language.VERDICT_TL, "", "", 100 * time.Millisecond, 128 * 1024 * 1024},
		{sandbox.NewDummy(), language.Get("pypy3"), PYTHON3_re, language.VERDICT_RE, "", "", 1000 * time.Millisecond, 128 * 1024 * 1024},
		{sandbox.NewDummy(), language.Get("pypy3"), PYTHON3_rediv0, language.VERDICT_RE, "", "", 1000 * time.Millisecond, 128 * 1024 * 1024},
	} {
		test.Run(t)
	}
}
