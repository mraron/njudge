package output_only

import (
	"github.com/mraron/njudge/pkg/language"
	"github.com/mraron/njudge/pkg/problems"
	"github.com/mraron/njudge/pkg/problems/evaluation"
)

func New(checker problems.Checker) problems.TaskType {
	return problems.NewTaskType(
		"outputonly",
		evaluation.CompileCheckSupported{
			List:         []language.Language{language.DefaultStore.Get("zip")},
			NextCompiler: evaluation.Compile{},
		},
		evaluation.NewLinearEvaluator(evaluation.NewZipRunner(checker)),
	)
}
