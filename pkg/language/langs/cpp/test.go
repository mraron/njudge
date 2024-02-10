package cpp

import (
	"github.com/mraron/njudge/pkg/language/memory"
	"github.com/mraron/njudge/pkg/language/sandbox"
	"testing"
	"time"

	"github.com/mraron/njudge/pkg/language"
)

const (
	TestCodeAplusb = `#include<iostream>
using namespace std;
int main() {
	int a,b;
	cin>>a>>b;
	cout<<a+b<<"\n";
}`
	TestCodeCompileError = `#include<lol>
lmao;
int main(a,b,c);
`
	TestCodeHelloWorld = `#include<iostream>
using namespace std;
int main() {
	cout<<"Hello world";
	return 0;
}`
	TestCodeTimeLimit = `#include<iostream>
using namespace std;
int main() {
	int n=0;
	while(1) n++; 
}`
	TestCodeRuntimeError = `#include<iostream>
using namespace std;
void dfs(int x){
	dfs(x+1);
	if(x==int(1e9)) cerr<<"lel\n";
}
int main() {
	dfs(-10000);
}`
	TestCodeRuntimeErrorDiv0 = `#include<iostream>
using namespace std;
int main() {
	cerr<<(1/0);
}`
	TestCodeLongSleep = `#include<unistd.h>
int main() {
	sleep(20);
}
`
	TestCodeShortSleep = `#include<unistd.h>
int main() {
	usleep(100);
}
`
)

func (c Cpp) Test(t *testing.T, s sandbox.Sandbox) error {
	for _, test := range []language.Test{
		{Name: c.ID() + "_latest_aplusb", Language: c, Source: TestCodeAplusb, ExpectedVerdict: sandbox.VerdictOK, Input: "1 2", ExpectedOutput: "3\n", TimeLimit: 1 * time.Second, MemoryLimit: 128 * memory.MiB},
		{Name: c.ID() + "_latest_ce", Language: c, Source: TestCodeCompileError, ExpectedVerdict: sandbox.VerdictCE, TimeLimit: 1 * time.Second, MemoryLimit: 128 * memory.MiB},
		{Name: c.ID() + "_latest_hello", Language: c, Source: TestCodeHelloWorld, ExpectedVerdict: sandbox.VerdictOK, ExpectedOutput: "Hello world", TimeLimit: 1 * time.Second, MemoryLimit: 128 * memory.MiB},
		{Name: c.ID() + "_latest_tl", Language: c, Source: TestCodeTimeLimit, ExpectedVerdict: sandbox.VerdictTL, TimeLimit: 100 * time.Millisecond, MemoryLimit: 128 * memory.MiB},
		{Name: c.ID() + "_latest_retl", Language: c, Source: TestCodeRuntimeError, ExpectedVerdict: sandbox.VerdictRE | sandbox.VerdictTL, TimeLimit: 1000 * time.Millisecond, MemoryLimit: 128 * memory.MiB},
		{Name: c.ID() + "_latest_rediv0", Language: c, Source: TestCodeRuntimeErrorDiv0, ExpectedVerdict: sandbox.VerdictRE, TimeLimit: 1000 * time.Millisecond, MemoryLimit: 128 * memory.MiB},
		{Name: c.ID() + "_latest_slepptl", Language: c, Source: TestCodeLongSleep, ExpectedVerdict: sandbox.VerdictTL, TimeLimit: 100 * time.Millisecond, MemoryLimit: 128 * memory.MiB},
		{Name: c.ID() + "_latest_sleepok", Language: c, Source: TestCodeShortSleep, ExpectedVerdict: sandbox.VerdictOK, TimeLimit: 200 * time.Millisecond, MemoryLimit: 128 * memory.MiB},
	} {
		t.Run(test.Name, func(t *testing.T) {
			if err := test.Run(s); err != nil {
				t.Error(err)
			}
		})
	}

	return nil
}
