package web

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/mraron/njudge/utils/problems/config/polygon"
	"github.com/mraron/njudge/web/models"
	. "github.com/volatiletech/sqlboiler/v4/queries/qm"
	"io/ioutil"
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

	return c.Render(http.StatusOK, "problemsetproblem.html", s.getProblem(problem))
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

	f, err := ioutil.ReadFile(fileLoc)
	if err != nil {
		return c.String(http.StatusInternalServerError, "error serving file")
	}

	return c.Blob(http.StatusOK, mime.TypeByExtension(filepath.Ext(fileLoc)), f)
}

func (s *Server) getProblemsetProblemAttachment(c echo.Context) error {
	p, attachment := s.getProblem(c.Param("problem")), c.Param("attachment")
	if p == nil {
		return c.String(http.StatusNotFound, "no such problem")
	}

	for _, val := range p.Attachments() {
		if val.Name == attachment {
			return c.Blob(http.StatusOK, mime.TypeByExtension(filepath.Ext(val.Name)), val.Contents)
		}
	}

	return c.String(http.StatusNotFound, "no such attachment")
}

type statusPage struct {
	Name string
	Active bool
	Disabled bool
	Url string
}
func (s *Server) getProblemsetStatus(c echo.Context) error {
	var (
		sbs []*models.Submission
		err error
		cnt int64
	)

	ac := c.QueryParam("ac")
	problem_set := c.QueryParam("problem_set")
	problem := c.QueryParam("problem")
	page, err := strconv.Atoi(c.QueryParam("page"))
	if  err != nil || page<=0 {
		page = 1
	}

	per_page := 20
	pagination := []QueryMod{Limit(per_page), Offset((page-1)*per_page)}

	query := []QueryMod{}
	if problem == "" {
		query = []QueryMod{OrderBy("id DESC")}
	} else {
		if ac == "1" {
			query = []QueryMod{OrderBy("id DESC"), Where("verdict = 0"), Where("problem = ?", problem), Where("problemset = ?", problem_set)}
		} else {
			query = []QueryMod{OrderBy("id DESC"), Where("problem = ?", problem), Where("problemset = ?", problem_set)}
		}
	}

	sbs, err = models.Submissions(append(pagination, query...)...).All(s.db)
	if err != nil {
		return s.internalError(c, err, "Belső hiba 1")
	}

	cnt, err = models.Submissions(query[1:]...).Count(s.db)
	if err != nil {
		return s.internalError(c, err, "Belső hiba 2")
	}

	qu := c.Request().URL.Query()


	pageCnt := (int(cnt)+per_page-1)/per_page
	pages := make([]statusPage, pageCnt+2)
	pages[0] = statusPage{"Előző", false, true, "#"}
	if page>1 {
		qu.Set("page", strconv.Itoa(page-1))

		pages[0].Disabled = false
		pages[0].Url = "?"+qu.Encode()
	}
	for i := 1; i < len(pages)-1; i++ {
		qu.Set("page", strconv.Itoa(i))
		pages[i] = statusPage{strconv.Itoa(i), false, false, "?"+qu.Encode()}
		if i==page {
			pages[i].Active = true
		}
	}
	pages[len(pages)-1] = statusPage{"Következő", false, true, "#"}
	if page<pageCnt {
		qu.Set("page", strconv.Itoa(page+1))

		pages[len(pages)-1].Disabled = false
		pages[len(pages)-1].Url = "?"+qu.Encode()
	}

	if page>len(pages) {
		return s.internalError(c, err, "Nincs ilyen oldal")
	}

	return c.Render(http.StatusOK, "status.html", struct {
		CurrentPage int
		Pages []statusPage
		Submissions []*models.Submission
	}{page, pages, sbs})
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

	return c.Render(http.StatusOK, "task_archive.html", roots)
}
