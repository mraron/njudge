package evaluation_test

import (
	"context"
	"github.com/mraron/njudge/pkg/language/langs/python3"
	"github.com/mraron/njudge/pkg/language/sandbox"
	"github.com/mraron/njudge/pkg/problems"
	"github.com/mraron/njudge/pkg/problems/evaluation"
	"github.com/mraron/njudge/pkg/problems/executable/checker"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBasicRunner_Run(t *testing.T) {
	type args struct {
		ctx             context.Context
		sandboxProvider sandbox.Provider
		testcase        *problems.Testcase
	}
	s, err := sandbox.NewDummy()
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
				sandboxProvider: sandbox.NewSandboxProvider().Put(s),
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
				sandboxProvider: sandbox.NewSandboxProvider().Put(s),
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
				sandboxProvider: sandbox.NewSandboxProvider().Put(s),
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
				sandboxProvider: sandbox.NewSandboxProvider().Put(s),
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
				sandboxProvider: sandbox.NewSandboxProvider().Put(s),
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
				sandboxProvider: sandbox.NewSandboxProvider().Put(s),
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
