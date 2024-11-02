package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/vyevs/vtools"
)

func main() {
	defer vtools.TimeIt(time.Now(), "everything")

	lines, err := vtools.ReadLines(os.Args[1])
	if err != nil {
		panic(err)
	}
	{
		seeds, mappings := parseInput1(lines)

		low := findLowestLocation(seeds, mappings)
		fmt.Printf("part 1: the lowest location is %d\n", low)
	}
	{
		seeds, mappings := parseInput2(lines)

		low := findLowestLocation2(seeds, mappings)
		fmt.Printf("part 2: the lowest location is %d\n", low)
	}
}

func findLowestLocation(seeds []uint64, mappings []mapping) uint64 {
	low := uint64(math.MaxUint64)
	for _, seed := range seeds {
		loc := getLocation(seed, mappings)
		if loc < low {
			low = loc
		}
	}
	return low
}

func findLowestLocation2(seeds []seed, mappings []mapping) uint64 {
	// Look through all possible location values, beginning at 0.
	// The first one that has a corresponding initial seed is our answer.
	for loc := uint64(0); loc < (1 << 36); loc++ {
		seed := getSeed(loc, mappings)

		if isSeed(seeds, seed) {
			return loc
		}
	}

	panic("no location found")
}

// isSeed returns whether v is in seeds.
func isSeed(seeds []seed, v uint64) bool {
	for _, s := range seeds {
		lo, hi := s.start, s.start+s.length
		if v >= lo && v < hi {
			return true
		}
	}

	return false
}

// getSeed returns the seed value corresponding to loc given the mappings.
func getSeed(loc uint64, mappings []mapping) uint64 {
	for _, m := range slices.Backward(mappings) {
		loc = m.applyReverse(loc)
	}
	return loc
}

// getLocation returns the location value corresponding to the provided seed given the mappings.
func getLocation(seed uint64, mappings []mapping) uint64 {
	for _, m := range mappings {
		seed = m.apply(seed)
	}

	return seed
}

type mapping struct {
	ranges []myRange
}

// apply takes src value v and turns it into a dest value using m.
func (m mapping) apply(v uint64) uint64 {
	for _, r := range m.ranges {
		lo, hi := r.srcStart, r.srcStart+r.length
		if v >= lo && v < hi {
			return r.destStart + (v - r.srcStart)
		}
	}

	return v
}

// apply takes dest value v and turns it into a src value using m.
func (m mapping) applyReverse(v uint64) uint64 {
	for _, r := range m.ranges {
		lo, hi := r.destStart, r.destStart+r.length
		if v >= lo && v < hi {
			return r.srcStart + (v - r.destStart)
		}
	}
	return v
}

type myRange struct {
	srcStart, destStart, length uint64
}

type seed struct {
	start, length uint64
}

func parseInput1(lines []string) ([]uint64, []mapping) {
	seeds := parseIntSeeds(lines[0])
	mappings := parseMappings(lines[2:])

	return seeds, mappings
}

func parseInput2(lines []string) ([]seed, []mapping) {
	seeds := parseSeeds(lines[0])
	mappings := parseMappings(lines[2:])
	return seeds, mappings
}

func parseSeeds(line string) []seed {
	parts := strings.Fields(line)
	seeds := make([]seed, 0, len(parts)/2)
	for i := 1; i < len(parts); i += 2 {
		seeds = append(seeds, seed{
			start:  atoiu64OrPanic(parts[i]),
			length: atoiu64OrPanic(parts[i+1]),
		})
	}

	return seeds
}

func parseMappings(lines []string) []mapping {
	mappings := make([]mapping, 0, 64)
	var curMapping mapping
	for _, line := range lines {
		if strings.Contains(line, "map") {
			mappings = append(mappings, curMapping)
			curMapping.ranges = nil
			continue
		}

		r := parseRange(line)
		curMapping.ranges = append(curMapping.ranges, r)
	}

	mappings = append(mappings, curMapping)
	return mappings
}

func parseIntSeeds(line string) []uint64 {
	parts := strings.Fields(line)
	seeds := make([]uint64, 0, len(parts)-1)
	for _, p := range parts[1:] {
		seed := atoiu64OrPanic(p)
		seeds = append(seeds, seed)
	}
	return seeds
}

func parseRange(line string) myRange {
	parts := strings.Fields(line)
	return myRange{
		destStart: atoiu64OrPanic(parts[0]),
		srcStart:  atoiu64OrPanic(parts[1]),
		length:    atoiu64OrPanic(parts[2]),
	}
}

func atoiu64OrPanic(s string) uint64 {
	v, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		panic(err)
	}
	return v
}
