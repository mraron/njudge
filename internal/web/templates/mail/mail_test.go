package mail_test

import (
	"bytes"
	"github.com/mraron/njudge/internal/web/templates/mail"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestForgotPassword(t *testing.T) {
	vm := mail.ForgotPasswordViewModel{
		Name: "teszt1",
		URL:  "teszt2",
		Key:  "teszt3",
	}
	buf := &bytes.Buffer{}
	assert.NoError(t, vm.Execute(buf))
}

func TestActivation(t *testing.T) {
	vm := mail.ActivationViewModel{
		Name:          "teszt1",
		URL:           "teszt2",
		ActivationKey: "teszt3",
	}
	buf := &bytes.Buffer{}
	assert.NoError(t, vm.Execute(buf))
}
