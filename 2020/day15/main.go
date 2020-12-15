package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	nums := readNumbers()

	fmt.Println("Part 1:")
	doPart1(nums)
	fmt.Println()

	fmt.Println("Part 2:")
	doPart2(nums)
}

func doPart1(nums []int) {
	numSpoken := numberSpokenOnTurn(nums, 2020)

	fmt.Printf("The 2020th number spoken is %d\n", numSpoken)
}

func doPart2(nums []int) {
	numSpoken := numberSpokenOnTurn(nums, 30000000)

	fmt.Printf("The 30000000th number spoken is %d\n", numSpoken)
}

func numberSpokenOnTurn(nums []int, targetTurn int) int {
	turnsNumWasSpoken := make(map[int][2]int, targetTurn)

	addTurnForNumber := func(num, turn int) {
		turns := turnsNumWasSpoken[num]
		turns[0] = turns[1]
		turns[1] = turn
		turnsNumWasSpoken[num] = turns
	}

	for i, v := range nums {
		addTurnForNumber(v, i+1)
	}

	lastNumSpoken := nums[len(nums)-1]

	for turn := len(nums) + 1; ; turn++ {
		numWasSpoken := turnsNumWasSpoken[lastNumSpoken]

		var thisTurnsNumber int
		if numWasSpoken[0] == 0 {
			thisTurnsNumber = 0
		} else {
			thisTurnsNumber = numWasSpoken[1] - numWasSpoken[0]
		}

		addTurnForNumber(thisTurnsNumber, turn)

		if turn == targetTurn {
			return thisTurnsNumber
		}

		lastNumSpoken = thisTurnsNumber
	}
}

func readNumbers() []int {
	scanner := bufio.NewScanner(os.Stdin)

	nums := make([]int, 0, 64)

	scanner.Scan()
	line := scanner.Text()

	numStrs := strings.Split(line, ",")

	for _, numStr := range numStrs {
		num, _ := strconv.Atoi(numStr)
		nums = append(nums, num)
	}

	return nums
}
