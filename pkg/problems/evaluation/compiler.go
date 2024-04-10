package evaluation

import (
	"bytes"
	"context"
	"fmt"
	"github.com/mraron/njudge/pkg/language"
	"github.com/mraron/njudge/pkg/language/sandbox"
	"github.com/mraron/njudge/pkg/problems"
	"io"
	"os"
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

func (c CompileCopyFile) Compile(ctx context.Context, solution problems.Solution, sandbox sandbox.Sandbox) (*problems.CompilationResult, error) {
	f, err := solution.GetFile(ctx)
	return &problems.CompilationResult{
		CompiledFile:       &f,
		CompilationMessage: "",
	}, err
}

type Compile struct{}

func truncate(s string, to int) string {
	if len(s) < to {
		return s
	}

	return s[:to-1] + "..."
}

func (c Compile) CompileWithExtras(ctx context.Context, solution problems.Solution, s sandbox.Sandbox, extras []sandbox.File) (*problems.CompilationResult, error) {
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
	}, stderrTruncated, extras); err != nil {
		if len(stderr.String()) == 0 {
			return nil, err
		}
		return &problems.CompilationResult{
			CompiledFile:       nil,
			CompilationMessage: truncate(err.Error()+"\n"+stderr.String(), 2048),
		}, nil
	}

	return &problems.CompilationResult{
		CompiledFile:       res,
		CompilationMessage: stderr.String(),
	}, nil
}

func (c Compile) Compile(ctx context.Context, solution problems.Solution, s sandbox.Sandbox) (*problems.CompilationResult, error) {
	return c.CompileWithExtras(ctx, solution, s, nil)
}

type CompileCheckSupported struct {
	List         []language.Language
	NextCompiler problems.Compiler
}

func (c CompileCheckSupported) Compile(ctx context.Context, solution problems.Solution, sandbox sandbox.Sandbox) (*problems.CompilationResult, error) {
	found := false
	for _, l := range c.List {
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

	return c.NextCompiler.Compile(ctx, solution, sandbox)
}

type CompileWithStubs struct {
	stubs map[string][]problems.EvaluationFile
}

func NewCompilerWithStubs() *CompileWithStubs {
	return &CompileWithStubs{
		make(map[string][]problems.EvaluationFile),
	}
}

func (c *CompileWithStubs) AddStub(lang language.Language, stub problems.EvaluationFile) {
	if _, ok := c.stubs[lang.ID()]; !ok {
		c.stubs[lang.ID()] = make([]problems.EvaluationFile, 0)
	}
	c.stubs[lang.ID()] = append(c.stubs[lang.ID()], stub)
}

func (c *CompileWithStubs) Compile(ctx context.Context, solution problems.Solution, box sandbox.Sandbox) (*problems.CompilationResult, error) {
	lang := solution.GetLanguage()
	neededFiles := c.stubs[lang.ID()]
	sandboxFiles := make([]sandbox.File, 0, len(neededFiles))
	for ind := range neededFiles {
		contents, err := os.ReadFile(neededFiles[ind].Path)
		if err != nil {
			return nil, err
		}

		sandboxFiles = append(sandboxFiles, sandbox.File{
			Name:   neededFiles[ind].Name,
			Source: io.NopCloser(bytes.NewBuffer(contents)),
		})
	}

	return Compile{}.CompileWithExtras(ctx, solution, box, sandboxFiles)

}
