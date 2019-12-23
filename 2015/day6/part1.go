package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("provide 1 input file")
	}

	inFile := os.Args[1]

	f, err := os.Open(inFile)
	if err != nil {
		log.Fatalf("Open: %v", err)
	}

	scanner := bufio.NewScanner(f)

	grid := make([][]int, 1000)
	for i := range grid {
		grid[i] = make([]int, 1000)
	}

	for lineNum := 1; scanner.Scan(); lineNum++ {
		line := scanner.Text()

		tokens := strings.Split(line, " ")

		var x0, y0, x1, y1, op int
		if tokens[0] == "toggle" {
			_, err := fmt.Sscanf(tokens[1], "%d,%d", &x0, &y0)
			if err != nil {
				log.Fatalf("error scanning initial coords on line %d: %v", lineNum, err)
			}

			_, err = fmt.Sscanf(tokens[3], "%d,%d", &x1, &y1)
			if err != nil {
				log.Fatalf("error scanning final coords on line %d: %v", lineNum, err)
			}

		} else if tokens[1] == "on" || tokens[1] == "off" {
			_, err := fmt.Sscanf(tokens[2], "%d,%d", &x0, &y0)
			if err != nil {
				log.Fatalf("error scanning initial coords on line %d: %v", lineNum, err)
			}
			_, err = fmt.Sscanf(tokens[4], "%d,%d", &x1, &y1)
			if err != nil {
				log.Fatalf("error scanning final coords on line %d: %v", lineNum, err)
			}

			if tokens[1] == "on" {
				op = 1
			} else if tokens[1] == "off" {
				op = 2
			}
		}

		for x := x0; x <= x1; x++ {
			for y := y0; y <= y1; y++ {
				if op == 0 {
					grid[x][y] += 2
				} else if op == 1 {
					grid[x][y]++
				} else {
					if grid[x][y] > 0 {
						grid[x][y]--
					}
				}
			}
		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("scanner: %v", err)
	}

	var on int
	for x := 0; x < 1000; x++ {
		for y := 0; y < 1000; y++ {
			on += grid[x][y]
		}
	}

	fmt.Println(on)
}
