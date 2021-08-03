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
	Mode string
	Hostname       string
	Url string
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

	CustomHead string

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

		if k.PrivateKey, err = x509.ParsePKCS1PrivateKey(block.Bytes); err != nil {
			return err
		}

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