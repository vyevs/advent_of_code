package main

import (
	"fmt"
	"hash/maphash"
	"slices"
	"strconv"
	"strings"
	"time"
)

func (p params) stateStr(state []byte) string {
	var b strings.Builder
	b.Grow(128)

	for floor := p.topFloor; floor >= p.bottomFloor; floor-- {

		b.WriteByte('F')
		b.WriteString(strconv.Itoa(int(floor)))
		b.WriteByte(' ')

		if state[0] == floor { // elevator.
			b.WriteString("E  ")
		} else {
			b.WriteString(".  ")
		}

		for item := 1; item <= p.nItems; item++ {
			if state[item] == floor {
				b.WriteString(p.names[item])
				b.WriteByte(' ')
			} else {
				b.WriteString(".  ")
			}
		}

		b.WriteByte('\n')
	}

	return b.String()
}

func (p params) validState(state []byte) bool {

	for item := 1; item <= p.nItems; item++ {
		itemFloor := state[item]
		itemName := p.names[item]

		if itemName[1] == 'G' {
			continue
		}

		var haveSameGen, haveOtherGen bool
		for other := 1; other <= p.nItems; other++ {
			if item == other {
				continue
			}
			otherFloor := state[other]

			if itemFloor == otherFloor {
				otherName := p.names[other]

				if otherName[1] == 'G' {
					if itemName[0] != otherName[0] {
						haveOtherGen = true
					} else {
						haveSameGen = true
					}
				}
			}
		}

		if haveOtherGen && !haveSameGen {
			return false
		}
	}

	return true
}

func shouldGoDown(state []byte) bool {
	elevatorFloor := state[0]
	for i := 1; i < len(state); i++ {
		if state[i] < elevatorFloor {
			return true
		}
	}
	return false
}

func (p params) get_unseen_next_states(state []byte, seen map[uint64]bool) [][]byte {
	nextStates := make([][]byte, 0, 32)

	elevator := 0
	elevatorFloor := state[elevator]

	shouldGoDown := shouldGoDown(state)

	for i := 1; i <= p.nItems; i++ {
		if state[i] != elevatorFloor {
			continue
		}

		var broughtTwoDown bool
		var broughtTwoUp bool
		for j := i + 1; j <= p.nItems; j++ {
			if state[j] != elevatorFloor {
				continue
			}

			if elevatorFloor < p.topFloor {
				nextState := slices.Clone(state)
				nextState[elevator]++
				nextState[i]++
				nextState[j]++
				if p.validState(nextState) && !seen[p.hashState(nextState)] {
					nextStates = append(nextStates, nextState)
					broughtTwoUp = true
				}
			}

			if elevatorFloor > p.bottomFloor && shouldGoDown {
				nextState := slices.Clone(state)
				nextState[elevator]--
				nextState[i]--
				nextState[j]--
				if p.validState(nextState) && !seen[p.hashState(nextState)] {
					nextStates = append(nextStates, nextState)
					broughtTwoDown = true
				}
			}
		}

		if state[elevator] < p.topFloor && !broughtTwoUp {
			nextState := slices.Clone(state)
			nextState[elevator]++
			nextState[i]++
			if p.validState(nextState) && !seen[p.hashState(nextState)] {
				nextStates = append(nextStates, nextState)
			}
		}

		if elevatorFloor > p.bottomFloor && shouldGoDown && !broughtTwoDown {
			nextState := slices.Clone(state)
			nextState[elevator]--
			nextState[i]--
			if p.validState(nextState) && !seen[p.hashState(nextState)] {
				nextStates = append(nextStates, nextState)
			}
		}

	}

	return nextStates
}

func bfs(p params, start []byte) [][]byte {
	desiredState := make([]byte, 0, p.nItems)
	for range p.nItems + 1 {
		desiredState = append(desiredState, p.topFloor)
	}

	fmt.Println("initial state:")
	fmt.Println(p.stateStr(start))

	fmt.Println(p.validState(start))

	type node struct {
		state  []byte
		parent *node
	}

	collect_result := func(end *node) [][]byte {
		path := make([][]byte, 0, 16)
		for cur := end; cur != nil; cur = cur.parent {
			path = append(path, cur.state)
		}
		slices.Reverse(path)
		return path
	}

	queue := make([]node, 0, 1<<16)
	queue = append(queue, node{state: start})

	seen := make(map[uint64]bool, 1<<16)

	for len(queue) > 0 {
		n := queue[0]
		queue = queue[1:]

		seen[p.hashState(n.state)] = true

		if slices.Equal(n.state, desiredState) {
			return collect_result(&n)
		}

		next_states := p.get_unseen_next_states(n.state, seen)
		for _, ns := range next_states {
			childNode := node{
				state:  ns,
				parent: &n,
			}
			queue = append(queue, childNode)
		}
	}

	return nil
}

func example() {
	p := params{
		bottomFloor: 1,
		topFloor:    4,

		nItems:   4,
		names:    []string{"E", "HM", "LM", "HG", "LG"},
		hashSeed: maphash.MakeSeed(),
	}

	initialState := []byte{1, 1, 1, 2, 3}

	path := bfs(p, initialState)
	if path == nil {
		fmt.Println("didn't make it")
	} else {
		fmt.Printf("made it in %d steps\n", len(path)-1)
	}
}

func input() {
	p := params{
		bottomFloor: 1,
		topFloor:    4,

		nItems:   10,
		names:    []string{"E", "SG", "SM", "PG", "PM", "TG", "RG", "RM", "CG", "CM", "TM"},
		hashSeed: maphash.MakeSeed(),
	}

	initialState := []byte{1, 1, 1, 1, 1, 2, 2, 2, 2, 2, 3}

	path := bfs(p, initialState)
	if path == nil {
		fmt.Println("didn't make it")
	} else {
		fmt.Printf("made it in %d steps\n", len(path)-1)
		for _, s := range path {
			fmt.Println(p.stateStr(s))
		}
	}
}

func main() {
	defer func(start time.Time) {
		fmt.Printf("that took %s\n", time.Since(start))
	}(time.Now())

	example()
	input()
}

type params struct {
	bottomFloor byte
	topFloor    byte

	nItems int
	names  []string

	hashSeed maphash.Seed
}

func (p params) hashState(state []byte) uint64 {
	return maphash.Bytes(p.hashSeed, state)
}
