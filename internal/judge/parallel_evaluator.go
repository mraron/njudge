package judge

import (
	"context"
	"errors"
	"github.com/mraron/njudge/pkg/language/sandbox"
	"github.com/mraron/njudge/pkg/problems"
	"github.com/mraron/njudge/pkg/problems/evaluation"
	"golang.org/x/sync/errgroup"
	"log/slog"
	"strconv"
)

type ParallelEvaluator struct {
	Runner problems.Runner
	Tokens chan struct{}
	Logger *slog.Logger
}

func (pe *ParallelEvaluator) Evaluate(ctx context.Context, skeleton problems.Status, compiledSolution problems.Solution, sandboxProvider sandbox.Provider, statusUpdater problems.StatusUpdater) (problems.Status, error) {
	type statusUpdateInfo struct {
		tc    problems.Testcase
		tcInd int
	}

	var (
		ans problems.Status
		err error
	)

	if err = pe.Runner.SetSolution(ctx, compiledSolution); err != nil {
		return ans, err
	}

	defer func() {
		_ = statusUpdater.Done(ctx)
	}()

	groupAC := make(map[string]bool)
	dependenciesOK := func(deps []string) bool {
		for _, dep := range deps {
			if !groupAC[dep] {
				return false
			}
		}
		return true
	}

	ans = evaluation.DeepCopyStatus(skeleton)
	ans.Compiled = true
	ans.FeedbackType = skeleton.FeedbackType

	for tsInd := range skeleton.Feedback {
		testset := &ans.Feedback[tsInd]
		pe.Logger.Info("⬜\tstarted testset", "testset", testset.Name)

		for gInd := range testset.Groups {
			group := &testset.Groups[gInd]
			pe.Logger.Info("◻️\tstarted group", "group", group.Name)

			groupCtx, groupCancel := context.WithCancel(ctx)
			executionGroup, _ := errgroup.WithContext(groupCtx)

			testcaseChan := make(chan statusUpdateInfo)
			groupDone := make(chan struct{})
			var updateStatusError error
			go func() {
				mxInd := -1
				for testcase := range testcaseChan {
					ans.Feedback[tsInd].Groups[gInd].Testcases[testcase.tcInd] = testcase.tc
					if testcase.tc.Index > mxInd {
						mxInd = testcase.tc.Index
					}

					if err := statusUpdater.UpdateStatus(ctx, strconv.Itoa(mxInd), ans); err != nil {
						updateStatusError = errors.Join(updateStatusError, err)
					}
				}
				groupDone <- struct{}{}
			}()

			for tcInd := range group.Testcases {
				tcInd := tcInd
				executionGroup.Go(func() error {
					<-pe.Tokens
					pe.Logger.Info("▫️\tstarted test", "testcase_ind", tcInd)
					defer func() {
						pe.Tokens <- struct{}{}
					}()

					if groupCtx.Err() != nil {
						return groupCtx.Err()
					}

					tc := group.Testcases[tcInd]
					if tc.VerdictName != problems.VerdictDR {
						return nil
					}

					if dependenciesOK(group.Dependencies) {
						err := pe.Runner.Run(ctx, sandboxProvider, &tc)
						if err != nil {
							tc.VerdictName = problems.VerdictXX
							testcaseChan <- statusUpdateInfo{tc: tc, tcInd: tcInd}
							return err
						} else if tc.VerdictName != problems.VerdictAC && tc.VerdictName != problems.VerdictPC {
							if ans.FeedbackType == problems.FeedbackLazyIOI || ans.FeedbackType == problems.FeedbackACM {
								groupCancel()
							}
						}
						testcaseChan <- statusUpdateInfo{tc: tc, tcInd: tcInd}
					} else {
						groupCancel()
					}
					return nil
				})
			}

			groupAC[group.Name] = groupCtx.Err() == nil
			if err := executionGroup.Wait(); err != nil && !errors.Is(err, context.Canceled) {
				groupCancel()
				close(testcaseChan)
				<-groupDone
				return ans, err
			}
			groupCancel()
			close(testcaseChan)
			<-groupDone
			if updateStatusError != nil {
				return ans, updateStatusError
			}
		}
	}

	return ans, nil
}
