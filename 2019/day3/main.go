package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	//doPart1()
	doPart2()
}

func doPart1()  {
	wire1Path, wire2Path := readInWirePaths()

	minY, maxY, minX, maxX := determineGridDimensions(wire1Path, wire2Path)
	fmt.Printf("minY: %d maxY: %d minX: %d maxX: %d\n", minY, maxY, minX, maxX)
	height, width := maxY-minY+1, maxX-minX+1

	fmt.Printf("grid size: height %d width %d\n", height, width)

	grid := makeSizedGrid(height, width)

	wireY0 := height - 1 - maxY
	wireX0 := width - 1 - maxX

	fmt.Printf("wire origin row %d col %d\n", wireY0, wireX0)

	drawWireOnGrid(grid, wire1Path, true, wireY0, wireX0)

	drawWireOnGrid(grid, wire2Path, false, wireY0, wireX0)

	dist := distanceToClosestIntersection(grid, wireY0, wireX0)

	fmt.Printf("distance to closest intersection: %d\n", dist)
}

func doPart2()  {
	wire1Path, wire2Path := readInWirePaths()

	minY, maxY, minX, maxX := determineGridDimensions(wire1Path, wire2Path)
	fmt.Printf("minY: %d maxY: %d minX: %d maxX: %d\n", minY, maxY, minX, maxX)
	height, width := maxY-minY+1, maxX-minX+1

	fmt.Printf("grid size: height %d width %d\n", height, width)

	grid := makeSizedGrid(height, width)

	wireY0 := height - 1 - maxY
	wireX0 := width - 1 - maxX

	fmt.Printf("wire origin row %d col %d\n", wireY0, wireX0)

	drawWireOnGrid(grid, wire1Path, true, wireY0, wireX0)

	drawWireOnGrid(grid, wire2Path, false, wireY0, wireX0)

	printGrid(grid)

	leastSteps := determineLeastTotalStepsToIntersection(grid, wire1Path, wire2Path, wireX0, wireY0)

	fmt.Printf("least total steps to intersection: %d\n", leastSteps)
}

func determineLeastTotalStepsToIntersection(grid [][]gridPoint, wire1Path, wire2Path []string, x0, y0 int) int {
	var minSum int
	for y, row := range grid {
		for x, p := range row {
			if p != gridPointIntersection {
				continue
			}

			stepsForW1 := determineStepsForWireToReachPoint(wire1Path, y0, x0, y, x)
			stepsForW2 := determineStepsForWireToReachPoint(wire2Path, y0, x0, y, x)

			sum := stepsForW1 + stepsForW2

			if minSum == 0 || sum < minSum {
				minSum = sum
			}
		}
	}

	return minSum
}

func determineStepsForWireToReachPoint(path []string, y0, x0, targetY, targetX int) int {
	curY, curX := y0, x0

	stepInDirection := func(dir byte) {
		switch dir {
		case 'R': curX++
		case 'L': curX--
		case 'U': curY--
		case 'D': curY++
		default: panic(fmt.Sprintf("unknown dir %c", dir))
		}
	}


	var nSteps int
	for _, section := range path {
		dir := section[0]
		dist, _ := strconv.Atoi(section[1:])

		for i := 0; i < dist; i++ {
			stepInDirection(dir)
			nSteps++

			if curY == targetY && curX == targetX {
				return nSteps
			}
		}
	}

	panic("did not reach target")
}

func distanceToClosestIntersection(grid [][]gridPoint, y0, x0 int) int {
	bestDist := len(grid) * 10000
	for y, row := range grid {
		for x, p := range row {
			if p != gridPointIntersection {
				continue
			}

			dist := intAbs(y-y0) + intAbs(x-x0)
			if dist < bestDist {
				bestDist = dist
			}
		}
	}
	return bestDist
}

func intAbs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func readInWirePaths() ([]string, []string) {
	scanner := bufio.NewScanner(os.Stdin)

	readWire := func() []string {
		scanner.Scan()
		line := scanner.Text()
		steps := strings.Split(line, ",")
		return steps
	}

	wire1Steps := readWire()
	wire2Steps := readWire()

	return wire1Steps, wire2Steps
}

func determineGridDimensions(w1Path, w2Path []string) (int, int, int, int) {
	w1MinY, w1MaxY, w1MinX, w1MaxX := wireSpans(w1Path)
	w2MinY, w2MaxY, w2MinX, w2MaxX := wireSpans(w2Path)

	minY := w1MinY
	if w2MinY < minY {
		minY = w2MinY
	}

	maxY := w1MaxY
	if w2MaxY > maxY {
		maxY = w2MaxY
	}

	minX := w1MinX
	if w2MinX < minX {
		minX = w2MinX
	}

	maxX := w1MaxX
	if w2MaxX > maxX {
		maxX = w2MaxX
	}

	return minY, maxY, minX, maxX
}

func wireSpans(wirePath []string) (minY, maxY, minX, maxX int) {
	var curX, curY int

	processMove := func(move string) {
		dir := move[0]
		dist, _ := strconv.Atoi(move[1:])

		switch dir {
		case 'R':
			curX += dist
			if curX > maxX {
				maxX = curX
			}
		case 'L':
			curX -= dist
			if curX < minX {
				minX = curX
			}
		case 'U':
			curY -= dist
			if curY < minY {
				minY = curY
			}
		case 'D':
			curY += dist
			if curY > maxY {
				maxY = curY
			}

		default: panic(fmt.Sprintf("unknown direction %c", dir))
		}
	}
	for _, move := range wirePath {
		processMove(move)
	}

	return
}

func makeSizedGrid(height, width int) [][]gridPoint {
	grid := make([][]gridPoint, height)
	widths := make([]gridPoint, height*width)
	for i := 0; i < height; i++ {
		grid[i] = widths[i*width : i*width+width]
		for j := 0; j < width; j++ {
			grid[i][j] = gridPointEmpty
		}
	}
	return grid
}

type gridPoint int

const (
	gridPointEmpty gridPoint = iota
	gridPointOrigin
	gridPointIntersection
	gridPointTurn

	gridPointW1Vertical
	gridPointW1Horizontal

	gridPointW2Vertical
	gridPointW2Horizontal
)

var gridPointRunes = map[gridPoint]rune{
	gridPointEmpty:        '.',
	gridPointOrigin:       'O',
	gridPointIntersection: 'X',
	gridPointTurn:         '+',
	gridPointW1Vertical:   '|',
	gridPointW1Horizontal: '-',
	gridPointW2Vertical:   '|',
	gridPointW2Horizontal: '-',
}

func drawWireOnGrid(grid [][]gridPoint, path []string, wire1 bool, y0, x0 int) {
	curY, curX := y0, x0

	stepInDirection := func(dir byte) {
		switch dir {
		case 'R': curX++
		case 'L': curX--
		case 'U': curY--
		case 'D': curY++
		default: panic(fmt.Sprintf("unknown dir %c", dir))
		}
	}

	determineGridPoint := func(dir byte) gridPoint {
		if curY == y0 && curX == x0 {
			return gridPointOrigin
		}
		if wire1 {
			if grid[curY][curX] == gridPointW2Vertical || grid[curY][curX] == gridPointW2Horizontal {
				return gridPointIntersection
			}
			switch dir {
				case 'U', 'D': return gridPointW1Vertical
				case 'L', 'R': return gridPointW1Horizontal
				default: panic(fmt.Sprintf("unknown dir %c", dir))
			}

		} else {
			if grid[curY][curX] == gridPointW1Vertical || grid[curY][curX] == gridPointW1Horizontal {
				return gridPointIntersection
			}
			switch dir {
				case 'U', 'D': return gridPointW2Vertical
				case 'L', 'R': return gridPointW2Horizontal
				default: panic(fmt.Sprintf("unknown dir %c", dir))
			}
		}
	}

	for moveNum, move := range path {
		dir := move[0]
		dist, _ := strconv.Atoi(move[1:])

		for i := 0; i < dist; i++ {
			dirGridPoint := determineGridPoint(dir)

			grid[curY][curX] = dirGridPoint

			stepInDirection(dir)
		}

		// the move we just made in a particular direction either ends with us making a turn and going in another direction so we write a '+'
		// or this move was the final one for the wire so we just write the final direction rune
		if moveNum < len(path)-1 {
			if grid[curY][curX] == '|' || grid[curY][curX] == '-' {
				grid[curY][curX] = 'X'
			} else {
				grid[curY][curX] = '+'
			}
		} else {
			if grid[curY][curX] == '|' || grid[curY][curX] == '-' {
				grid[curY][curX] = 'X'
			} else {
				grid[curY][curX] = determineGridPoint(dir)
			}
		}
	}
}

func printGrid(grid [][]gridPoint) {
	var buf strings.Builder

	for _, row := range grid {
		for _, p := range row {
			buf.WriteRune(gridPointRunes[p])
		}
		buf.WriteRune('\n')
	}

	fmt.Println(buf.String())
}
