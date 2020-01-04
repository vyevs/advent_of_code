package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	f, _ := os.Open(os.Args[1])

	scanner := bufio.NewScanner(f)

	lines := make([]string, 0, 64)
	for scanner.Scan() {
		line := scanner.Text()

		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("scanner: %v", err)
	}

	p1Code := part1Code(lines)

	fmt.Printf("part1 answer: %q\n", p1Code)

	p2Code := part2Code(lines)

	fmt.Printf("part2 answer: %q\n", p2Code)
}

func part1Code(lines []string) string {
	x, y := 1, 1

	code := make([]rune, 0, len(lines))

	for _, line := range lines {
		for _, r := range line {
			switch r {
			case 'U':
				if y != 0 {
					y--
				}
			case 'D':
				if y != 2 {
					y++
				}
			case 'L':
				if x != 0 {
					x--
				}
			case 'R':
				if x != 2 {
					x++
				}
			}
		}

		code = append(code, rune('0'+y*3+x+1))
	}

	return string(code)
}

func part2Code(lines []string) string {
	x, y := 0, 2

	code := make([]rune, 0, len(lines))

	for _, line := range lines {
		for _, r := range line {
			switch r {
			case 'U':
				y--
			case 'D':
				y++
			case 'L':
				x--
			case 'R':
				x++
			}

			if _, ok := coordToChar[coord{x, y}]; !ok {
				switch r {
				case 'U':
					y++
				case 'D':
					y--
				case 'L':
					x++
				case 'R':
					x--
				}
			}
		}

		char := coordToChar[coord{x, y}]
		if char == 0 {
			log.Fatalf("%d %d", x, y)
		}

		code = append(code, char)
	}

	return string(code)
}

type coord struct {
	x, y int
}

var coordToChar = map[coord]rune{
	coord{2, 0}: '1',
	coord{1, 1}: '2',
	coord{2, 1}: '3',
	coord{3, 1}: '4',
	coord{0, 2}: '5',
	coord{1, 2}: '6',
	coord{2, 2}: '7',
	coord{3, 2}: '8',
	coord{4, 2}: '9',
	coord{1, 3}: 'A',
	coord{2, 3}: 'B',
	coord{3, 3}: 'C',
	coord{2, 4}: 'D',
}
