package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/mraron/njudge/internal/judge"

	"github.com/mraron/njudge/pkg/language"
	"github.com/mraron/njudge/pkg/language/sandbox"
	"github.com/mraron/njudge/pkg/problems"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var TestProblemArgs struct {
	ProblemDir   string
	Language     string
	SolutionPath string
	Verbose      int
}

var TestProblemCmd = &cobra.Command{
	Use: "testproblem",
	RunE: func(cmd *cobra.Command, args []string) error {
		sp := sandbox.NewProvider()
		s1, _ := sandbox.NewIsolate(50)
		sp.Put(s1)
		s2, _ := sandbox.NewIsolate(51)
		sp.Put(s2)
		w := judge.NewWorker(1, sp)

		logger, err := zap.NewDevelopment()
		if err != nil {
			return err
		}
		p, err := problems.Parse(TestProblemArgs.ProblemDir)
		if err != nil {
			return err
		}

		l, _ := language.DefaultStore.Get(TestProblemArgs.Language)
		src, err := os.ReadFile(TestProblemArgs.SolutionPath)
		if err != nil {
			return err
		}

		st, err := w.Judge(context.Background(), logger, p, src, l, judge.NewWriterCallback(io.Discard, func() {}))
		for _, test := range st.Feedback[0].Testcases() {
			fmt.Println(test.Group, test.Index, test.VerdictName, test.Score, test.MaxScore)
		}
		fmt.Println(st.Feedback[0].Verdict(), st.Feedback[0].Score(), "/", st.Feedback[0].MaxScore())
		fmt.Println(st.FeedbackType)

		if TestProblemArgs.Verbose > 0 {
			res, err := json.Marshal(st)
			if err != nil {
				return err
			}

			fmt.Println(string(res))
		}

		return err
	},
}

func init() {
	TestProblemCmd.Flags().StringVar(&TestProblemArgs.ProblemDir, "problem", "", "problemdir")
	TestProblemCmd.Flags().StringVar(&TestProblemArgs.Language, "lang", "", "language")
	TestProblemCmd.Flags().StringVar(&TestProblemArgs.SolutionPath, "sol", "", "solutionpath")
	TestProblemCmd.Flags().CountVarP(&TestProblemArgs.Verbose, "verbose", "v", "verbose")

	TestProblemCmd.MarkFlagFilename("sol")
	TestProblemCmd.MarkFlagDirname("problem")

	TestProblemCmd.MarkFlagRequired("problem")
	TestProblemCmd.MarkFlagRequired("lang")
	TestProblemCmd.MarkFlagRequired("sol")

	RootCmd.AddCommand(TestProblemCmd)
}
