package cpp11

import (
	"github.com/mraron/njudge/utils/language"
	"testing"
	"time"
)

const (
	CPP11_newfromcpp11 = `#include<iostream>
using namespace std;
int main() {
	cerr<<(10'0110'0)<<"\n";
}`
)

func TestCompileAndRun(t *testing.T) {
	for _, test := range []language.LanguageTest{
		{language.Get("cpp11"), CPP11_newfromcpp11, language.VERDICT_CE, "", "", 1 * time.Second, 128 * 1024 * 1024},
	} {
		test.Run(t)
	}
}
