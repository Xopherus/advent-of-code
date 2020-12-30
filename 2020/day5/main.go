// https://adventofcode.com/2020/day/5
package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"sort"
)

func getRow(rowStr string) int {
	min, max := float64(0), float64(127)

	for _, ch := range rowStr {
		switch ch {
		case 'F':
			max -= math.Ceil((max - min) / 2)
		case 'B':
			min += math.Ceil((max - min) / 2)
		}
	}

	return int(max)
}

func getCol(colStr string) int {
	min, max := float64(0), float64(7)

	for _, ch := range colStr {
		switch ch {
		case 'L':
			max -= math.Ceil((max - min) / 2)
		case 'R':
			min += math.Ceil((max - min) / 2)
		}
	}

	return int(max)
}

func part1(filename string) int {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("could not open file: %s", err)
	}
	defer f.Close()

	var maxSeatID int

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		seat := scanner.Text()

		row, col := getRow(seat[:7]), getCol(seat[7:])
		seatID := (row * 8) + col

		if seatID > maxSeatID {
			maxSeatID = seatID
		}
	}

	return maxSeatID
}

// this problem reads like a similar binary search problem.
// find an empty seat which has adjacent seats filled (seat IDs +- 1)
func part2(filename string) int {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("could not open file: %s", err)
	}
	defer f.Close()

	var seats []int

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		seat := scanner.Text()
		row, col := getRow(seat[:7]), getCol(seat[7:])

		seats = append(seats, (row*8)+col)
	}

	// sort seats by ID (asc)
	sort.Ints(seats)

	// seats should be monotonically increasing (each seat ID is +1 of the previous).
	// so if we there a gap in seatIDs > 2, we know there should be 2 seats there but one is missing (our seat)
	//
	// also, i'm skipping the first and last seats because the prompt says our seat is not there
	for i := 1; i < len(seats)-1; i++ {
		prev, next := seats[i-1], seats[i+1]

		if next-prev >= 2 {
			return next - 1
		}
	}

	return 0
}

func main() {
	maxSeatID := part1("input.txt")
	log.Printf("highest seatID = %d", maxSeatID)

	mySeatID := part2("input.txt")
	log.Printf("my seatID = %d", mySeatID)
}
