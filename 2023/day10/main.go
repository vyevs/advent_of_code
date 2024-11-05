package main

import (
	"fmt"
	"os"
	"slices"

	"github.com/vyevs/vtools"
)

type cell struct {
	tile          byte
	distFromStart int // -1 indicates that the cell is not part of the loop.
	visited       bool
}

func main() {
	grid := getGrid()
	if len(grid) < 100 {
		printGrid(grid)
		fmt.Println()
	}

	steps := stepsToFarthestPosition(grid)
	fmt.Printf("part 1: %d steps to the farthest point\n", steps)

	area := areaEnclosedByPipe(grid)
	fmt.Printf("part 2: %d is the enclosed area\n", area)
}

var tileDirections = func() ['|' + 1][][2]int {
	var a ['|' + 1][][2]int
	a['-'] = [][2]int{{0, -1}, {0, 1}}
	a['7'] = [][2]int{{1, 0}, {0, -1}}
	a['F'] = [][2]int{{1, 0}, {0, 1}}
	a['J'] = [][2]int{{-1, 0}, {0, -1}}
	a['L'] = [][2]int{{-1, 0}, {0, 1}}
	a['S'] = [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	a['|'] = [][2]int{{-1, 0}, {1, 0}}
	return a
}()

// getLoopPath returns the list of points (in order) that make up the loop.
// The points may be positively or negatively oriented, depends on the loop.
func getLoopPath(grid [][]cell) [][2]int {
	y0, x0 := findStartPoint(grid)

	path := make([][2]int, 0, 128)
	grid[y0][x0].visited = true
	path = append(path, [2]int{y0, x0})

	dirs := tileDirections['S']
	for _, dir := range dirs {
		dy, dx := dir[0], dir[1]

		y1, x1 := y0+dy, x0+dx

		if y1 < 0 || x1 < 0 || y1 >= len(grid) || x1 >= len(grid[y1]) {
			continue
		}

		nextTile := grid[y1][x1].tile
		shouldMoveInDir := false
		if dy == -1 && (nextTile == '|' || nextTile == '7' || nextTile == 'F') {
			shouldMoveInDir = true
		} else if dy == 1 && (nextTile == '|' || nextTile == 'J' || nextTile == 'L') {
			shouldMoveInDir = true
		} else if dx == -1 && (nextTile == '-' || nextTile == 'F' || nextTile == 'L') {
			shouldMoveInDir = true
		} else if dx == 1 && (nextTile == '-' || nextTile == '7' || nextTile == 'J') {
			shouldMoveInDir = true
		}

		// The first direction we can move in is the one we take.
		if shouldMoveInDir {
			getLoopPathRec(grid, y1, x1, &path)
			break
		}
	}

	return path
}

func getLoopPathRec(grid [][]cell, y, x int, path *[][2]int) {
	if grid[y][x].visited {
		return
	}

	grid[y][x].visited = true
	*path = append(*path, [2]int{y, x})

	tile := grid[y][x].tile

	dirs := tileDirections[tile]
	for _, dir := range dirs {
		dy, dx := dir[0], dir[1]
		y1, x1 := y+dy, x+dx
		getLoopPathRec(grid, y1, x1, path)
	}
}

func areaEnclosedByPipe(grid [][]cell) int {
	loopPath := getLoopPath(grid)
	area := polygonArea(loopPath)

	// Pick's theorem that relates the number of integer points on a polygon's boundary and interior.
	// https://en.wikipedia.org/wiki/Pick%27s_theorem
	interiorArea := area - len(loopPath)/2 + 1
	return interiorArea
}

// returns the area of enclosed by the points, boundary included.
// the provided points must be positively oriented.
func polygonArea(points [][2]int) int {

	// Shoelace formula for calculating the area of a polygon with integer vertices.
	// https://en.wikipedia.org/wiki/Shoelace_formula

	var area int
	for i := 0; i < len(points)-1; i++ {
		p0, p1 := points[i], points[i+1]
		area += (p0[1] - p1[1]) * (p1[0] + p0[0])
	}

	// Extra case to connect the last point to the first point.
	{
		p0, p1 := points[len(points)-1], points[0]
		area += (p0[1] - p1[1]) * (p1[0] + p0[0])
	}

	// The math above is done assuming points are positively oriented AKA counter-clockwise so that when we walk them the polygon is to our left.
	// This means that for points that are negatively oriented, the result will be negative, so take the abs of area.
	return vtools.Abs(area) / 2
}

func stepsToFarthestPosition(grid [][]cell) int {
	y0, x0 := findStartPoint(grid)

	grid[y0][x0].distFromStart = 0

	queue := vtools.NewQueue[[2]int](2)
	queue.Push([2]int{y0, x0})

	for len(queue) != 0 {
		loc := queue.Pop()
		y, x := loc[0], loc[1]

		cell := grid[y][x]
		tile := cell.tile
		dist := cell.distFromStart

		dirs := tileDirections[tile]

		if tile == 'S' {

			for _, dir := range dirs {
				dy, dx := dir[0], dir[1]

				newR, newC := y+dy, x+dx

				if newR < 0 || newC < 0 || newR >= len(grid) || newC >= len(grid[newR]) {
					continue
				}

				nextTile := grid[newR][newC].tile
				shouldMoveInDir := false
				if dy == -1 && (nextTile == '|' || nextTile == '7' || nextTile == 'F') {
					shouldMoveInDir = true
				} else if dy == 1 && (nextTile == '|' || nextTile == 'J' || nextTile == 'L') {
					shouldMoveInDir = true
				} else if dx == -1 && (nextTile == '-' || nextTile == 'F' || nextTile == 'L') {
					shouldMoveInDir = true
				} else if dx == 1 && (nextTile == '-' || nextTile == '7' || nextTile == 'J') {
					shouldMoveInDir = true
				}

				if shouldMoveInDir {
					grid[newR][newC].distFromStart = 1
					queue.Push([2]int{newR, newC})
				}
			}

		} else {
			for _, dir := range dirs {
				dy, dx := dir[0], dir[1]

				newR, newC := y+dy, x+dx
				if grid[newR][newC].distFromStart == -1 {
					grid[newR][newC].distFromStart = dist + 1
					queue.Push([2]int{newR, newC})
				}
			}
		}
	}

	var maxDist int
	for _, row := range grid {
		rowMax := slices.MaxFunc(row, func(a, b cell) int { return a.distFromStart - b.distFromStart }).distFromStart
		if rowMax > maxDist {
			maxDist = rowMax
		}
	}

	return maxDist
}

// returns the (y, x) position of the start point 'S' in grid.
func findStartPoint(grid [][]cell) (int, int) {
	for y, row := range grid {
		for x, cell := range row {
			if cell.tile == 'S' {
				return y, x
			}
		}
	}

	panic("there is no start position")
}

func printGrid(grid [][]cell) {
	for _, row := range grid {
		for _, cell := range row {
			tile := cell.tile

			r := rune(tile)

			switch tile {
			case 'F':
				r = '┏'
			case '7':
				r = '┓'
			case 'J':
				r = '┛'
			case 'L':
				r = '┗'
			case '|':
				r = '┃'
			case '-':
				r = '━'
			}

			fmt.Printf("%c", r)
		}
		fmt.Println()
	}
}

func getGrid() [][]cell {
	lines, err := vtools.ReadLines(os.Args[1])
	if err != nil {
		panic(err)
	}

	grid := make([][]cell, 0, len(lines))
	for _, line := range lines {
		row := make([]cell, 0, len(line))

		for i := range line {
			tile := line[i]
			cell := cell{
				tile:          tile,
				distFromStart: -1,
			}
			row = append(row, cell)
		}

		grid = append(grid, row)
	}

	return grid
}
