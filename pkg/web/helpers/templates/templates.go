package templates

import (
	"crypto/md5"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"io"
	"io/fs"
	"math"
	"path/filepath"
	"strings"
	"time"

	"github.com/mraron/njudge/pkg/web/helpers/config"
	"github.com/mraron/njudge/pkg/web/helpers/i18n"
	"github.com/mraron/njudge/pkg/web/helpers/roles"
	"github.com/mraron/njudge/pkg/web/models"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/mraron/njudge/pkg/problems"
)

type Renderer struct {
	templates map[string]*template.Template
	cfg       config.Server
}

func New(cfg config.Server, problemStore problems.Store, db *sql.DB) *Renderer {
	renderer := &Renderer{make(map[string]*template.Template), cfg}

	layoutFiles := make([]string, 0)
	err := filepath.Walk(cfg.TemplatesDir, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			if strings.HasPrefix(info.Name(), "_") {
				layoutFiles = append(layoutFiles, path)
			} else {
				name, err := filepath.Rel(cfg.TemplatesDir, path)
				if err != nil {
					panic(err)
				}

				renderer.templates[name] = template.Must(template.New(info.Name()).Funcs(funcs(problemStore, db)).ParseFiles(append(layoutFiles, path)...))
			}
		}
		return nil
	})

	if err != nil {
		panic(err)
	}

	return renderer
}

func (t *Renderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	if !strings.HasSuffix(name, ".gohtml") {
		name += ".gohtml"
	}

	if _, ok := t.templates[name]; !ok {
		return fmt.Errorf("can find template %q", name)
	}

	return t.templates[name].ExecuteTemplate(w, filepath.Base(name), struct {
		Data       interface{}
		Context    echo.Context
		CustomHead string
	}{data, c, t.cfg.CustomHead})
}

func funcs(store problems.Store, db *sql.DB) template.FuncMap {
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
		"fixedlenFloat32": func(a float32, len int) string {
			return fmt.Sprintf(fmt.Sprintf("%%.%df", len), a)
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
		"dict": func(values ...interface{}) (map[string]interface{}, error) {
			if len(values)%2 != 0 {
				return nil, errors.New("invalid dict call")
			}
			dict := make(map[string]interface{}, len(values)/2)
			for i := 0; i < len(values); i += 2 {
				key, ok := values[i].(string)
				if !ok {
					return nil, errors.New("dict keys must be strings")
				}
				dict[key] = values[i+1]
			}
			return dict, nil
		},
		"partial": NewPartialCache(db, 30*time.Second).Get,
		"roundto": func(num float64, digs int) float64 {
			return math.Round(num*100) / 100
		},
		"tags": func() (models.TagSlice, error) {
			return models.Tags().All(db)
		},
	}
}