package services

import (
	"context"
	"errors"
	"github.com/mraron/njudge/internal/web/domain/email"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"gopkg.in/gomail.v2"
	"log"
	"regexp"
)

type MailService interface {
	Send(ctx context.Context, m email.Mail) error
}

type SMTPMailService struct {
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

func (s SMTPMailService) Send(ctx context.Context, mail email.Mail) error {
	m := gomail.NewMessage()

	m.SetHeader("From", s.From)
	m.SetHeader("To", mail.Recipients...)
	m.SetHeader("Subject", mail.Subject)

	m.SetBody("text/html", mail.Message)
	m.AddAlternative("text/plain", StripHTMLRegex(mail.Message))

	return gomail.NewDialer(s.Host, s.Port, s.User, s.Password).DialAndSend(m)
}

type SendgridMailService struct {
	SenderName    string
	SenderAddress string
	APIKey        string
}

func (s SendgridMailService) Send(ctx context.Context, m email.Mail) error {
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

type LogMailService struct {
	Logger *log.Logger
}

func (l LogMailService) Send(_ context.Context, m email.Mail) error {
	l.Logger.Println("sending message", m)
	return nil
}

type ErrorMailService struct{}

func (e ErrorMailService) Send(_ context.Context, _ email.Mail) error {
	return errors.New("can't send mail")
}
