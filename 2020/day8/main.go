// https://adventofcode.com/2020/day/8
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type instruction struct {
	op  string
	arg int

	// used to uniquely identify instruction
	number int
}

func (i *instruction) Unmarshal(data []byte) error {
	parts := strings.Split(string(data), " ")
	if len(parts) != 2 {
		return fmt.Errorf("invalid instruction, must be of the format op: argument")
	}

	arg, err := strconv.Atoi(parts[1])
	if err != nil {
		return err
	}

	i.op, i.arg = parts[0], arg

	return nil
}

func eval(instructions []instruction) (int, bool) {
	var next int
	var accumulator int

	// track visited instructions for infinite loop protection
	visited := make(map[int]bool)

	for {
		// just for sanity checking...
		if next < 0 {
			log.Fatalf("invalid instruction, array index out of bounds")
		}

		// program successfully terminates
		if next >= len(instructions) {
			return accumulator, true
		}

		// program detected infinite loop
		if _, ok := visited[next]; ok {
			// log.Printf("detected infinite loop at instruction %d, exiting...", next)
			return accumulator, false
		}

		visited[next] = true

		// log.Printf("at instruction %d: %s: %d", next, instructions[next].op, instructions[next].arg)

		switch inst := instructions[next]; inst.op {
		case "acc":
			accumulator += inst.arg
			next++
		case "jmp":
			next += inst.arg
		case "nop":
			next++
		}
	}
}

func part1() int {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("could not open file: %s", err)
	}
	defer f.Close()

	var instructions []instruction
	var instructionCounter int

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		i := instruction{number: instructionCounter}
		i.Unmarshal(scanner.Bytes())

		instructions = append(instructions, i)
		instructionCounter++
	}

	accumulator, _ := eval(instructions)
	return accumulator
}

func part2() int {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("could not open file: %s", err)
	}
	defer f.Close()

	var instructions []instruction
	var instructionCounter int

	// track which instructions are potentially corrupted so we can fix them
	var jumpOrNops []int

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		i := instruction{number: instructionCounter}
		i.Unmarshal(scanner.Bytes())

		instructions = append(instructions, i)
		instructionCounter++

		if i.op != "acc" {
			jumpOrNops = append(jumpOrNops, i.number)
		}
	}

	// try changing each jmp/nop to the other one and see if the program terminates
	for _, i := range jumpOrNops {
		prevOp := instructions[i].op

		switch prevOp {
		case "jmp":
			instructions[i].op = "nop"
		case "nop":
			instructions[i].op = "jmp"
		}

		// log.Printf("changed instruction %d (was %s now %+v)", i, prevOp, instructions[i])

		accumulator, success := eval(instructions)
		if success {
			return accumulator
		}

		// don't forget to reset the instruction before changing another one!
		instructions[i].op = prevOp
	}

	return 0
}

func main() {
	accumulator := part1()
	log.Printf("part1: accumulator = %d", accumulator)

	accumulator = part2()
	log.Printf("part2: accumulator = %d", accumulator)
}
