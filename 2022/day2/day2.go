package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	in := getInput()
	part1(in)
	part2(in)
}

func part1(in [][2]rune) {
	rounds := inputToRounds(in)
	total := finalScore(rounds)
	
	fmt.Printf("total score after all the rounds is %d\n", total)
}

func finalScore(rs []round) int {
	var total int
	for _, r := range rs {
		total += r.score()
	}
	return total
}

type move int
const (
	rock move = iota
	paper
	scissors
)

func (m move) score() int {
	switch m {
		case rock: return 1
		case paper: return 2
		case scissors: return 3
	}
	panic(fmt.Sprintf("unknown move %d", m))
}

func (m move) winsAgainst() move {
	switch m {
		case rock: return scissors
		case paper: return rock
		case scissors: return paper
	}
	panic(fmt.Sprintf("unknown move %d", m))
}

func (m move) losesAgainst() move {
	switch m {
		case rock: return paper
		case paper: return scissors
		case scissors: return rock
	}
	panic(fmt.Sprintf("unknown move %d", m))
}

const (
	loseScore = 0
	drawScore = 3
	winScore = 6
)

var scores = map[[2]move]int {
	{rock, paper}: loseScore,
	{paper, scissors}: loseScore,
	{scissors, rock}: loseScore,
	
	{rock, rock}: drawScore,
	{paper, paper}: drawScore,
	{scissors, scissors}: drawScore,
	
	{rock, scissors}: winScore,
	{paper, rock}: winScore,
	{scissors, paper}: winScore,
}

func (m move) result(other move) int {
	return scores[[2]move{m, other}]
}

type round struct {
	theirs, ours move
	wantOutcome int
}

func inputToRounds(in [][2]rune) []round {
	out := make([]round, 0, len(in))
	
	for _, r := range in {
		out = append(out, inputToRound(r))
	}
	
	return out
}

func inputToRoundsP2(in [][2]rune) []round {
	out := make([]round, 0, len(in))
	
	for _, r := range in {
		out = append(out, inputToRoundP2(r))
	}
	
	return out
}


var inputToMove = map[rune]move {
	'A': rock,
	'B': paper,
	'C': scissors,
	'X': rock,
	'Y': paper,
	'Z': scissors,
}

func inputToRound(r [2]rune) round {
	return round {
		theirs: inputToMove[r[0]],
		ours: inputToMove[r[1]],
	}
}

func inputToRoundP2(r [2]rune) round {
	theirs := inputToMove[r[0]]
	
	outcome := r[1]
	
	var ours move
	
	if outcome == 'X' { // lose
		ours = theirs.winsAgainst()
	} else if outcome == 'Y' { // draw
		ours = theirs
	} else { // win
		ours = theirs.losesAgainst()
	}
	
	return round {
		ours: ours,
		theirs: theirs,
	}
}

func (r round) score() int {	
	return r.ours.score() + scores[[2]move{r.ours, r.theirs}]
}

func part2(in [][2]rune) {
	rounds := inputToRoundsP2(in)
	
	total := finalScore(rounds)
	
	fmt.Printf("total score after all the rounds is %d\n", total)
}

func getInput() [][2]rune {
	s := bufio.NewScanner(os.Stdin)
	
	
	out := make([][2]rune, 0, 1024*10)
	
	for s.Scan() {
		line := s.Text()
		
		splt := strings.Split(line, " ")
		
		round := [2]rune{rune(splt[0][0]), rune(splt[1][0])}
		
		out = append(out, round)
	}
	
	
	return out
}