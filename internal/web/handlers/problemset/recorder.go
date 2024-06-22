package problemset

import (
	"context"
	"github.com/mraron/njudge/pkg/language"
	"github.com/mraron/njudge/pkg/language/sandbox"
	"github.com/mraron/njudge/pkg/problems"
	"github.com/mraron/njudge/pkg/problems/evaluation"
	"io"
	"io/fs"
	"strings"
)

type recordingSandbox struct {
	recorded string
}

func (r *recordingSandbox) Id() string {
	return "recording"
}

func (r *recordingSandbox) Init(ctx context.Context) error {
	return nil
}

func (r *recordingSandbox) Pwd() string {
	return ""
}

type closer struct {
	io.Writer
}

func (c closer) Close() error {
	return nil
}

func (r *recordingSandbox) CreateFile(file sandbox.File) error {
	return nil
}

func (r *recordingSandbox) Create(name string) (io.WriteCloser, error) {
	return closer{io.Discard}, nil
}

func (r *recordingSandbox) MakeExecutable(name string) error {
	return nil
}

type fakeFsFile struct {
}

func (f fakeFsFile) Stat() (fs.FileInfo, error) {
	return nil, nil
}

func (f fakeFsFile) Read(bytes []byte) (int, error) {
	return 0, io.EOF
}

func (f fakeFsFile) Close() error {
	return nil
}

func (r *recordingSandbox) Open(name string) (fs.File, error) {
	return fakeFsFile{}, nil
}

func (r *recordingSandbox) Run(ctx context.Context, config sandbox.RunConfig, command string, commandArgs ...string) (*sandbox.Status, error) {
	r.recorded = command + " " + strings.Join(commandArgs, " ")
	return nil, nil
}

func (r *recordingSandbox) Cleanup(ctx context.Context) error {
	return nil
}

func extractCompileAndRunCommand(ctx context.Context, taskType problems.TaskType, lang language.Language) (string, error) {
	rec := &recordingSandbox{}
	file, err := taskType.Compile(ctx, evaluation.NewByteSolution(lang, nil), rec)
	if err != nil {
		return "", err
	}
	res := rec.recorded
	_, err = lang.Run(ctx, rec, *file.CompiledFile, nil, nil, 0, 0)
	if err != nil {
		return "", err
	}
	if res != "" {
		res += "\n"
	}
	res += rec.recorded
	return res, nil
}
