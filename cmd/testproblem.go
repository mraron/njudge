package cmd

import (
	"context"
	"fmt"
	"github.com/mraron/njudge/internal/judge"
	"github.com/mraron/njudge/pkg/language"
	"github.com/mraron/njudge/pkg/language/sandbox"
	"github.com/mraron/njudge/pkg/problems"
	"github.com/spf13/cobra"
	"io"
	"log/slog"
	"os"
	"runtime"
	"time"
)

type TestProblemConfig struct {
	ProblemDir   string
	Language     string
	SolutionPath string
	Verbose      bool

	IsolateSandboxRange []int
	Isolate             bool
	Concurrency         int
}

var DefaultTestProblemConfig = TestProblemConfig{
	ProblemDir:   ".",
	Language:     "cpp17",
	SolutionPath: "./sol/solution.cpp",
	Verbose:      false,

	IsolateSandboxRange: []int{10, 99},
	Isolate:             true,
	Concurrency:         4,
}

type ProblemStore struct {
	problem problems.Problem
}

func (p ProblemStore) GetProblem(s string) (problems.Problem, error) {
	return p.problem, nil
}

func NewTestProblemCmd() *cobra.Command {
	cfg := TestProblemConfig{}
	cmd := &cobra.Command{
		Use:   "test_problem",
		Short: "Test a problem with a solution",
		RunE: func(cmd *cobra.Command, args []string) error {
			log := slog.Default()
			if !cfg.Verbose {
				log = slog.New(slog.NewJSONHandler(io.Discard, nil))
			}

			p, err := problems.Parse(cfg.ProblemDir)
			if err != nil {
				return err
			}

			provider := sandbox.NewProvider()
			sandboxCount := 0
			if cfg.Isolate {
				for i := cfg.IsolateSandboxRange[0]; i <= cfg.IsolateSandboxRange[1]; i++ {
					s, err := sandbox.NewIsolate(i, sandbox.IsolateOptionUseLogger(log))
					if err != nil {
						return err
					}
					provider.Put(s)
					sandboxCount++
				}
			} else {
				for i := 0; i < 10; i++ {
					s, err := sandbox.NewDummy(sandbox.DummyWithLogger(log))
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

			j := &judge.Judge{
				SandboxProvider: provider,
				ProblemStore:    ProblemStore{problem: p},
				LanguageStore:   language.DefaultStore,
				RateLimit:       200 * time.Millisecond,
				Tokens:          tokens,
				Logger:          log,
			}
			src, err := os.ReadFile(cfg.SolutionPath)
			if err != nil {
				return err
			}

			res, err := j.Judge(context.Background(), judge.Submission{
				Source:   src,
				Language: cfg.Language,
			}, func(result judge.Result) error {
				fmt.Println(result.Status.Feedback[0].Verdict(), result.Status.Feedback[0].Score(), "/", result.Status.Feedback[0].MaxScore())
				return nil
			})
			if err != nil {
				return err
			}
			for _, test := range res.Feedback[0].Testcases() {
				fmt.Println(test.Group, test.Index, test.VerdictName, test.Score, test.MaxScore)
			}
			fmt.Println(res.Feedback[0].Verdict(), res.Feedback[0].Score(), "/", res.Feedback[0].MaxScore())
			fmt.Println(res.FeedbackType)
			return nil
		},
	}
	cmd.Flags().StringVar(&cfg.ProblemDir, "problem_dir", DefaultTestProblemConfig.ProblemDir, "problem directory")
	cmd.Flags().StringVar(&cfg.Language, "lang", DefaultTestProblemConfig.Language, "language")
	cmd.Flags().StringVar(&cfg.SolutionPath, "sol", DefaultTestProblemConfig.SolutionPath, "sol")
	cmd.Flags().BoolVar(&cfg.Verbose, "verbose", DefaultTestProblemConfig.Verbose, "verbose")
	cmd.Flags().BoolVar(&cfg.Isolate, "isolate", DefaultTestProblemConfig.Isolate, "isolate")
	cmd.Flags().IntSliceVar(&cfg.IsolateSandboxRange, "isolate_sandbox_range", DefaultTestProblemConfig.IsolateSandboxRange, "inclusive interval of isolate sandbox IDs")
	cmd.Flags().IntVar(&cfg.Concurrency, "concurrency", DefaultTestProblemConfig.Concurrency, "the maximum number of concurrently executed testcase")

	return cmd
}

func init() {
	RootCmd.AddCommand(NewTestProblemCmd())
}
