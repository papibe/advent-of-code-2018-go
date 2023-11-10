package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func parse(filename string) [][]int {
	data, err := os.ReadFile(filename)

	if err != nil {
		panic("file error")
	}
	content := string(data)

	re := regexp.MustCompile(`#\d+ @ (\d+),(\d+): (\d+)x(\d+)`)
	claims := [][]int{}
	for _, line := range strings.Split(strings.Trim(content, "\n"), "\n") {
		matches := re.FindStringSubmatch(line)

		col, _ := strconv.Atoi(matches[1])
		row, _ := strconv.Atoi(matches[2])
		width, _ := strconv.Atoi(matches[3])
		height, _ := strconv.Atoi(matches[4])

		claims = append(claims, []int{col, row, width, height})
	}

	return claims
}

func solution(filename string) int {
	claims := parse(filename)

	type point struct {
		col int
		row int
	}

	fabric_cuts := make(map[point]int)

	for _, claim := range claims {
		col := claim[0]
		row := claim[1]
		width := claim[2]
		height := claim[3]

		for i := col; i < col+width; i++ {
			for j := row; j < row+height; j++ {
				key := point{col: i, row: j}
				value, ok := fabric_cuts[key]
				if ok {
					fabric_cuts[key] = value + 1
				} else {
					fabric_cuts[key] = 1
				}
			}
		}
	}

	for claim_number, claim := range claims {
		col := claim[0]
		row := claim[1]
		width := claim[2]
		height := claim[3]

		intersection := false
	main_loop:
		for i := col; i < col+width; i++ {
			for j := row; j < row+height; j++ {
				key := point{col: i, row: j}
				value, ok := fabric_cuts[key]
				if ok && value != 1 {
					intersection = true
					break main_loop
				}
				// value, ok := fabric_cuts[key]
				// if ok {
				// 	fabric_cuts[key] = value + 1
				// } else {
				// 	fabric_cuts[key] = 1
				// }
			}
		}
		if !intersection {
			return claim_number + 1
		}
	}
	return 0
}

func main() {
	fmt.Println(solution("./example.txt")) // 3
	fmt.Println(solution("./input.txt"))   // 445
}
