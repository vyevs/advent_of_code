package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
)

func main() {

	{
		rows := readSpreadsheet("part1_example.txt")
		fmt.Printf("part 1: example checksum is %d\n", checksum(rows))
	}
	{
		rows := readSpreadsheet("part2_example.txt")
		fmt.Printf("part 2: example sum of evenly disibles: %d\n", sumEvenlyDivibleResults(rows))
	}

	rows := readSpreadsheet("input.txt")

	checksum := checksum(rows)

	fmt.Printf("part 1: the checksum is %d\n", checksum)
	fmt.Printf("part 2: the sum is %d\n", sumEvenlyDivibleResults(rows))
}

func checksum(rows [][]int) int {
	var checksum int

	for _, row := range rows {
		minV := slices.Min(row)
		maxV := slices.Max(row)

		diff := maxV - minV

		checksum += diff
	}

	return checksum
}

func sumEvenlyDivibleResults(rows [][]int) int {
	var sum int

	for _, row := range rows {
		sum += evenlyDivisbleResult(row)
	}

	return sum
}

func evenlyDivisbleResult(row []int) int {
	var sum int

	for i, v1 := range row {
		for j := i + 1; j < len(row); j++ {
			v2 := row[j]

			if v1 > v2 && v1%v2 == 0 {
				sum += v1 / v2
			} else if v2 > v1 && v2%v1 == 0 {
				sum += v2 / v1
			}

		}
	}

	return sum
}

func readSpreadsheet(fName string) [][]int {
	f, _ := os.Open(fName)
	defer f.Close()

	sc := bufio.NewScanner(f)

	rows := make([][]int, 0, 64)

	for sc.Scan() {
		line := sc.Text()

		row := make([]int, 0, 64)

		parts := splitOnWhiteSpace(line)

		for _, p := range parts {
			v, _ := strconv.Atoi(p)
			row = append(row, v)
		}

		rows = append(rows, row)
	}

	return rows
}

func splitOnWhiteSpace(s string) []string {
	out := make([]string, 0, 32)

	buf := make([]byte, 0, 16)
	for i := range len(s) {
		c := s[i]

		if c == ' ' || c == '\t' || c == '\n' {
			if len(buf) > 0 {
				out = append(out, string(buf))
				buf = buf[:0]
			}
		} else {
			buf = append(buf, c)
		}
	}

	if len(buf) > 0 {
		out = append(out, string(buf))
	}

	return out
}
