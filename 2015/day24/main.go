package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	weights := readWeights()

	fmt.Printf("weights = %v\n", weights)

	groupings := possibleEqualWeightGroupings(weights)

	fmt.Printf("There are %d equal weight groupings\n", len(groupings))

	sort.Slice(groupings, func(i, j int) bool {
		return len(groupings[i][0]) < len(groupings[j][0])
	})

	firstLen := len(groupings[0][0])
	firstN := 0
	for i := 1; ; i++ {
		firstN++
		if len(groupings[i][0]) > firstLen {
			break
		}
	}

	sort.Slice(groupings[:firstN], func(i, j int) bool {
		return len(groupings[i][0]) < len(groupings[j][0])
	})

	fmt.Printf("Quantum entaglement of first group in ideal configuration is %d\n", product(groupings[0][0]))
}

func product(ns []int) int {
	prod := 1
	for _, v := range ns {
		prod *= v
	}
	return prod
}

func possibleEqualWeightGroupings(weights []int) [][3][]int {
	assignments := make([]int, len(weights))

	var recurse func(weightIdx int)

	groupings := make([][3][]int, 0, 1024)

	recurse = func(weightIdx int) {
		if weightIdx == len(weights) {
			if allSameWeight(weights, assignments) {

				var grouping [3][]int

				for i, assignment := range assignments {
					grouping[assignment-1] = append(grouping[assignment-1], weights[i])
				}

				groupings = append(groupings, grouping)
			}

		} else {

			for i := 1; i <= 3; i++ {
				assignments[weightIdx] = i
				recurse(weightIdx + 1)
			}

		}
	}

	recurse(0)

	return groupings
}

func allSameWeight(weights, assignments []int) bool {
	var sums [3]int

	for i, weight := range weights {
		sums[assignments[i]-1] += weight
	}

	return sums[0] == sums[1] && sums[1] == sums[2]
}

func readWeights() []int {
	scanner := bufio.NewScanner(os.Stdin)

	weights := make([]int, 0, 64)
	for scanner.Scan() {
		line := scanner.Text()

		weight, _ := strconv.Atoi(line)

		weights = append(weights, weight)
	}

	return weights
}
