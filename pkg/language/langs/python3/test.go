package python3

import (
	"github.com/mraron/njudge/pkg/language/memory"
	"github.com/mraron/njudge/pkg/language/sandbox"
	"testing"
	"time"

	"github.com/mraron/njudge/pkg/language"
)

const (
	TestCodeAplusb = `inp = input().split(' ')
a,b = int(inp[0]), int(inp[1]) 
print(a+b)`
	TestCodeSyntaxError = `inp = input(()`
	TestCodeHelloWorld  = `print("Hello world")`
	TestCodeTimeLimit   = `x = 0
while True:
	x = x+1`
	TestCodeRuntimeError = `x = [1,2,3]
print(x[4])`
	TestCodeRuntimeErrorDiv0 = `print(1/0)`
)

func (p python3) Test(t *testing.T, s sandbox.Sandbox) error {
	for _, test := range []language.Test{
		{Name: p.ID() + "_aplusb", Language: p, Source: TestCodeAplusb, ExpectedVerdict: sandbox.VerdictOK, Input: "1 2", ExpectedOutput: "3\n", TimeLimit: 1 * time.Second, MemoryLimit: 128 * memory.MiB},
		{Name: p.ID() + "_re", Language: p, Source: TestCodeSyntaxError, ExpectedVerdict: sandbox.VerdictRE, TimeLimit: 1 * time.Second, MemoryLimit: 128 * memory.MiB},
		{Name: p.ID() + "_print", Language: p, Source: TestCodeHelloWorld, ExpectedVerdict: sandbox.VerdictOK, ExpectedOutput: "Hello world\n", TimeLimit: 1 * time.Second, MemoryLimit: 128 * memory.MiB},
		{Name: p.ID() + "_tl", Language: p, Source: TestCodeTimeLimit, ExpectedVerdict: sandbox.VerdictTL, TimeLimit: 100 * time.Millisecond, MemoryLimit: 128 * memory.MiB},
		{Name: p.ID() + "_re", Language: p, Source: TestCodeRuntimeError, ExpectedVerdict: sandbox.VerdictRE, TimeLimit: 1000 * time.Millisecond, MemoryLimit: 128 * memory.MiB},
		{Name: p.ID() + "_rediv0", Language: p, Source: TestCodeRuntimeErrorDiv0, ExpectedVerdict: sandbox.VerdictRE, TimeLimit: 1000 * time.Millisecond, MemoryLimit: 128 * memory.MiB},
	} {
		t.Run(test.Name, func(t *testing.T) {
			if err := test.Run(s); err != nil {
				t.Error(err)
			}
		})
	}

	return nil
}
