package main

import (
	"testing"
)

func TestExample(t *testing.T) {
	expected := 12
	result := solution("./example.txt")
	if result != expected {
		t.Errorf("'./example.txt' -> %v; need '%v'", result, expected)
	}
}
