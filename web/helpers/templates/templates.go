package templates

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/mraron/njudge/utils/problems"
	"github.com/mraron/njudge/web/helpers/i18n"
	"github.com/mraron/njudge/web/helpers/roles"
	"github.com/mraron/njudge/web/models"
	"html/template"
	"io"
	"path/filepath"
	"strings"
	"time"
)

type Renderer struct {
	templates *template.Template
}

func New(templatesDir string, st problems.Store) *Renderer {
	return &Renderer{template.Must(template.New("").Funcs(funcs(st)).ParseGlob(filepath.Join(templatesDir, "*.gohtml")))}
}

func (t *Renderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	err := t.templates.ExecuteTemplate(w, name, struct {
		Data    interface{}
		Context echo.Context
	}{data, c})

	if err != nil {
		panic(err)
	}

	return nil
}

func funcs(store problems.Store) template.FuncMap {
	return template.FuncMap{
		"translateContent": i18n.TranslateContent,
		"problem":          store.Get,
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
		"divide": func(a, b int) int {
			return a / b
		},
		"tostring": func(b []byte) string {
			return string(b)
		},
		"gravatarHash": func(user *models.User) string {
			return fmt.Sprintf("%x", md5.Sum([]byte(strings.ToLower(strings.TrimSpace(user.Email)))))
		},
	}
}
