package main

import (
	"fmt"
)

func get_scores(str_scores string) []int {
	scores := []int{}
	for _, char := range str_scores {
		scores = append(scores, int(char-'0'))
	}
	return scores
}

func appears(scores, recipes []int, recipe_index int) bool {
	for score_index := len(scores) - 1; score_index >= 0; score_index-- {
		if scores[score_index] != recipes[recipe_index] {
			return false
		}
		recipe_index--
	}
	return true
}

func solve(scores []int) int {
	recipes := []int{3, 7}
	chef1 := 0
	chef2 := 1

	for {
		score_sum := recipes[chef1] + recipes[chef2]
		next_score1 := score_sum / 10
		next_score2 := score_sum % 10
		if next_score1 > 0 {
			recipes = append(recipes, next_score1)
		}
		recipes = append(recipes, next_score2)
		chef1 = (chef1 + 1 + recipes[chef1]) % len(recipes)
		chef2 = (chef2 + 1 + recipes[chef2]) % len(recipes)

		if appears(scores, recipes, len(recipes)-1) {
			return len(recipes) - len(scores)
		}
		if appears(scores, recipes, len(recipes)-2) {
			return len(recipes) - len(scores) - 1
		}
	}
}

func solution(str_scores string) int {
	scores := get_scores(str_scores)
	return solve(scores)
}

func main() {
	// fmt.Println(solution("01245")) // 5
	// fmt.Println(solution("51589")) // 9
	// fmt.Println(solution("92510")) // 18
	// fmt.Println(solution("59414")) // 1018
	fmt.Println(solution("503761")) // 20185425
}
