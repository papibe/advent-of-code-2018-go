package main

import (
	"testing"
)

func TestExample1(t *testing.T) {
	expected := 32
	result := solution(9, 25)
	if result != expected {
		t.Errorf("'./example.txt' -> %v; need '%v'", result, expected)
	}
}

func TestExample2(t *testing.T) {
	expected := 8317
	result := solution(10, 1618)
	if result != expected {
		t.Errorf("'./example.txt' -> %v; need '%v'", result, expected)
	}
}

func TestExample3(t *testing.T) {
	expected := 146373
	result := solution(13, 7999)
	if result != expected {
		t.Errorf("'./example.txt' -> %v; need '%v'", result, expected)
	}
}

func TestExample4(t *testing.T) {
	expected := 2764
	result := solution(17, 1105)
	if result != expected {
		t.Errorf("'./example.txt' -> %v; need '%v'", result, expected)
	}
}

func TestExample5(t *testing.T) {
	expected := 37305
	result := solution(30, 5807)
	if result != expected {
		t.Errorf("'./example.txt' -> %v; need '%v'", result, expected)
	}
}

func TestExample6(t *testing.T) {
	expected := 54718
	result := solution(21, 6111)
	if result != expected {
		t.Errorf("'./example.txt' -> %v; need '%v'", result, expected)
	}
}
