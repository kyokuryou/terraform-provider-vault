package client

import (
	"io/ioutil"
	"testing"
)

func TestNew(t *testing.T) {
	c := NewTest()
	println(c)
}

func NewTest() *Client {
	privatekey, err := ioutil.ReadFile("./secret_key.pem")
	if err != nil {
		println(err.Error())
	}
	c, err := New("./test", string(privatekey))
	if err != nil {
		println(err.Error())
	}
	return c
}
