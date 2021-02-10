package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	if add(2, 3) != 5 {
		t.Error("TestAdd failed")
	}
}

func add(a, b int) int {
	return a + b
}
