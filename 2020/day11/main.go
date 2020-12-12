package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input := readInput()

	fmt.Println("Part 1:")
	doPart1(input)
	fmt.Println()

	fmt.Println("Part 2:")
	doPart2(input)
}

func doPart1(grid [][]byte) {
	gridCopy := make([][]byte, len(grid))
	for i := 0; i < len(grid); i++ {
		gridCopy[i] = make([]byte, len(grid[i]))
		copy(gridCopy[i], grid[i])
	}

	curState := gridCopy
	var nextState [][]byte
	for {
		//printGrid(curState)
		nextState = doOneStep1(curState)

		if sameStates(curState, nextState) {
			break
		}

		curState = nextState
	}

	nOccupied := countOccupied(nextState)

	fmt.Printf("At the end there are %d seats occupied\n", nOccupied)
}

func sameStates(s1, s2 [][]byte) bool {
	for i := 0; i < len(s1); i++ {
		for j := 0; j < len(s1[i]); j++ {
			if s1[i][j] != s2[i][j] {
				return false
			}
		}
	}
	return true
}

func countOccupied(grid [][]byte) int {
	var n int
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '#' {
				n++
			}
		}
	}
	return n
}

func doOneStep1(grid [][]byte) [][]byte {
	newGrid := make([][]byte, len(grid))
	for i := 0; i < len(grid); i++ {
		newGrid[i] = make([]byte, len(grid[i]))
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			nNeighbors := nOccupiedNeighbors(grid, i, j)

			if grid[i][j] == 'L' && nNeighbors == 0 {
				newGrid[i][j] = '#'
			} else if grid[i][j] == '#' && nNeighbors >= 4 {
				newGrid[i][j] = 'L'
			} else {
				newGrid[i][j] = grid[i][j]
			}
		}
	}

	return newGrid
}

func doOneStep2(grid [][]byte) [][]byte {
	newGrid := make([][]byte, len(grid))
	for i := 0; i < len(grid); i++ {
		newGrid[i] = make([]byte, len(grid[i]))
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			nNeighbors := nOccupiedInDirections(grid, i, j)

			if grid[i][j] == 'L' && nNeighbors == 0 {
				newGrid[i][j] = '#'
			} else if grid[i][j] == '#' && nNeighbors >= 5 {
				newGrid[i][j] = 'L'
			} else {
				newGrid[i][j] = grid[i][j]
			}
		}
	}

	return newGrid
}

func nOccupiedNeighbors(grid [][]byte, row, col int) int {
	var n int

	if row-1 >= 0 {
		if col-1 >= 0 {
			if grid[row-1][col-1] == '#' {
				n++
			}
		}
		if col+1 < len(grid[row]) {
			if grid[row-1][col+1] == '#' {
				n++
			}
		}
		if grid[row-1][col] == '#' {
			n++
		}
	}

	if row+1 < len(grid) {
		if col-1 >= 0 {
			if grid[row+1][col-1] == '#' {
				n++
			}
		}
		if col+1 < len(grid[row]) {
			if grid[row+1][col+1] == '#' {
				n++
			}
		}
		if grid[row+1][col] == '#' {
			n++
		}
	}

	if col-1 >= 0 {
		if grid[row][col-1] == '#' {
			n++
		}
	}

	if col+1 < len(grid[row]) {
		if grid[row][col+1] == '#' {
			n++
		}
	}

	return n
}

func nOccupiedInDirections(grid [][]byte, row, col int) int {
	var n int

	if seatInDirection(grid, row, col, -1, -1) == '#' {
		n++
	}
	if seatInDirection(grid, row, col, -1, 0) == '#' {
		n++
	}
	if seatInDirection(grid, row, col, -1, 1) == '#' {
		n++
	}
	if seatInDirection(grid, row, col, 0, -1) == '#' {
		n++
	}
	if seatInDirection(grid, row, col, 0, 1) == '#' {
		n++
	}
	if seatInDirection(grid, row, col, 1, -1) == '#' {
		n++
	}
	if seatInDirection(grid, row, col, 1, 0) == '#' {
		n++
	}
	if seatInDirection(grid, row, col, 1, 1) == '#' {
		n++
	}

	return n
}

func seatInDirection(grid [][]byte, row, col, drow, dcol int) byte {
	for {
		row += drow
		col += dcol

		if row < 0 || row >= len(grid) || col < 0 || col >= len(grid[row]) {
			break
		}

		if grid[row][col] == 'L' || grid[row][col] == '#' {
			return grid[row][col]
		}
	}

	return '.'
}

func doPart2(grid [][]byte) {
	gridCopy := make([][]byte, len(grid))
	for i := 0; i < len(grid); i++ {
		gridCopy[i] = make([]byte, len(grid[i]))
		copy(gridCopy[i], grid[i])
	}

	curState := gridCopy
	var nextState [][]byte
	for {
		//printGrid(curState)
		nextState = doOneStep2(curState)

		if sameStates(curState, nextState) {
			break
		}

		curState = nextState
	}

	nOccupied := countOccupied(nextState)

	fmt.Printf("At the end there are %d seats occupied\n", nOccupied)
}

func readInput() [][]byte {
	scanner := bufio.NewScanner(os.Stdin)

	grid := make([][]byte, 0, 128)
	for scanner.Scan() {
		line := scanner.Text()

		row := []byte(line)
		grid = append(grid, row)
	}

	return grid

}
