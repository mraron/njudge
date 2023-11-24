package memory_test

import (
	"context"
	"testing"

	"github.com/mraron/njudge/internal/njudge"
	"github.com/mraron/njudge/internal/njudge/memory"
	"github.com/stretchr/testify/assert"
)

func TestUsers(t *testing.T) {
	u, err := njudge.NewUser("mraron", "asd@bsd.com", "user")
	assert.Nil(t, err)

	m := memory.NewUsers()
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

func TestUsersUnique(t *testing.T) {
	u1, err := njudge.NewUser("mraron", "asd@bsd.com", "user")
	assert.Nil(t, err)
	u2, err := njudge.NewUser("mraron", "csd@bsd.com", "user")
	assert.Nil(t, err)

	m := memory.NewUsers()
	_, err = m.Insert(context.TODO(), *u1)
	assert.Nil(t, err)
	_, err = m.Insert(context.TODO(), *u2)
	assert.Error(t, err)

	m = memory.NewUsers()
	u2.Name = "mraron2"
	u2.Email = u1.Email
	_, err = m.Insert(context.TODO(), *u1)
	assert.Nil(t, err)
	_, err = m.Insert(context.TODO(), *u2)
	assert.Error(t, err)
}
