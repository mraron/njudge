package cmd

import (
	"errors"
	"fmt"
	"github.com/mraron/njudge/internal/judge"
	"github.com/mraron/njudge/pkg/language"
	"github.com/mraron/njudge/pkg/language/sandbox"
	"github.com/mraron/njudge/pkg/problems"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"io/fs"
	"log/slog"
	"strings"
	"time"
)

type JudgeConfig struct {
	Port        int
	ProblemsDir string

	Isolate             bool
	IsolateSandboxRange []int

	UpdateStatusLimitEvery time.Duration
}

var DefaultJudgeConfig = JudgeConfig{
	Port:                   8080,
	ProblemsDir:            "/problems",
	Isolate:                true,
	IsolateSandboxRange:    []int{400, 444},
	UpdateStatusLimitEvery: 5 * time.Second,
}

func NewJudgeCmd(v *viper.Viper) *cobra.Command {
	cfg := JudgeConfig{}
	cmd := &cobra.Command{
		Use:   "judge",
		Short: "start judge server",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			v.SetConfigFile("judge.yaml")
			v.AddConfigPath(".")

			v.SetDefault("port", DefaultJudgeConfig.Port)
			v.SetDefault("problemsDir", DefaultJudgeConfig.ProblemsDir)
			v.SetDefault("isolate", DefaultJudgeConfig.Isolate)
			v.SetDefault("isolateSandboxRange", DefaultJudgeConfig.IsolateSandboxRange)
			v.SetDefault("updateStatusLimitEvery", DefaultJudgeConfig.UpdateStatusLimitEvery)

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
				fmt.Println(configName, flag.Changed, v.IsSet(configName))
				if !flag.Changed && v.IsSet(configName) {
					val := v.Get(configName)
					_ = cmd.Flags().Set(flag.Name, fmt.Sprintf("%v", val))
				}
			})

			return nil
		},

		RunE: func(cmd *cobra.Command, args []string) error {
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

			server := judge.NewServer(
				slog.Default(),
				&judge.Judge{
					SandboxProvider: provider,
					ProblemStore:    store,
					LanguageStore:   language.DefaultStore,
					RateLimit:       cfg.UpdateStatusLimitEvery,
				},
				store,
				judge.WithPortServerOption(cfg.Port),
			)

			return server.Run()
		},
	}

	cmd.Flags().IntVar(&cfg.Port, "port", DefaultJudgeConfig.Port, "port to listen on")
	cmd.Flags().StringVar(&cfg.ProblemsDir, "problems-dir", DefaultJudgeConfig.ProblemsDir, "directory of the problems")
	cmd.Flags().BoolVar(&cfg.Isolate, "isolate", DefaultJudgeConfig.Isolate, "use isolate (otherwise dummy sandboxes are used which are NOT secure)")
	cmd.Flags().IntSliceVar(&cfg.IsolateSandboxRange, "isolate-sandbox-range", DefaultJudgeConfig.IsolateSandboxRange, "inclusive interval of isolate sandbox IDs")
	cmd.Flags().DurationVar(&cfg.UpdateStatusLimitEvery, "updateStatus-limit-every", DefaultJudgeConfig.UpdateStatusLimitEvery, "the rate of status updates for the clients")

	return cmd
}

func init() {
	RootCmd.AddCommand(NewJudgeCmd(viper.GetViper()))
}
