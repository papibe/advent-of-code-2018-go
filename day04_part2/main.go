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
	date     int
	kind     entryType
	guard_id int
}

func parse(filename string) []LogEntry {
	data, err := os.ReadFile(filename)

	if err != nil {
		panic("file error")
	}
	content := string(data)

	re_main := regexp.MustCompile(`\[(\d\d\d\d)-(\d\d)-(\d\d) (\d\d):(\d\d)\] (.*)$`)
	re_msg := regexp.MustCompile(`Guard #(\d+) begins shift`)

	log_entries := []LogEntry{}

	for _, line := range strings.Split(strings.Trim(content, "\n"), "\n") {
		matches := re_main.FindStringSubmatch(line)

		month, _ := strconv.Atoi(matches[2])
		day, _ := strconv.Atoi(matches[3])
		hour, _ := strconv.Atoi(matches[4])
		minute, _ := strconv.Atoi(matches[5])
		log_msg := matches[6]

		entry := Starts
		guard := 0 // invalid guard if AFAIK

		if log_msg == "falls asleep" {
			entry = Falls
		} else if log_msg == "wakes up" {
			entry = Wakes
		} else {
			match := re_msg.FindStringSubmatch(log_msg)
			guard, _ = strconv.Atoi(match[1])
			entry = Starts
		}
		epoch := month*1000000 + day*10000 + hour*100 + minute
		log_entries = append(log_entries, LogEntry{date: epoch, kind: entry, guard_id: guard})
	}

	return log_entries
}

func strategy2(log_entries []LogEntry) int {
	guards_sleep := make(map[int]*[60]int)
	current_guard := 0
	falls_sleep_time := 0

	for _, entry := range log_entries {
		switch entry.kind {
		case Starts:
			_, ok := guards_sleep[entry.guard_id]
			if !ok {
				guards_sleep[entry.guard_id] = &[60]int{}
			}
			current_guard = entry.guard_id
		case Falls:
			falls_sleep_time = entry.date % 100
		case Wakes:
			wakes_up_time := entry.date % 100
			for minute := falls_sleep_time; minute < wakes_up_time; minute++ {
				guards_sleep[current_guard][minute] += 1
			}
		default:
			panic("blah!")

		}
	}

	sleepier_guard := 0
	max_slept_minute := 0
	max_sleep := 0
	for guard, sleep_record := range guards_sleep {
		for minute, sleep_count := range sleep_record {
			if sleep_count > max_sleep {
				max_sleep = sleep_count
				max_slept_minute = minute
				sleepier_guard = guard
			}
		}
	}

	return sleepier_guard * max_slept_minute
}

func solution(filename string) int {
	log_entries := parse(filename)

	// sort log entries by date
	sort.Slice(log_entries, func(i, j int) bool {
		return log_entries[i].date < log_entries[j].date
	})

	return strategy2(log_entries)
}

func main() {
	fmt.Println(solution("./input.txt"))
}
