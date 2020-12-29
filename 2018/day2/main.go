package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/arbovm/levenshtein"
)

// https://adventofcode.com/2018/day2
func main() {
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Panicf("could not open file %s", err)
	}

	part1(input)
	part2(input)
}

func part1(input []byte) {
	scanner := bufio.NewScanner(bytes.NewReader(input))
	wordsWith2Letters, wordsWith3Letters := 0, 0

	for scanner.Scan() {
		// count letters (runes) in ID
		runes := make(map[rune]int)
		for _, r := range scanner.Text() {
			runes[r] += 1
		}

		// check for letters which appear exactly 2 or 3 times
		hasTwo, hasThree := false, false

		for _, count := range runes {
			if count == 2 {
				hasTwo = true
			}
			if count == 3 {
				hasThree = true
			}
		}

		if hasTwo {
			wordsWith2Letters += 1
		}
		if hasThree {
			wordsWith3Letters += 1
		}
	}

	// multiply those together to get checksum
	fmt.Printf("Checksum is %d\n", wordsWith2Letters*wordsWith3Letters)
}

func part2(input []byte) {
	scanner := bufio.NewScanner(bytes.NewReader(input))
	var ids []string

	for scanner.Scan() {
		ids = append(ids, string(scanner.Text()))
	}

	for _, s1 := range ids {
		for _, s2 := range ids {
			if s1 != s2 {
				if dist := levenshtein.Distance(s1, s2); dist == 1 {
					fmt.Printf("Common letters are %s\n", intersection(s1, s2))
					return
				}
			}
		}
	}
}

func intersection(s1, s2 string) string {
	var b strings.Builder

	for i := range s1 {
		if s1[i] == s2[i] {
			b.WriteByte(s1[i])
		}
	}
	return b.String()
}
