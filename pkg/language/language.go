package language

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
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

	Test(Sandbox) error
}

type LanguageTest struct {
	Language        Language
	Source          string
	ExpectedVerdict Verdict
	Input           string
	ExpectedOutput  string
	TimeLimit       time.Duration
	MemoryLimit     int
}

func (test LanguageTest) Run(sandbox Sandbox) error {
	err := sandbox.Init(log.New(ioutil.Discard, "", 0))
	if err != nil {
		return err
	}

	src := bytes.NewBufferString(test.Source)
	bin := &bytes.Buffer{}
	stderr := &bytes.Buffer{}

	err = test.Language.Compile(sandbox, File{test.Language.DefaultFileName(), src}, bin, stderr, []File{})
	stderrContent := stderr.String()

	if (test.ExpectedVerdict&VERDICT_CE == 0 && err != nil) || (test.ExpectedVerdict&VERDICT_CE != 0 && err == nil && stderrContent == "") {
		return fmt.Errorf("error: %w stderr: %s", err, stderrContent)
	}

	err = sandbox.Cleanup()
	if err != nil {
		return fmt.Errorf("cleanup err: %w", err)
	}

	if test.ExpectedVerdict&VERDICT_CE == 0 {
		err := sandbox.Init(log.New(io.Discard, "", 0))
		if err != nil {
			return err
		}

		output := &bytes.Buffer{}
		status, err := test.Language.Run(sandbox, bin, bytes.NewBufferString(test.Input), output, test.TimeLimit, test.MemoryLimit)

		outputContent := output.String()
		if status.Verdict&test.ExpectedVerdict == 0 || err != nil || outputContent != test.ExpectedOutput {
			return fmt.Errorf("EXPECTED %s got %s, source %q\n error: %w status: %v output: %q expected output: %q", test.ExpectedVerdict, status.Verdict, test.Source, err, status, outputContent, test.ExpectedOutput)
		}

		err = sandbox.Cleanup()
		if err != nil {
			return fmt.Errorf("cleanup err: %w", err)
		}
	}

	return nil
}
