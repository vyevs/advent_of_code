package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
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
	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatalf("Open: %v", err)
	}

	scanner := bufio.NewScanner(f)

	cities := make([]string, 0, 64)
	distances := distanceMap(make(map[path]int, 1024))

	for i := 1; scanner.Scan(); i++ {
		line := scanner.Text()

		tokens := strings.Split(line, " ")

		city1 := tokens[0]
		city2 := tokens[2]
		distStr := tokens[4]

		dist, err := strconv.Atoi(distStr)
		if err != nil {
			log.Fatalf("error parsing distance on line %d: %v", i, err)
		}

		if !contains(cities, city1) {
			cities = append(cities, city1)
		}
		if !contains(cities, city2) {
			cities = append(cities, city2)
		}

		distances.addDistance(city1, city2, dist)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("scanner: %v", err)
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
	for i := 0; i < len(visitOrder)-1; i++ {
		dist += distances.distance(visitOrder[i], visitOrder[i+1])
	}
	return dist
}

func contains(strs []string, s string) bool {
	for _, str := range strs {
		if str == s {
			return true
		}
	}
	return false
}
