package csharp

import (
	"github.com/mraron/njudge/pkg/language/memory"
	"github.com/mraron/njudge/pkg/language/sandbox"
	"testing"
	"time"

	"github.com/mraron/njudge/pkg/language"
)

const TestCodeHelloWorld = `class Hello {         
	static void Main(string[] args)
	{
		System.Console.WriteLine("Hello world");
	}
}`

func (c csharp) Test(t *testing.T, s sandbox.Sandbox) error {
	for _, test := range []language.Test{
		{Name: c.ID() + "_print", Language: c, Source: TestCodeHelloWorld, ExpectedVerdict: sandbox.VerdictOK, Input: "", ExpectedOutput: "Hello world\n", TimeLimit: 1 * time.Second, MemoryLimit: 128 * memory.MiB},
	} {
		t.Run(test.Name, func(t *testing.T) {
			if err := test.Run(s); err != nil {
				t.Error(err)
			}
		})
	}

	return nil
}
