package web

import (
	"github.com/labstack/echo"
	"github.com/mraron/njudge/utils/problems"
	"github.com/mraron/njudge/web/models"
	"github.com/mraron/njudge/web/roles"
	"html/template"
	"io"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	err := t.templates.ExecuteTemplate(w, name, struct {
		Data    interface{}
		Context echo.Context
	}{data, c})

	if err != nil {
		panic(err)
	}

	return nil
}

func translateContent(locale string, cs []problems.Content) problems.Content {
	search := func(loc string) (problems.Content, bool) {
		for _, c := range cs {
			if locale == c.Locale {
				return c, true
			}
		}

		return problems.Content{}, false
	}

	if val, ok := search(locale); ok {
		return val
	}

	if val, ok := search("hungarian"); ok {
		return val
	}

	return cs[0]
}

func locales(cs []problems.Content) []string {
	lst := make(map[string]bool)
	for _, val := range cs {
		lst[val.Locale] = true
	}

	ans := make([]string, len(lst))

	ind := 0
	for key := range lst {
		ans[ind] = key
		ind++
	}

	return ans
}

func (s *Server) templatefuncs() template.FuncMap {
	return template.FuncMap{
		"locales":          locales,
		"translateContent": translateContent,
		"problem": func(name string) problems.Problem {
			return s.problems[name]
		},
		"str2html": func(s string) template.HTML {
			return template.HTML(s)
		},
		"logged": func(c echo.Context) bool {
			return nil != c.Get("user").(*models.User)
		},
		"user": func(c echo.Context) *models.User {
			return c.Get("user").(*models.User)
		},
		"canView": func(role roles.Role, entity roles.Entity) bool {
			return roles.Can(role, roles.ActionView, entity)
		},
	}
}
