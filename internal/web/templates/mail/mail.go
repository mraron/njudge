package mail

import (
	"embed"
	"html/template"
	"io"
)

//go:embed all:*.gohtml
var FS embed.FS

var forgotPasswordTemplate = template.Must(template.New("forgot_password.gohtml").ParseFS(FS, "_template.gohtml", "forgot_password.gohtml"))

type ForgotPasswordViewModel struct {
	Name string
	URL  string
	Key  string
}

func (vm ForgotPasswordViewModel) Execute(w io.Writer) error {
	return forgotPasswordTemplate.Execute(w, vm)
}

var activationTemplate = template.Must(template.New("activation.gohtml").ParseFS(FS, "_template.gohtml", "activation.gohtml"))

type ActivationViewModel struct {
	Name          string
	URL           string
	ActivationKey string
}

func (vm ActivationViewModel) Execute(w io.Writer) error {
	return activationTemplate.Execute(w, vm)
}
