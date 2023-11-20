package njudge_test

import (
	"context"
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

	m := njudge.NewMemoryUsers()
	u, err = m.Insert(context.TODO(), *u)
	assert.Nil(t, err)
	assert.Greater(t, u.ID, 0)

	id := u.ID

	u, err = m.Get(context.TODO(), u.ID)
	assert.Nil(t, err)
	assert.Equal(t, id, u.ID)

	dummy, err := njudge.NewUser("dummy", "asd@bsd2.com", "admin")
	assert.Nil(t, err)
	u2, err := m.Insert(context.TODO(), *dummy)
	assert.Nil(t, err)
	assert.Greater(t, u2.ID, 0)
	assert.NotEqual(t, u2.ID, id)

	u2.Name = "dummy2"
	u2, _ = m.Get(context.TODO(), u2.ID)
	assert.Equal(t, u2.Name, "dummy")

	u2.Name = "dummy2"
	err = m.Update(context.TODO(), *u2)
	assert.Nil(t, err)

	u2, _ = m.Get(context.TODO(), u2.ID)
	assert.Equal(t, u2.Name, "dummy2")

	err = m.Delete(context.TODO(), u2.ID)
	assert.Nil(t, err)

	_, err = m.Get(context.TODO(), u2.ID)
	assert.ErrorIs(t, err, njudge.ErrorUserNotFound)

	_, err = m.GetByName(context.TODO(), "mraron")
	assert.Nil(t, err)
}
