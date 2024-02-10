package cython3

import (
	"github.com/mraron/njudge/pkg/language/langs/python3"
	"github.com/mraron/njudge/pkg/language/memory"
	"github.com/mraron/njudge/pkg/language/sandbox"
	"testing"
	"time"

	"github.com/mraron/njudge/pkg/language"
)

func (c Cython3) Test(t *testing.T, s sandbox.Sandbox) error {
	for _, test := range []language.Test{
		{Name: c.ID() + "_aplusb", Language: c, Source: python3.TestCodeAplusb, ExpectedVerdict: sandbox.VerdictOK, Input: "1 2\n", ExpectedOutput: "3\n", TimeLimit: 1500 * time.Millisecond, MemoryLimit: 128 * memory.MiB},
		{Name: c.ID() + "_ce", Language: c, Source: python3.TestCodeSyntaxError, ExpectedVerdict: sandbox.VerdictCE, TimeLimit: 1 * time.Second, MemoryLimit: 128 * memory.MiB},
		{Name: c.ID() + "_print", Language: c, Source: python3.TestCodeHelloWorld, ExpectedVerdict: sandbox.VerdictOK, ExpectedOutput: "Hello world\n", TimeLimit: 1 * time.Second, MemoryLimit: 128 * memory.MiB},
		{Name: c.ID() + "_tl", Language: c, Source: python3.TestCodeTimeLimit, ExpectedVerdict: sandbox.VerdictTL, TimeLimit: 100 * time.Millisecond, MemoryLimit: 128 * memory.MiB},
		{Name: c.ID() + "_re", Language: c, Source: python3.TestCodeRuntimeError, ExpectedVerdict: sandbox.VerdictRE, TimeLimit: 1000 * time.Millisecond, MemoryLimit: 128 * memory.MiB},
		{Name: c.ID() + "_rediv0", Language: c, Source: python3.TestCodeRuntimeErrorDiv0, ExpectedVerdict: sandbox.VerdictRE, TimeLimit: 1000 * time.Millisecond, MemoryLimit: 128 * memory.MiB},
	} {
		t.Run(test.Name, func(t *testing.T) {
			if err := test.Run(s); err != nil {
				t.Error(err)
			}
		})
	}

	return nil
}
