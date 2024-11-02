package main

import (
	"fmt"
	"os"
	"time"

	"github.com/vyevs/vtools"
)

func main() {
	defer vtools.TimeIt(time.Now(), "everything")

	lines, err := vtools.ReadLines(os.Args[1])
	if err != nil {
		panic(err)
	}

	rects := vtools.MapSlice(lines, func(s string) rect {
		var r rect
		_, err = fmt.Sscanf(s, "#%d @ %d,%d: %dx%d", &r.id, &r.x0, &r.y0, &r.width, &r.height)
		if err != nil {
			panic(err)
		}
		return r
	})

	overlappingArea := calculateOverlappingArea(rects)
	fmt.Printf("part 1: %d square inches of fabric are within at least 2 claims\n", overlappingArea)

	soloClaimID := idOfClaimThatDoesntOverlapWithAnyOther(rects)
	fmt.Printf("part 2: the ID of the claim that doesn't overlap with any others is %d\n", soloClaimID)
}

func idOfClaimThatDoesntOverlapWithAnyOther(rects []rect) int {
	cts := countVisited(rects)

outer:
	for _, r := range rects {
		for x := r.x0; x < r.x0+r.width; x++ {
			for y := r.y0; y < r.y0+r.height; y++ {
				if cts[x][y] > 1 {
					continue outer
				}
			}
		}
		return r.id
	}

	panic("there is no claim that doesn't overlap with any other")
}

func calculateOverlappingArea(rects []rect) int {
	cts := countVisited(rects)

	var total int
	for _, row := range cts {
		for _, ct := range row {
			if ct >= 2 {
				total++
			}
		}
	}
	return total
}

func countVisited(rects []rect) [][]int {
	cts := make([][]int, 1000)
	for i := range 1000 {
		cts[i] = vtools.NewSliceValues(1000, 0)
	}

	for _, r := range rects {
		for x := r.x0; x < r.x0+r.width; x++ {
			for y := r.y0; y < r.y0+r.height; y++ {
				cts[x][y]++
			}
		}
	}
	return cts
}

type rect struct {
	id, x0, y0, width, height int
}
