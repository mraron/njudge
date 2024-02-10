package java

import (
	"bytes"
	"context"
	"github.com/mraron/njudge/pkg/language/memory"
	"github.com/mraron/njudge/pkg/language/sandbox"
	"io"
	"regexp"
	"strings"
	"time"

	"github.com/mraron/njudge/pkg/language"
)

type Java struct {
	className     string
	autoClassName bool

	jars []Jar
}

type Jar struct {
	Name     string
	Contents []byte
}

type Option func(*Java)

func AutoClassName() Option {
	return func(java *Java) {
		java.autoClassName = true
	}
}

func WithJars(jars ...Jar) Option {
	return func(java *Java) {
		java.jars = jars
	}
}

func New(className string, opts ...Option) *Java {
	java := &Java{className: className}
	for _, opt := range opts {
		opt(java)
	}
	return java
}

var classRegexp = regexp.MustCompile("public +class +(\\w+)")

func (j *Java) Rename(source sandbox.File) (*sandbox.File, string, error) {
	if !j.autoClassName {
		return &source, j.className, nil
	}

	src, err := io.ReadAll(source.Source)
	if err != nil {
		return nil, "", err
	}
	source.Source = bytes.NewBuffer(src)

	lst := classRegexp.FindSubmatch(src)
	if len(lst) <= 1 {
		source.Name = j.DefaultFilename()
		return &source, j.className, nil
	}

	source.Name = string(lst[1]) + ".java"
	return &source, string(lst[1]), nil
}

func (*Java) ID() string {
	return "java"
}

func (*Java) DisplayName() string {
	return "Java"
}

func (j *Java) DefaultFilename() string {
	return j.className + ".java"
}

func (j *Java) Compile(ctx context.Context, s sandbox.Sandbox, f sandbox.File, stderr io.Writer, extras []sandbox.File) (*sandbox.File, error) {
	renamed, className, err := j.Rename(f)
	if err != nil {
		return nil, err
	}

	err = sandbox.CreateFileFromSource(s, renamed.Name, renamed.Source)
	if err != nil {
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
		TimeLimit:        10 * time.Second,
		MemoryLimit:      1 * memory.GiB,
		Stdout:           stderr,
		Stderr:           stderr,
		WorkingDirectory: s.Pwd(),
		Args:             []string{"--open-files=2048"},
	}
	if _, err := s.Run(ctx, rc, "/usr/bin/javac", sandbox.SplitArgs("-cp "+classPath+" "+renamed.Name)...); err != nil {
		return nil, err
	}

	return sandbox.ExtractFile(s, className+".class")
}

func (j *Java) Run(ctx context.Context, s sandbox.Sandbox, binary sandbox.File, stdin io.Reader, stdout io.Writer, tl time.Duration, ml memory.Amount) (*sandbox.Status, error) {
	if err := sandbox.CreateFileFromSource(s, binary.Name, binary.Source); err != nil {
		return nil, err
	}

	classPath := "."
	for ind := range j.jars {
		if err := sandbox.CreateFileFromSource(s, j.jars[ind].Name, bytes.NewBuffer(j.jars[ind].Contents)); err != nil {
			return nil, err
		}
		classPath += ":" + j.jars[ind].Name
	}

	execName := strings.Split(binary.Name, ".")[0]
	rc := sandbox.RunConfig{
		MaxProcesses: -1,
		DirectoryMaps: []sandbox.DirectoryMap{
			{"/etc", "/etc", nil},
		},
		InheritEnv:       true,
		Stdin:            stdin,
		Stdout:           stdout,
		TimeLimit:        tl,
		MemoryLimit:      ml,
		WorkingDirectory: s.Pwd(),
	}
	return s.Run(ctx, rc, "/usr/bin/java", sandbox.SplitArgs("-cp "+classPath+" "+execName)...)
}

func init() {
	language.DefaultStore.Register("java", New("main", AutoClassName()))
}
