package client

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"os"
	"path/filepath"
)

type Client struct {
	BasePath string
	// APIKey is an optional API key.
	PrivateKey *rsa.PrivateKey
	// BasicAuth is optional basic auth credentials.
	PublicKey *rsa.PublicKey
}

func New(basePath string, privateKey string) (*Client, error) {
	var pk *rsa.PrivateKey
	var err error
	if pk, err = decodePrivateKeyFromPEM([]byte(privateKey)); err != nil {
		return nil, err
	}
	return &Client{
		BasePath:   basePath,
		PrivateKey: pk,
		PublicKey:  pk.Public().(*rsa.PublicKey),
	}, err
}

func decodePrivateKeyFromPEM(privateKeyPEM []byte) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode(privateKeyPEM)
	if block == nil || block.Type != "RSA PRIVATE KEY" {
		return nil, errors.New("unable to decode PEM")
	}
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return privateKey, nil
}

func fileJoin(elem ...string) (string, error) {
	var file string
	for index, value := range elem {
		if index < len(elem)-1 {
			file = filepath.Join(file, value)
			if _, err := os.Stat(file); err == nil {
				continue
			}
			if err := os.Mkdir(file, os.ModePerm); err != nil {
				return "", err
			}
		} else {
			file = filepath.Join(file, value)
		}
	}
	return file, nil
}
