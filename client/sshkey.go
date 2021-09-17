package client

import (
	"fmt"
	"golang.org/x/crypto/ssh"
	"strings"
)

func (d *Client) GenKey(comment string) (string, error) {
	rsaKey, err := ssh.NewPublicKey(d.PublicKey)
	if err != nil {
		return "", err
	}
	keyBytes := ssh.MarshalAuthorizedKey(rsaKey)
	keyString := strings.TrimRight(string(keyBytes[:]), "\r\n")
	if err != nil {
		return "", err
	}
	result := strings.TrimSpace(fmt.Sprintf("%s %s", keyString, comment))
	return result, nil
}
