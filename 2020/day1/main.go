// https://adventofcode.com/2020/day/1
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

// readInts is a helper function to convert input to a list of numbers
func readInts(filename string) ([]int, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("could not open file: %w", err)
	}

	scanner := bufio.NewScanner(f)
	nums := []int{}

	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, fmt.Errorf("could not read input: %w", err)
		}

		nums = append(nums, num)
	}

	return nums, nil
}

// After reading prompt, this seems like a classic two-sum algorithm question.
// Good thing I've been giving so many interviews lately...
func part1() (int, int) {
	nums, err := readInts("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	// for every num A in the list,
	// check if A exists in the map. if it doesn't exist, save remainder in map (2020 - A)
	// if it does, we know A and map[A] sum to 2020.
	diffs := make(map[int]int)

	for _, a := range nums {
		if b, ok := diffs[a]; ok {
			return a, b
		}

		diffs[2020-a] = a
	}

	return 0, 0
}

func part2() (int, int, int) {
	nums, err := readInts("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	sum := 2020

	for i := 0; i < len(nums); i++ {
		diffs := make(map[int]bool)

		for j := i + 1; j < len(nums); j++ {
			// have we seen (sum - nums[i] - nums[j]) before?
			remainder := sum - nums[i] - nums[j]

			for k := range diffs {
				if remainder == k {
					return nums[i], nums[j], remainder
				}
			}

			// if not, mark nums[j] as visited
			diffs[nums[j]] = true
		}
	}

	return 0, 0, 0
}

func main() {
	var a, b, c int

	a, b = part1()
	log.Printf("%d + %d == 2020, their product is %d", a, b, a*b)

	a, b, c = part2()
	log.Printf("%d + %d + %d == 2020, their product is %d", a, b, c, a*b*c)

}
