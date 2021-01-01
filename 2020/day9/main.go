// https://adventofcode.com/2020/day/9
package main

import (
	"bufio"
	"container/list"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

// twoSum returns true if there are 2 numbers in the list that add to sum.
func twoSum(numbers *list.List, sum int) bool {
	diffs := make(map[int]int)

	for elem := numbers.Front(); elem != nil; elem = elem.Next() {
		a := elem.Value.(int)

		if _, ok := diffs[a]; ok {
			return true
		}

		diffs[sum-a] = a
	}

	return false
}

func part1() int {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("could not open file: %s", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	// use container/list like a FIFO queue
	numbers := list.New()

	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatalf("text is not a number!")
		}

		// is list full? if so, check if number is valid
		if numbers.Len() == 25 {
			if !twoSum(numbers, num) {
				return num
			}

			numbers.Remove(numbers.Front())
		}

		// if it is, add to end of list
		numbers.PushBack(num)
	}

	return 0
}

func add(numbers []int) int {
	var t int

	for _, n := range numbers {
		t += n
	}

	return t
}

func part2(sum int) (int, error) {
	f, err := os.Open("input.txt")
	if err != nil {
		return 0, fmt.Errorf("could not open file: %w", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	numbers := []int{}

	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return 0, fmt.Errorf("text is not a number!")
		}

		numbers = append(numbers, num)
	}

	// prompt says contiguous range must have at least 2 numbers, so we'll start there
	windowSize := 2

	for {
		log.Printf("trying window size = %d", windowSize)

		if windowSize > len(numbers) {
			return 0, fmt.Errorf("no contiguous set of numbers which add to sum")
		}

		for i := 0; i < len(numbers); i++ {
			if i+windowSize > len(numbers) {
				continue
			}

			subset := numbers[i : i+windowSize]

			// log.Printf("subset [%d-%d]: %+v sum = %d", i, i+windowSize, subset, add(subset))

			if add(subset) == sum {
				sort.Ints(subset)
				return subset[0] + subset[len(subset)-1], nil
			}
		}

		windowSize++
	}
}

func main() {
	answer := part1()
	log.Printf("part1: first invalid number %d", answer)

	answer, err := part2(answer)
	if err != nil {
		log.Fatalf("%s", err)
	}
	log.Printf("part2: sum of smallest + largest number %d", answer)
}
