package checker

import (
	"github.com/mraron/njudge/pkg/problems"
	"golang.org/x/net/context"
)

// Noop doesn't perform any checking
type Noop struct{}

func (Noop) Name() string {
	return "noop"
}

func (Noop) Check(ctx context.Context, testcase *problems.Testcase) error {
	return nil
}
