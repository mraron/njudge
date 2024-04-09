package judge2

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/mraron/njudge/pkg/language"
	"github.com/mraron/njudge/pkg/language/sandbox"
	"github.com/mraron/njudge/pkg/problems"
	"github.com/mraron/njudge/pkg/problems/evaluation"
	"golang.org/x/time/rate"
	"io"
	"mime"
	"net/http"
	"time"
)

type ResultCallback func(Result) error

type Judger interface {
	Judge(ctx context.Context, sub Submission, callback ResultCallback) (*problems.Status, error)
}

type Judge struct {
	SandboxProvider sandbox.Provider
	ProblemStore    problems.Store
	LanguageStore   language.Store
	RateLimit       time.Duration
}

func (j *Judge) Judge(ctx context.Context, sub Submission, callback ResultCallback) (*problems.Status, error) {
	fmt.Println(sub)
	problem, err := j.ProblemStore.GetProblem(sub.Problem)
	if err != nil {
		return nil, err
	}
	lang, err := j.LanguageStore.Get(sub.Language)
	if err != nil {
		return nil, err
	}

	provider := sandbox.NewProvider()
	for i := 0; i < 2; i++ {
		box, _ := j.SandboxProvider.Get()
		provider.Put(box)
		defer j.SandboxProvider.Put(box)
	}

	res := problems.Status{}

	compileSandbox, _ := provider.Get()
	if err := compileSandbox.Init(ctx); err != nil {
		return nil, err
	}
	taskType := problem.GetTaskType()
	compilationResult, err := taskType.Compile(ctx, evaluation.NewByteSolution(lang, sub.Source), compileSandbox)
	provider.Put(compileSandbox)
	if err != nil {
		return nil, err
	}

	if compilationResult.CompiledFile == nil {
		res.Compiled = false
		res.CompilerOutput = compilationResult.CompilationMessage
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
	res, err = taskType.Evaluate(ctx, *st, evaluation.NewByteSolution(lang, binary), provider, updater)
	res.CompilerOutput = compilationResult.CompilationMessage
	<-done
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

	s := bufio.NewScanner(resp.Body)
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
	return res.Status, err
}

func NewClient(URL string) *Client {
	return &Client{URL: URL, HttpClient: http.DefaultClient} //TODO timeout and such
}
