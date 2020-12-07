package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	input := readInput()

	fmt.Println("Part 1:")
	doPart1(input)
	fmt.Println()

	fmt.Println("Part 2:")
	doPart2(input)

}

func doPart1(groups []group) {
	var sum int

	for _, group := range groups {
		sum += len(group.yesUnion())
	}

	fmt.Printf("A total of %d questions were answered yes\n", sum)
}

func (g group) yesUnion() map[byte]struct{} {
	union := make(map[byte]struct{}, 26)
	for _, person := range g.people {
		for yes := range person {
			union[yes] = struct{}{}
		}
	}
	return union
}

func doPart2(groups []group) {
	var sum int

	for _, group := range groups {
		sum += len(group.yesIntersection())
	}

	fmt.Printf("Sum: %d\n", sum)
}

func (g group) yesIntersection() map[byte]struct{} {
	intersection := make(map[byte]struct{}, 26)

	var charsToCheck map[byte]struct{}
	for _, person := range g.people {
		charsToCheck = person
		break
	}

outer:
	for c := range charsToCheck {
		for _, person := range g.people {
			_, yes := person[byte(c)]
			if !yes {
				continue outer
			}
		}

		intersection[byte(c)] = struct{}{}
	}

	return intersection
}

type group struct {
	people []map[byte]struct{} // people[i] is the set of answers to which person i answered yes
}

func readInput() []group {
	scanner := bufio.NewScanner(os.Stdin)

	groups := make([]group, 0, 1024)
	curGroup := group{
		people: make([]map[byte]struct{}, 0, 128),
	}
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)

		if line == "" {
			groups = append(groups, curGroup)
			curGroup = group{
				people: make([]map[byte]struct{}, 0, 128),
			}

		} else {
			person := make(map[byte]struct{}, 26)
			for i := 0; i < len(line); i++ {
				c := line[i]
				person[c] = struct{}{}
			}
			curGroup.people = append(curGroup.people, person)
		}
	}

	groups = append(groups, curGroup)

	return groups
}
