package main

import (
	"fmt"
	"os"
	"strings"
)

const PLANT = '#'
const NO_PLANT = '.'

func parse(filename string) ([]rune, map[string]rune) {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic("File not found")
	}

	lines := strings.Split(strings.Trim(string(data), "\n"), "\n\n")
	initial_state_str := lines[0]
	rules_str := lines[1]

	state := []rune(strings.Split(initial_state_str, ": ")[1])

	rules := make(map[string]rune)
	for _, line := range strings.Split(rules_str, "\n") {
		split_line := strings.Split(line, " => ")
		pattern := split_line[0]
		next_gen := []rune(split_line[1])[0]
		rules[pattern] = next_gen
	}

	return state, rules
}

func solve(state []rune, rules map[string]rune, generations int) int {

	current_shift := 0
	completed_cycles := 0
	for i := 0; i < generations; i++ {
		// save previous state
		prev_state := make([]rune, len(state))
		copy(prev_state, state)

		// add 3 at the beginning and at the end
		current_shift += 3
		state = append([]rune{'.', '.', '.'}, state...)
		state = append(state, []rune{'.', '.', '.'}...)

		// create empty next_state
		next_state := make([]rune, len(state))
		for j := 0; j < len(state); j++ {
			next_state[j] = '.'
		}

		for index := 0; index < len(state)-4; index++ {
			next_gen, ok := rules[string(state[index:index+5])]
			if ok {
				next_state[index+2] = next_gen
			}
		}
		state = next_state

		// remove unnecessary '.'
		start_index := 0
		for state[start_index] == NO_PLANT {
			start_index += 1
		}
		end_index := len(state) - 1
		for state[end_index] == NO_PLANT {
			end_index -= 1
		}
		state = state[start_index : end_index+1]
		current_shift -= start_index
		completed_cycles += 1

		if string(prev_state) == string(state) {
			break
		}
	}
	current_shift = current_shift - (generations - completed_cycles)
	index_sum := 0
	for index, char := range state {
		if char == PLANT {
			index_sum += index - current_shift
		}
	}
	return index_sum
}

func solution(filename string, generations int) int {
	state, rules := parse(filename)
	return solve(state, rules, generations)
}

func main() {
	fmt.Println(solution("./example.txt", 50000000000)) // 999999999374
	fmt.Println(solution("./input.txt", 50000000000))   // 2500000000695
}
