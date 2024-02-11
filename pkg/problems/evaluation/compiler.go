package evaluation

import (
	"bytes"
	"context"
	"fmt"
	"github.com/mraron/njudge/pkg/language"
	"github.com/mraron/njudge/pkg/language/sandbox"
	"github.com/mraron/njudge/pkg/problems"
	"io"
	"testing/iotest"
)

type BytesSolution struct {
	lang language.Language
	src  []byte
}

func NewByteSolution(lang language.Language, src []byte) *BytesSolution {
	return &BytesSolution{
		lang: lang,
		src:  src,
	}
}

func (b *BytesSolution) GetLanguage() language.Language {
	return b.lang
}

func (b *BytesSolution) GetFile(ctx context.Context) (sandbox.File, error) {
	return sandbox.File{
		Name:   b.lang.DefaultFilename(),
		Source: io.NopCloser(bytes.NewBuffer(b.src)),
	}, nil
}

type CompileCopyFile struct {
}

func (c CompileCopyFile) Compile(ctx context.Context, problem problems.Judgeable, solution problems.Solution, sandbox sandbox.Sandbox) (*problems.CompilationResult, error) {
	f, err := solution.GetFile(ctx)
	return &problems.CompilationResult{
		CompiledFile:       &f,
		CompilationMessage: "",
	}, err
}

type Compile struct{}

func (c Compile) Compile(ctx context.Context, problem problems.Judgeable, solution problems.Solution, s sandbox.Sandbox) (*problems.CompilationResult, error) {
	lang := solution.GetLanguage()

	f, err := solution.GetFile(ctx)
	if err != nil {
		return nil, err
	}

	stderr, res := &bytes.Buffer{}, &sandbox.File{}
	stderrTruncated := iotest.TruncateWriter(stderr, 1<<16)
	if res, err = lang.Compile(ctx, s, sandbox.File{
		Name:   lang.DefaultFilename(),
		Source: f.Source,
	}, stderrTruncated, nil); err != nil {
		return &problems.CompilationResult{
			CompiledFile:       nil,
			CompilationMessage: stderr.String(),
		}, err
	}

	return &problems.CompilationResult{
		CompiledFile:       res,
		CompilationMessage: stderr.String(),
	}, nil
}

type CompileCheckSupported struct{}

func (c CompileCheckSupported) Compile(ctx context.Context, problem problems.Judgeable, solution problems.Solution, sandbox sandbox.Sandbox) (*problems.CompilationResult, error) {
	lst, found := problem.Languages(), false

	for _, l := range lst {
		if l.ID() == solution.GetLanguage().ID() {
			found = true
		}
	}

	if !found {
		return &problems.CompilationResult{
			CompiledFile:       nil,
			CompilationMessage: "",
		}, fmt.Errorf("language is not supported: %s", solution.GetLanguage().ID())
	}

	return Compile{}.Compile(ctx, problem, solution, sandbox)
}
