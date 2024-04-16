package judge

import (
	"context"
	"github.com/mraron/njudge/pkg/language"
	"github.com/mraron/njudge/pkg/language/sandbox"
	"github.com/mraron/njudge/pkg/problems"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestJudge_Judge(t *testing.T) {
	s1, _ := sandbox.NewDummy()
	s2, _ := sandbox.NewDummy()
	store := problems.NewFsStore("testdata")
	_ = store.UpdateProblems()

	judge := Judge{
		SandboxProvider: sandbox.NewProvider().Put(s1).Put(s2),
		ProblemStore:    store,
		LanguageStore:   language.DefaultStore,
		RateLimit:       0,
	}

	assertProgression := func() ResultCallback {
		prev := 0
		count := 10000
		return func(result Result) error {
			assert.Greater(t, result.Index, prev)
			drs := 0
			for _, tc := range result.Status.Feedback[0].Testcases() {
				if tc.VerdictName == problems.VerdictDR {
					drs++
				}
			}
			assert.Less(t, drs, count)
			prev = result.Index
			count = drs
			return nil
		}
	}

	assertVerdicts := func(status *problems.Status, verdicts []problems.VerdictName) {
		testcases := status.Feedback[0].Testcases()
		assert.Equal(t, len(testcases), len(verdicts))
		for i := 0; i < len(verdicts); i++ {
			assert.Equal(t, verdicts[i], testcases[i].VerdictName)
		}
	}

	res, err := judge.Judge(context.Background(), Submission{
		ID:       "",
		Problem:  "aplusb",
		Language: "python3",
		Source: []byte(`
print("hello world!")
`),
	}, assertProgression())
	assert.NoError(t, err)
	assertVerdicts(res, []problems.VerdictName{problems.VerdictWA, problems.VerdictWA, problems.VerdictWA})

	res, err = judge.Judge(context.Background(), Submission{
		ID:       "",
		Problem:  "aplusb",
		Language: "python3",
		Source: []byte(`a, b = list(map(int, input().split()))
print(a+a)
`),
	}, assertProgression())
	assert.NoError(t, err)
	assertVerdicts(res, []problems.VerdictName{problems.VerdictWA, problems.VerdictAC, problems.VerdictWA})

	res, err = judge.Judge(context.Background(), Submission{
		ID:       "",
		Problem:  "aplusb",
		Language: "python3",
		Source: []byte(`a, b = list(map(int, input().split()))
print(a+b)
`),
	}, assertProgression())
	assert.NoError(t, err)
	assertVerdicts(res, []problems.VerdictName{problems.VerdictAC, problems.VerdictAC, problems.VerdictAC})

	res, err = judge.Judge(context.Background(), Submission{
		ID:       "",
		Problem:  "aplusb2",
		Language: "python3",
		Source: []byte(`a, b = list(map(int, input().split()))
print(a+b)
`),
	}, nil)
	assert.Nil(t, res)
	assert.ErrorIs(t, err, problems.ErrorProblemNotFound)

	res, err = judge.Judge(context.Background(), Submission{
		ID:       "",
		Problem:  "aplusb",
		Language: "python5",
		Source: []byte(`a, b = list(map(int, input().split()))
print(a+b)
`),
	}, nil)
	assert.Nil(t, res)
	assert.ErrorIs(t, err, language.ErrorLanguageNotFound)

	res, err = judge.Judge(context.Background(), Submission{
		ID:       "",
		Problem:  "aplusb",
		Language: "cpp14",
		Source: []byte(`int main() {}
`),
	}, func(result Result) error {
		return nil
	})
	assert.NoError(t, err)
	assert.NotNil(t, res)
	assert.True(t, res.Compiled)

}
