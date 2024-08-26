package evaluation_test

import (
	"archive/zip"
	"bytes"
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/mraron/njudge/pkg/internal/testutils"
	"github.com/mraron/njudge/pkg/language/langs/cpp"
	"github.com/mraron/njudge/pkg/language/langs/python3"
	zipLang "github.com/mraron/njudge/pkg/language/langs/zip"
	"github.com/mraron/njudge/pkg/language/sandbox"
	"github.com/mraron/njudge/pkg/problems"
	"github.com/mraron/njudge/pkg/problems/evaluation"
	"github.com/mraron/njudge/pkg/problems/executable/checker"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

func TestBasicRunner_Run(t *testing.T) {
	type args struct {
		ctx             context.Context
		sandboxProvider sandbox.Provider
		testcase        *problems.Testcase
	}
	var (
		s   sandbox.Sandbox
		err error
	)
	if !*testutils.UseIsolate {
		s, err = sandbox.NewDummy()
	} else {
		s, err = sandbox.NewIsolate(444)
	}
	assert.Nil(t, err)

	fs := afero.NewMemMapFs()
	assert.Nil(t, afero.WriteFile(fs, "input", []byte("1 2 3\n"), 0644))
	assert.Nil(t, afero.WriteFile(fs, "answer", []byte("6\n"), 0644))
	tests := []struct {
		name        string
		solution    problems.Solution
		br          *evaluation.BasicRunner
		args        args
		wantErr     bool
		wantVerdict problems.VerdictName
	}{
		{
			name: "stdin_stdout_ac",
			solution: evaluation.NewByteSolution(python3.Python3{}, []byte(`a,b,c = [int(x) for x in input().split()]
print(a+b+c, '   \n')`)),
			br: evaluation.NewBasicRunner(
				evaluation.BasicRunnerWithFs(fs),
				evaluation.BasicRunnerWithChecker(checker.NewWhitediff(checker.WhiteDiffWithFs(fs, afero.NewOsFs()))),
			),
			args: args{
				ctx:             context.TODO(),
				sandboxProvider: sandbox.NewProvider().Put(s),
				testcase: &problems.Testcase{
					InputPath:  "input",
					AnswerPath: "answer",
				},
			},
			wantErr:     false,
			wantVerdict: problems.VerdictAC,
		},
		{
			name: "stdin_stdout_wa",
			solution: evaluation.NewByteSolution(python3.Python3{}, []byte(`a,b,c = [int(x) for x in input().split()]
print(7)`)),
			br: evaluation.NewBasicRunner(
				evaluation.BasicRunnerWithFs(fs),
				evaluation.BasicRunnerWithChecker(checker.NewWhitediff(checker.WhiteDiffWithFs(fs, afero.NewOsFs()))),
			),
			args: args{
				ctx:             context.TODO(),
				sandboxProvider: sandbox.NewProvider().Put(s),
				testcase: &problems.Testcase{
					InputPath:  "input",
					AnswerPath: "answer",
				},
			},
			wantErr:     false,
			wantVerdict: problems.VerdictWA,
		},
		{
			name: "fileinput_stdout_ac",
			solution: evaluation.NewByteSolution(
				python3.Python3{},
				[]byte(`import sys
sys.stdin = open('bemenet', 'r')
a,b,c = [int(x) for x in input().split()]
print(a+b+c, '   \n')`)),
			br: evaluation.NewBasicRunner(
				evaluation.BasicRunnerWithFs(fs),
				evaluation.BasicRunnerWithFiles("bemenet", ""),
				evaluation.BasicRunnerWithChecker(checker.NewWhitediff(checker.WhiteDiffWithFs(fs, afero.NewOsFs()))),
			),
			args: args{
				ctx:             context.TODO(),
				sandboxProvider: sandbox.NewProvider().Put(s),
				testcase: &problems.Testcase{
					InputPath:  "input",
					AnswerPath: "answer",
				},
			},
			wantErr:     false,
			wantVerdict: problems.VerdictAC,
		},
		{
			name: "fileinput_fileoutput_ac",
			solution: evaluation.NewByteSolution(
				python3.Python3{},
				[]byte(`import sys
sys.stdin = open('bemenet', 'r')
a,b,c = [int(x) for x in input().split()]
with open('kimenet', 'w') as w:
	w.write(str(a+b+c)+'   \n')`)),
			br: evaluation.NewBasicRunner(
				evaluation.BasicRunnerWithFs(fs),
				evaluation.BasicRunnerWithFiles("bemenet", "kimenet"),
				evaluation.BasicRunnerWithChecker(checker.NewWhitediff(checker.WhiteDiffWithFs(fs, afero.NewOsFs()))),
			),
			args: args{
				ctx:             context.TODO(),
				sandboxProvider: sandbox.NewProvider().Put(s),
				testcase: &problems.Testcase{
					InputPath:  "input",
					AnswerPath: "answer",
				},
			},
			wantErr:     false,
			wantVerdict: problems.VerdictAC,
		},
		{
			name: "fileoutput_not_created_wa",
			solution: evaluation.NewByteSolution(
				python3.Python3{},
				[]byte(`import sys
sys.stdin = open('bemenet', 'r')
a,b,c = [int(x) for x in input().split()]
print(a+b+c)`)),
			br: evaluation.NewBasicRunner(
				evaluation.BasicRunnerWithFs(fs),
				evaluation.BasicRunnerWithFiles("bemenet", "kimenet"),
				evaluation.BasicRunnerWithChecker(checker.NewWhitediff(checker.WhiteDiffWithFs(fs, afero.NewOsFs()))),
			),
			args: args{
				ctx:             context.TODO(),
				sandboxProvider: sandbox.NewProvider().Put(s),
				testcase: &problems.Testcase{
					InputPath:  "input",
					AnswerPath: "answer",
				},
			},
			wantErr:     false,
			wantVerdict: problems.VerdictWA,
		},
		{
			name: "stdin_fileout_ac",
			solution: evaluation.NewByteSolution(
				python3.Python3{},
				[]byte(`a,b,c = [int(x) for x in input().split()]
with open('kimenet', 'w') as w:
	w.write(str(a+b+c)+'   \n')`)),
			br: evaluation.NewBasicRunner(
				evaluation.BasicRunnerWithFs(fs),
				evaluation.BasicRunnerWithFiles("", "kimenet"),
				evaluation.BasicRunnerWithChecker(checker.NewWhitediff(checker.WhiteDiffWithFs(fs, afero.NewOsFs()))),
			),
			args: args{
				ctx:             context.TODO(),
				sandboxProvider: sandbox.NewProvider().Put(s),
				testcase: &problems.Testcase{
					InputPath:  "input",
					AnswerPath: "answer",
				},
			},
			wantErr:     false,
			wantVerdict: problems.VerdictAC,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Nil(t, tt.br.SetSolution(tt.args.ctx, tt.solution))
			if err := tt.br.Run(tt.args.ctx, tt.args.sandboxProvider, tt.args.testcase); (err != nil) != tt.wantErr {
				t.Errorf("Run() error = %v, wantErr %v", err, tt.wantErr)
			}
			assert.Equal(t, tt.wantVerdict, tt.args.testcase.VerdictName)
		})
	}
}

func TestBasicRunnerMultipleRuns(t *testing.T) {
	var (
		s   sandbox.Sandbox
		err error
	)
	if !*testutils.UseIsolate {
		s, err = sandbox.NewDummy()
	} else {
		s, err = sandbox.NewIsolate(444)
	}
	assert.Nil(t, err)

	fs := afero.NewMemMapFs()
	assert.Nil(t, afero.WriteFile(fs, "input", []byte("1 2 3\n"), 0644))
	assert.Nil(t, afero.WriteFile(fs, "answer", []byte("6\n"), 0644))

	sol := evaluation.NewByteSolution(python3.Python3{}, []byte(`a,b,c = [int(x) for x in input().split()]
print(a+b+c, '   \n')`))
	br := evaluation.NewBasicRunner(
		evaluation.BasicRunnerWithFs(fs),
		evaluation.BasicRunnerWithChecker(checker.NewWhitediff(checker.WhiteDiffWithFs(fs, afero.NewOsFs()))),
	)

	assert.Nil(t, br.SetSolution(context.Background(), sol))
	tc := &problems.Testcase{
		InputPath:  "input",
		AnswerPath: "answer",
	}

	assert.Nil(t, br.Run(context.Background(), sandbox.NewProvider().Put(s), tc))
	assert.Equal(t, problems.VerdictAC, tc.VerdictName)
	tc.VerdictName = problems.VerdictDR
	assert.Nil(t, br.Run(context.Background(), sandbox.NewProvider().Put(s), tc))
	assert.Equal(t, problems.VerdictAC, tc.VerdictName)
}

func TestZipRunner_Run(t *testing.T) {
	fs := afero.NewMemMapFs()
	assert.Nil(t, afero.WriteFile(fs, "input", []byte("hello_world?\n"), 0644))
	assert.Nil(t, afero.WriteFile(fs, "answer", []byte("hello_world!\n"), 0644))

	zipBuf := &bytes.Buffer{}
	w := zip.NewWriter(zipBuf)
	f, _ := w.Create("output1")
	_, _ = f.Write([]byte("hello_world!"))
	_ = w.Close()

	type args struct {
		ctx             context.Context
		sandboxProvider sandbox.Provider
		testcase        *problems.Testcase
	}
	tests := []struct {
		name        string
		solution    problems.Solution
		checker     problems.Checker
		args        args
		wantErr     assert.ErrorAssertionFunc
		wantVerdict problems.VerdictName
	}{
		{
			name:     "zip_hello_world",
			solution: evaluation.NewByteSolution(zipLang.Zip{}, zipBuf.Bytes()),
			checker:  checker.NewWhitediff(checker.WhiteDiffWithFs(fs, afero.NewOsFs())),
			args: args{
				ctx:             context.TODO(),
				sandboxProvider: nil,
				testcase: &problems.Testcase{
					Index:      1,
					InputPath:  "input",
					OutputPath: "output1",
					AnswerPath: "answer",
				},
			},
			wantErr:     assert.NoError,
			wantVerdict: problems.VerdictAC,
		},
		{
			name:     "zip_hello_world_wa",
			solution: evaluation.NewByteSolution(zipLang.Zip{}, zipBuf.Bytes()),
			checker:  checker.NewWhitediff(checker.WhiteDiffWithFs(fs, afero.NewOsFs())),
			args: args{
				ctx:             context.TODO(),
				sandboxProvider: nil,
				testcase: &problems.Testcase{
					Index:      1,
					InputPath:  "input",
					OutputPath: "output1",
					AnswerPath: "input", // WA because of this
				},
			},
			wantErr:     assert.NoError,
			wantVerdict: problems.VerdictWA,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			z := evaluation.NewZipRunner(tt.checker)
			assert.Nil(t, z.SetSolution(context.TODO(), tt.solution))
			tt.wantErr(t, z.Run(tt.args.ctx, tt.args.sandboxProvider, tt.args.testcase), fmt.Sprintf("Run(%v, %v, %v)", tt.args.ctx, tt.args.sandboxProvider, tt.args.testcase))
			assert.Equal(t, tt.wantVerdict, tt.args.testcase.VerdictName)
		})
	}
}

func mustReadFile(name string) []byte {
	b, err := os.ReadFile(name)
	if err != nil {
		panic(err)
	}
	return b
}

func TestInteractiveRunner_Run(t *testing.T) {
	type args struct {
		ctx             context.Context
		sandboxProvider sandbox.Provider
		testcase        *problems.Testcase
	}
	var (
		s1, s2 sandbox.Sandbox
		err    error
	)
	if !*testutils.UseIsolate {
		s1, err = sandbox.NewDummy()
		assert.NoError(t, err)
		s2, err = sandbox.NewDummy()
		assert.NoError(t, err)
	} else {
		s1, err = sandbox.NewIsolate(444)
		assert.NoError(t, err)
		s2, err = sandbox.NewIsolate(445)
		assert.NoError(t, err)
	}

	taskYAMLExecutor := &evaluation.TaskYAMLUserInteractorExecute{}

	fs := afero.NewMemMapFs()
	assert.NoError(t, afero.WriteFile(fs, "input", []byte("11 12\n"), 0644))
	assert.NoError(t, afero.WriteFile(fs, "answer", []byte("23\n"), 0644))
	assert.NoError(t, afero.WriteFile(fs, "empty", []byte("\n"), 0644))

	assert.NoError(t, afero.WriteFile(fs, "input_multi", []byte("1\n11 12\n"), 0644))
	assert.NoError(t, afero.WriteFile(fs, "manager.cpp", mustReadFile("testdata/taskyaml_manager.cpp"), 0644))

	compileSandbox, _ := sandbox.NewDummy()
	assert.NoError(t, cpp.AutoCompile(context.TODO(), fs, compileSandbox, "", "manager.cpp", "manager"))
	managerBinary, err := afero.ReadFile(fs, "manager")
	assert.NoError(t, err)

	tests := []struct {
		name        string
		solution    problems.Solution
		ir          *evaluation.InteractiveRunner
		args        args
		wantErr     assert.ErrorAssertionFunc
		wantVerdict problems.VerdictName
		wantScore   float64
	}{
		{
			name:     "aplusb_python_interactor_polygon",
			solution: evaluation.NewByteSolution(python3.Python3{}, mustReadFile("testdata/aplusb_single.py")),
			ir:       evaluation.NewInteractiveRunner(mustReadFile("testdata/polygon_interactor.py"), checker.NewWhitediff(checker.WhiteDiffWithFs(fs, afero.NewOsFs())), evaluation.InteractiveRunnerWithFs(fs)),
			args: args{
				ctx:             context.TODO(),
				sandboxProvider: sandbox.NewProvider().Put(s1).Put(s2),
				testcase: &problems.Testcase{
					Index:      1,
					InputPath:  "input",
					OutputPath: "output",
					AnswerPath: "answer",
					TimeLimit:  1 * time.Second,
				},
			},
			wantErr:     assert.NoError,
			wantVerdict: problems.VerdictAC,
		},
		{
			name:     "aplusb_python_interactor_taskyaml",
			solution: evaluation.NewByteSolution(python3.Python3{}, mustReadFile("testdata/aplusb_single.py")),
			ir: evaluation.NewInteractiveRunner(
				mustReadFile("testdata/taskyaml_interactor.py"),
				taskYAMLExecutor,
				evaluation.InteractiveRunnerWithExecutor(taskYAMLExecutor),
				evaluation.InteractiveRunnerWithFs(fs),
			),
			args: args{
				ctx:             context.TODO(),
				sandboxProvider: sandbox.NewProvider().Put(s1).Put(s2),
				testcase: &problems.Testcase{
					Index:      1,
					InputPath:  "input",
					OutputPath: "output",
					AnswerPath: "answer",
					MaxScore:   10.0,
					TimeLimit:  1 * time.Second,
				},
			},
			wantErr:     assert.NoError,
			wantVerdict: problems.VerdictAC,
			wantScore:   10.0,
		},
		{
			name:     "aplusb_cpp_interactor_taskyaml",
			solution: evaluation.NewByteSolution(python3.Python3{}, mustReadFile("testdata/aplusb_multi.py")),
			ir: evaluation.NewInteractiveRunner(
				managerBinary,
				taskYAMLExecutor,
				evaluation.InteractiveRunnerWithExecutor(taskYAMLExecutor),
				evaluation.InteractiveRunnerWithFs(fs),
			),
			args: args{
				ctx:             context.TODO(),
				sandboxProvider: sandbox.NewProvider().Put(s1).Put(s2),
				testcase: &problems.Testcase{
					Index:      1,
					InputPath:  "input_multi",
					OutputPath: "output",
					AnswerPath: "answer",
					MaxScore:   10.0,
					TimeLimit:  1 * time.Second,
				},
			},
			wantErr:     assert.NoError,
			wantVerdict: problems.VerdictAC,
			wantScore:   10.0,
		},
		{
			name:     "printalot_polygon",
			solution: evaluation.NewByteSolution(python3.Python3{}, mustReadFile("testdata/empty.py")),
			ir: evaluation.NewInteractiveRunner(
				mustReadFile("testdata/printalot.py"),
				checker.NewWhitediff(checker.WhiteDiffWithFs(fs, afero.NewOsFs())),
				evaluation.InteractiveRunnerWithFs(fs),
			),
			args: args{
				ctx:             context.TODO(),
				sandboxProvider: sandbox.NewProvider().Put(s1).Put(s2),
				testcase: &problems.Testcase{
					Index:      1,
					InputPath:  "input",
					OutputPath: "output",
					AnswerPath: "empty",
					MaxScore:   10.0,
					TimeLimit:  1 * time.Second,
				},
			},
			wantErr:     assert.NoError,
			wantVerdict: problems.VerdictAC,
			wantScore:   10.0,
		},
		{
			name: "interactor_error",
			solution: evaluation.NewByteSolution(python3.Python3{}, mustReadFile("testdata/empty.py")),
			ir: evaluation.NewInteractiveRunner(
				mustReadFile("testdata/error.py"),
				checker.NewWhitediff(checker.WhiteDiffWithFs(fs, afero.NewOsFs())),
				evaluation.InteractiveRunnerWithFs(fs),
			),
			args: args{
				ctx:             context.TODO(),
				sandboxProvider: sandbox.NewProvider().Put(s1).Put(s2),
				testcase: &problems.Testcase{
					Index:      1,
					InputPath:  "input",
					OutputPath: "output",
					AnswerPath: "empty",
					MaxScore:   10.0,
					TimeLimit:  1 * time.Second,
				},
			},
			wantErr:     assert.Error,
			wantVerdict: problems.VerdictXX,
			wantScore:   0.0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.NoError(t, tt.ir.SetSolution(tt.args.ctx, tt.solution))
			tt.wantErr(t, tt.ir.Run(tt.args.ctx, tt.args.sandboxProvider, tt.args.testcase), fmt.Sprintf("Run(%v, %v, %v)", tt.args.ctx, tt.args.sandboxProvider, tt.args.testcase))
			assert.Equal(t, tt.wantVerdict, tt.args.testcase.VerdictName)
			assert.Equal(t, tt.wantScore, tt.args.testcase.Score)
		})
	}
}
