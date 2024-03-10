package main

import (
	"fmt"
)

type Coord struct {
	x int
	y int
}

var memo = make(map[Coord]int)

func geologic_index(x, y, depth, target_x, target_y int) int {
	if x == target_x && y == target_y {
		return 0
	}
	if x == 0 && y == 0 {
		return 0
	}
	if y == 0 {
		return x * 16807
	}
	if x == 0 {
		return y * 48271
	}
	key := Coord{x, y}
	result, seen_before := memo[key]
	if seen_before {
		return result
	}
	left_gindex := geologic_index(x-1, y, depth, target_x, target_y)
	upper_gindex := geologic_index(x, y-1, depth, target_x, target_y)
	memo[key] = erosion_level(left_gindex, depth) * erosion_level(upper_gindex, depth)

	return memo[key]
}

func erosion_level(gindex, depth int) int {
	return (gindex + depth) % 20183
}

func solution(depth, target_x, target_y int) int {
	region := []int{0, 0, 0}
	for x := 0; x <= target_x; x++ {
		for y := 0; y <= target_y; y++ {
			gindex := geologic_index(x, y, depth, target_x, target_y)
			elevel := erosion_level(gindex, depth)
			region_type := elevel % 3
			region[region_type] += 1
		}
	}
	return region[1] + region[2]*2
}

func main() {
	fmt.Println(solution(510, 10, 10))   // 114
	fmt.Println(solution(4080, 14, 785)) // 11843
}
