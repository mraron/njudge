package web

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/mraron/njudge/utils/problems/polygon"
	"github.com/mraron/njudge/web/models"
	. "github.com/volatiletech/sqlboiler/queries/qm"
	"io/ioutil"
	"mime"
	"net/http"
	"path/filepath"
)

func (s *Server) getProblemsetMain(c echo.Context) error {
	name := c.Param("name")
	lst, err := models.ProblemRels(Where("problemset=?", name)).All(s.db)
	//lst, err := models.ProblemsFromProblemset(s.db, name)
	if err != nil {
		return s.internalError(c, err, "Belső hiba.")
	}

	if len(lst) == 0 {
		return c.Render(http.StatusNotFound, "404.html", "Nem található.")
	}

	return c.Render(http.StatusOK, "problemsetmain.html", lst)
}

func (s *Server) getProblemsetProblem(c echo.Context) error {
	name, problem := c.Param("name"), c.Param("problem")

	lst, err := models.ProblemRels(Where("problemset=?", name)).All(s.db)
	if err != nil {
		return s.internalError(c, err, "Belső hiba.")
	}

	ok := false
	for _, val := range lst {
		if val.Problem == problem {
			ok = true
		}

		if ok {
			break
		}
	}

	if !ok {
		return c.JSON(http.StatusNotFound, nil)
	}

	return c.Render(http.StatusOK, "problemsetproblem.html", s.problems[problem])
}

func (s *Server) getProblemsetProblemPDFLanguage(c echo.Context) error {
	p, lang := s.problems[c.Param("problem")], c.Param("language")

	if p == nil {
		return c.String(http.StatusNotFound, "no such problem")
	}

	if len(p.PDFStatements()) == 0 {
		return c.String(http.StatusNotFound, "no pdf statement")
	}

	return c.Blob(http.StatusOK, "application/pdf", translateContent(lang, p.PDFStatements()).Contents)
}

func (s *Server) getProblemsetProblemFile(c echo.Context) error {
	p := s.problems[c.Param("problem")]

	if p == nil {
		return c.String(http.StatusNotFound, "not found")
	}

	fileLoc := ""

	switch p.(type) {
	case polygon.Problem:
		if len(p.HTMLStatements()) == 0 {
			return c.String(http.StatusNotFound, "file not found")
		}

		fileLoc = filepath.Join(s.ProblemsDir, p.Name(), "statements", ".html", p.HTMLStatements()[0].Locale, c.Param("file"))
	default:
		return c.String(http.StatusNotFound, "not found")
	}

	f, err := ioutil.ReadFile(fileLoc)
	if err != nil {
		return c.String(http.StatusInternalServerError, "error serving file")
	}

	return c.Blob(http.StatusOK, mime.TypeByExtension(filepath.Ext(fileLoc)), f)
}

func (s *Server) getProblemsetProblemAttachment(c echo.Context) error {
	p, attachment := s.problems[c.Param("problem")], c.Param("attachment")
	if p == nil {
		return c.String(http.StatusNotFound, "no such problem")
	}

	for _, val := range p.Attachments() {
		fmt.Println(val.Name, attachment)
		if val.Name == attachment {
			return c.Blob(http.StatusOK, mime.TypeByExtension(filepath.Ext(val.Name)), val.Contents)
		}
	}

	return c.String(http.StatusNotFound, "no such attachment")
}

func (s *Server) getProblemsetStatus(c echo.Context) error {
	sbs, err := models.Submissions(OrderBy("id DESC")).All(s.db)
	if err != nil {
		return s.internalError(c, err, "Belső hiba.")
	}

	return c.Render(http.StatusOK, "status.html", sbs)
}
