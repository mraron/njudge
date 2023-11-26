package memory

import (
	"context"
	"slices"
	"sync"

	"github.com/mraron/njudge/internal/njudge"
)

type Problems struct {
	sync.Mutex
	nextId int
	data   []njudge.Problem
}

func NewProblems() *Problems {
	return &Problems{
		nextId: 1,
		data:   make([]njudge.Problem, 0),
	}
}

func (m *Problems) GetByNames(ctx context.Context, problemset, problem string) (*njudge.Problem, error) {
	m.Lock()
	defer m.Unlock()
	for ind := range m.data {
		if m.data[ind].Problemset == problemset && m.data[ind].Problem == problem {
			res := m.data[ind]
			res.Tags = make([]njudge.ProblemTag, len(res.Tags))
			copy(res.Tags, m.data[ind].Tags)
			return &res, nil
		}
	}

	return nil, njudge.ErrorProblemNotFound
}

func (m *Problems) Get(ctx context.Context, ID int) (*njudge.Problem, error) {
	m.Lock()
	defer m.Unlock()
	for ind := range m.data {
		if m.data[ind].ID == ID {
			res := m.data[ind]
			res.Tags = make([]njudge.ProblemTag, len(res.Tags))
			copy(res.Tags, m.data[ind].Tags)
			return &res, nil
		}
	}

	return nil, njudge.ErrorProblemNotFound
}

func (m *Problems) GetAll(ctx context.Context) ([]njudge.Problem, error) {
	m.Lock()
	defer m.Unlock()
	res := make([]njudge.Problem, len(m.data))
	copy(res, m.data)
	for ind := range res {
		res[ind].Tags = make([]njudge.ProblemTag, len(res[ind].Tags))
		copy(res[ind].Tags, m.data[ind].Tags)
	}

	return res, nil
}

func (m *Problems) Insert(ctx context.Context, p njudge.Problem) (*njudge.Problem, error) {
	m.Lock()
	defer m.Unlock()
	p.ID = m.nextId
	m.nextId++

	m.data = append(m.data, p)

	res := m.data[len(m.data)-1]
	return &res, nil
}

func (m *Problems) Delete(ctx context.Context, id int) error {
	m.Lock()
	defer m.Unlock()
	for ind := range m.data {
		if m.data[ind].ID == id {
			m.data[ind] = m.data[len(m.data)-1]
			m.data = m.data[:len(m.data)-1]
			return nil
		}
	}

	return njudge.ErrorProblemNotFound
}

func (m *Problems) Update(ctx context.Context, p njudge.Problem, fields []string) error {
	m.Lock()
	defer m.Unlock()
	for ind := range m.data {
		if m.data[ind].ID == p.ID {
			if slices.Contains(fields, njudge.ProblemFields.Tags) {
				m.data[ind].Tags = p.Tags
			}

			if slices.Contains(fields, njudge.ProblemFields.Problemset) {
				m.data[ind].Problemset = p.Problemset
			}
			if slices.Contains(fields, njudge.ProblemFields.Problem) {
				m.data[ind].Problem = p.Problem
			}
			if slices.Contains(fields, njudge.ProblemFields.SolverCount) {
				m.data[ind].SolverCount = p.SolverCount
			}
			if slices.Contains(fields, njudge.ProblemFields.Category) {
				m.data[ind].Category = p.Category
			}

			return nil
		}
	}
	return njudge.ErrorProblemNotFound
}
