package main

import (
	"fmt"
	"os"
	"slices"
	"time"

	"github.com/vyevs/vtools"
)

func main() {
	defer vtools.TimeIt(time.Now(), "everything")

	heightMap := getHeightMap()

	srl := sumRiskLevels(heightMap)
	fmt.Printf("part 1: the sum of the risk levels is %d\n", srl)

	product := getProductOf3LargestBasins(heightMap)
	fmt.Printf("part 2: the product of the 3 largest basin sizes is %d\n", product)

}

func getProductOf3LargestBasins(grid [][]uint8) int {
	basins := make([]int, 0, 1024)
	for r, row := range grid {
		for c := range row {
			if isLowPoint(grid, r, c) {
				bs := getBasinSize(grid, r, c)
				basins = append(basins, bs)
			}
		}
	}
	slices.Sort(basins)
	l := len(basins) - 1
	return basins[l] * basins[l-1] * basins[l-2]
}

type basinSizer struct {
	visited [][]bool
}

func getBasinSize(grid [][]uint8, r, c int) int {
	visited := make([][]bool, len(grid))
	for i, row := range grid {
		visited[i] = make([]bool, len(row))
	}

	return basinSizer{visited: visited}.getBasinSizeRec(grid, r, c)
}

func (s basinSizer) getBasinSizeRec(grid [][]uint8, r, c int) int {
	if s.visited[r][c] || grid[r][c] == 9 {
		return 0
	}

	size := 1
	s.visited[r][c] = true
	v := grid[r][c]

	if r > 0 && grid[r-1][c] > v {
		size += s.getBasinSizeRec(grid, r-1, c)
	}
	if c > 0 && grid[r][c-1] > v {
		size += s.getBasinSizeRec(grid, r, c-1)
	}
	if r < len(grid)-1 && grid[r+1][c] > v {
		size += s.getBasinSizeRec(grid, r+1, c)
	}
	if c < len(grid[r])-1 && grid[r][c+1] > v {
		size += s.getBasinSizeRec(grid, r, c+1)
	}

	return size
}

func sumRiskLevels(grid [][]uint8) int {
	var sum int
	for r, row := range grid {
		for c, v := range row {
			if isLowPoint(grid, r, c) {
				sum += int(v + 1)
			}
		}
	}
	return sum
}

func isLowPoint(grid [][]uint8, r, c int) bool {
	v := grid[r][c]
	if r > 0 && grid[r-1][c] <= v {
		return false
	}
	if c > 0 && grid[r][c-1] <= v {
		return false
	}
	if r < len(grid)-1 && grid[r+1][c] <= v {
		return false
	}
	if c < len(grid[r])-1 && grid[r][c+1] <= v {
		return false
	}

	return true
}

func getHeightMap() [][]uint8 {
	lines, err := vtools.ReadLines(os.Args[1])
	if err != nil {
		panic(err)
	}

	hm := make([][]uint8, 0, len(lines))
	for _, line := range lines {
		row := make([]uint8, 0, len(lines))
		for _, c := range line {
			row = append(row, uint8(c-'0'))
		}

		hm = append(hm, row)
	}

	return hm
}

func makeBigInput() {
	lines, err := vtools.ReadLines(os.Args[1])
	if err != nil {
		panic(err)
	}

	data := make([]byte, 0, 1<<20)
	for i, line := range lines {
		for range 100 {
			data = append(data, line...)
		}

		if i < len(lines)-1 {
			data = append(data, '\n')
		}
	}

	if err := os.WriteFile("big_input.txt", data, 0777); err != nil {
		panic(fmt.Errorf("failed to write big input: %v", err))
	}
}
