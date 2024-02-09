package sandbox

import (
	"io"
	"io/fs"
	"os"
	"path/filepath"
)

type OsFS struct {
	base string
}

func (o OsFS) getPathTo(name string) string {
	return filepath.Join(o.base, name)
}

func (o OsFS) Pwd() string {
	return o.base
}

func (o OsFS) Create(name string) (io.WriteCloser, error) {
	return os.Create(o.getPathTo(name))
}

func (o OsFS) MakeExecutable(name string) error {
	filename := o.getPathTo(name)
	return os.Chmod(filename, 0755)
}

func (o OsFS) Open(name string) (fs.File, error) {
	return os.Open(o.getPathTo(name))
}
