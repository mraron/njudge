package templates

import (
	"database/sql"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/mraron/njudge/templates"
	"html/template"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/labstack/echo/v4"
	"github.com/mraron/njudge/internal/web/helpers/config"
	"github.com/mraron/njudge/internal/web/helpers/templates/partials"
	"github.com/mraron/njudge/pkg/problems"
)

type Renderer struct {
	templates     map[string]*template.Template
	cfg           config.Server
	problemStore  problems.Store
	db            *sql.DB
	partialsStore partials.Store

	sync.RWMutex
}

func New(cfg config.Server, problemStore problems.Store, db *sql.DB, partialsStore partials.Store) *Renderer {
	renderer := &Renderer{
		templates:     make(map[string]*template.Template),
		cfg:           cfg,
		problemStore:  problemStore,
		db:            db,
		partialsStore: partialsStore,
	}

	if err := renderer.Update(); err != nil {
		panic(err)
	}

	if cfg.Mode == "development" {
		renderer.startWatcher()
	}

	return renderer
}

func (t *Renderer) startWatcher() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		panic(err)
	}

	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}

				if event.Has(fsnotify.Write) {
					if err := t.Update(); err != nil {
						log.Println("error while parsing templates:", err)
					}
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("watcher error:", err)
			}
		}
	}()

	err = watcher.Add(t.cfg.TemplatesDir)
	if err != nil {
		panic(err)
	}
}

func (t *Renderer) Update() error {
	t.Lock()
	defer t.Unlock()

	layoutFiles := make([]string, 0)

	usedFS := os.DirFS(t.cfg.TemplatesDir)
	if t.cfg.Mode == "production" {
		usedFS = templates.FS
	}

	return fs.WalkDir(usedFS, ".", func(path string, info fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			if strings.HasPrefix(info.Name(), "_") {
				layoutFiles = append(layoutFiles, path)
			} else {
				t.templates[path] = template.Must(template.New(info.Name()).
					Funcs(contextFuncs(nil)).
					Funcs(statelessFuncs(t.problemStore, t.db, t.partialsStore)).
					ParseFS(usedFS, append(layoutFiles, path)...))
			}
		}

		return nil
	})
}

func (t *Renderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	t.RLock()
	defer t.RUnlock()

	if !strings.HasSuffix(name, ".gohtml") {
		name += ".gohtml"
	}

	if _, ok := t.templates[name]; !ok {
		return fmt.Errorf("can find template %q", name)
	}

	return t.templates[name].Funcs(contextFuncs(c)).ExecuteTemplate(w, filepath.Base(name), struct {
		Data    interface{}
		Context echo.Context
	}{data, c})
}
