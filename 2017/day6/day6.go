package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// returns the index of the highest value in the slice
func findMax(slice []int) (int, int) {
	max := 0
	maxIdx := 0

	for k, v := range slice {
		if v > max {
			maxIdx = k
			max = v
		}
	}
	return maxIdx, max
}

// converts the slice to a csv-string
func saveState(s []int) string {
	delim := ","
	return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(s)), delim), "[]")
}

func main() {
	f, err := os.Open("input")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	// load state of memory banks
	banks := []int{}
	fscanner := bufio.NewScanner(f)
	for fscanner.Scan() {
		for _, b := range strings.Split(fscanner.Text(), "\t") {
			b, _ := strconv.Atoi(b)
			banks = append(banks, b)
		}
	}

	// track previously seen states
	cycles := 0
	states := make(map[string]int)

	for {
		// choose highest memory bank, redistribute amongst other banks
		idx, blocks := findMax(banks)

		for i := 1; i < blocks+1; i++ {
			banks[(idx+i)%len(banks)]++
		}
		banks[idx] = 0 // highest memory bank is now empty

		// save state of the world
		cycles++

		if _, ok := states[saveState(banks)]; !ok {
			states[saveState(banks)] = 1
		} else {
			fmt.Printf("infinite loop detected after %d cycles\n", cycles)
			fmt.Printf("state: %s\n", saveState(banks))
			break // state seen before, exit
		}
	}
}
