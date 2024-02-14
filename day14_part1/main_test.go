package main

import (
	"testing"
)

func Test_after_5_should_be_0124515891(t *testing.T) {
	input := 5
	expected := "0124515891"
	result := solution(input)
	if result != expected {
		t.Errorf("after %v recipes -> %v; need '%v'", input, result, expected)
	}
}

func Test_after_9_should_be_5158916779(t *testing.T) {
	input := 9
	expected := "5158916779"
	result := solution(input)
	if result != expected {
		t.Errorf("after %v recipes -> %v; need '%v'", input, result, expected)
	}
}

func Test_after_18_should_be_9251071085(t *testing.T) {
	input := 18
	expected := "9251071085"
	result := solution(input)
	if result != expected {
		t.Errorf("after %v recipes -> %v; need '%v'", input, result, expected)
	}
}

func Test_after_2018_should_be_5941429882(t *testing.T) {
	input := 2018
	expected := "5941429882"
	result := solution(input)
	if result != expected {
		t.Errorf("after %v recipes -> %v; need '%v'", input, result, expected)
	}
}
