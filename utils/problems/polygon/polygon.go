package polygon

import (
	"github.com/mraron/njudge/utils/language"
	_ "github.com/mraron/njudge/utils/language/cpp11"

	"github.com/mraron/njudge/utils/problems"

	"os"
	"os/exec"

	"path/filepath"

	"encoding/xml"

	"io"
	"io/ioutil"

	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"text/template"
	"time"
)

const htmlTemplate = `<link href="problem-statement.css" rel="stylesheet" type="text/css"><div class="problem-statement">
<div class="header">
	<div class="title">{{.Name}}</div>
	<div class="time-limit"><div class="property-title">tesztenkénti időlimit</div> {{.TimeLimit}} ms</DIV>
	<div class="memory-limit"><div class="property-title">tesztenkénti memórialimit</div> {{div .MemoryLimit 1048576}} MiB</div>
	<div class="input-file"><div class="property-title">inputfájl</div> {{if .InputFile}} {{.InputFile}} {{else}} stdin {{end}}</div>
	<div class="output-file"><div class="property-title">outputfájl</div> {{if .OutputFile}} {{.OutputFile}} {{else}} stdout {{end}} </div>
</div><p></p><p></p>
{{if .Legend}}<div class="legend">{{.Legend}}</div><p></p><p></p>{{end}}
{{if .Input}}<div class="input-specification"><div class="section-title">Bemenet</div> {{.Input}}</div><p></p><p></p>{{end}}
{{if .Output}}<div class="input-specification"><div class="section-title">Kimenet</div> {{.Output}}</div><p></p><p></p>{{end}}
{{if .SampleTests}}
<div class="sample-tests">
	<div class="section-title">Példák</div>
	{{range $i := .SampleTests}}
		<div class="sample-test">
			<div class="input"><div class="title">Bemenet</div><pre class="content">{{$i.Input}}</pre></div>
			<div class="output"><div class="title">Kimenet</div><pre class="content">{{$i.Output}}</pre></div>
		</div>
		<p></p><p></p>
	{{end}}
</div>
{{end}}

{{if .Notes}}<div class="section-title">Megjegyzések</div> {{.Notes}}<p></p><p></p>{{end}}
</div>`

var htmlTmpl *template.Template

type SampleTest struct {
	Input  string
	Output string
}

type JSONStatement struct {
	Name        string
	TimeLimit   int
	MemoryLimit int
	InputFile   string
	OutputFile  string
	Legend      string
	Input       string
	Output      string
	SampleTests []SampleTest
	Notes       string
}

type Name struct {
	Language string `xml:"language,attr"`
	Value    string `xml:"value,attr"`
}

type Statement struct {
	Language string `xml:"language,attr"`
	Path     string `xml:"path,attr"`
	Type     string `xml:"type,attr"`
}

type Test struct {
	Method string `xml:"method,attr"`
	Cmd    string `xml:"cmd,attr"`
	Sample bool   `xml:"sample,attr"`
	Score  int    `xml:"score,attr"`
}

type Judging struct {
	CpuName    string `xml:"cpu-name,attr"`
	CpuSpeed   int    `xml:"cpu-speed,attr"`
	InputFile  string `xml:"input-file,attr"`
	OutputFile string `xml:"output-file,attr"`

	TestSet []struct {
		Name              string `xml:"name,attr"`
		TimeLimit         int    `xml:"time-limit"`
		MemoryLimit       int    `xml:"memory-limit"`
		TestCount         int    `xml:"test-count"`
		InputPathPattern  string `xml:"input-path-pattern"`
		AnswerPathPattern string `xml:"answer-path-pattern"`
		Scoring           string `xml:"scoring,attr"`
		Tests             []Test `xml:"tests>test"`
	} `xml:"testset"`
}

type Attachment struct {
	Name     string `xml:"name,attr"`
	Location string `xml:"location,attr"`
}

type Assets struct {
	Attachments []Attachment `xml:"attachments>attachment"`
}

type Problem struct {
	Path                   string
	JSONStatementList      []JSONStatement
	AttachmentsList        []problems.Attachment
	GeneratedStatementList []problems.Content

	FeedbackType  string      `xml:"feedback,attr"`
	Revision      int         `xml:"revision,attr"`
	ShortName     string      `xml:"short-name,attr"`
	Url           string      `xml:"url,attr"`
	Names         []Name      `xml:"names>name"`
	StatementList []Statement `xml:"statements>statement"`
	Judging       Judging     `xml:"judging"`
	Assets        Assets      `xml:"assets"`
	TagsList      []struct {
		Value string `xml:"value,attr"`
	} `xml:"tags>tag"`
}

func (p Problem) Name() string {
	return p.ShortName
}

func (p Problem) Titles() []problems.Content {
	ans := make([]problems.Content, len(p.Names))
	for i := 0; i < len(p.Names); i++ {
		ans[i] = problems.Content{p.Names[i].Language, []byte(p.Names[i].Value), "text"}
	}

	return ans
}

func (p Problem) Statements() []problems.Content {
	return p.GeneratedStatementList
}

func (p Problem) HTMLStatements() []problems.Content {
	lst := make([]problems.Content, 0)
	for _, val := range p.GeneratedStatementList {
		if val.Type == "text/html" {
			lst = append(lst, val)
		}
	}

	return lst
}

func (p Problem) PDFStatements() []problems.Content {
	lst := make([]problems.Content, 0)
	for _, val := range p.GeneratedStatementList {
		if val.Type == "application/pdf" {
			lst = append(lst, val)
		}
	}

	return lst
}

func (p Problem) MemoryLimit() int {
	return p.Judging.TestSet[0].MemoryLimit
}

func (p Problem) TimeLimit() int {
	return p.Judging.TestSet[0].TimeLimit
}

func (p Problem) InputOutputFiles() (string, string) {
	return p.Judging.InputFile, p.Judging.OutputFile
}

func (p Problem) Interactive() bool {
	return false
}

func (p Problem) Attachments() []problems.Attachment {
	return p.AttachmentsList
}

func (p Problem) Tags() (lst []string) {
	lst = make([]string, len(p.TagsList))
	for ind, val := range p.TagsList {
		lst[ind] = val.Value
	}

	return
}

func (p Problem) Compile(s language.Sandbox, lang language.Language, src io.Reader, e io.Writer) (io.Reader, error) {
	buf := &bytes.Buffer{}

	err := lang.Compile(s, src, buf, e)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

func (Problem) Languages() []language.Language {
	lst1 := language.List()

	lst2 := make([]language.Language, 0, len(lst1))
	for _, val := range lst1 {
		if val.Id() != "zip" {
			lst2 = append(lst2, val)
		}
	}

	return lst2
}

func truncate(s string) string {
	if len(s) < 256 {
		return s
	}

	return s[:255] + "..."
}

func (p Problem) Run(s language.Sandbox, lang language.Language, bin io.Reader, testNotifier chan string, statusNotifier chan problems.Status) (problems.Status, error) {
	var (
		ans            problems.Status
		err            error
		binaryContents []byte
	)

	defer func() {
		close(testNotifier)
		close(statusNotifier)
	}()

	ans.Compiled = true
	ans.Feedback = make([]problems.Testset, 0)
	ans.FeedbackType = problems.FeedbackFromString(p.FeedbackType)

	if binaryContents, err = ioutil.ReadAll(bin); err != nil {
		return ans, err
	}

	for tsid, ts := range p.Judging.TestSet {
		ans.Feedback = append(ans.Feedback, problems.Testset{Name: ts.Name})
		testset := &ans.Feedback[len(ans.Feedback)-1]
		testset.Scoring = problems.ScoringFromString(ts.Scoring)
		for tc := 1; tc <= ts.TestCount; tc++ {
			testNotifier <- strconv.Itoa(tc)
			statusNotifier <- ans

			testLocation, answerLocation := fmt.Sprintf(filepath.Join(p.Path, ts.InputPathPattern), tc), fmt.Sprintf(filepath.Join(p.Path, ts.AnswerPathPattern), tc)

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

			res, err := lang.Run(s, bytes.NewReader(binaryContents), testcase, stdout, time.Duration(ts.TimeLimit)*time.Millisecond, ts.MemoryLimit)
			if err != nil {
				return ans, err
			}

			if res.Verdict == language.VERDICT_OK {
				checkerOutput := &bytes.Buffer{}
				programOutput := stdout.String()

				expectedOutput := string(answerContents)

				cmd := exec.Command(filepath.Join(p.Path, "check"), testLocation, "/dev/stdin", answerLocation)
				cmd.Stdin = stdout
				cmd.Stdout = checkerOutput
				cmd.Stderr = checkerOutput

				err = cmd.Run()

				/*	if err != nil {
					ans.Feedback = append(ans.Feedback, problems.Testcase{VerdictName: problems.VERDICT_XX, MemoryUsed: res.Memory, TimeSpent: res.Time})
					return ans, err
				}*/

				testset.Testcases = append(testset.Testcases, problems.Testcase{MaxScore: float64(p.Judging.TestSet[tsid].Tests[tc-1].Score), Testset: ts.Name, MemoryUsed: res.Memory, TimeSpent: res.Time, Output: truncate(programOutput), ExpectedOutput: truncate(expectedOutput), CheckerOutput: truncate(checkerOutput.String())})

				if err == nil && cmd.ProcessState.Success() {
					testset.Testcases[len(testset.Testcases)-1].VerdictName = problems.VERDICT_AC
					testset.Testcases[len(testset.Testcases)-1].Score = testset.Testcases[len(testset.Testcases)-1].MaxScore
				} else {
					testset.Testcases[len(testset.Testcases)-1].VerdictName = problems.VERDICT_WA
					testset.Testcases[len(testset.Testcases)-1].Score = 0
					if problems.FeedbackFromString(p.FeedbackType) != problems.FEEDBACK_IOI {
						return ans, nil
					}
				}
			} else {
				curr := problems.Testcase{}
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

				curr.MemoryUsed = res.Memory
				curr.TimeSpent = res.Time
				curr.Score = 0
				curr.MaxScore = float64(p.Judging.TestSet[tsid].Tests[tc-1].Score)
				curr.Output = truncate(stdout.String()) //now it's stderr
				curr.ExpectedOutput = truncate(string(answerContents))

				testset.Testcases = append(testset.Testcases, curr)

				if problems.FeedbackFromString(p.FeedbackType) != problems.FEEDBACK_IOI {
					return ans, nil
				}
			}

		}
	}

	return ans, nil
}

func parser(path string) (problems.Problem, error) {
	problemXML := filepath.Join(path, "problem.xml")

	f, err := os.Open(problemXML)
	if err != nil {
		return nil, err
	}

	p := Problem{}

	dec := xml.NewDecoder(f)
	if err := dec.Decode(&p); err != nil {
		return nil, err
	}

	p.Path = path

	list, err := ioutil.ReadDir(filepath.Join(path, "statements"))
	if err != nil {
		return nil, err
	}

	for _, dir := range list {
		if !dir.IsDir() || strings.HasPrefix(dir.Name(), ".") {
			continue
		}

		jsonstmt := JSONStatement{}

		propertiesFile, err := os.Open(filepath.Join(path, "statements", dir.Name(), "problem-properties.json"))
		if err != nil {
			return nil, err
		}

		dec := json.NewDecoder(propertiesFile)
		if err := dec.Decode(&jsonstmt); err != nil {
			return nil, err
		}

		replace := func(str *string) {
			*str = strings.Replace(*str, "\n\n", "<p></p><p></p>", -1)
		}

		replace(&jsonstmt.Legend)
		replace(&jsonstmt.Input)
		replace(&jsonstmt.Output)
		replace(&jsonstmt.Notes)

		jsonstmt.InputFile, jsonstmt.OutputFile = p.InputOutputFiles()
		jsonstmt.TimeLimit = p.TimeLimit()
		jsonstmt.MemoryLimit = p.MemoryLimit()

		p.JSONStatementList = append(p.JSONStatementList, jsonstmt)

		buf := bytes.Buffer{}
		htmlTmpl.Execute(&buf, jsonstmt)

		p.GeneratedStatementList = append(p.GeneratedStatementList, problems.Content{Locale: dir.Name(), Contents: buf.Bytes(), Type: "text/html"})
	}

	for _, stmt := range p.StatementList {
		statementPath := filepath.Join(path, stmt.Path)

		cont, err := ioutil.ReadFile(statementPath)
		if err != nil {
			return nil, err
		}

		p.GeneratedStatementList = append(p.GeneratedStatementList, problems.Content{Locale: stmt.Language, Contents: cont, Type: stmt.Type})
	}

	if _, err := os.Stat(filepath.Join(path, "check")); os.IsNotExist(err) {
		if checkerBinary, err := os.Create(filepath.Join(path, "check")); err == nil {
			defer checkerBinary.Close()

			if lang := language.Get("cpp11"); lang != nil {
				if checkerFile, err := os.Open(filepath.Join(path, "files", "check.cpp")); err == nil {
					defer checkerFile.Close()

					if err := lang.InsecureCompile(filepath.Join(path, "files"), checkerFile, checkerBinary, os.Stderr); err != nil {
						return nil, err
					}

					if err := os.Chmod(filepath.Join(path, "check"), 0777); err != nil {
						return nil, err
					}
				} else {
					return nil, err
				}
			} else {
				return nil, errors.New("error while parsing polygon problem can't compile polygon checker because there's no cpp11 compiler")
			}
		} else {
			return nil, err
		}
	}

	for _, val := range p.Assets.Attachments {
		attachmentLocation := filepath.Join(path, val.Location)
		contents, err := ioutil.ReadFile(attachmentLocation)
		if err != nil {
			return nil, err
		}

		p.AttachmentsList = append(p.AttachmentsList, problems.Attachment{val.Name, contents})
	}

	return p, nil
}

func identifier(path string) bool {
	_, err := os.Stat(filepath.Join(path, "problem.xml"))
	return !os.IsNotExist(err)
}

func init() {
	problems.RegisterType("polygon", parser, identifier)

	if tmpl, err := template.New("polygonHtmlTemplate").Funcs(template.FuncMap{"div": func(a, b int) int { return a / b }}).Parse(htmlTemplate); err != nil {
		panic(err)
	} else {
		htmlTmpl = tmpl
	}
}
