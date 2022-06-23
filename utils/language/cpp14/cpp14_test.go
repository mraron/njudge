package cpp14

import (
	"github.com/mraron/njudge/utils/language"
	"testing"
	"time"
)

const (
	CPP14_aplusb = `#include<iostream>
using namespace std;
int main() {
	int a,b;
	cin>>a>>b;
	cout<<a+b<<"\n";
}`
	CPP14_ce = `#include<lol>
lmao;
int main(a,b,c);
`
	CPP14_print = `#include<iostream>
using namespace std;
int main() {
	cout<<"Hello world";
	return 0;
}`
	CPP14_tl = `#include<iostream>
using namespace std;
int main() {
	int n=0;
	while(1) n++; 
}`
	CPP14_re = `#include<iostream>
using namespace std;
void dfs(int x){
	dfs(x+1);
	if(x==int(1e9)) cerr<<"lel\n";
}
int main() {
	dfs(-10000);
}`
	CPP14_rediv0 = `#include<iostream>
using namespace std;
int main() {
	cerr<<(1/0);
}`
	CPP14_newfromcpp11 = `#include<iostream>
using namespace std;
int main() {
	cerr<<(10'0110'0)<<"\n";
}`
	CPP14_sleep = `#include<unistd.h>
int main() {
	sleep(20);
}
`
	CPP14_smallsleep = `#include<unistd.h>
int main() {
	sleep(1);
}
`
)

func TestCompileAndRun(t *testing.T) {
	for _, test := range []language.LanguageTest{
		{language.Get("cpp14"), CPP14_aplusb, language.VERDICT_OK, "1 2", "3\n", 1 * time.Second, 128 * 1024 * 1024},
		{language.Get("cpp14"), CPP14_ce, language.VERDICT_CE, "", "", 1 * time.Second, 128 * 1024 * 1024},
		{language.Get("cpp14"), CPP14_print, language.VERDICT_OK, "", "Hello world", 1 * time.Second, 128 * 1024 * 1024},
		{language.Get("cpp14"), CPP14_tl, language.VERDICT_TL, "", "", 100 * time.Millisecond, 128 * 1024 * 1024},
		{language.Get("cpp14"), CPP14_re, language.VERDICT_RE, "", "", 1000 * time.Millisecond, 128 * 1024 * 1024},
		{language.Get("cpp14"), CPP14_rediv0, language.VERDICT_RE, "", "", 1000 * time.Millisecond, 128 * 1024 * 1024},
		{language.Get("cpp14"), CPP14_newfromcpp11, language.VERDICT_OK, "", "", 100 * time.Millisecond, 128 * 1024 * 1024},
		{language.Get("cpp14"), CPP14_sleep, language.VERDICT_TL, "", "", 2 * time.Second, 128 * 1024 * 1024},
		{language.Get("cpp14"), CPP14_smallsleep, language.VERDICT_OK, "", "", 2 * time.Second, 128 * 1024 * 1024},
	} {
		test.Run(t)
	}
}
