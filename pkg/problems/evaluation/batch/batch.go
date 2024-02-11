package batch

import (
	"bytes"
	"context"
	"fmt"
	"github.com/mraron/njudge/pkg/language/memory"
	"github.com/mraron/njudge/pkg/language/sandbox"
	"io"
	"io/ioutil"
	"os"
	"strconv"

	"github.com/mraron/njudge/pkg/language"
	"github.com/mraron/njudge/pkg/problems"
)

type Batch struct {
	PrepareFilesF func(*CompileContext) (sandbox.File, []sandbox.File, error)

	InitF      func(*RunContext) error
	RunF       func(*RunContext, *problems.Group, *problems.Testcase) (sandbox.Status, error)
	CheckOKF   func(*RunContext, sandbox.Status, *problems.Group, *problems.Testcase) error
	CheckFailF func(*RunContext, sandbox.Status, *problems.Group, *problems.Testcase) error
	CleanupF   func(*RunContext) error
}

func New() Batch {
	return Batch{
		PrepareFilesF: PrepareFiles,
		InitF:         Init,
		RunF:          Run,
		CheckOKF:      CheckOK,
		CheckFailF:    CheckFail,
		CleanupF:      Cleanup,
	}
}

func truncate(s string) string {
	if len(s) < 256 {
		return s
	}

	return s[:255] + "..."
}

type RunContext struct {
	Problem         problems.Judgeable
	SandboxProvider *sandbox.ChanProvider
	Sandbox         sandbox.Sandbox
	Lang            language.Language
	Binary          []byte
	TestChan        chan string
	StatusChan      chan problems.Status
	Stdout          *bytes.Buffer

	Store map[string]interface{}
}

type CompileContext struct {
	Problem problems.Judgeable
	Sandbox sandbox.Sandbox
	Lang    language.Language
	Source  io.Reader
	Binary  io.Writer
}

func PrepareFiles(ctx *CompileContext) (sandbox.File, []sandbox.File, error) {
	lst, found := ctx.Problem.Languages(), false

	for _, l := range lst {
		if l.ID() == ctx.Lang.ID() {
			found = true
		}
	}

	if !found {
		return sandbox.File{}, nil, fmt.Errorf("language %s is not supported", ctx.Lang.ID())
	}

	return sandbox.File{Name: ctx.Lang.DefaultFilename(), Source: io.NopCloser(ctx.Source)}, nil, nil
}

func Init(*RunContext) error {
	return nil
}

func Run(ctx *RunContext, group *problems.Group, testcase *problems.Testcase) (sandbox.Status, error) {
	inputFile, _ := ctx.Problem.InputOutputFiles()
	testLocation, answerLocation := testcase.InputPath, testcase.AnswerPath
	input, err := os.Open(testLocation)
	if err != nil {
		testcase.VerdictName = problems.VerdictXX
		return sandbox.Status{}, err
	}
	defer input.Close()

	answerFile, err := os.Open(answerLocation)
	if err != nil {
		testcase.VerdictName = problems.VerdictXX
		return sandbox.Status{}, err
	}
	defer answerFile.Close()

	if inputFile != "" {
		if err := sandbox.CreateFile(ctx.Sandbox, sandbox.File{inputFile, input}); err != nil {
			testcase.VerdictName = problems.VerdictXX
			return sandbox.Status{}, err
		}
		input = nil
	}

	res, err := ctx.Lang.Run(context.TODO(), ctx.Sandbox, sandbox.File{
		"a.out",
		io.NopCloser(bytes.NewReader(ctx.Binary)),
	}, input, ctx.Stdout, testcase.TimeLimit, memory.Amount(testcase.MemoryLimit))

	if err != nil {
		testcase.VerdictName = problems.VerdictXX
		return *res, err
	}

	testcase.MemoryUsed = res.Memory
	testcase.TimeSpent = res.Time

	return *res, nil
}

func CheckOK(ctx *RunContext, res sandbox.Status, group *problems.Group, testcase *problems.Testcase) error {
	programOutput := ctx.Stdout.String()
	answerContents, err := ioutil.ReadFile(testcase.AnswerPath)
	if err != nil {
		testcase.VerdictName = problems.VerdictXX
		return err
	}

	tmpfile, err := os.CreateTemp("/tmp", "OutputOfProgram")
	if err != nil {
		testcase.VerdictName = problems.VerdictXX
		return err
	}

	if _, err := tmpfile.Write([]byte(programOutput)); err != nil {
		testcase.VerdictName = problems.VerdictXX
		return err
	}

	if err := tmpfile.Close(); err != nil {
		testcase.VerdictName = problems.VerdictXX
		return err
	}

	defer os.Remove(tmpfile.Name())

	testcase.OutputPath = tmpfile.Name()

	err = ctx.Problem.Checker().Check(nil, testcase)

	testcase.Output = truncate(programOutput)
	testcase.ExpectedOutput = truncate(string(answerContents))

	return err
}

func CheckFail(ctx *RunContext, res sandbox.Status, group *problems.Group, testcase *problems.Testcase) error {
	answerContents, err := ioutil.ReadFile(testcase.AnswerPath)
	if err != nil {
		testcase.VerdictName = problems.VerdictXX
		return err
	}

	switch res.Verdict {
	case sandbox.VerdictRE:
		testcase.VerdictName = problems.VerdictRE
	case sandbox.VerdictXX:
		testcase.VerdictName = problems.VerdictXX
	case sandbox.VerdictML:
		testcase.VerdictName = problems.VerdictML
	case sandbox.VerdictTL:
		testcase.VerdictName = problems.VerdictTL
	}

	testcase.Output = truncate(ctx.Stdout.String())
	testcase.ExpectedOutput = truncate(string(answerContents))

	testcase.Score = 0
	return nil
}

func Cleanup(ctx *RunContext) error {
	return nil
}

func (b Batch) Name() string {
	return "batch"
}

func (b Batch) Compile(jinfo problems.Judgeable, sandbox sandbox.Sandbox, lang language.Language, src io.Reader, dest io.Writer) (io.Reader, error) {
	file, extras, err := b.PrepareFilesF(&CompileContext{
		Problem: jinfo,
		Sandbox: sandbox,
		Lang:    lang,
		Source:  src,
		Binary:  dest,
	})
	if err != nil {
		return nil, err
	}

	buf := &bytes.Buffer{}
	if _, err := lang.Compile(context.TODO(), sandbox, file, dest, extras); err != nil {
		return nil, err
	}

	return buf, nil
}

func (b Batch) Run(judging problems.Judgeable, sp *sandbox.ChanProvider, lang language.Language, bin io.Reader, testNotifier chan string, statusNotifier chan problems.Status) (problems.Status, error) {
	var (
		ans            problems.Status
		skeleton       *problems.Status
		binaryContents []byte
		err            error
		s              sandbox.Sandbox
	)

	if skeleton, err = judging.StatusSkeleton(""); err != nil {
		return ans, err
	}

	s, err = sp.Get()
	if err != nil {
		ans.Compiled = false
		ans.CompilerOutput = err.Error()
		return ans, err
	}
	defer sp.Put(s)

	binaryContents, err = ioutil.ReadAll(bin)
	if err != nil {
		ans.Compiled = false
		ans.CompilerOutput = err.Error()
		return ans, err
	}
	ans.Compiled = true
	ans.FeedbackType = skeleton.FeedbackType

	defer func() {
		close(statusNotifier)
		close(testNotifier)
	}()

	groupAC := make(map[string]bool)

	dependenciesOK := func(deps []string) bool {
		for _, dep := range deps {
			if !groupAC[dep] {
				return false
			}
		}
		return true
	}

	ctx := RunContext{
		Problem:         judging,
		SandboxProvider: sp,
		Sandbox:         s,
		Lang:            lang,
		Binary:          binaryContents,
		TestChan:        testNotifier,
		StatusChan:      statusNotifier,

		Store: map[string]interface{}{},
	}

	if err := b.InitF(&ctx); err != nil {
		return ans, err
	}

	for _, ts := range skeleton.Feedback {
		ans.Feedback = append(ans.Feedback, problems.Testset{Name: ts.Name})
		testset := &ans.Feedback[len(ans.Feedback)-1]

		for _, g := range ts.Groups {
			testset.Groups = append(testset.Groups, problems.Group{Name: g.Name, Scoring: g.Scoring})

			group := &testset.Groups[len(testset.Groups)-1]
			group.Testcases = append(group.Testcases, g.Testcases...)
		}
	}

	testCache := make(map[string]*problems.Testcase)
	for tsind := range ans.Feedback {
		testset := &ans.Feedback[tsind]

		for gind := range testset.Groups {
			group := &testset.Groups[gind]

			ac := true

			for tcind := range group.Testcases {
				tc := &group.Testcases[tcind]

				statusNotifier <- ans
				testNotifier <- strconv.Itoa(tc.Index)

				if tc.VerdictName != problems.VerdictDR {
					continue
				}

				if _, ok := testCache[tc.InputPath]; ok {
					tmpIndex, tmpGroup := tc.Index, tc.Group
					*tc = *testCache[tc.InputPath]
					tc.Index = tmpIndex
					tc.Group = tmpGroup
					tc.Score = 0.0
					tc.MaxScore = 0.0
					continue
				}
				testCache[tc.InputPath] = tc

				if ans.FeedbackType == problems.FeedbackLazyIOI && !ac {
					continue
				}

				if dependenciesOK(group.Dependencies) {
					ctx.Stdout = &bytes.Buffer{}
					res, err := b.RunF(&ctx, group, tc)

					if err != nil {
						tc.VerdictName = problems.VerdictXX
						return ans, err
					} else if res.Verdict == sandbox.VerdictOK {
						if err := b.CheckOKF(&ctx, res, group, tc); err != nil {
							tc.VerdictName = problems.VerdictXX
							return ans, err
						} else {
							if tc.VerdictName == problems.VerdictWA || tc.VerdictName == problems.VerdictPE {
								ac = false
								if skeleton.FeedbackType != problems.FeedbackIOI && skeleton.FeedbackType != problems.FeedbackLazyIOI {
									return ans, nil
								}
							}
						}
					} else {
						if err := b.CheckFailF(&ctx, res, group, tc); err != nil {
							tc.VerdictName = problems.VerdictXX
							return ans, err
						}

						ac = false
						if skeleton.FeedbackType != problems.FeedbackIOI && skeleton.FeedbackType != problems.FeedbackLazyIOI {
							return ans, nil
						}
					}
				} else {
					ac = false
				}
			}

			groupAC[group.Name] = ac
		}
	}

	return ans, Cleanup(&ctx)
}