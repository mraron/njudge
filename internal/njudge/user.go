package njudge

import (
	"context"
	"errors"
	"math/rand"
	"sync"
	"time"
	"unicode"

	"golang.org/x/crypto/bcrypt"
)

type HashedPassword string

func NewHashedPassword(password string) (HashedPassword, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return HashedPassword(hashed), err
}

type UserActivationInfo struct {
	Activated bool
	Key       string
}

func GenerateActivationKey() UserActivationInfo {
	// TODO: move NewSource to rand.go as an init() func
	var (
		alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ012345678901234567890123456789"
		ans      = make([]byte, 32)
	)

	src := rand.NewSource(time.Now().UnixNano())

	for i := 0; i < 32; i++ {
		ans[i] = alphabet[(int(src.Int63()))%len(alphabet)]
	}

	return UserActivationInfo{
		Activated: false,
		Key:       string(ans),
	}
}

type UserSettings struct {
	ShowUnsolvedTags bool
}

type User struct {
	ID             int
	Name           string
	Password       HashedPassword
	Email          string
	ActivationInfo UserActivationInfo
	Role           string
	Points         float64
	Settings       UserSettings
	Created        time.Time
}

var (
	ErrorNonAlphanumeric = errors.New("njudge: string is not alphanumeric")
	ErrorFieldRequired   = errors.New("njudge: field must not be empty")
	ErrorUnknownRole     = errors.New("njudge: unknown role")
)

func isAlphanumeric(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) && !unicode.IsDigit(r) {
			return false
		}
	}

	return true
}

func NewUser(name, email, role string) (*User, error) {
	if !isAlphanumeric(name) {
		return nil, ErrorNonAlphanumeric
	}

	if len(name) == 0 || len(email) == 0 || len(role) == 0 {
		return nil, ErrorFieldRequired
	}

	if role != "user" && role != "admin" {
		return nil, ErrorUnknownRole
	}

	return &User{
		Name:           name,
		Email:          email,
		Role:           role,
		ActivationInfo: GenerateActivationKey(),
		Settings: UserSettings{
			ShowUnsolvedTags: true,
		},
		Created: time.Now(),
	}, nil
}

func (u *User) SetPassword(password string) error {
	if len(password) == 0 {
		return ErrorFieldRequired
	}

	var err error
	u.Password, err = NewHashedPassword(password)
	return err
}

func (u *User) Activate() {
	u.ActivationInfo.Activated = true
	u.ActivationInfo.Key = ""
}

func (u *User) AuthenticatePassword(password string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)); err == nil {
		return true
	}

	return false
}

var (
	ErrorUserNotFound = errors.New("njudge: user not found")
)

type Users interface {
	Get(ctx context.Context, ID int) (*User, error)
	GetByName(ctx context.Context, name string) (*User, error)
	Insert(ctx context.Context, u User) (*User, error)
	Delete(ctx context.Context, ID int) error
	Update(ctx context.Context, user User) error
}

type MemoryUsers struct {
	sync.Mutex
	nextId int
	data   []User
}

func NewMemoryUsers() *MemoryUsers {
	return &MemoryUsers{
		nextId: 1,
		data:   make([]User, 0),
	}
}

func (m *MemoryUsers) Get(ctx context.Context, ID int) (*User, error) {
	m.Lock()
	defer m.Unlock()
	for ind := range m.data {
		if m.data[ind].ID == ID {
			res := m.data[ind]
			return &res, nil
		}
	}

	return nil, ErrorUserNotFound
}

func (m *MemoryUsers) GetByName(ctx context.Context, name string) (*User, error) {
	m.Lock()
	defer m.Unlock()
	for ind := range m.data {
		if m.data[ind].Name == name {
			res := m.data[ind]
			return &res, nil
		}
	}

	return nil, ErrorUserNotFound
}

func (m *MemoryUsers) Insert(ctx context.Context, u User) (*User, error) {
	m.Lock()
	defer m.Unlock()
	u.ID = m.nextId
	m.nextId++
	m.data = append(m.data, u)

	res := m.data[len(m.data)-1]
	return &res, nil
}

func (m *MemoryUsers) Delete(ctx context.Context, ID int) error {
	m.Lock()
	defer m.Unlock()
	for ind := range m.data {
		if m.data[ind].ID == ID {
			m.data[ind] = m.data[len(m.data)-1]
			m.data = m.data[:len(m.data)-1]
			return nil
		}
	}

	return ErrorUserNotFound
}

func (m *MemoryUsers) Update(ctx context.Context, user User) error {
	m.Lock()
	defer m.Unlock()
	for ind := range m.data {
		if m.data[ind].ID == user.ID {
			m.data[ind] = user
			return nil
		}
	}
	return ErrorUserNotFound
}
