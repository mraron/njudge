package language

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"sort"
	"testing"
	"time"
)

type Verdict int

const (
	VERDICT_OK Verdict = iota
	VERDICT_TL
	VERDICT_ML
	VERDICT_RE
	VERDICT_XX
	VERDICT_CE
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
	DefaultFileName() string
	InsecureCompile(string, io.Reader, io.Writer, io.Writer) error
	Compile(Sandbox, File, io.Writer, io.Writer, []File) error
	Run(Sandbox, io.Reader, io.Reader, io.Writer, time.Duration, int) (Status, error)
}

type LanguageTest struct {
	Language Language
	Source string
	ExpectedVerdict Verdict
	Input string
	ExpectedOutput string
	TimeLimit time.Duration
	MemoryLimit int
}

func (test LanguageTest) Run(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	sandbox := NewIsolateSandbox(500+rand.Intn(100))

	sandbox.Init(log.New(ioutil.Discard, "", 0))

	src := bytes.NewBufferString(test.Source)
	bin := &bytes.Buffer{}
	stderr := &bytes.Buffer{}

	err := test.Language.Compile(sandbox, File{test.Language.DefaultFileName(), src}, bin, stderr, []File{})

	stderr_content := stderr.String()

	if (test.ExpectedVerdict != VERDICT_CE && err != nil) || (test.ExpectedVerdict == VERDICT_CE && (err==nil || stderr_content=="")){
		t.Fatalf("error: %s stderr: %s", err, stderr_content)
	}

	if test.ExpectedVerdict != VERDICT_CE {
		output := &bytes.Buffer{}
		status, err := test.Language.Run(sandbox, bin, bytes.NewBufferString(test.Input), output, test.TimeLimit, test.MemoryLimit)

		output_content := output.String()
		if status.Verdict != test.ExpectedVerdict || err != nil || output_content != test.ExpectedOutput {
			t.Fatalf("source %s\n error: %s status: %v output: %q", test.Source, err, status, output_content)
		}
	}

	err = sandbox.Cleanup()
	if err != nil {
		t.Fatalf("cleanup err: %s", err.Error())
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
