package main

import (
	"testing"
)

func TestFullname(t *testing.T) {
	n := fullname("a", "b")
	if n != "a b" {
		t.Errorf("Unmatch %v", n)
	}
}
