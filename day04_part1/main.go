package main

import (
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type entryType string

const (
	Falls  entryType = "F"
	Wakes  entryType = "W"
	Starts entryType = "S"
)

type LogEntry struct {
	date  string
	kind  entryType
	guard string
}

// type GuardEntries struct {
// 	name       string
// 	sleep_mins [60]int
// }

func parse(filename string) []LogEntry {
	data, err := os.ReadFile(filename)

	if err != nil {
		panic("file error")
	}
	content := string(data)

	re_main := regexp.MustCompile(`\[(\d\d\d\d)-(\d\d)-(\d\d) (\d\d):(\d\d)\] (.*)$`)
	re_msg := regexp.MustCompile(`Guard #(\d+) begins shift`)

	entries := []LogEntry{}

	for _, line := range strings.Split(strings.Trim(content, "\n"), "\n") {
		matches := re_main.FindStringSubmatch(line)

		// year, _ := strconv.Atoi(matches[1])
		// month, _ := strconv.Atoi(matches[2])
		month := matches[2]
		// day, _ := strconv.Atoi(matches[3])
		day := matches[3]
		hour := matches[4]
		minute := matches[5]
		log_msg := matches[6]

		entry := Starts
		guard := "N/A"

		// fmt.Println(line)
		// fmt.Print("\t", year, month, day, hour, minute)
		if log_msg == "falls asleep" {
			entry = Falls
			// fmt.Println("\tFalls")
		} else if log_msg == "wakes up" {
			entry = Wakes
			// fmt.Println("\tWakes")
		} else {
			match := re_msg.FindStringSubmatch(log_msg)
			// guard, _ := strconv.Atoi(match[1])
			guard = match[1]
			// fmt.Println("\tGuard", guard)

			entry = Starts
		}
		// epoc := month*1000000 + day*1000 + hour*100 + minute
		entries = append(entries, LogEntry{date: month + day + hour + minute, kind: entry, guard: guard})
	}

	return entries

}

func solution(filename string) int {
	log_entries := parse(filename)

	// fmt.Println(log_entries)

	sort.Slice(log_entries, func(i, j int) bool {
		return log_entries[i].date < log_entries[j].date
	})
	// fmt.Println("================================================")
	// fmt.Println(log_entries)

	guards_sleep := make(map[string]*[60]int)

	current_guard := ""
	_ = current_guard
	falls_sleep_time := 0

	for _, entry := range log_entries {
		switch entry.kind {
		case Starts:
			_, ok := guards_sleep[entry.guard]
			if !ok {
				guards_sleep[entry.guard] = &[60]int{}
			}
			current_guard = entry.guard
		case Falls:
			falls_sleep_time, _ = strconv.Atoi(entry.date)
			falls_sleep_time %= 100
		case Wakes:
			wakes_up_time, _ := strconv.Atoi(entry.date)
			wakes_up_time %= 100
			// fmt.Println(current_guard, falls_sleep_time, wakes_up_time)
			for minute := falls_sleep_time; minute < wakes_up_time; minute++ {
				// fmt.Println(current_guard, falls_sleep_time, wakes_up_time, minute)
				guards_sleep[current_guard][minute] += 1
			}

		default:
			panic("blah!")

		}
	}

	// for k, v := range guards_sleep {
	// 	fmt.Println(k, v)
	// }

	sleepier_guard := ""
	max_sleep_mins := 0
	for k, v := range guards_sleep {
		// fmt.Println("checking guard", k)
		guard_total_sleep := 0
		for _, value := range v {
			guard_total_sleep += value
		}
		// fmt.Println(k, guard_total_sleep)
		if guard_total_sleep > max_sleep_mins {
			max_sleep_mins = guard_total_sleep
			sleepier_guard = k
		}
	}
	// fmt.Println("sleepier is", sleepier_guard, "with", max_sleep_mins)

	max_worst_minute := 0
	max_sleep := 0
	for minute, sleep_count := range guards_sleep[sleepier_guard] {
		if sleep_count > max_sleep {
			max_sleep = sleep_count
			max_worst_minute = minute
		}
	}

	guard, _ := strconv.Atoi(sleepier_guard)
	fmt.Println("Part1: ", guard*max_worst_minute)

	max_worst_minute = 0
	max_sleep = 0
	sleepier_guard = ""
	for guard, sleep_record := range guards_sleep {
		for minute, sleep_count := range sleep_record {
			if sleep_count > max_sleep {
				max_sleep = sleep_count
				max_worst_minute = minute
				sleepier_guard = guard
			}
		}
	}

	guard, _ = strconv.Atoi(sleepier_guard)

	fmt.Println("Part2: ", guard*max_worst_minute)

	return guard * max_worst_minute
}

func main() {
	// fmt.Println(solution("./example.txt"))
	fmt.Println(solution("./input.txt"))
}
