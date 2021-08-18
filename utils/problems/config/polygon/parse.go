package polygon

import (
	"encoding/xml"
	"errors"
	"github.com/mraron/njudge/utils/language"
	"github.com/mraron/njudge/utils/problems"
	"github.com/spf13/afero"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func compileIfNotCompiled(fs afero.Fs, path string, src, dst string) error {
	if src == "" {
		return nil
	}

	if st, err := fs.Stat(filepath.Join(path, dst)); os.IsNotExist(err) || st.Size() == 0 {
		if binary, err := fs.Create(filepath.Join(path, dst)); err == nil {
			defer binary.Close()

			if lang := language.Get("cpp14"); lang != nil {
				if file, err := fs.Open(filepath.Join(path, src)); err == nil {
					defer file.Close()
					if err := lang.InsecureCompile(filepath.Join(path, "files"), file, binary, os.Stderr); err != nil {
						return err
					}

					if err := fs.Chmod(filepath.Join(path, dst), os.ModePerm); err != nil {
						return err
					}
				} else {
					return err
				}
			} else {
				return errors.New("can't compile file, no cpp14 compiler")
			}
		} else {
			return err
		}
	}

	return nil
}

type Option func(*config)

func UseFS(fs afero.Fs) Option {
	return func(c *config) {
		c.fs = fs
	}
}

func CompileBinaries(compile bool) Option {
	return func(c *config) {
		c.compileBinaries = compile
	}
}

type config struct {
	fs afero.Fs
	compileBinaries bool
}

func newConfig() *config {
	return &config{fs: afero.NewOsFs(), compileBinaries: true}
}

func ParserAndIdentifier(opts... Option) (problems.ConfigParser, problems.ConfigIdentifier) {
	cfg := newConfig()
	for _, opt := range opts {
		opt(cfg)
	}

	parser := func(path string) (problems.Problem, error) {
		problemXML := filepath.Join(path, "problem.xml")

		f, err := cfg.fs.Open(problemXML)
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

		list, err := ioutil.ReadDir(filepath.Join(path, "statements"))
		if err == nil {
			for _, dir := range list {
				if !dir.IsDir() || strings.HasPrefix(dir.Name(), ".") {
					continue
				}

				jsonStmt, err := ParseJSONStatement(cfg.fs, filepath.Join(path, "statements", dir.Name()))
				if err != nil {
					return nil, err
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

				p.GeneratedStatementList = append(p.GeneratedStatementList, problems.Content{Locale: dir.Name(), Contents: contents, Type: "text/html"})
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

		if cfg.compileBinaries {
			if err := compileIfNotCompiled(cfg.fs, p.Path, p.Assets.Checker.Source.Path, "check"); err != nil {
				return nil, err
			}
			if err := compileIfNotCompiled(cfg.fs, p.Path, p.Assets.Interactor.Source.Path, "files/interactor"); err != nil {
				return nil, err
			}
		}

		for _, val := range p.Assets.Attachments {
			attachmentLocation := filepath.Join(path, val.Location)
			contents, err := ioutil.ReadFile(attachmentLocation)
			if err != nil {
				return nil, err
			}

			p.AttachmentsList = append(p.AttachmentsList, problems.Attachment{Name:val.Name, Contents: contents})
		}

		return p, nil
	}

	identifier := func(path string) bool {
		_, err := cfg.fs.Stat(filepath.Join(path, "problem.xml"))
		return !os.IsNotExist(err)
	}

	return parser, identifier
}

func init() {
	parser, identifier := ParserAndIdentifier(CompileBinaries(true))
	if err := problems.RegisterConfigType("polygon", parser, identifier); err != nil {
		panic(err)
	}
}
