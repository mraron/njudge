package mail

import (
	"crypto/tls"
	"fmt"
	"github.com/mraron/njudge/internal/web/helpers/config"
	"io"
	"net/smtp"
	"strings"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type Mail struct {
	Recipients []string
	Subject    string
	Message    string
}

func (m Mail) Body(Sender string) string {
	return fmt.Sprintf("From: %s\r\nTo: %s\r\nSubject: %s\r\nContent-Type: text/html; charset=utf-8\r\n\r\n%s", Sender, strings.Join(m.Recipients, ";"), m.Subject, m.Message)
}

func (m Mail) Send(s config.Server) error {
	if s.SMTP.Enabled {
		auth := smtp.PlainAuth("", s.SMTP.MailAccount, s.SMTP.MailAccountPassword, s.SMTP.MailServerHost)

		tlsConfig := &tls.Config{
			InsecureSkipVerify: true,
			ServerName:         s.SMTP.MailServerHost + ":" + s.SMTP.MailServerPort,
		}

		conn, err := tls.Dial("tcp", s.SMTP.MailServerHost+":"+s.SMTP.MailServerPort, tlsConfig)
		if err != nil {
			return err
		}

		defer conn.Close()

		client, err := smtp.NewClient(conn, s.SMTP.MailServerHost)
		if err != nil {
			return err
		}

		if err = client.Auth(auth); err != nil {
			return err
		}

		defer client.Close()

		if err = client.Mail(s.SMTP.MailAccount); err != nil {
			return err
		}

		for _, r := range m.Recipients {
			if err = client.Rcpt(r); err != nil {
				return err
			}
		}

		w, err := client.Data()
		if err != nil {
			return err
		}

		if _, err = io.WriteString(w, m.Body(s.SMTP.MailAccount)); err != nil {
			return err
		}

		if err = w.Close(); err != nil {
			return err
		}

		return nil
	} else if s.Sendgrid.Enabled {
		from := mail.NewEmail(s.Sendgrid.SenderName, s.Sendgrid.SenderAddress)
		to := mail.NewEmail("", m.Recipients[0]) // @TODO erroneous

		plainTextContent := m.Message
		htmlContent := m.Message
		message := mail.NewSingleEmail(from, m.Subject, to, plainTextContent, htmlContent)

		client := sendgrid.NewSendClient(s.Sendgrid.ApiKey)
		resp, err := client.Send(message)
		if err != nil {
			return err
		}

		fmt.Println(resp)

		return nil
	}

	return fmt.Errorf("can't send mail")
}
