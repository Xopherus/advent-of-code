package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("input")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	// load maze into memory
	maze := make([]int, 0)
	fscanner := bufio.NewScanner(f)
	for fscanner.Scan() {
		num, _ := strconv.Atoi(fscanner.Text())
		maze = append(maze, num)
	}

	p := 0
	moves := 0

	for {
		if p < 0 || p > len(maze)-1 {
			break
		}
		fmt.Printf("instruction [%d] : %d\n", p, maze[p])
		jump := maze[p]

		maze[p]++ // update current instruction before jumping
		moves++

		p += jump // modify current position based on instruction
	}

	fmt.Printf("number of moves to exit maze: %d\n", moves)
}
