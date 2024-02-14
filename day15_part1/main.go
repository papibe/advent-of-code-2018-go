package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

const WALL = '#'
const SPACE = '.'
const ELF = 'E'
const GOBLIN = 'G'
const ATTACK_POWER = 3

type Coord struct {
	row int
	col int
}

type Cell struct {
	kind  rune
	hp    int
	moved bool
}

func (cell Cell) isEnemy(kind rune) bool {
	if kind == WALL || kind == SPACE {
		return false
	}

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
	for _, line := range board.grid {
		for _, cell := range line {
			if !cell.IsNature() {
				fmt.Printf("%c(%d)\n", cell.kind, cell.hp)
			}
		}
	}
	fmt.Println()
}

type QElement struct {
	coords   Coord
	distance int
}

type Q2Element struct {
	coords   Coord
	distance int
	path     []Coord
	visited  map[Coord]bool
}

func (board Board) get_target_in_range(row, col int) []QElement {
	rows := len(board.grid)
	cols := len(board.grid[0])

	// BFS init
	start := Coord{row, col}
	current := board.grid[row][col]
	queue := []QElement{{start, 0}}
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

func (board Board) get_next_step(row, col, distance int, destination Coord) (Coord, error) {
	rows := len(board.grid)
	cols := len(board.grid[0])

	// BFS init
	start := Coord{row, col}
	initial_visited := make(map[Coord]bool)
	initial_visited[start] = true
	queue := []Q2Element{{start, 0, []Coord{}, initial_visited}}

	paths := [][]Coord{}

	// fmt.Println("start", start)
	// fmt.Println("destination", destination)
	// fmt.Println("max distance", distance)

	for len(queue) > 0 {

		element := queue[0]
		queue = queue[1:]
		pos := element.coords
		path := element.path
		visited := element.visited
		current_distance := element.distance

		// fmt.Println(path)

		if current_distance > distance {
			continue
		}

		if pos.row == destination.row && pos.col == destination.col {
			// fmt.Println("arrived!")
			if len(path) > 0 {
				paths = append(paths, path)
			}
			continue
		}

		steps := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
		for _, step := range steps {
			new_row := pos.row + step[0]
			new_col := pos.col + step[1]

			if new_row >= 0 && new_row < rows && new_col >= 0 && new_col < cols {
				new_cell := board.grid[new_row][new_col]
				coords := Coord{new_row, new_col}

				_, is_visited := visited[coords]
				if is_visited {
					continue
				}
				if new_cell.kind != SPACE {
					continue
				}
				new_path := make([]Coord, len(path))
				_ = copy(new_path, path)
				new_path = append(new_path, coords)

				new_visited := deepcopy(visited) // copied?
				new_visited[coords] = true
				queue = append(queue, Q2Element{coords, current_distance + 1, new_path, new_visited})
			}

		}
	}

	// fmt.Println("---------------------------------")
	// fmt.Println(paths)
	// fmt.Println("---------------------------------")

	if len(paths) == 0 {
		return Coord{-1, -1}, errors.New("no where to go")
	}

	min_row := 10000
	for _, path := range paths {
		// if len(path) == 0 {
		// 	fmt.Println("paths", paths)
		// }
		row := path[0].row
		if row < min_row {
			min_row = row
		}
	}
	min_col := 10000
	for _, path := range paths {
		row := path[0].row
		col := path[0].col
		if row == min_row && row < min_col {
			min_col = col
		}
	}
	return Coord{min_row, min_col}, nil
}

func (board *Board) Move(row, col int, next_coord Coord) {
	current := &(board.grid[row][col])
	next := &(board.grid[next_coord.row][next_coord.col])
	kind := current.kind
	hp := current.hp

	next.kind = kind
	next.hp = hp
	next.moved = true

	current.kind = SPACE
	current.hp = 0
}

func (board *Board) resetMoves() {
	for row := 0; row < len(board.grid); row++ {
		for col := 0; col < len(board.grid[0]); col++ {
			current := &(board.grid[row][col])
			current.moved = false
		}
	}
}

func (board *Board) Attack(row, col int) {
	cell := board.grid[row][col]
	cell.hp = cell.hp - ATTACK_POWER
}

func deepcopy(original map[Coord]bool) map[Coord]bool {
	new_copy := make(map[Coord]bool)
	for key, value := range original {
		new_copy[key] = value
	}
	return new_copy
}

func parse(filename string) Board {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic("File not found")
	}

	board := Board{grid: [][]Cell{}}
	// elves := []Elf{}
	// goblins := []Goblin{}

	for _, line := range strings.Split(strings.Trim(string(data), "\n"), "\n") {
		new_row := []Cell{}
		for _, char := range line {
			cell := Cell{kind: char, hp: 0, moved: false}
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

func get_target(min_targets []QElement) QElement {
	min_row := 10000
	for _, target := range min_targets {
		row := target.coords.row
		if row < min_row {
			min_row = row
		}
	}
	min_col := 10000
	var min_target QElement
	for _, target := range min_targets {
		row := target.coords.row
		col := target.coords.col
		if row == min_row && row < min_col {
			min_col = col
			min_target = target
		}
	}
	return min_target
}

func get_attack_target(board Board, row, col int) (Coord, error) {
	attacker := board.grid[row][col]
	targets := []Coord{}
	steps := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	for _, step := range steps {
		new_row := row + step[0]
		new_col := col + step[1]
		target := board.grid[new_row][new_col]
		if attacker.isEnemy(target.kind) {
			targets = append(targets, Coord{new_row, new_col})
		}
	}
	if len(targets) == 0 {
		return Coord{-1, -1}, errors.New("no target in sight")
	}
	targets_min_hp := 10000
	for _, coord := range targets {
		target := board.grid[coord.row][coord.col]
		if target.hp < targets_min_hp {
			targets_min_hp = target.hp
		}
	}
	min_targets := []Coord{}
	for _, coord := range targets {
		target := board.grid[coord.row][coord.col]
		if target.hp == targets_min_hp {
			min_targets = append(min_targets, coord)
		}
	}
	if len(min_targets) == 1 {
		coord := min_targets[0]
		return Coord{coord.row, coord.col}, nil
	}
	// get read order target
	min_row := 10000
	for _, coords := range min_targets {
		row := coords.row
		if row < min_row {
			min_row = row
		}
	}
	min_col := 10000
	for _, coord := range min_targets {
		row := coord.row
		col := coord.col
		if row == min_row && row < min_col {
			min_col = col
		}
	}
	return Coord{min_row, min_col}, nil
}

func solve(board Board) int {
	round := 0
	for {
		round += 1
		for row := 0; row < len(board.grid); row++ {
			for col := 0; col < len(board.grid[0]); col++ {
				cell := board.grid[row][col]
				if cell.IsNature() {
					continue
				}
				if cell.moved {
					continue
				}
				// fmt.Println(row, col)
				targets := board.get_target_in_range(row, col)
				// fmt.Println(row, col, "targets", targets)
				min_targets := get_min_targets(targets)
				// fmt.Println("min_targets", min_targets)
				if len(min_targets) > 0 {
					target := get_target(min_targets)
					distance := target.distance
					// fmt.Println(row, col, "target", target)
					next_step, err := board.get_next_step(row, col, distance, target.coords)
					if err != nil {
						continue
					}
					// fmt.Println(row, col, "next step", next_step)
					board.Move(row, col, next_step)
					// board.Print()
				}
				attack_target, err := get_attack_target(board, row, col)
				if err != nil {
					continue
				}
				fmt.Println(row, col, "attack", attack_target)
				// (&board).Attack(attack_target.row, attack_target.col)
				attack_cell := board.grid[attack_target.row][attack_target.col]
				attack_cell.hp -= ATTACK_POWER
				// (&board.grid[attack_target.row][attack_target.col]).hp -= ATTACK_POWER
				// TODO: kill unit
				// break
			}
			// break
		}
		board.Print()
		board.resetMoves()
		if round >= 3 {

			break
		}
	}
	return 0
}

func solution(filename string) int {
	board := parse(filename)
	return solve(board)
}

func main() {
	// fmt.Println(solution("./example0.txt"))
	// fmt.Println(solution("./example1.txt"))
	// fmt.Println(solution("./example2.txt"))
	fmt.Println(solution("./example3.txt"))
	// fmt.Println(solution("./input.txt"))
}
