package main

import (
	"container/heap"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
)

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

func parse(filename string) [][]int {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic("File not found")
	}
	content := string(data)
	re := regexp.MustCompile(`Step (.) must be finished before step (.) can begin.`)

	output := [][]int{}

	for _, line := range strings.Split(strings.Trim(content, "\n"), "\n") {
		matches := re.FindStringSubmatch(line)

		dependency := matches[1]
		node := matches[2]

		output = append(output, []int{int(dependency[0]), int(node[0])})
	}
	return output
}

func SliceToString(a []int) string {
	if len(a) == 0 {
		return ""
	}

	b := make([]byte, len(a))
	for i, v := range a {
		b[i] = byte(v)
	}
	return string(b)
}

func create_workers_queues(n int) [][]int {
	a := make([][]int, n)
	for i := 0; i < n; i++ {
		a[i] = []int{}
	}
	return a
}

func is_available_worker(wq [][]int) int {
	for index, worker := range wq {
		if len(worker) == 0 {
			return index
		}
	}
	return -1
}

func workers_have_work(wq [][]int) bool {
	for _, worker := range wq {
		if len(worker) != 0 {
			return true
		}
	}
	return false
}

func solution(filename string, number_of_workers int, base_time int) int {
	data := parse(filename)
	unique_nodes := make(map[int]bool)
	for _, rule := range data {
		unique_nodes[rule[0]] = true
		unique_nodes[rule[1]] = false
	}

	nodes := []int{}
	for node, _ := range unique_nodes {
		nodes = append(nodes, node)
	}
	sort.Ints(nodes)
	number_of_nodes := len(nodes)

	nodes_map := make(map[int]int)
	for index, node := range nodes {
		nodes_map[node] = index
	}

	adjacency_matrix := make([][]int, len(nodes))
	for i := range adjacency_matrix {
		adjacency_matrix[i] = make([]int, len(nodes))
	}

	for _, rule := range data {
		dependency := rule[0]
		node := rule[1]
		dep_index := nodes_map[dependency]
		node_index := nodes_map[node]
		adjacency_matrix[dep_index][node_index] = 1
	}

	incomming_degree := make([]int, len(nodes))
	for j := 0; j < number_of_nodes; j++ {
		for i := 0; i < number_of_nodes; i++ {
			incomming_degree[j] += adjacency_matrix[i][j]
		}
	}

	hqueue := &IntHeap{}
	heap.Init(hqueue)

	for index, degree := range incomming_degree {
		if degree == 0 {
			heap.Push(hqueue, index)
		}
	}

	// workers_queue := []int{}
	output := []int{}
	_ = output
	workers := create_workers_queues(number_of_workers)
	seconds := 0

	// BFS with a heap instead of a queue
	for len(*hqueue) > 0 || workers_have_work(workers) {
		available_worker := is_available_worker(workers)
		if len(*hqueue) > 0 && available_worker >= 0 {
			pop_node := heap.Pop(hqueue)
			node, ok := pop_node.(int)
			if !ok {
				panic("blha")
			}

			// route work to workers
			for i := 0; i < base_time+node+1; i++ {
				workers[available_worker] = append(workers[available_worker], node)
			}
			continue
		}
		// process workers' work
		for i := 0; i < number_of_workers; i++ {
			wlen := len(workers[i])
			if wlen > 0 {
				// pop from worker queue
				var node int
				node, workers[i] = workers[i][wlen-1], workers[i][:wlen-1]
				// fmt.Println("worker->", worker)
				if len(workers[i]) == 0 {
					output = append(output, nodes[node])

					// update incomming degree
					for j := 0; j < len(adjacency_matrix); j++ {
						if adjacency_matrix[node][j] == 1 {
							incomming_degree[j] -= 1
							// push to main queue if dependencies met
							if incomming_degree[j] == 0 {
								heap.Push(hqueue, j)
							}
						}
					}

				}
			}
		}

		seconds += 1
		// fmt.Println(hqueue)
		// for _, worker := range workers {
		// 	fmt.Println(worker)
		// }
		// fmt.Println("===================================")
	}
	// fmt.Println(output)
	// fmt.Println(SliceToString(output))
	return seconds
}

func main() {
	// fmt.Println(solution("./example.txt", 2, 0))
	fmt.Println(solution("./input.txt", 5, 60))
}
