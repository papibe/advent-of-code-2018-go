package main

import (
	"fmt"
	"strconv"
	"strings"
)

func solution(rounds int) string {
	recipes := []int{3, 7}
	chef1 := 0
	chef2 := 1

	for len(recipes) < rounds+10 {
		score_sum := recipes[chef1] + recipes[chef2]
		next_score1 := score_sum / 10
		next_score2 := score_sum % 10
		if next_score1 > 0 {
			recipes = append(recipes, next_score1)
		}
		recipes = append(recipes, next_score2)
		chef1 = (chef1 + 1 + recipes[chef1]) % len(recipes)
		chef2 = (chef2 + 1 + recipes[chef2]) % len(recipes)
	}

	str_recipes := []string{}
	last_ten_recipes := recipes[rounds : rounds+10]
	for _, score := range last_ten_recipes {
		str_recipes = append(str_recipes, strconv.Itoa(score))
	}
	return strings.Join(str_recipes, "")
}

func main() {
	// fmt.Println(solution(5))      // 0124515891
	// fmt.Println(solution(9))      // 5158916779
	// fmt.Println(solution(18))     // 9251071085
	// fmt.Println(solution(2018))   // 5941429882
	fmt.Println(solution(503761)) // 1044257397
}
