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
	for _, frequency := range strings.Split(strings.Trim(content, "\n"), "\n") {
		frequency_change, _ := strconv.Atoi(frequency)
		current_frequency += frequency_change
	}

	return current_frequency
}

func main() {
	fmt.Println(solution("./input.txt"))
}
