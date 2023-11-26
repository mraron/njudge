package memory

import (
	"context"
	"sync"

	"github.com/mraron/njudge/internal/njudge"
)

type Categories struct {
	sync.Mutex
	nextId int
	data   []njudge.Category
}

func NewCategories() *Categories {
	return &Categories{
		nextId: 1,
		data:   make([]njudge.Category, 0),
	}
}

func (m *Categories) Get(ctx context.Context, ID int) (*njudge.Category, error) {
	m.Lock()
	defer m.Unlock()
	for ind := range m.data {
		if m.data[ind].ID == ID {
			res := m.data[ind]
			return &res, nil
		}
	}

	return nil, njudge.ErrorCategoryNotFound
}

func (m *Categories) GetAll(ctx context.Context) ([]njudge.Category, error) {
	m.Lock()
	defer m.Unlock()
	res := make([]njudge.Category, len(m.data))
	copy(res, m.data)

	return res, nil
}

func (m *Categories) GetAllWithParent(ctx context.Context, parentID int) ([]njudge.Category, error) {
	m.Lock()
	defer m.Unlock()
	res := make([]njudge.Category, 0)
	for ind := range m.data {
		if parentID > 0 {
			if m.data[ind].ParentID.Valid && m.data[ind].ParentID.Int == parentID {
				res = append(res, m.data[ind])
			}
		} else if !m.data[ind].ParentID.Valid {
			res = append(res, m.data[ind])
		}
	}

	return res, nil
}

func (m *Categories) GetByName(ctx context.Context, name string) (*njudge.Category, error) {
	m.Lock()
	defer m.Unlock()
	for ind := range m.data {
		if m.data[ind].Name == name {
			res := m.data[ind]
			return &res, nil
		}
	}

	return nil, njudge.ErrorCategoryNotFound
}

func (m *Categories) Insert(ctx context.Context, u njudge.Category) (*njudge.Category, error) {
	m.Lock()
	defer m.Unlock()
	u.ID = m.nextId
	m.nextId++
	m.data = append(m.data, u)

	res := m.data[len(m.data)-1]
	return &res, nil
}

func (m *Categories) Delete(ctx context.Context, ID int) error {
	m.Lock()
	defer m.Unlock()
	for ind := range m.data {
		if m.data[ind].ID == ID {
			m.data[ind] = m.data[len(m.data)-1]
			m.data = m.data[:len(m.data)-1]
			return nil
		}
	}

	return njudge.ErrorCategoryNotFound
}

func (m *Categories) Update(ctx context.Context, cat njudge.Category) error {
	m.Lock()
	defer m.Unlock()
	for ind := range m.data {
		if m.data[ind].ID == cat.ID {
			m.data[ind] = cat
			return nil
		}
	}
	return njudge.ErrorCategoryNotFound
}
