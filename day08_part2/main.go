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

func solve(lf []int, start int) (int, int) {
	internal_nodes := lf[start]
	meta_size := lf[start+1]

	if internal_nodes == 0 {
		meta_sum := 0
		for i := 0; i < meta_size; i++ {
			meta_sum += lf[start+2+i]
		}
		return meta_sum, start + 1 + meta_size
	}

	children_value := []int{}
	new_start := start + 2
	for i := 0; i < internal_nodes; i++ {
		sum, end := solve(lf, new_start)
		children_value = append(children_value, sum)
		new_start = end + 1
	}
	meta_sum := 0
	for i := 0; i < meta_size; i++ {
		child := lf[new_start+i] - 1
		if 0 <= child && child < len(children_value) {
			meta_sum += children_value[child]
		}
	}
	return meta_sum, new_start + meta_size - 1
}

func solution(filename string) int {
	license := parse(filename)
	sum, _ := solve(license, 0)
	return sum
}

func main() {
	fmt.Println(solution("./input.txt")) // 30063
}
