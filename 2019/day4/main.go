package main

import (
	"fmt"
	"strconv"
)

func main() {
	doPart1()
	doPart2()
}

func doPart1() {
	var nValid int
	for pw := 178416; pw <= 676461; pw++ {
		if part1pwSatisfies(pw) {
			nValid++
		}
	}

	fmt.Printf("%d passwords in the range satisfy the criteria\n", nValid)
}

func doPart2() {
	var nValid int
	for pw := 178416; pw <= 676461; pw++ {
		if part2pwSatisfies(pw) {
			nValid++
		}
	}

	fmt.Printf("%d passwords in the range satisfy the criteria\n", nValid)
}

func part1pwSatisfies(pw int) bool {
	return hasAdjacentSameDigits(pw) && isNonDecreasing(pw)
}

func part2pwSatisfies(pw int) bool {
	return hasAdjacentSameDigitsButNot3InARow(pw) && isNonDecreasing(pw)
}

func hasAdjacentSameDigits(pw int) bool {
	var digit int
	prevDigit := -1
	for pw > 0 {
		digit = pw % 10

		pw /= 10

		if prevDigit != -1 && digit == prevDigit {
			return true
		}

		prevDigit = digit
	}

	return false
}

func hasAdjacentSameDigitsButNot3InARow(pw int) bool {
	pwStr := strconv.Itoa(pw)

	for i := 0; i < len(pwStr)-1; i++ {
		dig1 := pwStr[i]
		dig2 := pwStr[i+1]

		if dig1 != dig2 {
			continue
		}

		if i > 0 {
			prevDig := pwStr[i-1]

			if prevDig == dig1 {
				continue
			}
		}

		if i+2 < len(pwStr) {
			nextDig := pwStr[i+2]

			if nextDig == dig1 {
				continue
			}
		}

		return true
	}

	return false
}

func isNonDecreasing(pw int) bool {
	var digit int
	prevDigit := -1
	for pw > 0 {
		digit = pw % 10

		pw /= 10

		if prevDigit != -1 && prevDigit < digit {
			return false
		}

		prevDigit = digit
	}

	return true
}
