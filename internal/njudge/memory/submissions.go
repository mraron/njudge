package memory

import (
	"context"
	"sync"

	"github.com/mraron/njudge/internal/njudge"
)

type Submissions struct {
	sync.Mutex
	nextId int
	data   []njudge.Submission
}

func NewSubmissions() *Submissions {
	return &Submissions{
		nextId: 1,
		data:   make([]njudge.Submission, 0),
	}
}

func (m *Submissions) Get(ctx context.Context, ID int) (*njudge.Submission, error) {
	m.Lock()
	defer m.Unlock()
	for ind := range m.data {
		if m.data[ind].ID == ID {
			res := m.data[ind]
			return &res, nil
		}
	}

	return nil, njudge.ErrorSubmissionNotFound
}
func (m *Submissions) GetAll(ctx context.Context) ([]njudge.Submission, error) {
	m.Lock()
	defer m.Unlock()
	res := make([]njudge.Submission, len(m.data))
	copy(res, m.data)

	return res, nil
}

func (m *Submissions) Insert(ctx context.Context, s njudge.Submission) (*njudge.Submission, error) {
	m.Lock()
	defer m.Unlock()
	s.ID = m.nextId
	m.nextId++

	m.data = append(m.data, s)

	res := m.data[len(m.data)-1]
	return &res, nil
}

func (m *Submissions) Delete(ctx context.Context, ID int) error {
	m.Lock()
	defer m.Unlock()
	for ind := range m.data {
		if m.data[ind].ID == ID {
			m.data[ind] = m.data[len(m.data)-1]
			m.data = m.data[:len(m.data)-1]
			return nil
		}
	}

	return njudge.ErrorSubmissionNotFound
}

func (m *Submissions) Update(ctx context.Context, s njudge.Submission) error {
	m.Lock()
	defer m.Unlock()
	for ind := range m.data {
		if m.data[ind].ID == s.ID {
			m.data[ind] = s
			return nil
		}
	}
	return njudge.ErrorSubmissionNotFound
}
