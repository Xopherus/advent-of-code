// https://adventofcode.com/2020/day/7
package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type Rule struct {
	Color string
	Bags  map[string]int
}

func parseRule(rule string) Rule {
	parts := strings.Split(rule, "bags contain")
	if len(parts) != 2 {
		log.Fatalf("cannot parse rule, must be of fmt: <color> bags contain <contents>")
	}

	r := Rule{Color: strings.TrimSpace(parts[0])}
	if strings.Contains(parts[1], "no other bags") {
		return r
	}

	r.Bags = make(map[string]int)

	for _, part := range strings.Split(parts[1], ",") {
		inner := strings.Split(strings.TrimSpace(part), " ")

		innerCount, _ := strconv.Atoi(inner[0])
		innerColor := strings.Join(inner[1:len(inner)-1], " ")

		r.Bags[innerColor] = innerCount
	}

	return r
}

func part1(startingColor string) int {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("could not open file: %s", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	// map allows us to lookup a bag and find its parents (double nested map to prevent duplicates)
	bags := make(map[string]map[string]bool)

	for scanner.Scan() {
		rule := parseRule(scanner.Text())

		if _, ok := bags[rule.Color]; !ok {
			bags[rule.Color] = make(map[string]bool)
		}

		for color := range rule.Bags {
			if _, ok := bags[color]; !ok {
				bags[color] = make(map[string]bool)
			}

			bags[color][rule.Color] = true
		}
	}

	colors := []string{startingColor}
	possibleColors := make(map[string]bool)

	for {
		var nextColors []string

		for _, color := range colors {
			// add parents to possibleColors
			for parent := range bags[color] {
				possibleColors[parent] = true
				nextColors = append(nextColors, parent)
			}
		}

		if len(nextColors) == 0 {
			return len(possibleColors)
		}

		colors = nextColors
	}
}

// countBags recursively counts the number of bags contained within the bag.
func countBags(bags map[string]Rule, color string) int {
	var total int

	// just count the single bag
	if len(bags[color].Bags) == 0 {
		// log.Printf("%s has 0 bags in it", color)
		return total
	}

	for color, count := range bags[color].Bags {
		total += count + (count * countBags(bags, color))
	}

	// log.Printf("%s has %d bags in it: %+v", color, total, bags[color].Bags)

	return total
}

func part2(startingColor string) int {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("could not open file: %s", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	// map allow us to quickly lookup rules by starting color
	bags := make(map[string]Rule)

	for scanner.Scan() {
		rule := parseRule(scanner.Text())

		bags[rule.Color] = rule
	}

	return countBags(bags, startingColor)
}

func main() {
	var startingColor string = "shiny gold"

	possibleColors := part1(startingColor)
	log.Printf("part1: %d possible colors", possibleColors)

	requiredBags := part2(startingColor)
	log.Printf("part2: %d required bags", requiredBags)
}
