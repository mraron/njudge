package golang

import (
	"github.com/mraron/njudge/pkg/language/memory"
	"github.com/mraron/njudge/pkg/language/sandbox"
	"testing"
	"time"

	"github.com/mraron/njudge/pkg/language"
)

const (
	aplusb = `package main
import (
	"fmt"
)
func main() {
	a,b := 0,0
	fmt.Scanf("%d %d", &a, &b)
	fmt.Println(a+b)
}`
	ce = `pkgace main
import (
	"fmt"
)
func main() {
	fmt.Println("lol")
}
`
	print = `package main
import (
	"fmt"
)
func main() {
	fmt.Println("Hello world")
}`
	tl = `package main
func main() {
	a := 0
	for 1==1 {
		a++
	}
}`
	re = `package main
func dfs(x int) {
	dfs(x+1)
	if x==1000000000 {
		return 
	}
}
func main() {
	dfs(-1000)
}`
	rediv0 = `package main
import (
	"fmt"
)
func main() {
	a := 1
	b := 0
	fmt.Println(a/b)
}`
)

func (g golang) Test(t *testing.T, s sandbox.Sandbox) error {
	for _, test := range []language.Test{
		{"golang_aplusb", g, aplusb, sandbox.VerdictOK, "1 2", "3\n", 1 * time.Second, 128 * memory.MiB},
		{"golang_ce", g, ce, sandbox.VerdictCE, "", "", 1 * time.Second, 128 * memory.MiB},
		{"golang_print", g, print, sandbox.VerdictOK, "", "Hello world\n", 1 * time.Second, 128 * memory.MiB},
		{"golang_tl", g, tl, sandbox.VerdictTL, "", "", 100 * time.Millisecond, 128 * memory.MiB},
		{"golang_re", g, re, sandbox.VerdictRE | sandbox.VerdictTL, "", "", 1000 * time.Millisecond, 128 * memory.MiB},
		{"golang_rediv0", g, rediv0, sandbox.VerdictRE, "", "", 1000 * time.Millisecond, 128 * memory.MiB},
	} {
		t.Run(test.Name, func(t *testing.T) {
			if err := test.Run(s); err != nil {
				t.Error(err)
			}
		})
	}

	return nil
}
