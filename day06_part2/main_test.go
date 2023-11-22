package main

import (
	"testing"
)

func TestExample(t *testing.T) {
	expected := 16
	result := solution("./example.txt", 32)
	if result != expected {
		t.Errorf("'./example.txt' -> %v; need '%v'", result, expected)
	}
}
