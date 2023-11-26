package email

import (
	"context"
	"errors"
	"log"
	"regexp"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"gopkg.in/gomail.v2"
)

// Service sends and email
type Service interface {
	Send(ctx context.Context, m Mail) error
}

// SMTPService sends an email via SMTP
type SMTPService struct {
	From     string
	Host     string
	Port     int
	User     string
	Password string
}

func StripHTMLRegex(s string) string {
	r := regexp.MustCompile("<.*?>")
	return r.ReplaceAllString(s, "")
}

func (s SMTPService) Send(ctx context.Context, mail Mail) error {
	m := gomail.NewMessage()

	m.SetHeader("From", s.From)
	m.SetHeader("To", mail.Recipients...)
	m.SetHeader("Subject", mail.Subject)

	m.SetBody("text/html", mail.Message)
	m.AddAlternative("text/plain", StripHTMLRegex(mail.Message))

	return gomail.NewDialer(s.Host, s.Port, s.User, s.Password).DialAndSend(m)
}

// SendgridService sends and email via the Sendgrid API
type SendgridService struct {
	SenderName    string
	SenderAddress string
	APIKey        string
}

func (s SendgridService) Send(ctx context.Context, m Mail) error {
	if len(m.Recipients) > 1 {
		return errors.New("sendgridMailService doesn't support multiple recipients")
	}

	from := mail.NewEmail(s.SenderName, s.SenderAddress)
	to := mail.NewEmail("", m.Recipients[0])

	htmlContent := m.Message
	plainTextContent := StripHTMLRegex(m.Message)
	message := mail.NewSingleEmail(from, m.Subject, to, plainTextContent, htmlContent)

	client := sendgrid.NewSendClient(s.APIKey)
	_, err := client.SendWithContext(ctx, message)
	return err
}

// LogService logs the email to a *log.Logger
type LogService struct {
	Logger *log.Logger
}

func (l LogService) Send(_ context.Context, m Mail) error {
	l.Logger.Println("sending message", m)
	return nil
}

// ErrorService always return an error when sending an email
type ErrorService struct{}

func (e ErrorService) Send(_ context.Context, _ Mail) error {
	return errors.New("can't send mail")
}
