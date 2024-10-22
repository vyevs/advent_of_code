package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"

	"github.com/vyevs/vtools"
)

func main() {
	moves := getInput()

	part1(moves)
	part2(moves)
}

func part1(moves []move) {
	grid := make([][]bool, 1000)
	for i := range grid {
		grid[i] = make([]bool, 1000)
	}

	for _, m := range moves {
		for x := m.x0; x <= m.x1; x++ {
			for y := m.y0; y <= m.y1; y++ {
				if m.toggle {
					grid[x][y] = !grid[x][y]
				} else {
					grid[x][y] = m.turnOn
				}
			}
		}
	}

	var on int
	for _, row := range grid {
		on += vtools.CountSlice(row, true)
	}

	fmt.Printf("part 1: %d lights are on\n", on)
}

func part2(moves []move) {
	grid := make([][]int, 1000)
	for i := range grid {
		grid[i] = make([]int, 1000)
	}

	for _, m := range moves {
		for x := m.x0; x <= m.x1; x++ {
			for y := m.y0; y <= m.y1; y++ {
				if m.toggle {
					grid[x][y] += 2
				} else if m.turnOn {
					grid[x][y] = grid[x][y] + 1
				} else {
					grid[x][y] = max(grid[x][y]-1, 0)
				}
			}
		}
	}

	var totalBrightness int
	for _, row := range grid {
		for _, v := range row {
			totalBrightness += v
		}
	}

	fmt.Printf("part 2: total brightness is %d\n", totalBrightness)
}

type move struct {
	toggle         bool
	turnOn         bool
	x0, y0, x1, y1 int
}

func getInput() []move {
	lines, err := vtools.ReadLines(os.Args[1])
	if err != nil {
		log.Fatalf("ReadLines: %v", err)
	}

	moves := make([]move, 0, 1024)
	for _, line := range lines {
		moves = append(moves, parseLine(line))
	}

	return moves
}

var re = regexp.MustCompile(`^(turn|toggle)\s*(on|off)?\s*(\d+),(\d+) through (\d+),(\d+)$`)

func parseLine(line string) move {
	groups := re.FindStringSubmatch(line)

	m := move{
		toggle: groups[1] == "toggle",
	}
	if !m.toggle {
		m.turnOn = groups[2] == "on"
	}

	x0, _ := strconv.Atoi(groups[3])
	y0, _ := strconv.Atoi(groups[4])
	x1, _ := strconv.Atoi(groups[5])
	y1, _ := strconv.Atoi(groups[6])

	m.x0 = x0
	m.y0 = y0
	m.x1 = x1
	m.y1 = y1

	return m
}
