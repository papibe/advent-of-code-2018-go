package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Coord struct {
	w int
	x int
	y int
	z int
}

func parse(filename string) []Coord {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic("File error")
	}
	content := string(data)
	raw_lines := strings.Trim(content, "\n")
	lines := strings.Split(raw_lines, "\n")

	re := regexp.MustCompile(`(-*\d+),(-*\d+),(-*\d+),(-*\d+)`)

	points := []Coord{}
	for _, line := range lines {
		matches := re.FindStringSubmatch(line)
		w, _ := strconv.Atoi(matches[1])
		x, _ := strconv.Atoi(matches[2])
		y, _ := strconv.Atoi(matches[3])
		z, _ := strconv.Atoi(matches[4])
		points = append(points, Coord{w, x, y, z})
	}
	return points
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func manhattan(p1, p2 Coord) int {
	return abs(p1.w-p2.w) + abs(p1.x-p2.x) + abs(p1.y-p2.y) + abs(p1.z-p2.z)
}

func find(p int, parents []int) int {
	root := p
	for root != parents[root] {
		root = parents[root]
	}
	return root
}

func union(p, q int, parents []int) {
	root1 := find(p, parents)
	root2 := find(q, parents)
	parents[root1] = root2
}

func solve(points []Coord) int {
	n := len(points)
	parents := make([]int, n)
	for i := 0; i < n; i++ {
		parents[i] = i
	}
	for i, coord1 := range points {
		for j := i + 1; j < n; j++ {
			coord2 := points[j]
			if manhattan(coord1, coord2) <= 3 {
				union(i, j, parents)
			}
		}
	}
	constellation := make(map[int]bool)
	for i := 0; i < n; i++ {
		key := find(i, parents)
		constellation[key] = true
	}
	return len(constellation)
}

func solution(filename string) int {
	coordinates := parse(filename)
	return solve(coordinates)
}

func main() {
	fmt.Println(solution("input.txt")) // 420
}
