package polygon

import (
	"bytes"
	"encoding/json"
	"github.com/spf13/afero"
	"html"
	"os/exec"
	"path/filepath"
	"strings"
	"text/template"
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

func convertPandoc(str *string) (err error) {
	buf := &bytes.Buffer{}

	cmd := exec.Command("pandoc", "--mathjax", "-f", "latex", "-t", "html")
	cmd.Stdin = strings.NewReader(*str)
	cmd.Stdout = buf

	err = cmd.Run()
	if err == nil {
		*str = html.UnescapeString(buf.String())
	}

	return
}

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
	Interaction string
	Output      string
	Scoring     string
	SampleTests []SampleTest
	Notes       string
}

func ParseJSONStatement(fs afero.Fs, path string) (*JSONStatement, error) {
	var (
		stmt JSONStatement
		err  error
	)

	propsFile, err := fs.Open(filepath.Join(path, "problem-properties.json"))
	if err != nil {
		return nil, err
	}

	defer func() {
		err = propsFile.Close()
	}()

	dec := json.NewDecoder(propsFile)
	if err = dec.Decode(&stmt); err != nil {
		return nil, err
	}

	for _, elem := range []*string{&stmt.Legend, &stmt.Input, &stmt.Output, &stmt.Notes, &stmt.Scoring, &stmt.Interaction} {
		if err == nil {
			err = convertPandoc(elem)
		}
	}

	return &stmt, err
}

func (j JSONStatement) Html() ([]byte, error) {
	buf := bytes.Buffer{}
	if err := htmlTmpl.Execute(&buf, j); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func init() {
	if tmpl, err := template.New("polygonHtmlTemplate").Funcs(template.FuncMap{"div": func(a, b int) int { return a / b }}).Parse(htmlTemplate); err != nil {
		panic(err)
	} else {
		htmlTmpl = tmpl
	}
}
