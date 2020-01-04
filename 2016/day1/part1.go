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

type cardinalDirection int

const (
	north cardinalDirection = iota
	east
	south
	west
)

type turnDirection int

const (
	left turnDirection = iota
	right
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("provide 1 input file")
	}

	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatalf("open: %v", err)
	}

	moves := parseInput(f)

	x, y := doMoves(moves)

	fmt.Printf("ended at (%d, %d)\n", x, y)

	absX := abs(x)
	absY := abs(y)

	fmt.Println(absX + absY)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func doMoves(moves []move) (int, int) {
	direction := north
	var x, y int

	type coord struct {
		x, y int
	}

	visited := make(map[coord]struct{}, 128)
	visited[coord{0, 0}] = struct{}{}

	for _, move := range moves {
		if move.turnDir == left {
			direction--
			if direction < north {
				direction = west
			}
		} else {
			direction++
			if direction > west {
				direction = north
			}
		}

		/*
			switch direction {
			case north:
				//y += move.steps
				for i := 0; i < move.steps; i++ {
					y++
				}
			case east:
				x += move.steps
				for i := 0; i < move.steps; i++ {
					x++
				}
			case south:
				y -= move.steps
				for i := 0; i < move.steps; i++ {
					y--
				}
			case west:
				x -= move.steps
				for i := 0; i < move.steps; i++ {
					x--
				}
			}*/

		for i := 0; i < move.steps; i++ {
			switch direction {
			case north:
				y++
			case east:
				x++
			case south:
				y--
			case west:
				x--
			}

			if _, ok := visited[coord{x, y}]; ok {
				fmt.Printf("visited (%d, %d) twice\n", x, y)
			} else {
				visited[coord{x, y}] = struct{}{}
			}
		}

	}

	return x, y
}

type move struct {
	turnDir turnDirection
	steps   int
}

func parseInput(r io.Reader) []move {
	moves := make([]move, 0, 256)

	scanner := bufio.NewScanner(r)
	scanner.Scan()
	line := scanner.Text()

	if err := scanner.Err(); err != nil {
		log.Fatalf("scanner: %v", err)
	}

	moveStrs := strings.Split(line, ", ")

	for i, moveStr := range moveStrs {

		var move move

		if moveStr[0] == 'R' {
			move.turnDir = right
		} else {
			move.turnDir = left
		}

		steps, err := strconv.Atoi(moveStr[1:])
		if err != nil {
			log.Fatalf("error parsing number of steps from move %s (line %d)", moveStr, i+1)
		}

		move.steps = steps

		moves = append(moves, move)
	}

	return moves
}
