package batch

import (
	"github.com/mraron/njudge/pkg/problems"
	"github.com/mraron/njudge/pkg/problems/evaluation"
)

const Name = "batch"

func New(compiler problems.Compiler, options ...evaluation.BasicRunnerOption) problems.TaskType {
	return problems.NewTaskType(
		Name,
		compiler,
		evaluation.NewLinearEvaluator(evaluation.NewBasicRunner(options...)),
	)
}
