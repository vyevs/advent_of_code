package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"slices"
	"strconv"

	"github.com/vyevs/vtools"
)

type pair struct {
	person string
	nextTo string
}

func main() {
	lines, err := vtools.ReadLines(os.Args[1])
	if err != nil {
		log.Fatalf("ReadLines: %v", err)
	}

	m := make(map[pair]int, 1024)
	people := make([]string, 0, 1024)

	for _, line := range lines {
		person, amt, nextTo := parseLine(line)

		pair := pair{person: person, nextTo: nextTo}
		m[pair] = amt

		if !slices.Contains(people, person) {
			people = append(people, person)
		}
		if !slices.Contains(people, nextTo) {
			people = append(people, nextTo)
		}
	}

	for _, person := range people {
		m[pair{person: "vlad", nextTo: person}] = 0
		m[pair{person: person, nextTo: "vlad"}] = 0
	}
	people = append(people, "vlad")

	seating := make([]string, 0, 1024)
	seated := make([]bool, len(people))
	var best int

	var fn func()
	fn = func() {
		if len(seating) == len(people) {
			hc := happinessChange(seating, m)
			if hc > best {
				best = hc
			}
			return
		}

		for i, person := range people {
			if !seated[i] {
				seating = append(seating, person)
				seated[i] = true
				fn()
				seating = seating[:len(seating)-1]
				seated[i] = false
			}
		}
	}

	fn()

	fmt.Println(best)
}

func happinessChange(seating []string, gain map[pair]int) int {
	var hc int

	for i := range len(seating) - 1 {
		hc += gain[pair{person: seating[i], nextTo: seating[i+1]}]
		hc += gain[pair{person: seating[i+1], nextTo: seating[i]}]
	}

	hc += gain[pair{person: seating[0], nextTo: seating[len(seating)-1]}]
	hc += gain[pair{person: seating[len(seating)-1], nextTo: seating[0]}]

	return hc
}

var re = regexp.MustCompile(`^(\w+) would (gain|lose) (\d+) happiness units by sitting next to (\w+)\.$`)

func parseLine(line string) (string, int, string) {
	captures := re.FindStringSubmatch(line)

	p1 := captures[1]

	amtStr := captures[3]
	amt, _ := strconv.Atoi(amtStr)

	if captures[2] == "lose" {
		amt = -amt
	}

	p2 := captures[4]

	return p1, amt, p2
}
