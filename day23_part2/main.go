package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"sort"
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

func is_inside(nb Nanobot, x, y, z int) bool {
	d := int(math.Abs(float64(nb.x)-float64(x)) + math.Abs(float64(nb.y)-float64(y)) + math.Abs(float64(nb.z)-float64(z)))
	return d <= nb.radius
}

func sort_points_by_intersections(points []Coord, intersections map[Coord]int) {
	sort.Slice(points, func(i, j int) bool {
		return intersections[points[i]] > intersections[points[j]]
	})
}

func get_intersections(point Coord, nanobots []Nanobot) int {
	counter := 0
	for j := 0; j < len(nanobots); j++ {
		if is_inside(nanobots[j], point.x, point.y, point.z) {
			counter += 1
		}
	}
	return counter
}

func man_to_origin(next_point Coord) int {
	return int(math.Abs(float64(next_point.x)) + math.Abs(float64(next_point.y)) + math.Abs(float64(next_point.z)))
}

func bfs(
	point Coord,
	visited *map[Coord]bool,
	current_intersections, max_man_distance, max_bfs_rec int,
	nanobots []Nanobot,
) (int, int, Coord) {
	queue := []Coord{point}
	(*visited)[point] = true
	local_min_man := math.MaxInt
	min_point := point

	for len(queue) > 0 && max_bfs_rec > 0 {
		p := queue[0]
		x, y, z := p.x, p.y, p.z
		queue = queue[1:]
		next_points := []Coord{
			{x, y, z + 1},
			{x, y, z - 1},

			{x, y + 1, z},
			{x, y - 1, z},

			{x + 1, y, z},
			{x - 1, y, z},
		}
		for _, next_point := range next_points {
			_, is_visited := (*visited)[next_point]
			if is_visited {
				continue
			}
			intersections := get_intersections(next_point, nanobots)
			if intersections >= current_intersections {
				current_intersections = intersections
				current_man := man_to_origin(next_point)
				// fmt.Println("    ", next_point, current_intersections, current_man)
				// local_min_man = min(local_min_man, current_man)
				if current_man < local_min_man {
					local_min_man = current_man
					min_point = next_point
				}
			}

			(*visited)[next_point] = true
			queue = append(queue, next_point)
			// fmt.Println("    ", next_point, current_intersections, local_max_man)
		}

		max_bfs_rec--
	}
	return local_min_man, current_intersections, min_point
}

func solve(nanobots []Nanobot) int {
	// var best_coord = Coord{0, 0, 0}
	// max_num_nanobots := math.MinInt

	max_intersections := math.MinInt
	intersections := make(map[Coord]int)
	for i := 0; i < len(nanobots); i++ {
		corners := [][]int{
			{nanobots[i].x, nanobots[i].y, nanobots[i].z + nanobots[i].radius},
			{nanobots[i].x, nanobots[i].y, nanobots[i].z - nanobots[i].radius},

			{nanobots[i].x, nanobots[i].y + nanobots[i].radius, nanobots[i].z},
			{nanobots[i].x, nanobots[i].y - nanobots[i].radius, nanobots[i].z},

			{nanobots[i].x + nanobots[i].radius, nanobots[i].y, nanobots[i].z},
			{nanobots[i].x - nanobots[i].radius, nanobots[i].y, nanobots[i].z},

			{nanobots[i].x, nanobots[i].y, nanobots[i].z},
		}

		for _, corner := range corners {
			counter := 0
			for j := 0; j < len(nanobots); j++ {
				if is_inside(nanobots[j], corner[0], corner[1], corner[2]) {
					counter += 1
				}
			}
			value, seen_point_before := intersections[Coord{corner[0], corner[1], corner[2]}]
			if seen_point_before {
				fmt.Println("---------><--------------", value, counter)
			}
			intersections[Coord{corner[0], corner[1], corner[2]}] = counter
			max_intersections = max(max_intersections, counter)
		}
	}

	points := []Coord{}
	for point, _ := range intersections {
		points = append(points, point)
	}
	sort_points_by_intersections(points, intersections)

	//
	// TODO: take every point and do BFS
	// each point shound be equal or better intersections
	// for each point calculate distance to origin

	max_bfs_rec := 5_500_000
	// max_bfs_rec := 1_000
	max_list_proc := 1
	// max_list_proc := 6000
	min_man_distance := man_to_origin(points[0])
	max_intersections = intersections[points[0]]
	visited := make(map[Coord]bool)

	fmt.Println("max inter", max_intersections)
	fmt.Println("max man", man_to_origin(points[0]))

	for i := 0; i < max_list_proc; i++ {
		fmt.Println(points[i], intersections[points[i]])
		// visited := make(map[Coord]bool)

		local_min_man, current_intersections, min_point := bfs(
			points[i],
			&visited,
			// max_intersections,
			intersections[points[i]],
			min_man_distance,
			max_bfs_rec,
			nanobots,
		)
		if current_intersections >= max_intersections {
			max_intersections = current_intersections
			min_man_distance = min(min_man_distance, local_min_man)
			fmt.Println("--->", current_intersections, min_man_distance, min_point, len(visited))
		} else {
			fmt.Println("--><", current_intersections, min_man_distance, len(visited))
		}
	}

	return min_man_distance
}

func solution(filename string) int {
	nanobots := parse(filename)
	// fmt.Println("num nano", len(nanobots))
	return solve(nanobots)
}

func main() {
	// fmt.Println(solution("./example.txt")) // 36
	fmt.Println(solution("./input.txt")) //
	// 82444530 too high
}
