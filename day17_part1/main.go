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

func parse(filename string) (map[Coord]bool, int) {
	raw_data, err := os.ReadFile(filename)
	if err != nil {
		panic("File not found")
	}
	data := strings.Trim(string(raw_data), "\n")

	re_x := regexp.MustCompile(`x=(\d+), y=(\d+)\.\.(\d+)`)
	re_y := regexp.MustCompile(`y=(\d+), x=(\d+)\.\.(\d+)`)

	max_y := 0
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
				if y > max_y {
					max_y = y
				}
			}
			continue
		}

		// try re_y
		y_matches := re_y.FindStringSubmatch(line)
		if len(y_matches) > 0 {
			y, _ := strconv.Atoi(y_matches[1])
			x1, _ := strconv.Atoi(y_matches[2])
			x2, _ := strconv.Atoi(y_matches[3])
			if y > max_y {
				max_y = y
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
	return reservoir, max_y
}

func go_right(reservoir map[Coord]bool, x, y, max_y int, wet, settled *map[Coord]bool) (Coord, bool) {
	for {
		(*wet)[Coord{x, y}] = true
		_, lower_right_is_clay := reservoir[Coord{x + 1, y + 1}]
		_, lower_right_is_settled := (*settled)[Coord{x + 1, y + 1}]
		if !lower_right_is_clay && !lower_right_is_settled {
			// fmt.Println("gright", x, y)
			// panic("what?")
			// go_down(reservoir, x+1, y, max_y, wet, settled)
			return Coord{x + 1, y}, false
		}

		_, next_is_clay := reservoir[Coord{x + 1, y}]
		_, next_is_settled := (*settled)[Coord{x + 1, y}]
		if next_is_clay || next_is_settled {
			// (*settled)[Coord{x, y}] = true
			return Coord{x, y}, true
		}
		x += 1
	}
}

func go_left(reservoir map[Coord]bool, x, y, max_y int, wet, settled *map[Coord]bool) (Coord, bool) {
	for {
		(*wet)[Coord{x, y}] = true
		_, lower_left_is_clay := reservoir[Coord{x - 1, y + 1}]
		_, lower_left_is_settled := (*settled)[Coord{x - 1, y + 1}]
		if !lower_left_is_clay && !lower_left_is_settled {
			// fmt.Println("gleft", x, y)
			// panic("what?")
			// go_down(reservoir, x-1, y, max_y, wet, settled)
			return Coord{x - 1, y}, false
		}

		_, next_is_clay := reservoir[Coord{x - 1, y}]
		_, next_is_settled := (*settled)[Coord{x - 1, y}]
		if next_is_clay || next_is_settled {
			// (*settled)[Coord{x, y}] = true
			return Coord{x, y}, true
		}
		x -= 1
	}
}

func go_down(reservoir map[Coord]bool, x, y, max_y int, wet, settled *map[Coord]bool) {
	for {
		if y > max_y {
			return
		}
		(*wet)[Coord{x, y}] = true
		_, next_is_clay := reservoir[Coord{x, y + 1}]
		_, next_is_settled := (*settled)[Coord{x, y + 1}]
		if next_is_clay || next_is_settled {
			// fmt.Println("floor", x, y)
			rcoord, rsettled := go_right(reservoir, x, y, max_y, wet, settled)
			lcoord, lsettled := go_left(reservoir, x, y, max_y, wet, settled)

			// fmt.Println("row", y, "lsettled", lsettled, "rsettled", rsettled)

			if lsettled && rsettled {
				// fmt.Println("settling row", y)
				for row := lcoord.x; row <= rcoord.x; row++ {
					(*settled)[Coord{row, y}] = true
				}
			} else {
				if !rsettled {
					// fmt.Println("going down right", rcoord)
					go_down(reservoir, rcoord.x, rcoord.y, max_y, wet, settled)
				}
				if !lsettled {
					// fmt.Println("going down left", lcoord)
					go_down(reservoir, lcoord.x, lcoord.y, max_y, wet, settled)
				}
			}

			return
		}

		y += 1
	}
}

func reservoir_print(reservoir map[Coord]bool, wet, settled *map[Coord]bool) {
	max_x := -1
	min_x := 10000
	max_y := -1
	min_y := 10000

	for coord, _ := range reservoir {
		min_x = min(min_x, coord.x)
		max_x = max(max_x, coord.x)
		min_y = min(min_y, coord.y)
		max_y = max(max_y, coord.y)
	}
	fmt.Println(min_x, max_x, min_y, max_y)
	for row := min_y; row <= max_y; row++ {
		for col := min_x; col <= max_x; col++ {
			coord := Coord{col, row}
			_, is_clay := reservoir[coord]
			_, is_settled := (*settled)[coord]
			_, is_wet := (*wet)[coord]

			// fmt.Println(is_clay, is_settled, is_wet)
			output := "."

			if is_clay {
				output = "#"
			}
			if is_wet {
				output = "|"
			}
			if is_settled {
				output = "~"
			}
			fmt.Print(output)

		}
		fmt.Println()
	}
}

func solve(reservoir map[Coord]bool, max_y int, wet, settled *map[Coord]bool) int {
	counter := 0
	for {
		current_wet := len(*wet)
		go_down(reservoir, 500, 1, max_y, wet, settled)

		// fmt.Println("len(wet)", len(*wet))
		// fmt.Println("len(settled)", len(*settled))
		// fmt.Println()
		counter += 1
		if counter > 500 {
			fmt.Println("max counter")
			break
		}

		// reservoir_print(reservoir, wet, settled)

		if len(*wet) == current_wet {
			fmt.Println("finishing?")
			break
		}
	}
	// for coord, _ := range *wet {
	// 	fmt.Println("wet", coord)
	// }
	return len(*wet)
}

func solution(filename string) int {
	reservoir, max_y := parse(filename)
	wet1 := make(map[Coord]bool)
	settled1 := make(map[Coord]bool)
	wet := &wet1
	settled := &settled1
	return solve(reservoir, max_y, wet, settled)
}

func main() {
	fmt.Println(solution("example.txt")) // 57
	// fmt.Println(solution("input.txt")) //
}
