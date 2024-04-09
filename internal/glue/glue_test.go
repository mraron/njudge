package glue_test

import (
	"context"
	"errors"
	"github.com/mraron/njudge/internal/glue"
	"github.com/mraron/njudge/internal/judge"
	"github.com/mraron/njudge/internal/njudge"
	"github.com/mraron/njudge/internal/njudge/memory"
	"github.com/mraron/njudge/pkg/language"
	"github.com/mraron/njudge/pkg/language/sandbox"
	"github.com/mraron/njudge/pkg/problems"
	"github.com/stretchr/testify/assert"
	"github.com/volatiletech/null/v8"
	"testing"
	"time"
)

type TestJudger struct {
	f func(ctx context.Context, sub judge.Submission, callback judge.ResultCallback) (*problems.Status, error)
}

func (t TestJudger) Judge(ctx context.Context, sub judge.Submission, callback judge.ResultCallback) (*problems.Status, error) {
	return t.f(ctx, sub, callback)
}

type FakeProblems struct {
	Problem njudge.Problem
}

func (f FakeProblems) Get(ctx context.Context, ID int) (*njudge.Problem, error) {
	return &f.Problem, nil
}

func (f FakeProblems) GetAll(ctx context.Context) ([]njudge.Problem, error) {
	panic("implement me")
}

func (f FakeProblems) Insert(ctx context.Context, p njudge.Problem) (*njudge.Problem, error) {
	panic("implement me")
}

func (f FakeProblems) Delete(ctx context.Context, ID int) error {
	panic("implement me")
}

func (f FakeProblems) Update(ctx context.Context, p njudge.Problem, fields []string) error {
	panic("implement me")
}

func TestGlue_ProcessSubmission(t *testing.T) {
	type fields struct {
		Judge       judge.Judger
		Submissions njudge.Submissions
		Problems    njudge.Problems
	}
	type args struct {
		ctx context.Context
		sub njudge.Submission
	}
	tests := []struct {
		name        string
		fields      fields
		args        args
		wantErr     bool
		wantVerdict njudge.Verdict
		wantScore   float32
	}{
		{
			name: "compilation error",
			fields: fields{
				Judge: TestJudger{
					f: func(ctx context.Context, sub judge.Submission, callback judge.ResultCallback) (*problems.Status, error) {
						return &problems.Status{
							Compiled:       false,
							CompilerOutput: "compilation error",
						}, nil
					},
				},
				Submissions: memory.NewSubmissions(),
				Problems: FakeProblems{Problem: njudge.Problem{
					Problem: "aplusb",
				}},
			},
			args: args{
				ctx: context.Background(),
				sub: njudge.Submission{
					ID: 1,
				},
			},
			wantErr:     false,
			wantVerdict: njudge.VerdictCE,
		},
		{
			name: "error while running",
			fields: fields{
				Judge: TestJudger{f: func(ctx context.Context, sub judge.Submission, callback judge.ResultCallback) (*problems.Status, error) {
					_ = callback(judge.Result{
						Index: 1,
						Test:  "test1",
						Status: &problems.Status{
							Compiled:       true,
							CompilerOutput: "",
							FeedbackType:   problems.FeedbackIOI,
							Feedback: []problems.Testset{
								{
									Name: "tests",
									Groups: []problems.Group{
										{
											Name:    "group1",
											Scoring: problems.ScoringMin,
											Testcases: []problems.Testcase{
												{
													Index:       1,
													Testset:     "tests",
													Group:       "group1",
													VerdictName: problems.VerdictAC,
													Score:       5,
													MaxScore:    5,
												},
											},
											Dependencies: nil,
										},
									},
								},
							},
						},
						Error: "",
					})
					return nil, errors.New("test error")
				}},
				Submissions: memory.NewSubmissions(),
				Problems: FakeProblems{Problem: njudge.Problem{
					Problem: "aplusb",
				}},
			},
			args: args{
				ctx: context.TODO(),
				sub: njudge.Submission{
					ID: 1,
				},
			},
			wantErr:     true,
			wantVerdict: njudge.VerdictRU,
			wantScore:   0.0,
		},
		{
			name: "no callback run RE",
			fields: fields{
				Judge: TestJudger{f: func(ctx context.Context, sub judge.Submission, callback judge.ResultCallback) (*problems.Status, error) {
					return &problems.Status{
						Compiled:       true,
						CompilerOutput: "",
						FeedbackType:   problems.FeedbackIOI,
						Feedback: []problems.Testset{
							{
								Name: "tests",
								Groups: []problems.Group{
									{
										Name:    "group1",
										Scoring: problems.ScoringMin,
										Testcases: []problems.Testcase{
											{
												Index:       1,
												Testset:     "tests",
												Group:       "group1",
												VerdictName: problems.VerdictRE,
												Score:       0,
												MaxScore:    5,
											},
										},
										Dependencies: nil,
									},
								},
							},
						},
					}, nil
				}},
				Submissions: memory.NewSubmissions(),
				Problems: FakeProblems{Problem: njudge.Problem{
					Problem: "aplusb",
				}},
			},
			args: args{
				ctx: context.TODO(),
				sub: njudge.Submission{
					ID: 1,
				},
			},
			wantErr:     false,
			wantVerdict: njudge.VerdictRE,
			wantScore:   0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &glue.Glue{
				Judge:            tt.fields.Judge,
				Submissions:      tt.fields.Submissions,
				Problems:         tt.fields.Problems,
				SubmissionsQuery: memory.NewSubmissionsQuery(tt.fields.Submissions),
			}

			_, _ = tt.fields.Submissions.Insert(tt.args.ctx, tt.args.sub)

			if err := g.ProcessSubmission(tt.args.ctx, tt.args.sub); (err != nil) != tt.wantErr {
				t.Errorf("ProcessSubmission() error = %v, wantErr %v", err, tt.wantErr)
			}
			if s, err := tt.fields.Submissions.Get(context.TODO(), 1); tt.wantVerdict != s.Verdict || tt.wantScore != s.Score || err != nil {
				t.Errorf("ProcessSubmission() verdict = %v, wantVerdict %v (err = %v)", s.Verdict, tt.wantVerdict, err)
			}
		})
	}
}

func TestJudgeIntegration(t *testing.T) {
	s1, _ := sandbox.NewDummy()
	s2, _ := sandbox.NewDummy()
	store := problems.NewFsStore("../judge/testdata")
	_ = store.UpdateProblems()

	judge := &judge.Judge{
		SandboxProvider: sandbox.NewProvider().Put(s1).Put(s2),
		ProblemStore:    store,
		LanguageStore:   language.DefaultStore,
		RateLimit:       0,
	}

	probsMem := memory.NewProblems()
	aplusb, _ := probsMem.Insert(context.TODO(), njudge.Problem{
		Problemset: "main",
		Problem:    "aplusb",
	})

	g, _ := glue.New(judge)
	g.Submissions = memory.NewSubmissions()
	g.Problems = probsMem
	g.SubmissionsQuery = memory.NewSubmissionsQuery(g.Submissions)
	sub, _ := g.Submissions.Insert(context.TODO(), njudge.Submission{
		ID:        0,
		UserID:    1,
		ProblemID: aplusb.ID,
		Language:  "python3",
		Source: []byte(`a, b = list(map(int, input().split()))
print(a+b)
`),
		Private:   false,
		Started:   false,
		Verdict:   njudge.VerdictUP,
		Ontest:    null.String{},
		Submitted: time.Time{},
		Status:    problems.Status{},
		Judged:    null.Time{},
		Score:     0,
	})
	assert.NoError(t, g.ProcessSubmission(context.TODO(), *sub))
	sub, _ = g.Submissions.Get(context.TODO(), sub.ID)
	assert.Equal(t, njudge.VerdictAC, sub.Verdict)
}
