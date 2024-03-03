package main

import (
	"fmt"
	"os"
	"strings"
)

const OPEN = '.'
const TREES = '|'
const LUMBERYARD = '#'

var values = map[rune]int64{
	'.': 0,
	'|': 1,
	'#': 2,
}

func print_area(area [][]rune) {
	for _, line := range area {
		for _, char := range line {
			fmt.Print(string(char))
		}
		fmt.Println()
	}
}

func parse(filename string) [][]rune {
	raw_data, err := os.ReadFile(filename)
	if err != nil {
		panic("File not found")
	}
	data := strings.Trim(string(raw_data), "\n")

	area := [][]rune{}
	for _, line := range strings.Split(data, "\n") {
		row := []rune(line)
		area = append(area, row)
	}

	return area
}

func hash(area [][]rune) string {
	var sb strings.Builder
	for _, line := range area {
		for _, char := range line {
			sb.WriteRune(char)
		}
	}
	return sb.String()
}

func deep_copy(area [][]rune) [][]rune {
	new_area := [][]rune{}
	for _, line := range area {
		new_row := []rune{}
		for _, char := range line {
			new_row = append(new_row, char)
		}
		new_area = append(new_area, new_row)
	}
	return new_area
}

func get_neigbhors(area [][]rune, row, col int) (int, int, int) {
	var steps = [][]int{
		{-1, -1},
		{-1, 0},
		{-1, 1},

		{0, -1},
		{0, 1},

		{1, -1},
		{1, 0},
		{1, 1},
	}

	opens := 0
	trees := 0
	lumberyards := 0

	for _, step := range steps {
		new_row := row + step[0]
		new_col := col + step[1]
		if new_row >= 0 && new_row < len(area) && new_col >= 0 && new_col < len(area[0]) {
			item := area[new_row][new_col]
			switch item {
			case OPEN:
				opens++
			case TREES:
				trees++
			case LUMBERYARD:
				lumberyards++
			}
		}
	}

	return opens, trees, lumberyards
}

func minute_cycle(area [][]rune) [][]rune {
	new_area := deep_copy(area)
	for row, line := range area {
		for col, char := range line {
			opens, trees, lumberyards := get_neigbhors(area, row, col)
			_ = opens
			switch char {
			case OPEN:
				if trees >= 3 {
					new_area[row][col] = TREES
				}
			case TREES:
				if lumberyards >= 3 {
					new_area[row][col] = LUMBERYARD
				}
			case LUMBERYARD:
				if lumberyards >= 1 && trees >= 1 {
					new_area[row][col] = LUMBERYARD
				} else {
					new_area[row][col] = OPEN
				}
			}
		}
	}
	return new_area
}

func solve(area [][]rune, minutes int) int {
	current_minute := 0
	repeated_minute := 0
	states := make(map[string]int)
	for minute := 0; minute < minutes; minute++ {

		current_hash := hash(area)
		past_minute, is_repeated_state := states[current_hash]

		if is_repeated_state {
			current_minute = minute
			repeated_minute = past_minute
			break
		}

		states[current_hash] = minute

		// Process rules
		new_area := minute_cycle(area)
		area = new_area
	}

	// calculate what is the cycle and how much is missing
	cycle := current_minute - repeated_minute
	missing := (minutes - repeated_minute) % cycle

	for minute := 0; minute < missing; minute++ {
		new_area := minute_cycle(area)
		area = new_area
	}

	opens := 0
	trees := 0
	lumberyards := 0
	for _, line := range area {
		for _, char := range line {
			switch char {
			case OPEN:
				opens++
			case TREES:
				trees++
			case LUMBERYARD:
				lumberyards++
			}
		}
	}
	return trees * lumberyards
}

func solution(filename string, minutes int) int {
	area := parse(filename)
	return solve(area, minutes)
}

func main() {
	fmt.Println(solution("./input.txt", 1000000000)) // 210796
}
