package memory

import (
	"context"
	"slices"
	"sync"

	"github.com/mraron/njudge/internal/njudge"
)

type Submissions struct {
	mutex  sync.Mutex
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
	m.mutex.Lock()
	defer m.mutex.Unlock()
	for ind := range m.data {
		if m.data[ind].ID == ID {
			res := m.data[ind]
			return &res, nil
		}
	}

	return nil, njudge.ErrorSubmissionNotFound
}
func (m *Submissions) GetAll(ctx context.Context) ([]njudge.Submission, error) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	res := make([]njudge.Submission, len(m.data))
	copy(res, m.data)

	return res, nil
}

func (m *Submissions) Insert(ctx context.Context, s njudge.Submission) (*njudge.Submission, error) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	s.ID = m.nextId
	m.nextId++

	m.data = append(m.data, s)

	res := m.data[len(m.data)-1]
	return &res, nil
}

func (m *Submissions) Delete(ctx context.Context, ID int) error {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	for ind := range m.data {
		if m.data[ind].ID == ID {
			m.data[ind] = m.data[len(m.data)-1]
			m.data = m.data[:len(m.data)-1]
			return nil
		}
	}

	return njudge.ErrorSubmissionNotFound
}

func (m *Submissions) Update(ctx context.Context, s njudge.Submission, fields []string) error {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	for ind := range m.data {
		if m.data[ind].ID == s.ID {
			if slices.Contains(fields, njudge.SubmissionFields.UserID) {
				m.data[ind].UserID = s.UserID
			}
			if slices.Contains(fields, njudge.SubmissionFields.ProblemID) {
				m.data[ind].ProblemID = s.ProblemID
			}
			if slices.Contains(fields, njudge.SubmissionFields.Language) {
				m.data[ind].Language = s.Language
			}
			if slices.Contains(fields, njudge.SubmissionFields.Source) {
				m.data[ind].Source = s.Source
			}
			if slices.Contains(fields, njudge.SubmissionFields.Private) {
				m.data[ind].Private = s.Private
			}
			if slices.Contains(fields, njudge.SubmissionFields.Started) {
				m.data[ind].Started = s.Started
			}
			if slices.Contains(fields, njudge.SubmissionFields.Verdict) {
				m.data[ind].Verdict = s.Verdict
			}
			if slices.Contains(fields, njudge.SubmissionFields.Ontest) {
				m.data[ind].Ontest = s.Ontest
			}
			if slices.Contains(fields, njudge.SubmissionFields.Submitted) {
				m.data[ind].Submitted = s.Submitted
			}
			if slices.Contains(fields, njudge.SubmissionFields.Status) {
				m.data[ind].Status = s.Status
			}
			if slices.Contains(fields, njudge.SubmissionFields.Judged) {
				m.data[ind].Judged = s.Judged
			}
			if slices.Contains(fields, njudge.SubmissionFields.Score) {
				m.data[ind].Score = s.Score
			}

			return nil
		}
	}
	return njudge.ErrorSubmissionNotFound
}
