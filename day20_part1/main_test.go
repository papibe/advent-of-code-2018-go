package main

import (
	"testing"
)

var area2 = `
#########
#.|.|.|.#
#-#######
#.|.|.|.#
#-#####-#
#.#.#X|.#
#-#-#####
#.|.|.|.#
#########
`

var area3 = `
###########
#.|.#.|.#.#
#-###-#-#-#
#.|.|.#.#.#
#-#####-#-#
#.#.#X|.#.#
#-#-#####-#
#.#.|.|.|.#
#-###-###-#
#.|.|.#.|.#
###########
`

var area4 = `
#############
#.|.|.|.|.|.#
#-#####-###-#
#.#.|.#.#.#.#
#-#-###-#-#-#
#.#.#.|.#.|.#
#-#-#-#####-#
#.#.#.#X|.#.#
#-#-#-###-#-#
#.|.#.|.#.#.#
###-#-###-#-#
#.|.#.|.|.#.#
#############
`

var area5 = `
###############
#.|.|.|.#.|.|.#
#-###-###-#-#-#
#.|.#.|.|.#.#.#
#-#########-#-#
#.#.|.|.|.|.#.#
#-#-#########-#
#.#.#.|X#.|.#.#
###-#-###-#-#-#
#.|.#.#.|.#.|.#
#-###-#####-###
#.|.#.|.|.#.#.#
#-#-#####-#-#-#
#.#.|.|.|.#.|.#
###############
`

func Test_example2_should_draw_area2(t *testing.T) {
	input := "example2.txt"
	expected := area2
	regex := parse(input)
	doors, rooms, walls := read_regex(regex, 0, 0, 0)
	result := draw(doors, rooms, walls)
	if result != expected {
		t.Errorf("area discovered from %v was %v; need '%v'", input, result, expected)
	}
}

func Test_example3_should_draw_area3(t *testing.T) {
	input := "example3.txt"
	expected := area3
	regex := parse(input)
	doors, rooms, walls := read_regex(regex, 0, 0, 0)
	result := draw(doors, rooms, walls)
	if result != expected {
		t.Errorf("area discovered from %v was %v; need '%v'", input, result, expected)
	}
}

func Test_example4_should_draw_area4(t *testing.T) {
	input := "example4.txt"
	expected := area4
	regex := parse(input)
	doors, rooms, walls := read_regex(regex, 0, 0, 0)
	result := draw(doors, rooms, walls)
	if result != expected {
		t.Errorf("area discovered from %v was %v; need '%v'", input, result, expected)
	}
}

func Test_example5_should_draw_area4(t *testing.T) {
	input := "example5.txt"
	expected := area5
	regex := parse(input)
	doors, rooms, walls := read_regex(regex, 0, 0, 0)
	result := draw(doors, rooms, walls)
	if result != expected {
		t.Errorf("area discovered from %v was %v; need '%v'", input, result, expected)
	}
}

func Test_example1_should_be_3(t *testing.T) {
	input := "example1.txt"
	expected := 3
	result := solution(input)
	if result != expected {
		t.Errorf("largest number of doors from %v was %v; need '%v'", input, result, expected)
	}
}

func Test_example2_should_be_10(t *testing.T) {
	input := "example2.txt"
	expected := 10
	result := solution(input)
	if result != expected {
		t.Errorf("largest number of doors from %v was %v; need '%v'", input, result, expected)
	}
}

func Test_example3_should_be_18(t *testing.T) {
	input := "example3.txt"
	expected := 18
	result := solution(input)
	if result != expected {
		t.Errorf("largest number of doors from %v was %v; need '%v'", input, result, expected)
	}
}

func Test_example4_should_be_23(t *testing.T) {
	input := "example4.txt"
	expected := 23
	result := solution(input)
	if result != expected {
		t.Errorf("largest number of doors from %v was %v; need '%v'", input, result, expected)
	}
}

func Test_example5_should_be_31(t *testing.T) {
	input := "example5.txt"
	expected := 31
	result := solution(input)
	if result != expected {
		t.Errorf("largest number of doors from %v was %v; need '%v'", input, result, expected)
	}
}
