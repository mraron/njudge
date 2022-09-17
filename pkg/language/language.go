package language

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"sort"
	"testing"
	"time"
)

type Verdict int

const (
	VERDICT_OK Verdict = 1 << iota
	VERDICT_TL
	VERDICT_ML
	VERDICT_RE
	VERDICT_XX
	VERDICT_CE
)

func (v Verdict) String() string {
	switch v {
	case VERDICT_OK:
		return "OK"
	case VERDICT_TL:
		return "TL"
	case VERDICT_ML:
		return "ML"
	case VERDICT_RE:
		return "RE"
	case VERDICT_XX:
		return "XX"
	case VERDICT_CE:
		return "CE"
	}
	return "??"
}

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
	DefaultFileName() string
	Compile(Sandbox, File, io.Writer, io.Writer, []File) error
	Run(Sandbox, io.Reader, io.Reader, io.Writer, time.Duration, int) (Status, error)
}

type LanguageTest struct {
	Sandbox         Sandbox
	Language        Language
	Source          string
	ExpectedVerdict Verdict
	Input           string
	ExpectedOutput  string
	TimeLimit       time.Duration
	MemoryLimit     int
}

func (test LanguageTest) Run(t *testing.T) {
	err := test.Sandbox.Init(log.New(ioutil.Discard, "", 0))
	if err != nil {
		t.Error(err)
	}

	src := bytes.NewBufferString(test.Source)
	bin := &bytes.Buffer{}
	stderr := &bytes.Buffer{}

	err = test.Language.Compile(test.Sandbox, File{test.Language.DefaultFileName(), src}, bin, stderr, []File{})
	stderrContent := stderr.String()

	if (test.ExpectedVerdict&VERDICT_CE == 0 && err != nil) || (test.ExpectedVerdict&VERDICT_CE != 0 && err == nil && stderrContent == "") {
		t.Errorf("error: %v stderr: %s", err, stderrContent)
	}

	if test.ExpectedVerdict&VERDICT_CE == 0 {
		output := &bytes.Buffer{}
		status, err := test.Language.Run(test.Sandbox, bin, bytes.NewBufferString(test.Input), output, test.TimeLimit, test.MemoryLimit)

		outputContent := output.String()
		if status.Verdict&test.ExpectedVerdict == 0 || err != nil || outputContent != test.ExpectedOutput {
			t.Errorf("EXPECTED %s got %s, source %q\n error: %v status: %v output: %q", test.ExpectedVerdict, status.Verdict, test.Source, err, status, outputContent)
		}
	}

	err = test.Sandbox.Cleanup()
	if err != nil {
		t.Errorf("cleanup err: %v", err.Error())
	}

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
