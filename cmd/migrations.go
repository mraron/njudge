package cmd

import (
	"log"
	"os"

	"github.com/mraron/njudge/pkg/web"
	"github.com/mraron/njudge/pkg/web/helpers/config"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type migrateLogger struct {
	*log.Logger
	verbose bool
}

func (m migrateLogger) Verbose() bool {
	return m.verbose
}

var MigrateCmdArgs struct {
	Up    bool
	Down  bool
	Steps int
}

var MigrateCmd = &cobra.Command{
	Use: "migrate",
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg := config.Server{}

		err := viper.Unmarshal(&cfg)
		if err != nil {
			return err
		}

		server := web.Server{Server: cfg}
		server.ConnectToDB()
		driver, err := postgres.WithInstance(server.DB.DB, &postgres.Config{})
		if err != nil {
			return err
		}

		m, err := migrate.NewWithDatabaseInstance("file://migrations", "postgres", driver)
		if err != nil {
			return err
		}

		m.Log = &migrateLogger{log.New(os.Stdout, "[migrate]", 0), false}

		if MigrateCmdArgs.Up {
			err = m.Up()
			if err != nil {
				return err
			}
		} else if MigrateCmdArgs.Down {
			err = m.Down()
			if err != nil {
				return err
			}
		} else if MigrateCmdArgs.Steps != 0 {
			err = m.Steps(MigrateCmdArgs.Steps)
			if err != nil {
				return err
			}
		}

		return nil
	},
}

func init() {
	MigrateCmd.Flags().BoolVar(&MigrateCmdArgs.Up, "up", false, "runs up migrations")
	MigrateCmd.Flags().BoolVar(&MigrateCmdArgs.Down, "down", false, "runs down migrations")
	MigrateCmd.Flags().IntVar(&MigrateCmdArgs.Steps, "steps", 0, "runs `x` up/down migrations depending on the positivity")

	WebCmd.AddCommand(MigrateCmd)
}
