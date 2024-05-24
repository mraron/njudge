package stub

import (
	"github.com/mraron/njudge/pkg/problems"
	"github.com/mraron/njudge/pkg/problems/evaluation"
)

const Name = "stub"

func New(stubCompiler *evaluation.CompileWithStubs, options ...evaluation.BasicRunnerOption) problems.TaskType {
	return problems.NewTaskType(
		Name,
		stubCompiler,
		evaluation.NewLinearEvaluator(evaluation.NewBasicRunner(options...)),
	)
}
