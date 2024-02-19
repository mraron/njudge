package zip

import (
	"context"
	"github.com/mraron/njudge/pkg/language/memory"
	"github.com/mraron/njudge/pkg/language/sandbox"
	"io"
	"time"

	"github.com/mraron/njudge/pkg/language"
)

type Zip struct{}

func (Zip) ID() string {
	return "zip"
}

func (Zip) DisplayName() string {
	return "ZIP archívum"
}

func (Zip) DefaultFilename() string {
	return "main.zip"
}

func (Zip) Compile(_ context.Context, _ sandbox.Sandbox, f sandbox.File, _ io.Writer, _ []sandbox.File) (*sandbox.File, error) {
	return &f, nil
}

func (Zip) Run(_ context.Context, _ sandbox.Sandbox, _ sandbox.File, _ io.Reader, _ io.Writer, _ time.Duration, _ memory.Amount) (*sandbox.Status, error) {
	return &sandbox.Status{}, nil
}

func init() {
	language.DefaultStore.Register("zip", Zip{})
}
