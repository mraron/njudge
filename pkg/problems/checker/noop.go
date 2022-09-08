package checker

import "github.com/mraron/njudge/pkg/problems"

type Noop struct{}

func (Noop) Name() string {
	return "noop"
}

func (Noop) Check(testcase *problems.Testcase) error {
	return nil
}
