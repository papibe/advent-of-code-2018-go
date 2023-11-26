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

func Test_solution_18_should_be_90_269_16(t *testing.T) {
	expected_x := 90
	expected_y := 269
	expected_size := 16
	result_x, result_y, result_size := solution(18)
	if result_x != expected_x || result_y != expected_y || result_size != expected_size {
		t.Errorf("'./example.txt' -> %v,%v; need '%v,%v'", result_x, result_y, expected_x, expected_y)
	}
}

func Test_solution_42_should_be_232_251_12(t *testing.T) {
	expected_x := 232
	expected_y := 251
	expected_size := 12
	result_x, result_y, result_size := solution(18)
	if result_x != expected_x || result_y != expected_y || result_size != expected_size {
		t.Errorf("'./example.txt' -> %v,%v; need '%v,%v'", result_x, result_y, expected_x, expected_y)
	}
}
