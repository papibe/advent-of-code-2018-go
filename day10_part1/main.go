package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func parse(filename string) [][][]int {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic("File not found")
	}
	re_main := regexp.MustCompile(`position=<(.*)> velocity=<(.*)>`)
	re_points := regexp.MustCompile(`([- ]*)(\d+), ([- ]*)(\d+)`)
	content := string(data)
	_ = re_main
	_ = re_points

	for _, line := range strings.Split(strings.Trim(content, "\n"), "\n") {
		// fmt.Println(line)
		main_matches := re_main.FindStringSubmatch(line)
		// fmt.Println(main_matches[1])
		point_matches := re_points.FindStringSubmatch(main_matches[1])
		// speed_matches := re_points.FindStringSubmatch(main_matches[2])
		// fmt.Println(point_matches[1], "<>", point_matches[2])
		pointx, _ := strconv.Atoi(point_matches[2])
		if point_matches[1] == "-" {
			pointx *= -1
		}
		pointy, _ := strconv.Atoi(point_matches[4])
		if point_matches[3] == "-" {
			pointy *= -1
		}
		fmt.Println(line, "\t", pointx, pointy)

		// fmt.Println(speed_matches[1], "<>", speed_matches[2])
	}

	return [][][]int{}
}

func solution(filename string) int {
	points := parse(filename)
	fmt.Println(points)
	return 0
}

func main() {
	fmt.Println(solution("./example.txt"))
}
