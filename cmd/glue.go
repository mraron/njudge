package cmd

import (
	"errors"
	"fmt"
	"github.com/mraron/njudge/internal/glue"
	"github.com/mraron/njudge/internal/web"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"golang.org/x/net/context"
	"io/fs"
	"log/slog"
	"strings"
	"time"
)

type GlueConfig struct {
	Database web.DatabaseConfig `mapstructure:"db"`
}

var DefaultGlueConfig = GlueConfig{
	Database: _dockerDatabaseConfig,
}

func NewGlueCmd(v *viper.Viper) *cobra.Command {
	cfg := DefaultGlueConfig
	cmd := &cobra.Command{
		Use:   "glue",
		Short: "start glue service",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			v.SetConfigType("yaml")
			v.SetConfigFile("glue.yaml")
			v.AddConfigPath(".")

			v.SetEnvPrefix("njudge")
			v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
			BindEnvs(GlueConfig{})

			if err := v.ReadInConfig(); err != nil {
				var res *fs.PathError
				if !errors.As(err, &res) {
					return err
				}
			}

			cmd.Flags().VisitAll(func(flag *pflag.Flag) {
				configName := flag.Name
				if !flag.Changed && v.IsSet(configName) {
					val := v.Get(configName)
					_ = cmd.Flags().Set(flag.Name, fmt.Sprintf("%v", val))
				}
			})

			return nil
		},

		RunE: func(cmd *cobra.Command, args []string) error {
			if err := v.Unmarshal(&cfg); err != nil {
				return err
			}

			conn, err := cfg.Database.ConnectAndPing(slog.Default())
			if err != nil {
				return err
			}
			judges := glue.NewJudges(conn, slog.Default())
			go func() {
				for {
					judges.Update(context.Background())
					time.Sleep(10 * time.Second)
				}
			}()
			g, err := glue.New(
				judges,
				glue.WithDatabaseOption(conn),
				glue.WithLogger(slog.Default()),
			)
			if err != nil {
				return err
			}
			g.Start(cmd.Context())
			return nil
		},
	}

	cmd.Flags().StringVar(&cfg.Database.User, "db.user", DefaultGlueConfig.Database.User, "database user")
	cmd.Flags().StringVar(&cfg.Database.Password, "db.password", DefaultGlueConfig.Database.Password, "database password")
	cmd.Flags().StringVar(&cfg.Database.Host, "db.host", DefaultGlueConfig.Database.Host, "database host")
	cmd.Flags().StringVar(&cfg.Database.Name, "db.name", DefaultGlueConfig.Database.Name, "database name")
	cmd.Flags().IntVar(&cfg.Database.Port, "db.port", DefaultGlueConfig.Database.Port, "database port")
	cmd.Flags().BoolVar(&cfg.Database.SSLMode, "db.ssl_mode", DefaultGlueConfig.Database.SSLMode, "database sslmode")

	return cmd
}

func init() {
	RootCmd.AddCommand(NewGlueCmd(viper.GetViper()))
}
