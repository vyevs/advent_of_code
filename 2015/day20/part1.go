package main

import (
	"fmt"
	"math"
)

func main() {
	for i := 1; ; i++ {
		presents := presentsToHouse(i)
		if presents >= 34000000 {
			fmt.Printf("house %d got %d presents\n", i, presents)
			break
		}
	}
}

func presentsToHouse(n int) int {
	ds := divisors(n)

	var presents int
	for _, ds := range ds {
		presents += ds * 10
	}

	return presents
}

func divisors(n int) []int {
	divisorSet := make(map[int]struct{}, 512)

	high := int(math.Sqrt(float64(n)))

	for i := 2; i <= high; i++ {
		if n%i == 0 {
			divisorSet[i] = struct{}{}
			divisorSet[n/i] = struct{}{}
		}
	}

	divisors := make([]int, 0, 1024)
	divisors = append(divisors, 1)
	for divisor := range divisorSet {
		divisors = append(divisors, divisor)
	}
	divisors = append(divisors, n)

	return divisors
}
