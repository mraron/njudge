package memory

import (
	"context"
	"sync"

	"github.com/mraron/njudge/internal/njudge"
)

type Users struct {
	sync.Mutex
	nextId int
	data   []njudge.User
}

func NewUsers() *Users {
	return &Users{
		nextId: 1,
		data:   make([]njudge.User, 0),
	}
}

func (m *Users) Get(ctx context.Context, ID int) (*njudge.User, error) {
	m.Lock()
	defer m.Unlock()
	for ind := range m.data {
		if m.data[ind].ID == ID {
			res := m.data[ind]
			return &res, nil
		}
	}

	return nil, njudge.ErrorUserNotFound
}

func (m *Users) GetByName(ctx context.Context, name string) (*njudge.User, error) {
	m.Lock()
	defer m.Unlock()
	for ind := range m.data {
		if m.data[ind].Name == name {
			res := m.data[ind]
			return &res, nil
		}
	}

	return nil, njudge.ErrorUserNotFound
}

func (m *Users) Insert(ctx context.Context, u njudge.User) (*njudge.User, error) {
	m.Lock()
	defer m.Unlock()
	u.ID = m.nextId
	m.nextId++
	m.data = append(m.data, u)

	res := m.data[len(m.data)-1]
	return &res, nil
}

func (m *Users) Delete(ctx context.Context, ID int) error {
	m.Lock()
	defer m.Unlock()
	for ind := range m.data {
		if m.data[ind].ID == ID {
			m.data[ind] = m.data[len(m.data)-1]
			m.data = m.data[:len(m.data)-1]
			return nil
		}
	}

	return njudge.ErrorUserNotFound
}

func (m *Users) Update(ctx context.Context, user njudge.User) error {
	m.Lock()
	defer m.Unlock()
	for ind := range m.data {
		if m.data[ind].ID == user.ID {
			m.data[ind] = user
			return nil
		}
	}
	return njudge.ErrorUserNotFound
}
