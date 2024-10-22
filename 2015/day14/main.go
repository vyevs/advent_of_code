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

func main() {
	part1()
	part2()
}

type reindeer struct {
	name           string
	speed          int
	travelDuration int
	restDuration   int
}

func (r reindeer) String() string {
	return fmt.Sprintf("%s travels at %d for %ds, rests for %ds", r.name, r.speed, r.travelDuration, r.restDuration)
}

func part1() {
	lines, err := vtools.ReadLines(os.Args[1])
	if err != nil {
		log.Fatalf("ReadLine: %v", err)
	}

	rs := parseLines(lines)

	travel := simulateTravel1(rs, 2503)

	for _, result := range travel {
		fmt.Printf("%s travelled %d\n", result.reindeerName, result.distance)
	}

	best := travel[0]
	for i := 1; i < len(travel); i++ {
		if travel[i].distance > best.distance {
			best = travel[i]
		}
	}

	fmt.Printf("\n%s travelled the furthest %d\n", best.reindeerName, best.distance)
}

func simulateTravel2(rs []reindeer, time int) []travelResult {
	travelling := make([]bool, 0, len(rs))
	travelled := make([]int, len(rs))
	actionTimes := make([]int, len(rs))

	for range rs {
		travelling = append(travelling, true)
	}

	for range time {
		for i := range len(rs) {
			if travelling[i] {
				if actionTimes[i] == rs[i].travelDuration {
					travelling[i] = false
					actionTimes[i] = 1
				} else {
					travelled[i] += rs[i].speed
					actionTimes[i]++
				}
			} else {
				if actionTimes[i] == rs[i].restDuration {
					travelling[i] = true
					actionTimes[i] = 1
					travelled[i] += rs[i].speed
				} else {
					actionTimes[i]++
				}
			}
		}
	}

	results := make([]travelResult, 0, len(travelled))
	for i, t := range travelled {
		result := travelResult{reindeerName: rs[i].name, distance: t}
		results = append(results, result)
	}

	return results
}

func part2() {
	lines, err := vtools.ReadLines(os.Args[1])
	if err != nil {
		log.Fatalf("ReadLine: %v", err)
	}

	rs := parseLines(lines)

	travel := simulateTravel2(rs, 2503)

	for _, result := range travel {
		fmt.Printf("%s travelled %d (%d points)\n", result.reindeerName, result.distance, result.points)
	}

	bestDist := travel[0]
	bestPts := travel[0]
	for i := 1; i < len(travel); i++ {
		if travel[i].distance > bestDist.distance {
			bestDist = travel[i]
		}
		if travel[i].points > bestPts.points {
			bestPts = travel[i]
		}
	}

	fmt.Printf("\n%s travelled the furthest %d\n", bestDist.reindeerName, bestDist.distance)
	fmt.Printf("%s got the most points %d\n", bestPts.reindeerName, bestPts.points)
}

type travelResult struct {
	reindeerName string
	distance     int
	points       int
}

func simulateTravel1(rs []reindeer, time int) []travelResult {
	travelling := make([]bool, 0, len(rs))
	travelled := make([]int, len(rs))
	actionTimes := make([]int, len(rs))
	points := make([]int, len(rs))

	for range len(rs) {
		travelling = append(travelling, true)
	}

	for range time {
		for i := range len(rs) {
			if travelling[i] {
				if actionTimes[i] == rs[i].travelDuration {
					travelling[i] = false
					actionTimes[i] = 1
				} else {
					travelled[i] += rs[i].speed
					actionTimes[i]++
				}
			} else {
				if actionTimes[i] == rs[i].restDuration {
					travelling[i] = true
					actionTimes[i] = 1
					travelled[i] += rs[i].speed
				} else {
					actionTimes[i]++
				}
			}
		}

		bestSoFar := slices.Max(travelled)
		for i := range rs {
			if travelled[i] == bestSoFar {
				points[i]++
			}
		}
	}

	results := make([]travelResult, 0, len(travelled))
	for i, t := range travelled {
		result := travelResult{reindeerName: rs[i].name, distance: t, points: points[i]}
		results = append(results, result)
	}

	return results
}

var re = regexp.MustCompile(`^(\w+) can fly (\d+) km/s for (\d+) seconds, but then must rest for (\d+) seconds\.$`)

func parseLine(line string) reindeer {
	groups := re.FindStringSubmatch(line)

	r := reindeer{
		name: groups[1],
	}
	r.speed, _ = strconv.Atoi(groups[2])
	r.travelDuration, _ = strconv.Atoi(groups[3])
	r.restDuration, _ = strconv.Atoi(groups[4])

	return r
}

func parseLines(lines []string) []reindeer {
	return vtools.Map(lines, parseLine)
}
