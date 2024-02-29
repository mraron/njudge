package evaluation

import (
	"context"
	"github.com/mraron/njudge/pkg/problems"
	"golang.org/x/time/rate"
)

type IgnoreStatusUpdate struct{}

func (i IgnoreStatusUpdate) UpdateStatus(ctx context.Context, testcase string, status problems.Status) error {
	return nil
}

func (i IgnoreStatusUpdate) Done(ctx context.Context) error {
	return nil
}

type Status struct {
	Testcase string
	Status   problems.Status
}

type ChanStatusUpdate struct {
	res chan Status
}

func NewChanStatusUpdate() (ChanStatusUpdate, <-chan Status) {
	u := ChanStatusUpdate{make(chan Status)}
	return u, u.res
}

func (c ChanStatusUpdate) UpdateStatus(_ context.Context, testcase string, status problems.Status) error {
	c.res <- Status{testcase, DeepCopyStatus(status)}
	return nil
}

func (c ChanStatusUpdate) Done(_ context.Context) error {
	close(c.res)
	return nil
}

type RateLimitStatusUpdate struct {
	su      problems.StatusUpdater
	limiter *rate.Limiter
}

func NewRateLimitStatusUpdate(innerUpdater problems.StatusUpdater, r rate.Limit) *RateLimitStatusUpdate {
	return &RateLimitStatusUpdate{
		su:      innerUpdater,
		limiter: rate.NewLimiter(r, 1),
	}
}

func (r RateLimitStatusUpdate) UpdateStatus(ctx context.Context, testcase string, status problems.Status) error {
	if r.limiter.Allow() {
		return r.su.UpdateStatus(ctx, testcase, status)
	}
	return nil
}

func (r RateLimitStatusUpdate) Done(ctx context.Context) error {
	return r.su.Done(ctx)
}
