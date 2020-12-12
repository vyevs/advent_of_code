package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const windowSize = 25

func main() {
	input := readInput()

	fmt.Println("Part 1:")
	doPart1(input)
	fmt.Println()

	fmt.Println("Part 2:")
	doPart2(input)
}

func doPart1(ints []int) {
	invalidNumber := findInvalidNumber(ints)

	fmt.Printf("There are no 2 nums in the previous %d numbers that sum to %d\n", windowSize, invalidNumber)
}

func findInvalidNumber(ints []int) int {
	windowStart := 0
	windowEnd := windowSize
	for i := windowSize; i < len(ints); i++ {
		target := ints[i]

		window := ints[windowStart:windowEnd]

		if !areThere2IntsThatSumToTarget(window, target) {
			return target
		}

		windowStart++
		windowEnd++
	}

	panic("There is no invalid number")
}

func areThere2IntsThatSumToTarget(ints []int, target int) bool {
	l := len(ints)
	for i := 0; i < l; i++ {
		for j := 0; j < l; j++ {
			if i != j && ints[i] != ints[j] && ints[i]+ints[j] == target {
				return true
			}
		}
	}
	return false
}

func doPart2(ints []int) {
	invalidNumber := findInvalidNumber(ints)

	sumNums := findContiguousSum(ints, invalidNumber)

	min := intMin(sumNums)
	max := intMax(sumNums)

	fmt.Printf("Encryption weakness is %d\n", min+max)
}

func findContiguousSum(ints []int, target int) []int {
	windowStart, windowEnd := 0, 1
	windowSum := ints[0]

	for {
		if windowEnd > len(ints) {
			break
		}

		if windowSum == target {
			ret := make([]int, windowEnd-windowStart)
			copy(ret, ints[windowStart:windowEnd])
			return ret
		}

		if windowSum < target {
			windowSum += ints[windowEnd]
			windowEnd++
		} else {
			windowSum -= ints[windowStart]
			windowStart++
		}
	}

	panic(fmt.Sprintf("There is no contiguous window that sums to %d", target))
}

func intMin(ints []int) int {
	min := ints[0]
	for i := 0; i < len(ints); i++ {
		if ints[i] < min {
			min = ints[i]
		}
	}
	return min
}

func intMax(ints []int) int {
	max := ints[0]
	for i := 0; i < len(ints); i++ {
		if ints[i] > max {
			max = ints[i]
		}
	}
	return max
}

func readInput() []int {
	scanner := bufio.NewScanner(os.Stdin)

	ints := make([]int, 0, 1024)

	for scanner.Scan() {
		line := scanner.Text()

		v, _ := strconv.Atoi(line)

		ints = append(ints, v)
	}

	return ints
}
