package main

import (
	"fmt"
	"math"
)

const GRID_SIZE = 300

func power_at(x int, y int, serial int) int {
	rack_id := x + 10
	// fmt.Println("rack_id:", rack_id)
	power_level := (rack_id * y)
	// fmt.Println("power starts:", power_level)
	power_level += serial
	// fmt.Println("adding serial:", power_level)
	power_level *= rack_id
	// fmt.Println("mult rack id:", power_level)
	hundreds := (power_level / 100) % 10
	// fmt.Println("hundreds:", hundreds)
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
	// fmt.Println(power_at(3, 5, 8))      // 4
	// fmt.Println(power_at(122, 79, 57))  // -5
	// fmt.Println(power_at(217, 196, 39)) // 0
	// fmt.Println(power_at(101, 153, 71)) // 4
	x, y := solution(18)
	fmt.Printf("%d,%d\n", x, y)

	x, y = solution(42)
	fmt.Printf("%d,%d\n", x, y)

	x, y = solution(5235)
	fmt.Printf("%d,%d\n", x, y)
}
