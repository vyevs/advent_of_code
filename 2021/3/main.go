package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := getInput()
	{
		pc := determinePowerConsumption(in)
		fmt.Printf("1: Power consumption is %d\n", pc)
	}
	{
		lsr := determineLifeSupportRating(in)
		fmt.Printf("2: Life support rating is %d\n", lsr)
	}
}

func determinePowerConsumption(in []string) int {
	g := determineGammaRate(in)
	e := determineEpsilonRate(in)
	return e * g
}

func determineLifeSupportRating(in []string) int {
	ogr := determineOxygenGeneratorRating(in)
	co2sr := determineCO2ScrubberRating(in)
	return ogr * co2sr
}

func determineGammaRate(in []string) int {
	count0s := count0sByColumn(in)

	var res int
	for i, ct := range count0s {
		if ct <= len(in)/2 {
			res |= (1 << (len(in[0]) - 1 - i))
		}
	}
	return res
}

func determineEpsilonRate(in []string) int {
	count0s := count0sByColumn(in)

	var res int
	for i, ct := range count0s {
		if ct >= len(in)/2 {
			res |= (1 << (len(in[0]) - 1 - i))
		}
	}
	return res
}

func count0sByColumn(in []string) []int {
	count0s := make([]int, len(in[0]))
	for _, v := range in {
		for i, c := range v {
			if c == '0' {
				count0s[i]++
			}
		}
	}
	return count0s
}

func determineOxygenGeneratorRating(in []string) int {
	set := make(map[string]struct{}, len(in))
	for _, v := range in {
		set[v] = struct{}{}
	}

	var bitPos int
	for len(set) > 1 {
		bit := mostCommonBitAtPosition(set, bitPos)

		for v := range set {
			if v[bitPos] != bit {
				delete(set, v)
			}
		}

		bitPos++
	}

	var lastVal string
	for v := range set {
		lastVal = v
	}
	return binStrToInt(lastVal)
}

func binStrToInt(s string) int {
	var res int
	for i, v := range s {
		if v == '1' {
			res |= (1 << (len(s) - i - 1))
		}
	}
	return res
}

func determineCO2ScrubberRating(in []string) int {
	set := make(map[string]struct{}, len(in))
	for _, v := range in {
		set[v] = struct{}{}
	}

	var bitPos int
	for len(set) > 1 {
		bit := leastCommonBitAtPosition(set, bitPos)

		for v := range set {
			if v[bitPos] != bit {
				delete(set, v)
			}
		}

		bitPos++
	}

	var lastVal string
	for v := range set {
		lastVal = v
	}
	return binStrToInt(lastVal)
}

func mostCommonBitAtPosition(set map[string]struct{}, pos int) byte {
	var count0s int
	for v := range set {
		if v[pos] == '0' {
			count0s += 1
		}
	}
	if count0s > len(set)/2 {
		return '0'
	}
	return '1'
}

func leastCommonBitAtPosition(set map[string]struct{}, pos int) byte {
	var count0s int
	for v := range set {
		if v[pos] == '0' {
			count0s += 1
		}
	}
	if count0s <= len(set)/2 {
		return '0'
	}
	return '1'
}

func getInput() []string {
	scanner := bufio.NewScanner(os.Stdin)

	in := make([]string, 0, 1024)
	for scanner.Scan() {
		line := scanner.Text()
		in = append(in, line)
	}

	return in
}
