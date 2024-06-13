package judge

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/karrick/gobls"
	"github.com/mraron/njudge/pkg/language"
	"github.com/mraron/njudge/pkg/language/sandbox"
	"github.com/mraron/njudge/pkg/problems"
	"github.com/mraron/njudge/pkg/problems/evaluation"
	"golang.org/x/time/rate"
	"io"
	"log/slog"
	"mime"
	"net/http"
	"time"
)

type ResultCallback func(Result) error

type Judger interface {
	Judge(ctx context.Context, sub Submission, callback ResultCallback) (*problems.Status, error)
}

type ProblemStore interface {
	GetProblem(string) (problems.Problem, error)
}

type Judge struct {
	SandboxProvider sandbox.Provider
	ProblemStore    ProblemStore
	LanguageStore   language.Store
	RateLimit       time.Duration

	Tokens chan struct{}
	Logger *slog.Logger
}

func (j *Judge) Judge(ctx context.Context, sub Submission, callback ResultCallback) (*problems.Status, error) {
	if j.Logger == nil {
		j.Logger = slog.New(slog.NewJSONHandler(io.Discard, nil))
	}
	j.Logger.Info("🏃\tjudging started", "submission_id", sub.ID)
	problem, err := j.ProblemStore.GetProblem(sub.Problem)
	if err != nil {
		return nil, err
	}
	lang, err := j.LanguageStore.Get(sub.Language)
	if err != nil {
		return nil, err
	}

	res := problems.Status{}

	j.Logger.Info("🏗️\tcompilation step", "submission_id", sub.ID)
	compileSandbox, _ := j.SandboxProvider.Get()
	if err := compileSandbox.Init(ctx); err != nil {
		return nil, err
	}
	taskType := problem.GetTaskType()
	res.CompilationStatus = problems.DuringCompilation // this has no real effect
	compilationResult, err := taskType.Compile(ctx, evaluation.NewByteSolution(lang, sub.Source), compileSandbox)
	j.SandboxProvider.Put(compileSandbox)
	if err != nil {
		return nil, err
	}

	if compilationResult.CompiledFile == nil {
		res.Compiled = false
		res.CompilationStatus = problems.AfterCompilation
		res.CompilerOutput = problems.Base64String(compilationResult.CompilationMessage)
		return &res, nil
	}

	binary, err := io.ReadAll(compilationResult.CompiledFile.Source)
	if err != nil {
		return nil, fmt.Errorf("failed to read binary: %w", err)
	}
	_ = compilationResult.CompiledFile.Source.Close()

	st, err := problem.StatusSkeleton("")
	if err != nil {
		return nil, err
	}

	j.Logger.Info("🔥\tstart evaluation", "submission_id", sub.ID)
	innerUpdater, updates := evaluation.NewChanStatusUpdate()
	updater := evaluation.NewRateLimitStatusUpdate(innerUpdater, rate.Every(j.RateLimit))
	done := make(chan struct{})
	go func() {
		ind := 1
		for update := range updates {
			_ = callback(Result{
				Index:  ind,
				Test:   update.Testcase,
				Status: &update.Status,
			})
			ind++
		}
		close(done)
	}()

	eval := taskType.Evaluator
	if j.Tokens != nil && (st.FeedbackType == problems.FeedbackIOI || st.FeedbackType == problems.FeedbackLazyIOI) {
		j.Logger.Info("🔀\tusing parallel evaluation", "submission_id", sub.ID)
		eval = &ParallelEvaluator{
			Runner: evaluation.NewCachedRunner(taskType.Evaluator.(*evaluation.LinearEvaluator).Runner),
			Tokens: j.Tokens,
			Logger: j.Logger.With("submission_id", sub.ID),
		}
	}
	res, err = eval.Evaluate(ctx, *st, evaluation.NewByteSolution(lang, binary), j.SandboxProvider, updater)
	res.CompilerOutput = problems.Base64String(compilationResult.CompilationMessage)
	<-done
	j.Logger.Info("🏁\tdone", "submission_id", sub.ID)
	return &res, err
}

type Client struct {
	URL        string
	HttpClient *http.Client
}

func (c Client) Judge(ctx context.Context, sub Submission, callback ResultCallback) (*problems.Status, error) {
	buf, err := json.Marshal(sub)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequestWithContext(ctx, "POST", c.URL+"/judge", bytes.NewBuffer(buf))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", mime.TypeByExtension(".json"))

	resp, err := c.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}

	s := gobls.NewScanner(resp.Body)
	var res Result
	for s.Scan() {
		if err2 := json.Unmarshal(s.Bytes(), &res); err2 != nil {
			err = errors.Join(err, err2)
		}

		if res.Index == 0 {
			break
		} else {
			if err2 := callback(res); err2 != nil {
				err = errors.Join(err, err2)
			}
		}
	}
	err = errors.Join(err, s.Err(), resp.Body.Close())
	if res.Error != "" {
		err = errors.Join(err, errors.New(res.Error))
	}
	return res.Status, err
}

func NewClient(URL string) *Client {
	return &Client{URL: URL, HttpClient: http.DefaultClient} //TODO timeout and such
}
