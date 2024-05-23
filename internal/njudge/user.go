package njudge

import (
	"context"
	"errors"
	"math/rand"
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

func genRandBytes(l int) []byte {
	// TODO: move NewSource to rand.go as an init() func
	var (
		alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ012345678901234567890123456789"
		ans      = make([]byte, l)
	)

	src := rand.NewSource(time.Now().UnixNano())

	for i := 0; i < 32; i++ {
		ans[i] = alphabet[(int(src.Int63()))%len(alphabet)]
	}

	return ans
}

func GenerateActivationKey() UserActivationInfo {
	ans := genRandBytes(32)

	return UserActivationInfo{
		Activated: false,
		Key:       string(ans),
	}
}

type UserSettings struct {
	ShowUnsolvedTags bool
}

type ForgottenPasswordKey struct {
	ID         int
	UserID     int
	Key        string
	ValidUntil time.Time
}

func NewForgottenPasswordKey(validDuration time.Duration) ForgottenPasswordKey {
	return ForgottenPasswordKey{
		Key:        string(genRandBytes(32)),
		ValidUntil: time.Now().Add(validDuration),
	}
}

func (f *ForgottenPasswordKey) IsValid() bool {
	return time.Now().Before(f.ValidUntil)
}

var UserFields = struct {
	ID                   string
	Name                 string
	Password             string
	Email                string
	ActivationInfo       string
	Role                 string
	Points               string
	Settings             string
	Created              string
	ForgottenPasswordKey string
}{
	ID:                   "id",
	Name:                 "name",
	Password:             "password",
	Email:                "email",
	ActivationInfo:       "activated",
	Role:                 "role",
	Points:               "points",
	Settings:             "settings",
	Created:              "created",
	ForgottenPasswordKey: "forgotten_password_key",
}

type User struct {
	ID                   int
	Name                 string
	Password             HashedPassword
	Email                string
	ActivationInfo       UserActivationInfo
	Role                 string
	Points               float32
	Settings             UserSettings
	Created              time.Time
	ForgottenPasswordKey *ForgottenPasswordKey
}

var (
	ErrorNonAlphanumeric = errors.New("njudge: string is not alphanumeric")
	ErrorFieldRequired   = errors.New("njudge: field must not be empty")
	ErrorUnknownRole     = errors.New("njudge: unknown role")
	ErrorSameName        = errors.New("njudge: name already in use")
	ErrorSameEmail       = errors.New("njudge: email already in use")
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

type RegisterRequest struct {
	Name     string
	Email    string
	Password string
}

func RegisterUser(ctx context.Context, users Users, req RegisterRequest, postRegisterFunc func(*User) error) (*User, error) {
	u, err := NewUser(req.Name, req.Email, "user")
	if err != nil {
		return nil, err
	}

	if err := u.SetPassword(req.Password); err != nil {
		return nil, err
	}

	u, err = users.Insert(ctx, *u)
	if err != nil {
		return nil, err
	}
	if err = postRegisterFunc(u); err != nil {
		return u, err
	}
	return u, nil
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

func (u *User) SetForgottenPasswordKey(fpkey ForgottenPasswordKey) {
	u.ForgottenPasswordKey = &fpkey
	u.ForgottenPasswordKey.UserID = u.ID
}

func (u *User) DeleteForgottenPasswordKey() {
	u.ForgottenPasswordKey = nil
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
	GetByEmail(ctx context.Context, email string) (*User, error)
	Insert(ctx context.Context, u User) (*User, error)
	Delete(ctx context.Context, ID int) error
	Update(ctx context.Context, user *User, fields []string) error
}
