package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	steps := getSteps()
	
	{
		prod := solvePart1(steps)
		fmt.Printf("1: The final depth * final horizontal position = %d\n", prod)
	}
	{
		prod := solvePart2(steps)
		fmt.Printf("2: The final depth * final horizontal position = %d\n", prod)
	}
}

func solvePart1(steps []step) int {
	var depth, horizontalPos int
	
	for _, s := range steps {
		if s.dir == "forward" {
			horizontalPos += s.mag
		} else if s.dir == "down" {
			depth += s.mag
		} else {
			depth -= s.mag
		}
	}
	
	return depth * horizontalPos
}

func solvePart2(steps []step) int {
	var depth, horizontalPos, aim int
	
	for _, s := range steps {
		if s.dir == "forward" {
			horizontalPos += s.mag
			depth += s.mag * aim
		} else if s.dir == "down" {
			aim += s.mag
		} else {
			aim -= s.mag
		}
	}
	
	return depth * horizontalPos
}

type step struct {
	dir string
	mag int
}

func getSteps() []step {
	scanner := bufio.NewScanner(os.Stdin)
	
	steps := make([]step, 0, 1024)
	for scanner.Scan() {
		line := scanner.Text()
		
		parts := strings.Split(line, " ")
		
		mag, _ := strconv.Atoi(parts[1])
		step := step {
			dir: parts[0],
			mag: mag,
		}
		steps = append(steps, step)
	}
	
	return steps
}