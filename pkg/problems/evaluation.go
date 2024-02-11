package problems

import (
	"context"
	"github.com/mraron/njudge/pkg/language"
	"github.com/mraron/njudge/pkg/language/sandbox"
)

type CompilationResult struct {
	CompiledFile       *sandbox.File
	CompilationMessage string
}

type Solution interface {
	GetLanguage() language.Language
	GetFile(ctx context.Context) (sandbox.File, error)
}

type Compiler interface {
	Compile(ctx context.Context, problem Judgeable, solution Solution, sandbox sandbox.Sandbox) (*CompilationResult, error)
}

type StatusUpdater interface {
	UpdateStatus(ctx context.Context, testcase string, status Status) error
	Done(ctx context.Context) error
}

type Evaluator interface {
	Evaluate(ctx context.Context, skeleton Status, compiledSolution Solution, sandboxProvider sandbox.Provider, statusUpdater StatusUpdater) (Status, error)
}

type Runner interface {
	SetSolution(ctx context.Context, solution Solution) error
	Run(ctx context.Context, sandboxProvider sandbox.Provider, testcase *Testcase) error
}

type TaskType struct {
	name string
	Compiler
	Evaluator
}

func NewTaskType(name string, compiler Compiler, eval Evaluator) TaskType {
	return TaskType{
		name:      name,
		Compiler:  compiler,
		Evaluator: eval,
	}
}

func (t TaskType) Name() string {
	return t.name
}
