package main

import (
	"fmt"
	"math"
)

// type Coord struct {
// 	x int
// 	y int
// 	z int
// }

func bla() {
	x0 := 10
	y0 := 12
	z0 := 2
	r := 3

	points := make(map[Coord]bool)
	for x := x0; x <= x0+r; x++ {
		max_y := y0 + r + x0 - x

		for y := y0; y <= max_y; y++ {
			d := int(float64(r) - math.Abs(float64(x0)-float64(x)) - math.Abs(float64(y0)-float64(y)))

			// fmt.Println(x, y, d)
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

	// check distances
	for p, _ := range points {
		d := math.Abs(float64(x0)-float64(p.x)) + math.Abs(float64(y0)-float64(p.y)) + math.Abs(float64(z0)-float64(p.z))
		if int(d) > r {
			fmt.Println("error at ", p, r, d)
		}
	}

	// for z := -2; z <= 6; z++ {
	// 	fmt.Println("z", z, "=====================")
	// 	for y := 25; y >= 0; y-- {
	// 		for x := 0; x < 25; x++ {
	// 			_, is_point := points[Coord{x, y, z}]
	// 			if is_point {
	// 				fmt.Print("#")
	// 			} else {
	// 				fmt.Print(".")
	// 			}
	// 		}
	// 		fmt.Println()
	// 	}
	// }
	// fmt.Println(points)
}
