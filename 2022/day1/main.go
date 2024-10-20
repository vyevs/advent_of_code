package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	in := getInput()
	part1(in)
	part2(in)
}

func part1(in [][]int) {
	var best int
	for _, e := range in {
		var totalCals int
		for _, c := range e {
			totalCals += c
		}
		
		if totalCals > best {
			best = totalCals
		}
	}
	
	fmt.Printf("the elf with the most calories is carrying %d calories\n", best)
}

func part2(in [][]int) {
	calTotals := make([]int, len(in))
	
	for i, e := range in {
		for _, c := range e {
			calTotals[i] += c
		}
	}
	
	sort.Ints(calTotals)
	l := len(calTotals)
	top3Total := calTotals[l-1] + calTotals[l-2] + calTotals[l-3]
	
	fmt.Printf("the 3 elves with the most calories are carrying a total of %d calories\n", top3Total)
}


func getInput() [][]int {
	s := bufio.NewScanner(os.Stdin)
	
	out := make([][]int, 0, 1024)
	haveMoreText := true
	for haveMoreText {
		calsList := make([]int, 0, 64)
		
		for {
			more := s.Scan()
			if !more {
				haveMoreText = false
				break
			}
			
			line := s.Text()
			if line == "" {
				break
			}
			
			cals, _ := strconv.Atoi(line)
			calsList = append(calsList, cals)
		}
		
		out = append(out, calsList)
	}
	
	
	return out
}