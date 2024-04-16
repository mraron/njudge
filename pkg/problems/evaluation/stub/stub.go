package stub

import (
	"github.com/mraron/njudge/pkg/problems"
	"github.com/mraron/njudge/pkg/problems/evaluation"
)

func New(stubCompiler *evaluation.CompileWithStubs, options ...evaluation.BasicRunnerOption) problems.TaskType {
	return problems.NewTaskType(
		"stub",
		stubCompiler,
		evaluation.NewLinearEvaluator(evaluation.NewBasicRunner(options...)),
	)
}
