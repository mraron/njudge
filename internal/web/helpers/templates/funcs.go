package templates

import (
	"context"
	"crypto/md5"
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/mraron/njudge/internal/njudge"
	"github.com/mraron/njudge/internal/web/domain/problem"
	"github.com/mraron/njudge/internal/web/helpers"
	"github.com/mraron/njudge/internal/web/helpers/i18n"
	"github.com/mraron/njudge/internal/web/helpers/roles"
	"github.com/mraron/njudge/internal/web/helpers/templates/partials"
	"github.com/mraron/njudge/internal/web/models"
	"github.com/mraron/njudge/pkg/problems"
	"golang.org/x/text/message"
)

func contextFuncs(c echo.Context) template.FuncMap {
	return template.FuncMap{
		"logged": func() bool {
			if _, ok := c.Get("user").(*njudge.User); ok {
				return nil != c.Get("user").(*njudge.User)
			}

			return false
		},
		"user": func() *njudge.User {
			if _, ok := c.Get("user").(*njudge.User); ok {
				return c.Get("user").(*njudge.User)
			}

			return nil
		},
		"get": func(key string) interface{} {
			return c.Get(key)
		},
		"getCookie": func(name string) *string {
			val, err := c.Cookie(name)
			if err != nil {
				return nil
			}

			return &val.Value
		},
		"getFlash": func(name string) interface{} {
			return helpers.GetFlash(c, name)
		},
		"csrf": func() string {
			return c.Get(middleware.DefaultCSRFConfig.ContextKey).(string)
		},
		"Tr": func(key message.Reference, args ...interface{}) string {
			return c.Get(i18n.TranslatorContextKey).(i18n.Translator).Translate(key, args...)
		},
	}
}

func statelessFuncs(store problems.Store, tags njudge.Tags, store2 partials.Store) template.FuncMap {
	return template.FuncMap{
		"translateContent": i18n.TranslateContent,
		"problem":          store.Get,
		"str2html": func(s string) template.HTML {
			return template.HTML(s)
		},
		"canView": func(role string, entity roles.Entity) bool {
			return roles.Can(roles.Role(role), roles.ActionView, entity)
		},
		"fixedLen": func(a int, len int) string {
			return fmt.Sprintf(fmt.Sprintf("%%0%dd", len), a)
		},
		"fixedLenFloat32": func(a float32, len int) string {
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
				panic(err)
			}

			return st
		},
		"divide": func(a, b int) int {
			return a / b
		},
		"toString": func(b interface{}) string {
			switch b := b.(type) {
			case []byte:
				return string(b)
			case int:
				return strconv.Itoa(b)
			}

			return ""
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
		"partial": func(name string) string {
			c, _ := store2.Get(name)
			return c
		},
		"roundTo": func(num float64, digs int) float64 {
			return math.Round(num*100) / 100
		},
		"tags": func() ([]njudge.Tag, error) {
			return tags.GetAll(context.Background())
		},
		"contextTODO": func() context.Context {
			return context.TODO()
		},
		"verdict": problem.VerdictFromProblemsVerdictName,
	}
}
