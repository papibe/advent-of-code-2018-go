package main

import (
	"testing"
)

func Test_power_at_3_5_should_be_8(t *testing.T) {
	expected := 4
	result := power_at(3, 5, 8)
	if result != expected {
		t.Errorf("'./example.txt' -> %v; need '%v'", result, expected)
	}
}

func Test_power_at_122_79_should_be_57(t *testing.T) {
	expected := -5
	result := power_at(122, 79, 57)
	if result != expected {
		t.Errorf("'./example.txt' -> %v; need '%v'", result, expected)
	}
}

func Test_power_at_217_196_should_be_39(t *testing.T) {
	expected := 0
	result := power_at(217, 196, 39)
	if result != expected {
		t.Errorf("'./example.txt' -> %v; need '%v'", result, expected)
	}
}

func Test_power_at_101_153_should_be_71(t *testing.T) {
	expected := 4
	result := power_at(101, 153, 71)
	if result != expected {
		t.Errorf("'./example.txt' -> %v; need '%v'", result, expected)
	}
}

func Test_solution_18_should_be_33_45(t *testing.T) {
	expected_x := 33
	expected_y := 45
	result_x, result_y := solution(18)
	if result_x != expected_x || result_y != expected_y {
		t.Errorf("'./example.txt' -> %v,%v; need '%v,%v'", result_x, result_y, expected_x, expected_y)
	}
}

func Test_solution_42_should_be_21_61(t *testing.T) {
	expected_x := 21
	expected_y := 61
	result_x, result_y := solution(42)
	if result_x != expected_x || result_y != expected_y {
		t.Errorf("'./example.txt' -> %v,%v; need '%v,%v'", result_x, result_y, expected_x, expected_y)
	}
}
