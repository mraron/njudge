package taskarchive

import (
	"net/http"

	"github.com/mraron/njudge/internal/njudge"
	"github.com/mraron/njudge/internal/web/helpers/i18n"

	"github.com/labstack/echo/v4"
	"github.com/mraron/njudge/pkg/problems"
)

type TaskArchive struct {
	Roots []TreeNode
}

type TreeNode struct {
	ID           int
	Type         string
	Name         string
	Link         string
	SolvedStatus njudge.SolvedStatus
	Children     []TreeNode
	Visible      bool
}

func Get(cats njudge.Categories, problemQuery njudge.ProblemQuery, solvedStatusQuery njudge.SolvedStatusQuery, problemStore problems.Store) echo.HandlerFunc {
	return func(c echo.Context) error {
		tr := c.Get(i18n.TranslatorContextKey).(i18n.Translator)

		u := c.Get("user").(*njudge.User)

		lst, err := cats.GetAllWithParent(c.Request().Context(), 0)
		if err != nil {
			return err
		}

		taskArchive := TaskArchive{Roots: make([]TreeNode, 0)}

		var dfs func(category njudge.Category, node *TreeNode) error
		id := 1000

		dfs = func(root njudge.Category, tree *TreeNode) error {
			problemList, err := problemQuery.GetProblemsWithCategory(c.Request().Context(), njudge.NewCategoryIDFilter(root.ID))
			if err != nil {
				return err
			}

			for _, p := range problemList {
				elem := TreeNode{
					ID:           id,
					Type:         "problem",
					Name:         tr.TranslateContent(problemStore.MustGetProblem(p.Problem).Titles()).String(),
					Link:         c.Echo().Reverse("getProblemMain", p.Problemset, p.Problem),
					Children:     make([]TreeNode, 0),
					SolvedStatus: -1,
					Visible:      p.Visible,
				}

				if u != nil {
					solvedStatus, err := solvedStatusQuery.GetSolvedStatus(c.Request().Context(), p.ID, u.ID)
					if err != nil {
						return err
					}
					elem.SolvedStatus = solvedStatus
				}

				tree.Children = append(tree.Children, elem)

				id++
			}

			subCategories, err := cats.GetAllWithParent(c.Request().Context(), root.ID)
			if err != nil {
				return err
			}

			for _, cat := range subCategories {
				if !cat.Visible {
					if u == nil || u.Role != "admin" {
						continue
					}
				}

				tree.Children = append(tree.Children, TreeNode{
					ID:           cat.ID,
					Type:         "category",
					Name:         cat.Name,
					Link:         "",
					Children:     make([]TreeNode, 0),
					SolvedStatus: -1,
					Visible:      cat.Visible,
				})

				if err := dfs(cat, &tree.Children[len(tree.Children)-1]); err != nil {
					return err
				}
			}

			return nil
		}

		for _, start := range lst {
			if !start.Visible {
				if u == nil || u.Role != "admin" {
					continue
				}
			}
			taskArchive.Roots = append(taskArchive.Roots, TreeNode{
				ID:           start.ID,
				Type:         "category",
				Name:         start.Name,
				Link:         "",
				Children:     make([]TreeNode, 0),
				SolvedStatus: -1,
				Visible:      start.Visible,
			})

			if dfs(start, &taskArchive.Roots[len(taskArchive.Roots)-1]) != nil {
				return err
			}
		}

		c.Set("title", tr.Translate("Archive"))
		return c.Render(http.StatusOK, "task_archive.gohtml", taskArchive)
	}
}
