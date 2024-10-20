package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := getInput()
	part1(in)
	part2(in)
}

func part1(in []string) {
	total := sumPrioritiesOfCommons(in)
	
	fmt.Printf("sum of the priorities of the common items: %d\n", total)
}

func part2(in []string) {
	total := sumGroupBadgePriorities(in)
	
	fmt.Printf("sum of the priorities of the group badge items: %d\n", total)
}

func sumGroupBadgePriorities(in []string) int {
	badgeItems := determineBadgeItems(in)
	var total int
	for _, v := range badgeItems {
		total += itemPriority(v)
	}
	return total
}

func determineBadgeItems(in []string) []rune {
	out := make([]rune, 0, len(in)/3)
	for i := 0; i < len(in); i += 3 {
		group := in[i:i+3]
		out = append(out, determineBadgeItem(group))
	}
	
	return out
}

func determineBadgeItem(group []string) rune {
	for _, r1 := range group[0] {
		for _, r2 := range group[1] {
			for _, r3 := range group[2] {
				if r1 == r2 && r2 == r3 {
					return r1
				}
			}
		}
	}
	
	panic(fmt.Sprintf("group %v has no badge item", group))
}

func sumPrioritiesOfCommons(in []string) int {
	var total int
	for _, v := range in {
		common := getCommon(v)
		total += itemPriority(common)
	}
	
	return total
}

func getCommon(s string) rune {
	halfwayIdx := len(s)/2
	p1 := s[:halfwayIdx]
	p2 := s[halfwayIdx:]
	
	for _, r1 := range p1 {
		for _, r2 := range p2 {
			if r1 == r2 {
				return r1
			}
		}
	}
	
	panic(fmt.Sprintf("item %s contains no commonality between the 2 halves", s))
}

func itemPriority(i rune) int {
	if i >= 'a' && i <= 'z' { // a through z
		return int(i - 'a' + 1)
	}
	
	return int(i - 'A' + 27) // A through Z
}


func getInput() []string {
	s := bufio.NewScanner(os.Stdin)
	
	out := make([]string, 0, 1024)

	for s.Scan() {
		l := s.Text()
		out = append(out, l)
	}
	
	return out
}