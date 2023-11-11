package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

type Node struct {
	unit rune
	next *Node
	prev *Node
}

func parse(filename string) string {
	data, err := os.ReadFile(filename)

	if err != nil {
		panic("file error")
	}
	content := string(data)

	return strings.Trim(content, "\n")
}

func reduce(polymer *Node) bool {
	current := polymer
	removed := false

	for current != nil && current.next != nil {
		current_char := current.unit
		next_char := current.next.unit
		// handle of comples logic

		units_are_same_type := unicode.ToLower(current_char) == unicode.ToLower(next_char)
		first_unit_low_and_second_upper := unicode.IsLower(current_char) && unicode.IsUpper(next_char)
		first_unit_upper_and_second_low := unicode.IsUpper(current_char) && unicode.IsLower(next_char)
		units_are_opposite_polarity := first_unit_low_and_second_upper || first_unit_upper_and_second_low

		if units_are_same_type && units_are_opposite_polarity {
			// get relevant nodes for removal
			previous := current.prev
			next := current.next.next

			// remove both nodes
			previous.next = next
			next.prev = previous

			removed = true
			current = previous
		} else {
			current = current.next
		}
	}

	return removed
}

func create_list(polymer string) *Node {
	dummy_head := &Node{unit: '>', next: nil, prev: nil}
	current := dummy_head

	for _, rune := range polymer {
		node := &Node{unit: rune, next: nil, prev: current}
		current.next = node
		current = current.next
	}
	dummy_tail := &Node{unit: '<', next: nil, prev: current}
	current.next = dummy_tail

	return dummy_head
}

func polymer_len(polymer *Node) int {
	length := 0
	current := polymer

	for current != nil {
		length += 1
		current = current.next
	}
	return length - 2
}

func solve(polymer *Node) int {

	removed := reduce(polymer)
	for removed {
		removed = reduce(polymer)
	}

	return polymer_len(polymer)
}

func solution(filename string) int {
	polymer := parse(filename)
	polymer_list := create_list(polymer)
	return solve(polymer_list)
}

func main() {
	fmt.Println(solution("./input.txt"))
}
