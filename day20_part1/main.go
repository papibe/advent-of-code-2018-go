package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

type Coord struct {
	row int
	col int
}

type Queue struct {
	row      int
	col      int
	distance int
}

func parse(filename string) string {
	raw_data, err := os.ReadFile(filename)
	if err != nil {
		panic("File not found")
	}
	regex := strings.Trim(string(raw_data), "\n")

	return regex
}

func draw(doors map[Coord]rune, rooms, walls map[Coord]bool) string {
	min_row := math.MaxInt
	max_row := math.MinInt
	min_col := math.MaxInt
	max_col := math.MinInt

	for door := range doors {
		min_row = min(min_row, door.row)
		min_col = min(min_col, door.col)
		max_row = max(max_row, door.row)
		max_col = max(max_col, door.col)
	}
	for room := range rooms {
		min_row = min(min_row, room.row)
		min_col = min(min_col, room.col)
		max_row = max(max_row, room.row)
		max_col = max(max_col, room.col)
	}
	for wall := range walls {
		min_row = min(min_row, wall.row)
		min_col = min(min_col, wall.col)
		max_row = max(max_row, wall.row)
		max_col = max(max_col, wall.col)
	}
	output := []string{}
	output = append(output, "\n")

	for row := min_row; row <= max_row; row++ {
		for col := min_col; col <= max_col; col++ {
			if row == 0 && col == 0 {
				// fmt.Print("X")
				output = append(output, "X")
				continue
			}
			_, is_room := rooms[Coord{row, col}]
			door, is_door := doors[Coord{row, col}]
			_, is_wall := walls[Coord{row, col}]
			if is_room {
				// fmt.Print(".")
				output = append(output, ".")
			} else if is_door {
				// fmt.Print(string(door))
				output = append(output, string(door))

			} else if is_wall {
				// fmt.Print("#")
				output = append(output, "#")
			} else {
				walls[Coord{row, col}] = true
				// fmt.Print("#")
				output = append(output, "#")
			}
		}
		// fmt.Println()
		output = append(output, "\n")
	}
	return strings.Join(output, "")
}

func read_regex(regex string, index, row, col int) (map[Coord]rune, map[Coord]bool, map[Coord]bool) {
	doors := make(map[Coord]rune)
	rooms := make(map[Coord]bool)
	walls := make(map[Coord]bool)

	stack := []Coord{}

outer:
	for {
		// fmt.Print("current ", row, col, " ")
		switch regex[index] {
		case '^':
			// fmt.Println("saving", row, col)
			stack = append(stack, Coord{row, col})
		case '$':
			_ = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if len(stack) != 0 {
				panic("stack is not empty")
			}
			break outer
		case 'N':
			// fmt.Println("N")
			rooms[Coord{row, col}] = true         // current room
			rooms[Coord{row - 2, col}] = true     // next room
			doors[Coord{row - 1, col}] = '-'      // pass through door
			walls[Coord{row - 1, col - 1}] = true // wall sustain door
			walls[Coord{row - 1, col + 1}] = true // wall sustain door
			row -= 2
		case 'E':
			// fmt.Println("E")
			rooms[Coord{row, col}] = true         // current room
			rooms[Coord{row, col + 2}] = true     // next room
			doors[Coord{row, col + 1}] = '|'      // pas through door
			walls[Coord{row - 1, col + 1}] = true // wall sustain door
			walls[Coord{row + 1, col + 1}] = true // wall sustain door
			col += 2
		case 'S':
			// fmt.Println("S")
			rooms[Coord{row, col}] = true         // current room
			rooms[Coord{row + 2, col}] = true     // next room
			doors[Coord{row + 1, col}] = '-'      // pass through door
			walls[Coord{row + 1, col - 1}] = true // wall sustain door
			walls[Coord{row + 1, col + 1}] = true // wall sustain door
			row += 2

		case 'W':
			// fmt.Println("W")
			rooms[Coord{row, col}] = true         // current room
			rooms[Coord{row, col - 2}] = true     // next room
			doors[Coord{row, col - 1}] = '|'      // pas through door
			walls[Coord{row - 1, col - 1}] = true // wall sustain door
			walls[Coord{row + 1, col - 1}] = true // wall sustain door
			col -= 2

		case '(':
			// break outer
			// fmt.Println("( saving", row, col)
			stack = append(stack, Coord{row, col})
		case ')':
			// check if prev == '|'
			coord := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			row, col = coord.row, coord.col
			// fmt.Println(") popping", row, col)
		case '|':
			coord := stack[len(stack)-1]
			// stack = stack[:len(stack)-1]
			row, col = coord.row, coord.col
			// fmt.Println("| gettop", row, col)

		}
		index += 1
	}

	return doors, rooms, walls
}

func get_max_distance(doors map[Coord]rune, rooms, walls map[Coord]bool) int {
	// BFS init
	queue := []Queue{{0, 0, 0}}
	visited := make(map[Coord]bool)
	max_distance := 0

	// BFS
	for len(queue) > 0 {
		item := queue[len(queue)-1]
		queue = queue[:len(queue)-1]
		row, col, distance := item.row, item.col, item.distance
		_, is_room := rooms[Coord{row, col}]
		if is_room {
			max_distance = max(max_distance, distance)
		}

		steps := []Coord{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
		for _, step := range steps {
			new_row := row + step.row
			new_col := col + step.col
			new_coord := Coord{new_row, new_col}

			// identify what kind of cell it is
			_, is_room := rooms[new_coord]
			_, is_door := doors[new_coord]
			_, is_wall := walls[new_coord]
			_, is_visited := visited[new_coord]

			if is_wall || is_visited {
				continue
			}
			if !is_door && !is_room {
				continue
			}
			var new_distance = distance
			if is_room {
				new_distance += 1
			}
			queue = append(queue, Queue{new_row, new_col, new_distance})
			visited[new_coord] = true
		}

	}
	return max_distance
}

func solution(filename string) int {
	regex := parse(filename)
	doors, rooms, walls := read_regex(regex, 0, 0, 0)
	// draw(doors, rooms, walls)
	return get_max_distance(doors, rooms, walls)
}

func main() {
	fmt.Println(solution("example1.txt")) // 3
	fmt.Println(solution("example2.txt")) // 10
	fmt.Println(solution("example3.txt")) // 18
	fmt.Println(solution("example4.txt")) // 23
	fmt.Println(solution("example5.txt")) // 31

	fmt.Println(solution("input.txt")) //

	// TODO
	// - make tests for all examples
}
