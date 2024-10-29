package main

import (
	"fmt"
	"math/bits"
	"os"
	"slices"
	"strings"
	"time"

	"github.com/vyevs/vtools"
)

// 0:      1:      2:      3:      4:
// 	aaaa    ....    aaaa    aaaa    ....
// b    c  .    c  .    c  .    c  b    c
// b    c  .    c  .    c  .    c  b    c
// 	....    ....    dddd    dddd    dddd
// e    f  .    f  e    .  .    f  .    f
// e    f  .    f  e    .  .    f  .    f
// 	gggg    ....    gggg    gggg    ....
//
//  5:      6:      7:      8:      9:
// 	aaaa    aaaa    aaaa    aaaa    aaaa
// b    .  b    .  .    c  b    c  b    c
// b    .  b    .  .    c  b    c  b    c
// 	dddd    dddd    ....    dddd    dddd
// .    f  e    f  .    f  e    f  .    f
// .    f  e    f  .    f  e    f  .    f
// 	gggg    gggg    ....    gggg    gggg
//
// bit patterns, where bits are mapped to the segments a through g:
// 0: 1110111
// 1: 0100100
// 2: 1011101
// 3: 1011011
// 4: 0111010
// 5: 1101011
// 6: 1101111
// 7: 1010010
// 8: 1111111
// 9: 1111011
//
// ones counts:
// 1: 2
// 7: 3
// 4: 4
// 2, 3, 5: 5
// 0, 6, 9: 6
// 8: 7

// determineDigitPatterns returns a mapping from digits 0-9 to the bit patterns that represent them in the provided patterns.
// patterns is expected to be sorted by the number of bits in each uint8.
func determineDigitPatterns(patterns [10]uint8) [10]uint8 {
	var digitToPattern [10]uint8    // digitToPattern[i] is the pattern of bits that represent that digit. It is not fully filled in from the start.
	digitToPattern[1] = patterns[0] // 1 has 2 bits on
	digitToPattern[7] = patterns[1] // 7 has 3 bits on
	digitToPattern[4] = patterns[2] // 4 has 4 bits on
	digitToPattern[8] = patterns[9] // 8 has 7 bits on

	// Determine which pattern represents 3. It's the only one with 5 bits on that has all the bits that represent 7.
	{
		p7 := digitToPattern[7]
		for _, p := range patterns[3:6] {
			if bits.OnesCount8(p) == 5 && (p&p7) == p7 {
				digitToPattern[3] = p
				break
			}
		}
	}

	// Determine which pattern represents 9. It's the only one with 6 bits on that has all the bits that represent 4.
	{
		p4 := digitToPattern[4]
		for _, p := range patterns[6:9] {
			if bits.OnesCount8(p) == 6 && (p&p4) == p4 {
				digitToPattern[9] = p
				break
			}
		}
	}

	// Determine which pattern represents 6. It's the only one with 6 bits on that DOES NOT have all the bits that represent 7.
	{
		p7 := digitToPattern[7]
		for _, p := range patterns[6:9] {
			if bits.OnesCount8(p) == 6 && (p&p7) != p7 {
				digitToPattern[6] = p
				break
			}
		}
	}

	// Determine which pattern represents 0. By elimination, its the one with 6 bits on that is not equal to the patterns that represent 6 or 9
	{
		p6 := digitToPattern[6]
		p9 := digitToPattern[9]
		for _, p := range patterns[6:9] {
			if bits.OnesCount8(p) == 6 && p != p6 && p != p9 {
				digitToPattern[0] = p
				break
			}
		}
	}

	// Determine which pattern represents 5. It's the one with 5 bits that has 5 bits in common with the pattern for 9 AND is not 3.
	{
		p9 := digitToPattern[9]
		p3 := digitToPattern[3]
		for _, p := range patterns[3:6] {
			if bits.OnesCount8(p) == 5 && p != p3 && bits.OnesCount8(p&p9) == 5 {
				digitToPattern[5] = p
				break
			}
		}
	}

	// Determine which pattern represents 2. By elimination, it's the pattern with 5 bits on that is not equal to the patterns of 3 or 5.
	{
		p3 := digitToPattern[3]
		p5 := digitToPattern[5]
		for _, p := range patterns[3:6] {
			if bits.OnesCount8(p) == 5 && p != p3 && p != p5 {
				digitToPattern[2] = p
				break
			}
		}
	}

	return digitToPattern
}

func determineMapping(patterns [10]uint8) [255]uint8 {
	digitToPattern := determineDigitPatterns(patterns)

	var patternToDigit [255]uint8
	for d, p := range digitToPattern {
		patternToDigit[p] = uint8(d)
	}
	return patternToDigit
}

type note struct {
	patterns [10]uint8
	output   [4]uint8
}

func (n note) getOutput() int {
	m := determineMapping(n.patterns)

	var digits [4]uint8
	for i, it := range n.output {
		digits[i] = m[it]
	}

	return 1000*int(digits[0]) + 100*int(digits[1]) + 10*int(digits[2]) + int(digits[3])
}

func sumOutputs(ns []note) int {
	var sum int
	for _, n := range ns {
		sum += n.getOutput()
	}
	return sum
}

func main() {
	defer vtools.TimeIt(time.Now(), "everything")

	notes := getInput()

	defer vtools.TimeIt(time.Now(), "solving")

	num1478s := count1478s(notes)
	fmt.Printf("part 1: the numbers 1, 4, 7, 8 appear %d times in the output\n", num1478s)

	sum := sumOutputs(notes)
	fmt.Printf("part 2: the sum of the outputs is %d\n", sum)
}

func count1478s(notes []note) int {
	var ct int

	for _, note := range notes {
		for _, it := range note.output {
			ones := bits.OnesCount8(it)
			if ones == 2 /* 1 */ || ones == 3 /* 7 */ || ones == 4 /* 4 */ || ones == 7 /* 8 */ {
				ct++
			}
		}
	}

	return ct
}

func getInput() []note {
	lines, err := vtools.ReadLines(os.Args[1])
	if err != nil {
		panic(err)
	}

	out := make([]note, 0, len(lines))
	for _, line := range lines {
		n := parseNote(line)
		out = append(out, n)
	}

	return out
}

func strToBits(s string) uint8 {
	var out uint8
	for i := range len(s) {
		b := s[i]

		out |= (1 << (b - 'a'))
	}
	return out
}

func parseNote(s string) note {
	parts := strings.Split(s, " | ")

	patterns := strings.Split(parts[0], " ")
	output := strings.Split(parts[1], " ")
	slices.SortFunc(patterns, func(a, b string) int {
		if len(a) < len(b) {
			return -1
		}
		return 1
	})
	n := note{
		patterns: [10]uint8(vtools.MapSlice(patterns, strToBits)),
		output:   [4]uint8(vtools.MapSlice(output, strToBits)),
	}
	return n
}
