package database

import (
	"fmt"
	"testing"
)

func TestConnect(t *testing.T) {
	c := GetConnect()
	if c != nil {
		r, e := c.Do("GET", "aaa")
		if e == nil {
			fmt.Printf("get aaa %s\n", string(r.([]byte)))
		}
	}
}
