package main

import (
	"testing"
)

func TestExample(t *testing.T) {
	expected := 240
	result := solution("./example.txt")
	if result != expected {
		t.Errorf("'./example.txt' -> %v; need '%v'", result, expected)
	}
}
