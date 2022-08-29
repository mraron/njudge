package config

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"io/ioutil"
)

type Server struct {
	Mode         string
	Hostname     string
	Url          string
	Port         string
	ProblemsDir  string
	TemplatesDir string

	CookieSecret string

	GoogleAuth struct {
		Enabled   bool   `json:"enabled" mapstructure:"enabled"`
		ClientKey string `json:"client_key" mapstructure:"clientkey"`
		Secret    string `json:"secret" mapstructure:"secret"`
		Callback  string `json:"callback" mapstructure:"callback"`
	} `json:"googleAuth" mapstructure:"googleAuth"`

	Sendgrid struct {
		Enabled       bool   `json:"enabled" mapstructure:"enabled"`
		ApiKey        string `json:"api_key" mapstructure:"apikey"`
		SenderName    string `json:"sender_name" mapstructure:"sendername"`
		SenderAddress string `json:"sender_address" mapstructure:"senderaddress"`
	} `json:"sendgrid" mapstructure:"sendgrid"`

	SMTP struct {
		Enabled             bool
		MailAccount         string `json:"mail_account" mapstructure:"mailaccount"`
		MailServerHost      string `json:"mail_server" mapstructure:"mailserver"`
		MailServerPort      string `json:"mail_port" mapstructure:"mailport"`
		MailAccountPassword string `json:"mail_password" mapstructure:"mailpassword"`
	} `json:"smtp" mapstructure:"smtp"`

	CustomHead string

	DBAccount  string
	DBPassword string
	DBHost     string
	DBName     string
	DBPort     int
	DBSSLMode  bool

	GluePort string

	Keys Keys
}

type Keys struct {
	PrivateKeyLocation string `json:"private_key" mapstructure:"privatekey"`
	PublicKeyLocation  string `json:"public_key" mapstructure:"publickey"`
	PrivateKey         *rsa.PrivateKey
	PublicKey          *rsa.PublicKey
}

func (k *Keys) Parse() error {
	if k.PrivateKeyLocation != "" {
		if k.PublicKeyLocation == "" {
			return errors.New("private key filled, public not")
		}

		privateKeyContents, err := ioutil.ReadFile(k.PrivateKeyLocation)
		if err != nil {
			return err
		}

		block, _ := pem.Decode(privateKeyContents)
		if block == nil {
			return fmt.Errorf("can't parse pem private key file: %s", k.PrivateKeyLocation)
		}

		var pKey any
		if pKey, err = x509.ParsePKCS8PrivateKey(block.Bytes); err != nil {
			return err
		}
		k.PrivateKey = pKey.(*rsa.PrivateKey)

		publicKeyContents, err := ioutil.ReadFile(k.PublicKeyLocation)
		if err != nil {
			return err
		}

		block, _ = pem.Decode(publicKeyContents)
		if block == nil {
			return fmt.Errorf("can't parse pem public key file: %s", k.PrivateKeyLocation)
		}

		if k.PublicKey, err = x509.ParsePKCS1PublicKey(block.Bytes); err != nil {
			return err
		}
	}

	return nil
}

func (k *Keys) MustParse() {
	if err := k.Parse(); err != nil {
		panic(err)
	}
}
