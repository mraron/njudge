package java

import (
	"time"

	"github.com/mraron/njudge/pkg/language"
)

const print = `public class main {
    public static void main(String[] args) {
        System.out.println("Hello world"); 
    }
}`

func (j java) Test(s language.Sandbox) error {
	for _, test := range []language.LanguageTest{
		{Language: j, Source: print, ExpectedVerdict: language.VERDICT_OK, Input: "", ExpectedOutput: "Hello world\n", TimeLimit: 1 * time.Second, MemoryLimit: 50 * 128 * 1024 * 1024},
	} {
		if err := test.Run(s); err != nil {
			return err
		}
	}

	return nil
}
