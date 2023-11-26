package njudge_test

import (
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
