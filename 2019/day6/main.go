package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func timeIt(start time.Time, msg string) {
	log.Printf("%s (%s)", msg, time.Since(start))
}

func main() {
	defer timeIt(time.Now(), "done")
	orbits := readInput()
	doPart1(orbits)
	doPart2(orbits)
}

type orbit struct {
	orbiter string
	orbitee string
}

func readInput() []orbit {
	defer timeIt(time.Now(), "read input")

	scanner := bufio.NewScanner(os.Stdin)

	orbits := make([]orbit, 0, 1024)

	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.Split(line, ")")

		orbits = append(orbits, orbit{orbitee: parts[0], orbiter: parts[1]})
	}

	return orbits
}

func doPart1(orbits []orbit) {
	orbiters := makeOrbiters(orbits)

	fmt.Println(orbiters)

	numOrbits := countOrbits(orbiters)

	fmt.Printf("total orbits: %d\n", numOrbits)
}

func countOrbits(orbiters map[string][]string) int {

	//visited := make(map[string]struct{}, len(orbiters))

	var recurse func(orbitee string, depth int) int

	recurse = func(orbitee string, depth int) int {
		/*if _, alreadyVisited := visited[orbitee]; alreadyVisited {
			return 0
		}*/
		var sumOfOrbiters int
		currentOrbiters := orbiters[orbitee]
		for _, orbiter := range currentOrbiters {
			sumOfOrbiters += recurse(orbiter, depth+1)
		}

		return depth + sumOfOrbiters
	}

	return recurse("COM", 0)
}

// creates a mapping of objects to the objects that orbit them
func makeOrbiters(orbits []orbit) map[string][]string {
	orbiters := make(map[string][]string, len(orbits))
	for _, orbit := range orbits {
		currentOrbiters := orbiters[orbit.orbitee]
		orbiters[orbit.orbitee] = append(currentOrbiters, orbit.orbiter)
	}
	return orbiters
}

// creates a mapping of objects to the object that it directly orbits
func makeOrbitees(orbits []orbit) map[string]string {
	orbitees := make(map[string]string, len(orbits))
	for _, orbit := range orbits {
		orbitees[orbit.orbiter] = orbit.orbitee
	}
	return orbitees
}

func doPart2(orbits []orbit) {
	orbiters := makeOrbiters(orbits)
	orbitees := makeOrbitees(orbits)

	visited := make(map[string]struct{}, len(orbits))

	haveVisited := func(obj string) bool {
		_, hv := visited[obj]
		return hv
	}

	var recurse func(orbitee string, depth int) int

	recurse = func(current string, depth int) int {
		if haveVisited(current) {
			return -1
		}
		visited[current] = struct{}{}

		orbitersOfCurrent := orbiters[current]

		for _, orbiter := range orbitersOfCurrent {
			if orbiter == "SAN" {
				return depth
			}
		}

		minSteps := -1
		for _, orbiter := range orbitersOfCurrent {
			if d := recurse(orbiter, depth+1); d != -1 && (minSteps == -1 || d < minSteps) {
				minSteps = d
			}
		}
		if minSteps != -1 {
			return minSteps
		}

		if orbitee, haveOrbitee := orbitees[current]; haveOrbitee {
			if d := recurse(orbitee, depth+1); d != -1 {
				return d
			}
		}

		delete(visited, current)

		return -1
	}

	steps := recurse(orbitees["YOU"], 0)

	fmt.Printf("took %d steps to get to SAN\n", steps)
}
