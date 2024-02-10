package judge

import (
	"context"
	"errors"
	"fmt"
	"github.com/mraron/njudge/pkg/problems/evaluation"
	"io"

	"github.com/mraron/njudge/pkg/language"
	"github.com/mraron/njudge/pkg/language/sandbox"
	"github.com/mraron/njudge/pkg/problems"
	"go.uber.org/zap"
)

type Worker struct {
	id              int
	sandboxProvider *sandbox.ChanProvider
}

func NewWorker(id int, sandboxProvider *sandbox.ChanProvider) *Worker {
	return &Worker{id: id, sandboxProvider: sandboxProvider}
}

func (w Worker) Judge(ctx context.Context, plogger *zap.Logger, p problems.Judgeable, src []byte, lang language.Language, c Callbacker) (st problems.Status, err error) {
	logger := plogger.With(zap.Int("worker", w.id))
	logger.Info("started to judge")

	sandboxes := sandbox.NewSandboxProvider()
	for i := 0; i < 2; i++ {
		var s sandbox.Sandbox
		s, err = w.sandboxProvider.Get()
		if err != nil {
			logger.Error("can't get sandbox", zap.Error(err))
			return
		}
		defer w.sandboxProvider.Put(s)

		err = s.Init(context.TODO())
		if err != nil {
			return
		}
		sandboxes.Put(s)

		defer func(sandbox sandbox.Sandbox) {
			err = errors.Join(err, sandbox.Cleanup(context.TODO()))
		}(s)
	}

	tt := p.GetTaskType()

	logger.Info("compiling")
	compileSandbox, _ := sandboxes.Get()
	compileRes, err := tt.Compile(context.Background(), p, evaluation.NewByteSolution(lang, src), compileSandbox)
	sandboxes.Put(compileSandbox)

	if err != nil {
		logger.Error("compilation error", zap.Error(err))
		st.Compiled = false
		st.CompilerOutput = err.Error() + "\n" + truncate(compileRes.CompilationMessage, 1024)

		return st, nil
	}
	st.Compiled = true

	var (
		testNotifier   = make(chan string)
		statusNotifier = make(chan problems.Status)
		errRun         error
		test           string = "1"

		waiter = make(chan struct{})
	)

	go func() {
		skeleton, _ := p.StatusSkeleton("")

		bin, _ := io.ReadAll(compileRes.CompiledFile)
		defer compileRes.CompiledFile.Close()

		st, errRun = tt.Evaluate(context.Background(), *skeleton, evaluation.NewByteSolution(lang, bin), sandboxes, evaluation.IgnoreStatusUpdate{})
		close(testNotifier)
		close(statusNotifier)
		waiter <- struct{}{}
	}()

	for status := range statusNotifier {
		test = <-testNotifier
		err = c.Callback(Response{test, status, false, ""})
		if err != nil {
			logger.Error("error while calling callback", zap.Error(err))
			return
		}
	}

	<-waiter
	err = errors.Join(err, errRun)
	if err == nil {
		logger.Info("successful judging")
	} else {
		logger.Info("got error while judging", zap.Error(err))
	}

	return
}

type WorkerProvider interface {
	Get() *Worker
	Put(*Worker)
}

type IsolateWorkerProvider struct {
	minSandboxId, maxSandboxId int
	sandboxIdUsed              map[int]struct{}
	workers                    chan *Worker
	workerCount                int
}

func NewIsolateWorkerProvider(minSandboxId, maxSandboxId, workerCount int) (*IsolateWorkerProvider, error) {
	wp := &IsolateWorkerProvider{
		minSandboxId:  minSandboxId,
		maxSandboxId:  maxSandboxId,
		workerCount:   workerCount,
		workers:       make(chan *Worker, workerCount),
		sandboxIdUsed: make(map[int]struct{}),
	}

	for i := 0; i < wp.workerCount; i++ {
		provider := sandbox.NewSandboxProvider()
		if err := wp.populateProvider(provider, 2); err != nil {
			return nil, err
		}

		wp.workers <- NewWorker(i+1, provider)
	}

	return wp, nil
}

func (wp *IsolateWorkerProvider) Get() *Worker {
	return <-wp.workers
}

func (wp *IsolateWorkerProvider) Put(w *Worker) {
	wp.workers <- w
}

func (wp *IsolateWorkerProvider) populateProvider(provider *sandbox.ChanProvider, cnt int) error {
	for i := wp.minSandboxId; i <= wp.maxSandboxId; i++ {
		if _, ok := wp.sandboxIdUsed[i]; !ok {
			s, _ := sandbox.NewIsolate(i)
			provider.Put(s)
			cnt -= 1
			wp.sandboxIdUsed[i] = struct{}{}
		}

		if cnt == 0 {
			break
		}
	}

	if cnt != 0 {
		return fmt.Errorf("not enough sandboxes")
	}

	return nil
}

func truncate(s string, to int) string {
	if len(s) < to {
		return s
	}

	return s[:to-1] + "..."
}
