package njudge_test

import (
	"context"
	"errors"
	"github.com/mraron/njudge/internal/njudge/memory"
	"testing"

	"github.com/mraron/njudge/internal/njudge"
	"github.com/stretchr/testify/assert"
)

func TestUser(t *testing.T) {
	u, err := njudge.NewUser("mraron", "asd@bsd.com", "user")
	assert.Nil(t, err)
	assert.Nil(t, u.SetPassword("abc"))
	assert.False(t, u.ActivationInfo.Activated)
	u.Activate()
	assert.True(t, u.ActivationInfo.Activated)
	assert.Error(t, u.SetPassword(""))

	_, err = njudge.NewUser("mraron", "asd@bsd.com", "userrr")
	assert.ErrorIs(t, err, njudge.ErrorUnknownRole)
	_, err = njudge.NewUser("mra ron", "asd@bsd.com", "user")
	assert.ErrorIs(t, err, njudge.ErrorNonAlphanumeric)
	_, err = njudge.NewUser("mraron!", "asd@bsd.com", "user")
	assert.ErrorIs(t, err, njudge.ErrorNonAlphanumeric)
	_, err = njudge.NewUser("mráron", "asd@bsd.com", "user")
	assert.Nil(t, err)
}

func TestRegister(t *testing.T) {
	users := memory.NewUsers()
	u, err := njudge.RegisterUser(context.Background(), users, njudge.RegisterRequest{
		Name:     "mraron",
		Email:    "",
		Password: "",
	}, func(user *njudge.User) error {
		return nil
	})
	assert.Nil(t, u)
	assert.ErrorIs(t, err, njudge.ErrorFieldRequired)
	u, err = njudge.RegisterUser(context.Background(), users, njudge.RegisterRequest{
		Name:     "mraron",
		Email:    "asd@bsd.com",
		Password: "abc",
	}, func(user *njudge.User) error {
		return nil
	})
	assert.NotNil(t, u)
	assert.NoError(t, err)
	u2, err := njudge.RegisterUser(context.Background(), users, njudge.RegisterRequest{
		Name:     "mraron",
		Email:    "bsd@bsd.com",
		Password: "abc",
	}, func(user *njudge.User) error {
		return nil
	})
	assert.Nil(t, u2)
	assert.ErrorIs(t, err, njudge.ErrorSameName)
	u2, err = njudge.RegisterUser(context.Background(), users, njudge.RegisterRequest{
		Name:     "mraron2",
		Email:    "asd@bsd.com",
		Password: "abc",
	}, func(user *njudge.User) error {
		return nil
	})
	assert.Nil(t, u2)
	assert.ErrorIs(t, err, njudge.ErrorSameEmail)
	u3, err := njudge.RegisterUser(context.Background(), users, njudge.RegisterRequest{
		Name:     "other",
		Email:    "shalala",
		Password: "abc",
	}, func(user *njudge.User) error {
		return errors.New("postReg error")
	})
	assert.NotNil(t, u3)
	assert.Error(t, err)
}
