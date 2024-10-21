package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/vyevs/vtools"
)

func main() {
	example := []int{0, 2, 7, 0}
	{
		fmt.Printf("example banks required %d redistributions\n", stepsToDuplicate(example))
	}
	input := []int{0, 5, 10, 0, 11, 14, 13, 4, 11, 8, 8, 7, 1, 4, 12, 11}
	{
		fmt.Printf("input banks required %d redistributions\n", stepsToDuplicate(input))
	}

	{
		fmt.Printf("example banks infinite loop is %d steps long\n", stepsToDuplicate(example))
	}
	{
		fmt.Printf("example banks infinite loop is %d steps long\n", stepsToDuplicate(input))
	}

}

func stepsToDuplicate(banks []int) int {
	seen := vtools.NewSet[string](16)

	var steps int
	for ; ; steps++ {
		banksStr := toStr(banks)
		if seen.Contains(banksStr) {
			return steps
		}
		seen.Add(banksStr)
		redistribute(banks)
	}
}

func redistribute(banks []int) {
	max, maxI := vtools.MaxIndex(banks)

	banks[maxI] = 0

	for i := (maxI + 1) % len(banks); max > 0; max, i = max-1, (i+1)%len(banks) {
		banks[i]++
	}
}

func toStr(banks []int) string {
	var b strings.Builder
	b.Grow(32)
	for _, v := range banks {
		_, _ = b.WriteString(strconv.Itoa(v))
	}
	return b.String()
}
