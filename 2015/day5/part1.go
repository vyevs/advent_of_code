package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("give input file")
	}

	inFile := os.Args[1]

	f, err := os.Open(inFile)
	if err != nil {
		log.Fatalf("Open: %v", err)
	}

	scanner := bufio.NewScanner(f)

	var niceStrs int

	for scanner.Scan() {
		line := scanner.Text()

		if niceStringV2(line) {
			fmt.Println(line)
			niceStrs++
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("scanner: %v", err)
	}

	fmt.Println(niceStrs)

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

		case 'b':
			if lastRune == 'a' {
				return false
			}

		case 'd':
			if lastRune == 'c' {
				return false
			}

		case 'q':
			if lastRune == 'p' {
				return false
			}

		case 'y':
			if lastRune == 'x' {
				return false
			}

		}

		lastRune = r
	}

	return seen2InARow && vowelCount >= 3
}

func oneRepeatWithBetween(s string) bool {
	for i := 0; i < len(s)-2; i++ {
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
