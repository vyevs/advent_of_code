package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
	"time"

	"github.com/vyevs/vtools"
)

func main() {
	defer vtools.TimeIt(time.Now(), "everything")

	hands := getHands()

	tw := totalWinnings(hands, false)
	fmt.Printf("part 1: the total winnings are %d\n", tw)

	tw = totalWinnings(hands, true)
	fmt.Printf("part 2: the total winnings are %d\n", tw)
}

func totalWinnings(hands []hand, part2 bool) int {
	for i, h := range hands {
		if part2 {
			hands[i].strength = handStrength2(h.cards)
		} else {
			hands[i].strength = handStrength(h.cards)
		}
	}

	slices.SortFunc(hands, handComparer(part2))

	var tw int
	for i, h := range hands {
		tw += (i + 1) * h.bid
	}
	return tw
}

func handComparer(part2 bool) func(a, b hand) int {
	return func(h1, h2 hand) int {
		str1 := h1.strength
		str2 := h2.strength

		if str1 == str2 {
			for i := range h1.cards {
				c1 := h1.cards[i]
				c2 := h2.cards[i]

				var c1Str, c2Str int
				if part2 {
					c1Str = cardStrength2(c1)
					c2Str = cardStrength2(c2)
				} else {
					c1Str = cardStrength(c1)
					c2Str = cardStrength(c2)
				}

				if c1Str == c2Str {
					continue
				}
				return c1Str - c2Str
			}
		}

		return str1 - str2
	}
}

type hand struct {
	cards    string
	bid      int
	strength int
}

func handStrength(cards string) int {
	var cts [13]int
	for i := range cards {
		c := cards[i]
		cts[cardStrength(c)]++
	}

	if countValue(cts, 5) == 1 {
		return 6
	}
	if countValue(cts, 4) == 1 {
		return 5
	}
	if countValue(cts, 3) == 1 && countValue(cts, 2) == 1 {
		return 4
	}
	if countValue(cts, 3) == 1 {
		return 3
	}
	if countValue(cts, 2) == 2 {
		return 2
	}
	if countValue(cts, 2) == 1 {
		return 1
	}

	return 0
}

func handStrength2(cards string) int {
	var cts [13]int
	for i := range cards {
		c := cards[i]
		cts[cardStrength2(c)]++
	}

	jackCount := cts[0]

	const (
		fiveOfAKind  = 6
		fourOfAKind  = 5
		fullHouse    = 4
		threeOfAKind = 3
		twoPairs     = 2
		onePair      = 1
	)

	if countValue(cts, 5) == 1 {
		return fiveOfAKind
	}
	if countValue(cts, 4) == 1 {
		if jackCount == 1 || jackCount == 4 {
			return fiveOfAKind
		}
	}
	if countValue(cts, 3) == 1 && countValue(cts, 2) == 1 {
		if jackCount == 2 || jackCount == 3 {
			return fiveOfAKind
		}
	}
	if countValue(cts, 3) == 1 {
		if jackCount == 1 || jackCount == 3 {
			return fourOfAKind
		}
	}
	if countValue(cts, 2) == 2 {
		if jackCount == 2 {
			return fourOfAKind
		}
		if jackCount == 1 {
			return fullHouse
		}
	}
	if countValue(cts, 2) == 1 {
		if jackCount == 1 || jackCount == 2 {
			return threeOfAKind
		}
	}

	if countValue(cts, 4) == 1 {
		return fourOfAKind
	}
	if countValue(cts, 3) == 1 && countValue(cts, 2) == 1 {
		return fullHouse
	}
	if countValue(cts, 3) == 1 {
		return threeOfAKind
	}
	if countValue(cts, 2) == 2 {
		return twoPairs
	}
	if countValue(cts, 2) == 1 {
		return onePair
	}

	if jackCount == 1 {
		return onePair
	}

	return 0
}

func cardStrength(c byte) int {
	switch c {
	case 'A':
		return 12
	case 'K':
		return 11
	case 'Q':
		return 10
	case 'J':
		return 9
	case 'T':
		return 8
	default:
		return int(c - '2')
	}
}

func cardStrength2(c byte) int {
	switch c {
	case 'A':
		return 12
	case 'K':
		return 11
	case 'Q':
		return 10
	case 'J':
		return 0
	case 'T':
		return 9
	default:
		return int(c - '1')
	}
}

func countValue(cts [13]int, target int) int {
	var ct int
	for _, v := range cts {
		if v == target {
			ct++
		}
	}
	return ct
}

func getHands() []hand {
	lines, err := vtools.ReadLines(os.Args[1])
	if err != nil {
		panic(err)
	}

	hands := make([]hand, 0, len(lines))
	for _, line := range lines {
		parts := strings.Fields(line)
		h := hand{
			cards: parts[0],
			bid:   vtools.AtoiOrPanic(parts[1]),
		}
		hands = append(hands, h)
	}

	return hands
}
