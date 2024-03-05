package main

import (
	"fmt"
	"os"
	"strings"
)

const OPEN = '.'
const TREE = '|'
const LUMBERYARD = '#'

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
			case TREE:
				trees++
			case LUMBERYARD:
				lumberyards++
			}
		}
	}

	return opens, trees, lumberyards
}

func solve(area [][]rune) int {
	for minute := 1; minute <= 10; minute++ {

		// Process rules
		new_area := deep_copy(area)
		for row, line := range area {
			for col, char := range line {
				opens, trees, lumberyards := get_neigbhors(area, row, col)
				_ = opens
				switch char {
				case OPEN:
					if trees >= 3 {
						new_area[row][col] = TREE
					}
				case TREE:
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
			case TREE:
				trees++
			case LUMBERYARD:
				lumberyards++
			}
		}
	}
	return trees * lumberyards
}

func solution(filename string) int {
	area := parse(filename)
	// print_area(area)
	return solve(area)
}

func main() {
	fmt.Println(solution("./example.txt")) // 1147
	fmt.Println(solution("./input.txt"))   // 606416
}
