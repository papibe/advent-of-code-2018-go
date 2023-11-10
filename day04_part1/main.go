package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
	"strconv"
)

func parse(filename string) []int {
	data, err := os.ReadFile(filename)

	if err != nil {
		panic("file error")
	}
	content := string(data)

	re_main := regexp.MustCompile(`\[(\d\d\d\d)-(\d\d)-(\d\d) (\d\d):(\d\d)\] (.*)$` )
	re_msg := regexp.MustCompile(`Guard #(\d+) begins shift`)

	for _, line := range strings.Split(strings.Trim(content, "\n"), "\n") {
		matches := re_main.FindStringSubmatch(line)

		year, _ := strconv.Atoi(matches[1])
		month, _ := strconv.Atoi(matches[2])
		day, _ := strconv.Atoi(matches[3])
		hour, _ := strconv.Atoi(matches[4])
		minute, _ := strconv.Atoi(matches[5])
		log_msg := matches[6]

		fmt.Println(line)
		fmt.Println(year, month, day, hour, minute)
		if log_msg == "falls asleep" {
			fmt.Println("\tFalls")
		} else if log_msg == "wakes up" {
			fmt.Println("\tWakes")
		} else {
			match := re_msg.FindStringSubmatch(log_msg)
			guard, _ := strconv.Atoi(match[1])
			fmt.Println("\tGuard", guard	)
		}
	}

	return []int{1, 2, 3}

}

func solution(filename string) int {
	data := parse(filename)
	fmt.Println(data)
	return 0
}

func main() {
	fmt.Println(solution("./example.txt"))
}