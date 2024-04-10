package cpp_test

import (
	"bytes"
	"context"
	"github.com/mraron/njudge/pkg/internal/testutils"
	"github.com/mraron/njudge/pkg/language/langs/cpp"
	"github.com/mraron/njudge/pkg/language/memory"
	"github.com/mraron/njudge/pkg/language/sandbox"
	"github.com/stretchr/testify/assert"
	"io"
	"testing"
	"testing/iotest"
	"time"
)

const (
	TestCodeExtraFilesMain = `#include "extra.h"
int main() {
	extra();
}`
	TestCodeExtraHeader = `void extra();`
	TestCodeExtraSource = `#include<iostream>
void extra() {
std::cout<<"Hello world\n";
}
`
	TestCodeCompileError = `#include<iostream>
using namespace std;

int main() {
    int n;
    cin>>n;
    vector<int> t(n);
    cout<<"Hello world!\n";    return 0;
}`
)

func TestExtraFiles(t *testing.T) {
	var (
		s   sandbox.Sandbox
		err error
	)
	if *testutils.UseIsolate {
		s, err = sandbox.NewIsolate(557)
	} else {
		s, err = sandbox.NewDummy()
	}
	if err != nil {
		t.Error(err)
	}

	testcases := []struct {
		name       string
		source     sandbox.File
		extras     []sandbox.File
		wantStdout string
	}{
		{
			name:   "only_header",
			source: sandbox.File{Name: "main.cpp", Source: io.NopCloser(bytes.NewBufferString(TestCodeExtraFilesMain))},
			extras: []sandbox.File{
				{"extra.h", io.NopCloser(bytes.NewBufferString(TestCodeExtraSource))},
			},
			wantStdout: "Hello world\n",
		},
		{
			name:   "source_and_header",
			source: sandbox.File{Name: "main.cpp", Source: io.NopCloser(bytes.NewBufferString(TestCodeExtraFilesMain))},
			extras: []sandbox.File{
				{"extra.h", io.NopCloser(bytes.NewBufferString(TestCodeExtraHeader))},
				{"extra.cpp", io.NopCloser(bytes.NewBufferString(TestCodeExtraSource))},
			},
			wantStdout: "Hello world\n",
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			if err := s.Init(context.TODO()); err != nil {
				t.Error(err)
			}

			bin, err := cpp.Std17.Compile(context.TODO(), s, tc.source, io.Discard, tc.extras)
			if err != nil {
				t.Error(err)
			}

			stdout := &bytes.Buffer{}
			st, err := cpp.Std17.Run(context.TODO(), s, *bin, iotest.ErrReader(io.EOF), stdout, 1*time.Second, 512*memory.MiB)
			if err != nil {
				t.Error(err)
			}

			assert.Equal(t, sandbox.VerdictOK, st.Verdict)
			assert.Equal(t, tc.wantStdout, stdout.String())
		})
	}

}

func TestCompileError(t *testing.T) {
	var (
		s   sandbox.Sandbox
		err error
	)
	if *testutils.UseIsolate {
		s, err = sandbox.NewIsolate(557)
	} else {
		s, err = sandbox.NewDummy()
	}
	if err != nil {
		t.Error(err)
	}

	_ = s.Init(context.TODO())

	file, err := cpp.Std17.Compile(context.Background(), s, sandbox.File{
		Name:   "main.cpp",
		Source: io.NopCloser(bytes.NewBufferString(TestCodeCompileError)),
	}, io.Discard, nil)
	assert.Nil(t, file)
	assert.Nil(t, err)
}
