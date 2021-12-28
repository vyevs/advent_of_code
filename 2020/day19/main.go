package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	input := readInput()

	fmt.Println("Part 1:")
	doPart1(input)
	fmt.Println()

	fmt.Println("Part 2:")
	doPart2(input)
}

func createCNFGrammar(rules []inputRule) cnfGrammar {
	cnf := cnfGrammar{
		productions:     make([]cnfProduction, 0, len(rules)*len(rules)),
		lhsToProduction: make(map[string]*cnfProduction, len(rules)*len(rules)),
	}

	for _, rawRule := range rules {
		p := parseRawRule(rawRule.lhs, rawRule.rhs)
		cnf.productions = append(cnf.productions, p)
		cnf.lhsToProduction[p.from] = &cnf.productions[len(cnf.productions)-1]
	}

	changedGrammar := true
	for changedGrammar {
		changedGrammar = false

		if eliminateNonSolitaryTerminals(&cnf) {
			changedGrammar = true
		}
		if eliminateMultipleNonterminals(&cnf) {
			changedGrammar = true
		}
		if elimiateUnitRules(&cnf) {
			changedGrammar = true
		}
	}

	return cnf
}

// TERM: Eliminate rules with nonsolitary terminals.
// Returns whether at least one modification was made to the grammar.
func eliminateNonSolitaryTerminals(cnf *cnfGrammar) bool {
	var changedGrammar bool
	for i := range cnf.productions {
		p := &cnf.productions[i]
		if eliminateNonSolitaryTerminalsFromProduction(cnf, p) {
			changedGrammar = true
		}
	}

	return changedGrammar
}

// TERM: Eliminate rules with nonsolitary terminals.
// Returns whether at least one modification was made to the production.
func eliminateNonSolitaryTerminalsFromProduction(cnf *cnfGrammar, p *cnfProduction) bool {
	var changedProduction bool
	for i := range p.tos {
		if eliminateNonSolitaryTerminalsFromRule(cnf, p, i) {
			changedProduction = true
		}
	}
	return changedProduction
}

// TERM: Eliminate rules with nonsolitary terminals.
// Returns whether at least one modification was made to the rule for the production.
func eliminateNonSolitaryTerminalsFromRule(cnf *cnfGrammar, p *cnfProduction, ruleIdx int) bool {
	hasMultipleTerminals := func() bool {
		ts := p.isTerminal[ruleIdx]
		var ct int
		for _, isTerm := range ts {
			if isTerm {
				ct++
			}
		}
		return ct > 1
	}

	if !hasMultipleTerminals() {
		return false
	}

	rhs := p.tos[ruleIdx]
	isTerminal := p.isTerminal[ruleIdx]
	for i, r := range rhs {
		if !isTerminal[i] {
			continue
		}

		newRuleLHS := "BIN" + r

		// the symbol at index i is no longer terminal, and now refers to another rule
		rhs[i] = newRuleLHS
		isTerminal[i] = false

		_, haveLHS := cnf.lhsToProduction[newRuleLHS]
		if !haveLHS {
			newP := cnfProduction{
				from:       newRuleLHS,
				tos:        [][]string{[]string{r}},
				isTerminal: [][]bool{[]bool{true}},
			}
			cnf.productions = append(cnf.productions, newP)
			cnf.lhsToProduction[newP.from] = &cnf.productions[len(cnf.productions)-1]
		}

	}

	return true
}

// BIN: Eliminate right-hand sides with more than 2 nonterminals
// Returns whether the grammar was changed.
// This is intended to be called after a TERM has been done on the grammar,
// which makes sure that there are no non-solitary terminals.
func eliminateMultipleNonterminals(cnf *cnfGrammar) bool {
	var changedGrammar bool
	for i := range cnf.productions {
		p := &cnf.productions[i]
		if eliminateMultipleNonterminalsFromProduction(cnf, p) {
			changedGrammar = true
		}
	}
	return changedGrammar
}

// BIN: Eliminate right-hand sides with more than 2 nonterminals
// Returns whether the grammar was changed.
func eliminateMultipleNonterminalsFromProduction(cnf *cnfGrammar, p *cnfProduction) bool {
	var changedProduction bool
	for i := range p.tos {
		if eliminateMultipleNonterminalsFromRule(cnf, p, i) {
			changedProduction = true
		}
	}
	return changedProduction
}

// BIN: Eliminate right-hand sides with more than 2 nonterminals
// Returns whether the rule for the grammar was changed.
func eliminateMultipleNonterminalsFromRule(cnf *cnfGrammar, p *cnfProduction, ruleIdx int) bool {
	// This is true because this should be run only AFTER a BIN step has been performed.
	// Which guarantees that there are no intermingled terminals and non-terminals
	nonTerminalCt := len(p.tos[ruleIdx])

	if nonTerminalCt < 3 {
		return false
	}

	for i := len(p.tos[ruleIdx]) - 2; i >= 1; i-- {
		lhs1ToReplace := p.tos[ruleIdx][i]
		lhs2ToReplace := p.tos[ruleIdx][i+1]

		newRuleLHS := "TERM" + lhs1ToReplace + "/" + lhs2ToReplace

		p.tos[ruleIdx][i] = newRuleLHS

		_, haveLHS := cnf.lhsToProduction[newRuleLHS]
		if !haveLHS {
			newP := cnfProduction{
				from:       newRuleLHS,
				tos:        [][]string{[]string{lhs1ToReplace, lhs2ToReplace}},
				isTerminal: [][]bool{[]bool{false, false}},
			}

			cnf.productions = append(cnf.productions, newP)
			cnf.lhsToProduction[newP.from] = &cnf.productions[len(cnf.productions)-1]
		}

		p.tos[ruleIdx] = p.tos[ruleIdx][:len(p.tos[ruleIdx])-1]
	}

	return true
}

// UNIT: Eliminate unit rules
// Returns whether at least one modification was made to the grammar.
func elimiateUnitRules(cnf *cnfGrammar) bool {
	var changedGrammar bool
	for i := 0; ; i++ {
		// can't use for range here because the for range argument is a copy of cnf.productions
		// as it is in this function, but elimiateUnitRulesFromProduction possibly modifies it's length
		if i >= len(cnf.productions) {
			break
		}
		if elimiateUnitRulesFromProduction(cnf, i) {
			changedGrammar = true
		}
	}

	return changedGrammar
}

// UNIT: Eliminate unit rules
// Returns whether at least one modification was made to the production.
func elimiateUnitRulesFromProduction(cnf *cnfGrammar, pIdx int) bool {
	p := &cnf.productions[pIdx]
	for i, to := range p.tos {
		if len(to) == 1 && !p.isTerminal[i][0] {

			unitLHS := p.from
			unitRHS := to[0]

			for _, p := range cnf.productions {
				for _, to := range p.tos {
					for i, s := range to {
						if s == unitLHS {
							to[i] = unitRHS
						}
					}
				}
			}

			p.tos[i], p.tos[len(p.tos)-1] = p.tos[len(p.tos)-1], p.tos[i]
			p.tos = p.tos[:len(p.tos)-1]
			p.isTerminal[i], p.isTerminal[len(p.isTerminal)-1] = p.isTerminal[len(p.isTerminal)-1], p.isTerminal[i]
			p.isTerminal = p.isTerminal[:len(p.isTerminal)-1]

			cnf.productions[pIdx], cnf.productions[len(cnf.productions)-1] = cnf.productions[len(cnf.productions)-1], cnf.productions[pIdx]
			cnf.productions = cnf.productions[:len(cnf.productions)-1]
			delete(cnf.lhsToProduction, unitLHS)
			return true
		}
	}
	return false
}

func parseRawRule(lhs, rhs string) cnfProduction {
	p := cnfProduction{
		from: lhs,
		tos:  make([][]string, 0, 128),
	}

	rhsSplit := strings.Split(rhs, " ")

	to := make([]string, 0, 3)
	terminals := make([]bool, 0, 3)
	for _, str := range rhsSplit {
		if str == "|" {
			p.tos = append(p.tos, to)
			p.isTerminal = append(p.isTerminal, terminals)

			to = make([]string, 0, 3)
			terminals = make([]bool, 0, 3)
			continue
		}

		if str[0] == '"' {
			to = append(to, str[1:len(str)-1])
			terminals = append(terminals, true)
		} else {
			to = append(to, str)
			terminals = append(terminals, false)
		}
	}

	p.tos = append(p.tos, to)
	p.isTerminal = append(p.isTerminal, terminals)

	return p
}

func doPart1(input input) {
	defer func(s time.Time) {
		fmt.Printf("That took %s\n", time.Since(s))
	}(time.Now())

	cnf := createCNFGrammar(input.rules)
	fmt.Printf("%s\n", cnf)

	var nValid int
	for _, m := range input.messages {
		isValid := doesStrSatisfyGrammar(cnf, m)

		if isValid {
			nValid++
		}
	}

	fmt.Printf("%d messages are valid\n", nValid)
}

func doesStrSatisfyGrammar(cnf cnfGrammar, str string) bool {
	n := len(str)
	r := len(cnf.productions)

	P := make([][][]bool, len(str))
	for i := 0; i < n; i++ {
		P[i] = make([][]bool, len(str))
		for j := 0; j < n; j++ {
			P[i][j] = make([]bool, r)
		}
	}

	for s := 0; s < n; s++ {
		var v int

		as := str[s : s+1]
		for _, p := range cnf.productions {
			for _, to := range p.tos {
				if len(to) == 1 && to[0] == as {
					P[0][s][v] = true
				}
				v++
			}
		}
	}

	for spanLen := 2; spanLen <= n; spanLen++ {

	}

	return P[n-1][0][0]
}

func doPart2(input input) {

}

type cnfGrammar struct {
	productions     []cnfProduction
	lhsToProduction map[string]*cnfProduction
}

// A cnfProduction is a production rule that in CNF form,
type cnfProduction struct {
	from string

	tos        [][]string
	isTerminal [][]bool // whether to[i][j] is a terminal
}

func (cnf cnfGrammar) String() string {
	var buf strings.Builder
	buf.Grow(128)

	for i, p := range cnf.productions {
		buf.WriteString(p.String())

		if i < len(cnf.productions)-1 {
			buf.WriteByte('\n')
		}
	}

	return buf.String()
}

func (p cnfProduction) String() string {
	var buf strings.Builder
	buf.Grow(64)

	for i, to := range p.tos {
		buf.WriteString(fmt.Sprintf("%-3s", p.from))
		buf.WriteString(" -> ")

		for j, t := range to {

			if p.isTerminal[i][j] {
				buf.WriteByte('"')
			}
			buf.WriteString(t)
			if p.isTerminal[i][j] {
				buf.WriteByte('"')
			}

			if j < len(to)-1 {
				buf.WriteByte(' ')
			}
		}

		if i < len(p.tos)-1 {
			buf.WriteByte('\n')
		}
	}

	return buf.String()
}

type input struct {
	rules    []inputRule
	messages []string
}

type inputRule struct {
	lhs string
	rhs string
}

func readInput() input {
	scanner := bufio.NewScanner(os.Stdin)

	var input input

	input.rules = make([]inputRule, 0, 128)
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			break
		}

		colonIdx := strings.IndexByte(line, ':')

		ruleLHS := line[:colonIdx]
		ruleRHS := line[colonIdx+2:]

		inputRule := inputRule{
			lhs: ruleLHS,
			rhs: ruleRHS,
		}

		input.rules = append(input.rules, inputRule)
	}

	input.messages = make([]string, 0, 512)
	for scanner.Scan() {
		line := scanner.Text()

		input.messages = append(input.messages, line)
	}

	return input
}
