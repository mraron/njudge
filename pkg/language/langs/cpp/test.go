package cpp

import (
	"time"

	"github.com/mraron/njudge/pkg/language"
)

const (
	aplusb = `#include<iostream>
using namespace std;
int main() {
	int a,b;
	cin>>a>>b;
	cout<<a+b<<"\n";
}`
	compilerError = `#include<lol>
lmao;
int main(a,b,c);
`
	print = `#include<iostream>
using namespace std;
int main() {
	cout<<"Hello world";
	return 0;
}`
	timelimitExceeded = `#include<iostream>
using namespace std;
int main() {
	int n=0;
	while(1) n++; 
}`
	runtimeError = `#include<iostream>
using namespace std;
void dfs(int x){
	dfs(x+1);
	if(x==int(1e9)) cerr<<"lel\n";
}
int main() {
	dfs(-10000);
}`
	runtimeErrorDiv0 = `#include<iostream>
using namespace std;
int main() {
	cerr<<(1/0);
}`
	longSleep = `#include<unistd.h>
int main() {
	sleep(20);
}
`
	shortSleep = `#include<unistd.h>
int main() {
	sleep(1);
}
`
)

func (c Cpp) Test(s language.Sandbox) error {
	for _, test := range []language.LanguageTest{
		{latest, aplusb, language.VerdictOK, "1 2", "3\n", 1 * time.Second, 128 * 1024 * 1024},
		{latest, compilerError, language.VerdictCE, "", "", 1 * time.Second, 128 * 1024 * 1024},
		{latest, print, language.VerdictOK, "", "Hello world", 1 * time.Second, 128 * 1024 * 1024},
		{latest, timelimitExceeded, language.VerdictTL, "", "", 100 * time.Millisecond, 128 * 1024 * 1024},
		{latest, runtimeError, language.VerdictRE | language.VerdictTL, "", "", 1000 * time.Millisecond, 128 * 1024 * 1024},
		{latest, runtimeErrorDiv0, language.VerdictRE, "", "", 1000 * time.Millisecond, 128 * 1024 * 1024},
		{latest, longSleep, language.VerdictTL, "", "", 2 * time.Second, 128 * 1024 * 1024},
		{latest, shortSleep, language.VerdictOK, "", "", 2 * time.Second, 128 * 1024 * 1024},
	} {
		if err := test.Run(s); err != nil {
			return err
		}
	}

	return nil
}
