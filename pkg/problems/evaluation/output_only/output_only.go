package output_only

import (
	"github.com/mraron/njudge/pkg/language"
	"github.com/mraron/njudge/pkg/language/langs/zip"
	"github.com/mraron/njudge/pkg/problems"
	"github.com/mraron/njudge/pkg/problems/evaluation"
)

const Name = "outputonly"

func New(checker problems.Checker) problems.TaskType {
	return problems.NewTaskType(
		Name,
		evaluation.CompileCheckSupported{
			List:         []language.Language{zip.Zip{}},
			NextCompiler: evaluation.Compile{},
		},
		evaluation.NewLinearEvaluator(evaluation.NewZipRunner(checker)),
	)
}
