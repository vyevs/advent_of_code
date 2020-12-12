package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	input := readInput()

	fmt.Println("Part 1:")
	doPart1(input)
	fmt.Println()

	fmt.Println("Part 2:")
	doPart2(input)
}

type direction int

const (
	directionEast direction = iota
	directionNorth
	directionWest
	directionSouth
)

func (d direction) turnLeft(degs int) direction {
	turns := degs / 90

	leftDir := d
	for turns > 0 {
		leftDir += 1
		leftDir %= (directionSouth + 1)
		turns--
	}
	return leftDir
}

func (d direction) turnRight(degs int) direction {
	turns := degs / 90

	rightDir := d
	for turns > 0 {
		rightDir -= 1
		if rightDir < 0 {
			rightDir = directionSouth
		}
		turns--
	}
	return rightDir
}

func doPart1(moves []move) {
	x, y := 0, 0

	curDir := directionEast

	for _, move := range moves {
		if move.typ == 'L' {
			curDir = curDir.turnLeft(move.mag)
		} else if move.typ == 'R' {
			curDir = curDir.turnRight(move.mag)
		} else if move.typ == 'F' {
			switch curDir {
			case directionNorth:
				y += move.mag
			case directionSouth:
				y -= move.mag
			case directionEast:
				x += move.mag
			case directionWest:
				x -= move.mag
			}
		} else if move.typ == 'N' {
			y += move.mag
		} else if move.typ == 'S' {
			y -= move.mag
		} else if move.typ == 'E' {
			x += move.mag
		} else if move.typ == 'W' {
			x -= move.mag
		}
	}

	fmt.Printf("We end up at [%d, %d], which is %d manhattan distance from the start\n", x, y, intAbs(x)+intAbs(y))
}

func intAbs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func doPart2(moves []move) {
	x, y := 0, 0
	wpx, wpy := 10, 1

	for _, move := range moves {
		if move.typ == 'L' {
			turns := move.mag / 90

			for turns > 0 {
				wpx, wpy = rotateCcwAboutOrigin(wpx, wpy)
				turns--
			}

		} else if move.typ == 'R' {
			turns := move.mag / 90

			for turns > 0 {
				wpx, wpy = rotateCwAboutOrigin(wpx, wpy)
				turns--
			}

		} else if move.typ == 'F' {
			x += move.mag * wpx
			y += move.mag * wpy

		} else if move.typ == 'N' {
			wpy += move.mag
		} else if move.typ == 'S' {
			wpy -= move.mag
		} else if move.typ == 'E' {
			wpx += move.mag
		} else if move.typ == 'W' {
			wpx -= move.mag
		}
	}

	fmt.Printf("We end up at [%d, %d], which is %d manhattan distance from the start\n", x, y, intAbs(x)+intAbs(y))
}

func rotateCwAboutOrigin(x, y int) (int, int) {
	return y, -x
}

func rotateCcwAboutOrigin(x, y int) (int, int) {
	return -y, x
}

type move struct {
	typ byte
	mag int
}

func readInput() []move {
	scanner := bufio.NewScanner(os.Stdin)

	moves := make([]move, 0, 1024)
	for scanner.Scan() {
		line := scanner.Text()

		mag, _ := strconv.Atoi(line[1:])
		move := move{
			typ: line[0],
			mag: mag,
		}
		moves = append(moves, move)
	}

	return moves
}
