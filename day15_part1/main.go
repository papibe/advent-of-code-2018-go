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
	for row, line := range board.grid {
		for col, cell := range line {
			if !cell.IsNature() {
				fmt.Printf("%d,%d\t%c(%d)\n", row, col, cell.kind, cell.hp)
			}
		}
	}
	fmt.Println()
}

type QElement struct {
	coords   Coord
	distance int
}

func (board Board) get_targets(row, col int) ([]Coord, int) {
	rows := len(board.grid)
	cols := len(board.grid[0])

	// BFS init
	start := Coord{row, col}
	current := board.grid[row][col]
	queue := []QElement{{start, 0}}
	visited := make(map[Coord]bool)
	visited[start] = true

	min_distance := 10000

	min_targets := make(map[int][]Coord)

	// BFS
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
						if distance < min_distance {
							min_distance = distance
						}
						_, distance_is_in_map := min_targets[distance]
						if distance_is_in_map {
							min_targets[distance] = append(min_targets[distance], pos)
						} else {
							min_targets[distance] = []Coord{pos}
						}
					}
					continue
				}

				queue = append(queue, QElement{coords, distance + 1})
				visited[coords] = true
			}

		}
	}
	return min_targets[min_distance], min_distance
}

func get_min_distances(board Board, row, col, max_distance int, destination Coord) map[Coord]int {
	rows := len(board.grid)
	cols := len(board.grid[0])

	distances := make(map[Coord]int)

	steps := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	for _, step := range steps {
		new_row := row + step[0]
		new_col := col + step[1]
		if new_row >= 0 && new_row < rows && new_col >= 0 && new_col < cols {
			if board.grid[new_row][new_col].kind == SPACE {
				distances[Coord{new_row, new_col}] = -1
			}
		}
	}

	// BFS
	start := Coord{destination.row, destination.col}
	visited := make(map[Coord]bool)
	visited[start] = true
	queue := []QElement{{start, 0}}

	for len(queue) > 0 {
		element := queue[0]
		queue = queue[1:]
		pos := element.coords
		current_distance := element.distance

		if current_distance > max_distance {
			continue
		}
		_, is_in_distances := distances[pos]
		if is_in_distances {
			distances[pos] = current_distance
		}
		ready := true
		for _, distance := range distances {
			if distance < 0 {
				ready = false
			}
		}
		if ready {
			return distances
		}

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
				queue = append(queue, QElement{coords, current_distance + 1})
				visited[coords] = true
			}
		}

	}
	trim_distances := make(map[Coord]int)
	for coord, distance := range distances {
		if distance >= 0 {
			trim_distances[coord] = distance
		}
	}
	return trim_distances
}

func (board Board) get_next_step(row, col, max_distance int, destination Coord) (Coord, error) {
	distances := get_min_distances(board, row, col, max_distance, destination)

	min_distance := 10000
	for _, distance := range distances {
		if distance < min_distance {
			min_distance = distance
		}
	}
	min_coords := make(map[Coord]int)
	for coord, distance := range distances {
		if distance == min_distance {
			min_coords[coord] = distance
		}
	}
	steps := [][]int{{-1, 0}, {0, -1}, {0, 1}, {1, 0}}

	for _, step := range steps {
		new_row := row + step[0]
		new_col := col + step[1]
		coord := Coord{new_row, new_col}
		_, in_map := min_coords[coord]
		if in_map {
			return coord, nil
		}
	}
	return Coord{-1, -1}, errors.New("not good")
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

func (board *Board) kill(row, col int) {
	cell := &(board.grid[row][col])
	cell.hp = 0
	cell.kind = SPACE
}

func (board Board) endCombat() (bool, int) {
	elves := 0
	goblins := 0
	for _, line := range board.grid {
		for _, cell := range line {
			if cell.IsElf() {
				elves += cell.hp
			}
			if cell.IsGoblin() {
				goblins += cell.hp
			}
		}
	}
	if elves == 0 {
		return true, goblins
	}
	if goblins == 0 {
		return true, elves
	}
	return false, 0
}

func parse(filename string) Board {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic("File not found")
	}

	board := Board{grid: [][]Cell{}}

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
	return board
}

func get_target(min_targets []Coord) Coord {
	min_row := 10000
	for _, coord := range min_targets {
		row := coord.row
		if row < min_row {
			min_row = row
		}
	}
	min_col := 10000
	var min_target Coord
	for _, coord := range min_targets {
		row := coord.row
		col := coord.col
		if row == min_row && row < min_col {
			min_col = col
			min_target = coord
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
				if cell.IsNature() || cell.moved {
					continue
				}
				// Combat ends condition
				end, total_hp := board.endCombat()
				if end {
					return (round - 1) * total_hp
				}
				min_targets, distance := board.get_targets(row, col)

				next_row := row
				next_col := col
				if len(min_targets) > 0 {
					target := get_target(min_targets)
					next_step, err := board.get_next_step(row, col, distance, target)
					if err == nil {
						board.Move(row, col, next_step)
						next_row = next_step.row
						next_col = next_step.col
					}
				}
				attack_target, err := get_attack_target(board, next_row, next_col)
				if err != nil {
					continue
				}
				hp := board.grid[attack_target.row][attack_target.col].hp - ATTACK_POWER
				if hp > 0 {
					board.grid[attack_target.row][attack_target.col].hp = hp
				} else {
					// kill
					(&board).kill(attack_target.row, attack_target.col)
				}
			}
		}
		board.resetMoves()
	}
}

func solution(filename string) int {
	board := parse(filename)
	return solve(board)
}

func main() {
	fmt.Println(solution("./input.txt"))
}
