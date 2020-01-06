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

	var fn func()

	curPerm := "e"
	var steps int

	out, _ := os.OpenFile("output.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0755)

	fn = func() {
		fmt.Fprintln(out, curPerm)
		if len(curPerm) == len(molecule) && curPerm == molecule {
			fmt.Println(steps)
			os.Exit(0)
		}

		for _, r := range replacements {
			indices := indexAll(curPerm, r.str)
			for _, repl := range r.replacements {
				for _, ind := range indices {
					oldCurPerm := curPerm
					curPerm = curPerm[:ind] + repl + curPerm[ind+len(r.str):]
					if len(curPerm) <= len(molecule) {
						steps++
						fn()
						steps--
					}
					curPerm = oldCurPerm
				}
			}
		}
	}

	fn()
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
