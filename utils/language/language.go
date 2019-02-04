package language

import (
	"io"
	"sort"
	"time"
)

type Verdict int

const (
	VERDICT_OK Verdict = iota
	VERDICT_TL
	VERDICT_ML
	VERDICT_RE
	VERDICT_XX
)

type File struct {
	Name   string
	Source io.Reader
}

type Status struct {
	Verdict Verdict
	Signal  int
	Memory  int
	Time    time.Duration
}

type Language interface {
	Id() string
	Name() string
	InsecureCompile(string, io.Reader, io.Writer, io.Writer) error
	Compile(Sandbox, File, io.Writer, io.Writer, []File) error
	Run(Sandbox, io.Reader, io.Reader, io.Writer, time.Duration, int) (Status, error)
}

var langList map[string]Language

func Register(name string, l Language) {
	langList[name] = l
}

func List() []Language {
	ans := make([]Language, len(langList))

	ind := 0
	for _, val := range langList {
		ans[ind] = val
		ind++
	}

	sort.Slice(ans, func(i, j int) bool {
		return ans[i].Name() < ans[j].Name()
	})

	return ans
}

func Get(name string) Language {
	if val, ok := langList[name]; ok {
		return val
	}

	return nil
}

func init() {
	langList = make(map[string]Language)
}
