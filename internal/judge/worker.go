package judge

import (
	"bytes"
	"context"
	"fmt"
	"time"

	"github.com/mraron/njudge/pkg/language"
	"github.com/mraron/njudge/pkg/language/sandbox"
	"github.com/mraron/njudge/pkg/problems"
	"go.uber.org/multierr"
	"go.uber.org/zap"
)

type SubmissionStatus struct {
	Test   string
	Status problems.Status
	Done   bool
	Time   time.Time
}

type Worker struct {
	id              int
	sandboxProvider *language.SandboxProvider
}

func NewWorker(id int, sandboxProvider *language.SandboxProvider) *Worker {
	return &Worker{id: id, sandboxProvider: sandboxProvider}
}

func (w Worker) Judge(ctx context.Context, plogger *zap.Logger, p problems.Judgeable, src []byte, lang language.Language, c Callbacker) (st problems.Status, err error) {
	logger := plogger.With(zap.Int("worker", w.id))
	logger.Info("started to judge")

	sandboxes := language.NewSandboxProvider()
	for i := 0; i < 2; i++ {
		var sandbox language.Sandbox
		sandbox, err = w.sandboxProvider.Get()
		if err != nil {
			logger.Error("can't get sandbox", zap.Error(err))
			return
		}
		defer w.sandboxProvider.Put(sandbox)

		err = sandbox.Init(zap.NewStdLog(logger))
		if err != nil {
			return
		}
		sandboxes.Put(sandbox)

		defer func(sandbox language.Sandbox) {
			err = multierr.Append(err, sandbox.Cleanup())
		}(sandbox)
	}

	var (
		tt     problems.TaskType = p.GetTaskType()
		stderr bytes.Buffer      = bytes.Buffer{}
	)

	logger.Info("compiling")

	compileSandbox := sandboxes.MustGet()
	bin, err := tt.Compile(p, compileSandbox, lang, bytes.NewReader(src), &stderr)
	sandboxes.Put(compileSandbox)

	if err != nil {
		logger.Error("compilation error", zap.Error(err))
		st.Compiled = false
		st.CompilerOutput = err.Error() + "\n" + truncate(stderr.String(), 1024)

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
		st, errRun = tt.Run(p, sandboxes, lang, bin, testNotifier, statusNotifier)
		waiter <- struct{}{}
	}()

	for status := range statusNotifier {
		test = <-testNotifier
		err = c.Callback(test, status, false)
		if err != nil {
			logger.Error("error while calling callback", zap.Error(err))
			return
		}
	}

	<-waiter
	err = multierr.Combine(err, errRun)
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
	workerCount int
}

func NewIsolateWorkerProvider(minSandboxId, maxSandboxId, workerCount int) (*IsolateWorkerProvider, error) {
	wp := &IsolateWorkerProvider{
		minSandboxId: minSandboxId,
		maxSandboxId: maxSandboxId,
		workerCount: workerCount,
		workers: make(chan *Worker, workerCount),
		sandboxIdUsed: make(map[int]struct{}),
	}

	for i := 0; i < wp.workerCount; i++ {
		provider := language.NewSandboxProvider()
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

func (wp *IsolateWorkerProvider) populateProvider(provider *language.SandboxProvider, cnt int) error {
	for i := wp.minSandboxId; i <= wp.maxSandboxId; i++ {
		if _, ok := wp.sandboxIdUsed[i]; !ok {
			provider.Put(sandbox.NewIsolate(i))
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
