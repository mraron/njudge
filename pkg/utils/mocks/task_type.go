package mocks

import (
	"io"

	"github.com/mraron/njudge/pkg/language"
	"github.com/mraron/njudge/pkg/problems"
)

type TaskType struct {
	FName func() string
	FCompile func(problems.Judgeable, language.Sandbox, language.Language, io.Reader, io.Writer) (io.Reader, error)
	FRun func(problems.Judgeable, *language.SandboxProvider, language.Language, io.Reader, chan string, chan problems.Status) (problems.Status, error)
}

func (tt *TaskType) Name() string {
	if tt.FName == nil {
		panic("not implemented")
	}
	
	return tt.FName()
}

func (tt *TaskType) Compile(j problems.Judgeable, s language.Sandbox, l language.Language, src io.Reader, errStream io.Writer) (io.Reader, error) {
	if tt.FCompile == nil {
		panic("not implemented")
	}

	return tt.FCompile(j, s, l, src, errStream)
}

func (tt *TaskType) Run(j problems.Judgeable, sbp *language.SandboxProvider, l language.Language, bin io.Reader, resTest chan string, resStatus chan problems.Status) (problems.Status, error) {
	if tt.FRun == nil {
		panic("not implemented")
	}

	return tt.FRun(j, sbp, l, bin, resTest, resStatus)
}

