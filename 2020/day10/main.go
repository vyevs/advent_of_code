package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	input := readInput()

	fmt.Println("Part 1:")
	doPart1(input)
	fmt.Println()

	fmt.Println("Part 2:")
	doPart2(input)
}

func doPart1(joltages []int) {
	sort.Ints(joltages)

	var oneDiffs, threeDiffs int

	if joltages[0] == 1 {
		oneDiffs++
	} else if joltages[0] == 3 {
		threeDiffs++
	}

	for i := 0; i < len(joltages)-1; i++ {
		v := joltages[i]
		next := joltages[i+1]

		if next-v == 1 {
			oneDiffs++
		} else if next-v == 3 {
			threeDiffs++
		}
	}

	threeDiffs++

	fmt.Printf("One joltage jumps: %d, three joltage jumps: %d\n", oneDiffs, threeDiffs)
	fmt.Printf("%d * %d = %d\n", oneDiffs, threeDiffs, oneDiffs*threeDiffs)
}

func doPart2(joltages []int) {
	sort.Ints(joltages)

	// we must start with 0 because that is the initial joltage we must jump from
	joltages = append([]int{0}, joltages...)

	// waysToArrangeFrom[i] tells you the ways to arrange the adapters starting from i
	waysToArrangeFrom := make([]int, len(joltages))
	for i := 0; i < len(waysToArrangeFrom); i++ {
		waysToArrangeFrom[i] = -1
	}
	// if you are the last adapter, there is only 1 way to get to the end
	waysToArrangeFrom[len(waysToArrangeFrom)-1] = 1

	var recurse func(i int) int

	recurse = func(i int) int {
		if waysToArrangeFrom[i] != -1 {
			return waysToArrangeFrom[i]
		}

		var ct int

		// the ways to go from joltage i is just the sum of the ways
		// to go from the joltages you can jump to
		for j := i + 1; j < len(joltages); j++ {
			if joltages[j]-joltages[i] > 3 {
				break
			}
			ct += recurse(j)
		}
		waysToArrangeFrom[i] = ct
		return ct
	}

	total := recurse(0)

	fmt.Printf("There are %d ways to arrange the adapters\n", total)

}

func readInput() []int {
	scanner := bufio.NewScanner(os.Stdin)

	ints := make([]int, 0, 1024)

	for scanner.Scan() {
		line := scanner.Text()

		i, _ := strconv.Atoi(line)

		ints = append(ints, i)
	}

	return ints
}
