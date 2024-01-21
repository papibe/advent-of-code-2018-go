package main

import (
	"fmt"
	"os"
	"strings"
)

var CARTS_SYMBOLS = []rune{'>', '<', '^', 'v'}

// type Direction [2]int

type Direction struct {
	row int
	col int
}

// Directions
var RIGHT = Direction{0, 1}
var LEFT = Direction{0, -1}
var UP = Direction{-1, 0}
var DOWN = Direction{1, 0}

var translate_to_symbol = map[Direction]rune{
	RIGHT: '>',
	LEFT:  '<',
	UP:    '^',
	DOWN:  'v',
}

var translate_to_direction = map[rune]Direction{
	'>': RIGHT,
	'<': LEFT,
	'^': UP,
	'v': DOWN,
}

var translate_to_grid = map[rune]rune{
	'>': '-',
	'<': '-',
	'^': '|',
	'v': '|',
}

var next_turn = map[int]map[Direction]Direction{
	0: {
		RIGHT: UP,
		LEFT:  DOWN,
		UP:    LEFT,
		DOWN:  RIGHT,
	},
	1: {
		RIGHT: RIGHT,
		LEFT:  LEFT,
		UP:    UP,
		DOWN:  DOWN,
	},
	2: {
		RIGHT: DOWN,
		LEFT:  UP,
		UP:    RIGHT,
		DOWN:  LEFT,
	},
}

var simple_next_dir = map[Direction]map[rune]Direction{
	RIGHT: {
		'-':  RIGHT,
		'\\': DOWN,
		'/':  UP,
	},
	LEFT: {
		'-':  LEFT,
		'\\': UP,
		'/':  DOWN,
	},
	DOWN: {
		'|':  DOWN,
		'\\': RIGHT,
		'/':  LEFT,
	},
	UP: {
		'|':  UP,
		'\\': LEFT,
		'/':  RIGHT,
	},
}

type Cart struct {
	cart_id int
	pos_row int
	pos_col int
	dir_row int
	dir_col int
	counter int
}

func (cart *Cart) Move(grid [][]rune) Position {
	cart.pos_row += cart.dir_row
	cart.pos_col += cart.dir_col
	grid_cell := grid[cart.pos_row][cart.pos_col]
	if grid_cell == ' ' {
		panic("what!")
	}
	if grid_cell != '+' {
		next_dir := simple_next_dir[Direction{cart.dir_row, cart.dir_col}][grid_cell]
		// fmt.Println(cart, next_dir, string(grid_location))
		if next_dir.row == 0 && next_dir.col == 0 {
			panic("wrong direction")
		}
		cart.dir_row = next_dir.row
		cart.dir_col = next_dir.col
	} else {
		next_dir := next_turn[cart.counter][Direction{cart.dir_row, cart.dir_col}]
		if next_dir.row == 0 && next_dir.col == 0 {
			panic("wrong direction")
		}
		cart.dir_row = next_dir.row
		cart.dir_col = next_dir.col
		cart.counter = (cart.counter + 1) % 3
	}

	return Position{row: cart.pos_row, col: cart.pos_col}
}

type Position struct {
	row int
	col int
}

func is_cart(char rune) bool {
	for _, item := range CARTS_SYMBOLS {
		if item == char {
			return true
		}
	}
	return false
}

func draw(grid [][]rune, carts map[Position]Cart) {
	for row, line := range grid {
		for col, cell := range line {
			cart, ok := carts[Position{row: row, col: col}]
			if ok {
				fmt.Print(string(translate_to_symbol[Direction{row: cart.dir_row, col: cart.dir_col}]))
			} else {
				fmt.Print(string(cell))
			}
		}
		fmt.Println()
	}
	fmt.Println("==============================================")
}

func parse(filename string) ([][]rune, map[Position]Cart) {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic("File not found")
	}

	grid := [][]rune{}
	carts := make(map[Position]Cart)
	cart_id := 0
	for row, line := range strings.Split(strings.Trim(string(data), "\n"), "\n") {
		new_row := []rune{}
		for col, char := range line {
			if is_cart(char) {
				dir, ok1 := translate_to_direction[char]
				if !ok1 {
					panic("wrong direction translation")
				}
				carts[Position{row: row, col: col}] = Cart{
					cart_id: cart_id,
					pos_row: row,
					pos_col: col,
					dir_row: dir.row,
					dir_col: dir.col,
					counter: 0,
				}
				track_char, ok2 := translate_to_grid[char]
				if !ok2 {
					panic("wrong track char")
				}
				new_row = append(new_row, track_char)
			} else {
				new_row = append(new_row, char)
			}
			cart_id += 1
		}
		grid = append(grid, new_row)
	}
	draw(grid, carts)
	return grid, carts
}

func solve(grid [][]rune, carts map[Position]Cart) (int, int) {
	fmt.Println("n carts", len(carts))
	for {
		// tick movements

		visited := make(map[int]bool)
		clear(visited)
		// fmt.Println("n carts", len(carts))
		// fmt.Println(carts)
		// fmt.Println("visited", visited)
		// fmt.Println("grid rows", len(grid), "grid cols", len(grid[0]))
		for row := 0; row < len(grid); row++ {
			for col := 0; col < len(grid[row]); col++ {
				// fmt.Print(row, col, " , ")
				position := Position{row: row, col: col}
				cart, theres_a_car_here := carts[position]
				if !theres_a_car_here {
					continue
				}
				// fmt.Println("found car at", row, col)
				_, is_visited := visited[cart.cart_id]
				if is_visited {
					continue
				}
				// fmt.Println("n carts", len(carts))
				visited[cart.cart_id] = true
				// fmt.Println("working at", row, col)
				next_position := cart.Move(grid)

				if next_position == position {
					panic("same position")
				}

				_, already_a_car_here := carts[next_position]
				if already_a_car_here {
					fmt.Println("crash at ", next_position)
					delete(carts, position)
					delete(carts, next_position)
				} else {
					delete(carts, position)
					carts[next_position] = cart
				}
				// fmt.Println("n carts", len(carts))
			}
		}
		// draw(grid, carts)

		if len(carts) == 1 {
			for position, _ := range carts {
				return position.col, position.row
			}
		}
	}
}

func solution(filename string) (int, int) {
	grid, carts := parse(filename)
	return solve(grid, carts)
}

func main() {
	fmt.Println(solution("./example.txt")) // 6,4
	fmt.Println(solution("./input.txt"))   // 29,74
}
