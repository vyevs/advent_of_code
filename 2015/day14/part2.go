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

type reindeer struct {
	name           string
	speed          int
	travelDuration int
	restDuration   int
}

func (r reindeer) String() string {
	return fmt.Sprintf("%s travels at %d for %ds, rests for %ds", r.name, r.speed, r.travelDuration, r.restDuration)
}

func main() {
	f, _ := os.Open(os.Args[1])

	rs := parseInput(f)
	f.Close()

	travel := simulateTravel(rs, 2503)

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

func simulateTravel(rs []reindeer, time int) []travelResult {
	travelling := make([]bool, 0, 64)
	travelled := make([]int, 0, 64)
	actionTimes := make([]int, 0, 64)
	points := make([]int, 0, 64)

	for range rs {
		travelling = append(travelling, true)
		travelled = append(travelled, 0)
		actionTimes = append(actionTimes, 0)
		points = append(points, 0)
	}

	for i := 0; i < time; i++ {
		for i := 0; i < len(rs); i++ {
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

		bestSoFar := max(travelled)
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

func max(s []int) int {
	max := s[0]
	for i := 1; i < len(s); i++ {
		if s[i] > max {
			max = s[i]
		}
	}
	return max
}

func parseInput(r io.Reader) []reindeer {
	scanner := bufio.NewScanner(r)

	rs := make([]reindeer, 0, 64)

	for scanner.Scan() {
		line := scanner.Text()

		tokens := strings.Split(line, " ")

		var reindeer reindeer

		reindeer.name = tokens[0]
		reindeer.speed, _ = strconv.Atoi(tokens[3])
		reindeer.travelDuration, _ = strconv.Atoi(tokens[6])
		reindeer.restDuration, _ = strconv.Atoi(tokens[13])

		rs = append(rs, reindeer)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("scanner: %v", err)
	}

	return rs
}
