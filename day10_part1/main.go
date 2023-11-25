package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Tuple struct {
	x int
	y int
}

func get_points(matches []string) (int, int) {
	point_x, _ := strconv.Atoi(matches[2])
	if matches[1] == "-" {
		point_x *= -1
	}
	point_y, _ := strconv.Atoi(matches[4])
	if matches[3] == "-" {
		point_y *= -1
	}
	return point_x, point_y
}

func parse(filename string) ([]Tuple, []Tuple) {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic("File not found")
	}

	re_main := regexp.MustCompile(`position=<(.*)> velocity=<(.*)>`)
	re_points := regexp.MustCompile(`([- ]*)(\d+), ([- ]*)(\d+)`)
	content := string(data)

	points := []Tuple{}
	speeds := []Tuple{}

	for _, line := range strings.Split(strings.Trim(content, "\n"), "\n") {
		main_matches := re_main.FindStringSubmatch(line)
		point_matches := re_points.FindStringSubmatch(main_matches[1])
		speed_matches := re_points.FindStringSubmatch(main_matches[2])

		point_x, point_y := get_points(point_matches)
		speed_x, speed_y := get_points(speed_matches)
		// fmt.Println(line, "\t", point_x, point_y, "\t", speed_x, speed_y)

		points = append(points, Tuple{x: point_x, y: point_y})
		speeds = append(speeds, Tuple{x: speed_x, y: speed_y})

	}
	return points, speeds
}

func print_points(current_points []Tuple) bool {
	points := make(map[Tuple]bool)
	for _, point := range current_points {
		points[Tuple{x: point.x, y: point.y}] = true
	}
	max_x := math.MinInt
	min_x := math.MaxInt
	max_y := math.MinInt
	min_y := math.MaxInt
	for point, _ := range points {
		if point.x > max_x {
			max_x = point.x
		}
		if point.x < min_x {
			min_x = point.x
		}
		if point.y > max_y {
			max_y = point.y
		}
		if point.y < min_y {
			min_y = point.y
		}
	}
	// fmt.Println(min_x, max_x, min_y, max_y, "\t", max_x-min_x, max_y-min_y)

	if max_y-min_y > 10 {
		return false
	}
	for y := min_y; y <= max_y; y++ {
		for x := min_x; x <= max_x; x++ {
			_, ok := points[Tuple{x: x, y: y}]
			if ok {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}

	return true
}

func solve(current_points []Tuple, speeds []Tuple) int {
	counter := 0
	for {
		if print_points(current_points) {
			break
		}
		next_points := []Tuple{}
		for index := 0; index < len(current_points); index++ {
			next_point := Tuple{
				x: current_points[index].x + speeds[index].x,
				y: current_points[index].y + speeds[index].y,
			}
			next_points = append(next_points, next_point)
		}

		counter += 1
		current_points = next_points
	}
	return counter
}

func solution(filename string) int {
	points, speeds := parse(filename)
	return solve(points, speeds)
}

func main() {
	// fmt.Println(solution("./example.txt")) // HI
	fmt.Println(solution("./input.txt")) // GPJLLLLH
}
