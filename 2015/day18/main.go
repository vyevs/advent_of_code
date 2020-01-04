package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("provide 1 input file")
	}

	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatalf("open: %v", err)
	}

	grid := parseInput(f)
	grid[0][0] = true
	grid[0][len(grid[0])-1] = true
	grid[len(grid)-1][0] = true
	grid[len(grid)-1][len(grid[len(grid)-1])-1] = true

	//printGrid(grid)

	for i := 0; i < 100; i++ {
		grid = nextStep2(grid)
		//printGrid(grid)
	}

	fmt.Println(numOn(grid))
}

func numOn(grid [][]bool) int {
	var on int
	for _, row := range grid {
		for _, cell := range row {
			if cell {
				on++
			}
		}
	}
	return on
}

func nextStep(grid [][]bool) [][]bool {
	newGrid := make([][]bool, len(grid))
	for i, row := range grid {
		newGrid[i] = make([]bool, len(row))
	}

	for i, row := range grid {
		for j, cell := range row {
			on := onNeighbors(grid, i, j)

			if on == 3 || (cell && on == 2) {
				newGrid[i][j] = true
			}
		}
	}

	return newGrid
}

func nextStep2(grid [][]bool) [][]bool {
	newGrid := make([][]bool, len(grid))
	for i, row := range grid {
		newGrid[i] = make([]bool, len(row))
	}

	newGrid[0][0] = true
	newGrid[0][len(grid[0])-1] = true
	newGrid[len(grid)-1][0] = true
	newGrid[len(grid)-1][len(grid[len(grid)-1])-1] = true

	for i, row := range grid {
		for j, cell := range row {
			if isCorner(grid, i, j) {
				continue
			}
			on := onNeighbors(grid, i, j)

			if on == 3 || (cell && on == 2) {
				newGrid[i][j] = true
			}
		}
	}

	return newGrid
}

func isCorner(grid [][]bool, i, j int) bool {
	return (i == 0 && j == 0) ||
		(i == len(grid)-1 && j == 0) ||
		(i == 0 && j == len(grid[i])-1) ||
		(i == len(grid)-1 && j == len(grid[len(grid)-1])-1)
}

func printGrid(grid [][]bool) {
	for _, row := range grid {
		for _, cell := range row {
			if cell {
				fmt.Printf("%c", '#')
			} else {
				fmt.Printf("%c", '.')
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func onNeighbors(grid [][]bool, i, j int) int {
	upRow := i - 1
	downRow := i + 1
	leftCol := j - 1
	rightCol := j + 1

	var on int
	if upRow != -1 {
		if leftCol != -1 && grid[upRow][leftCol] {
			on++
		}
		if rightCol != len(grid[upRow]) && grid[upRow][rightCol] {
			on++
		}
		if grid[upRow][j] {
			on++
		}
	}

	if downRow != len(grid) {
		if leftCol != -1 && grid[downRow][leftCol] {
			on++
		}
		if rightCol != len(grid[downRow]) && grid[downRow][rightCol] {
			on++
		}
		if grid[downRow][j] {
			on++
		}
	}

	if leftCol != -1 && grid[i][leftCol] {
		on++
	}
	if rightCol != len(grid[i]) && grid[i][rightCol] {
		on++
	}

	return on
}

func parseInput(r io.Reader) [][]bool {
	reader := bufio.NewReader(r)

	grid := make([][]bool, 0, 128)
	grid = append(grid, make([]bool, 0, 128))

	var i int
	for {
		b, err := reader.ReadByte()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				log.Fatalf("read error: %v", err)
			}
		}

		if b == '#' {
			grid[i] = append(grid[i], true)

		} else if b == '.' {
			grid[i] = append(grid[i], false)

		} else if b == 13 { // carriage return
			i++
			grid = append(grid, make([]bool, 0, 128))

			_, err = reader.ReadByte()
			if err != nil {
				log.Fatalf("trying to read line feed: %v", err)
			}
		} else {
			log.Fatalf("unknown byte %d", b)
		}
	}

	return grid
}
