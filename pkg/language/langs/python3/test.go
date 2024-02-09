package python3

import (
	"github.com/mraron/njudge/pkg/language/memory"
	"github.com/mraron/njudge/pkg/language/sandbox"
	"testing"
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

func (p python3) Test(t *testing.T, s sandbox.Sandbox) error {
	for _, test := range []language.Test{
		{"python3_aplusb", p, aplusb, sandbox.VerdictOK, "1 2", "3\n", 1 * time.Second, 128 * memory.MiB},
		{"python3_re", p, ce, sandbox.VerdictRE, "", "", 1 * time.Second, 128 * memory.MiB},
		{"python3_print", p, print, sandbox.VerdictOK, "", "Hello world\n", 1 * time.Second, 128 * memory.MiB},
		{"python3_tl", p, tl, sandbox.VerdictTL, "", "", 100 * time.Millisecond, 128 * memory.MiB},
		{"python3_re", p, re, sandbox.VerdictRE, "", "", 1000 * time.Millisecond, 128 * memory.MiB},
		{"python3_rediv0", p, rediv0, sandbox.VerdictRE, "", "", 1000 * time.Millisecond, 128 * memory.MiB},
	} {
		t.Run(test.Name, func(t *testing.T) {
			if err := test.Run(s); err != nil {
				t.Error(err)
			}
		})
	}

	return nil
}
