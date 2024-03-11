package main

import (
	"container/heap"
	"fmt"
	"math"
)

type Coord struct {
	x int
	y int
}

const ROCKY = 0
const WET = 1
const NARROW = 2

const TORCH = 0
const CLIMBING_GEAR = 1
const NEITHER = 2

var right_tool = map[int]map[int]bool{
	ROCKY: {
		TORCH:         true,
		CLIMBING_GEAR: true,
		NEITHER:       false,
	},
	WET: {
		TORCH:         false,
		CLIMBING_GEAR: true,
		NEITHER:       true,
	},
	NARROW: {
		TORCH:         true,
		CLIMBING_GEAR: false,
		NEITHER:       true,
	},
}

type State struct {
	x           int
	y           int
	tool        int
	region_type int
	switch_gear bool
}

// An Item is something we manage in a priority queue.
type Item struct {
	value    State // The value of the item; arbitrary.
	priority int   // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) update(item *Item, value State, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
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

func get_cave(depth, x, y, target_x, target_y int) int {
	if x == target_x && y == target_y {
		return ROCKY
	}
	gindex := geologic_index(x, y, depth, target_x, target_y)
	elevel := erosion_level(gindex, depth)
	region_type := elevel % 3
	return region_type
}

// func solve_sample(cave map[Coord]int) int {
// 	s1 := State{0, 1, 0}
// 	s2 := State{0, 2, 0}
// 	s3 := State{0, 3, 0}

// 	// Some items and their priorities.
// 	items := map[State]int{
// 		s1: 3, s2: 2, s3: 4,
// 	}

// 	// Create a priority queue, put the items in it, and
// 	// establish the priority queue (heap) invariants.
// 	pq := PriorityQueue{}
// 	heap.Init(&pq)

// 	i := 0
// 	for value, priority := range items {
// 		heap.Push(&pq, &Item{value: value, priority: priority, index: i})
// 		i++
// 	}

// 	// Insert a new item and then modify its priority.
// 	item := &Item{
// 		value:    State{0, 4, 0},
// 		priority: 1,
// 	}
// 	heap.Push(&pq, item)
// 	// pq.update(item, item.value, 5)

// 	// Take the items out; they arrive in decreasing priority order.
// 	for pq.Len() > 0 {
// 		item := heap.Pop(&pq).(*Item)
// 		fmt.Printf("%.2d:%v\n", item.priority, item.value)
// 	}

// 	return 0
// }

func solve(depth, target_x, target_y int) int {
	// Dijstra init
	pq := PriorityQueue{}
	heap.Init(&pq)
	heap.Push(&pq, &Item{value: State{0, 0, TORCH, ROCKY, false}, priority: 0})
	visited := make(map[Coord]bool)
	visited[Coord{0, 0}] = true

	distances := make(map[Coord]int)
	distances[Coord{0, 0}] = 0

	// BFS Dijstra
	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)
		state := item.value
		x, y, tool, current_region_type := state.x, state.y, state.tool, state.region_type
		switch_gear := state.switch_gear
		minutes := item.priority

		fmt.Println(x, y, minutes)

		if x == target_x && y == target_y && tool == TORCH {
			fmt.Println("arriving. tool", tool)
			return minutes
		}

		steps := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
		for _, step := range steps {
			// only >= 0 coordinates
			new_x := x + step[0]
			new_y := y + step[1]
			if new_x < 0 || new_y < 0 {
				continue
			}
			_, is_visited := visited[Coord{new_x, new_y}]
			if is_visited {
				continue
			}

			next_region_type := get_cave(depth, new_x, new_y, target_x, target_y)

			if right_tool[next_region_type][tool] {
				// keep same tool
				new_minutes := minutes + 1
				current_distance, is_seen := distances[Coord{new_x, new_y}]
				if !is_seen {
					current_distance = math.MaxInt
				}
				if new_minutes < current_distance {
					item := &Item{
						value:    State{new_x, new_y, tool, next_region_type, false},
						priority: new_minutes,
					}
					heap.Push(&pq, item)
					distances[Coord{new_x, new_y}] = new_minutes
					visited[Coord{new_x, new_y}] = true
				}
				// continue
			}
		}

		if switch_gear {
			continue
		}

		// change tool
		for new_tool, is_valid_tool := range right_tool[current_region_type] {
			// fmt.Println(new_tool, is_valid_tool)
			if is_valid_tool && new_tool != tool {
				new_minutes := minutes + 7
				item := &Item{
					value:    State{x, y, new_tool, current_region_type, true},
					priority: new_minutes,
				}
				heap.Push(&pq, item)
				// distances[Coord{x, y}] = new_minutes
				// visited[Coord{x, y}] = true
			}
		}
		// break
	}
	return distances[Coord{target_x, target_y}]
}

func solution(depth, target_x, target_y int) int {
	// cave := create_cave_map(depth, target_x, target_y)
	return solve(depth, target_x, target_y)
}

func main() {
	fmt.Println(solution(510, 10, 10)) // 45
	// fmt.Println(solution(4080, 14, 785)) // 1098?
}
