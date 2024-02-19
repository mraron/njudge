package judge_test

import (
	"context"
	"errors"
	problemsMock "github.com/mraron/njudge/mocks/github.com/mraron/njudge/pkg/problems"
	"github.com/stretchr/testify/mock"
	"io"
	"sync"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/mraron/njudge/internal/judge"
	"github.com/mraron/njudge/pkg/language/sandbox"
	"github.com/mraron/njudge/pkg/problems"
	"go.uber.org/zap"
)

func TestWorker(t *testing.T) {
	tests := []struct {
		Name              string
		Judgeable         func() problems.EvaluationInfo
		JudgeReturnStatus problems.Status
		JudgeReturnErr    error
		Responses         []judge.Response
	}{
		{
			"TestWorkerRunning",
			func() problems.EvaluationInfo {
				var tasktype problemsMock.TaskType
				tasktype.On("Compile", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil, nil)
				tasktype.On("Run", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Run(func(args mock.Arguments) {
					c1, c2 := args.Get(4).(chan string), args.Get(5).(chan problems.Status)
					c2 <- problems.Status{CompilerOutput: "hehe"}
					c1 <- "1"
					c2 <- problems.Status{CompilerOutput: "huhu"}
					c1 <- "2"
					close(c1)
					close(c2)
				}).Return(problems.Status{}, nil)

				var judgeable problemsMock.Judgeable
				judgeable.On("GetTaskType").Return(&tasktype)
				return &judgeable
			},
			problems.Status{},
			nil,
			[]judge.Response{
				{Test: "1", Status: problems.Status{CompilerOutput: "hehe"}},
				{Test: "2", Status: problems.Status{CompilerOutput: "huhu"}},
			},
		},
		{
			"TestWorkerCompileError",
			func() problems.EvaluationInfo {
				var tasktype problemsMock.TaskType
				tasktype.On("Compile", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Run(func(args mock.Arguments) {
					w := args.Get(4).(io.Writer)
					w.Write([]byte("a"))
				}).Return(nil, errors.New(""))
				tasktype.On("Run", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(problems.Status{}, nil)

				var judgeable problemsMock.Judgeable
				judgeable.On("GetTaskType").Return(&tasktype)
				return &judgeable
			},
			problems.Status{CompilerOutput: "\na", Compiled: false},
			nil,
			[]judge.Response{},
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			sbp := sandbox.NewSandboxProvider()
			s1, _ := sandbox.NewDummy()
			sbp.Put(s1)
			s2, _ := sandbox.NewDummy()
			sbp.Put(s2)

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
			ret, err := w.Judge(context.Background(), zap.NewNop(), test.Judgeable(), []byte(""), nil, cb)
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
