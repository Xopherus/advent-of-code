// https://adventofcode.com/2020/day/3
package main

import (
	"bufio"
	"log"
	"os"
)

const (
	treeRune = '#'
)

func traverse(stepRight, stepDown int) int {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("could not open file: %s", err)
	}

	grid := [][]rune{}
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		grid = append(grid, []rune(scanner.Text()))
	}

	var treesEncountered int
	var i, j int

	for {
		if grid[i][j] == treeRune {
			treesEncountered++
		}

		// we've hit the bottom of the map
		if i+stepDown >= len(grid) {
			return treesEncountered
		}

		i += stepDown
		j = (j + stepRight) % len(grid[i])
	}
}

func main() {
	// part 1
	treesEncountered := traverse(3, 1)
	log.Printf("Encountered %d trees", treesEncountered)

	// part 2
	traversals := []struct {
		right, down int
	}{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}

	// make sure this is set to 1, not 0!!
	product := 1

	for _, t := range traversals {
		product *= traverse(t.right, t.down)
	}

	log.Printf("Product of all traversals is %d", product)
}
