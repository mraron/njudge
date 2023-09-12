package taskarchive

import (
	"net/http"

	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	"github.com/mraron/njudge/internal/web/domain/problem"
	"github.com/mraron/njudge/internal/web/helpers"
	"github.com/mraron/njudge/internal/web/helpers/i18n"
	"github.com/mraron/njudge/internal/web/models"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/mraron/njudge/pkg/problems"
)

type TaskArchive struct {
	Roots []TreeNode `json:"categories"`
}

type TreeNode struct {
	ID           int                  `json:"-"`
	Type         string               `json:"type"`
	Name         string               `json:"title"`
	Link         string               `json:"link"`
	SolvedStatus problem.SolvedStatus `json:"solvedStatus"`
	Children     []TreeNode           `json:"children"`
}

func MakeTaskArchive(c echo.Context, tr i18n.Translator, DB *sqlx.DB, problemStore problems.Store, u *models.User) (*TaskArchive, error) {
	lst, err := models.ProblemCategories(models.ProblemCategoryWhere.ParentID.IsNull()).All(c.Request().Context(), DB)
	if err != nil {
		return nil, err
	}

	taskArchive := &TaskArchive{Roots: make([]TreeNode, 0)}

	var dfs func(category *models.ProblemCategory, node *TreeNode) error
	id := 1000

	dfs = func(root *models.ProblemCategory, tree *TreeNode) error {
		problemList, err := models.ProblemRels(models.ProblemRelWhere.CategoryID.EQ(null.Int{
			Int:   root.ID,
			Valid: true,
		}), qm.OrderBy("problem")).All(c.Request().Context(), DB)
		if err != nil {
			return err
		}

		for _, p := range problemList {
			elem := TreeNode{
				ID:           id,
				Type:         "problem",
				Name:         tr.TranslateContent(problemStore.MustGet(p.Problem).Titles()).String(),
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

		subCategories, err := models.ProblemCategories(models.ProblemCategoryWhere.ParentID.EQ(null.Int{
			Int:   root.ID,
			Valid: true,
		}), qm.OrderBy("name")).All(c.Request().Context(), DB)

		if err != nil {
			return err
		}

		for _, cat := range subCategories {
			tree.Children = append(tree.Children, TreeNode{
				ID:           cat.ID,
				Type:         "category",
				Name:         cat.Name,
				Link:         "",
				Children:     make([]TreeNode, 0),
				SolvedStatus: -1,
			})

			if err := dfs(cat, &tree.Children[len(tree.Children)-1]); err != nil {
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

		if dfs(start, &taskArchive.Roots[len(taskArchive.Roots)-1]) != nil {
			return nil, err
		}
	}

	return taskArchive, nil
}

func Get(DB *sqlx.DB, problemStore problems.Store) echo.HandlerFunc {
	return func(c echo.Context) error {
		tr := c.Get(i18n.TranslatorContextKey).(i18n.Translator)

		u := c.Get("user").(*models.User)

		taskArchive, err := MakeTaskArchive(c, tr, DB, problemStore, u)
		if err != err {
			return err
		}

		c.Set("title", tr.Translate("Archive"))
		return c.Render(http.StatusOK, "task_archive.gohtml", taskArchive)
	}
}
