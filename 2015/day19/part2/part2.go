package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func main() {
	moleculeMappings, target := readInput()

	log.Printf("target: %q", target)
	for to, from := range moleculeMappings {
		log.Printf("%s -> %s", to, from)
	}

	/*
		allMolecules := make([]string, 0, len(moleculeMappings)*2)
		for from, allTos := range moleculeMappings {
			allMolecules = append(allMolecules, from)
			allMolecules = append(allMolecules, allTos...)
		}

		log.Printf("all molecules: %v", allMolecules)

		allPossibleDivisions := makeAllPossibleDivisions(target, allMolecules)
		log.Printf("there are %d possible divisions", len(allPossibleDivisions))
	*/
}

func makeAllPossibleDivisions(str string, possibleMolecules []string) [][]string {
	moleculesSet := make(map[string]struct{}, len(possibleMolecules))
	for _, m := range possibleMolecules {
		moleculesSet[m] = struct{}{}
	}

	var recurse func(str string)

	divisions := make([][]string, 0, 1024)
	curDivision := make([]string, 0, len(str))
	recurse = func(str string) {
		log.Printf("recurse(%q)", str)
		if len(str) == 0 {
			return
		}
		if _, isSingleMolecule := moleculesSet[str]; isSingleMolecule {
			cpy := make([]string, len(curDivision))
			copy(cpy, curDivision)
			cpy = append(cpy, str)
			divisions = append(divisions, cpy)
		}

		partStart := 0
		partEnd := 1
		for ; partEnd <= len(str); partEnd++ {
			part := str[partStart:partEnd]
			if _, isMolecule := moleculesSet[part]; isMolecule {
				curDivision = append(curDivision, part)
				recurse(str[partEnd:])
				curDivision = curDivision[:len(curDivision)-1]
			}
		}
	}

	recurse(str)

	return divisions
}

func readInput() (map[string][]string, string) {
	scanner := bufio.NewScanner(os.Stdin)

	moleculeMappings := make(map[string][]string, 128)

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			break
		}

		tokens := strings.Split(line, " ")

		from := tokens[0]
		to := tokens[2]

		moleculeMappings[from] = append(moleculeMappings[from], to)
	}

	scanner.Scan()
	target := scanner.Text()

	return moleculeMappings, target
}
