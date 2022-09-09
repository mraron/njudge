package checker

import "github.com/mraron/njudge/pkg/problems"

// Noop doesn't perform any checking
type Noop struct{}

func (Noop) Name() string {
	return "noop"
}

func (Noop) Check(testcase *problems.Testcase) error {
	return nil
}
