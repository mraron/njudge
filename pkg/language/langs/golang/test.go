package golang

import (
	"github.com/mraron/njudge/pkg/language/memory"
	"github.com/mraron/njudge/pkg/language/sandbox"
	"testing"
	"time"

	"github.com/mraron/njudge/pkg/language"
)

const (
	TestCodeAplusb = `package main
import (
	"fmt"
)
func main() {
	a,b := 0,0
	fmt.Scanf("%d %d", &a, &b)
	fmt.Println(a+b)
}`
	TestCodeCompilationError = `pkgace main
import (
	"fmt"
)
func main() {
	fmt.Println("lol")
}
`
	TestCodeHelloWorld = `package main
import (
	"fmt"
)
func main() {
	fmt.Println("Hello world")
}`
	TestCodeTimeLimit = `package main
func main() {
	a := 0
	for 1==1 {
		a++
	}
}`
	TestCodeRuntimeError = `package main
func dfs(x int) {
	dfs(x+1)
	if x==1000000000 {
		return 
	}
}
func main() {
	dfs(-1000)
}`
	TestCodeRuntimeErrorDiv0 = `package main
import (
	"fmt"
)
func main() {
	a := 1
	b := 0
	fmt.Println(a/b)
}`
)

func (g Golang) Test(t *testing.T, s sandbox.Sandbox) error {
	for _, test := range []language.Test{
		{Name: g.ID() + "_aplusb", Language: g, Source: TestCodeAplusb, ExpectedVerdict: sandbox.VerdictOK, Input: "1 2", ExpectedOutput: "3\n", TimeLimit: 1 * time.Second, MemoryLimit: 128 * memory.MiB},
		{Name: g.ID() + "_ce", Language: g, Source: TestCodeCompilationError, ExpectedVerdict: sandbox.VerdictCE, TimeLimit: 1 * time.Second, MemoryLimit: 128 * memory.MiB},
		{Name: g.ID() + "_print", Language: g, Source: TestCodeHelloWorld, ExpectedVerdict: sandbox.VerdictOK, ExpectedOutput: "Hello world\n", TimeLimit: 1 * time.Second, MemoryLimit: 128 * memory.MiB},
		{Name: g.ID() + "_tl", Language: g, Source: TestCodeTimeLimit, ExpectedVerdict: sandbox.VerdictTL, TimeLimit: 100 * time.Millisecond, MemoryLimit: 128 * memory.MiB},
		{Name: g.ID() + "_re", Language: g, Source: TestCodeRuntimeError, ExpectedVerdict: sandbox.VerdictRE | sandbox.VerdictTL, TimeLimit: 1000 * time.Millisecond, MemoryLimit: 128 * memory.MiB},
		{Name: g.ID() + "_rediv0", Language: g, Source: TestCodeRuntimeErrorDiv0, ExpectedVerdict: sandbox.VerdictRE, TimeLimit: 1000 * time.Millisecond, MemoryLimit: 128 * memory.MiB},
	} {
		t.Run(test.Name, func(t *testing.T) {
			if err := test.Run(s); err != nil {
				t.Error(err)
			}
		})
	}

	return nil
}
