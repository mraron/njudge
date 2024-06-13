package cmd

import (
	"errors"
	"fmt"
	"github.com/mraron/njudge/internal/web"
	"github.com/mraron/njudge/pkg/problems"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"io/fs"
	"strings"
)

var _dockerDatabaseConfig = web.DatabaseConfig{
	User:     "postgres",
	Password: "postgres",
	Host:     "db",
	Name:     "postgres",
	Port:     5432,
	SSLMode:  false,
}

type WebConfig struct {
	ProblemsDir       string `mapstructure:"problems_dir" yaml:"problems_dir"`
	DiscordWebhookURL string `mapstructure:"discord_webhook_url" yaml:"discord_webhook_url"`

	web.Config `mapstructure:",squash"`
}

var DefaultWebConfig = WebConfig{
	ProblemsDir: "/njudge_problems",
	Config: web.Config{
		Mode:           web.ModeDevelopment,
		Url:            "http://localhost:5555",
		Port:           "5555",
		TimeZone:       "Europe/Budapest",
		CookieSecret:   "svp3r_s3cr3t",
		GoogleAuth:     web.GoogleAuthConfig{},
		Sendgrid:       web.SendgridConfig{},
		SMTP:           web.SMTPConfig{},
		DatabaseConfig: _dockerDatabaseConfig,
	},
}

func NewWebCmd(v *viper.Viper) *cobra.Command {
	cfg := DefaultWebConfig
	cmd := &cobra.Command{
		Use:   "web",
		Short: "Run the web server",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			v.SetConfigFile("web.yaml")
			v.AddConfigPath(".")

			v.SetEnvPrefix("njudge")
			v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
			BindEnvs(WebConfig{})

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
			if err := cfg.Valid(); err != nil {
				return err
			}

			logger := getHookedLogger(cfg.DiscordWebhookURL)

			emailService := cfg.EmailService()
			db, err := cfg.ConnectAndPing(logger)
			if err != nil {
				return err
			}

			store := problems.NewFsStore(cfg.ProblemsDir)

			var dataAccess *web.DataAccess
			if cfg.Mode != web.ModeDemo {
				if dataAccess, err = web.NewDBDataAccess(cmd.Context(), store, db, emailService); err != nil {
					return err
				}
			} else {
				if dataAccess, err = web.NewDemoDataAccess(cmd.Context(), store, emailService); err != nil {
					return err
				}
			}
			s, err := web.NewServer(logger, cfg.Config, *dataAccess, db)
			if err != nil {
				return err
			}

			return s.Run(cmd.Context())
		},
	}

	cmd.Flags().StringVar(&cfg.Port, "port", DefaultWebConfig.Port, "port to listen on")
	cmd.Flags().StringVar(&cfg.ProblemsDir, "problems_dir", DefaultWebConfig.ProblemsDir, "directory of problems")

	cmd.Flags().StringVar(&cfg.DatabaseConfig.User, "db.user", DefaultWebConfig.DatabaseConfig.User, "database user")
	cmd.Flags().StringVar(&cfg.DatabaseConfig.Password, "db.password", DefaultWebConfig.DatabaseConfig.Password, "database password")
	cmd.Flags().StringVar(&cfg.DatabaseConfig.Host, "db.host", DefaultWebConfig.DatabaseConfig.Host, "database host")
	cmd.Flags().StringVar(&cfg.DatabaseConfig.Name, "db.name", DefaultWebConfig.DatabaseConfig.Name, "database name")
	cmd.Flags().IntVar(&cfg.DatabaseConfig.Port, "db.port", DefaultWebConfig.DatabaseConfig.Port, "database port")
	cmd.Flags().BoolVar(&cfg.DatabaseConfig.SSLMode, "db.ssl_mode", DefaultWebConfig.DatabaseConfig.SSLMode, "database sslmode")

	cmd.AddCommand(NewMigrateCommand(v))
	return cmd
}

/*
	var WebCmd = &cobra.Command{
		Use:   "web",
		Short: "manage web related parts, for example start webserver, run migrations etc.",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			BindEnvs(web.Config{})
			viper.SetEnvPrefix("njudge")

			viper.SetConfigName("web")
			viper.AddConfigPath(".")
			viper.AutomaticEnv()
			return viper.MergeInConfig()
		},

		RunE: func(cmd *cobra.Command, args []string) error {
			cfg := web.Config{}
			err := viper.Unmarshal(&cfg)
			if err != nil {
				return err
			}
			fmt.Println(cfg)
			return nil

			s := web.Server{Config: cfg}
			s.Run()

			return nil
		},
	}

	var SubmitCmdArgs struct {
		File       string
		Problemset string
		Problem    string
		Language   string
		User       int
	}

	var SubmitCmd = &cobra.Command{
		Use: "submit",
		RunE: func(cmd *cobra.Command, args []string) error {
			cfg := web.Config{}

			err := viper.Unmarshal(&cfg)
			if err != nil {
				return err
			}

			s := web.Server{Config: cfg}
			src, err := os.ReadFile(SubmitCmdArgs.File)
			if err != nil {
				return err
			}

			err = s.SetupEnvironment(cmd.Context())
			if err != nil {
				return err
			}

			web.NewDBDataAccess(cmd.Context())

			sub, err := s.SubmitService.Submit(cmd.Context(), njudge.SubmitRequest{
				UserID:     SubmitCmdArgs.User,
				Problemset: SubmitCmdArgs.Problem,
				Problem:    SubmitCmdArgs.Problem,
				Language:   SubmitCmdArgs.Language,
				Source:     src,
			})
			if err != nil {
				return err
			}

			log.Print("submission received with id ", sub.ID)
			return nil
		},
	}

	var ActivateCmdArgs struct {
		Name string
	}

	var ActivateCmd = &cobra.Command{
		Use: "activate",
		RunE: func(cmd *cobra.Command, args []string) error {
			cfg := web.Config{}

			err := viper.Unmarshal(&cfg)
			if err != nil {
				return err
			}

			s := web.Server{Config: cfg}

			s.SetupEnvironment()
			s.ConnectToDB()
			_, err = models.Users(qm.Where("name = ?", ActivateCmdArgs.Name)).UpdateAll(context.Background(), s.DB, models.M{"activation_key": nil})

			return err
		},
	}

	func RenameInDB(from, to string, tx *sql.Tx) error {
		log.Print("Renaming ", from, " to ", to)

		n, err := models.ProblemRels(qm.Where("problem = ?", from)).UpdateAll(context.Background(), tx, models.M{"problem": to})
		fmt.Println(n, err)
		if err != nil {
			return err
		}

		n, err = models.Submissions(qm.Where("problem = ?", from)).UpdateAll(context.Background(), tx, models.M{"problem": to})
		fmt.Println(n, err)
		return err
	}

	var PrefixCmdArgs struct {
		DryRun bool
		Reset  bool
	}

	var PrefixCmd = &cobra.Command{
		Use: "prefix_problems",
		RunE: func(cmd *cobra.Command, args []string) error {
			cfg := web.Config{}

			err := viper.Unmarshal(&cfg)
			if err != nil {
				return err
			}
			server := web.Server{Config: cfg}
			server.ConnectToDB()

			withoutPrefixes := problems.NewFsStore(cfg.ProblemsDir, problems.FsStoreIgnorePrefix())
			err = withoutPrefixes.UpdateProblems()
			if err != nil {
				return nil
			}

			withPrefixes := problems.NewFsStore(cfg.ProblemsDir)
			err = withPrefixes.UpdateProblems()
			if err != nil {
				return nil
			}

			if PrefixCmdArgs.Reset {
				withPrefixes, withoutPrefixes = withoutPrefixes, withPrefixes
			}

			tx, err := server.DB.Begin()
			if err != nil {
				return err
			}

			for path, id := range withPrefixes.ByPath {
				if withoutPrefixes.ByPath[path] != id {
					if err := RenameInDB(withoutPrefixes.ByPath[path], id, tx); err != nil {
						return err
					}
				}
			}

			if PrefixCmdArgs.DryRun {
				return tx.Rollback()
			}

			return tx.Commit()
		},
	}

	var PrefixRunCmd = &cobra.Command{
		Use: "run",
	}
*/
func init() {
	/*	RootCmd.AddCommand(WebCmd)

		SubmitCmd.Flags().IntVar(&SubmitCmdArgs.User, "user", 0, "ID of user on behalf we make the submission")
		SubmitCmd.Flags().StringVar(&SubmitCmdArgs.Problemset, "problemset", "main", "Problemset of problem")
		SubmitCmd.Flags().StringVar(&SubmitCmdArgs.Problem, "problem", "", "Problem")
		SubmitCmd.Flags().StringVar(&SubmitCmdArgs.Language, "language", "cpp14", "Language")
		SubmitCmd.Flags().StringVar(&SubmitCmdArgs.File, "file", "", "file to submit")

		SubmitCmd.MarkFlagRequired("user")
		SubmitCmd.MarkFlagRequired("problem")
		SubmitCmd.MarkFlagRequired("file")

		WebCmd.AddCommand(SubmitCmd)

		ActivateCmd.Flags().StringVar(&ActivateCmdArgs.Name, "name", "", "name os the user to activate")
		ActivateCmd.MarkFlagRequired("name")

		WebCmd.AddCommand(ActivateCmd)

		PrefixCmd.Flags().BoolVar(&PrefixCmdArgs.DryRun, "dry-run", false, "dry run")
		PrefixCmd.Flags().BoolVar(&PrefixCmdArgs.Reset, "reset", false, "reset prefixes")
		WebCmd.AddCommand(PrefixCmd)*/
}
