package csharp

import (
	"time"

	"github.com/mraron/njudge/pkg/language"
)

const print = `class Hello {         
	static void Main(string[] args)
	{
		System.Console.WriteLine("Hello world");
	}
}`

func (c csharp) Test(s language.Sandbox) error {
	for _, test := range []language.LanguageTest{
		{Language: c, Source: print, ExpectedVerdict: language.VerdictOK, Input: "", ExpectedOutput: "Hello world\n", TimeLimit: 1 * time.Second, MemoryLimit: 50 * 128 * 1024 * 1024},
	} {
		if err := test.Run(s); err != nil {
			return err
		}
	}

	return nil
}
