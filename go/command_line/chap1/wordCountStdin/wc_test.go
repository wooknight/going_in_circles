package main

import (
	"strings"
	"testing"
)

func TestCount(t *testing.T) {
	str := "the pan galactic gargle blaster"
	cnt := count(strings.NewReader(str))
	if cnt != 5 {
		t.Fatal("count does not match ")
	}
	t.Log("count match")
}
