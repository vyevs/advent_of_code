package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
	"time"

	"github.com/vyevs/vtools"
)

func main() {
	defer vtools.TimeIt(time.Now(), "everything")

	seqs := getSequences()

	sum := sumNextValues(seqs, false)
	fmt.Printf("part 1: the sum of the extrapolated values is %d\n", sum)

	sum = sumNextValues(seqs, true)
	fmt.Printf("part 2: the sum of the previous values is %d\n", sum)
}

func sumNextValues(seqs [][]int, part2 bool) int {
	var sum int
	for _, s := range seqs {
		if part2 {
			prev := determinePreviousValue(s)
			sum += prev
		} else {
			sum += determineNextValue(s)
		}
	}
	return sum
}

func determineNextValue(seq []int) int {
	seq = slices.Clone(seq) // So we don't modify the original.

	var v int
	for {
		v += seq[len(seq)-1]

		seq = nextSeqInPlace(seq)

		if vtools.AllSlice(seq, 0) {
			break
		}
	}

	return v
}

func determinePreviousValue(seq []int) int {
	firstColumn := make([]int, 0, len(seq))
	for {
		firstColumn = append(firstColumn, seq[0])

		seq = nextSeqInPlace(seq)

		if vtools.AllSlice(seq, 0) {
			break
		}
	}

	var v int
	for _, colVal := range slices.Backward(firstColumn) {
		v = colVal - v
	}

	return v
}

func nextSeqInPlace(seq []int) []int {
	for i := 0; i < len(seq)-1; i++ {
		seq[i] = seq[i+1] - seq[i]
	}
	seq = seq[:len(seq)-1]
	return seq
}

func getSequences() [][]int {
	lines, err := vtools.ReadLines(os.Args[1])
	if err != nil {
		panic(err)
	}

	seqs := make([][]int, 0, len(lines))
	for _, line := range lines {
		parts := strings.Fields(line)

		seq := vtools.MapSlice(parts, vtools.AtoiOrPanic)
		seqs = append(seqs, seq)
	}

	return seqs
}
