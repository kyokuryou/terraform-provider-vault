package client

import "testing"

func TestGenKey(t *testing.T) {
	c := NewTest()
	r, err := c.GenKey("test1")
	if err != nil {
		println(err.Error())
	}
	println(r)
}
