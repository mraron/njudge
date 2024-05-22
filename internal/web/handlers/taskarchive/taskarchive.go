package taskarchive

import (
	"github.com/a-h/templ"
	"github.com/mraron/njudge/internal/njudge"
	"github.com/mraron/njudge/internal/web/helpers/i18n"
	"github.com/mraron/njudge/internal/web/templates"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Get(tas njudge.TaskArchiveService) echo.HandlerFunc {
	return func(c echo.Context) error {
		tr := c.Get(i18n.TranslatorContextKey).(i18n.Translator)

		u := c.Get("user").(*njudge.User)
		res, err := tas.CreateFromAllTopLevel(c.Request().Context(), u)
		if err != nil {
			return err
		}
		vm := templates.TaskArchiveViewModel{
			Root:         *res,
			ProblemLinks: make(map[int]templates.Link),
		}
		res.Search(func(node njudge.TaskArchiveNode) bool {
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
