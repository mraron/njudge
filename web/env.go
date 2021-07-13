package web

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/google"
	"github.com/mraron/njudge/utils/problems"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"io/ioutil"
	"time"
)

func (s *Server) SetupEnvironment() {
	if s.Mode == "development" {
		boil.DebugMode = true
	}

	loc, err := time.LoadLocation("Europe/Budapest")
	if err != nil {
		panic(err)
	}
	time.Local = loc
	boil.SetLocation(loc)

	if s.GoogleAuth.Enabled {
		goth.UseProviders(
			google.New(s.GoogleAuth.ClientKey, s.GoogleAuth.Secret, s.GoogleAuth.Callback, "email", "profile"),
		)
	}

	s.ConnectToDB()
	s.ParseKeys()

	s.ProblemStore = problems.NewFsStore(s.ProblemsDir)
}

func (s *Server) ParseKeys() {
	if s.Keys.PrivateKeyLocation != "" {
		if s.Keys.PublicKeyLocation == "" {
			panic("private key filled, public not")
		}

		privateKeyContents, err := ioutil.ReadFile(s.Keys.PrivateKeyLocation)
		if err != nil {
			panic(err)
		}

		block, _ := pem.Decode(privateKeyContents)
		if block == nil {
			panic(fmt.Sprintf("can't parse pem private key file: %s", s.Keys.PrivateKeyLocation))
		}

		if s.Keys.PrivateKey, err = x509.ParsePKCS1PrivateKey(block.Bytes); err != nil {
			panic(err)
		}

		publicKeyContents, err := ioutil.ReadFile(s.Keys.PublicKeyLocation)
		if err != nil {
			panic(err)
		}

		block, _ = pem.Decode(publicKeyContents)
		if block == nil {
			panic(fmt.Sprintf("can't parse pem public key file: %s", s.Keys.PrivateKeyLocation))
		}

		if s.Keys.PublicKey, err = x509.ParsePKCS1PublicKey(block.Bytes); err != nil {
			panic(err)
		}
	}
}

func (s *Server) ConnectToDB() {
	var err error
	s.DB, err = sqlx.Open("postgres", "postgres://"+s.DBAccount+":"+s.DBPassword+"@"+s.DBHost+"/"+s.DBName)

	if err != nil {
		panic(err)
	}

	boil.SetDB(s.DB)
}