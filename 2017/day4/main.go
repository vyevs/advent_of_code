package main

import (
	"fmt"
	"maps"
	"strings"
	"time"

	"github.com/vyevs/vtools"
)

func main() {
	defer vtools.TimeIt(time.Now(), "everything")

	if err := doIt(); err != nil {
		fmt.Printf("error: %v", err)
	}
}

func doIt() error {
	{
		fmt.Println("part 1 examples:")
		examples := []string{"aa bb cc dd ee", "aa bb cc dd aa", "aa bb cc dd aaa"}
		for _, pp := range examples {
			valid := !containsDuplicateWords(pp)
			validStr := validStr(valid)
			fmt.Printf("\t%q is %s\n", pp, validStr)
		}
	}

	passphrases, err := vtools.ReadLines("input.txt")
	if err != nil {
		return fmt.Errorf("failed to read file lines: %v", err)
	}

	{
		fmt.Println("part 1:")
		ctValid := vtools.CountFunc(passphrases, func(s string) bool {
			return !containsDuplicateWords(s)
		})
		fmt.Printf("\tthere are %d valid passphrases\n", ctValid)
	}

	{
		fmt.Println("part 2 examples:")
		examples := []string{"abcde fghij", "abcde xyz ecdab", "a ab abc abd abf abj", "iiii oiii ooii oooi oooo", "oiii ioii iioi iiio"}
		for _, pp := range examples {
			valid := !containsAnagrams(pp)
			validStr := validStr(valid)
			fmt.Printf("\t%q is %s\n", pp, validStr)
		}
	}

	{
		fmt.Println("part 2:")
		ctValid := vtools.CountFunc(passphrases, func(s string) bool {
			return !containsAnagrams(s)
		})
		fmt.Printf("\t%d passphrases are valid\n", ctValid)
	}

	return nil
}

func validStr(valid bool) string {
	if valid {
		return "valid"
	}
	return "invalid"
}

func containsDuplicateWords(s string) bool {
	parts := strings.Split(s, " ")

	cts := vtools.CounterSlice(parts) // Count number of occurrences of each string.
	values := maps.Values(cts)        // Take only the counts.

	// If a string appeared more than once in s, then s contains duplicate words.
	return vtools.Any(values, func(i int) bool { return i > 1 })
}

func containsAnagrams(s string) bool {
	parts := strings.Split(s, " ")

	for i, p1 := range parts {
		for j, p2 := range parts {
			if i == j {
				continue
			}

			if areAnagrams(p1, p2) {
				return true
			}
		}
	}

	return false
}

func areAnagrams(s1, s2 string) bool {
	cts1 := vtools.Counter(vtools.StrBytes(s1))
	cts2 := vtools.Counter(vtools.StrBytes(s2))

	return maps.Equal(cts1, cts2)
}
