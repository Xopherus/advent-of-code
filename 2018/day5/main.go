package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"strings"
)

const units = "abcdefghijklmnopqrstuvwxyz"

func main() {
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Panicf("could not open file %s", err)
	}

	// part1
	reaction := func(u1, u2 string) bool {
		return u1 != u2 && strings.EqualFold(u1, u2)
	}
	polymer := collapse(string(input), reaction)

	fmt.Printf("Smallest polymer has %d units\n", len(polymer))

	// part 2
	unit, minLength := "", math.MaxInt64

	for _, u := range strings.Split(units, "") {
		removeUnit := func(r rune) rune {
			if strings.EqualFold(string(r), u) {
				return -1
			}
			return r
		}
		p := collapse(strings.Map(removeUnit, string(input)), reaction)

		if len(p) < minLength {
			unit, minLength = u, len(p)
		}
	}
	fmt.Printf("Smallest polymer can be created by removing `%s`: resulting length of %d\n", unit, minLength)
}

// collapse collapses a polymer
func collapse(polymer string, reactionFunc func(u1, u2 string) bool) string {
	states := make(map[string]int)

	for {
		var intermediate strings.Builder
		i := 0

		for {
			if i >= len(polymer)-1 {
				break
			}
			u1, u2 := string(polymer[i]), string(polymer[i+1])

			if reactionFunc(u1, u2) {
				i += 2
				continue
			}
			intermediate.WriteString(u1)
			i++

			// if u2 is last character, we need to write that out
			if i == len(polymer)-1 {
				intermediate.WriteString(u2)
			}
		}
		polymer = intermediate.String()

		if _, ok := states[polymer]; ok {
			return polymer
		}
		states[polymer]++
	}
}
