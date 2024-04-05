package main

import (
	"testing"
)

func TestExample1ShouldBe2(t *testing.T) {
	expected := 2
	result := solution("./example1.txt")
	if result != expected {
		t.Errorf("'./example1.txt' -> %v; need '%v'", result, expected)
	}
}

func TestExample2ShouldBe4(t *testing.T) {
	expected := 4
	result := solution("./example2.txt")
	if result != expected {
		t.Errorf("'./example2.txt' -> %v; need '%v'", result, expected)
	}
}

func TestExample3ShouldBe3(t *testing.T) {
	expected := 3
	result := solution("./example3.txt")
	if result != expected {
		t.Errorf("'./example3.txt' -> %v; need '%v'", result, expected)
	}
}

func TestExample4ShouldBe8(t *testing.T) {
	expected := 8
	result := solution("./example4.txt")
	if result != expected {
		t.Errorf("'./example4.txt' -> %v; need '%v'", result, expected)
	}
}
