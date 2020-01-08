package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, _ := os.Open(os.Args[1])

	instructions := parseInput(f)
	f.Close()

	var grid [6][50]bool

	for _, inst := range instructions {

		switch inst.op {
		case rect:

			for row := 0; row < inst.row; row++ {
				for col := 0; col < inst.col; col++ {
					grid[row][col] = true
				}
			}

		case rotateRow:

			shftAmt := inst.shftAmt

			oldRow := grid[inst.row]

			for col := 0; col < 50; col++ {
				grid[inst.row][(col+shftAmt)%50] = oldRow[col]
			}

		case rotateCol:
			shftAmt := inst.shftAmt

			var oldCol [6]bool
			for i := 0; i < 6; i++ {
				oldCol[i] = grid[i][inst.col]
			}

			for row := 0; row < 6; row++ {
				grid[(row+shftAmt)%6][inst.col] = oldCol[row]
			}
		}
	}

	printGrid(grid)

	var on int
	for _, row := range grid {
		for _, col := range row {
			if col {
				on++
			}
		}
	}

	fmt.Println(on)
}

func printGrid(grid [6][50]bool) {
	for _, row := range grid {
		for _, col := range row {
			if col {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

type op int

const (
	rect op = iota
	rotateRow
	rotateCol
)

type instruction struct {
	op op

	row, col int

	shftAmt int
}

func parseInput(r io.Reader) []instruction {
	scanner := bufio.NewScanner(r)

	instructions := make([]instruction, 0, 512)

	for scanner.Scan() {
		line := scanner.Text()

		tokens := strings.Split(line, " ")

		var inst instruction

		if tokens[0] == "rect" {
			inst.op = rect

			dimensionStrs := strings.Split(tokens[1], "x")

			inst.col, _ = strconv.Atoi(dimensionStrs[0])
			inst.row, _ = strconv.Atoi(dimensionStrs[1])

		} else if tokens[1] == "column" {
			inst.op = rotateCol

			colStr := strings.Split(tokens[2], "=")[1]

			inst.col, _ = strconv.Atoi(colStr)
			inst.shftAmt, _ = strconv.Atoi(tokens[4])

		} else {
			inst.op = rotateRow

			rowStr := strings.Split(tokens[2], "=")[1]

			inst.row, _ = strconv.Atoi(rowStr)
			inst.shftAmt, _ = strconv.Atoi(tokens[4])
		}

		instructions = append(instructions, inst)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("scanner: %v", err)
	}

	return instructions
}
