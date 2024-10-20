package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	stacks, moves := getInput()
	part2(stacks, moves)
}

func part1(stacks [][]rune, moves []move) {
	for _, m := range moves {
		makeMoveP1(stacks, m)
	}
	
	for _, s := range stacks {
		fmt.Printf("%c", s[len(s)-1])
	}
	fmt.Println()
}

func part2(stacks [][]rune, moves []move) {
	for _, m := range moves {
		makeMoveP2(stacks, m)
	}
	
	for _, s := range stacks {
		fmt.Printf("%c", s[len(s)-1])
	}
	fmt.Println()
}

func makeMoveP1(stacks [][]rune, m move) {
	for i := 0; i < m.ct; i++ {
		fromIdx := len(stacks[m.from]) - 1
		stacks[m.to] = append(stacks[m.to], stacks[m.from][fromIdx])
		stacks[m.from] = stacks[m.from][:fromIdx]
	}
}

func makeMoveP2(stacks [][]rune, m move) {
	fromIdx := len(stacks[m.from]) - m.ct
	stacks[m.to] = append(stacks[m.to], stacks[m.from][fromIdx:]...)
	stacks[m.from] = stacks[m.from][:fromIdx]
}

type move struct {
	from, to, ct int
}

func getInput() ([][]rune, []move) {
	
	lines := readLines()
	
	stacks := parseStacks(lines)
	moves := parseMoves(lines)
		
	return stacks, moves
}

func parseStacks(lines []string) [][]rune {
	emptyStrIdx := indexEmptyStr(lines)
	stackLines := lines[:emptyStrIdx]
	
	stackNumsLine := stackLines[len(stackLines)-1]
	var nStacks int
	r := strings.NewReader(stackNumsLine)
	for {
		var stackNum int
		_, err := fmt.Fscanf(r, "%d", &stackNum)
		if err == io.EOF {
			break
		}
		nStacks += 1
	}

	stacks := make([][]rune, nStacks)
	
	
	crateLines := stackLines[:len(stackLines)-1]
	for _, l := range crateLines {
		for i := 0; i < nStacks; i++ {
			crateIdx := (4*i)+1
		
			crate := l[crateIdx]
			
			if crate == ' ' {
				continue
			}
		
			stacks[i] = append([]rune{rune(crate)}, stacks[i]...)
		}
	}
	
	
	return stacks
}

func parseMoves(lines []string) []move {
	ms := make([]move, 0, 1024)
	
	emptyStrIdx := indexEmptyStr(lines)
	moveLines := lines[emptyStrIdx+1:]
	
	for _, l := range moveLines {
		m := parseMoveLine(l)
		ms = append(ms, m)
	}
	
	
	return ms
}

func parseMoveLine(l string) move {
	var m move
	_, _ = fmt.Sscanf(l, "move %d from %d to %d", &m.ct, &m.from, &m.to)
	m.from -= 1
	m.to -= 1
	return m
}

func indexEmptyStr(strs []string) int {
	for i, v := range strs {
		if v == "" {
			return i
		}
	}
	panic(fmt.Sprintf("%v does not contain an empty string", strs))
}

func readLines() []string {
	lines := make([]string, 0, 1024)
	
	s := bufio.NewScanner(os.Stdin)
	
	for s.Scan() {
		lines = append(lines, s.Text())
	}
	
	return lines
}
