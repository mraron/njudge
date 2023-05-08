package judge

import (
	"context"
	"fmt"

	"github.com/mraron/njudge/pkg/language"
	"github.com/mraron/njudge/pkg/problems"
	"go.uber.org/multierr"
	"go.uber.org/zap"
)

type Enqueuer interface {
	Enqueue(context.Context, Submission) (<-chan Response, error)
} 

type Response struct {
	Test string
	Status problems.Status
	Done bool
}

type Queue struct {
	problemStore               problems.Store
	languageStore              language.Store
	workerProvider WorkerProvider

	queue                      chan Submission
	
	logger                     *zap.Logger
}

func NewQueue(logger *zap.Logger, problemStore problems.Store, languageStore language.Store, workerProvider WorkerProvider)  (*Queue, error) {
	queue := &Queue{
		problemStore: problemStore,
		languageStore: languageStore,
		workerProvider: workerProvider,
		queue:  make(chan Submission, 128),
		logger: logger,
	}


	return queue, nil
}

func (j *Queue) Enqueue(ctx context.Context, sub Submission) (<-chan Response, error) {
	channel := make(chan Response)
	
	sub.c = NewChanCallback(channel)
	j.queue <- sub

	return channel, nil
}


func (j *Queue) Run() {
	judge := func(worker *Worker, sub Submission) error {
		p, err := j.problemStore.Get(sub.Problem)
		if err != nil {
			return err
		}

		lang := j.languageStore.Get(sub.Language)
		if lang == nil {
			return fmt.Errorf("no such language: %s", sub.Language)
		}

		st, err := worker.Judge(context.Background(), j.logger, p, sub.Source, lang, sub.c)
		if err != nil {
			j.logger.Error("judge error", zap.Error(err))

			st.Compiled = false
			st.CompilerOutput = "internal error: " + err.Error()
			return multierr.Combine(sub.c.Callback("", st, true), err)
		} else {
			return sub.c.Callback("", st, true)
		}
	}

	for {
		w := j.workerProvider.Get()
		sub := <-j.queue

		if err := judge(w, sub); err != nil {
			j.logger.Error("judging error", zap.Error(err))
		}

		j.workerProvider.Put(w)
	}
}