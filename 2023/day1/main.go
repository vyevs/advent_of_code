package main

import (
	"errors"
	"fmt"
	"os"
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
	if len(os.Args) != 2 {
		return errors.New("provide one command line arg that specifies the input file")
	}

	lines, err := vtools.ReadLines(os.Args[1])
	if err != nil {
		return fmt.Errorf("failed to read input file lines: %v", err)
	}

	if true {
		sumOfCVs := vtools.SumSlice(vtools.MapSlice(lines, getCalibrationValue1))
		fmt.Printf("part 1: the sum of the calibration values is %d\n", sumOfCVs)
	}

	if true {
		sumOfCVs := vtools.SumSlice(vtools.MapSlice(lines, getCalibrationValue2))
		fmt.Printf("part 2: the sum of the calibration values is %d\n", sumOfCVs)
	}

	return nil
}

type wordAndDigit struct {
	word  string
	digit byte
}

var wordsAndDigits = []wordAndDigit{
	{"one", '1'},
	{"two", '2'},
	{"six", '6'},
	{"four", '4'},
	{"five", '5'},
	{"nine", '9'},
	{"three", '3'},
	{"seven", '7'},
	{"eight", '8'},
}

func getCalibrationValue2(line string) int {
	firstDigit := getFirstDigit2(line)
	lastDigit := getLastDigit2(line)

	return int(10*(firstDigit-'0') + (lastDigit - '0'))
}

func getFirstDigit2(line string) byte {
outer:
	for i := range len(line) {
		c := line[i]
		if isDigit(c) {
			return c
		}

		for _, it := range wordsAndDigits {
			if len(it.word) > i+1 {
				continue outer
			}
			if it.word == line[i-len(it.word)+1:i+1] {
				return it.digit
			}
		}

	}

	panic(fmt.Sprintf("the string %q does not contain a digit"))
}

func getLastDigit2(line string) byte {
outer:
	for i := len(line) - 1; i >= 0; i-- {
		c := line[i]
		if isDigit(c) {
			return c
		}

		for _, it := range wordsAndDigits {
			if len(it.word) > len(line)-i {
				continue outer
			}
			if it.word == line[i:i+len(it.word)] {
				return it.digit
			}
		}
	}

	panic(fmt.Sprintf("the string %q does not contain a digit", line))
}

func getCalibrationValue1(line string) int {
	firstDigit := getFirstDigit(line)
	lastDigit := getLastDigit(line)
	return int(10*(firstDigit-'0') + (lastDigit - '0'))
}

func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func getFirstDigit(s string) byte {
	for i := range len(s) {
		c := s[i]
		if isDigit(c) {
			return c
		}
	}
	panic(fmt.Sprintf("the string %q does not contain a digit", s))
}

func getLastDigit(s string) byte {
	for i := len(s) - 1; i >= 0; i-- {
		c := s[i]
		if isDigit(c) {
			return c
		}
	}
	panic(fmt.Sprintf("the string %q does not contain a digit", s))
}
