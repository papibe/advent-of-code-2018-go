package main

import (
	"container/heap"
	"fmt"
)

type Coord struct {
	x int
	y int
}

type Seen struct {
	x    int
	y    int
	tool int
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

func solve(depth, target_x, target_y int) int {
	// Dijstra init
	pq := PriorityQueue{}
	heap.Init(&pq)
	heap.Push(&pq, &Item{value: State{0, 0, TORCH, ROCKY}, priority: 0})
	visited := make(map[Coord]bool)
	visited[Coord{0, 0}] = true

	distances := make(map[Coord]int)
	distances[Coord{0, 0}] = 0

	seen := make(map[Seen]bool)

	// BFS Dijstra
	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)
		state := item.value
		x, y, tool, current_region_type := state.x, state.y, state.tool, state.region_type
		minutes := item.priority

		if x == target_x && y == target_y && tool == TORCH {
			return minutes
		}

		_, is_seen := seen[Seen{x, y, tool}]
		if is_seen {
			continue
		}
		seen[Seen{x, y, tool}] = true

		steps := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
		for _, step := range steps {
			// only >= 0 coordinates
			new_x := x + step[0]
			new_y := y + step[1]
			if new_x < 0 || new_y < 0 {
				continue
			}

			next_region_type := get_cave(depth, new_x, new_y, target_x, target_y)

			if right_tool[next_region_type][tool] {
				// keep same tool
				new_minutes := minutes + 1
				item := &Item{
					value:    State{new_x, new_y, tool, next_region_type},
					priority: new_minutes,
				}
				heap.Push(&pq, item)
			}
		}
		// change tool
		for new_tool, is_valid_tool := range right_tool[current_region_type] {
			if is_valid_tool && new_tool != tool {
				new_minutes := minutes + 7
				item := &Item{
					value:    State{x, y, new_tool, current_region_type},
					priority: new_minutes,
				}
				heap.Push(&pq, item)
			}
		}
	}
	return -1
}

func solution(depth, target_x, target_y int) int {
	clear(memo)
	return solve(depth, target_x, target_y)
}

func main() {
	fmt.Println(solution(510, 10, 10))   // 45
	fmt.Println(solution(4080, 14, 785)) // 1078
}
