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

// type Q2Element struct {
// 	coords   Coord
// 	distance int
// 	path     []Coord
// 	visited  map[Coord]bool
// }

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

func get_min_distance(board Board, row, col, max_distance int, destination Coord) int {
	rows := len(board.grid)
	cols := len(board.grid[0])

	// BFS init
	start := Coord{row, col}
	visited := make(map[Coord]bool)
	visited[start] = true
	queue := []QElement{{start, 0}}

	for len(queue) > 0 {
		element := queue[0]
		queue = queue[1:]
		pos := element.coords
		current_distance := element.distance
		fmt.Println("pop\t", pos, current_distance, "md", max_distance)

		if current_distance > max_distance {
			continue
		}
		if pos.row == destination.row && pos.col == destination.col {
			return current_distance
		}
		steps := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
		for _, step := range steps {
			new_row := pos.row + step[0]
			new_col := pos.col + step[1]
			fmt.Println("1re\t", new_row, new_col)

			if new_row >= 0 && new_row < rows && new_col >= 0 && new_col < cols {
				new_cell := board.grid[new_row][new_col]
				coords := Coord{new_row, new_col}
				fmt.Println("2re\t\t", coords)

				_, is_visited := visited[coords]
				if is_visited {
					continue
				}
				if new_cell.kind != SPACE {
					continue
				}
				fmt.Println("pus\t\t\t", coords)
				queue = append(queue, QElement{coords, current_distance + 1})
				visited[coords] = true
			}
		}

	}
	return -1
}

func (board Board) get_next_step(row, col, max_distance int, destination Coord) (Coord, error) {
	rows := len(board.grid)
	cols := len(board.grid[0])

	distances := make(map[Coord]int)

	// steps := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}, {0, 0}}
	steps := [][]int{{-1, 0}, {0, -1}, {0, 1}, {1, 0}}
	for _, step := range steps {
		new_row := row + step[0]
		new_col := col + step[1]
		if new_row >= 0 && new_row < rows && new_col >= 0 && new_col < cols {
			if board.grid[new_row][new_col].kind == SPACE {
				distance := get_min_distance(board, new_row, new_col, max_distance, destination)
				fmt.Println("distance", new_row, new_col, distance)
				if distance >= 0 {
					distances[Coord{new_row, new_col}] = get_min_distance(board, new_row, new_col, max_distance, destination)
				}
			}
		}
	}
	fmt.Println("distances", distances)
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
	for _, step := range steps {
		new_row := row + step[0]
		new_col := col + step[1]
		coord := Coord{new_row, new_col}
		_, in_map := min_coords[coord]
		if in_map {
			return coord, nil
		}
	}
	return Coord{-1, -1}, errors.New("well...")
}

func (board Board) get_next_step_(row, col, distance int, destination Coord) (Coord, error) {
	// if row == destination.row && col == destination.col {
	// 	return Coord{row, col}, nil
	// }
	// fmt.Println(row, col, "->", destination.row, destination.col)
	rows := len(board.grid)
	cols := len(board.grid[0])

	// BFS init
	start := Coord{row, col}
	visited := make(map[Coord]bool)
	visited[start] = true
	// queue := []Q2Element{{start, 0, []Coord{}, initial_visited}}

	paths := [][]Coord{}
	path := []Coord{{row, col}}

	var get_paths func(row, col, distance int)

	get_paths = func(row, col, distance int) {
		if distance < 0 {
			return
		}
		if row == destination.row && col == destination.col {
			new_path := make([]Coord, len(path))
			_ = copy(new_path, path)
			paths = append(paths, new_path)
			return
		}

		steps := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
		for _, step := range steps {
			new_row := row + step[0]
			new_col := col + step[1]

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
				path = append(path, coords)
				visited[coords] = true
				get_paths(new_row, new_col, distance-1)
				_, path = path[len(path)-1], path[:len(path)-1]
				delete(visited, coords)

			}
		}
	}
	get_paths(start.row, start.col, distance)

	// fmt.Println("start", start)
	// fmt.Println("destination", destination)
	// fmt.Println("max distance", distance)

	// for len(queue) > 0 {

	// 	fmt.Println("len(queue)", len(queue), distance)

	// 	element := queue[0]
	// 	queue = queue[1:]
	// 	pos := element.coords
	// 	path := element.path
	// 	visited := element.visited
	// 	current_distance := element.distance

	// 	// fmt.Println(path)

	// 	if current_distance > distance {
	// 		continue
	// 	}

	// 	if pos.row == destination.row && pos.col == destination.col {
	// 		// fmt.Println("arrived!")
	// 		if len(path) > 0 {
	// 			paths = append(paths, path)
	// 		}
	// 		continue
	// 	}

	// 	steps := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	// 	for _, step := range steps {
	// 		new_row := pos.row + step[0]
	// 		new_col := pos.col + step[1]

	// 		if new_row >= 0 && new_row < rows && new_col >= 0 && new_col < cols {
	// 			new_cell := board.grid[new_row][new_col]
	// 			coords := Coord{new_row, new_col}

	// 			_, is_visited := visited[coords]
	// 			if is_visited {
	// 				continue
	// 			}
	// 			if new_cell.kind != SPACE {
	// 				continue
	// 			}
	// 			new_path := make([]Coord, len(path))
	// 			_ = copy(new_path, path)
	// 			new_path = append(new_path, coords)

	// 			new_visited := deepcopy(visited) // copied?
	// 			new_visited[coords] = true
	// 			queue = append(queue, Q2Element{coords, current_distance + 1, new_path, new_visited})
	// 		}

	// 	}
	// }

	// fmt.Println("---------------------------------")
	// fmt.Println(paths)
	// fmt.Println("---------------------------------")

	if len(paths) == 0 {
		return Coord{-1, -1}, errors.New("no where to go")
	}

	min_row := 10000
	for _, path := range paths {
		if len(path) == 0 {
			fmt.Println("paths", paths)
		}
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

// func deepcopy(original map[Coord]bool) map[Coord]bool {
// 	new_copy := make(map[Coord]bool)
// 	for key, value := range original {
// 		new_copy[key] = value
// 	}
// 	return new_copy
// }

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
	// out:
	for {
		round += 1
		for row := 0; row < len(board.grid); row++ {
			for col := 0; col < len(board.grid[0]); col++ {
				cell := board.grid[row][col]
				if cell.IsNature() {
					// fmt.Println(row, col, "nature")
					continue
				}
				if cell.moved {
					// fmt.Println(row, col, "already moved")
					continue
				}
				end, total_hp := board.endCombat()
				if end {
					return (round - 1) * total_hp
				}
				fmt.Println(row, col)
				targets := board.get_target_in_range(row, col)
				// Combat ends condition

				fmt.Println(row, col, "targets", targets)
				min_targets := get_min_targets(targets)
				fmt.Println("min_targets", min_targets)
				next_row := row
				next_col := col
				if len(min_targets) > 0 {
					target := get_target(min_targets)
					fmt.Println("target", target)
					distance := target.distance
					// fmt.Println(row, col, "target", target)
					next_step, err := board.get_next_step(row, col, distance, target.coords)
					fmt.Println("next_step", next_step)
					if err == nil {
						// fmt.Println(row, col, "next step", next_step)
						board.Move(row, col, next_step)
						board.Print()
						next_row = next_step.row
						next_col = next_step.col
						fmt.Println(row, col, "-move->", next_row, next_col)
					} else {
						fmt.Println(row, col, "no move")
						// continue
					}
				} else {
					fmt.Println(row, col, "no move")
				}
				attack_target, err := get_attack_target(board, next_row, next_col)
				if err != nil {
					board.Print()
					continue
				}
				fmt.Println(next_row, next_col, "attack", attack_target)
				hp := board.grid[attack_target.row][attack_target.col].hp - ATTACK_POWER
				if hp > 0 {
					board.grid[attack_target.row][attack_target.col].hp = hp
				} else {
					// kill
					(&board).kill(attack_target.row, attack_target.col)
					fmt.Println("KILLING", attack_target.row, attack_target.col)
				}
				board.Print()
			}
			// break
		}
		// board.Print()
		board.resetMoves()
		// if round >= 100000 {
		// 	break
		// }
		fmt.Println("-------------")
		fmt.Println(round)
		fmt.Println("-------------")
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
	// fmt.Println("sol-3", solution("./example3.txt"))
	// fmt.Println("sol-4", solution("./example4.txt"))
	// fmt.Println("sol-5", solution("./example5.txt"))
	// fmt.Println("sol-6", solution("./example6.txt"))
	// fmt.Println("sol-7", solution("./example7.txt"))
	// fmt.Println("sol-8", solution("./example8.txt"))
	fmt.Println("sol-i", solution("./input.txt"))
}
