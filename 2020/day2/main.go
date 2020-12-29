// https://adventofcode.com/2020/day/2
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Password struct {
	Policy struct {
		Char     string
		Min, Max int
	}

	Plaintext string
}

func part1Validator(pwd Password) bool {
	charCount := strings.Count(pwd.Plaintext, pwd.Policy.Char)

	return charCount >= pwd.Policy.Min && charCount <= pwd.Policy.Max
}

func part2Validator(pwd Password) bool {
	pwdRunes := []rune(pwd.Plaintext)
	charRune := []rune(pwd.Policy.Char)[0]

	pos1 := pwdRunes[pwd.Policy.Min-1] == charRune
	pos2 := pwdRunes[pwd.Policy.Max-1] == charRune

	return (pos1 || pos2) && (pos1 != pos2)
}

func checkPasswords(pwdValidator func(Password) bool) (int, int) {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("could not open file: %s", err)
	}

	var total, validPasswords int
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		pwd := Password{}

		n, err := fmt.Sscanf(scanner.Text(), "%d-%d %1s: %s", &pwd.Policy.Min, &pwd.Policy.Max, &pwd.Policy.Char, &pwd.Plaintext)
		if err != nil || n != 4 {
			log.Fatalf("could not read line: %s", err)
		}

		if pwdValidator(pwd) {
			validPasswords++
		}

		total++
	}

	return validPasswords, total
}

func main() {
	var total, validPasswords int

	validPasswords, total = checkPasswords(part1Validator)
	log.Printf("Part1: %d / %d passwords valid", validPasswords, total)

	validPasswords, total = checkPasswords(part2Validator)
	log.Printf("Part2: %d / %d passwords valid", validPasswords, total)
}
