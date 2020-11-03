package polygon

import (
	"github.com/mraron/njudge/utils/language"
	_ "github.com/mraron/njudge/utils/language/cpp11"
	_ "github.com/mraron/njudge/utils/language/cpp14"
	"github.com/mraron/njudge/utils/problems"
	"os/exec"
	"syscall"

	"os"
	"path/filepath"

	"encoding/xml"

	"io/ioutil"

	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"text/template"
	"time"
)

//@TODO add njudge prefix to non standard problem xml additions
//@TODO validate group/test points and create a logs directory maybe for problem parsin logs

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
{{if .Scoring}}<div class="input-specification"><div class="section-title">Pontozás</div> {{.Scoring}}</div><p></p><p></p>{{end}}
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

// @TODO: add scoring, explore what is it like

type JSONStatement struct {
	Name        string
	TimeLimit   int
	MemoryLimit int
	InputFile   string
	OutputFile  string
	Legend      string
	Input       string
	Output      string
	Scoring     string
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
	Method string  `xml:"method,attr"`
	Cmd    string  `xml:"cmd,attr"`
	Sample bool    `xml:"sample,attr"`
	Points float64 `xml:"points,attr"`
	Group  string  `xml:"group,attr"`

	Input  string
	Answer string
	Index  int
}

//@TODO FeedbackPolicy
type Group struct {
	Name         string  `xml:"name,attr"`
	PointsPolicy string  `xml:"points-policy,attr"`
	Points       float64 `xml:"points,attr"`
}

type Judging struct {
	CpuName    string `xml:"cpu-name,attr"`
	CpuSpeed   int    `xml:"cpu-speed,attr"`
	InputFile  string `xml:"input-file,attr"`
	OutputFile string `xml:"output-file,attr"`

	TestSet []struct {
		Name              string  `xml:"name,attr"`
		TimeLimit         int     `xml:"time-limit"`
		MemoryLimit       int     `xml:"memory-limit"`
		TestCount         int     `xml:"test-count"`
		InputPathPattern  string  `xml:"input-path-pattern"`
		AnswerPathPattern string  `xml:"answer-path-pattern"`
		Tests             []Test  `xml:"tests>test"`
		Groups            []Group `xml:"groups>group"`
	} `xml:"testset"`
}

type Attachment struct {
	Name     string `xml:"name,attr"`
	Location string `xml:"location,attr"`
}

type Stub struct {
	Name     string `xml:"name,attr"`
	Path     string `xml:"path,attr"`
	Language string `xml:"language,attr"`
}

type Checker struct {
	Type   string `xml:"type,attr"`
	Source struct {
		Path string `xml:"path,attr"`
	} `xml:"source"`
}

type Assets struct {
	Attachments []Attachment `xml:"attachments>attachment"`
	Stubs       []Stub       `xml:"stubs>stub"`
	Checker     Checker      `xml:"checker"`
}

type Problem struct {
	Path                   string
	JSONStatementList      []JSONStatement
	AttachmentsList        []problems.Attachment
	GeneratedStatementList []problems.Content

	TaskType      string      `xml:"njudge-task-type,attr"`
	FeedbackType  string      `xml:"njudge-feedback-type,attr"`
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

func (p Problem) StatusSkeleton() problems.Status {
	ans := problems.Status{false, "status skeleton", problems.FeedbackFromString(p.FeedbackType), make([]problems.Testset, 0)}

	for _, ts := range p.Judging.TestSet {
		ans.Feedback = append(ans.Feedback, problems.Testset{Name: ts.Name})
		testset := &ans.Feedback[len(ans.Feedback)-1]

		tcbygroup := make(map[string][]Test)
		for ind, tc := range ts.Tests {
			tc.Input, tc.Answer = fmt.Sprintf(filepath.Join(p.Path, ts.InputPathPattern), ind+1), fmt.Sprintf(filepath.Join(p.Path, ts.AnswerPathPattern), ind+1)
			tc.Index = ind + 1

			if len(tcbygroup[tc.Group]) == 0 {
				tcbygroup[tc.Group] = make([]Test, 0)
			}

			tcbygroup[tc.Group] = append(tcbygroup[tc.Group], tc)
		}

		if len(ts.Groups) == 0 {
			fmt.Println(len(ts.Groups), "wuwtuaofjsdléfklék")
			ts.Groups = append(ts.Groups, Group{"", "sum", -1.0})
		}

		idx := 1
		for _, grp := range ts.Groups {
			testset.Groups = append(testset.Groups, problems.Group{})
			group := &testset.Groups[len(testset.Groups)-1]

			group.Name = grp.Name
			if grp.PointsPolicy == "complete-group" {
				group.Scoring = problems.SCORING_GROUP
			}else {
				group.Scoring = problems.SCORING_SUM
			}

			for _, tc := range tcbygroup[grp.Name] {
				testcase := problems.Testcase{idx, tc.Input, "", tc.Answer, ts.Name, group.Name, problems.VERDICT_DR, float64(0.0), float64(tc.Points), "-", "-", "-", 0 * time.Millisecond, 0, time.Duration(p.TimeLimit()) * time.Millisecond, p.MemoryLimit()}
				group.Testcases = append(group.Testcases, testcase)
				testset.Testcases = append(testset.Testcases, testcase)

				idx++
			}
		}
	}

	return ans
}

func (p Problem) Check(tc *problems.Testcase) error {
	output := &bytes.Buffer{}

	cmd := exec.Command(filepath.Join(p.Path, "check"), tc.InputPath, tc.OutputPath, tc.AnswerPath)
	cmd.Stdout = output
	cmd.Stderr = output

	err := cmd.Run()

	tc.CheckerOutput = problems.Truncate(output.String())
	if err != nil {
		if exit_err, ok := err.(*exec.ExitError); ok {
			if status, ok := exit_err.Sys().(syscall.WaitStatus); ok {
				if status.ExitStatus() == 1 {
					tc.VerdictName = problems.VERDICT_WA
				} else if status.ExitStatus() == 2 {
					tc.VerdictName = problems.VERDICT_PE
				} else if status.ExitStatus() == 7 {
					tc.VerdictName = problems.VERDICT_PC

					rel := 0
					fmt.Sscanf(output.String(), "points %d", &rel)
					rel -= 16

					tc.Score = float64(rel) / (200.0 * tc.MaxScore)
				} else { //3 -> fail
					tc.VerdictName = problems.VERDICT_XX
				}
			}
		} else {
			tc.VerdictName = problems.VERDICT_XX
			return err
		}
	} else {
		tc.Score = tc.MaxScore
		tc.VerdictName = problems.VERDICT_AC
	}

	return nil
}

func (p Problem) Files() []problems.File {
	res := make([]problems.File, 0)
	for _, stub := range p.Assets.Stubs {
		res = append(res, problems.File{stub.Name, "stub_" + stub.Language, filepath.Join(p.Path, stub.Path)})
	}

	return res
}

func (p Problem) TaskTypeName() string {
	if p.TaskType == "" {
		return "batch"
	}

	return p.TaskType
}

//@TODO actually respect problem.xml with statements and checkers

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
	if err == nil {
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

			convert_pandoc := func(str *string) {
				if err != nil {
					return
				}

				buf := &bytes.Buffer{}

				cmd := exec.Command("pandoc", "--mathjax", "-f", "latex", "-t", "html")
				cmd.Stdin = strings.NewReader(*str)
				cmd.Stdout = buf

				err = cmd.Run()
				if err == nil {
					*str = buf.String()
				}
			}

			convert_pandoc(&jsonstmt.Legend)
			convert_pandoc(&jsonstmt.Input)
			convert_pandoc(&jsonstmt.Output)
			convert_pandoc(&jsonstmt.Notes)
			convert_pandoc(&jsonstmt.Scoring)

			if err != nil {
				return nil, err
			}

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
	}

	if _, err := os.Stat(filepath.Join(path, "check")); os.IsNotExist(err) {
		if checkerBinary, err := os.Create(filepath.Join(path, "check")); err == nil {
			defer checkerBinary.Close()

			if lang := language.Get("cpp14"); lang != nil {
				if checkerFile, err := os.Open(filepath.Join(path, p.Assets.Checker.Source.Path)); err == nil {
					defer checkerFile.Close()
					fmt.Println(filepath.Join(path, p.Assets.Checker.Source.Path), "!!!")
					if err := lang.InsecureCompile(filepath.Join(path, "files"), checkerFile, checkerBinary, os.Stderr); err != nil {
						return nil, err
					}

					if err := os.Chmod(filepath.Join(path, "check"), os.ModePerm); err != nil {
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
	problems.RegisterConfigType("polygon", parser, identifier)

	if tmpl, err := template.New("polygonHtmlTemplate").Funcs(template.FuncMap{"div": func(a, b int) int { return a / b }}).Parse(htmlTemplate); err != nil {
		panic(err)
	} else {
		htmlTmpl = tmpl
	}
}
