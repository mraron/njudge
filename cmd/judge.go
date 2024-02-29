package cmd

import (
	"github.com/mraron/njudge/internal/judge2"
	"github.com/mraron/njudge/pkg/language"
	"github.com/mraron/njudge/pkg/language/sandbox"
	"github.com/mraron/njudge/pkg/problems"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log/slog"
	"time"
)

type JudgeConfig struct {
	Port        int
	ProblemsDir string

	Isolate             bool
	IsolateSandboxRange []int

	UpdateStatusLimitEvery time.Duration
}

var JudgeCmd = &cobra.Command{
	Use:   "judge",
	Short: "start judge server",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		viper.SetConfigFile("judge.yaml")
		viper.AddConfigPath(".")

		viper.SetDefault("port", 8080)
		viper.SetDefault("problemsDir", "/problems")
		viper.SetDefault("isolate", true)
		viper.SetDefault("isolateSandboxRange", []int{400, 444})
		viper.SetDefault("updateStatusLimitEvery", 5*time.Second)

		viper.AutomaticEnv()
		viper.SetEnvPrefix("njudge")

		return viper.MergeInConfig()
	},

	RunE: func(cmd *cobra.Command, args []string) error {
		cfg := JudgeConfig{}

		if err := viper.Unmarshal(&cfg); err != nil {
			return err
		}

		store := problems.NewFsStore(cfg.ProblemsDir)
		provider := sandbox.NewProvider()
		if cfg.Isolate {
			for i := cfg.IsolateSandboxRange[0]; i <= cfg.IsolateSandboxRange[1]; i++ {
				s, err := sandbox.NewIsolate(i)
				if err != nil {
					return err
				}
				provider.Put(s)
			}
		} else {
			for i := 0; i < 10; i++ {
				s, err := sandbox.NewDummy()
				if err != nil {
					return err
				}
				provider.Put(s)
			}
		}

		server := judge2.NewServer(
			slog.Default(),
			&judge2.Judge{
				SandboxProvider: provider,
				ProblemStore:    store,
				LanguageStore:   language.DefaultStore,
				RateLimit:       cfg.UpdateStatusLimitEvery,
			},
			store,
			judge2.WithPortServerOption(cfg.Port),
		)

		return server.Run()
	},
}

func init() {
	RootCmd.AddCommand(JudgeCmd)
}
