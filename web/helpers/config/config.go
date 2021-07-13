package config

import "crypto/rsa"

type Server struct {
	Mode string
	Hostname       string
	Port           string
	ProblemsDir    string
	SubmissionsDir string
	TemplatesDir   string

	CookieSecret string

	GoogleAuth struct {
		Enabled   bool
		ClientKey string
		Secret    string
		Callback  string
	}

	Sendgrid struct {
		Enabled bool
		ApiKey string `json:"api_key"`
		SenderName string `json:"sender_name"`
		SenderAddress string `json:"sender_address"`
	}

	SMTP struct {
		Enabled bool
		MailAccount         string `json:"mail_account"`
		MailServerHost      string `json:"mail_server"`
		MailServerPort      string `json:"mail_port"`
		MailAccountPassword string `json:"mail_password"`
	} `json:"smtp"`

	DBAccount  string
	DBPassword string
	DBHost     string
	DBName     string

	GluePort string

	Keys Keys
}

type Keys struct {
	PrivateKeyLocation string `json:"private_key"`
	PublicKeyLocation  string `json:"public_key"`
	PrivateKey         *rsa.PrivateKey
	PublicKey          *rsa.PublicKey
}