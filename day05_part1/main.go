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

func reduce(polymer *Node) int {
	dummy_head := Node{unit: '>', next: polymer}
	_ = dummy_head
	current := polymer.next

	for current != nil && current.next != nil {
		// fmt.Print(string(current.unit))
		previous := current.prev
		if unicode.ToLower(previous.unit) == unicode.ToLower(current.unit) {
			fmt.Println(string(previous.unit))
		}
		// previous = current.unit
		current = current.next
	}
	fmt.Println()

	return 0
}

func create_list(polymer string) *Node {
	current := &Node{unit: '>', next: nil, prev: nil}
	dummy_head := current
	previous := current

	for _, rune := range polymer {
		node := &Node{unit: rune, next: nil, prev: previous}
		current.next = node
		previous = current
		current = current.next
	}
	// h := dummy_head.next
	// for h != nil {
	// 	fmt.Print(string(h.unit))
	// 	h = h.next
	// }
	// fmt.Println()
	dummy_head.next.prev = nil
	return dummy_head.next
}

func solution(filename string) int {
	polymer := parse(filename)
	polymer_list := create_list(polymer)
	return reduce(polymer_list)
}

func main() {
	fmt.Println(solution("./example1.txt"))
	// fmt.Println(solution("./input.txt"))
}
