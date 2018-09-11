package web

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"github.com/mraron/njudge/utils/problems"
	"github.com/mraron/njudge/web/models"
	"github.com/mraron/njudge/web/roles"
	"html/template"
	"io"
	"time"
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
		"canView": func(role string, entity roles.Entity) bool {
			return roles.Can(roles.Role(role), roles.ActionView, entity)
		},
		"get": func(c echo.Context, key string) interface{} {
			return c.Get(key)
		},
		"fixedlen": func(a int, len int) string {
			return fmt.Sprintf(fmt.Sprintf("%%0%dd", len), a)
		},
		"month2int": func(month time.Month) int {
			return int(month)
		},
		"decr": func(val int) int {
			return val - 1
		},
		"add": func(a, b int) int {
			return a + b
		},
		"parseStatus": func(s string) *problems.Status {
			st := &problems.Status{}
			err := json.Unmarshal([]byte(s), st)
			if err != nil {
				log.Debug(err)
			}

			return st
		},
	}
}
