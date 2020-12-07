package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	input := readInput()

	fmt.Println("Part 1:")
	doPart1(input)
	fmt.Println()

	fmt.Println("Part 2:")
	doPart2(input)
}

func doPart1(field [][]byte) {
	dx, dy := 3, 1

	nTreesHit := howManyTreesHit(field, dx, dy)

	fmt.Printf("We hit %d trees on the way down\n", nTreesHit)
}

func doPart2(field [][]byte) {
	slopes := [5][]int{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}

	var nTreesHit [5]int

	for i, slope := range slopes {
		nTreesHit[i] = howManyTreesHit(field, slope[0], slope[1])
	}

	for i, slope := range slopes {
		fmt.Printf("we hit %d trees going down slope [%d, %d]\n", nTreesHit[i], slope[0], slope[1])
	}

	fmt.Printf("%d * %d * %d * %d * %d = %d\n", nTreesHit[0], nTreesHit[1], nTreesHit[2], nTreesHit[3], nTreesHit[4], product(nTreesHit))
}

func product(ns [5]int) int {
	prod := 1
	for _, v := range ns {
		prod *= v
	}
	return prod
}

func howManyTreesHit(field [][]byte, dx, dy int) int {
	x, y := 0, 0

	var nTreesHit int
	for {
		if field[y][x] == '#' {
			nTreesHit++
		}

		x += dx
		if x >= len(field[0]) {
			x %= len(field[0])
		}

		y += dy
		if y >= len(field) {
			break
		}
	}

	return nTreesHit
}

func readInput() [][]byte {
	scanner := bufio.NewScanner(os.Stdin)

	input := make([][]byte, 0, 1024)

	for scanner.Scan() {
		line := scanner.Text()

		row := make([]byte, 0, 64)
		for i := 0; i < len(line); i++ {
			row = append(row, line[i])
		}
		input = append(input, row)
	}

	return input
}
