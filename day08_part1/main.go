package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parse(filename string) []int {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic("File not found")
	}
	content := string(data)

	numbers := []int{}
	for _, str_number := range strings.Split(strings.Trim(content, "\n"), " ") {
		number, _ := strconv.Atoi(str_number)
		numbers = append(numbers, number)
	}
	return numbers
}

func solve(lf []int, start int, nodes int) (int, int) {
	internal_nodes := lf[start]
	meta_size := lf[start+1]
	// fmt.Println("index: {}, ", start, internal_nodes, meta_size)
	fmt.Println("index:", start, "nodes:", nodes)

	if internal_nodes == 0 {
		meta_sum := 0
		for i := 0; i < meta_size; i++ {
			meta_sum += lf[start+2+i]
		}
		return meta_sum, start + 1 + meta_size
	}

	current_sum := 0
	new_start := start + 2
	for i := 0; i < nodes; i++ {
		fmt.Println("call")
		sum, end := solve(lf, new_start, internal_nodes)
		current_sum += sum
		new_start = end + 1
	}
	meta_sum := 0
	for i := 0; i < meta_size; i++ {
		meta_sum += lf[new_start+i]
	}

	return current_sum + meta_sum, new_start + meta_size
}

func solution(filename string) int {
	licence := parse(filename)
	sum, _ := solve(licence, 0, 2)
	return sum
}

func main() {
	fmt.Println(solution("./example.txt"))
}
