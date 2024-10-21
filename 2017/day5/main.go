package main

import (
	"fmt"
	"slices"
	"strconv"

	"github.com/vyevs/vtools"
)

func main() {
	example := []int{0, 3, 0, 1, -3}
	{
		fmt.Println("part 1 example:")
		fmt.Printf("\twe exit program %v after %d steps\n", example, stepsToExit1(example))
	}

	inputLines, _ := vtools.ReadLines("input.txt")
	input := vtools.Map(inputLines, func(s string) int {
		v, _ := strconv.Atoi(s)
		return v
	})
	fmt.Printf("the input contains %d instructions\n", len(input))

	{
		fmt.Println("part 1:")
		fmt.Printf("\tthe program exits after %d steps\n", stepsToExit1(input))
	}

	{
		fmt.Println("part 2 example:")
		fmt.Printf("\the program exits after %d steps\n", stepsToExit2(example))
	}

	{
		fmt.Println("part 2:")
		fmt.Printf("\tthe program exits after %d steps\n", stepsToExit2(input))
	}
}

func stepsToExit1(instrs []int) int {
	jmps := slices.Clone(instrs)

	var steps, pc int
	for pc < len(jmps) {
		jmp := jmps[pc]
		jmps[pc]++
		pc += jmp

		steps++
	}

	return steps
}

func stepsToExit2(instrs []int) int {
	jmps := slices.Clone(instrs)

	var steps, pc int
	for pc < len(jmps) {
		jmp := jmps[pc]

		if jmp >= 3 {
			jmps[pc]--
		} else {
			jmps[pc]++
		}

		pc += jmp

		steps++
	}

	return steps
}
