package polygon

import (
	"bytes"
	"encoding/xml"
	"errors"
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/mraron/njudge/pkg/language/langs/cpp"
	"github.com/mraron/njudge/pkg/language/sandbox"

	"github.com/mraron/njudge/pkg/language"
	"github.com/mraron/njudge/pkg/problems"
	"github.com/spf13/afero"
	"go.uber.org/multierr"
)

func compileIfNotCompiled(fs afero.Fs, wd, src, dst string) error {
	if src == "" {
		return nil
	}

	if st, err := fs.Stat(dst); os.IsNotExist(err) || st.Size() == 0 {
		if binary, err := fs.Create(dst); err == nil {
			if file, err := fs.Open(src); err == nil {
				var buf bytes.Buffer
				s := sandbox.NewDummy()
				if err := s.Init(log.New(ioutil.Discard, "", 0)); err != nil {
					return multierr.Combine(err, binary.Close(), file.Close())
				}

				//hacky solution
				var headers []language.File
				conts, err := afero.ReadFile(fs, src)
				if err != nil {
					return multierr.Combine(err, file.Close(), binary.Close())
				}
				if bytes.Contains(conts, []byte("testlib.h")) {
					f, err := fs.Open(filepath.Join(wd, "testlib.h"))
					if err != nil {
						return multierr.Combine(err, f.Close(), file.Close(), binary.Close())
					}

					headers = append(headers, language.File{
						Name:   "testlib.h",
						Source: f,
					})
				}

				if err := cpp.Std14.Compile(s, language.File{
					Name:   filepath.Base(src),
					Source: file,
				}, binary, &buf, headers); err != nil {
					return multierr.Combine(err, binary.Close(), file.Close(), fmt.Errorf("compile error: %v", buf.String()))
				}

				if err := fs.Chmod(dst, os.ModePerm); err != nil {
					return multierr.Combine(err, binary.Close(), file.Close())
				}

				return multierr.Combine(binary.Close(), file.Close())
			} else {
				return multierr.Combine(err, binary.Close())
			}
		} else {
			return err
		}
	} else {
		return err
	}
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
	fs              afero.Fs
	compileBinaries bool
}

func newConfig() *config {
	return &config{fs: afero.NewOsFs(), compileBinaries: true}
}

func ParserAndIdentifier(opts ...Option) (problems.ConfigParser, problems.ConfigIdentifier) {
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

		list, err := afero.ReadDir(cfg.fs, filepath.Join(path, "statements"))
		if err == nil {
			for _, dir := range list {
				if !dir.IsDir() || strings.HasPrefix(dir.Name(), ".") {
					continue
				}

				jsonStmt, err := ParseJSONStatement(cfg.fs, filepath.Join(path, "statements", dir.Name()))
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

		for _, stmt := range p.StatementList {
			statementPath := filepath.Join(path, stmt.Path)
			cont, err := afero.ReadFile(cfg.fs, statementPath)
			if err != nil {
				return nil, err
			}

			p.GeneratedStatementList = append(p.GeneratedStatementList, problems.BytesData{Loc: stmt.Language, Val: cont, Typ: stmt.Type})
		}

		if cfg.compileBinaries {
			workingDirectory := p.Path
			if _, err := cfg.fs.Stat(filepath.Join(p.Path, "files")); !errors.Is(err, fs.ErrNotExist) {
				workingDirectory = filepath.Join(p.Path, "files")
			}

			checkerPath := p.Assets.Checker.Source.Path
			if checkerPath == "" {
				checkerPath = "check.cpp"
			}
			if err := compileIfNotCompiled(cfg.fs, workingDirectory, filepath.Join(p.Path, checkerPath), filepath.Join(p.Path, "check")); err != nil {
				return nil, err
			}

			if p.Assets.Interactor.Source.Path != "" {
				if err := compileIfNotCompiled(cfg.fs, workingDirectory, filepath.Join(p.Path, p.Assets.Interactor.Source.Path), filepath.Join(p.Path, "files/interactor")); err != nil {
					return nil, err
				}
			}
		}

		for _, val := range p.Assets.Attachments {
			attachmentLocation := filepath.Join(path, val.Location)
			contents, err := afero.ReadFile(cfg.fs, attachmentLocation)
			if err != nil {
				return nil, err
			}

			p.AttachmentsList = append(p.AttachmentsList, problems.BytesData{Nam: val.Name, Val: contents})
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
