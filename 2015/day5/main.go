package main

import (
	"fmt"
	"log"
	"os"

	"github.com/vyevs/vtools"
)

func main() {
	lines, err := vtools.ReadLines(os.Args[1])
	if err != nil {
		log.Fatalf("ReadLine: %v", err)
	}

	{
		niceCt := vtools.CountFunc(lines, niceString)
		fmt.Printf("part 1: %d nice strings\n", niceCt)
	}

	{
		niceCt := vtools.CountFunc(lines, niceStringV2)
		fmt.Printf("part 2: %d nice strings\n", niceCt)
	}

}

func niceString(s string) bool {
	var vowelCount int
	var lastRune rune
	var seen2InARow bool

	for _, r := range s {
		if r == lastRune {
			seen2InARow = true
		}

		switch r {
		case 'a', 'e', 'i', 'o', 'u':
			vowelCount++

		case 'b', 'd', 'q', 'y':
			if lastRune == r-1 {
				return false
			}
		}

		lastRune = r
	}

	return seen2InARow && vowelCount >= 3
}

func oneRepeatWithBetween(s string) bool {
	for i := range len(s) - 2 {
		if s[i] == s[i+2] {
			return true
		}
	}
	return false
}

func hasPairThatAppearsTwice(s string) bool {
	pairs := make(map[string]int, 26*26)

	pairs[s[:2]]++

	for i := 1; i < len(s)-1; i++ {
		if s[i-1] != s[i] || s[i] != s[i+1] {
			pairs[s[i:i+2]]++
		}
	}

	for _, v := range pairs {
		if v >= 2 {
			return true
		}
	}

	return false
}

func niceStringV2(s string) bool {
	return oneRepeatWithBetween(s) && hasPairThatAppearsTwice(s)
}
