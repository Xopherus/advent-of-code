package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
)

// https://adventofcode.com/2018/day1
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
	freq := 0

	for scanner.Scan() {
		delta, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Panicf("could not convert to int %s", err)
		}
		freq += delta
	}
	fmt.Printf("Part1: The resulting frequency is %d\n", freq)
}

func part2(input []byte) {
	freq := 0
	freqCache := make(map[int]struct{})

	for {
		scanner := bufio.NewScanner(bytes.NewReader(input))

		for scanner.Scan() {
			delta, err := strconv.Atoi(scanner.Text())
			if err != nil {
				log.Panicf("could not convert to int %s", err)
			}
			freq += delta

			if _, exists := freqCache[freq]; !exists {
				freqCache[freq] = struct{}{}
				continue
			}
			fmt.Printf("Part2: The first frequency to be seen twice is %d\n", freq)
			return
		}
	}
}
