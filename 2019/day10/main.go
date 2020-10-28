package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	grid := make([][]int, 0, 64)
	for scanner.Scan() {
		line := scanner.Text()

		row := make([]int, 0, 64)
		for _, r := range line {
			if r == '#' {
				row = append(row, 1)
			} else {
				row = append(row, 0)
			}
		}

		grid = append(grid, row)
	}

	var bestX, bestY, best int
	for x, row := range grid {
		for y, p := range row {
			if p == 1 {
				seen := visibility(grid, x, y)

				if best == 0 || seen > best {
					bestX = x
					bestY = y
					best = seen
				}
			}
		}
	}

	fmt.Printf("best at [%d, %d] sees %d\n", bestX, bestY, best)
}

type direction struct {
	dx, dy int
}

func visibility(grid [][]int, x, y int) int {
	seenAlready := make(map[direction]struct{}, len(grid)*len(grid[0]))
	for otherX, row := range grid {
		for otherY, p := range row {
			if otherX == x && otherY == y {
				continue
			}

			dir := gcdDirection(x, y, otherX, otherY)

			if p == 1 {
				seenAlready[dir] = struct{}{}
			}

		}
	}

	return len(seenAlready)
}

func gcdDirection(x, y, otherX, otherY int) direction {
	dx := otherX - x
	dy := otherY - y

	if dx == 0 {
		// if dx is 0, we want to reduce to direction [0,1] or [0,-1]
		if dy < 0 {
			dy /= -dy
		} else {
			dy /= dy
		}
		return direction{
			dx: 0,
			dy: dy,
		}
	} else if dy == 0 {
		// if dy is 0, we want to reduce to direction [1,0] or [-1,0]
		if dx < 0 {
			dx /= -dx
		} else {
			dx /= dx
		}
		return direction{
			dx: dx,
			dy: 0,
		}
	}

	g := gcd(intAbs(dx), intAbs(dy))

	// reduce the dx & dy to the lowest possible representation that maintains the same slope
	dx = dx / g
	dy = dy / g

	return direction{
		dx: dx,
		dy: dy,
	}
}

// returns gcd of 2 non-zero, positive ints
func gcd(a, b int) int {
	if a <= 0 {
		panic("gcd a < 0")
	}
	if b <= 0 {
		panic("gcd b < 0")
	}

	if a == b {
		return a
	}

	if a > b {
		return gcd(a-b, b)
	}
	return gcd(a, b-a)
}

func intAbs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}
