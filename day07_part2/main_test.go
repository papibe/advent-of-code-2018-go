package main

import (
	"testing"
)

func TestExample(t *testing.T) {
	expected := 15
	result := solution("./example.txt", 2, 0)
	if result != expected {
		t.Errorf("'./example.txt' -> %v; need '%v'", result, expected)
	}
}
