package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type pair struct {
	person string
	nextTo string
}

func main() {
	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatalf("Open: %v", err)
	}

	scanner := bufio.NewScanner(f)

	m := make(map[pair]int, 1024)
	people := make([]string, 0, 1024)

	for i := 1; scanner.Scan(); i++ {
		line := scanner.Text()

		tokens := strings.Split(line, " ")

		person := tokens[0]
		nextTo := tokens[10]
		nextTo = nextTo[:len(nextTo)-1]

		amt, err := strconv.Atoi(tokens[3])
		if err != nil {
			log.Fatalf("error parsing int on line %d", i)
		}

		pair := pair{person: person, nextTo: nextTo}
		if tokens[2] == "gain" {
			m[pair] = amt
		} else {
			m[pair] = -amt
		}

		if !contains(people, person) {
			people = append(people, person)
		}
		if !contains(people, nextTo) {
			people = append(people, nextTo)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("scanner: %v", err)
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

func contains(strs []string, target string) bool {
	for _, s := range strs {
		if s == target {
			return true
		}
	}
	return false
}

func happinessChange(seating []string, gain map[pair]int) int {
	var hc int

	for i := 0; i < len(seating)-1; i++ {
		hc += gain[pair{person: seating[i], nextTo: seating[i+1]}]
		hc += gain[pair{person: seating[i+1], nextTo: seating[i]}]
	}

	hc += gain[pair{person: seating[0], nextTo: seating[len(seating)-1]}]
	hc += gain[pair{person: seating[len(seating)-1], nextTo: seating[0]}]

	return hc
}
