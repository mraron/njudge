package judge

import (
	"bytes"
	"context"
	"time"

	"github.com/mraron/njudge/pkg/language"
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

func (w Worker) Judge(ctx context.Context, plogger *zap.Logger, p problems.Problem, src []byte, lang language.Language, c Callbacker) (st problems.Status, err error) {
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

	tt := p.GetTaskType()

	stderr := bytes.Buffer{}

	logger.Info("compiling")

	compileSandbox := sandboxes.MustGet()
	bin, err := tt.Compile(p, compileSandbox, lang, bytes.NewReader(src), &stderr)
	sandboxes.Put(compileSandbox)

	if err != nil {
		logger.Error("compilation error", zap.Error(err))
		st.Compiled = false
		st.CompilerOutput = err.Error() + "\n" + truncate(stderr.String(), 1024)
		err = multierr.Append(err, c.Callback("", st, true))
		return
	}
	st.Compiled = true

	var (
		testNotifier   = make(chan string)
		statusNotifier = make(chan problems.Status)
		ran            = make(chan bool)
		errRun         error
	)

	go func() {
		st, errRun = tt.Run(p, sandboxes, lang, bin, testNotifier, statusNotifier)
		ran <- true
		close(ran)
	}()

	run := true
	for run {
		select {
		case test, ok := <-testNotifier:
			if ok {
				status := <-statusNotifier

				err = c.Callback(test, status, false)
				if err != nil {
					logger.Error("error while calling callback", zap.Error(err))
					return
				}
			}
		case <-ran:
			run = false
			break

		}
	}
	err = multierr.Combine(err, errRun)

	if err == nil {
		logger.Info("successful judging")
	} else {
		logger.Info("got error while judging", zap.Error(err))
	}

	return
}

func truncate(s string, to int) string {
	if len(s) < to {
		return s
	}

	return s[:to-1] + "..."
}
