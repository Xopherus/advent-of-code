// https://adventofcode.com/2020/day/6
package main

import (
	"bufio"
	"bytes"
	"log"
	"os"
	"runtime"
	"strings"
)

// scanBlock scans a block of text until it hits a double newline.
func scanBlock(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}

	// running this on windows so need to use carriage return...
	sep := []byte("\n\n")
	if runtime.GOOS == "windows" {
		sep = []byte("\r\n\r\n")
	}

	if i := bytes.Index(data, sep); i >= 0 {
		return i + len(sep), data[0:i], nil
	}

	if atEOF {
		return len(data), data, nil
	}

	return
}

func charMap(s string) map[rune]int {
	chars := make(map[rune]int)

	for _, ch := range s {
		// make sure to ignore non-letter characters
		if ch >= 97 && ch <= 132 {
			chars[ch]++
		}
	}

	return chars
}

func part1() int {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("could not open file: %s", err)
	}
	defer f.Close()

	var totalQuestions int

	scanner := bufio.NewScanner(f)
	scanner.Split(scanBlock)

	for scanner.Scan() {
		block := scanner.Text()
		totalQuestions += len(charMap(block))
	}

	return totalQuestions
}

func part2() int {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("could not open file: %s", err)
	}
	defer f.Close()

	var totalQuestions int

	scanner := bufio.NewScanner(f)
	scanner.Split(scanBlock)

	for scanner.Scan() {
		block := scanner.Text()

		// people == number of newlines + 1 (scanner strips off last one)
		people, answers := strings.Count(block, "\n")+1, charMap(block)

		questions := 0
		for _, count := range answers {
			if count == people {
				questions++
			}
		}
		totalQuestions += questions
	}

	return totalQuestions
}

func main() {
	totalQuestions := part1()
	log.Printf("part 1: %d questions", totalQuestions)

	totalQuestions = part2()
	log.Printf("part 2: %d questions", totalQuestions)
}
