package main

import (
	"fmt"
	"os"
	"strings"
)

const WALL = '#'
const SPACE = '.'
const ELF = 'E'
const GOBLIN = 'G'

type Board struct {
	grid [][]rune
}

type Elf struct {
	hp  int
	row int
	col int
}

type Goblin struct {
	hp  int
	row int
	col int
}

func parse(filename string) (Board, []Elf, []Goblin) {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic("File not found")
	}

	board := Board{grid: [][]rune{}}
	elves := []Elf{}
	goblins := []Goblin{}

	for row, line := range strings.Split(strings.Trim(string(data), "\n"), "\n") {
		new_row := []rune{}
		for col, char := range line {
			// form grid row
			if char == WALL {
				new_row = append(new_row, char)
			} else {
				new_row = append(new_row, SPACE)
			}
			// get list of players
			if char == ELF {
				elves = append(elves, Elf{200, row, col})
			}
			if char == GOBLIN {
				goblins = append(goblins, Goblin{200, row, col})
			}
		}
		board.grid = append(board.grid, new_row)
	}

	fmt.Println(board)
	fmt.Println(elves)
	fmt.Println(goblins)
	return board, elves, goblins
}

func solve(board Board, elves []Elf, goblins []Goblin) int {
	round := 0
	_ = round
	for {
		break
	}
	return 0
}

func solution(filename string) int {
	board, elves, goblins := parse(filename)
	return solve(board, elves, goblins)
}

func main() {
	fmt.Println(solution("./example1.txt"))
}
