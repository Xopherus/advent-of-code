package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"time"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Panicf("could not open file %s", err)
	}

	scanner := bufio.NewScanner(f)

	// sort logs by timestamp asc
	var logs []string
	for scanner.Scan() {
		logs = append(logs, scanner.Text())
	}
	sort.Strings(logs)

	schedule := make(map[int][]int)

	var guardOnDuty int
	var sleepChange time.Time

	for _, l := range logs {
		ts, err := time.Parse("2006-01-02 15:04", l[1:17])
		if err != nil {
			log.Panicf("could not parse timestamp %s", err)
		}
		action := l[19:]

		switch action {
		case "falls asleep":
			sleepChange = ts
		case "wakes up":
			for i := sleepChange.Minute(); i < ts.Minute(); i++ {
				schedule[guardOnDuty][i]++
			}
		default:
			// change guard on duty
			if _, err := fmt.Sscanf(action, "Guard #%d begins shift", &guardOnDuty); err != nil {
				log.Panicf("could not get guard %s", err)
			}
			if _, ok := schedule[guardOnDuty]; !ok {
				schedule[guardOnDuty] = make([]int, 60)
			}
		}
	}

	// part 1
	part1(schedule)
	part2(schedule)
}

func part1(schedule map[int][]int) {
	guard, total, minute := 0, 0, 0

	for g, minutes := range schedule {
		sum, id, max := 0, 0, 0

		for i := range minutes {
			if minutes[i] > max {
				id, max = i, minutes[i]
			}
			sum += minutes[i]
		}

		if sum > total {
			guard, total, minute = g, sum, id
		}
	}

	fmt.Printf("Product %d * %d = %d\n", guard, minute, guard*minute)
}

// Of all guards, which guard is most frequently asleep on the same minute?
func part2(schedule map[int][]int) {
	max, minute, guard := 0, 0, 0

	for i := 0; i < 60; i++ {
		localMax, id := 0, 0

		for g, minutes := range schedule {
			if minutes[i] > localMax {
				localMax, id = minutes[i], g
			}
		}

		if localMax > max {
			max, minute, guard = localMax, i, id
		}
	}

	fmt.Printf("Product %d * %d = %d\n", guard, minute, guard*minute)
}
