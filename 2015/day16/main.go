package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"

	"github.com/vyevs/vtools"
)

type sue struct {
	children    int
	cats        int
	samoyeds    int
	pomeranians int
	akitas      int
	vizslas     int
	goldfish    int
	trees       int
	cars        int
	perfumes    int
}

func matchesPart1(sue, target sue) bool {
	return (sue.children == -1 || sue.children == target.children) &&
		(sue.cats == -1 || sue.cats == target.cats) &&
		(sue.samoyeds == -1 || sue.samoyeds == target.samoyeds) &&
		(sue.pomeranians == -1 || sue.pomeranians == target.pomeranians) &&
		(sue.akitas == -1 || sue.akitas == target.akitas) &&
		(sue.vizslas == -1 || sue.vizslas == target.vizslas) &&
		(sue.goldfish == -1 || sue.goldfish == target.goldfish) &&
		(sue.trees == -1 || sue.trees == target.trees) &&
		(sue.cars == -1 || sue.cars == target.cars) &&
		(sue.perfumes == -1 || sue.perfumes == target.perfumes)
}

func matchesPart2(sue, target sue) bool {
	return (sue.children == -1 || sue.children == target.children) &&
		(sue.cats == -1 || sue.cats > target.cats) &&
		(sue.samoyeds == -1 || sue.samoyeds == target.samoyeds) &&
		(sue.pomeranians == -1 || sue.pomeranians < target.pomeranians) &&
		(sue.akitas == -1 || sue.akitas == target.akitas) &&
		(sue.vizslas == -1 || sue.vizslas == target.vizslas) &&
		(sue.goldfish == -1 || sue.goldfish < target.goldfish) &&
		(sue.trees == -1 || sue.trees > target.trees) &&
		(sue.cars == -1 || sue.cars == target.cars) &&
		(sue.perfumes == -1 || sue.perfumes == target.perfumes)
}

func main() {
	lines, err := vtools.ReadLines(os.Args[1])
	if err != nil {
		log.Fatalf("ReadLines: %v", err)
	}

	sues := vtools.Map(lines, parseLine)

	target := sue{
		children:    3,
		cats:        7,
		samoyeds:    2,
		pomeranians: 3,
		akitas:      0,
		vizslas:     0,
		goldfish:    5,
		trees:       3,
		cars:        2,
		perfumes:    1,
	}

	for i, sue := range sues {
		if matchesPart1(sue, target) {
			fmt.Printf("part 1: it's Sue #%d\n", i+1)
		}
	}

	for i, sue := range sues {
		if matchesPart2(sue, target) {
			fmt.Printf("part 2: it's Sue #%d\n", i+1)
		}
	}
}

var re = regexp.MustCompile(`^Sue \d+: (\w+): (\d+), (\w+): (\d+), (\w+): (\d+)$`)

func parseLine(line string) sue {
	groups := re.FindStringSubmatch(line)

	s := sue{
		children:    -1,
		cats:        -1,
		samoyeds:    -1,
		pomeranians: -1,
		akitas:      -1,
		vizslas:     -1,
		goldfish:    -1,
		trees:       -1,
		cars:        -1,
		perfumes:    -1,
	}
	setField(&s, groups[1], atoi(groups[2]))
	setField(&s, groups[3], atoi(groups[4]))
	setField(&s, groups[5], atoi(groups[6]))
	return s
}

func setField(s *sue, field string, v int) {
	switch field {
	case "children":
		s.children = v
	case "cats":
		s.cats = v
	case "samoyeds":
		s.samoyeds = v
	case "pomeranians":
		s.pomeranians = v
	case "akitas":
		s.akitas = v
	case "vizslas":
		s.vizslas = v
	case "goldfish":
		s.goldfish = v
	case "trees":
		s.trees = v
	case "cars":
		s.cars = v
	case "perfumes":
		s.perfumes = v
	default:
		panic(field)
	}

}

func atoi(s string) int {
	v, err := strconv.Atoi(s)
	if err != nil {
		panic("strconv.Atoi: %v")
	}
	return v
}
