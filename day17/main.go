package main

import (
	"fmt"
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

	min_y := 1000000
	max_y := 0
	reservoir := make(map[Coord]bool)
	for _, line := range strings.Split(data, "\n") {

		// try re_X
		x_matches := re_x.FindStringSubmatch(line)
		if len(x_matches) > 0 {
			x, _ := strconv.Atoi(x_matches[1])
			y1, _ := strconv.Atoi(x_matches[2])
			y2, _ := strconv.Atoi(x_matches[3])
			if y1 > y2 {
				panic("wrong order")
			}
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
			if x1 > x2 {
				panic("wrong order")
			}
			for x := x1; x <= x2; x++ {
				reservoir[Coord{x, y}] = true
			}
			continue
		}
		panic("the what?")
	}

	// fmt.Println(reservoir)
	// fmt.Println(len(reservoir))
	return reservoir, min_y, max_y
}

func reservoir_print(reservoir map[Coord]bool, wet, settled *map[Coord]bool) {
	max_x := -1
	min_x := 10000
	max_y := -1
	min_y := 10000

	total := 0

	for coord, _ := range reservoir {
		min_x = min(min_x, coord.x)
		max_x = max(max_x, coord.x)
		min_y = min(min_y, coord.y)
		max_y = max(max_y, coord.y)
	}
	fmt.Println(min_x, max_x, min_y, max_y)
	// for y := min_y - 1; y <= max_y+1; y++ {
	// for y := 1; y <= max_y+1; y++ {
	// 	for x := min_x - 1; x <= max_x+1; x++ {
	// 		coord := Coord{x, y}
	// 		_, is_clay := reservoir[coord]
	// 		_, is_settled := (*settled)[coord]
	// 		_, is_wet := (*wet)[coord]

	// 		// fmt.Println(is_clay, is_settled, is_wet)
	// 		output := "."

	// 		if is_clay {
	// 			output = "#"
	// 		}
	// 		if is_wet {
	// 			output = "|"
	// 		}
	// 		if is_settled {
	// 			output = "~"
	// 		}
	// 		if output != "#" && output != "." {
	// 			total += 1
	// 		}
	// 		fmt.Print(output)

	// 	}
	// 	fmt.Println()
	// }
	fmt.Println("total", total)
}

func go_right(reservoir map[Coord]bool, x, y, max_y int, wet, rest *map[Coord]bool) (Coord, bool) {
	for {
		// (*wet)[Coord{x, y}] = true
		_, lower_right_is_clay := reservoir[Coord{x + 1, y + 1}]
		_, lower_right_is_settled := (*rest)[Coord{x + 1, y + 1}]
		if !lower_right_is_clay && !lower_right_is_settled {
			// fmt.Println("gright", x, y)
			// panic("what?")
			// go_down(reservoir, x+1, y, max_y, wet, settled)
			_, next_is_clay := reservoir[Coord{x + 1, y}]
			if next_is_clay {
				panic("counting clay")
			}
			// (*wet)[Coord{x + 1, y}] = true
			return Coord{x + 1, y}, false
		}

		_, next_is_clay := reservoir[Coord{x + 1, y}]
		_, next_is_settled := (*rest)[Coord{x + 1, y}]
		if next_is_clay || next_is_settled {
			// (*settled)[Coord{x, y}] = true
			return Coord{x, y}, true
		}
		x += 1
	}
}

func go_left(reservoir map[Coord]bool, x, y, max_y int, wet, rest *map[Coord]bool) (Coord, bool) {
	for {
		// (*wet)[Coord{x, y}] = true
		_, lower_left_is_clay := reservoir[Coord{x - 1, y + 1}]
		_, lower_left_is_settled := (*rest)[Coord{x - 1, y + 1}]
		if !lower_left_is_clay && !lower_left_is_settled {
			// fmt.Println("gleft", x, y)
			// panic("what?")
			// go_down(reservoir, x-1, y, max_y, wet, settled)
			_, next_is_clay := reservoir[Coord{x - 1, y}]
			if next_is_clay {
				panic("counting clay")
			}
			// (*wet)[Coord{x - 1, y}] = true
			return Coord{x - 1, y}, false
		}

		_, next_is_clay := reservoir[Coord{x - 1, y}]
		_, next_is_settled := (*rest)[Coord{x - 1, y}]
		if next_is_clay || next_is_settled {
			// (*settled)[Coord{x, y}] = true
			return Coord{x, y}, true
		}
		x -= 1
	}
}

func solve(reservoir map[Coord]bool, min_y, max_y int, wet, rest *map[Coord]bool) (int, int) {
	counter := 0
	stack := []Coord{{500, 1}}
	if 1 >= min_y {
		(*wet)[Coord{500, 1}] = true
	}
	for len(stack) > 0 {
		// reservoir_print(reservoir, wet, rest)

		// pop stack
		coord := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		// fmt.Println("pop", coord)
		// fmt.Println("stack", stack)

		lower := Coord{coord.x, coord.y + 1}
		_, lower_is_clay := reservoir[lower]
		_, lower_is_at_rest := (*rest)[lower]
		_, lower_is_wet := (*wet)[lower]

		if lower.y > max_y {
			// reservoir_print(reservoir, wet, rest)
			continue
		}

		// if lower is free go down
		if !lower_is_clay && !lower_is_at_rest && !lower_is_wet {
			// fmt.Println("going down")
			if lower.y >= min_y {

				(*wet)[lower] = true
			}
			stack = append(stack, coord) // for later processing
			stack = append(stack, lower)
			// reservoir_print(reservoir, wet, rest)
			continue
		}

		// hit clay or water at rest
		if lower_is_clay || lower_is_at_rest {
			// fmt.Println("hit something")
			rcoord, rsettled := go_right(reservoir, coord.x, coord.y, max_y, wet, rest)
			lcoord, lsettled := go_left(reservoir, coord.x, coord.y, max_y, wet, rest)

			// fmt.Println(rsettled, lsettled)
			if lsettled && rsettled {
				// fmt.Println("settling row", coord.y)
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
		// reservoir_print(reservoir, wet, rest)

		counter += 1
		if counter > 7000 {
			// fmt.Println("max counter")
			break
		}

	}

	// no clay is wet
	for coord, _ := range reservoir {
		_, is_in_wet := (*wet)[coord]
		if is_in_wet {
			panic("clay in wet")
		}
	}
	// no wet is clay
	for coord, _ := range *wet {
		_, is_clay := reservoir[coord]
		if is_clay {
			panic("clay in wet")
		}
	}
	// all rest are wet
	for coord, _ := range *rest {
		_, is_in_wet := (*wet)[coord]
		if !is_in_wet {
			panic("rest part is not wet")
		}
	}
	// not wet is out of range
	for coord, _ := range *wet {
		if coord.y < min_y || coord.y > max_y {
			panic("<---NOT---->")
		}
	}
	// reservoir_print(reservoir, wet, rest)

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
	// 474 too low
	// 7134 too low
	// 37863 too high
}
