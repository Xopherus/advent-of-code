package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
)

type Target struct {
	x, y float64
}

type Coordinate struct {
	t          *Target
	d2ct, d2at float64
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Panicf("could not open file %s", err)
	}

	maxX, maxY := float64(0), float64(0)
	targets := []*Target{}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		t := Target{}

		if _, err := fmt.Sscanf(scanner.Text(), "%f, %f", &t.x, &t.y); err != nil {
			log.Panicf("could not read coordinates %s", err)
		}
		targets = append(targets, &t)

		if t.x > maxX {
			maxX = t.x
		}
		if t.y > maxY {
			maxY = t.y
		}
	}

	area := make(map[*Target]int)
	infinite := make(map[*Target]bool)
	region := 0

	for y := float64(0); y < maxY+1; y++ {
		for x := float64(0); x < maxX+1; x++ {
			coord := Coordinate{}

			// find closest target to coordinate
			for _, t := range targets {
				dist := math.Abs(t.x-x) + math.Abs(t.y-y)

				if coord.d2ct > dist || (coord.d2ct == 0 && coord.t == nil) {
					coord.d2ct, coord.t = dist, t
				} else if coord.d2ct == dist {
					coord.t = nil
				}

				coord.d2at += dist
			}

			// if no closest target, skip
			if coord.t != nil {
				area[coord.t]++
			}

			// if on edge of map, mark that target as having infinite area
			if x == 0 || x == maxX || y == 0 || y == maxY {
				infinite[coord.t] = true
			}

			if coord.d2at < 10000 {
				region++
			}
		}
	}

	maxArea := 0

	for k, v := range area {
		if _, ok := infinite[k]; v > maxArea && !ok {
			maxArea = v
		}
	}

	fmt.Printf("part1: maxArea = %d\n", maxArea)
	fmt.Printf("part2: region = %d\n", region)

}
