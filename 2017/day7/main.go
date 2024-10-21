package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/vyevs/vtools"
)

func main() {
	defer vtools.TimeIt(time.Now(), "everything")

	if err := doIt(); err != nil {
		fmt.Printf("error: %v\n", err)
	}
}

func doIt() error {
	programs, err := getPrograms(os.Args[1])
	if err != nil {
		return fmt.Errorf("failed to get programs: %v", err)
	}
	root := getRootProgram(programs)
	fmt.Printf("part 1: the input's root program is %s\n", root.name)
	connectPrograms(programs)

	calculateSums(root)

	fmt.Printf("part 2: the example's root's sum is %d\n", root.sum)

	findImbalance(root)

	return nil
}

func findImbalance(root *program) int {
	childWeights := vtools.Map(root.children, func(p *program) int {
		return p.sum
	})

	
}

func getRootProgram(programs []program) *program {
	children := vtools.NewSet[string](len(programs))

	for _, p := range programs {
		for _, c := range p.childNames {
			children.Add(c)
		}
	}

	// The program that is not a child of any other is the root.
	for i, p := range programs {
		if !children.Contains(p.name) {
			return &programs[i]
		}
	}

	panic("no root found")
}

func getPrograms(fPath string) ([]program, error) {
	lines, err := vtools.ReadLines(fPath)
	if err != nil {
		return nil, fmt.Errorf("ReadLines failed: %v", err)
	}

	return vtools.Map(lines, parseProgramStr), nil
}

func parseProgramStr(pStr string) program {
	var p program

	parts := strings.Split(pStr, " -> ")

	_, _ = fmt.Sscanf(parts[0], "%s (%d)", &p.name, &p.weight)

	if len(parts) > 1 {
		childStr := parts[1]

		p.childNames = strings.Split(childStr, ", ")
	}

	return p
}

func connectPrograms(programs []program) {
	nameToP := make(map[string]*program, len(programs))

	for i, p := range programs {
		nameToP[p.name] = &programs[i]
	}

	for i, p := range programs {
		for _, childName := range p.childNames {
			programs[i].children = append(programs[i].children, nameToP[childName])
		}
	}
}

func calculateSums(root *program) int {
	root.sum = root.weight
	for _, c := range root.children {
		root.sum += calculateSums(c)
	}
	return root.sum
}

type program struct {
	name       string
	weight     int
	childNames []string
	children   []*program
	sum        int
}
