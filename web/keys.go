package web

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"github.com/golang-jwt/jwt"
	"io/ioutil"
	"time"
)

func (s *Server) getJWT() (string, error) {
	claims := &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(10 * time.Minute).Unix(),
		NotBefore: time.Now().Unix(),
		Issuer:    "njudge web",
		IssuedAt:  time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS512, claims)
	jwt, err := token.SignedString(s.Keys.PrivateKey)
	if err != nil {
		return "", err
	}

	return jwt, nil
}

func (s *Server) parseKeys() {
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
