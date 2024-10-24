package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"time"

	"github.com/vyevs/vtools"
)

func main() {
	defer vtools.TimeIt(time.Now(), "everything")

	weights := readWeights()

	const nBuckets = 4

	// The input comes sorted. Reverse it so it's in descending order. We want to consider heavy weights first.
	slices.Reverse(weights)

	fmt.Printf("weights = %v\n", weights)

	s := newSolver(weights, nBuckets)
	solution := s.solve()

	fmt.Printf("the best package grouping is %v\n", solution)
	fmt.Printf("the quantum entanglement of the 1st group is %d\n", product(solution[0]))
}

func smallestSubsetThatSumsTo(items []int, target int) []int {
	subset := make([]int, 0, len(items))
	var sum int
	bestLen := len(items)

	var rec func(itemsLeft []int)
	rec = func(itemsLeft []int) {
		if len(itemsLeft) == 0 {
			return
		}

		for i, v := range items {
			subset = append(subset, v)
			sum += v
			if sum == target {
				return
			}

			rec(itemsLeft[i:])

			sum -= v
		}

	}

	best := rec(items)
}

func newSolver(weights []int, nBuckets int) *solver {
	weightSum := vtools.SumSlice(weights)

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
	lines, err := vtools.ReadLines(os.Args[1])
	if err != nil {
		log.Fatalf("ReadLines: %v", err)
	}

	return vtools.Map(lines, func(s string) int {
		v, err := strconv.Atoi(s)
		if err != nil {
			panic(fmt.Sprintf("atoi: %v", err))
		}
		return v
	})
}

func product(s []int) int {
	p := 1
	for _, v := range s {
		p *= v
	}
	return p
}
