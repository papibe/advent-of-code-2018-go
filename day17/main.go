package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Coord struct {
	x int
	y int
}

func parse(filename string) (map[Coord]bool, int, int) {
	raw_data, err := os.ReadFile(filename)
	if err != nil {
		panic("File not found")
	}
	data := strings.Trim(string(raw_data), "\n")

	re_x := regexp.MustCompile(`x=(\d+), y=(\d+)\.\.(\d+)`)
	re_y := regexp.MustCompile(`y=(\d+), x=(\d+)\.\.(\d+)`)

	min_y := math.MaxInt
	max_y := math.MinInt
	reservoir := make(map[Coord]bool)
	for _, line := range strings.Split(data, "\n") {

		// try re_X
		x_matches := re_x.FindStringSubmatch(line)
		if len(x_matches) > 0 {
			x, _ := strconv.Atoi(x_matches[1])
			y1, _ := strconv.Atoi(x_matches[2])
			y2, _ := strconv.Atoi(x_matches[3])
			for y := y1; y <= y2; y++ {
				reservoir[Coord{x, y}] = true
				min_y = min(min_y, y)
				max_y = max(max_y, y)
			}
			continue
		}

		// try re_y
		y_matches := re_y.FindStringSubmatch(line)
		if len(y_matches) > 0 {
			y, _ := strconv.Atoi(y_matches[1])
			x1, _ := strconv.Atoi(y_matches[2])
			x2, _ := strconv.Atoi(y_matches[3])
			min_y = min(min_y, y)
			max_y = max(max_y, y)
			for x := x1; x <= x2; x++ {
				reservoir[Coord{x, y}] = true
			}
			continue
		}
	}
	return reservoir, min_y, max_y
}

func go_side(reservoir map[Coord]bool, x, y int, wet, rest *map[Coord]bool, step int) (Coord, bool) {
	for {
		_, lower_left_is_clay := reservoir[Coord{x + step, y + 1}]
		_, lower_left_is_settled := (*rest)[Coord{x + step, y + 1}]
		if !lower_left_is_clay && !lower_left_is_settled {
			return Coord{x + step, y}, false
		}

		_, next_is_clay := reservoir[Coord{x + step, y}]
		_, next_is_settled := (*rest)[Coord{x + step, y}]
		if next_is_clay || next_is_settled {
			return Coord{x, y}, true
		}
		x += step
	}
}

func go_left(reservoir map[Coord]bool, x, y int, wet, rest *map[Coord]bool) (Coord, bool) {
	return go_side(reservoir, x, y, wet, rest, -1)
}

func go_right(reservoir map[Coord]bool, x, y int, wet, rest *map[Coord]bool) (Coord, bool) {
	return go_side(reservoir, x, y, wet, rest, 1)
}

func solve(reservoir map[Coord]bool, min_y, max_y int, wet, rest *map[Coord]bool) (int, int) {
	counter := 0
	stack := []Coord{{500, 1}}
	if 1 >= min_y {
		(*wet)[Coord{500, 1}] = true
	}
	for len(stack) > 0 {

		// pop stack
		coord := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		// check status for lower position
		lower := Coord{coord.x, coord.y + 1}
		_, lower_is_clay := reservoir[lower]
		_, lower_is_at_rest := (*rest)[lower]
		_, lower_is_wet := (*wet)[lower]

		if lower.y > max_y {
			continue
		}

		if !lower_is_clay && !lower_is_at_rest && !lower_is_wet {
			if lower.y >= min_y {
				(*wet)[lower] = true
			}
			stack = append(stack, coord) // for later processing
			stack = append(stack, lower)
			continue
		}

		// hit clay or water at rest
		if lower_is_clay || lower_is_at_rest {
			rcoord, rsettled := go_right(reservoir, coord.x, coord.y, wet, rest)
			lcoord, lsettled := go_left(reservoir, coord.x, coord.y, wet, rest)

			if lsettled && rsettled {
				for row := lcoord.x; row <= rcoord.x; row++ {
					(*rest)[Coord{row, coord.y}] = true
					(*wet)[Coord{row, coord.y}] = true
				}
			} else {
				for row := lcoord.x; row <= rcoord.x; row++ {
					(*wet)[Coord{row, coord.y}] = true
				}
				if !rsettled {
					stack = append(stack, rcoord)
				}
				if !lsettled {
					stack = append(stack, lcoord)
				}
			}
		}

		counter += 1
		if counter > 6500 {
			fmt.Println("max counter")
			break
		}

	}
	return len(*wet), len(*rest)
}

func solution(filename string) (int, int) {
	reservoir, min_y, max_y := parse(filename)
	wet1 := make(map[Coord]bool)
	settled1 := make(map[Coord]bool)
	wet := &wet1
	settled := &settled1
	return solve(reservoir, min_y, max_y, wet, settled)
}

func main() {
	fmt.Println(solution("example.txt")) // 57
	fmt.Println(solution("input.txt"))   // 37858, 30410
}
