package polygon

import (
	"bytes"
	"context"
	"encoding/xml"
	"errors"
	"github.com/mraron/njudge/pkg/problems/config/polygon/internal"
	"os"
	"path/filepath"
	"strings"

	"github.com/mraron/njudge/pkg/language/langs/cpp"
	"github.com/mraron/njudge/pkg/language/sandbox"
	"github.com/mraron/njudge/pkg/problems"
	"github.com/spf13/afero"
)

type Parser struct {
	CompileBinaries bool
	EnableHTMLGen   bool
}

func (parser Parser) parseStatements(fs afero.Fs, p *Problem) error {
	statementsDir := filepath.Join(p.Path, "statements")
	locales, err := afero.ReadDir(fs, statementsDir)
	if err == nil {
		for _, dir := range locales {
			if !dir.IsDir() || strings.HasPrefix(dir.Name(), ".") {
				continue
			}

			if parser.EnableHTMLGen {
				jsonStmt, err := ParseJSONStatement(fs, filepath.Join(statementsDir, dir.Name()), dir.Name())
				if err != nil {
					return err
				}

				if jsonStmt == nil {
					continue
				}

				// problem-properties.json might be outdated. problem.xml should take priority
				jsonStmt.InputFile, jsonStmt.OutputFile = p.InputOutputFiles()
				jsonStmt.TimeLimit = p.TimeLimit()
				jsonStmt.MemoryLimit = int(p.MemoryLimit())

				p.JSONStatementList = append(p.JSONStatementList, *jsonStmt)

				contents, err := jsonStmt.Html()
				if err != nil {
					return err
				}
				res := bytes.NewBuffer(nil)
				if err = internal.InlineHTML(
					afero.NewBasePathFs(
						fs,
						filepath.Join(statementsDir, dir.Name()),
					), bytes.NewReader(contents), res); err != nil {
					return err
				}
				p.GeneratedStatementList = append(p.GeneratedStatementList,
					problems.BytesData{
						Loc: dir.Name(),
						Val: res.Bytes(),
						Typ: problems.DataTypeHTML,
					},
				)
			}
		}
	}

	for _, stmt := range p.StatementList {
		if stmt.Type == problems.DataTypeHTML && len(p.HTMLStatements().FilterByLocale(stmt.Language)) > 0 {
			continue
		}
		statementPath := filepath.Join(p.Path, stmt.Path)
		cont, err := afero.ReadFile(fs, statementPath)
		if err != nil {
			return err
		}

		p.GeneratedStatementList = append(p.GeneratedStatementList,
			problems.BytesData{
				Loc: stmt.Language,
				Val: cont,
				Typ: stmt.Type,
			},
		)
	}
	return nil
}

func (parser Parser) compileBinaries(fs afero.Fs, p *Problem) error {
	if !parser.CompileBinaries {
		return nil
	}

	var (
		workingDirectory = p.Path
		err              error
	)
	if _, err := fs.Stat(filepath.Join(p.Path, "files")); !errors.Is(err, os.ErrNotExist) {
		workingDirectory = filepath.Join(p.Path, "files")
	}

	checkerPath := p.Assets.Checker.Source.Path
	if checkerPath == "" {
		checkerPath = "check.cpp"
	}

	s, _ := sandbox.NewDummy()
	if err := cpp.AutoCompile(context.TODO(), fs, s, workingDirectory, filepath.Join(p.Path, checkerPath), filepath.Join(p.Path, "check")); err != nil {
		return err
	}

	if p.Assets.Interactor.Source.Path != "" {
		s, _ := sandbox.NewDummy()
		if err := cpp.AutoCompile(context.TODO(), fs, s, workingDirectory, filepath.Join(p.Path, p.Assets.Interactor.Source.Path), filepath.Join(p.Path, "files/interactor")); err != nil {
			return err
		}
		p.Assets.Interactor.binary, err = os.ReadFile(filepath.Join(p.Path, "files/interactor"))
		if err != nil {
			return err
		}
	}

	return nil
}

func (parser Parser) attachmentsList(fs afero.Fs, p *Problem) error {
	for _, val := range p.Assets.Attachments {
		attachmentLocation := filepath.Join(p.Path, val.Location)
		contents, err := afero.ReadFile(fs, attachmentLocation)
		if err != nil {
			return err
		}

		p.AttachmentsList = append(p.AttachmentsList, problems.BytesData{Nam: val.Name, Val: contents})
	}
	return nil
}

func (parser Parser) Parse(fs afero.Fs, path string) (problems.Problem, error) {
	problemXML := filepath.Join(path, "problem.xml")

	f, err := fs.Open(problemXML)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	p := Problem{}

	dec := xml.NewDecoder(f)
	if err := dec.Decode(&p); err != nil {
		return nil, err
	}

	p.Path = path
	if err = parser.parseStatements(fs, &p); err != nil {
		return nil, err
	}
	if err = parser.compileBinaries(fs, &p); err != nil {
		return nil, err
	}
	if err = parser.attachmentsList(fs, &p); err != nil {
		return nil, err
	}

	if err = p.Valid(); err != nil {
		return nil, err
	}
	return p, nil
}

func (parser Parser) Identifier(fs afero.Fs, path string) bool {
	_, err := fs.Stat(filepath.Join(path, "problem.xml"))
	return !os.IsNotExist(err)
}

func init() {
	parser := Parser{
		CompileBinaries: true,
		EnableHTMLGen:   true,
	}
	_ = problems.RegisterConfigType("polygon", parser.Parse, parser.Identifier)
}
