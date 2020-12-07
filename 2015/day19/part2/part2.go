package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	moleculeMappings, target := readInput()

	var recurse func(current string, nSteps uint64)

	var bestSteps uint64 = ^uint64(0)

	recurse = func(current string, nSteps uint64) {

		if current == "e" {
			if nSteps < bestSteps {
				bestSteps = nSteps
				fmt.Printf("found a way to get to the target in %d steps\n", nSteps)
			}

		} else {
			for from, tos := range moleculeMappings {
				for _, to := range tos {
					if strings.Contains(current, to) {
						newCurrent := strings.Replace(current, to, from, 1)

						recurse(newCurrent, nSteps+1)
					}
				}
			}
		}

	}

	recurse(target, 0)

	fmt.Printf("least steps is %d\n", bestSteps)
}

func readInput() (map[string][]string, string) {
	scanner := bufio.NewScanner(os.Stdin)

	moleculeMappings := make(map[string][]string, 128)

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			break
		}

		tokens := strings.Split(line, " ")

		from := tokens[0]
		to := tokens[2]

		moleculeMappings[from] = append(moleculeMappings[from], to)
	}

	scanner.Scan()
	target := scanner.Text()

	return moleculeMappings, target
}
