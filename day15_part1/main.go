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

type Cell interface {
	Move()
}

type Board struct {
	grid [][]Cell
}

type Nature struct {
	kind rune
}

func (n Nature) Move() {

}

type Elf struct {
	hp  int
	row int
	col int
}

func (e Elf) Move() {

}

type Goblin struct {
	hp  int
	row int
	col int
}

func (g Goblin) Move() {

}

func parse(filename string) (Board, []Elf, []Goblin) {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic("File not found")
	}

	board := Board{grid: [][]Cell{}}
	elves := []Elf{}
	goblins := []Goblin{}

	for row, line := range strings.Split(strings.Trim(string(data), "\n"), "\n") {
		new_row := []Cell{}
		for col, char := range line {
			// switch char {
			// case WALL:
			// 	new_row = append(new_row, Nature{WALL})
			// case SPACE:
			// 	new_row = append(new_row, Nature{SPACE})
			// case ELF:
			// 	new_row = append(new_row, Elf{200, row, col})
			// case GOBLIN:
			// 	new_row = append(new_row, Goblin{200, row, col})
			// }
			var cell Cell
			switch char {
			case WALL:
				cell = Nature{WALL}
			case SPACE:
				cell = Nature{SPACE}
			case ELF:
				cell = Elf{200, row, col}
			case GOBLIN:
				cell = Goblin{200, row, col}
			}
			new_row = append(new_row, cell)

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
