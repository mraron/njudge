package output_only

import (
	"github.com/mraron/njudge/pkg/language"
	"github.com/mraron/njudge/pkg/language/langs/zip"
	"github.com/mraron/njudge/pkg/problems"
	"github.com/mraron/njudge/pkg/problems/evaluation"
)

func New(checker problems.Checker) problems.TaskType {
	return problems.NewTaskType(
		"outputonly",
		evaluation.CompileCheckSupported{
			List:         []language.Language{zip.Zip{}},
			NextCompiler: evaluation.Compile{},
		},
		evaluation.NewLinearEvaluator(evaluation.NewZipRunner(checker)),
	)
}
