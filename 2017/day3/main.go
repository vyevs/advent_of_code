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
	odd := 1

	for {
		if at == sq {
			return abs(x) + abs(y)
		}

		stepsToNextCorner := odd*odd - (at - 1)
		stepsPerSide := stepsToNextCorner / 4

		// The -1 is because moving up always has 1 fewer steps because we step into a square 1 above the square's bottom right corner.
		for range stepsPerSide - 1 {
			if at == sq {
				return abs(x) + abs(y)
			}

			y++ // up
			at++
		}
		for range stepsPerSide {
			if at == sq {
				return abs(x) + abs(y)
			}
			x-- // left
			at++
		}
		for range stepsPerSide {
			if at == sq {
				return abs(x) + abs(y)
			}

			y-- // down
			at++
		}

		// The + 1 is so we step into the next bigger square, e.g. 1 -> 2, 9 -> 10, 25 -> 26, 49 -> 50 ...
		for range stepsPerSide + 1 {
			if at == sq {
				return abs(x) + abs(y)
			}
			x++ // right
			at++
		}

		odd += 2
	}

	panic("AHHHH")
}

func adjacentValues(target int) int {
	at := 1
	x, y := 0, 0
	odd := 1

	seen := make(map[[2]int]int, 1<<16)
	seen[[2]int{x, y}] = at

	sumOfAdjacents := func(x, y int) int {
		var sum int
		for dx := -1; dx <= 1; dx++ {
			for dy := -1; dy <= 1; dy++ {
				if dx == 0 && dy == 0 {
					continue
				}

				sum += seen[[2]int{x + dx, y + dy}]
			}
		}

		return sum
	}

	for {
		stepsToNextCorner := odd*odd - (at - 1)
		stepsPerSide := stepsToNextCorner / 4

		// The -1 is because moving up always has 1 fewer steps. This is because we step into a square 1 above the square's bottom right corner.
		for range stepsPerSide - 1 {
			y++ // up
			at++

			s := sumOfAdjacents(x, y)
			if s > target {
				return s
			}
			seen[[2]int{x, y}] = s
		}
		for range stepsPerSide {
			x-- // left
			at++

			s := sumOfAdjacents(x, y)
			if s > target {
				return s
			}
			seen[[2]int{x, y}] = s
		}
		for range stepsPerSide {
			y-- // down
			at++

			s := sumOfAdjacents(x, y)
			if s > target {
				return s
			}
			seen[[2]int{x, y}] = s
		}

		// The + 1 is so we step into the next bigger square, e.g. moves between squares 1 -> 2, 9 -> 10, 25 -> 26, 49 -> 50 ...
		for range stepsPerSide + 1 {
			x++ // right
			at++

			s := sumOfAdjacents(x, y)
			if s > target {
				return s
			}
			seen[[2]int{x, y}] = s
		}

		odd += 2
	}

	panic("EVEN MORE AHHHH")
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
