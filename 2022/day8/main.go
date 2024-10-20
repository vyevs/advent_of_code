package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	g := getGrid(os.Stdin)
	part1(g)
	part2(g)
}

func part2(g [][]int) {
	bss := bestScenicScore(g)
	
	fmt.Printf("the best scenic score is %d\n", bss)
}

func bestScenicScore(g [][]int) int {
	var best int
	for ri, row := range g {
		for ci := range row {
			ss := scenicScore(g, ri, ci)
			if ss > best {
				best = ss
			}
		}
	}
	return best
}

func scenicScore(g [][]int, r, c int) int {
	h, w := len(g), len(g[0])
	// Trees on the edges have a scenic score of 0 as at least one of their direction scores is 0.
	if r == 0 || r == h-1 || c == 0 || c == w-1 {
		return 0
	}
	
	uss := viewingDistanceInDirection(g, r, c, -1, 0)
	dss := viewingDistanceInDirection(g, r, c, 1, 0)
	lss := viewingDistanceInDirection(g, r, c, 0, -1)
	rss := viewingDistanceInDirection(g, r, c, 0, 1)
	
	return uss * dss * lss * rss
}

func viewingDistanceInDirection(g [][]int, c, r, dr, dc int) int {
	h, w := len(g), len(g[0])
	
	var vis int
	for ri, ci := r+dr, c+dc; ri >= 0 && ri < h && ci >= 0 && ci < w; ri, ci = ri+dr, ci+dc {
		vis++
		if g[ri][ci] >= g[r][c] {
			break
		}
	}
	return vis
}

func part1(g [][]int) {
	ctVisible := countVisibleTrees(g)
	
	fmt.Printf("%d trees are visible\n", ctVisible)
}

// countVisibleTreesNaive returns the number of visible trees in g according to part 1 of the problem.
func countVisibleTrees(g [][]int) int {
	var ct int
	for ri, r := range g {
		for ci := range r {
			vis := isTreeVisible(g, ri, ci)
			if vis {
				ct++
			}
		}
	}
	return ct
}

// isTreeVisible counts the number of trees visible from tree [r][c].
func isTreeVisible(g [][]int, r, c int) bool {
	h, w := len(g), len(g[0])
	ourHt := g[r][c]
		
	// Looking leftwards.
	for ci := w - 1; ci >= c; ci-- {
		if ci == c {
			return true
		}
		if g[r][ci] >= ourHt {
			break
		}
	}
	
	// Looking rightwards.
	for ci := 0; ci <= c; ci++ {
		if ci == c {
			return true
		}
		if g[r][ci] >= ourHt {
			break
		}
	}
	
	// Looking upwards.
	for ri := h - 1; ri >= r; ri-- {
		if ri == r {
			return true
		}
		if g[ri][c] >= ourHt {
			break
		}
	}
	
	// Looking downwards.
	for ri := 0; ri <= r; ri++ {
		if ri == r {
			return true
		}
		if g[ri][c] >= ourHt {
			break
		}
	}
	
	return false
}

func getGrid(r io.Reader) [][]int {
	s := bufio.NewScanner(r)
	
	grid := make([][]int, 0, 128)
	for s.Scan() {
		l := s.Text()
		
		row := make([]int, 0, 128)
		for _, r := range l {
			row = append(row, int(r - '0'))
		}
		grid = append(grid, row)
	}
	
	return grid
}