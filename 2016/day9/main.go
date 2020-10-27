package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, _ := os.Open(os.Args[1])

	lines := readLines(f)
	f.Close()

	{
		var decompressedLen int
		for _, line := range lines {
			//decompressedLine := decompress(line)

			decompressedLen += decompressedLength(line)
		}

		fmt.Printf("part 1: %d\n", decompressedLen)
	}

	{
		var decompressedLen int
		for _, line := range lines {
			//decompressedLine := decompress(line)

			decompressedLen += decompressedLength2(line, len(line), 1)
		}

		fmt.Printf("part 2: %d\n", decompressedLen)
	}

}

func decompressedLength(s string) int {
	var decompressedLen int

	for i := 0; i < len(s); {
		if s[i] == '(' {
			markerStr, ok := tryToReadMarker(s[i:])
			if ok {
				i += len(markerStr)

				markerStr = markerStr[1 : len(markerStr)-1]

				markerSplit := strings.Split(markerStr, "x")

				nCharsStr := markerSplit[0]
				nRepeatsStr := markerSplit[1]

				nChars, _ := strconv.Atoi(nCharsStr)
				nRepeats, _ := strconv.Atoi(nRepeatsStr)

				charsUntil := nChars + i
				if charsUntil > len(s) {
					charsUntil = len(s)
				}

				decompressedLen += (charsUntil - i) * nRepeats

				i = charsUntil

			} else {
				i++
				decompressedLen++
			}
		} else {
			i++
			decompressedLen++
		}
	}

	return decompressedLen
}

func decompressedLength2(s string, chars, repeats int) int {

	var decompressedLen int

	for rep := 0; rep < repeats; rep++ {
		for i := 0; i < len(s) && i < chars; {

			if s[i] == '(' {
				marker, ok := tryToReadMarker(s[i:])
				if ok {

					markerNoParens := marker[1 : len(marker)-1]

					split := strings.Split(markerNoParens, "x")

					nChars, _ := strconv.Atoi(split[0])
					nRepeats, _ := strconv.Atoi(split[1])

					decompressedLen += decompressedLength2(s[i+len(marker):], nChars, nRepeats)

					i += len(marker) + nChars
				} else {
					decompressedLen++
					i++
				}

			} else {
				decompressedLen++
				i++
			}
		}
	}

	return decompressedLen
}

func tryToReadMarker(s string) (string, bool) {
	if s == "" {
		return "", false
	}

	var i int
	if s[i] != '(' {
		return "", false
	}

	i++

	var seenDigit bool
	for ; i < len(s) && s[i] >= '0' && s[i] <= '9'; i++ {
		seenDigit = true
	}

	if i == len(s) || !seenDigit {
		return "", false
	}

	if s[i] != 'x' {
		fmt.Println("r4")
		return "", false
	}

	i++

	for seenDigit = false; i < len(s) && s[i] >= '0' && s[i] <= '9'; i++ {
		seenDigit = true
	}

	if i == len(s) || !seenDigit {
		return "", false
	}

	if s[i] != ')' {
		return "", false
	}

	i++

	return s[:i], true
}

func readLines(r io.Reader) []string {
	scanner := bufio.NewScanner(r)

	lines := make([]string, 0, 128)

	for scanner.Scan() {
		line := scanner.Text()

		lines = append(lines, line)
	}

	return lines
}
