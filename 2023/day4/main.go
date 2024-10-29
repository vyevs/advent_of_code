package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/vyevs/vtools"
)

func main() {
	defer vtools.TimeIt(time.Now(), "everything")

	cards := getCards()
	pts := calculateTotalPoints(cards)
	fmt.Printf("part 1: %d total points\n", pts)

	totalCards := calculateTotalScratchcards(cards)
	fmt.Printf("part 2: %d total scratchcards\n", totalCards)
}

func calculateTotalScratchcards(cards []card) int {
	dp := make([]int, len(cards)) // dp[i] is the count of scratchcard i.
	vtools.SetSliceValues(dp, 1)

	for i, c := range cards {
		nWins := c.countWinners()

		for j := i + 1; j <= i+nWins; j++ {
			dp[j] += dp[i]
		}
	}

	return vtools.SumSlice(dp)
}

func calculateTotalPoints(cards []card) int {
	var total int

	for _, it := range cards {
		total += it.calculatePoints()
	}

	return total
}

func (c card) countWinners() int {
	var n int
	for _, have := range c.have {
		for _, w := range c.winners {
			if w == have {
				n++
			}
		}
	}
	return n
}

func (c card) calculatePoints() int {
	var pts int
	nWins := c.countWinners()
	if nWins > 0 {
		pts = 1
	}
	if nWins > 1 {
		pts <<= nWins - 1
	}

	return pts
}

type card struct {
	winners []int
	have    []int
}

func getCards() []card {
	lines, err := vtools.ReadLines(os.Args[1])
	if err != nil {
		panic(err)
	}

	cards := make([]card, 0, len(lines))
	for _, line := range lines {
		colonI := strings.Index(line, ":")

		rest := line[colonI+2:]

		parts := strings.Split(rest, "|")

		var card card
		{
			winnerStr := parts[0]
			winnerStrs := strings.Fields(winnerStr)
			card.winners = vtools.MapSlice(winnerStrs, atoi)
		}

		{
			haveStr := parts[1]
			haveStrs := strings.Fields(haveStr)
			card.have = vtools.MapSlice(haveStrs, atoi)
		}
		cards = append(cards, card)
	}

	return cards
}

func atoi(s string) int {
	v, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return v
}
