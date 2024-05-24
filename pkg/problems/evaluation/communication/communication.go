package communication

import (
	"github.com/mraron/njudge/pkg/problems"
	"github.com/mraron/njudge/pkg/problems/evaluation"
)

const Name = "communication"

func New(compiler problems.Compiler, interactorBinary []byte, checker problems.Checker, options ...evaluation.InteractiveRunnerOption) problems.TaskType {
	return problems.NewTaskType(
		Name,
		compiler,
		evaluation.NewLinearEvaluator(evaluation.NewInteractiveRunner(interactorBinary, checker, options...)),
	)
}
