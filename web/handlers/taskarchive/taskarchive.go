package taskarchive

import (
	"fmt"
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/mraron/njudge/pkg/problems"
	"github.com/mraron/njudge/web/helpers"
	"github.com/mraron/njudge/web/helpers/i18n"
	"github.com/mraron/njudge/web/models"
	. "github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type TreeNode struct {
	Id           int
	Type         string
	Name         string
	Link         string
	SolvedStatus helpers.SolvedStatus
	Children     []*TreeNode
}

//@TODO optimize this to use less queries, most likely caching it
func Get(DB *sqlx.DB, problemStore problems.Store) echo.HandlerFunc {
	return func(c echo.Context) error {
		u := c.Get("user").(*models.User)

		lst, err := models.ProblemCategories(Where("parent_id IS NULL")).All(DB)
		if err != nil {
			return err
		}

		roots := make([]*TreeNode, 0)

		var dfs func(category *models.ProblemCategory, node *TreeNode) error
		id := 1000
		dfs = func(root *models.ProblemCategory, tree *TreeNode) error {
			problems, err := models.ProblemRels(Where("category_id = ?", root.ID), OrderBy("problem")).All(DB)
			if err != nil {
				return err
			}

			for _, problem := range problems {
				elem := &TreeNode{Id: id, Type: "problem", Name: i18n.TranslateContent("hungarian", problemStore.MustGet(problem.Problem).Titles()).String(), Link: fmt.Sprintf("/problemset/%s/%s/", problem.Problemset, problem.Problem), Children: make([]*TreeNode, 0), SolvedStatus: -1}
				if u != nil {
					elem.SolvedStatus, err = helpers.HasUserSolved(DB, u, problem.Problemset, problem.Problem)
					if err != nil {
						return err
					}
				}

				tree.Children = append(tree.Children, elem)

				id++
			}

			//@TODO make a way to control sorting order from DB (add migrations etc.)
			subcats, err := models.ProblemCategories(Where("parent_id = ?", root.ID), OrderBy("name")).All(DB)
			if err != nil {
				return err
			}

			for _, cat := range subcats {
				akt := &TreeNode{Id: cat.ID, Type: "category", Name: cat.Name, Link: "", Children: make([]*TreeNode, 0), SolvedStatus: -1}
				tree.Children = append(tree.Children, akt)
				if err := dfs(cat, akt); err != nil {
					return err
				}
			}

			return nil
		}

		for _, start := range lst {
			roots = append(roots, &TreeNode{Id: start.ID, Type: "category", Name: start.Name, Link: "", Children: make([]*TreeNode, 0), SolvedStatus: -1})
			if dfs(start, roots[len(roots)-1]) != nil {
				return err
			}
		}

		c.Set("title", "Archívum")
		return c.Render(http.StatusOK, "task_archive.gohtml", roots)
	}
}
