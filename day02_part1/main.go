package main

import (
	"fmt"
	"os"
	"strings"
)

func solution(filename string) int {
	data, err := os.ReadFile(filename)

	if err != nil {
		panic("file error")
	}
	content := string(data)

	twice_reps := 0
	three_reps := 0
	for _, box_id := range strings.Split(strings.Trim(content, "\n"), "\n") {
		frequency := make(map[uint8]int)
		for i := 0; i < len(box_id); i++ {
			frequency[box_id[i]] += 1
		}
		for _, value := range frequency {
			if value == 2 {
				twice_reps += 1
				break
			}
		}
		for _, value := range frequency {
			if value == 3 {
				three_reps += 1
				break
			}
		}
	}

	return twice_reps * three_reps
}

func main() {
	fmt.Println(solution("./example.txt")) // 12
	fmt.Println(solution("./input.txt"))   // 5000
}
