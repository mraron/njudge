package java

import (
	"github.com/mraron/njudge/pkg/language/memory"
	"github.com/mraron/njudge/pkg/language/sandbox"
	"testing"
	"time"

	"github.com/mraron/njudge/pkg/language"
)

const print = `public class main {
    public static void main(String[] args) {
        System.out.println("Hello world"); 
    }
}`

func (j *Java) Test(t *testing.T, s sandbox.Sandbox) error {
	for _, test := range []language.Test{
		{Name: "java_print", Language: j, Source: print, ExpectedVerdict: sandbox.VerdictOK, Input: "", ExpectedOutput: "Hello world\n", TimeLimit: 1 * time.Second, MemoryLimit: 128 * memory.MiB},
	} {
		t.Run(test.Name, func(t *testing.T) {
			if err := test.Run(s); err != nil {
				t.Error(err)
			}
		})
	}

	return nil
}
