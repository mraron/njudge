package cmd

import (
	"fmt"
	"log/slog"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"github.com/mraron/njudge/internal/njudge/db/migrations"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type MigrateCmdArgs struct {
	Up           bool
	Down         bool
	Steps        int
	ForceVersion int
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

			m, err := migrations.NewMigrate(db, true)
			if err != nil {
				return err
			}
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
			} else if migrateArgs.ForceVersion >= 0 {
				err = m.Force(migrateArgs.ForceVersion)
				if err != nil {
					return err
				}
			} else {
				v, dirty, err := m.Version()
				fmt.Println("version:", v, "dirty:", dirty, "err:", err)
			}

			return nil
		},
	}

	cmd.Flags().BoolVar(&migrateArgs.Up, "up", false, "runs up migrations")
	cmd.Flags().BoolVar(&migrateArgs.Down, "down", false, "runs down migrations")
	cmd.Flags().IntVar(&migrateArgs.Steps, "steps", 0, "runs `x` up/down migrations depending on the positivity")
	cmd.Flags().IntVar(&migrateArgs.ForceVersion, "force", -1, "forces version in the db and resets dirty")
	return cmd
}
