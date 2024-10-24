package main

import (
	"fmt"
	"os"

	"github.com/vyevs/vtools"
)

type line struct {
	x0, y0, x1, y1 int
}

func main() {
	inputLines, err := vtools.ReadLines(os.Args[1])
	if err != nil {
		panic(err)
	}

	lines := parseInput(inputLines)

	overlaps := countOverlaps(lines, false)
	fmt.Printf("part 1: %d points are crossed multiple times\n", overlaps)

	overlaps2 := countOverlaps(lines, true)
	fmt.Printf("part 2: %d points are crossed multiple times\n", overlaps2)
}

func countOverlaps(lines []line, part2 bool) int {
	walked := make(map[[2]int]int, 1024)
	for _, it := range lines {

		dx, dy := 1, 1
		if it.x0 == it.x1 {
			dx = 0
		} else if it.y0 == it.y1 {
			dy = 0
		} else if !part2 {
			// This line is a diagonal and we're solving part 1, then skip this line.
			continue
		}

		// We must move right to left.
		if it.x1 < it.x0 {
			dx = -1
		}
		// We must move up to down.
		if it.y1 < it.y0 {
			dy = -1
		}

		for x, y := it.x0, it.y0; ; x, y = x+dx, y+dy {
			walked[[2]int{x, y}]++

			// If we've reached the end of the line in either dimension and we are stepping in that dimension then the loop ends.
			if x == it.x1 && dx != 0 {
				break
			}
			if y == it.y1 && dy != 0 {
				break
			}
		}
	}

	var ct int
	for _, v := range walked {
		if v > 1 {
			ct++
		}
	}
	return ct
}

func parseInput(lines []string) []line {
	out := make([]line, 0, len(lines))

	for _, it := range lines {
		var l line
		_, _ = fmt.Sscanf(it, "%d,%d -> %d,%d", &l.x0, &l.y0, &l.x1, &l.y1)
		out = append(out, l)
	}

	return out
}
