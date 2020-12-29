package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
)

type claim struct {
	ID            string
	X, Y          int
	Width, Height int
}

// https://adventofcode.com/2018/day3
func main() {
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Panicf("could not open file %s", err)
	}

	// part 1
	fabric := make([][]int, 1000)
	for i := range fabric {
		fabric[i] = make([]int, 1000)
	}
	allClaims := []*claim{}

	scanner := bufio.NewScanner(bytes.NewBuffer(input))

	for scanner.Scan() {
		c := claim{}

		_, err := fmt.Sscanf(scanner.Text(), "#%s @ %d,%d: %dx%d", &c.ID, &c.X, &c.Y, &c.Width, &c.Height)
		if err != nil {
			log.Panicf("could not scan line %s", err)
		}
		allClaims = append(allClaims, &c)

		// while we're in here let's add the claim to the fabric array
		for i := c.X; i < c.X+c.Width; i++ {
			for j := c.Y; j < c.Y+c.Height; j++ {
				fabric[i][j]++
			}
		}
	}

	collisions := 0

	for i := range fabric {
		for j := range fabric[i] {
			if fabric[i][j] > 1 {
				collisions++
			}
		}
	}
	fmt.Printf("Collisions: %d\n", collisions)

	// part 2
claimLoop:
	for _, c := range allClaims {
		collision := false

		for i := c.X; i < c.X+c.Width; i++ {
			for j := c.Y; j < c.Y+c.Height; j++ {
				if fabric[i][j] > 1 {
					collision = true
					continue claimLoop
				}
			}
		}

		if !collision {
			fmt.Printf("Claim with no overlaps: %s\n", c.ID)
			return
		}
	}
}
