package main

import (
	"fmt"
	"os"
)

func main() {
	digits, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Printf("ReadFile failed: %v", err)
		return
	}

	{
		sum := solveCaptcha(digits, func(i, len int) int {
			return (i + 1) % len
		})
		fmt.Printf("part 1: %d\n", sum)
	}

	{
		sum := solveCaptcha(digits, func(i, len int) int {
			return (i + len/2) % len
		})
		fmt.Printf("part 2: %d\n", sum)
	}
}

func solveCaptcha(digits []byte, mapIdx func(i, len int) int) int {
	var sum int

	dLen := len(digits)

	for i := range dLen {
		nextIdx := mapIdx(i, dLen)

		if digits[i] == digits[nextIdx] {
			sum += int(digits[i] - '0')
		}
	}

	return sum
}
