package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func isValid(passphrase string) bool {
	words := strings.Split(passphrase, " ")
	uniqueWords := make(map[string]int)

	for _, word := range words {
		if _, ok := uniqueWords[word]; ok {
			return false
		}
		uniqueWords[word] = 1
	}
	return true
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	total, valid := 0, 0

	fscanner := bufio.NewScanner(f)
	for fscanner.Scan() {
		total++

		if isValid(fscanner.Text()) {
			valid++
		}
	}

	fmt.Printf("%d / %d valid passphrases\n", valid, total)
}
