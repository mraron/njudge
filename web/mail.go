package web

import (
	"crypto/tls"
	"fmt"
	"io"
	"net/smtp"
	"strings"
)

type Mail struct {
	Recipients []string
	Subject    string
	Message    string
}

func (m Mail) Body(Sender string) string {
	return fmt.Sprintf("From: %s\r\nTo: %s\r\nSubject: %s\r\nContent-Type: text/html; charset=utf-8\r\n\r\n%s", Sender, strings.Join(m.Recipients, ";"), m.Subject, m.Message)
}

func (s *Server) SendMail(mail Mail) error {
	auth := smtp.PlainAuth("", s.MailAccount, s.MailAccountPassword, s.MailServerHost)

	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         s.MailServerHost + ":" + s.MailServerPort,
	}

	conn, err := tls.Dial("tcp", s.MailServerHost+":"+s.MailServerPort, tlsConfig)
	if err != nil {
		return err
	}

	defer conn.Close()

	client, err := smtp.NewClient(conn, s.MailServerHost)
	if err != nil {
		return err
	}

	if err = client.Auth(auth); err != nil {
		return err
	}

	defer client.Close()

	if err = client.Mail(s.MailAccount); err != nil {
		return err
	}

	for _, r := range mail.Recipients {
		if err = client.Rcpt(r); err != nil {
			return err
		}
	}

	w, err := client.Data()
	if err != nil {
		return err
	}

	if _, err = io.WriteString(w, mail.Body(s.MailAccount)); err != nil {
		return err
	}

	if err = w.Close(); err != nil {
		return err
	}

	return nil
}
