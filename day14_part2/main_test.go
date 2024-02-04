package main

import (
	"testing"
)

func Test_scores_01245_appears_after_5(t *testing.T) {
	input := "01245"
	expected := 5
	result := solution(input)
	if result != expected {
		t.Errorf("after %v recipes -> %v; need '%v'", input, result, expected)
	}
}

func Test_scores_51589_appears_after_9(t *testing.T) {
	input := "51589"
	expected := 9
	result := solution(input)
	if result != expected {
		t.Errorf("after %v recipes -> %v; need '%v'", input, result, expected)
	}
}

func Test_scores_92510_appears_after_18(t *testing.T) {
	input := "92510"
	expected := 18
	result := solution(input)
	if result != expected {
		t.Errorf("after %v recipes -> %v; need '%v'", input, result, expected)
	}
}

func Test_scores_59414_appears_after_2018(t *testing.T) {
	input := "59414"
	expected := 2018
	result := solution(input)
	if result != expected {
		t.Errorf("after %v recipes -> %v; need '%v'", input, result, expected)
	}
}
