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

func solution(filename string) string {
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

	output := []int{}

	// BFS with a heap instead of a queue
	for len(*hqueue) > 0 {
		pop_node := heap.Pop(hqueue)
		node, ok := pop_node.(int)
		if ok {
			output = append(output, nodes[node])
		}
		for j := 0; j < len(adjacency_matrix); j++ {
			if adjacency_matrix[node][j] == 1 {
				incomming_degree[j] -= 1
				if incomming_degree[j] == 0 {

					heap.Push(hqueue, j)
				}
			}
		}
	}
	return SliceToString(output)
}

func main() {
	fmt.Println(solution("./input.txt"))
}
