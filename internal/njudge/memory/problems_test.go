package memory_test

import (
	"context"
	"testing"

	"github.com/mraron/njudge/internal/njudge"
	"github.com/mraron/njudge/internal/njudge/memory"
	"github.com/stretchr/testify/assert"
)

func TestMemoryProblems(t *testing.T) {
	m := memory.NewProblems()
	p, err := m.Insert(context.TODO(), njudge.NewProblem("main", "aplusb"))
	assert.NotNil(t, p)
	currID := p.ID
	assert.Greater(t, currID, 0)
	assert.Nil(t, p.Category)
	assert.Nil(t, err)

	p.SetCategory(njudge.Category{ID: 5})
	err = m.Update(context.TODO(), *p)
	assert.Nil(t, err)

	p, err = m.Get(context.TODO(), currID)
	assert.NotNil(t, p)
	assert.Equal(t, 5, p.Category.ID)
	assert.Equal(t, currID, p.ID)
	assert.Equal(t, "aplusb", p.Problem)
	assert.Equal(t, "main", p.Problemset)
	assert.Nil(t, err)

	err = m.Delete(context.TODO(), currID)
	assert.Nil(t, err)

	err = m.Delete(context.TODO(), currID)
	assert.EqualError(t, err, njudge.ErrorProblemNotFound.Error())

	m.Insert(context.TODO(), njudge.NewProblem("main", "a"))
	m.Insert(context.TODO(), njudge.NewProblem("main", "b"))
	m.Insert(context.TODO(), njudge.NewProblem("main", "c"))

	ps, _ := m.GetAll(context.TODO())
	prevProblem := ps[0].Problem
	ps[0].Problem = "lulz"

	p, _ = m.Get(context.TODO(), ps[0].ID)
	assert.Equal(t, prevProblem, p.Problem)
}

func TestProblemTags(t *testing.T) {
	m := memory.NewProblems()
	prob := njudge.NewProblem("main", "aplusb")
	prob.AddTag(njudge.Tag{ID: 6, Name: "dp"}, 1)

	p, err := m.Insert(context.TODO(), prob)
	assert.Nil(t, err)
	p.AddTag(njudge.Tag{ID: 5, Name: "greedy"}, 1)
	err = m.Update(context.TODO(), *p)
	assert.Nil(t, err)

	res, err := m.Get(context.TODO(), 1)
	assert.Nil(t, err)
	assert.Equal(t, len(res.Tags), 2)
	assert.Equal(t, res.Tags[0].ID, 1)
	assert.Equal(t, res.Tags[1].ID, 2)

	res.DeleteTag(njudge.Tag{ID: 6, Name: "dp"})
	assert.Nil(t, m.Update(context.TODO(), *res))

	res, err = m.Get(context.TODO(), 1)
	assert.Nil(t, err)
	assert.Equal(t, len(res.Tags), 1)
	assert.Equal(t, res.Tags[0].ID, 2)
}
