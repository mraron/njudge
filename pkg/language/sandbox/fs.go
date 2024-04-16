package sandbox

import (
	"io"
	"io/fs"
	"os"
	"path/filepath"
)

// OsFS implements an FS with and underlying base path.
// All of its methods are executed as if the current working directory is this base path.
// Should be only created via NewOsFS
type OsFS struct {
	base string

	inited bool
}

func NewOsFS(base string) OsFS {
	return OsFS{
		base:   base,
		inited: true,
	}
}

func (o OsFS) getPathTo(name string) string {
	return filepath.Join(o.base, name)
}

func (o OsFS) Pwd() string {
	return o.base
}

func (o OsFS) Create(name string) (io.WriteCloser, error) {
	if !o.inited {
		return nil, ErrorSandboxNotInitialized
	}
	return os.Create(o.getPathTo(name))
}

func (o OsFS) MakeExecutable(name string) error {
	if !o.inited {
		return ErrorSandboxNotInitialized
	}

	filename := o.getPathTo(name)
	return os.Chmod(filename, 0755)
}

func (o OsFS) Open(name string) (fs.File, error) {
	if !o.inited {
		return nil, ErrorSandboxNotInitialized
	}
	return os.Open(o.getPathTo(name))
}
