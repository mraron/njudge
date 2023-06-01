package taskarchive

import (
	"github.com/mraron/njudge/internal/web/domain/problem"
	"github.com/mraron/njudge/internal/web/helpers"
	"github.com/mraron/njudge/internal/web/helpers/i18n"
	"github.com/mraron/njudge/internal/web/models"
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/mraron/njudge/pkg/problems"
	. "github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type TaskArchive struct {
	Roots []TreeNode
}

type TreeNode struct {
	ID           int
	Type         string
	Name         string
	Link         string
	SolvedStatus problem.SolvedStatus
	Children     []TreeNode
}

func Get(DB *sqlx.DB, problemStore problems.Store) echo.HandlerFunc {
	return func(c echo.Context) error {
		tr := c.Get("translator").(i18n.Translator)

		u := c.Get("user").(*models.User)

		lst, err := models.ProblemCategories(Where("parent_id IS NULL")).All(c.Request().Context(), DB)
		if err != nil {
			return err
		}

		taskArchive := TaskArchive{Roots: make([]TreeNode, 0)}

		var dfs func(category *models.ProblemCategory, node TreeNode) error
		id := 1000

		dfs = func(root *models.ProblemCategory, tree TreeNode) error {
			problemList, err := models.ProblemRels(Where("category_id = ?", root.ID), OrderBy("problem")).All(c.Request().Context(), DB)
			if err != nil {
				return err
			}

			for _, p := range problemList {
				elem := TreeNode{
					ID:           id,
					Type:         "problem",
					Name:         i18n.TranslateContent("hungarian", problemStore.MustGet(p.Problem).Titles()).String(),
					Link:         c.Echo().Reverse("getProblemMain", p.Problemset, p.Problem),
					Children:     make([]TreeNode, 0),
					SolvedStatus: -1,
				}

				if u != nil {
					elem.SolvedStatus, err = helpers.HasUserSolved(DB.DB, u.ID, p.Problemset, p.Problem)
					if err != nil {
						return err
					}
				}

				tree.Children = append(tree.Children, elem)

				id++
			}

			subCategories, err := models.ProblemCategories(Where("parent_id = ?", root.ID), OrderBy("name")).All(c.Request().Context(), DB)
			if err != nil {
				return err
			}

			for _, cat := range subCategories {
				akt := TreeNode{
					ID:           cat.ID,
					Type:         "category",
					Name:         cat.Name,
					Link:         "",
					Children:     make([]TreeNode, 0),
					SolvedStatus: -1,
				}

				tree.Children = append(tree.Children, akt)
				if err := dfs(cat, akt); err != nil {
					return err
				}
			}

			return nil
		}

		for _, start := range lst {
			taskArchive.Roots = append(taskArchive.Roots, TreeNode{
				ID:           start.ID,
				Type:         "category",
				Name:         start.Name,
				Link:         "",
				Children:     make([]TreeNode, 0),
				SolvedStatus: -1,
			})

			if dfs(start, taskArchive.Roots[len(taskArchive.Roots)-1]) != nil {
				return err
			}
		}

		c.Set("title", tr.Translate("Archive"))
		return c.Render(http.StatusOK, "task_archive.gohtml", taskArchive)
	}
}
