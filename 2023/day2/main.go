package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"

	"github.com/vyevs/vtools"
)

// reveal contains the number of different cubes revealed at any one time.
type reveal struct {
	r, g, b int
}

type game struct {
	id      int
	reveals []reveal
}

// isPossible returns whether g was possible given the limits rMax, gMax, bMax.
func (g game) possible(rMax, gMax, bMax int) bool {
	for _, it := range g.reveals {
		if it.r > rMax || it.g > gMax || it.b > bMax {
			return false
		}
	}

	return true
}

// minRequired returns the minimum number of red, blue, green cubes required to have played the game.
func (g game) minRequired() (int, int, int) {
	var rMin, gMin, bMin int
	for _, it := range g.reveals {
		if it.r > rMin {
			rMin = it.r
		}
		if it.g > gMin {
			gMin = it.g
		}
		if it.b > bMin {
			bMin = it.b
		}
	}
	return rMin, gMin, bMin
}

func main() {
	lines, err := vtools.ReadLines(os.Args[1])
	if err != nil {
		panic(err)
	}
	_ = lines

	games := vtools.MapSlice(lines, parseGame)

	{
		const rMax, gMax, bMax = 12, 13, 14

		var idSum int
		for _, it := range games {
			if it.possible(rMax, gMax, bMax) {
				idSum += it.id
			}
		}
		fmt.Printf("part 1: the sum of IDs of possible games is %d\n", idSum)
	}

	{
		var sum int
		for _, it := range games {
			rMin, gMin, bMin := it.minRequired()
			p := rMin * gMin * bMin
			sum += p
		}
		fmt.Printf("part 2: the sum of the power of the sets is %d\n", sum)
	}
}

var re1 = regexp.MustCompile(`^Game (\d+): (.*)$`)
var re2 = regexp.MustCompile(`(\d+) (red|blue|green)(,|;)? ?`)

func parseGame(s string) game {
	var g game

	groups1 := re1.FindStringSubmatch(s)

	{
		idStr := groups1[1]
		id, err := strconv.Atoi(idStr)
		if err != nil {
			panic(err)
		}
		g.id = id
	}
	{
		revealsStr := groups1[2]
		groups2 := re2.FindAllStringSubmatch(revealsStr, -1)

		rs := make([]reveal, 0, 8)
		var cur reveal
		for _, group := range groups2 {
			amtStr := group[1]
			amt, err := strconv.Atoi(amtStr)
			if err != nil {
				panic(err)
			}

			switch group[2] {
			case "red":
				cur.r = amt
			case "green":
				cur.g = amt
			case "blue":
				cur.b = amt
			}

			if group[3] == ";" || group[3] == "" {
				rs = append(rs, cur)
				cur = reveal{}
			}
		}

		g.reveals = rs
	}

	return g
}
