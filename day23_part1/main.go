package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Nanobot struct {
	x      int
	y      int
	z      int
	radius int
}

func parse(filename string) []Nanobot {
	raw_data, err := os.ReadFile(filename)
	if err != nil {
		panic("File not found")
	}
	data := strings.Trim(string(raw_data), "\n")

	re := regexp.MustCompile(`pos=<(-*\d+),(-*\d+),(-*\d+)>, r=(\d+)`)

	nanobots := []Nanobot{}
	for _, line := range strings.Split(data, "\n") {
		matches := re.FindStringSubmatch(line)
		x, _ := strconv.Atoi(matches[1])
		y, _ := strconv.Atoi(matches[2])
		z, _ := strconv.Atoi(matches[3])
		radius, _ := strconv.Atoi(matches[4])

		nanobots = append(nanobots, Nanobot{x, y, z, radius})
	}
	return nanobots
}

func solve(nanobots []Nanobot) int {
	var strongest_nanobot Nanobot
	max_radius := math.MinInt
	for _, nanobot := range nanobots {
		if nanobot.radius > max_radius {
			max_radius = nanobot.radius
			strongest_nanobot = nanobot
		}
	}
	in_range := 0
	for _, nanobot := range nanobots {
		distance := math.Abs(float64(nanobot.x)-float64(strongest_nanobot.x)) +
			math.Abs(float64(nanobot.y)-float64(strongest_nanobot.y)) +
			math.Abs(float64(nanobot.z)-float64(strongest_nanobot.z))

		if distance <= float64(max_radius) {
			in_range += 1
		}
	}
	return in_range
}

func solution(filename string) int {
	nanobots := parse(filename)
	return solve(nanobots)
}

func main() {
	fmt.Println(solution("./example.txt")) // 7
	fmt.Println(solution("./input.txt"))   // 305
}
