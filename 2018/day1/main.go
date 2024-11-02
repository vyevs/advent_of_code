package main

import (
	"fmt"
	"os"

	"github.com/vyevs/vtools"
)

func main() {
	lines, err := vtools.ReadLines(os.Args[1])
	if err != nil {
		panic(err)
	}

	moves := vtools.MapSlice(lines, vtools.AtoiOrPanic)
	fmt.Printf("part 1: the final frequency is %d\n", finalFrequency(moves))

	fmt.Printf("part 2: the first frequency reached twice is %d\n", firstFrequencyReachedTwice(moves))
}

func finalFrequency(moves []int) int {
	var f int
	for _, m := range moves {
		f += m
	}
	return f
}

func firstFrequencyReachedTwice(moves []int) int {
	var f int

	seen := vtools.NewSet[int](len(moves))
	seen.Add(f)
	for m := range vtools.Cycle(moves) {
		f += m
		if seen.Contains(f) {
			return f
		}
		seen.Add(f)
	}

	panic("no frequency is reached twice")
}
