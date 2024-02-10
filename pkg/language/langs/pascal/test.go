package pascal

import (
	"github.com/mraron/njudge/pkg/language/memory"
	"github.com/mraron/njudge/pkg/language/sandbox"
	"testing"
	"time"

	"github.com/mraron/njudge/pkg/language"
)

const TestCodeHelloWorld = `begin
    writeln('Hello world');
end.
`

func (p Pascal) Test(t *testing.T, s sandbox.Sandbox) error {
	for _, test := range []language.Test{
		{Name: p.ID() + "_print", Language: p, Source: TestCodeHelloWorld, ExpectedVerdict: sandbox.VerdictOK, ExpectedOutput: "Hello world\n", TimeLimit: 1 * time.Second, MemoryLimit: 128 * memory.MiB},
	} {
		t.Run(test.Name, func(t *testing.T) {
			if err := test.Run(s); err != nil {
				t.Error(err)
			}
		})
	}

	return nil
}
