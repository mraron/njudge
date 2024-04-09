package cmd

import (
	"errors"
	"fmt"
	"github.com/mraron/njudge/internal/glue"
	"github.com/mraron/njudge/internal/judge"
	"github.com/mraron/njudge/internal/web/helpers/config"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"io/fs"
	"strings"
)

type GlueConfig struct {
	Database config.Database
}

var DefaultGlueConfig = GlueConfig{
	Database: config.Database{
		User:     "postgres",
		Password: "postgres",
		Host:     "db",
		Name:     "postgres",
		Port:     5432,
		SSLMode:  true,
	},
}

func NewGlueCmd(v *viper.Viper) *cobra.Command {
	cfg := GlueConfig{}
	cmd := &cobra.Command{
		Use:   "glue",
		Short: "start glue service",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			v.SetConfigFile("glue.yaml")
			v.AddConfigPath(".")

			v.SetDefault("db.user", DefaultGlueConfig.Database.User)
			v.SetDefault("db.password", DefaultGlueConfig.Database.Password)
			v.SetDefault("db.host", DefaultGlueConfig.Database.Host)
			v.SetDefault("db.name", DefaultGlueConfig.Database.Name)
			v.SetDefault("db.port", DefaultGlueConfig.Database.Port)
			v.SetDefault("db.sslmode", DefaultGlueConfig.Database.SSLMode)

			v.AutomaticEnv()
			v.SetEnvPrefix("njudge")

			if err := v.ReadInConfig(); err != nil {
				var res *fs.PathError
				if !errors.As(err, &res) {
					return err
				}
			}

			cmd.Flags().VisitAll(func(flag *pflag.Flag) {
				configName := strings.ReplaceAll(flag.Name, "-", "")
				if !flag.Changed && v.IsSet(configName) {
					val := v.Get(configName)
					_ = cmd.Flags().Set(flag.Name, fmt.Sprintf("%v", val))
				}
			})

			return nil
		},

		RunE: func(cmd *cobra.Command, args []string) error {
			g, err := glue.New(judge2.NewClient("http://localhost:8888"), glue.WithDatabaseOption(cfg.Database))
			if err != nil {
				return err
			}
			g.Start(cmd.Context())
			return nil
		},
	}

	cmd.Flags().StringVar(&cfg.Database.User, "db.user", DefaultGlueConfig.Database.User, "database user")
	cmd.Flags().StringVar(&cfg.Database.Password, "db.password", DefaultGlueConfig.Database.Password, "database password")
	cmd.Flags().StringVar(&cfg.Database.Host, "db.host", DefaultGlueConfig.Database.Password, "database host")
	cmd.Flags().StringVar(&cfg.Database.Name, "db.name", DefaultGlueConfig.Database.Name, "database name")
	cmd.Flags().IntVar(&cfg.Database.Port, "db.port", DefaultGlueConfig.Database.Port, "database port")
	cmd.Flags().BoolVar(&cfg.Database.SSLMode, "db.sslmode", DefaultGlueConfig.Database.SSLMode, "database sslmode")

	return cmd
}

func init() {
	RootCmd.AddCommand(NewGlueCmd(viper.GetViper()))
}
