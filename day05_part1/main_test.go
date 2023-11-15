package main

import (
	"testing"
)

func TestExample1(t *testing.T) {
	expected := 0
	result := solution("./example1.txt")
	if result != expected {
		t.Errorf("'./example1.txt' -> %v; need '%v'", result, expected)
	}
}

func TestExample2(t *testing.T) {
	expected := 0
	result := solution("./example2.txt")
	if result != expected {
		t.Errorf("'./example2.txt' -> %v; need '%v'", result, expected)
	}
}

func TestExample3(t *testing.T) {
	expected := 4
	result := solution("./example3.txt")
	if result != expected {
		t.Errorf("'./example3.txt' -> %v; need '%v'", result, expected)
	}
}

func TestExample4(t *testing.T) {
	expected := 6
	result := solution("./example4.txt")
	if result != expected {
		t.Errorf("'./example4.txt' -> %v; need '%v'", result, expected)
	}
}

func TestExample5(t *testing.T) {
	expected := 10
	result := solution("./example5.txt")
	if result != expected {
		t.Errorf("'./example5.txt' -> %v; need '%v'", result, expected)
	}
}
