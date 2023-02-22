package problem

import (
	"errors"
	"path/filepath"
	"strings"

	"github.com/mraron/njudge/internal/web/helpers"
	"github.com/mraron/njudge/internal/web/helpers/i18n"
	"github.com/mraron/njudge/internal/web/models"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/mraron/njudge/pkg/problems"
	"github.com/mraron/njudge/pkg/problems/config/polygon"
	. "github.com/volatiletech/sqlboiler/v4/queries/qm"
)

var (
	ErrorFileNotFound      = errors.New("file not found")
	ErrorStatementNotFound = errors.New("statement not found")
)

type Problem struct {
	problems.Problem
	ProblemRel *models.ProblemRel

	SolverCount  int
	SolvedStatus helpers.SolvedStatus
	LastLanguage string
	CategoryLink helpers.Link
	CategoryId   int
	Tags         []*models.Tag
	Submissions  []*models.Submission
}

func New(c echo.Context) *Problem {
	return &Problem{Problem: c.Get("problem").(problems.Problem), ProblemRel: c.Get("ProblemRel").(*models.ProblemRel)}
}

func (p *Problem) FillFields(c echo.Context, DB *sqlx.DB) error {
	var err error
	p.SolverCount = p.ProblemRel.SolverCount
	if u := c.Get("user").(*models.User); u != nil {
		p.LastLanguage = helpers.GetUserLastLanguage(c, DB)
		p.SolvedStatus, err = helpers.HasUserSolved(DB, u, p.ProblemRel.Problemset, p.ProblemRel.Problem)
		if err != nil {
			return err
		}
		p.Submissions, err = models.Submissions(Where("problemset = ?", p.ProblemRel.Problemset), Where("problem = ?", p.ProblemRel.Problem), Where("user_id = ?", u.ID), OrderBy("id DESC"), Limit(5)).All(DB)
		if err != nil {
			return err
		}
	}

	if p.ProblemRel.CategoryID.Valid {
		p.CategoryId = p.ProblemRel.CategoryID.Int
		p.CategoryLink, err = helpers.TopCategoryLink(p.CategoryId, DB)
		if err != nil {
			return err
		}
	}

	tags, err := models.Tags(InnerJoin("problem_tags pt on pt.tag_id = tags.id"), Where("pt.problem_id = ?", p.ProblemRel.ID)).All(DB)
	if err != nil {
		return err
	}
	p.Tags = tags

	return nil
}

func (p *Problem) GetPDF(lang string) ([]byte, error) {
	if len(p.PDFStatements()) == 0 {
		return nil, ErrorStatementNotFound
	}

	return i18n.TranslateContent(lang, p.PDFStatements()).Value()
}

func (p *Problem) GetFile(file string) (fileLoc string, err error) {
	switch p := p.Problem.(problems.ProblemWrapper).Problem.(type) {
	case polygon.Problem:
		if len(p.HTMLStatements()) == 0 || strings.HasSuffix(file, ".tex") || strings.HasSuffix(file, ".json") {
			err = ErrorFileNotFound
		}

		if strings.HasSuffix(file, ".css") {
			fileLoc = filepath.Join(p.Path, "statements", ".html", p.HTMLStatements()[0].Locale(), file)
		} else {
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
