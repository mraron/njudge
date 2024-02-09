package pascal

import (
	"github.com/mraron/njudge/pkg/language/memory"
	"github.com/mraron/njudge/pkg/language/sandbox"
	"testing"
	"time"

	"github.com/mraron/njudge/pkg/language"
)

const print = `begin
    writeln('Hello world');
end.
`

func (p pascal) Test(t *testing.T, s sandbox.Sandbox) error {
	for _, test := range []language.Test{
		{"pascal_print", p, print, sandbox.VerdictOK, "", "Hello world\n", 1 * time.Second, 128 * memory.MiB},
	} {
		t.Run(test.Name, func(t *testing.T) {
			if err := test.Run(s); err != nil {
				t.Error(err)
			}
		})
	}

	return nil
}
