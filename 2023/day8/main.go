package main

import (
	"fmt"
	"os"
	"time"

	"github.com/vyevs/vtools"
)

func main() {
	defer vtools.TimeIt(time.Now(), "everything")

	instructions, nodes := getInput()

	nodeMap := make(map[string]node, len(nodes))
	for _, n := range nodes {
		nodeMap[n.v] = n
	}

	if true {
		nSteps := stepsToZZZ(instructions, nodeMap)
		fmt.Printf("part 1: the number of steps to ZZZ is %d\n", nSteps)
	}

	if true {
		nSteps := stepsToAllZ(instructions, nodes, nodeMap)
		fmt.Printf("part 2: the number of steps until we're only on nodes that end with Z is %d\n", nSteps)
	}
}

func stepsToAllZ(instructions string, nodes []node, m map[string]node) int {
	// The total number of steps is the least common multiple of the distances from each start node to a node that ends with Z.

	starters := vtools.FilterSlice(nodes, func(n node) bool { return n.v[2] == 'A' })
	distances := make([]int, len(starters))
	for i, start := range starters {
		dist := stepsToTarget(instructions, start, m, func(s string) bool { return s[2] == 'Z' })
		distances[i] = dist
	}

	return vtools.LCMAll(distances...)
}

func stepsToTarget(instructions string, start node, m map[string]node, reached func(s string) bool) int {
	var steps int

	cur := start
	for c := range vtools.Cycle([]byte(instructions)) {
		steps++

		if c == 'L' {
			cur = m[cur.left]
		} else {
			cur = m[cur.right]
		}

		if reached(cur.v) {
			break
		}
	}

	return steps
}

func stepsToZZZ(instructions string, m map[string]node) int {
	return stepsToTarget(instructions, m["AAA"], m, func(s string) bool { return s == "ZZZ" })
}

type node struct {
	v, left, right string
}

func getInput() (string, []node) {
	lines, err := vtools.ReadLines(os.Args[1])
	if err != nil {
		panic(err)
	}

	nodes := make([]node, 0, len(lines)-1)
	for _, line := range lines[1:] {
		var n node
		_, _ = fmt.Sscanf(line, "%3s = (%3s, %3s)", &n.v, &n.left, &n.right)
		nodes = append(nodes, n)
	}

	return lines[0], nodes
}
