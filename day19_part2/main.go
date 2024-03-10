package main

import (
	"fmt"
)

func solution(number int) int {
	// the program does this
	//
	// 	factor_sum := 0
	// 	for i := 1; i <= number; i++ {
	// 		for j := 1; j <= number; j++ {
	// 			if i*j == number {
	// 				factor_sum += i
	// 			}
	// 		}
	// 	}
	// 	return factor_sum
	// }

	// which refactored is
	factor_sum := 0
	for i := 1; i <= number; i++ {
		if number%i == 0 {
			factor_sum += i
		}
	}
	return factor_sum
}

func main() {
	fmt.Println(solution(10551430)) // 18992592
}
