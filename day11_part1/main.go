package main

import (
	"fmt"
	"math"
)

const GRID_SIZE = 300

func power_at(x int, y int, serial int) int {
	rack_id := x + 10
	power_level := (rack_id * y)
	power_level += serial
	power_level *= rack_id
	hundreds := (power_level / 100) % 10
	hundreds -= 5
	return hundreds
}

func solution(serial int) (int, int) {
	size := GRID_SIZE + 1
	power := make([][]int, size)
	for i := 0; i < size; i++ {
		row := make([]int, size)
		power[i] = row
	}
	for x := 1; x <= GRID_SIZE; x++ {
		for y := 1; y <= GRID_SIZE; y++ {
			power[x][y] = power_at(x, y, serial)
		}
	}
	max_x := 0
	max_y := 0
	max_power := math.MinInt
	for x := 1; x <= GRID_SIZE-3; x++ {
		for y := 1; y <= GRID_SIZE-3; y++ {
			power_sum := 0
			for row := 0; row < 3; row++ {
				for col := 0; col < 3; col++ {
					power_sum += power[x+row][y+col]
				}
			}
			if power_sum > max_power {
				max_power = power_sum
				max_x = x
				max_y = y
			}
		}
	}

	return max_x, max_y
}

func main() {
	x, y := solution(5235)
	fmt.Printf("%d,%d\n", x, y) // 33,54
}
