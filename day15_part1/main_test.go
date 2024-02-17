package main

import (
	"testing"
)

func Test_example3_should_be_27730(t *testing.T) {
	input := "example3.txt"
	expected := 27730
	result := solution(input)
	if result != expected {
		t.Errorf("combat outcome from %v was %v; need '%v'", input, result, expected)
	}
}

func Test_example4_should_be_36334(t *testing.T) {
	input := "example4.txt"
	expected := 36334
	result := solution(input)
	if result != expected {
		t.Errorf("combat outcome from %v was %v; need '%v'", input, result, expected)
	}
}

func Test_example5_should_be_39514(t *testing.T) {
	input := "example5.txt"
	expected := 39514
	result := solution(input)
	if result != expected {
		t.Errorf("combat outcome from %v was %v; need '%v'", input, result, expected)
	}
}

func Test_example6_should_be_27755(t *testing.T) {
	input := "example6.txt"
	expected := 27755
	result := solution(input)
	if result != expected {
		t.Errorf("combat outcome from %v was %v; need '%v'", input, result, expected)
	}
}

func Test_example7_should_be_28944(t *testing.T) {
	input := "example7.txt"
	expected := 28944
	result := solution(input)
	if result != expected {
		t.Errorf("combat outcome from %v was %v; need '%v'", input, result, expected)
	}
}

func Test_example8_should_be_18740(t *testing.T) {
	input := "example8.txt"
	expected := 18740
	result := solution(input)
	if result != expected {
		t.Errorf("combat outcome from %v was %v; need '%v'", input, result, expected)
	}
}

func Test_example5_should_be_27755(t *testing.T) {
	input := "example6.txt"
	expected := 27755
	result := solution(input)
	if result != expected {
		t.Errorf("combat outcome from %v was %v; need '%v'", input, result, expected)
	}
}
