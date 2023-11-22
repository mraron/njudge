package njudge_test

import (
	"context"
	"testing"

	"github.com/mraron/njudge/internal/njudge"
	"github.com/mraron/njudge/pkg/language"
	"github.com/mraron/njudge/pkg/problems"
	"github.com/stretchr/testify/assert"
)

func TestSubmissions(t *testing.T) {
	mp := njudge.NewMemoryProblems()
	p, _ := mp.Insert(context.TODO(), njudge.NewProblem("main", "aplusb"))

	mu := njudge.NewMemoryUsers()

	u, _ := njudge.NewUser("mraron", "asd@bsd", "admin")
	u, _ = mu.Insert(context.TODO(), *u)

	ms := njudge.NewMemorySubmissions()
	s, err := njudge.NewSubmission(*u, *p, language.DefaultStore.Get("cpp14"))
	assert.Nil(t, err)
	s, err = ms.Insert(context.TODO(), *s)
	assert.Nil(t, err)
	assert.Greater(t, s.ID, 0)

	s.SetSource([]byte("hehe"))
	s, _ = ms.Get(context.TODO(), 1)
	assert.Empty(t, s.Source)
	s.SetSource([]byte("hehe"))
	_ = ms.Update(context.TODO(), *s)
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
