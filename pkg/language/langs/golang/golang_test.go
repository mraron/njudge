package golang

import (
	"github.com/mraron/njudge/pkg/language"
	"github.com/mraron/njudge/pkg/language/sandbox"
	"testing"
	"time"
)

const (
	GOLANG_aplusb = `package main
import (
	"fmt"
)
func main() {
	a,b := 0,0
	fmt.Scanf("%d %d", &a, &b)
	fmt.Println(a+b)
}`
	GOLANG_ce = `pkgace main
import (
	"fmt"
)
func main() {
	fmt.Println("lol")
}
`
	GOLANG_print = `package main
import (
	"fmt"
)
func main() {
	fmt.Println("Hello world")
}`
	GOLANG_tl = `package main
func main() {
	a := 0
	for 1==1 {
		a++
	}
}`
	GOLANG_re = `package main
func dfs(x int) {
	dfs(x+1)
	if x==1000000000 {
		return 
	}
}
func main() {
	dfs(-1000)
}`
	GOLANG_rediv0 = `package main
import (
	"fmt"
)
func main() {
	a := 1
	b := 0
	fmt.Println(a/b)
}`
)

func TestCompileAndRun(t *testing.T) {
	for _, test := range []language.LanguageTest{
		{sandbox.NewDummy(), language.Get("golang"), GOLANG_aplusb, language.VERDICT_OK, "1 2", "3\n", 1 * time.Second, 2 * 512 * 1024 * 1024},
		{sandbox.NewDummy(), language.Get("golang"), GOLANG_ce, language.VERDICT_CE, "", "", 1 * time.Second, 2 * 512 * 1024 * 1024},
		{sandbox.NewDummy(), language.Get("golang"), GOLANG_print, language.VERDICT_OK, "", "Hello world\n", 1 * time.Second, 2 * 512 * 1024 * 1024},
		{sandbox.NewDummy(), language.Get("golang"), GOLANG_tl, language.VERDICT_TL, "", "", 100 * time.Millisecond, 2 * 512 * 1024 * 1024},
		{sandbox.NewDummy(), language.Get("golang"), GOLANG_re, language.VERDICT_RE | language.VERDICT_TL, "", "", 1000 * time.Millisecond, 2 * 512 * 1024 * 1024},
		{sandbox.NewDummy(), language.Get("golang"), GOLANG_rediv0, language.VERDICT_RE, "", "", 1000 * time.Millisecond, 2 * 512 * 1024 * 1024},
	} {
		test.Run(t)
	}
}
