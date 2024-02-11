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

type Coord struct {
	row int
	col int
}

type Cell struct {
	kind  rune
	hp    int
	coord Coord
}

func (cell Cell) isEnemy(kind rune) bool {
	if cell.kind != kind {
		return true
	} else {
		return false
	}
}

func (cell Cell) IsNature() bool {
	if cell.kind == WALL || cell.kind == SPACE {
		return true
	} else {
		return false
	}
}

func (cell Cell) IsElf() bool {
	if cell.kind == ELF {
		return true
	} else {
		return false
	}
}

func (cell Cell) IsGoblin() bool {
	if cell.kind == GOBLIN {
		return true
	} else {
		return false
	}
}

type Board struct {
	grid [][]Cell
}

func (board Board) Print() {
	for _, line := range board.grid {
		for _, cell := range line {
			fmt.Print(string(cell.kind))
		}
		fmt.Println()
	}
}

type QElement struct {
	coords   Coord
	distance int
}

func (board Board) get_target_in_range(row, col int) []QElement {
	rows := len(board.grid)
	cols := len(board.grid[0])

	// BFS init
	start := Coord{row, col}
	current := board.grid[row][col]
	queue := []QElement{QElement{start, 0}}
	visited := make(map[Coord]bool)
	visited[start] = true

	targets := []QElement{}

	for len(queue) > 0 {
		element := queue[0]
		pos := element.coords
		distance := element.distance
		queue = queue[1:]

		steps := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
		for _, step := range steps {
			new_row := pos.row + step[0]
			new_col := pos.col + step[1]

			if new_row >= 0 && new_row < rows && new_col >= 0 && new_col < cols {
				new_player := board.grid[new_row][new_col]
				coords := Coord{new_row, new_col}

				_, is_visited := visited[coords]
				if is_visited {
					continue
				}
				if new_player.kind == WALL {
					continue
				}

				if new_player.IsElf() || new_player.IsGoblin() {
					if current.isEnemy(new_player.kind) {
						targets = append(targets, QElement{pos, distance})
					}
					continue
				}

				queue = append(queue, QElement{coords, distance + 1})
				visited[coords] = true
			}

		}
	}

	return targets
}

func parse(filename string) Board {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic("File not found")
	}

	board := Board{grid: [][]Cell{}}
	// elves := []Elf{}
	// goblins := []Goblin{}

	for row, line := range strings.Split(strings.Trim(string(data), "\n"), "\n") {
		new_row := []Cell{}
		for col, char := range line {
			cell := Cell{kind: char, coord: Coord{row, col}, hp: 0}
			if char == ELF || char == GOBLIN {
				cell.hp = 200
			}

			new_row = append(new_row, cell)
		}
		board.grid = append(board.grid, new_row)
	}
	board.Print()
	return board
}

func get_min_targets(targets []QElement) []QElement {
	// get min distance
	min_distance := 10000
	for _, target := range targets {
		distance := target.distance
		if distance < min_distance {
			min_distance = distance
		}
	}
	// get all points with min distance
	min_targets := []QElement{}
	for _, target := range targets {
		distance := target.distance
		if distance <= min_distance {
			min_targets = append(min_targets, target)
		}
	}
	return min_targets
}

func get_target(min_targets []QElement) Coord {
	min_row := 10000
	for _, target := range min_targets {
		row := target.coords.row
		if row < min_row {
			min_row = row
		}
	}
	min_col := 10000
	for _, target := range min_targets {
		row := target.coords.row
		col := target.coords.col
		if row == min_row && row < min_col {
			min_col = col
		}
	}
	return Coord{min_row, min_col}
}

func solve(board Board) int {
	round := 0
	_ = round
	for {
		for row, line := range board.grid {
			for col, cell := range line {
				if cell.IsNature() {
					continue
				}
				_ = row
				_ = col
				targets := board.get_target_in_range(row, col)
				fmt.Println(row, col, "targets", targets)
				min_targets := get_min_targets(targets)
				fmt.Println("min_targets", min_targets)
				if len(min_targets) > 0 {
					target := get_target(min_targets)
					fmt.Println("target", target)
				}
				// break
			}
			// break
		}
		break
	}
	return 0
}

func solution(filename string) int {
	board := parse(filename)
	return solve(board)
}

func main() {
	fmt.Println(solution("./example0.txt"))
	fmt.Println(solution("./example1.txt"))
}
