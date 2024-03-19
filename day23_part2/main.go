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

type Coord struct {
	x int
	y int
	z int
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

func get_points_inside(nanobot Nanobot) map[Coord]bool {
	x0 := nanobot.x
	y0 := nanobot.y
	z0 := nanobot.z
	r := nanobot.radius

	points := make(map[Coord]bool)
	for x := x0; x <= x0+r; x++ {
		max_y := y0 + r + x0 - x

		for y := y0; y <= max_y; y++ {
			d := int(float64(r) - math.Abs(float64(x0)-float64(x)) - math.Abs(float64(y0)-float64(y)))

			for z := z0 - d; z <= z0+d; z++ {
				points[Coord{x, y, z}] = true
			}

			d = int(float64(r) - math.Abs(float64(x0)-float64(x)) - math.Abs(float64(y0)-float64(y0-(y-y0))))
			for z := z0 - d; z <= z0+d; z++ {
				points[Coord{x, y0 - (y - y0), z}] = true
			}

			d = int(float64(r) - math.Abs(float64(x0)-float64(x0-(x-x0))) - math.Abs(float64(y0)-float64(y)))
			for z := z0 - d; z <= z0+d; z++ {
				points[Coord{x0 - (x - x0), y, z}] = true
			}

			d = int(float64(r) - math.Abs(float64(x0)-float64(x0-(x-x0))) - math.Abs(float64(y0)-float64(y0-(y-y0))))
			for z := z0 - d; z <= z0+d; z++ {
				points[Coord{x0 - (x - x0), y0 - (y - y0), z}] = true
			}

		}
	}
	return points
}

func distance(point Coord, nanobot Nanobot) int {
	d := math.Abs(float64(nanobot.x)-float64(point.x)) + math.Abs(float64(nanobot.y)-float64(point.y)) + math.Abs(float64(nanobot.z)-float64(point.z))

	return int(d)
}

func ndistance(nb1, nb2 Nanobot) int {
	d := math.Abs(float64(nb1.x)-float64(nb2.x)) + math.Abs(float64(nb1.y)-float64(nb2.y)) + math.Abs(float64(nb1.z)-float64(nb2.z))

	return int(d)
}

func solve(nanobots []Nanobot) int {
	var best_coord = Coord{0, 0, 0}
	// max_num_nanobots := math.MinInt

	counter := 0
	for i := 0; i < len(nanobots); i++ {
		for j := i + 1; j < len(nanobots); j++ {
			if ndistance(nanobots[i], nanobots[j]) < nanobots[i].radius+nanobots[j].radius {
				counter += 1
			}
		}
	}
	fmt.Println("counter", counter)

	d := math.Abs(float64(0)-float64(best_coord.x)) + math.Abs(float64(0)-float64(best_coord.y)) + math.Abs(float64(0)-float64(best_coord.z))
	return int(d)
}

func solution(filename string) int {
	nanobots := parse(filename)
	fmt.Println("num nano", len(nanobots))
	return solve(nanobots)
}

func main() {
	fmt.Println(solution("./example.txt")) //
	fmt.Println(solution("./input.txt"))   //
}
