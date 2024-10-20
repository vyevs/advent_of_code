package main

import (
	"fmt"
)

func main() {
	input := 277678
	sqrs := []int{1, 12, 23, 1024, input}

	fmt.Println("part 1:")
	for _, sq := range sqrs {
		dist := distanceTo(sq)

		fmt.Printf("%d steps to square %d\n", dist, sq)
	}

	fmt.Printf("part 2: the first value written greater than %d is %d\n", input, adjacentValues(input))
}

func distanceTo(sq int) int {
	at := 1
	x, y := 0, 0
	sqSideLen := 1

	move := func(dx, dy int) {
		x += dx
		y += dy
		at++
	}

	for at != sq {
		// At the beginning of each iteration, we are on the cell above the bottom right corner of the square we're on.
		// We traverse this square's sides during this iteration.

		// It takes n-1 steps to move across a segment of length n
		stepsPerSide := sqSideLen - 1

		// The -1 is because moving up always has 1 fewer steps because we step into a square 1 above the square's bottom right corner.
		stepsUp := stepsPerSide - 1
		for range stepsUp {
			move(0, 1)

			if at == sq {
				return abs(x) + abs(y)
			}
		}

		stepsLeft := stepsPerSide
		for range stepsLeft {
			move(-1, 0)

			if at == sq {
				return abs(x) + abs(y)
			}
		}

		stepsDown := stepsPerSide
		for range stepsDown {
			move(0, -1)

			if at == sq {
				return abs(x) + abs(y)
			}
		}

		// The + 1 is so we step into the next bigger square, e.g. 1 -> 2, 9 -> 10, 25 -> 26, 49 -> 50 ...
		stepsRight := stepsPerSide + 1
		for range stepsRight {
			move(1, 0)

			if at == sq {
				return abs(x) + abs(y)
			}

		}

		sqSideLen += 2
	}

	return abs(x) + abs(y)
}

func adjacentValues(target int) int {
	x, y := 0, 0
	sqSideLen := 1

	seen := make(map[[2]int]int, 1<<16)
	seen[[2]int{x, y}] = 1

	// sumOfAdjacents returns the sum of all squares adjacent to [x, y]. Some of these squares may not have a value, they don't matter.
	sumOfAdjacents := func(x, y int) int {
		var sum int
		for dx := -1; dx <= 1; dx++ {
			for dy := -1; dy <= 1; dy++ {
				sum += seen[[2]int{x + dx, y + dy}]
			}
		}

		return sum
	}

	move := func(dx, dy int) {
		x += dx
		y += dy
	}

	for {
		stepsPerSide := sqSideLen - 1

		// The -1 is because moving up always has 1 fewer steps. This is because we step into a square 1 above the square's bottom right corner.
		for range stepsPerSide - 1 {
			move(0, 1) // up

			s := sumOfAdjacents(x, y)
			if s > target {
				return s
			}
			seen[[2]int{x, y}] = s
		}
		for range stepsPerSide {
			move(-1, 0) // left

			s := sumOfAdjacents(x, y)
			if s > target {
				return s
			}
			seen[[2]int{x, y}] = s
		}
		for range stepsPerSide {
			move(0, -1) // down

			s := sumOfAdjacents(x, y)
			if s > target {
				return s
			}
			seen[[2]int{x, y}] = s
		}

		// The + 1 is so we step into the next bigger square, e.g. moves between squares 1 -> 2, 9 -> 10, 25 -> 26, 49 -> 50 ...
		for range stepsPerSide + 1 {
			move(1, 0) // right

			s := sumOfAdjacents(x, y)
			if s > target {
				return s
			}
			seen[[2]int{x, y}] = s
		}

		sqSideLen += 2
	}

	panic("EVEN MORE AHHHH")
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
