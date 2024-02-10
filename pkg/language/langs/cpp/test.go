package cpp

import (
	"github.com/mraron/njudge/pkg/language/memory"
	"github.com/mraron/njudge/pkg/language/sandbox"
	"testing"
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
	usleep(100);
}
`
)

func (c Cpp) Test(t *testing.T, s sandbox.Sandbox) error {
	for _, test := range []language.Test{
		{c.Id() + "_latest_aplusb", c, aplusb, sandbox.VerdictOK, "1 2", "3\n", 1 * time.Second, 128 * memory.MiB},
		{c.Id() + "_latest_ce", c, compilerError, sandbox.VerdictCE, "", "", 1 * time.Second, 128 * memory.MiB},
		{c.Id() + "_latest_hello", c, print, sandbox.VerdictOK, "", "Hello world", 1 * time.Second, 128 * memory.MiB},
		{c.Id() + "_latest_tl", c, timelimitExceeded, sandbox.VerdictTL, "", "", 100 * time.Millisecond, 128 * memory.MiB},
		{c.Id() + "_latest_retl", c, runtimeError, sandbox.VerdictRE | sandbox.VerdictTL, "", "", 1000 * time.Millisecond, 128 * memory.MiB},
		{c.Id() + "_latest_rediv0", c, runtimeErrorDiv0, sandbox.VerdictRE, "", "", 1000 * time.Millisecond, 128 * memory.MiB},
		{c.Id() + "_latest_slepptl", c, longSleep, sandbox.VerdictTL, "", "", 100 * time.Millisecond, 128 * memory.MiB},
		{c.Id() + "_latest_sleepok", c, shortSleep, sandbox.VerdictOK, "", "", 200 * time.Millisecond, 128 * memory.MiB},
	} {
		t.Run(test.Name, func(t *testing.T) {
			if err := test.Run(s); err != nil {
				t.Error(err)
			}
		})
	}

	return nil
}
