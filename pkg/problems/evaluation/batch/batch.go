package batch

import (
	"github.com/mraron/njudge/pkg/problems"
	"github.com/mraron/njudge/pkg/problems/evaluation"
)

func New(compiler problems.Compiler, options ...evaluation.BasicRunnerOption) problems.TaskType {
	return problems.NewTaskType(
		"batch",
		compiler,
		evaluation.NewLinearEvaluator(evaluation.NewBasicRunner(options...)),
	)
}
