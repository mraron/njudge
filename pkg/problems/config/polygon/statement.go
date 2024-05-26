package polygon

import (
	"bytes"
	"embed"
	"encoding/json"
	"github.com/spf13/afero"
	"html"
	"io"
	"os/exec"
	"path/filepath"
	"strings"
	"text/template"
)

var (
	//go:embed statements/*
	statementTemplates embed.FS

	defaultLocale  = "english"
	templates      map[string]*template.Template
	statementFuncs = template.FuncMap{
		"div": func(a, b int) int {
			return a / b
		},
		"needSection": func(s *string) bool {

			return s != nil && strings.TrimSpace(*s) != ""
		},
	}
)

func convertPandoc(str *string) (err error) {
	if str == nil {
		return nil
	}
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
	Locale string `json:"-"`

	Name        string
	TimeLimit   int
	MemoryLimit int
	InputFile   string
	OutputFile  string
	Legend      *string `json:"legend"`
	Input       *string `json:"input"`
	Interaction *string `json:"interaction"`
	Output      *string `json:"output"`
	Scoring     *string `json:"scoring"`
	SampleTests []SampleTest
	Notes       *string `json:"notes"`
}

func ParseJSONStatement(fs afero.Fs, path string, locale string) (*JSONStatement, error) {
	problemProps := filepath.Join(path, "problem-properties.json")
	if exists, err := afero.Exists(fs, problemProps); !exists || err != nil {
		return nil, err
	}

	propsFile, err := fs.Open(problemProps)
	if err != nil {
		return nil, err
	}
	return NewJSONStatement(propsFile, locale)
}

func NewJSONStatement(r io.ReadCloser, locale string) (*JSONStatement, error) {
	var (
		stmt JSONStatement
		err  error
	)

	dec := json.NewDecoder(r)
	defer r.Close()
	if err = dec.Decode(&stmt); err != nil {
		return nil, err
	}

	if _, ok := templates[locale]; !ok {
		locale = defaultLocale
	}

	stmt.Locale = locale
	for _, elem := range []*string{stmt.Legend, stmt.Input, stmt.Output, stmt.Notes, stmt.Scoring, stmt.Interaction} {
		if err == nil {
			err = convertPandoc(elem)
		}
	}

	return &stmt, err
}

func (j JSONStatement) Html() ([]byte, error) {
	buf := bytes.Buffer{}
	if err := templates[j.Locale].Execute(&buf, j); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func init() {
	elems, err := statementTemplates.ReadDir("statements")
	if err != nil {
		panic(err)
	}
	templates = make(map[string]*template.Template)
	for _, elem := range elems {

		if !elem.IsDir() && filepath.Ext(elem.Name()) == ".gohtml" {
			name := strings.TrimSuffix(filepath.Base(elem.Name()), ".gohtml")
			templates[name] = template.Must(template.New(elem.Name()).Funcs(statementFuncs).ParseFS(statementTemplates, filepath.Join("statements", elem.Name())))
		}
	}
}
