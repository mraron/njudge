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
	"runtime"
	"strings"
	"time"
)

type JudgeConfig struct {
	Port        int    `mapstructure:"port" yaml:"port"`
	ProblemsDir string `mapstructure:"problems_dir" yaml:"problems_dir"`

	Isolate             bool  `mapstructure:"isolate" yaml:"isolate"`
	IsolateSandboxRange []int `mapstructure:"isolate_sandbox_range" yaml:"isolate_sandbox_range"`

	UpdateStatusLimitEvery time.Duration `mapstructure:"update_status_limit_every" yaml:"update_status_limit_every"`

	Concurrency int `mapstructure:"concurrency" yaml:"concurrency"`
}

var DefaultJudgeConfig = JudgeConfig{
	Port:                   8080,
	ProblemsDir:            "/njudge_problems",
	Isolate:                true,
	IsolateSandboxRange:    []int{400, 444},
	UpdateStatusLimitEvery: 5 * time.Second,
	Concurrency:            4,
}

func NewJudgeCmd(v *viper.Viper) *cobra.Command {
	cfg := DefaultJudgeConfig
	cmd := &cobra.Command{
		Use:   "judge",
		Short: "Run the judge server",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			v.SetConfigFile("judge.yaml")
			v.AddConfigPath(".")

			v.SetEnvPrefix("njudge")
			v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
			BindEnvs(JudgeConfig{})

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

			log := slog.Default()

			store := problems.NewFsStore(cfg.ProblemsDir)
			provider := sandbox.NewProvider()
			sandboxCount := 0
			if cfg.Isolate {
				for i := cfg.IsolateSandboxRange[0]; i <= cfg.IsolateSandboxRange[1]; i++ {
					s, err := sandbox.NewIsolate(i, sandbox.IsolateOptionUseLogger(slog.Default()))
					if err != nil {
						return err
					}
					provider.Put(s)
					sandboxCount++
				}
			} else {
				for i := 0; i < 10; i++ {
					s, err := sandbox.NewDummy(sandbox.DummyWithLogger(slog.Default()))
					if err != nil {
						return err
					}
					provider.Put(s)
					sandboxCount++
				}
			}

			if 2*cfg.Concurrency > sandboxCount {
				log.Warn("sandbox count is low for concurrency")
			}
			if cfg.Concurrency > runtime.GOMAXPROCS(0) {
				log.Warn("concurrency is higher than GOMAXPROCS")
			}

			tokens := make(chan struct{}, cfg.Concurrency)
			for range cfg.Concurrency {
				tokens <- struct{}{}
			}

			server := judge.NewServer(
				log,
				&judge.Judge{
					SandboxProvider: provider,
					ProblemStore:    store,
					LanguageStore:   language.DefaultStore,
					RateLimit:       cfg.UpdateStatusLimitEvery,
					Tokens:          tokens,
					Logger:          log,
				},
				store,
				judge.WithPortServerOption(cfg.Port),
			)

			return server.Run()
		},
	}

	cmd.Flags().IntVar(&cfg.Port, "port", DefaultJudgeConfig.Port, "Port to listen on")
	cmd.Flags().StringVar(&cfg.ProblemsDir, "problems_dir", DefaultJudgeConfig.ProblemsDir, "directory of problems")
	cmd.Flags().BoolVar(&cfg.Isolate, "isolate", DefaultJudgeConfig.Isolate, "use isolate (otherwise dummy sandboxes are used which are NOT secure)")
	cmd.Flags().IntSliceVar(&cfg.IsolateSandboxRange, "isolate_sandbox_range", DefaultJudgeConfig.IsolateSandboxRange, "inclusive interval of isolate sandbox IDs")
	cmd.Flags().DurationVar(&cfg.UpdateStatusLimitEvery, "updateStatus_limit_every", DefaultJudgeConfig.UpdateStatusLimitEvery, "the rate of status updates for the clients")
	cmd.Flags().IntVar(&cfg.Concurrency, "concurrency", DefaultJudgeConfig.Concurrency, "the maximum number of concurrently executed testcase")

	return cmd
}
