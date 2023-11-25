package problem

import (
	"errors"
	"os"
	"path/filepath"
	"strings"

	"github.com/mraron/njudge/internal/njudge/db/models"
	"github.com/mraron/njudge/internal/web/domain/submission"
	"github.com/mraron/njudge/internal/web/helpers/i18n"
	"github.com/mraron/njudge/internal/web/helpers/ui"
	"github.com/mraron/njudge/pkg/problems"
	"github.com/mraron/njudge/pkg/problems/config/polygon"
)

var (
	ErrorFileNotFound      = errors.New("file not found")
	ErrorStatementNotFound = errors.New("statement not found")
)

type Problem struct {
	problems.Problem
	models.ProblemRel
}

func (p *Problem) GetPDF(lang string) ([]byte, error) {
	if len(p.Statements().FilterByType(problems.DataTypePDF)) == 0 {
		return nil, ErrorStatementNotFound
	}

	return i18n.TranslateContent(lang, p.Statements().FilterByType(problems.DataTypePDF)).Value()
}

func (p *Problem) GetFile(file string) (fileLoc string, err error) {
	switch p := p.Problem.(problems.ProblemWrapper).Problem.(type) {
	case polygon.Problem:
		if len(p.Statements().FilterByType(problems.DataTypeHTML)) == 0 || strings.HasSuffix(file, ".tex") || strings.HasSuffix(file, ".json") {
			err = ErrorFileNotFound
		}

		fileLoc = filepath.Join(p.Path, "statements", ".html", p.HTMLStatements()[0].Locale(), file)
		if _, err := os.Stat(fileLoc); err != nil {
			fileLoc = filepath.Join(p.Path, "statements", p.HTMLStatements()[0].Locale(), file)
		}
	default:
		err = ErrorFileNotFound
	}

	return
}

func (p *Problem) GetAttachment(attachment string) (problems.NamedData, error) {
	for _, val := range p.Attachments() {
		if val.Name() == attachment {
			return val, nil
		}
	}

	return nil, ErrorFileNotFound
}

type StatsData struct {
	SolverCount  int
	SolvedStatus SolvedStatus
	LastLanguage string
	CategoryLink ui.Link
	CategoryID   int
	Tags         []Tag
	Submissions  []submission.Submission
}
