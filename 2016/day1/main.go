package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	north byte = iota
	east
	south
	west
	allDirections
)

func turnLeft(d byte) byte {
	return (d - 1) % allDirections
}

func turnRight(d byte) byte {
	return (d + 1) % allDirections
}

func main() {
	defer func(start time.Time) {
		fmt.Printf("that took %v\n", time.Since(start))
	}(time.Now())

	bs, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("failed to read input file: %v", err)
	}

	moves := parseInput(string(bs))

	x, y := doMoves(moves)

	fmt.Printf("ended at (%d, %d)\n", x, y)

	absX := abs(x)
	absY := abs(y)

	fmt.Println(absX + absY)

	x, y = firstDuplicateLocation(moves)
	fmt.Printf("The first location we visit twice is (%d, %d), which is %d away\n", x, y, abs(x)+abs(y))
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func firstDuplicateLocation(moves []move) (int, int) {
	dir := north
	var here [2]int

	visited := make(map[[2]int]struct{}, 1024)
	visited[here] = struct{}{}

	for _, move := range moves {
		switch move.turnDir {
		case 'L':
			dir = turnLeft(dir)
		case 'R':
			dir = turnRight(dir)
		default:
			panic(fmt.Sprintf("unknown dir %q", move.turnDir))
		}

		for range move.steps {
			switch dir {
			case north:
				here[1]++
			case east:
				here[0]++
			case south:
				here[1]--
			case west:
				here[0]--
			default:
				panic(fmt.Sprintf("unknown dir %q", dir))
			}

			_, beenHere := visited[here]
			if beenHere {
				return here[0], here[1]
			}
			visited[here] = struct{}{}
		}

	}

	panic("no location was visited twice")
}

func doMoves(moves []move) (int, int) {
	dir := north
	var here [2]int

	for _, move := range moves {
		switch move.turnDir {
		case 'L':
			dir = turnLeft(dir)
		case 'R':
			dir = turnRight(dir)
		default:
			panic(fmt.Sprintf("unknown dir %q", move.turnDir))
		}

		switch dir {
		case north:
			here[1] += move.steps
		case east:
			here[0] += move.steps
		case south:
			here[1] -= move.steps
		case west:
			here[0] -= move.steps
		default:
			panic(fmt.Sprintf("unknown dir %q", dir))
		}
	}

	return here[0], here[1]
}

type move struct {
	turnDir byte
	steps   int
}

func parseInput(in string) []move {
	moves := make([]move, 0, 256)

	moveStrs := strings.Split(in, ", ")

	for i, moveStr := range moveStrs {
		steps, err := strconv.Atoi(moveStr[1:])
		if err != nil {
			log.Fatalf("error parsing number of steps from move %s (line %d)", moveStr, i+1)
		}

		m := move{
			turnDir: moveStr[0],
			steps:   steps,
		}

		moves = append(moves, m)
	}

	return moves
}
