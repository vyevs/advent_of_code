package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/vyevs/vtools"
)

// S = { a0, a1, ... an-1, an }
// find x such that sum abs(x - ai) for all ai is minimized.

func main() {
	crabs := getCrabPositions()

	fuelUsed := getTotalFuelToBestPosition(crabs, false)
	fmt.Printf("part 1: fuel used to best position is %d\n", fuelUsed)

	fuelUsed = getTotalFuelToBestPosition(crabs, true)
	fmt.Printf("part 2: fuel used to best position is %d\n", fuelUsed)

}

func getTotalFuelToBestPosition(crabs []int, part2 bool) int {
	min := slices.Min(crabs)
	max := slices.Max(crabs)

	bestDist := int(10e9)
	for p := min; p <= max; p++ {
		dist := getTotalDistance(crabs, p, part2)
		if dist < bestDist {
			bestDist = dist
		}
	}

	return bestDist
}

func getTotalDistance(crabs []int, to int, part2 bool) int {
	var sum int
	for _, c := range crabs {
		if part2 {
			d := vtools.Abs(c - to)
			sum += (d * (d + 1)) / 2
		} else {
			sum += vtools.Abs(to - c)
		}
	}
	return sum
}

func getCrabPositions() []int {
	fBs, err := os.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}

	parts := strings.Split(string(fBs), ",")
	out := make([]int, 0, len(parts))
	for _, it := range parts {
		pos, err := strconv.Atoi(it)
		if err != nil {
			panic(err)
		}
		out = append(out, pos)
	}

	return out
}
