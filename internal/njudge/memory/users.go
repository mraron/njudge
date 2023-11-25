package memory

import (
	"context"
	"slices"
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

func (m *Users) getByName(ctx context.Context, name string) (*njudge.User, error) {
	for ind := range m.data {
		if m.data[ind].Name == name {
			res := m.data[ind]
			return &res, nil
		}
	}

	return nil, njudge.ErrorUserNotFound
}

func (m *Users) GetByName(ctx context.Context, name string) (*njudge.User, error) {
	m.Lock()
	defer m.Unlock()
	return m.getByName(ctx, name)
}

func (m *Users) getByEmail(ctx context.Context, email string) (*njudge.User, error) {
	for ind := range m.data {
		if m.data[ind].Email == email {
			res := m.data[ind]
			return &res, nil
		}
	}

	return nil, njudge.ErrorUserNotFound
}

func (m *Users) GetByEmail(ctx context.Context, email string) (*njudge.User, error) {
	m.Lock()
	defer m.Unlock()
	return m.getByEmail(ctx, email)
}

func (m *Users) Insert(ctx context.Context, u njudge.User) (*njudge.User, error) {
	m.Lock()
	defer m.Unlock()

	if _, err := m.getByName(ctx, u.Name); err == nil {
		return nil, njudge.ErrorSameName
	}
	if _, err := m.getByEmail(ctx, u.Email); err == nil {
		return nil, njudge.ErrorSameEmail
	}

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

func (m *Users) Update(ctx context.Context, user njudge.User, fields []string) error {
	m.Lock()
	defer m.Unlock()
	for ind := range m.data {
		if m.data[ind].ID == user.ID {
			if slices.Contains(fields, njudge.UserFields.Name) {
				m.data[ind].Name = user.Name
			}
			if slices.Contains(fields, njudge.UserFields.Password) {
				m.data[ind].Password = user.Password
			}
			if slices.Contains(fields, njudge.UserFields.ActivationInfo) {
				m.data[ind].ActivationInfo = user.ActivationInfo
			}
			if slices.Contains(fields, njudge.UserFields.Role) {
				m.data[ind].Role = user.Role
			}
			if slices.Contains(fields, njudge.UserFields.Points) {
				m.data[ind].Points = user.Points
			}
			if slices.Contains(fields, njudge.UserFields.Settings) {
				m.data[ind].Settings = user.Settings
			}
			if slices.Contains(fields, njudge.UserFields.ForgottenPasswordKey) {
				m.data[ind].ForgottenPasswordKey = user.ForgottenPasswordKey
			}

			return nil
		}
	}
	return njudge.ErrorUserNotFound
}
