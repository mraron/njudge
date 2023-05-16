package golang

import (
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

func (g golang) Test(s language.Sandbox) error {
	for _, test := range []language.LanguageTest{
		{g, aplusb, language.VerdictOK, "1 2", "3\n", 1 * time.Second, 2 * 512 * 1024 * 1024},
		{g, ce, language.VerdictCE, "", "", 1 * time.Second, 2 * 512 * 1024 * 1024},
		{g, print, language.VerdictOK, "", "Hello world\n", 1 * time.Second, 2 * 512 * 1024 * 1024},
		{g, tl, language.VerdictTL, "", "", 100 * time.Millisecond, 2 * 512 * 1024 * 1024},
		{g, re, language.VerdictRE | language.VerdictTL, "", "", 1000 * time.Millisecond, 2 * 512 * 1024 * 1024},
		{g, rediv0, language.VerdictRE, "", "", 1000 * time.Millisecond, 2 * 512 * 1024 * 1024},
	} {
		if err := test.Run(s); err != nil {
			return err
		}
	}

	return nil
}
