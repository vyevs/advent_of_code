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

func matches(sue, target sue) bool {
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
	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatalf("open: %v", err)
	}

	sues := parseInput(f)
	f.Close()

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
		if matches(sue, target) {
			fmt.Printf("it's Sue #%d\n", i+1)
		}
	}
}

func parseInput(r io.Reader) []sue {
	scanner := bufio.NewScanner(r)

	sues := make([]sue, 0, 1024)

	for scanner.Scan() {
		line := scanner.Text()

		tokens := strings.Split(line, " ")

		sue := sue{
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

		setField(tokens[2], tokens[3], &sue)
		setField(tokens[4], tokens[5], &sue)
		setField(tokens[6], tokens[7], &sue)

		sues = append(sues, sue)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("scanner: %v", err)
	}

	return sues
}

func setField(field, value string, sue *sue) {
	if value[len(value)-1] == ',' {
		value = value[:len(value)-1]
	}

	v, err := strconv.Atoi(value)
	if err != nil {
		log.Fatalf("atoi: %v", err)
	}

	switch field {
	case "children:":
		sue.children = v
	case "cats:":
		sue.cats = v
	case "samoyeds:":
		sue.samoyeds = v
	case "pomeranians:":
		sue.pomeranians = v
	case "akitas:":
		sue.akitas = v
	case "vizslas:":
		sue.vizslas = v
	case "goldfish:":
		sue.goldfish = v
	case "trees:":
		sue.trees = v
	case "cars:":
		sue.cars = v
	case "perfumes:":
		sue.perfumes = v
	}
}
