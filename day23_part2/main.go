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

func BronKerbosch(r, p, x map[int]bool, graph []map[int]bool, result *[][]int) {
	if len(p) == 0 && len(x) == 0 {
		clique := []int{}
		for v := range r {
			clique = append(clique, v)
		}
		*result = append(*result, clique)
		return
	}
	// pivoting
	var u int
	if len(p) > 0 {
		for v := range p {
			u = v
			break
		}
	} else {
		for v := range x {
			u = v
			break
		}
	}
	p_minus_nu := make(map[int]bool)
	for v := range p {
		_, is_connected_to_v := graph[u][v]
		if !is_connected_to_v {
			p_minus_nu[v] = true
		}
	}
	for v := range p_minus_nu {
		new_r := make(map[int]bool)
		for key, value := range r {
			new_r[key] = value
		}
		new_r[v] = true

		new_p := make(map[int]bool)
		for key, value := range p {
			_, is_connected_to_v := graph[v][key]
			if is_connected_to_v {
				new_p[key] = value
			}
		}

		new_x := make(map[int]bool)
		for key, value := range x {
			_, is_connected_to_v := graph[v][key]
			if is_connected_to_v {
				new_x[key] = value
			}
		}

		BronKerbosch(new_r, new_p, new_x, graph, result)
		delete(p, v)
		x[v] = true
	}

}

func get_max_cliques(graph []map[int]bool) []int {
	r := make(map[int]bool)
	x := make(map[int]bool)

	p := make(map[int]bool)
	for i := 0; i < len(graph); i++ {
		p[i] = true
	}
	cliques := [][]int{}
	BronKerbosch(r, p, x, graph, &cliques)
	// get bigest cliques
	max_len := math.MinInt
	max_index := -1
	for index, clique := range cliques {
		if len(clique) > max_len {
			max_len = len(clique)
			max_index = index
		}
	}
	return cliques[max_index]
}

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}

func intersects(n1, n2 Nanobot) bool {
	distance := abs(n1.x-n2.x) + abs(n1.y-n2.y) + abs(n1.z-n2.z)
	return distance <= n1.radius+n2.radius
}

func form_graph(nanobots []Nanobot) []map[int]bool {
	graph := []map[int]bool{}
	for i, nbot1 := range nanobots {
		connections := make(map[int]bool)
		for j, nbot2 := range nanobots {
			if i == j {
				continue
			}
			if intersects(nbot1, nbot2) {
				connections[j] = true
			}
		}
		graph = append(graph, connections)
	}
	return graph
}

func solution(filename string) int {
	nanobots := parse(filename)
	graph := form_graph(nanobots)
	clique := get_max_cliques(graph)

	max_distance := math.MinInt
	for _, bot_index := range clique {
		nanobot := nanobots[bot_index]
		manhatan := abs(nanobot.x) + abs(nanobot.y) + abs(nanobot.z)
		distance := manhatan - nanobot.radius
		max_distance = max(max_distance, distance)
	}

	return max_distance
}

func main() {
	fmt.Println(solution("./input.txt")) // 78687716
}
