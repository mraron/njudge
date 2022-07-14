package communication

import (
	"bytes"
	"github.com/mraron/njudge/pkg/language"
	"github.com/mraron/njudge/pkg/problems"
	"github.com/mraron/njudge/pkg/problems/tasktype/stub"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"syscall"
)

type Communication struct {
}

func (b Communication) Name() string {
	return "communication"
}

func (b Communication) Compile(jinfo problems.Judgeable, sandbox language.Sandbox, lang language.Language, src io.Reader, dest io.Writer) (io.Reader, error) {
	return stub.Stub{}.Compile(jinfo, sandbox, lang, src, dest)
}

func truncate(s string) string {
	if len(s) < 256 {
		return s
	}

	return s[:255] + "..."
}

func (b Communication) Run(jinfo problems.Judgeable, sp *language.SandboxProvider, lang language.Language, bin io.Reader, testNotifier chan string, statusNotifier chan problems.Status) (problems.Status, error) {
	var (
		ans            problems.Status
		skeleton       *problems.Status
		binaryContents []byte
		err            error
	)

	if skeleton, err = jinfo.StatusSkeleton(""); err != nil {
		return ans, err
	}

	interactorSandbox, err := sp.Get()
	if err != nil {
		ans.Compiled = false
		ans.CompilerOutput = err.Error()
		return ans, err
	}
	defer sp.Put(interactorSandbox)

	interactorPath := ""
	for _, f := range jinfo.Files() {
		if f.Role == "interactor" {
			interactorPath = f.Path
		}
	}

	f, err := os.Open(interactorPath)
	if err != nil {
		ans.Compiled = false
		ans.CompilerOutput = "Can't find interactor"
		return ans, err
	}
	defer f.Close()

	interactorSandbox.CreateFile("interactor", f)
	interactorSandbox.MakeExecutable("interactor")

	s, err := sp.Get()
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
		close(testNotifier)
		close(statusNotifier)
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

	for _, ts := range skeleton.Feedback {
		ans.Feedback = append(ans.Feedback, problems.Testset{Name: ts.Name})
		testset := &ans.Feedback[len(ans.Feedback)-1]

		for _, g := range ts.Groups {
			testset.Groups = append(testset.Groups, problems.Group{Name: g.Name, Scoring: g.Scoring})
			group := &testset.Groups[len(testset.Groups)-1]

			ac := true

			for _, tc := range g.Testcases {
				testNotifier <- strconv.Itoa(tc.Index)
				statusNotifier <- ans

				if dependenciesOK(g.Dependencies) {
					testLocation, answerLocation := tc.InputPath, tc.AnswerPath

					testFile, err := os.Open(testLocation)
					if err != nil {
						tc.VerdictName = problems.VerdictXX
						return ans, err
					}
					defer testFile.Close()

					err = interactorSandbox.CreateFile("inp", testFile)
					if err != nil {
						tc.VerdictName = problems.VerdictXX
						return ans, err
					}

					stdout := &bytes.Buffer{}

					answerFile, err := os.Open(answerLocation)
					if err != nil {
						tc.VerdictName = problems.VerdictXX
						return ans, err
					}
					defer answerFile.Close()

					answerContents, err := ioutil.ReadAll(answerFile)
					if err != nil {
						tc.VerdictName = problems.VerdictXX
						return ans, err
					}

					var res language.Status

					os.Remove("/tmp/fifo1" + interactorSandbox.Id())
					os.Remove("/tmp/fifo2" + interactorSandbox.Id())

					err = syscall.Mkfifo(filepath.Join("/tmp", "fifo1"+interactorSandbox.Id()), 0766)
					if err != nil {
						tc.VerdictName = problems.VerdictXX
						return ans, err
					}

					err = syscall.Mkfifo(filepath.Join("/tmp", "fifo2"+interactorSandbox.Id()), 0766)
					if err != nil {
						tc.VerdictName = problems.VerdictXX
						return ans, err
					}

					fifo1, err := os.OpenFile(filepath.Join("/tmp", "fifo1"+interactorSandbox.Id()), os.O_RDWR, 0766)
					if err != nil {
						tc.VerdictName = problems.VerdictXX
						return ans, err
					}
					defer fifo1.Close()

					fifo2, err := os.OpenFile(filepath.Join("/tmp", "fifo2"+interactorSandbox.Id()), os.O_RDWR, 0766)
					if err != nil {
						tc.VerdictName = problems.VerdictXX
						return ans, err
					}
					defer fifo2.Close()

					done := make(chan int, 1)

					go func() {
						// @TODO check res and err of interactor
						interactorSandbox.Stdin(fifo1).Stdout(fifo2).Stderr(os.Stderr).TimeLimit(tc.TimeLimit).MemoryLimit(tc.MemoryLimit).Run("interactor inp out", true)

						done <- 1
					}()

					res, err = lang.Run(s, bytes.NewReader(binaryContents), fifo2, fifo1, tc.TimeLimit, tc.MemoryLimit)
					<-done

					if err != nil {
						tc.VerdictName = problems.VerdictXX
						return ans, err
					}

					if res.Verdict == language.VERDICT_OK {
						expectedOutput := string(answerContents)

						tc.OutputPath = filepath.Join(interactorSandbox.Pwd(), "out")

						err = jinfo.Check(&tc)

						tc.Output = truncate(stdout.String())
						tc.ExpectedOutput = truncate(expectedOutput)
						tc.MemoryUsed = res.Memory
						tc.TimeSpent = res.Time

						group.Testcases = append(group.Testcases, tc)

						if err == nil {
							if tc.VerdictName == problems.VerdictWA || tc.VerdictName == problems.VerdictPE {
								ac = false
								if skeleton.FeedbackType != problems.FeedbackIOI {
									return ans, nil
								}
							}
						} else {
							return ans, err
						}
					} else {
						ac = false

						curr := tc
						curr.Testset = ts.Name
						switch res.Verdict {
						case language.VERDICT_RE:
							curr.VerdictName = problems.VerdictRE
						case language.VERDICT_XX:
							curr.VerdictName = problems.VerdictXX
						case language.VERDICT_ML:
							curr.VerdictName = problems.VerdictML
						case language.VERDICT_TL:
							curr.VerdictName = problems.VerdictTL
						}

						curr.Group = g.Name
						curr.MemoryUsed = res.Memory
						curr.TimeSpent = res.Time
						curr.Score = 0
						curr.Output = truncate(stdout.String()) //now it's stderr
						curr.ExpectedOutput = truncate(string(answerContents))

						group.Testcases = append(group.Testcases, curr)

						if skeleton.FeedbackType != problems.FeedbackIOI {
							return ans, nil
						}
					}
				} else {
					group.Testcases = append(group.Testcases, tc)
				}
			}

			groupAC[g.Name] = ac
		}
	}

	return ans, nil
}

func init() {
	problems.RegisterTaskType(Communication{})
}
