package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"time"
)

func main() {
	defer func(start time.Time) {
		fmt.Printf("that took %v\n", time.Since(start))
	}(time.Now())

	weights := readWeights()

	const nBuckets = 4

	slices.SortFunc(weights, func(a, b int) int {
		if a < b {
			return 1
		}
		return -1
	})

	fmt.Printf("weights = %v\n", weights)

	s := newSolver(weights, nBuckets)
	solution := s.solve()

	fmt.Printf("the best package grouping is %v\n", solution)
	fmt.Printf("the quantum entanglement of the 1st group is %d\n", product(solution[0]))
}

func newSolver(weights []int, nBuckets int) *solver {
	weightSum := sum(weights)

	if weightSum%nBuckets != 0 {
		log.Fatalf("the sum of the weights is not divisible by the number of buckets %d, there is no solution", nBuckets)
	}

	groups := make([][]int, nBuckets)
	for i := range nBuckets {
		groups[i] = make([]int, 0, len(weights))
	}

	s := solver{
		weights:       weights,
		oneThirdTotal: weightSum / nBuckets,
		groups:        groups,
		weightTotals:  make([]int, nBuckets),
		bestLenSoFar:  len(weights),
		bestQESoFar:   10e10,
	}
	for i := range nBuckets {
		s.groups[i] = make([]int, 0, len(weights))
	}

	return &s
}

type solver struct {
	weights       []int
	oneThirdTotal int

	groups       [][]int
	weightTotals []int

	bestLenSoFar int
	bestQESoFar  int
	solution     [][]int
}

func (s *solver) solve() [][]int {
	s.recurse(0)
	return s.solution
}

func (s *solver) recurse(idx int) {
	if idx == len(s.weights) {
		if allEqual(s.weightTotals) {
			solLength := len(s.groups[0])

			if solLength < s.bestLenSoFar {
				s.solution = cloneGroups(s.groups)
				s.bestLenSoFar = solLength

			} else {
				quantumEntanglement := product(s.groups[0])

				if quantumEntanglement < s.bestQESoFar {
					s.solution = cloneGroups(s.groups)
					s.bestQESoFar = quantumEntanglement
				}

			}
		}
		return
	}

	if len(s.groups[0]) > s.bestLenSoFar {
		return
	}

	weight := s.weights[idx]

	for i := range s.groups {
		s.groups[i] = append(s.groups[i], weight)
		s.weightTotals[i] += weight

		if s.weightTotals[i] <= s.oneThirdTotal {
			s.recurse(idx + 1)
		}

		s.groups[i] = s.groups[i][:len(s.groups[i])-1]
		s.weightTotals[i] -= weight
	}
}

func cloneGroups(gs [][]int) [][]int {
	out := make([][]int, len(gs))
	for i, g := range gs {
		out[i] = slices.Clone(g)
	}
	return out
}

func allEqual(vs []int) bool {
	for i := range len(vs) - 1 {
		if vs[i] != vs[i+1] {
			return false
		}
	}
	return true
}

func readWeights() []int {
	bs, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("failed to read file: %v", err)
	}

	sc := bufio.NewScanner(bytes.NewReader(bs))

	weights := make([]int, 0, 64)
	for sc.Scan() {
		line := sc.Text()

		weight, _ := strconv.Atoi(line)

		weights = append(weights, weight)
	}

	return weights
}

func sum(s []int) int {
	sm := 0
	for _, v := range s {
		sm += v
	}
	return sm
}

func product(s []int) int {
	p := 1
	for _, v := range s {
		p *= v
	}
	return p
}
