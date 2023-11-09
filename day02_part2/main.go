package main

import (
	"fmt"
	"os"
	"strings"
)

func solution(filename string) string {
	data, err := os.ReadFile(filename)

	if err != nil {
		panic("file error")
	}
	content := string(data)

	box_ids := strings.Split(strings.Trim(content, "\n"), "\n")

	for i := 0; i < len(box_ids); i++ {
		for j := i + 1; j < len(box_ids); j++ {

			diffs := 0
			same_chars := []byte{}

			for index := 0; index < len(box_ids[i]) && diffs <= 1; index++ {
				if box_ids[i][index] != box_ids[j][index] {
					diffs += 1
				} else {
					same_chars = append(same_chars, box_ids[j][index])
				}
			}
			if diffs == 1 {
				return string(same_chars[:])
			}
		}
	}

	return ""
}

func main() {
	fmt.Println(solution("./input.txt")) // "ymdrchgpvwfloluktajxijsqb"
}
