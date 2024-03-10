package main

import (
	"container/heap"
	"fmt"
)

type Coord struct {
	x int
	y int
}

const ROCKY = 0
const WET = 1
const NARROW = 2

type State struct {
	x       int
	y       int
	tool    int
	minutes int
}

// An IntHeap is a min-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
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

func create_cave_map(depth, target_x, target_y int) map[Coord]int {
	cave := make(map[Coord]int)
	for x := 0; x <= target_x; x++ {
		for y := 0; y <= target_y; y++ {
			gindex := geologic_index(x, y, depth, target_x, target_y)
			elevel := erosion_level(gindex, depth)
			region_type := elevel % 3
			cave[Coord{x, y}] = region_type
		}
	}
	return cave

}

func solve(cave map[Coord]int) int {
	s1 := State{0, 0, 0, 5}
	s2 := State{0, 0, 0, 1}
	s3 := State{0, 0, 0, 3}

	h := &IntHeap{s1, s2, s3}
	heap.Init(h)
	heap.Push(h, State{0, 0, 0, 5})
	heap.Push(h, State{0, 0, 0, 1})
	heap.Push(h, State{0, 0, 0, 3})

	return 0
}

func solution(depth, target_x, target_y int) int {
	cave := create_cave_map(depth, target_x, target_y)
	return solve(cave)
}

func main() {
	fmt.Println(solution(510, 10, 10)) // 114
	// fmt.Println(solution(4080, 14, 785)) // 11843
}
