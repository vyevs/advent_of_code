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
		indices := kmpIndexAll(molecule, replacement.str)
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

func kmpTable(s string) []int {
	pos := 1
	cnd := 0

	table := make([]int, len(s)+1)
	table[0] = -1

	for pos < len(s) {
		if s[pos] == s[cnd] {
			table[pos] = table[cnd]
		} else {
			table[pos] = cnd
			cnd = table[cnd]
			for cnd >= 0 && s[pos] != s[cnd] {
				cnd = table[cnd]
			}
		}
		pos++
		cnd++
	}

	table[pos] = cnd

	return table
}

func kmpIndexAll(s, w string) []int {
	table := kmpTable(w)

	var j, k int

	positions := make([]int, 0, 8)

	for j < len(s) {
		if w[k] == s[j] {
			j++
			k++
			if k == len(w) {
				positions = append(positions, j-k)
				k = table[k]
			}
		} else {
			k = table[k]
			if k < 0 {
				j++
				k++
			}
		}
	}

	return positions
}
