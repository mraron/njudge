package judge

import (
	"context"
	"errors"
	"fmt"

	"github.com/mraron/njudge/pkg/language"
	"github.com/mraron/njudge/pkg/problems"
	"go.uber.org/zap"
)

type Enqueuer interface {
	Enqueue(context.Context, Submission) (<-chan Response, error)

	SupportedProblems() ([]string, error)
	SupportedLanguages() ([]string, error)
}

type Response struct {
	Test   string
	Status problems.Status
	Done   bool
	Error  string
}

type queueSubmission struct {
	Submission
	c Callbacker
}

type Queue struct {
	problemStore   problems.Store
	languageStore  language.Store
	workerProvider WorkerProvider

	queue chan queueSubmission

	logger *zap.Logger
}

func NewQueue(logger *zap.Logger, problemStore problems.Store, languageStore language.Store, workerProvider WorkerProvider) (*Queue, error) {
	queue := &Queue{
		problemStore:   problemStore,
		languageStore:  languageStore,
		workerProvider: workerProvider,
		queue:          make(chan queueSubmission, 128),
		logger:         logger,
	}

	return queue, nil
}

func (j *Queue) Enqueue(ctx context.Context, sub Submission) (<-chan Response, error) {
	channel := make(chan Response)

	qs := queueSubmission{Submission: sub}
	qs.c = NewChanCallback(channel)
	j.queue <- qs

	return channel, nil
}

func (q *Queue) SupportedProblems() ([]string, error) {
	return q.problemStore.List()
}

func (q *Queue) SupportedLanguages() ([]string, error) {
	res := []string{}
	for _, val := range q.languageStore.List() {
		res = append(res, val.ID())
	}

	return res, nil
}

func (j *Queue) Run() {
	judge := func(worker *Worker, sub queueSubmission) error {
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
			return errors.Join(sub.c.Callback(Response{"", st, true, err.Error()}), err)
		} else {
			return sub.c.Callback(Response{"", st, true, ""})
		}
	}

	for sub := range j.queue {
		sub := sub
		go func() {
			w := j.workerProvider.Get()
			if err := judge(w, sub); err != nil {
				j.logger.Error("judging error", zap.Error(err))
			}

			j.workerProvider.Put(w)
		}()
	}
}
