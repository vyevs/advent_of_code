package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := readInput()

	fmt.Println("Part 1:")
	doPart1(input)
	fmt.Println()

	fmt.Println("Part 2:")
	doPart2(input)
}

func doPart1(input input) {
	var errRate int
	for _, t := range input.nearbyTickets {
		for _, f := range t.fields {
			if !valueSatisfiesRestrictions(f, input.restrictions) {
				errRate += f
			}
		}
	}

	fmt.Printf("The ticket scanning error rate is %d\n", errRate)
}

func valueSatisfiesRestrictions(v int, rs []restriction) bool {
	for _, r := range rs {
		if valueSatisfiesRestriction(v, r) {
			return true
		}
	}
	return false
}

func valueSatisfiesRestriction(v int, r restriction) bool {
	return (v >= r.r1Min && v <= r.r1Max) || (v >= r.r2Min && v <= r.r2Max)
}

func doPart2(input input) {
	input.nearbyTickets = removeInvalidTickets(input)

	// fieldNamePossibilities[i] is the possible names that field number i of a ticket can have
	fieldNamePossibilities := buildInitialFieldNamePossibilities(input.restrictions)

	for _, nt := range input.nearbyTickets {
		for i, f := range nt.fields {
			possibiliesForI := fieldNamePossibilities[i]

			for _, r := range input.restrictions {
				// if a ticket's field does not satisfy a restriction, that particular
				// field column cannot have that restriction's name
				if !valueSatisfiesRestriction(f, r) {
					possibiliesForI = removeStr(possibiliesForI, r.fieldName)
				}
			}

			fieldNamePossibilities[i] = possibiliesForI
		}
	}

	for removedAtLeastOne := true; removedAtLeastOne; {
		removedAtLeastOne = false

		for _, possibleForI := range fieldNamePossibilities {

			// the ith field's name is already determined by the fact we only have one choice
			// then we can remove this field name from all other fields' possibilities
			if len(possibleForI) == 1 {
				toRemove := possibleForI[0]
				if removeStrFromAllStrSlices(fieldNamePossibilities, toRemove) {
					removedAtLeastOne = true
				}
			}
		}
	}

	product := 1
	for i, posForI := range fieldNamePossibilities {
		fieldName := posForI[0]

		if strings.HasPrefix(fieldName, "departure") {
			product *= input.yourTicket.fields[i]
		}
	}

	fmt.Printf("Product of the six fields that begin with %q is %d\n", "departure", product)
}

func buildInitialFieldNamePossibilities(rs []restriction) [][]string {
	pos := make([][]string, 0, len(rs))
	for i := 0; i < len(rs); i++ {
		iPos := make([]string, 0, len(rs))

		for _, r := range rs {
			iPos = append(iPos, r.fieldName)
		}

		pos = append(pos, iPos)
	}
	return pos
}

// returns whether at least one was removed
func removeStrFromAllStrSlices(ss [][]string, toRemove string) bool {
	var removedAtLeastOne bool
	for i, s := range ss {
		if len(s) != 1 {
			newS := removeStr(s, toRemove)
			ss[i] = newS

			if len(newS) != len(s) {
				removedAtLeastOne = true
			}
		}
	}

	return removedAtLeastOne
}

func removeStr(strs []string, str string) []string {
	for i, v := range strs {
		if v == str {
			strs[i], strs[len(strs)-1] = strs[len(strs)-1], strs[i]
			strs = strs[:len(strs)-1]
			return strs
		}
	}
	return strs
}

func removeInvalidTickets(input input) []ticket {
	nts := input.nearbyTickets

	for i := 0; i < len(nts); {
		t := nts[i]
		if !ticketSatifiesRestrictions(t, input.restrictions) {
			// removal of nt[i] from nt
			nts[i], nts[len(nts)-1] = nts[len(nts)-1], nts[i]
			nts = nts[:len(nts)-1]
		} else {
			i++
		}
	}

	return nts
}

func ticketSatifiesRestrictions(t ticket, rs []restriction) bool {
	for _, f := range t.fields {
		if !valueSatisfiesRestrictions(f, rs) {
			return false
		}
	}
	return true
}

type input struct {
	restrictions  []restriction
	yourTicket    ticket
	nearbyTickets []ticket
}

type restriction struct {
	fieldName    string
	r1Min, r1Max int
	r2Min, r2Max int
}

type ticket struct {
	fields []int
}

func readInput() input {
	scanner := bufio.NewScanner(os.Stdin)

	var input input
	input.restrictions = readRestrictions(scanner)
	input.yourTicket = readYourTicket(scanner)
	input.nearbyTickets = readNearbyTickets(scanner)

	return input
}

func readRestrictions(scanner *bufio.Scanner) []restriction {
	restrictions := make([]restriction, 0, 64)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if line == "" {
			break
		}

		restriction := parseRestriction(line)

		restrictions = append(restrictions, restriction)
	}

	return restrictions
}

func parseRestriction(line string) restriction {
	var restriction restriction

	colonIdx := strings.IndexByte(line, ':')
	restriction.fieldName = line[:colonIdx]

	bothRangeStrs := line[colonIdx+2:]
	rangeStrs := strings.Split(bothRangeStrs, " or ")

	restriction.r1Min, restriction.r1Max = parseRangeStr(rangeStrs[0])

	restriction.r2Min, restriction.r2Max = parseRangeStr(rangeStrs[1])

	return restriction
}

// parses a string of the form number1-number2 and returns number1, number2
func parseRangeStr(str string) (int, int) {
	parts := strings.Split(str, "-")

	v1Str := parts[0]
	v1, _ := strconv.Atoi(v1Str)

	v2Str := parts[1]
	v2, _ := strconv.Atoi(v2Str)

	return v1, v2
}

func readYourTicket(scanner *bufio.Scanner) ticket {
	scanner.Scan() // to eat the "your ticket:" line

	scanner.Scan()
	fieldsStr := scanner.Text()

	var ticket ticket
	ticket.fields = parseFieldsString(fieldsStr)

	scanner.Scan() // to eat the empty line after your ticket fields

	return ticket
}

func readNearbyTickets(scanner *bufio.Scanner) []ticket {
	scanner.Scan() // to eat the "nearby tickets:" line

	tickets := make([]ticket, 0, 128)
	for scanner.Scan() {
		line := scanner.Text()

		var ticket ticket
		ticket.fields = parseFieldsString(line)

		tickets = append(tickets, ticket)
	}

	return tickets
}

func parseFieldsString(fieldsStr string) []int {
	parts := strings.Split(fieldsStr, ",")

	fields := make([]int, 0, len(fieldsStr)/2+1)

	for _, part := range parts {
		field, _ := strconv.Atoi(part)
		fields = append(fields, field)
	}

	return fields
}
