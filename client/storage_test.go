package client

import (
	"encoding/json"
	"testing"
)

var data = map[string]interface{}{
	"authc_type": "ssh",
	"image_type": "linux",
	"ip_address": "10.10.10.1",
	"username":   "ydxuser",
}

func TestSign(t *testing.T) {
	c := NewTest()
	j, err := json.Marshal(data)
	if err != nil {
		println(err.Error())
	}
	s, err := c.Sign(j)
	if err != nil {
		println(err.Error())
	}
	println(s)
}

func TestEncode(f *testing.T) {
	c := NewTest()
	j, err := json.Marshal(data)
	if err != nil {
		println(err.Error())
	}
	if err := c.Encode("test1", j); err != nil {
		println(err.Error())
	}
}

func TestDecode(f *testing.T) {
	c := NewTest()

	text, err := c.Decode("test1")
	if err != nil {
		println(err.Error())
	}
	m := make(map[string]interface{})
	if err := json.Unmarshal(text, &m); err != nil {
		println(err.Error())
	}
	println(string(text))
	println(m)
}
