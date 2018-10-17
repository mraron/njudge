package batch

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/mraron/njudge/utils/language"
	"github.com/mraron/njudge/utils/problems"
	"io"
	"io/ioutil"
	"os"
	"strconv"
)

type Batch struct {
}

func (b Batch) Name() string {
	return "batch"
}

func (b Batch) Compile(jinfo problems.JudgingInformation, sandbox language.Sandbox, lang language.Language, src io.Reader, dest io.Writer) (io.Reader, error) {
	lst, found := jinfo.Languages(), false

	for _, l := range lst {
		if l.Name() == lang.Name() {
			found = true
		}
	}

	if !found {
		return nil, errors.New(fmt.Sprintf("running problem %s on %s tasktype, language %s is not supported", jinfo.Name(), b.Name(), lang.Name()))
	}

	buf := &bytes.Buffer{}

	err := lang.Compile(sandbox, language.File{"main", src}, buf, dest, nil)
	fmt.Println(err)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

func truncate(s string) string {
	if len(s) < 256 {
		return s
	}

	return s[:255] + "..."
}

func (b Batch) Run(jinfo problems.JudgingInformation, s language.Sandbox, lang language.Language, bin io.Reader, testNotifier chan string, statusNotifier chan problems.Status) (problems.Status, error) {
	var (
		ans            problems.Status
		skeleton       = jinfo.StatusSkeleton()
		binaryContents []byte
		err            error
	)

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

	fmt.Println(skeleton)
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

					testcase, err := os.Open(testLocation)
					if err != nil {
						return ans, err
					}

					stdout := &bytes.Buffer{}

					answerFile, err := os.Open(answerLocation)
					if err != nil {
						return ans, err
					}

					answerContents, err := ioutil.ReadAll(answerFile)
					if err != nil {
						return ans, err
					}

					res, err := lang.Run(s, bytes.NewReader(binaryContents), testcase, stdout, tc.TimeLimit, tc.MemoryLimit)
					if err != nil {
						return ans, err
					}

					fmt.Println(res, res.Verdict, language.VERDICT_OK, "!!!!!!!!!!!!!!")

					if res.Verdict == language.VERDICT_OK {
						checkerOutput := &bytes.Buffer{}
						programOutput := stdout.String()

						expectedOutput := string(answerContents)

						tmpfile, err := ioutil.TempFile("/tmp", "OutputOfProgram")
						if err != nil {
							return ans, err
						}

						if _, err := tmpfile.Write([]byte(programOutput)); err != nil {
							return ans, err
						}

						if err := tmpfile.Close(); err != nil {
							return ans, err
						}

						defer os.Remove(tmpfile.Name())

						err = jinfo.Check(testLocation, tmpfile.Name(), answerLocation, checkerOutput, checkerOutput)

						testset.Testcases = append(testset.Testcases, tc)
						group.Testcases = append(group.Testcases, tc)

						testset.Testcases[len(testset.Testcases)-1].CheckerOutput = truncate(checkerOutput.String())
						testset.Testcases[len(testset.Testcases)-1].Output = truncate(stdout.String())
						testset.Testcases[len(testset.Testcases)-1].ExpectedOutput = truncate(expectedOutput)
						testset.Testcases[len(testset.Testcases)-1].MemoryUsed = res.Memory
						testset.Testcases[len(testset.Testcases)-1].TimeSpent = res.Time

						group.Testcases[len(group.Testcases)-1].CheckerOutput = truncate(checkerOutput.String())
						group.Testcases[len(group.Testcases)-1].Output = truncate(stdout.String())
						group.Testcases[len(group.Testcases)-1].ExpectedOutput = truncate(expectedOutput)
						group.Testcases[len(group.Testcases)-1].MemoryUsed = res.Memory
						group.Testcases[len(group.Testcases)-1].TimeSpent = res.Time

						if err == nil {
							testset.Testcases[len(testset.Testcases)-1].VerdictName = problems.VERDICT_AC
							testset.Testcases[len(testset.Testcases)-1].Score = testset.Testcases[len(testset.Testcases)-1].MaxScore

							group.Testcases[len(group.Testcases)-1].VerdictName = problems.VERDICT_AC
							group.Testcases[len(group.Testcases)-1].Score = group.Testcases[len(group.Testcases)-1].MaxScore
						} else {
							ac = false

							testset.Testcases[len(testset.Testcases)-1].VerdictName = problems.VERDICT_WA
							testset.Testcases[len(testset.Testcases)-1].Score = 0

							group.Testcases[len(group.Testcases)-1].VerdictName = problems.VERDICT_WA
							group.Testcases[len(group.Testcases)-1].Score = 0

							if skeleton.FeedbackType != problems.FEEDBACK_IOI {
								return ans, nil
							}
						}
					} else {
						ac = false

						curr := tc
						curr.Testset = ts.Name
						switch res.Verdict {
						case language.VERDICT_RE:
							curr.VerdictName = problems.VERDICT_RE
						case language.VERDICT_XX:
							curr.VerdictName = problems.VERDICT_XX
						case language.VERDICT_ML:
							curr.VerdictName = problems.VERDICT_ML
						case language.VERDICT_TL:
							curr.VerdictName = problems.VERDICT_TL
						}

						curr.Group = g.Name
						curr.MemoryUsed = res.Memory
						curr.TimeSpent = res.Time
						curr.Score = 0
						curr.Output = truncate(stdout.String()) //now it's stderr
						curr.ExpectedOutput = truncate(string(answerContents))

						testset.Testcases = append(testset.Testcases, curr)
						group.Testcases = append(group.Testcases, curr)

						if skeleton.FeedbackType != problems.FEEDBACK_IOI {
							return ans, nil
						}
					}
				} else {
					group.Testcases = append(group.Testcases, tc)
					testset.Testcases = append(testset.Testcases, tc)
				}
			}

			groupAC[g.Name] = ac
		}
	}

	return ans, nil
}

func init() {
	problems.RegisterTaskType(Batch{})

}
