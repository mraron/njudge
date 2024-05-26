package cmd

import (
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	_ "github.com/lib/pq"
	"github.com/mraron/njudge/internal/njudge/db/migrations"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"log/slog"
	"os"
)

type migrateLogger struct {
	*log.Logger
	verbose bool
}

func (m migrateLogger) Verbose() bool {
	return m.verbose
}

type MigrateCmdArgs struct {
	Up    bool
	Down  bool
	Steps int
}

func NewMigrateCommand(v *viper.Viper) *cobra.Command {
	migrateArgs := MigrateCmdArgs{}
	cmd := &cobra.Command{
		Use: "migrate",
		RunE: func(cmd *cobra.Command, args []string) error {
			cfg := DefaultWebConfig

			err := viper.Unmarshal(&cfg)
			if err != nil {
				return err
			}

			db, err := cfg.DatabaseConfig.ConnectAndPing(slog.Default())
			if err != nil {
				return err
			}
			driver, err := postgres.WithInstance(db, &postgres.Config{})
			if err != nil {
				return err
			}

			d, err := iofs.New(migrations.FS, ".")
			if err != nil {
				return err
			}

			m, err := migrate.NewWithInstance("iofs", d, "postgres", driver)
			if err != nil {
				return err
			}

			m.Log = &migrateLogger{log.New(os.Stdout, "[migrate]", 0), false}

			if migrateArgs.Up {
				err = m.Up()
				if err != nil {
					return err
				}
			} else if migrateArgs.Down {
				if !askForConfirmation("This might DESTROY your data! Are you sure you want to down migrate?") {
					return nil
				}
				err = m.Down()
				if err != nil {
					return err
				}
			} else if migrateArgs.Steps != 0 {
				if !askForConfirmation(fmt.Sprintf("This might DESTROY your data! Are you sure you want to migrate %d steps?", migrateArgs.Steps)) {
					return nil
				}
				err = m.Steps(migrateArgs.Steps)
				if err != nil {
					return err
				}
			}

			return nil
		},
	}

	cmd.Flags().BoolVar(&migrateArgs.Up, "up", false, "runs up migrations")
	cmd.Flags().BoolVar(&migrateArgs.Down, "down", false, "runs down migrations")
	cmd.Flags().IntVar(&migrateArgs.Steps, "steps", 0, "runs `x` up/down migrations depending on the positivity")
	return cmd
}
