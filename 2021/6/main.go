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

	in, err := os.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}

	fish := parseFish(in)

	initialCount := countFish(fish)
	simFish(initialCount, 18)
	simFish(initialCount, 80)
	simFish(initialCount, 256)
	simFish(initialCount, 512)
	simFish(initialCount, 1024)
}

func simFish(count []uint64, steps int) {
	count = slices.Clone(count) // So we don't overwrite the initial state.
	for range steps {
		step(count)
	}

	fmt.Printf("there are %d fish after %d steps\n", vtools.SumSlice(count), steps)
}

func step(count []uint64) {
	zeroCt := count[0]

	for i := 1; i < 9; i++ {
		count[i-1] = count[i]
	}
	count[6] += zeroCt
	count[8] = zeroCt

}

func countFish(fish []int) []uint64 {
	ct := make([]uint64, 9)
	for _, f := range fish {
		ct[f]++
	}
	return ct
}

func parseFish(in []byte) []int {
	str := string(in)
	parts := strings.Split(str, ",")
	out := make([]int, 0, 128)

	for _, part := range parts {
		out = append(out, int(part[0]-'0'))
	}

	return out
}
