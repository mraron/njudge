package web

import (
	"bytes"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/mraron/njudge/utils/problems"
	"github.com/mraron/njudge/utils/problems/config/polygon"
	"github.com/mraron/njudge/web/models"
	. "github.com/volatiletech/sqlboiler/v4/queries/qm"
	"io"
	"mime"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"
)

func (s *Server) getProblemsetMain(c echo.Context) error {
	name := c.Param("name")
	lst, err := models.ProblemRels(Where("problemset=?", name)).All(s.db)
	//lst, err := models.ProblemsFromProblemset(s.db, name)
	fmt.Println(lst)
	if err != nil {
		return s.internalError(c, err, "Belső hiba.")
	}

	if len(lst) == 0 {
		return c.Render(http.StatusNotFound, "404.gohtml", "Nem található.")
	}

	return c.Render(http.StatusOK, "problemset_list.gohtml", lst)
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

	lastLanguage := ""
	if u := c.Get("user").(*models.User); u != nil {
		fmt.Println(u)
		sub, err := models.Submissions(Select("language"), Where("user_id = ?", u.ID), OrderBy("id DESC"), Limit(1)).One(s.db)
		if err == nil {
			lastLanguage = sub.Language
		}
	}

	return c.Render(http.StatusOK, "problemset_problem.gohtml",struct{
		Problem problems.Problem
		LastLanguage string
	}{s.getProblem(problem), lastLanguage})
}

func (s *Server) getProblemsetProblemPDFLanguage(c echo.Context) error {
	p, lang := s.getProblem(c.Param("problem")), c.Param("language")

	if p == nil {
		return c.String(http.StatusNotFound, "no such problem")
	}

	if len(p.PDFStatements()) == 0 {
		return c.String(http.StatusNotFound, "no pdf statement")
	}

	return c.Blob(http.StatusOK, "application/pdf", translateContent(lang, p.PDFStatements()).Contents)
}

func (s *Server) getProblemsetProblemFile(c echo.Context) error {
	p := s.getProblem(c.Param("problem"))

	if p == nil {
		return c.String(http.StatusNotFound, "not found")
	}

	fileLoc := ""

	switch p.(type) {
	case polygon.Problem:
		if len(p.HTMLStatements()) == 0 {
			return c.String(http.StatusNotFound, "file not found")
		}

		if strings.HasSuffix(c.Param("file"), ".css") {
			fileLoc = filepath.Join(s.ProblemsDir, p.Name(), "statements", ".html", p.HTMLStatements()[0].Locale, c.Param("file"))
		} else {
			fileLoc = filepath.Join(s.ProblemsDir, p.Name(), "statements", p.HTMLStatements()[0].Locale, c.Param("file"))
		}

	default:
		return c.String(http.StatusNotFound, "not found")
	}

	return c.Attachment(fileLoc, c.Param("file"))
}

func (s *Server) getProblemsetProblemAttachment(c echo.Context) error {
	p, attachment := s.getProblem(c.Param("problem")), c.Param("attachment")
	if p == nil {
		return c.String(http.StatusNotFound, "no such problem")
	}

	for _, val := range p.Attachments() {
		if val.Name == attachment {
			c.Response().Header().Set("Content-Disposition", "attachment; filename="+val.Name)
			c.Response().Header().Set("Content-Type", mime.TypeByExtension(filepath.Ext(val.Name)))
			c.Response().Header().Set("Content-Length", strconv.Itoa(len(val.Contents)))

			io.Copy(c.Response(), bytes.NewReader(val.Contents))

			return c.NoContent(http.StatusOK)
		}
	}

	return c.String(http.StatusNotFound, "no such attachment")
}

type taskArchiveTreeNode struct {
	Id       int
	Type     string
	Name     string
	Link     string
	Children []*taskArchiveTreeNode
}

//@TODO optimize this to use less queries
func (s *Server) getTaskArchive(c echo.Context) error {
	lst, err := models.ProblemCategories(Where("parent_id IS NULL")).All(s.db)
	if err != nil {
		return s.internalError(c, err, "Belső hiba.")
	}

	roots := make([]*taskArchiveTreeNode, 0)

	var dfs func(category *models.ProblemCategory, node *taskArchiveTreeNode) error
	id := 1000
	dfs = func(root *models.ProblemCategory, tree *taskArchiveTreeNode) error {
		problems, err := models.ProblemRels(Where("category_id = ?", root.ID), OrderBy("problem")).All(s.db)
		if err != nil {
			return err
		}

		for _, problem := range problems {
			tree.Children = append(tree.Children, &taskArchiveTreeNode{id, "problem", translateContent("hungarian", s.getProblem(problem.Problem).Titles()).String(), fmt.Sprintf("/problemset/%s/%s/", problem.Problemset, problem.Problem), make([]*taskArchiveTreeNode, 0)})
			id++
		}

		//@TODO make a way to control sorting order from db (add migrations etc.)
		subcats, err := models.ProblemCategories(Where("parent_id = ?", root.ID), OrderBy("name")).All(s.db)
		if err != nil {
			return err
		}

		for _, cat := range subcats {
			akt := &taskArchiveTreeNode{cat.ID, "category", cat.Name, "", make([]*taskArchiveTreeNode, 0)}
			tree.Children = append(tree.Children, akt)
			if err := dfs(cat, akt); err != nil {
				return err
			}
		}

		return nil
	}

	for _, start := range lst {
		roots = append(roots, &taskArchiveTreeNode{start.ID, "category", start.Name, "", make([]*taskArchiveTreeNode, 0)})
		if dfs(start, roots[len(roots)-1]) != nil {
			return s.internalError(c, err, "Belső hiba.")
		}
	}

	return c.Render(http.StatusOK, "task_archive.gohtml", roots)
}
