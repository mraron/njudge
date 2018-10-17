package outputonly

import (
	"archive/zip"
	"bytes"
	"fmt"
	"github.com/mraron/njudge/utils/language"
	"github.com/mraron/njudge/utils/problems"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

type OutputOnly struct {
}

func (o OutputOnly) Name() string {
	return "outputonly"
}

func (o OutputOnly) Compile(jinfo problems.JudgingInformation, s language.Sandbox, l language.Language, src io.Reader, errw io.Writer) (io.Reader, error) {
	zipContents, err := ioutil.ReadAll(src)
	if err != nil {
		errw.Write([]byte(err.Error()))
		return nil, err
	}

	readerAt := bytes.NewReader(zipContents)

	_, err = zip.NewReader(readerAt, int64(len(zipContents)))
	if err != nil {
		errw.Write([]byte(err.Error()))
		return nil, err
	}

	return bytes.NewReader(zipContents), nil
}

func (o OutputOnly) Run(jinfo problems.JudgingInformation, s language.Sandbox, lang language.Language, bin io.Reader, testNotifier chan string, statusNotifier chan problems.Status) (problems.Status, error) {
	defer func() {
		close(testNotifier)
		close(statusNotifier)
	}()

	skeleton := jinfo.StatusSkeleton()

	fmt.Println(skeleton)
	ans := problems.Status{}

	ans.Compiled = true

	ans.Feedback = make([]problems.Testset, 1)
	ans.Feedback[0] = problems.Testset{"main", make([]problems.Group, 0), make([]problems.Testcase, 0)}

	ans.FeedbackType = problems.FEEDBACK_IOI

	zipContents, err := ioutil.ReadAll(bin)
	if err != nil {
		ans.Compiled = false
		fmt.Println(err, "err1")
		return ans, err
	}

	readerAt := bytes.NewReader(zipContents)

	zip, err := zip.NewReader(readerAt, int64(len(zipContents)))
	if err != nil {
		fmt.Println(err, "err2")
		ans.Compiled = false
		return ans, err
	}

	ans.Feedback[0].Groups = append(ans.Feedback[0].Groups, problems.Group{"subtask1", problems.SCORING_SUM, make([]problems.Testcase, 0), make([]string, 0)})
	for _, tc := range skeleton.Feedback[0].Testcases {
		inputName := tc.InputPath
		outputName := tc.AnswerPath

		ans.Feedback[0].Testcases = append(ans.Feedback[0].Testcases, problems.Testcase{Testset: "main", VerdictName: problems.VERDICT_RE, Score: 0.0, MaxScore: 0.0})
		currentCase := &ans.Feedback[0].Testcases[len(ans.Feedback[0].Testcases)-1]

		for _, file := range zip.File {
			fmt.Println(file.Name, "!!!!!!!!!!!!!!")
			if file.Name == filepath.Base(outputName) {
				stdout, stderr := &bytes.Buffer{}, &bytes.Buffer{}

				fileinzip, err := file.Open()
				if err != nil {
					fmt.Println(err, "err1")
					currentCase.VerdictName = problems.VERDICT_XX
					currentCase.CheckerOutput = err.Error()

					ans.Feedback[0].Groups[0].Testcases = append(ans.Feedback[0].Groups[0].Testcases, *currentCase)
					break
				}

				full, err := ioutil.ReadAll(fileinzip)
				if err != nil {
					fmt.Println(err, "err2")
					currentCase.VerdictName = problems.VERDICT_XX
					currentCase.CheckerOutput = err.Error()

					ans.Feedback[0].Groups[0].Testcases = append(ans.Feedback[0].Groups[0].Testcases, *currentCase)
					break
				}

				tmpfile, err := ioutil.TempFile("/tmp", "FileInZip")
				if err != nil {
					fmt.Println(err, "err25")
					currentCase.VerdictName = problems.VERDICT_XX
					currentCase.CheckerOutput = err.Error()

					ans.Feedback[0].Groups[0].Testcases = append(ans.Feedback[0].Groups[0].Testcases, *currentCase)
					break
				}

				if _, err := tmpfile.Write([]byte(full)); err != nil {
					fmt.Println(err, "err3")
					currentCase.VerdictName = problems.VERDICT_XX
					currentCase.CheckerOutput = err.Error()

					ans.Feedback[0].Groups[0].Testcases = append(ans.Feedback[0].Groups[0].Testcases, *currentCase)
					break
				}

				if err := tmpfile.Close(); err != nil {
					fmt.Println(err, "err4")
					currentCase.VerdictName = problems.VERDICT_XX
					currentCase.CheckerOutput = err.Error()

					ans.Feedback[0].Groups[0].Testcases = append(ans.Feedback[0].Groups[0].Testcases, *currentCase)
					break
				}

				defer os.Remove(tmpfile.Name())

				err = jinfo.Check(inputName, tmpfile.Name(), outputName, stdout, stderr)
				fmt.Println(err, "LALALAAL")
				currentCase.CheckerOutput = stderr.String()
				fmt.Sscanf(stdout.String(), "%f/%f", &currentCase.Score, &currentCase.MaxScore)

				if err == nil {
					currentCase.VerdictName = problems.VERDICT_AC
					ans.Feedback[0].Groups[0].Testcases = append(ans.Feedback[0].Groups[0].Testcases, *currentCase)
				} else {
					currentCase.VerdictName = problems.VERDICT_WA
					ans.Feedback[0].Groups[0].Testcases = append(ans.Feedback[0].Groups[0].Testcases, *currentCase)
				}

				break
			}
		}
	}

	return ans, nil
}

func init() {
	problems.RegisterTaskType(OutputOnly{})
}
