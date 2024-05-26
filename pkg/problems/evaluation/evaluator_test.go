package evaluation_test

import (
	"context"
	"fmt"
	"github.com/mraron/njudge/pkg/language/langs/python3"
	"github.com/mraron/njudge/pkg/language/memory"
	"github.com/mraron/njudge/pkg/language/sandbox"
	"github.com/mraron/njudge/pkg/problems"
	"github.com/mraron/njudge/pkg/problems/evaluation"
	"github.com/mraron/njudge/pkg/problems/executable/checker"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func makeTestcases(testset, group string, tl time.Duration, ml memory.Amount, inputs, answers []string, verdicts []problems.VerdictName) []problems.Testcase {
	var res []problems.Testcase
	for ind := range inputs {
		if len(verdicts) <= ind {
			verdicts = append(verdicts, problems.VerdictDR)
		}
		res = append(res, problems.Testcase{
			Index:       ind + 1,
			InputPath:   inputs[ind],
			AnswerPath:  answers[ind],
			Testset:     testset,
			Group:       group,
			VerdictName: verdicts[ind],
			TimeLimit:   tl,
			MemoryLimit: ml,
		})
	}
	return res
}

func TestLinearEvaluator_Evaluate(t *testing.T) {
	s, _ := sandbox.NewDummy()
	fs := afero.NewMemMapFs()
	assert.Nil(t, afero.WriteFile(fs, "input", []byte("1 2 3\n"), 0644))
	assert.Nil(t, afero.WriteFile(fs, "answer", []byte("6\n"), 0644))

	assert.Nil(t, afero.WriteFile(fs, "input100", []byte("1 2 3\n"), 0644))
	assert.Nil(t, afero.WriteFile(fs, "answer100", []byte("6\n"), 0644))
	assert.Nil(t, afero.WriteFile(fs, "input101", []byte("2 3 3\n"), 0644))
	assert.Nil(t, afero.WriteFile(fs, "answer101", []byte("8\n"), 0644))
	assert.Nil(t, afero.WriteFile(fs, "input102", []byte("1 2 10\n"), 0644))
	assert.Nil(t, afero.WriteFile(fs, "answer102", []byte("13\n"), 0644)) //wrong answer
	type args struct {
		ctx              context.Context
		skeleton         problems.Status
		compiledSolution problems.Solution
		sandboxProvider  sandbox.Provider
		statusUpdater    problems.StatusUpdater
	}
	tests := []struct {
		name      string
		evaluator problems.Evaluator
		args      args
		want      []problems.VerdictName
		wantErr   assert.ErrorAssertionFunc
	}{
		{
			name:      "acrunner",
			evaluator: evaluation.NewLinearEvaluator(evaluation.ACRunner{}),
			args: args{
				ctx: context.TODO(),
				skeleton: problems.Status{
					FeedbackType: problems.FeedbackACM,
					Feedback: []problems.Testset{

						{
							Name: "testset23",
							Groups: []problems.Group{
								{
									Name:    "group1",
									Scoring: problems.ScoringSum,
									Testcases: makeTestcases(
										"testset23", "group1", 0, 0,
										[]string{"", "", ""},
										[]string{"", "", ""},
										nil,
									),
									Dependencies: nil,
								},
							},
						},
					},
				},
				compiledSolution: evaluation.NewByteSolution(python3.Python3{}, []byte(`print("Hello world")`)),
				sandboxProvider:  sandbox.NewProvider().Put(s),
				statusUpdater:    evaluation.IgnoreStatusUpdate{},
			},
			want:    []problems.VerdictName{problems.VerdictAC, problems.VerdictAC, problems.VerdictAC},
			wantErr: assert.NoError,
		},
		{
			name: "basicrunner",
			evaluator: evaluation.NewLinearEvaluator(
				evaluation.NewBasicRunner(
					evaluation.BasicRunnerWithFs(fs),
					evaluation.BasicRunnerWithChecker(checker.NewWhitediff(checker.WhiteDiffWithFs(fs, afero.NewOsFs()))),
				),
			),
			args: args{
				ctx: context.TODO(),
				skeleton: problems.Status{
					FeedbackType: problems.FeedbackACM,
					Feedback: []problems.Testset{
						{
							Name: "testset23",
							Groups: []problems.Group{
								{
									Name:    "group1",
									Scoring: problems.ScoringSum,
									Testcases: makeTestcases(
										"testset23", "group1", 0, 0,
										[]string{"input", "input", "input"},
										[]string{"answer", "answer", "answer"},
										nil,
									),
									Dependencies: nil,
								},
							},
						},
					},
				},
				compiledSolution: evaluation.NewByteSolution(python3.Python3{}, []byte(`print("Hello world")`)),
				sandboxProvider:  sandbox.NewProvider().Put(s),
				statusUpdater:    evaluation.IgnoreStatusUpdate{},
			},
			want:    []problems.VerdictName{problems.VerdictWA, problems.VerdictSK, problems.VerdictSK},
			wantErr: assert.NoError,
		},
		{
			name: "basicrunner_acwa",
			evaluator: evaluation.NewLinearEvaluator(
				evaluation.NewBasicRunner(
					evaluation.BasicRunnerWithFs(fs),
					evaluation.BasicRunnerWithChecker(checker.NewWhitediff(checker.WhiteDiffWithFs(fs, afero.NewOsFs()))),
				),
			),
			args: args{
				ctx: context.TODO(),
				skeleton: problems.Status{
					FeedbackType: problems.FeedbackACM,
					Feedback: []problems.Testset{
						{
							Name: "testset23",
							Groups: []problems.Group{
								{
									Name:    "group1",
									Scoring: problems.ScoringSum,
									Testcases: makeTestcases("testset23", "group1", 0, 0,
										[]string{"input100", "input101", "input102"},
										[]string{"answer100", "answer101", "answer102"},
										nil,
									),
									Dependencies: nil,
								},
							},
						},
					},
				},
				compiledSolution: evaluation.NewByteSolution(python3.Python3{}, []byte(`a,b,c = map(int, input().split())
print(a+b+3)`)),
				sandboxProvider: sandbox.NewProvider().Put(s),
				statusUpdater:   evaluation.IgnoreStatusUpdate{},
			},
			want:    []problems.VerdictName{problems.VerdictAC, problems.VerdictAC, problems.VerdictWA},
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			le := tt.evaluator
			got, err := le.Evaluate(tt.args.ctx, tt.args.skeleton, tt.args.compiledSolution, tt.args.sandboxProvider, tt.args.statusUpdater)
			if !tt.wantErr(t, err, fmt.Sprintf("Evaluate(%v, %v, %v, %v, %v)", tt.args.ctx, tt.args.skeleton, tt.args.compiledSolution, tt.args.sandboxProvider, tt.args.statusUpdater)) {
				return
			}
			verdictsWant := tt.want
			statusOrig := tt.args.skeleton
			testcasesOrig := statusOrig.Feedback[0].Testcases()
			testcasesGot := got.Feedback[0].Testcases()
			for ind := range verdictsWant {
				assert.Equal(t, verdictsWant[ind], testcasesGot[ind].VerdictName)

				assert.Equal(t, testcasesOrig[ind].Index, testcasesGot[ind].Index)
				assert.Equal(t, testcasesOrig[ind].Group, testcasesGot[ind].Group)
				assert.Equal(t, testcasesOrig[ind].MaxScore, testcasesGot[ind].MaxScore)
				assert.Equal(t, testcasesOrig[ind].InputPath, testcasesGot[ind].InputPath)
				assert.Equal(t, testcasesOrig[ind].AnswerPath, testcasesGot[ind].AnswerPath)
				assert.Equal(t, testcasesOrig[ind].TimeLimit, testcasesGot[ind].TimeLimit)
				assert.Equal(t, testcasesOrig[ind].MemoryLimit, testcasesGot[ind].MemoryLimit)
			}
		})
	}
}
