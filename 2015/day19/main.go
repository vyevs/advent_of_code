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
	if len(os.Args) != 2 {
		log.Fatalf("provide 1 input file")
	}

	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatalf("open: %v", err)
	}

	molecule, replacements := parseInput(f)

	/*
		fmt.Println(molecule)
		for _, r := range replacements {
			fmt.Printf("%s -> ", r.str)
			for _, replacement := range r.replacements {
				fmt.Printf("%s ", replacement)
			}
			fmt.Println()
		}
	*/

	seen := make(map[string]struct{}, 1024)
	for _, replacement := range replacements {
		indices := indexAll(molecule, replacement.str)
		for _, r := range replacement.replacements {
			for _, i := range indices {

				newMolecule := molecule[:i] + r + molecule[i+len(replacement.str):]

				seen[newMolecule] = struct{}{}
			}
		}
	}

	fmt.Println(len(seen))
	/*
		for m := range seen {
			fmt.Println(m)
		}*/
}

func indexAll(s, substr string) []int {
	indices := make([]int, 0, 64)
	for i := 0; i < len(s)-len(substr)+1; i++ {
		if s[i:i+len(substr)] == substr {
			indices = append(indices, i)
		}
	}
	return indices
}

type replacement struct {
	str          string
	replacements []string
}

func parseInput(r io.Reader) (string, []replacement) {
	scanner := bufio.NewScanner(r)

	replacements := make([]replacement, 0, 128)
	replacementsIndex := -1

	var curStr string

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			break
		}

		tokens := strings.Split(line, " ")

		if tokens[0] != curStr {
			curStr = tokens[0]
			replacementsIndex++
			replacements = append(replacements, replacement{str: curStr, replacements: make([]string, 0, 128)})
		}

		replacements[replacementsIndex].replacements = append(replacements[replacementsIndex].replacements, tokens[2])
	}

	scanner.Scan()
	molecule := scanner.Text()

	if err := scanner.Err(); err != nil {
		log.Fatalf("Scanner: %v", err)
	}

	return molecule, replacements
}
