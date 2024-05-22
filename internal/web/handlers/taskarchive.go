package handlers

import (
	"github.com/a-h/templ"
	"github.com/mraron/njudge/internal/njudge"
	"github.com/mraron/njudge/internal/web/helpers/i18n"
	"github.com/mraron/njudge/internal/web/templates"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetTaskArchive(tas njudge.TaskArchiveService) echo.HandlerFunc {
	return func(c echo.Context) error {
		tr := c.Get(i18n.TranslatorContextKey).(i18n.Translator)

		u := c.Get("user").(*njudge.User)

		var (
			rootStr = c.QueryParam("root")
			root    = 0

			res *njudge.TaskArchiveNode
			cat *njudge.Category
			err error
		)
		vm := templates.TaskArchiveViewModel{
			ProblemLinks: make(map[int]templates.Link),
		}
		if root, err = strconv.Atoi(rootStr); err == nil {
			cat, err = tas.Categories.Get(c.Request().Context(), root)
			if err != nil {
				return err
			}
			res, err = tas.CreateWithCategory(c.Request().Context(), *cat, u)
			if err != nil {
				return err
			}
			vm.Root = *res
		} else {
			res, err = tas.CreateTopLevel(c.Request().Context(), u)
			if err != nil {
				return err
			}
			vm.Root = *res
		}

		res.Search(func(node *njudge.TaskArchiveNode) bool {
			if node.Type == njudge.TaskArchiveNodeProblem {
				vm.ProblemLinks[node.ID] = templates.Link{
					Text: tr.TranslateContent(tas.ProblemStore.MustGetProblem(node.Problem.Problem).Titles()).String(),
					Href: templ.SafeURL(c.Echo().Reverse("getProblemMain", node.Problem.Problemset, node.Problem.Problem)),
				}
			}
			return false
		})
		c.Set("title", tr.Translate("Archive"))
		return templates.Render(c, http.StatusOK, templates.TaskArchive(vm))
	}
}
