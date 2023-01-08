package java

import (
	"bytes"
	"errors"
	"io"
	"os"
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

func (*Java) Name() string {
	return "Java"
}

func (j *Java) DefaultFileName() string {
	return j.sourceName
}

func (*Java) InsecureCompile(wd string, r io.Reader, w io.Writer, e io.Writer) error {
	return errors.New("can't insecure compile java")
}

func (j *Java) Compile(s language.Sandbox, r language.File, w io.Writer, e io.Writer, extras []language.File) error {
	err := s.CreateFile(j.sourceName, r.Source)
	if err != nil {
		return err
	}

	classPath := "."
	for ind := range j.jars {
		if err := s.CreateFile(j.jars[ind].Name, bytes.NewBuffer(j.jars[ind].Contents)); err != nil {
			return err
		}
		classPath += ":"+j.jars[ind].Name
	}

	if _, err := s.AddArg("--open-files=2048").SetMaxProcesses(-1).Env().TimeLimit(10*time.Second).MemoryLimit(4*256000).Stdout(e).Stderr(e).WorkingDirectory(s.Pwd()).Run("/usr/bin/javac -cp "+classPath+" "+j.sourceName, false); err != nil {
		return err
	}

	bin, err := s.GetFile(j.className)
	if err != nil {
		return err
	}

	_, err = io.Copy(w, bin)

	return err
}

func (j *Java) Run(s language.Sandbox, binary, stdin io.Reader, stdout io.Writer, tl time.Duration, ml int) (language.Status, error) {
	stat := language.Status{}
	stat.Verdict = language.VERDICT_XX

	if err := s.CreateFile(j.className, binary); err != nil {
		return stat, err
	}

	classPath := "."
	for ind := range j.jars {
		if err := s.CreateFile(j.jars[ind].Name, bytes.NewBuffer(j.jars[ind].Contents)); err != nil {
			return stat, err
		}
		classPath += ":"+j.jars[ind].Name
	}

	return s.SetMaxProcesses(-1).Env().Stdin(stdin).Stdout(stdout).Stderr(os.Stderr).TimeLimit(tl).MemoryLimit(ml/1024).WorkingDirectory(s.Pwd()).Run("/usr/bin/java -cp "+classPath+" "+j.execName, true)
}

func init() {
	language.Register("java", New("main"))
}
