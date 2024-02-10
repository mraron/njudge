package evaluation

import (
	"context"
	"github.com/mraron/njudge/pkg/problems"
)

type IgnoreStatusUpdate struct{}

func (i IgnoreStatusUpdate) UpdateStatus(ctx context.Context, testcase string, status problems.Status) error {
	return nil
}

func (i IgnoreStatusUpdate) Done(ctx context.Context) error {
	return nil
}
