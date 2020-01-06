package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	f, _ := os.Open(os.Args[1])
	lines := readInput(f)
	f.Close()

	msgMostCommon := errorCorrectedMessageMostCommon(lines)
	fmt.Printf("error corrected most common: %s\n", msgMostCommon)

	msgLeastCommon := errorCorrectedMessageLeastCommon(lines)
	fmt.Printf("error corrected least common: %s\n", msgLeastCommon)
}

func errorCorrectedMessageMostCommon(lines []string) string {
	columnCounts := make([][26]int, len(lines[0]))

	for _, line := range lines {
		for i, r := range line {
			columnCounts[i][r-'a']++
		}
	}

	var buf strings.Builder
	for _, column := range columnCounts {
		maxIdx := 0
		max := column[0]
		for i := 1; i < len(column); i++ {
			if column[i] > max {
				maxIdx = i
				max = column[i]
			}
		}

		buf.WriteRune(rune('a' + maxIdx))
	}

	return buf.String()
}

func errorCorrectedMessageLeastCommon(lines []string) string {
	columnCounts := make([][26]int, len(lines[0]))

	for _, line := range lines {
		for i, r := range line {
			columnCounts[i][r-'a']++
		}
	}

	var buf strings.Builder
	for _, column := range columnCounts {
		minIdx := 0
		min := column[0]
		for i := 1; i < len(column); i++ {
			if column[i] > 0 && column[i] < min {
				minIdx = i
				min = column[i]
			}
		}

		buf.WriteRune(rune('a' + minIdx))
	}

	return buf.String()
}

func readInput(r io.Reader) []string {
	scanner := bufio.NewScanner(r)

	lines := make([]string, 0, 512)

	for scanner.Scan() {
		line := scanner.Text()

		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("scanner: %v", err)
	}

	return lines
}
