package java

import (
	"github.com/mraron/njudge/pkg/language/memory"
	"github.com/mraron/njudge/pkg/language/sandbox"
	"testing"
	"time"

	"github.com/mraron/njudge/pkg/language"
)

const (
	TestCodePrintHelloWorld = `public class main {
    public static void main(String[] args) {
        System.out.println("Hello world"); 
    }
}`
	TestCodeDifferentClass = `public class feladat {
    public static void main(String[] args) {
        System.out.println("Hello world"); 
    }
}`
)

func (j *Java) Test(t *testing.T, s sandbox.Sandbox) error {
	for _, test := range []language.Test{
		{Name: j.ID() + "_print", Language: j, Source: TestCodePrintHelloWorld, ExpectedVerdict: sandbox.VerdictOK, Input: "", ExpectedOutput: "Hello world\n", TimeLimit: 1 * time.Second, MemoryLimit: 128 * memory.MiB},
		{Name: j.ID() + "_print2", Language: j, Source: TestCodeDifferentClass, ExpectedVerdict: sandbox.VerdictOK, Input: "", ExpectedOutput: "Hello world\n", TimeLimit: 1 * time.Second, MemoryLimit: 128 * memory.MiB},
	} {
		t.Run(test.Name, func(t *testing.T) {
			if err := test.Run(s); err != nil {
				t.Error(err)
			}
		})
	}

	return nil
}
