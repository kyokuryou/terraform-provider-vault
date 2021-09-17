package client

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"io/fs"
	"io/ioutil"
	"os"
)

func (d *Client) Sign(plainBytes []byte) (string, error) {
	h := sha1.New()
	h.Write(plainBytes)
	hash := h.Sum(nil)
	signature, err := rsa.SignPKCS1v15(rand.Reader, d.PrivateKey, crypto.SHA1, hash[:])
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(signature), nil
}

func (d *Client) Encode(name string, plainBytes []byte) error {
	filename, err := fileJoin(d.BasePath, "storage", name)
	if err != nil {
		return err
	}
	cipherBytes, err := rsa.EncryptPKCS1v15(rand.Reader, d.PublicKey, plainBytes)
	if err != nil {
		return err
	}
	buf := make([]byte, base64.StdEncoding.EncodedLen(len(cipherBytes)))
	base64.StdEncoding.Encode(buf, cipherBytes)
	if err = ioutil.WriteFile(filename, buf, fs.ModePerm); err != nil {
		return err
	}
	return nil
}

func (d *Client) Decode(name string) ([]byte, error) {
	filename, err := fileJoin(d.BasePath, "storage", name)
	if err != nil {
		return nil, err
	}
	cipherBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	dbuf := make([]byte, base64.StdEncoding.DecodedLen(len(cipherBytes)))
	n, err := base64.StdEncoding.Decode(dbuf, cipherBytes)
	if err != nil {
		return nil, err
	}
	plainBytes, err := rsa.DecryptPKCS1v15(rand.Reader, d.PrivateKey, dbuf[:n])
	return plainBytes, err
}

func (d *Client) Delete(name string) error {
	filename, err := fileJoin(d.BasePath, "storage", name)
	if err != nil {
		return err
	}
	return os.Remove(filename)
}
