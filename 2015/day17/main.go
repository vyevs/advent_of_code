package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("provide one input file")
	}

	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatalf("open: %v", err)
	}

	target, containers := parseInput(f)

	min := minNumberContainers(target, containers)

	numCombosUsingMin := numCombinationsUsingMinContainers(target, min, containers)

	fmt.Println(min)
	fmt.Println(numCombosUsingMin)
}

func countTrues(bs []bool) int {
	var c int
	for _, b := range bs {
		if b {
			c++
		}
	}
	return c
}

func minNumberContainers(target int, containers []int) int {
	used := make([]bool, len(containers))

	var sum int
	var low int
	min := -1

	var fn func()

	fn = func() {
		if sum > target {
			return
		}

		if sum == target {
			c := countTrues(used)
			if min == -1 || c < min {
				min = c
			}
			return
		}

		for i := low; i < len(containers); i++ {
			if !used[i] {
				c := containers[i]
				used[i] = true
				sum += c
				oldLow := low
				low = i + 1
				fn()
				low = oldLow
				sum -= c
				used[i] = false
			}
		}
	}

	fn()

	return min
}

func numCombinationsUsingMinContainers(target, minContainers int, containers []int) int {
	used := make([]bool, len(containers))

	var sum int
	var low int
	var count int

	var fn func()

	fn = func() {
		if sum > target {
			return
		}

		if sum == target && countTrues(used) == minContainers {
			count++
			return
		}

		for i := low; i < len(containers); i++ {
			if !used[i] {
				c := containers[i]
				used[i] = true
				sum += c
				oldLow := low
				low = i + 1
				fn()
				low = oldLow
				sum -= c
				used[i] = false
			}
		}
	}

	fn()

	return count
}

func parseInput(r io.Reader) (int, []int) {
	scanner := bufio.NewScanner(r)

	containers := make([]int, 0, 128)

	scanner.Scan()
	target, err := strconv.Atoi(scanner.Text())
	if err != nil {
		log.Fatalf("error parsing target: %v", err)
	}

	for lineNum := 1; scanner.Scan(); lineNum++ {
		line := scanner.Text()

		container, err := strconv.Atoi(line)
		if err != nil {
			log.Fatalf("error parsing container on line %d", lineNum)
		}

		containers = append(containers, container)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("scanner: %v", err)
	}

	return target, containers
}
