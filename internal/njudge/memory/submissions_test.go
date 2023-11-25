package memory_test

import (
	"context"
	"testing"

	"github.com/mraron/njudge/internal/njudge"
	"github.com/mraron/njudge/internal/njudge/memory"
	"github.com/mraron/njudge/pkg/language"
	"github.com/mraron/njudge/pkg/problems"
	"github.com/stretchr/testify/assert"
)

func TestSubmissions(t *testing.T) {
	var mp njudge.Problems = memory.NewProblems()
	p, _ := mp.Insert(context.TODO(), njudge.NewProblem("main", "aplusb"))

	mu := memory.NewUsers()

	u, _ := njudge.NewUser("mraron", "asd@bsd", "admin")
	u, _ = mu.Insert(context.TODO(), *u)

	var ms njudge.Submissions = memory.NewSubmissions()
	s, err := njudge.NewSubmission(*u, *p, language.DefaultStore.Get("cpp14"))
	assert.Nil(t, err)
	s, err = ms.Insert(context.TODO(), *s)
	assert.Nil(t, err)
	assert.Greater(t, s.ID, 0)

	s.SetSource([]byte("hehe"))
	s, _ = ms.Get(context.TODO(), 1)
	assert.Empty(t, s.Source)
	s.SetSource([]byte("hehe"))
	_ = ms.Update(context.TODO(), *s, njudge.Fields(njudge.SubmissionFields.Source))
	s, _ = ms.Get(context.TODO(), 1)
	assert.NotEmpty(t, s.Source)

	s.Judged.Valid = true
	s, _ = ms.Get(context.TODO(), 1)
	assert.False(t, s.Judged.Valid)

	s.Status.Feedback = append(s.Status.Feedback, problems.Testset{
		Name: "tests",
	})

	s, _ = ms.Get(context.TODO(), 1)
	assert.Empty(t, s.Status.Feedback)

	err = ms.Delete(context.TODO(), s.ID)
	assert.Nil(t, err)
}
