package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func solution(filename string) int {
	data, err := os.ReadFile(filename)

	if err != nil {
		panic("file error")
	}

	content := string(data)

	current_frequency := 0
	frequencies := map[int]bool{0: true}
	for {
		for _, frequency := range strings.Split(strings.Trim(content, "\n"), "\n") {
			frequency_change, _ := strconv.Atoi(frequency)
			current_frequency += frequency_change
			_, seen_before := frequencies[current_frequency]
			if seen_before {
				return current_frequency
			}
			frequencies[current_frequency] = true
		}
	}

	return 0
}

func main() {
	fmt.Println(solution("./input.txt"))
}
