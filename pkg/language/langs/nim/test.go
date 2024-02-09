package nim

import (
	"github.com/mraron/njudge/pkg/language/memory"
	"github.com/mraron/njudge/pkg/language/sandbox"
	"testing"
	"time"

	"github.com/mraron/njudge/pkg/language"
)

const print = `echo "Hello world"`

func (n nim) Test(t *testing.T, s sandbox.Sandbox) error {
	for _, test := range []language.Test{
		{Name: "nim_print", Language: n, Source: print, ExpectedVerdict: sandbox.VerdictOK, Input: "", ExpectedOutput: "Hello world\n", TimeLimit: 1 * time.Second, MemoryLimit: 128 * memory.MiB},
	} {
		t.Run(test.Name, func(t *testing.T) {
			if err := test.Run(s); err != nil {
				t.Error(err)
			}
		})
	}

	return nil
}
