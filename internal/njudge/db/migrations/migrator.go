package migrations

import (
	"database/sql"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/golang-migrate/migrate/v4/source/fn"
	_ "github.com/lib/pq"
	"log"
	"os"
)

type migrateLogger struct {
	*log.Logger
	verbose bool
}

func (m migrateLogger) Verbose() bool {
	return m.verbose
}

func NewMigrate(s *sql.DB, verbose bool) (*migrate.Migrate, error) {
	driver, err := postgres.WithInstance(s, &postgres.Config{})
	if err != nil {
		return nil, err
	}

	funcMigrations := map[string]*fn.Migration{
		"13_update_submission_status_to_base64": {
			Up:   source.ExecutorFunc(up13),
			Down: source.ExecutorFunc(down13),
		},
	}
	elems, err := FS.ReadDir(".")
	if err != nil {
		return nil, err
	}
	for _, e := range elems {
		if e.IsDir() {
			continue
		}
		m, err := source.Parse(e.Name())
		if err != nil {
			return nil, err
		}
		id := fmt.Sprintf("%d_%s", m.Version, m.Identifier)
		if _, ok := funcMigrations[id]; !ok {
			funcMigrations[id] = &fn.Migration{}
		}
		fun := source.ExecutorFunc(func(i interface{}) error {
			f, err := FS.Open(e.Name())
			if err != nil {
				return err
			}
			return driver.Run(f)
		})
		if m.Direction == source.Up {
			funcMigrations[id].Up = fun
		} else {
			funcMigrations[id].Down = fun
		}
	}

	d, err := fn.WithInstance(funcMigrations)
	if err != nil {
		return nil, err
	}

	m, err := migrate.NewWithInstance("fn", d, "postgres", driver)
	if err != nil {
		return nil, err
	}
	m.Log = &migrateLogger{log.New(os.Stderr, "[migrate]", 0), verbose}

	return m, nil
}
