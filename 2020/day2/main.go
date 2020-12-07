package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	pws := readInInput()

	fmt.Println("Part 1:")
	doPart1(pws)
	fmt.Println()

	fmt.Println("Part 2:")
	doPart2(pws)
}

func doPart1(pws []pw) {
	var nValid int
	for _, p := range pws {
		if p.isValid1() {
			nValid++
		}
	}

	fmt.Printf("%d passwords are valid\n", nValid)
}

func (p pw) isValid1() bool {
	var targetCharCt int

	for i := 0; i < len(p.str); i++ {
		if p.str[i] == p.targetChar {
			targetCharCt++
		}
	}

	return targetCharCt >= p.min && targetCharCt <= p.max
}

func (p pw) isValid2() bool {
	minIdx := p.min - 1
	maxIdx := p.max - 1

	return (p.str[minIdx] == p.targetChar && p.str[maxIdx] != p.targetChar) ||
		(p.str[minIdx] != p.targetChar && p.str[maxIdx] == p.targetChar)
}

func doPart2(pws []pw) {
	var nValid int
	for _, p := range pws {
		if p.isValid2() {
			nValid++
		}
	}

	fmt.Printf("%d passwords are valid\n", nValid)
}

type pw struct {
	targetChar byte
	str        string
	min, max   int
}

func readInInput() []pw {
	scanner := bufio.NewScanner(os.Stdin)

	pws := make([]pw, 0, 1024)
	for scanner.Scan() {
		line := scanner.Text()

		tokens := strings.Split(line, " ")

		seg1Tokens := strings.Split(tokens[0], "-")
		minStr := seg1Tokens[0]
		maxStr := seg1Tokens[1]

		min, _ := strconv.Atoi(minStr)
		max, _ := strconv.Atoi(maxStr)

		targetChar := tokens[1][0]

		pwStr := tokens[2]

		pws = append(pws, pw{str: pwStr, min: min, max: max, targetChar: targetChar})
	}

	return pws
}
