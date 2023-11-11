package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const INIT_CELL int = -1
const TIED int = -2
const VISITED int = -3

type manhattan struct {
	coord    int
	distance int
}

type bfs_node struct {
	row int
	col int
	// count int
}

func parse(filename string) [][]int {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic("File error")
	}
	content := string(data)

	coordinates := [][]int{}
	for _, line := range strings.Split(strings.Trim(content, "\n"), "\n") {
		fields := strings.Split(line, ",")
		col, _ := strconv.Atoi(strings.Trim(fields[0], " "))
		row, _ := strconv.Atoi(strings.Trim(fields[1], " "))
		coordinates = append(coordinates, []int{row, col})
		// fmt.Printf("x: %v, y: %v\n", x, y)
	}
	return coordinates
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func calculate_area(board [][]manhattan, row int, col int) int {
	coord_id := board[row][col].coord
	if coord_id == TIED {
		return 0
	}
	q := []bfs_node{bfs_node{row: row, col: col}}
	// queue = append(queue, bfs_node{row: row, col: col, count: 1})
	visited := make(map[bfs_node]bool)
	visited[bfs_node{row: row, col: col}] = true
	node := bfs_node{}
	// _ = node

	steps := [][]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}
	// _ = steps

	is_infinite := false
	area := 0
	// outer:
	for len(q) > 0 {
		node, q = q[len(q)-1], q[:len(q)-1]
		// fmt.Println(node.row, node.col)
		// area = max(area, node.count)
		board[node.row][node.col].coord = VISITED
		area += 1
		for _, step := range steps {
			new_row := node.row + step[0]
			new_col := node.col + step[1]

			if !(0 <= new_row && new_row < len(board) && 0 <= new_col && new_col < len(board[0])) {
				is_infinite = true
				continue
				// fmt.Print(" infinite ")
				// break outer
			}

			if board[new_row][new_col].coord == VISITED {
				continue
			}

			_, ok := visited[bfs_node{row: new_row, col: new_col}]
			if ok {
				continue
			}

			if board[new_row][new_col].coord != coord_id {
				continue
			}

			visited[bfs_node{row: new_row, col: new_col}] = true
			q = append(q, bfs_node{row: new_row, col: new_col})
		}

	}
	if is_infinite {
		return 0
	}
	return area
}

func solve(coords [][]int, safe_limit int) int {
	max_row := 0
	max_col := 0
	for _, coord := range coords {
		// fmt.Println(coord[0], coord[1])
		max_row = max(max_col, coord[0])
		max_col = max(max_row, coord[1])
	}
	max_row += 1
	max_col += 1

	board := [][]manhattan{}
	for row := 0; row < max_row; row++ {
		new_row := []manhattan{}
		for col := 0; col < max_col; col++ {
			new_row = append(new_row, manhattan{coord: INIT_CELL, distance: 0})
		}
		board = append(board, new_row)
	}

	for _, coord := range coords {
		coord_row := coord[0]
		coord_col := coord[1]

		for row := 0; row < max_row; row++ {
			for col := 0; col < max_col; col++ {
				man_distance := abs(coord_row-row) + abs(coord_col-col)
				board[row][col].distance += man_distance
			}
		}
	}

	// for row := 0; row < max_row; row++ {
	// 	for col := 0; col < max_col; col++ {
	// 		if board[row][col].coord == TIED {
	// 			fmt.Print(".")
	// 		} else {
	// 			fmt.Printf("%v", board[row][col].coord)
	// 		}
	// 	}
	// 	fmt.Println()
	// }

	safe_region := 0
	for row := 0; row < max_row; row++ {
		for col := 0; col < max_col; col++ {
			if board[row][col].distance < safe_limit {
				safe_region += 1
			}
		}
	}

	return safe_region
}

func solution(filename string, safe_limit int) int {
	coordinates := parse(filename)
	return solve(coordinates, safe_limit)
}

func main() {
	fmt.Println(solution("example.txt", 32))  // 16
	fmt.Println(solution("input.txt", 10000)) // 42998
}
