package evaluation

import (
	"context"
	"github.com/mraron/njudge/pkg/language/sandbox"
	"github.com/mraron/njudge/pkg/problems"
	"strconv"
)

type LinearEvaluator struct {
	Runner problems.Runner
}

func NewLinearEvaluator(runner problems.Runner) *LinearEvaluator {
	return &LinearEvaluator{
		Runner: runner,
	}
}

func DeepCopyStatus(skeleton problems.Status) (ans problems.Status) {
	ans.CompilerOutput = skeleton.CompilerOutput
	ans.Compiled = skeleton.Compiled
	ans.FeedbackType = skeleton.FeedbackType
	for _, ts := range skeleton.Feedback {
		ans.Feedback = append(ans.Feedback, problems.Testset{Name: ts.Name})
		testset := &ans.Feedback[len(ans.Feedback)-1]

		for _, g := range ts.Groups {
			testset.Groups = append(testset.Groups, problems.Group{Name: g.Name, Scoring: g.Scoring})

			group := &testset.Groups[len(testset.Groups)-1]
			group.Testcases = append(group.Testcases, g.Testcases...)
		}
	}
	return
}

func (le *LinearEvaluator) Evaluate(ctx context.Context, skeleton problems.Status, compiledSolution problems.Solution, sandboxProvider sandbox.Provider, statusUpdater problems.StatusUpdater) (problems.Status, error) {
	var (
		ans problems.Status
		err error
	)

	if err = le.Runner.SetSolution(ctx, compiledSolution); err != nil {
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

	ans = DeepCopyStatus(skeleton)
	ans.Compiled = true
	ans.FeedbackType = skeleton.FeedbackType

	for tsInd := range ans.Feedback {
		testset := &ans.Feedback[tsInd]

		for gInd := range testset.Groups {
			group, currAC := &testset.Groups[gInd], true

			for tcInd := range group.Testcases {
				if ctx.Err() != nil {
					return ans, ctx.Err()
				}

				tc := &group.Testcases[tcInd]
				if tc.VerdictName != problems.VerdictDR {
					continue
				}

				if ans.FeedbackType == problems.FeedbackLazyIOI && !currAC {
					continue
				}

				if dependenciesOK(group.Dependencies) {
					if err := statusUpdater.UpdateStatus(ctx, strconv.Itoa(tc.Index), ans); err != nil {
						return ans, err
					}

					if err := le.Runner.Run(ctx, sandboxProvider, tc); err != nil {
						tc.VerdictName = problems.VerdictXX
						return ans, err
					} else if tc.VerdictName != problems.VerdictAC {
						currAC = false
						continue
					}
				} else {
					currAC = false
					continue
				}
			}

			groupAC[group.Name] = currAC
		}
	}

	return ans, nil
}
