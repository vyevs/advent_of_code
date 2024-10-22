package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/vyevs/vtools"
)

type distanceMap map[path]int

func (m distanceMap) addDistance(from, to string, d int) {
	m[path{from: from, to: to}] = d
	m[path{from: to, to: from}] = d
}

func (m distanceMap) distance(from, to string) int {
	return m[path{from: from, to: to}]
}

type path struct {
	from string
	to   string
}

func main() {
	lines, err := vtools.ReadLines(os.Args[1])
	if err != nil {
		log.Fatalf("ReadLines: %v", err)
	}

	cities := make([]string, 0, 64)
	distances := distanceMap(make(map[path]int, 1024))

	for i, line := range lines {

		tokens := strings.Split(line, " ")

		city1 := tokens[0]
		city2 := tokens[2]
		distStr := tokens[4]

		dist, err := strconv.Atoi(distStr)
		if err != nil {
			log.Fatalf("error parsing distance on line %d: %v", i, err)
		}

		if !slices.Contains(cities, city1) {
			cities = append(cities, city1)
		}
		if !slices.Contains(cities, city2) {
			cities = append(cities, city2)
		}

		distances.addDistance(city1, city2, dist)
	}

	bestTraversalDist := travellingSalesman(cities, distances)

	fmt.Println(bestTraversalDist)
}

func travellingSalesman(cities []string, distances distanceMap) int {
	visitOrder := make([]string, 0, len(cities))
	visited := make([]bool, len(cities)) // whether cities[i] has been visited
	var best int

	var fn func()
	fn = func() {
		if len(visitOrder) == len(cities) {
			dist := traversalDistance(visitOrder, distances)
			if dist > best {
				best = dist
			}
			return
		}

		for i, city := range cities {
			if !visited[i] {
				visitOrder = append(visitOrder, city)
				visited[i] = true
				fn()
				visited[i] = false
				visitOrder = visitOrder[:len(visitOrder)-1]
			}
		}
	}

	fn()

	return best
}

func traversalDistance(visitOrder []string, distances distanceMap) int {
	var dist int
	for i := range len(visitOrder) - 1 {
		dist += distances.distance(visitOrder[i], visitOrder[i+1])
	}
	return dist
}
