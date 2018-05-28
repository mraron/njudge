package zipoo

import (
	"github.com/mraron/njudge/utils/problems"
	"github.com/mraron/njudge/utils/language"
	"io"
	"io/ioutil"
	"bytes"
	"archive/zip"
	"fmt"
	"os/exec"
	"path/filepath"
	"os"
	"encoding/json"
)

type Name struct {
	Language string
	Value string
}

type Statement struct {
	Language string
	Path string
	Type string
}

type Attachment struct {
	Name string
	Path string
}

type Problem struct {
	Path string
	GeneratedStatementList []problems.Content
	AttachmentList []problems.Attachment

	ShortName string `json:"shortname"`
	Names []Name
	StatementList []Statement `json:"statements"`
	AttachmentInfo []Attachment `json:"attachments"`

	TestCount int `json:"testcount"`
	InputPattern string `json:"inputpattern"`
	OutputPattern string `json:"outputpattern"`
}

func (p Problem) Name() string {
	return p.ShortName
}

func (p Problem) Titles() []problems.Content {
	ans := make([]problems.Content, len(p.Names))
	for ind, val := range p.Names {
		ans[ind] = problems.Content{val.Language, []byte(val.Value), "text"}
	}

	return ans
}

func (p Problem) Statements() []problems.Content {
	return p.GeneratedStatementList
}

func (p Problem) HTMLStatements() []problems.Content {
	return problems.FilterContentArray(p.GeneratedStatementList, "text/html")
}

func (p Problem) PDFStatements() []problems.Content {
	return problems.FilterContentArray(p.GeneratedStatementList, "application/pdf")
}

func (p Problem) MemoryLimit() int {
	return 0
}

func (p Problem) TimeLimit() int {
	return 0
}

func (p Problem) InputOutputFiles() (string, string) {
	return "", ""
}

func (p Problem) Interactive() bool {
	return false
}

func (p Problem) Languages() []language.Language {
	return []language.Language{language.Get("zip")}
}

func (p Problem) Attachments() []problems.Attachment {
	return p.AttachmentList
}

func (p Problem) Tags() []string {
	return []string{}
}

func (p Problem) Compile(s language.Sandbox, l language.Language, src io.Reader, errw io.Writer) (io.Reader, error) {
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

func (p Problem) Run(s language.Sandbox, lang language.Language, bin io.Reader, testNotifier chan string, statusNotifier chan problems.Status) (problems.Status, error) {
	defer func() {
		close(testNotifier)
		close(statusNotifier)
	}()

	ans := problems.Status{}

	ans.Compiled = true

	ans.Feedback = make([]problems.Testset, 1)
	ans.Feedback[0] = problems.Testset{"main", problems.SCORING_SUM, make([]problems.Testcase, 0)}

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

	for tc := 1; tc <= p.TestCount; tc++ {
		inputName := fmt.Sprintf(filepath.Join(p.Path, p.InputPattern), tc)
		outputName := fmt.Sprintf(p.OutputPattern, tc)

		ans.Feedback[0].Testcases = append(ans.Feedback[0].Testcases, problems.Testcase{Testset: "main", VerdictName: problems.VERDICT_RE, Score: 0.0, MaxScore: 0.0})
		currentCase := &ans.Feedback[0].Testcases[len(ans.Feedback[0].Testcases)-1]

		for _, file := range zip.File {
			if file.Name == outputName {
				stdout, stderr := &bytes.Buffer{}, &bytes.Buffer{}

				cmd := exec.Command(filepath.Join(p.Path, "check"), inputName, "/dev/stdin")

				cmd.Stdin, err = file.Open()
				if err != nil {
					currentCase.VerdictName = problems.VERDICT_XX
					currentCase.CheckerOutput = err.Error()
					break
				}

				cmd.Stdout = stdout
				cmd.Stderr = stderr

				err = cmd.Run()


				currentCase.CheckerOutput = stderr.String()
				fmt.Sscanf(stdout.String(), "%f/%f", &currentCase.Score, &currentCase.MaxScore)

				if err == nil {
					currentCase.VerdictName = problems.VERDICT_AC
				}else {
					currentCase.VerdictName = problems.VERDICT_WA
				}

				break
			}
		}
	}

	return ans, nil
}

func parser(path string) (problems.Problem, error){
	problemJSON, err := os.Open(filepath.Join(path, "problem.json"))
	if err != nil {
		return nil, err
	}

	defer problemJSON.Close()

	p := Problem{}

	if err := json.NewDecoder(problemJSON).Decode(&p); err!= nil {
		return nil, err
	}

	p.Path = path

	p.AttachmentList = make([]problems.Attachment, len(p.AttachmentInfo))
	for ind, val := range p.AttachmentInfo {
		contents, err := ioutil.ReadFile(filepath.Join(path, val.Path))

		if err != nil {
			return nil, err
		}

		p.AttachmentList[ind] = problems.Attachment{val.Name,contents}
	}

	p.GeneratedStatementList = make([]problems.Content, len(p.StatementList))
	for ind, val := range p.StatementList {
		contents, err := ioutil.ReadFile(filepath.Join(path, val.Path))

		if err != nil {
			return nil, err
		}

		p.GeneratedStatementList[ind] = problems.Content{Locale: val.Language, Contents: contents, Type: val.Type}
	}

	return p, nil
}

func init() {
	problems.RegisterType("zipoo", parser)
}