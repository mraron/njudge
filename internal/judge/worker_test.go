package judge_test

import (
	"context"
	"errors"
	"io"
	"sync"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/mraron/njudge/internal/judge"
	"github.com/mraron/njudge/pkg/language"
	"github.com/mraron/njudge/pkg/language/sandbox"
	"github.com/mraron/njudge/pkg/problems"
	"github.com/mraron/njudge/pkg/utils/mocks"
	"go.uber.org/zap"
)




func TestWorker(t *testing.T) {
	tests := []struct{
		Name string
		Judgeable problems.Judgeable
		JudgeReturnStatus problems.Status
		JudgeReturnErr error
		Responses []judge.Response
	}{
		{
			"TestWorkerRunning",
			&mocks.Judgeable{FGetTaskType: func() problems.TaskType {
				return &mocks.TaskType{
					FCompile: func(j problems.Judgeable, s language.Sandbox, l language.Language, r io.Reader, w io.Writer) (io.Reader, error) {
						return nil, nil
					},
					FRun: func(j problems.Judgeable, sp *language.SandboxProvider, l language.Language, r io.Reader, c1 chan string, c2 chan problems.Status) (problems.Status, error) {
						c2 <- problems.Status{CompilerOutput: "hehe"}
						c1 <- "1"
						c2 <- problems.Status{CompilerOutput: "huhu"}
						c1 <- "2"
						close(c1)
						close(c2)
						return problems.Status{}, nil
					},
				}
			}},
			problems.Status{},
			nil,
			[]judge.Response{
				{Test: "1", Status: problems.Status{CompilerOutput: "hehe"}},
				{Test: "2", Status: problems.Status{CompilerOutput: "huhu"}},
			},
		},
		{
			"TestWorkerCompileError",
			&mocks.Judgeable{FGetTaskType: func() problems.TaskType {
				return &mocks.TaskType{
					FCompile: func(j problems.Judgeable, s language.Sandbox, l language.Language, r io.Reader, w io.Writer) (io.Reader, error) {
						w.Write([]byte("a"))
						return nil, errors.New("")
					},
					FRun: func(j problems.Judgeable, sp *language.SandboxProvider, l language.Language, r io.Reader, c1 chan string, c2 chan problems.Status) (problems.Status, error) {
						return problems.Status{}, nil
					},
				}
			}},
			problems.Status{CompilerOutput: "\na", Compiled: false},
			nil,
			[]judge.Response{},
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			sbp := language.NewSandboxProvider()
			sbp.Put(sandbox.NewDummy())
			sbp.Put(sandbox.NewDummy())
			
			ch := make(chan judge.Response)
			cb := judge.NewChanCallback(ch)

			wg := sync.WaitGroup{}
			wg.Add(1)
			go func() {
				ind := 0
				for resp := range ch {
					if !cmp.Equal(test.Responses[ind], resp) {
						t.Errorf("%v != %v", test.Responses[ind], resp)
					}
					ind++
				}

				if ind != len(test.Responses) {
					t.Error("wrong number of responses")
				}

				wg.Done()
			}()

			w := judge.NewWorker(1, sbp)
			ret, err := w.Judge(context.Background(), zap.NewNop(), test.Judgeable, []byte(""), nil, cb)
			close(ch)

			if !cmp.Equal(ret, test.JudgeReturnStatus) {
				t.Errorf("%v != %v", ret, test.JudgeReturnStatus)
			}
			
			if err != test.JudgeReturnErr {
				t.Errorf("%s != %s", err, test.JudgeReturnErr)
			}

			wg.Wait()
		})
	}
}

