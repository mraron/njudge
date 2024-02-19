package communication

import (
	"github.com/mraron/njudge/pkg/problems"
	"github.com/mraron/njudge/pkg/problems/evaluation"
)

func New(compiler problems.Compiler, interactorBinary []byte, checker problems.Checker) problems.TaskType {
	return problems.NewTaskType(
		"communication",
		compiler,
		evaluation.NewLinearEvaluator(evaluation.NewInteractiveRunner(interactorBinary, checker)),
	)
}
