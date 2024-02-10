package java

import (
	"bytes"
	"context"
	"github.com/mraron/njudge/pkg/language/memory"
	"github.com/mraron/njudge/pkg/language/sandbox"
	"io"
	"io/fs"
	"time"

	"github.com/mraron/njudge/pkg/language"
)

type Java struct {
	sourceName string
	className  string
	execName   string

	jars []Jar
}

type Jar struct {
	Name     string
	Contents []byte
}

func New(className string, classPath ...Jar) *Java {
	return &Java{className + ".java", className + ".class", className, classPath}
}

func (*Java) Id() string {
	return "java"
}

func (*Java) DisplayName() string {
	return "Java"
}

func (j *Java) DefaultFilename() string {
	return j.sourceName
}

func (j *Java) Compile(s sandbox.Sandbox, r language.File, w io.Writer, e io.Writer, extras []language.File) error {
	err := sandbox.CreateFileFromSource(s, j.sourceName, r.Source)
	if err != nil {
		return err
	}

	classPath := "."
	for ind := range j.jars {
		if err := sandbox.CreateFileFromSource(s, j.jars[ind].Name, bytes.NewBuffer(j.jars[ind].Contents)); err != nil {
			return err
		}
		classPath += ":" + j.jars[ind].Name
	}

	rc := sandbox.RunConfig{
		MaxProcesses: -1,
		DirectoryMaps: []sandbox.DirectoryMap{
			{"/etc", "/etc", nil},
		},
		InheritEnv:       true,
		TimeLimit:        10 * time.Second,
		MemoryLimit:      1 * memory.GiB,
		Stdout:           e,
		Stderr:           e,
		WorkingDirectory: s.Pwd(),
		Args:             []string{"--open-files=2048"},
	}
	if _, err := s.Run(context.TODO(), rc, "/usr/bin/javac", sandbox.SplitArgs("-cp "+classPath+" "+j.sourceName)...); err != nil {
		return err
	}

	bin, err := s.Open(j.className)
	if err != nil {
		return err
	}
	defer func(bin fs.File) {
		_ = bin.Close()
	}(bin)

	_, err = io.Copy(w, bin)

	return err
}

func (j *Java) Run(s sandbox.Sandbox, binary io.Reader, stdin io.Reader, stdout io.Writer, tl time.Duration, ml memory.Amount) (*sandbox.Status, error) {
	if err := sandbox.CreateFileFromSource(s, j.className, binary); err != nil {
		return nil, err
	}

	classPath := "."
	for ind := range j.jars {
		if err := sandbox.CreateFileFromSource(s, j.jars[ind].Name, bytes.NewBuffer(j.jars[ind].Contents)); err != nil {
			return nil, err
		}
		classPath += ":" + j.jars[ind].Name
	}

	rc := sandbox.RunConfig{
		MaxProcesses: -1,
		DirectoryMaps: []sandbox.DirectoryMap{
			{"/etc", "/etc", nil},
		},
		InheritEnv:       true,
		Stdin:            stdin,
		Stdout:           stdout,
		TimeLimit:        tl,
		MemoryLimit:      memory.Amount(ml) * memory.KiB,
		WorkingDirectory: s.Pwd(),
	}
	return s.Run(context.TODO(), rc, "/usr/bin/java", sandbox.SplitArgs("-cp "+classPath+" "+j.execName)...)
}

func init() {
	language.DefaultStore.Register("java", New("main"))
}
