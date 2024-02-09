package polygon

import (
	"encoding/xml"
	"errors"
	context2 "golang.org/x/net/context"
	"os"
	"path/filepath"
	"strings"

	"github.com/mraron/njudge/pkg/language/langs/cpp"
	"github.com/mraron/njudge/pkg/language/sandbox"
	"github.com/mraron/njudge/pkg/problems"
	"github.com/spf13/afero"
)

type Option func(*config)

func CompileBinaries(compile bool) Option {
	return func(c *config) {
		c.compileBinaries = compile
	}
}

func DontGenHTML() Option {
	return func(c *config) {
		c.disableHTMLGen = true
	}
}

type config struct {
	compileBinaries bool
	disableHTMLGen  bool
}

func newConfig() *config {
	return &config{compileBinaries: true}
}

func ParserAndIdentifier(opts ...Option) (problems.ConfigParser, problems.ConfigIdentifier) {
	cfg := newConfig()
	for _, opt := range opts {
		opt(cfg)
	}

	parser := func(fs afero.Fs, path string) (problems.Problem, error) {
		problemXML := filepath.Join(path, "problem.xml")

		f, err := fs.Open(problemXML)
		if err != nil {
			return nil, err
		}
		defer f.Close()

		p := Problem{config: cfg}

		dec := xml.NewDecoder(f)
		if err := dec.Decode(&p); err != nil {
			return nil, err
		}

		p.Path = path

		list, err := afero.ReadDir(fs, filepath.Join(path, "statements"))
		if err == nil {
			for _, dir := range list {
				if !dir.IsDir() || strings.HasPrefix(dir.Name(), ".") {
					continue
				}

				if !cfg.disableHTMLGen {
					jsonStmt, err := ParseJSONStatement(fs, filepath.Join(path, "statements", dir.Name()))
					if err != nil {
						return nil, err
					}

					if jsonStmt == nil {
						continue
					}

					// problem-properties.json might be outdated. problem.xml should take priority
					jsonStmt.InputFile, jsonStmt.OutputFile = p.InputOutputFiles()
					jsonStmt.TimeLimit = p.TimeLimit()
					jsonStmt.MemoryLimit = p.MemoryLimit()

					p.JSONStatementList = append(p.JSONStatementList, *jsonStmt)

					contents, err := jsonStmt.Html()
					if err != nil {
						return nil, err
					}

					p.GeneratedStatementList = append(p.GeneratedStatementList, problems.BytesData{Loc: dir.Name(), Val: contents, Typ: "text/html"})
				}
			}
		}

		for _, stmt := range p.StatementList {
			statementPath := filepath.Join(path, stmt.Path)
			cont, err := afero.ReadFile(fs, statementPath)
			if err != nil {
				return nil, err
			}

			p.GeneratedStatementList = append(p.GeneratedStatementList, problems.BytesData{Loc: stmt.Language, Val: cont, Typ: stmt.Type})
		}

		if cfg.compileBinaries {
			workingDirectory := p.Path
			if _, err := fs.Stat(filepath.Join(p.Path, "files")); !errors.Is(err, os.ErrNotExist) {
				workingDirectory = filepath.Join(p.Path, "files")
			}

			checkerPath := p.Assets.Checker.Source.Path
			if checkerPath == "" {
				checkerPath = "check.cpp"
			}

			s, _ := sandbox.NewDummy()
			if err := cpp.AutoCompile(context2.TODO(), fs, s, workingDirectory, filepath.Join(p.Path, checkerPath), filepath.Join(p.Path, "check")); err != nil {
				return nil, err
			}

			if p.Assets.Interactor.Source.Path != "" {
				s, _ := sandbox.NewDummy()
				if err := cpp.AutoCompile(context2.TODO(), fs, s, workingDirectory, filepath.Join(p.Path, p.Assets.Interactor.Source.Path), filepath.Join(p.Path, "files/interactor")); err != nil {
					return nil, err
				}
			}
		}

		for _, val := range p.Assets.Attachments {
			attachmentLocation := filepath.Join(path, val.Location)
			contents, err := afero.ReadFile(fs, attachmentLocation)
			if err != nil {
				return nil, err
			}

			p.AttachmentsList = append(p.AttachmentsList, problems.BytesData{Nam: val.Name, Val: contents})
		}

		return p, nil
	}

	identifier := func(fs afero.Fs, path string) bool {
		_, err := fs.Stat(filepath.Join(path, "problem.xml"))
		return !os.IsNotExist(err)
	}

	return parser, identifier
}

func init() {
	parser, identifier := ParserAndIdentifier(CompileBinaries(true), DontGenHTML())
	if err := problems.RegisterConfigType("polygon", parser, identifier); err != nil {
		panic(err)
	}
}
