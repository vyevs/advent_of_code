package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
	"time"
)

const (
	elevator byte = iota
	hg
	hm
	lg
	lm

	firstItem      = hg
	lastItem       = lm
	bottom    byte = 1
	top       byte = 4
)

func itemStr(item byte) string {
	switch item {
	case hg:
		return "HG"
	case hm:
		return "HM"
	case lg:
		return "LG"
	case lm:
		return "LM"
	}
	panic(fmt.Sprintf("unknown item %c", item))
}

func stateStr(state [5]byte) string {
	var b strings.Builder
	b.Grow(128)

	for floor := top; floor >= bottom; floor-- {

		b.WriteByte('F')
		b.WriteString(strconv.Itoa(int(floor)))
		b.WriteByte(' ')

		if state[elevator] == floor {
			b.WriteString("E  ")
		} else {
			b.WriteString(".  ")
		}

		for item := firstItem; item <= lastItem; item++ {
			if state[item] == floor {
				b.WriteString(itemStr(item))
				b.WriteByte(' ')
			} else {
				b.WriteString(".  ")
			}
		}

		b.WriteByte('\n')
	}

	return b.String()
}

var initialState = [5]byte{1, 2, 1, 3, 1}
var desiredState = [5]byte{4, 4, 4, 4, 4}

func hm_ok(state [5]byte) bool {
	return state[hm] == state[hg] || state[hm] != state[lg]
}

func lm_ok(state [5]byte) bool {
	return state[lm] == state[lg] || state[lm] != state[hg]
}

func valid_state(state [5]byte) bool {
	return hm_ok(state) && lm_ok(state)
}

func shouldMoveDown(state [5]byte) bool {
	for i := firstItem; i <= lastItem; i++ {
		if state[i] < state[elevator] {
			return true
		}
	}
	return false
}

func get_unseen_next_states(state [5]byte, seen map[[5]byte]bool) [][5]byte {
	nextStates := make([][5]byte, 0, 32)

	shouldMoveDown := shouldMoveDown(state)

	for i := firstItem; i <= lastItem; i++ {
		if state[i] != state[elevator] {
			continue
		}

		var broughtOneDown bool
		if shouldMoveDown && state[elevator] > bottom {
			nextState := state
			nextState[elevator]--
			nextState[i]--
			if valid_state(nextState) && !seen[nextState] {
				nextStates = append(nextStates, nextState)
				broughtOneDown = true
			}
		}

		var broughtTwoUp bool
		for j := firstItem; j <= lastItem; j++ {
			if i == j {
				continue
			}
			if state[j] != state[elevator] {
				continue
			}

			if state[elevator] < top {
				nextState := state
				nextState[elevator]++
				nextState[i]++
				nextState[j]++
				if valid_state(nextState) && !seen[nextState] {
					nextStates = append(nextStates, nextState)
					broughtTwoUp = true
				}
			}

			if shouldMoveDown && !broughtOneDown && state[elevator] > bottom {
				nextState := state
				nextState[elevator]--
				nextState[i]--
				nextState[j]--
				if valid_state(nextState) && !seen[nextState] {
					nextStates = append(nextStates, nextState)
				}
			}
		}

		if !broughtTwoUp && state[elevator] < top {
			nextState := state
			nextState[elevator]++
			nextState[i]++
			if valid_state(nextState) && !seen[nextState] {
				nextStates = append(nextStates, nextState)
			}
		}

	}

	return nextStates
}

func bfs(start [5]byte) [][5]byte {
	type node struct {
		state  [5]byte
		parent *node
	}

	collect_result := func(end *node) [][5]byte {
		path := make([][5]byte, 0, 16)
		for cur := end; cur != nil; cur = cur.parent {
			path = append(path, cur.state)
		}
		slices.Reverse(path)
		return path
	}

	queue := make([]node, 0, 1024)
	queue = append(queue, node{state: start})

	seen := make(map[[5]byte]bool, 1024)

	for len(queue) > 0 {
		n := queue[0]
		queue = queue[1:]

		seen[n.state] = true

		if n.state == desiredState {
			return collect_result(&n)
		}

		next_states := get_unseen_next_states(n.state, seen)
		for _, ns := range next_states {
			if !seen[ns] {
				n := node{
					state:  ns,
					parent: &n,
				}
				queue = append(queue, n)
			}
		}

	}

	return nil
}

func main() {
	defer func(start time.Time) {
		fmt.Printf("that took %s\n", time.Since(start))
	}(time.Now())

	path := bfs(initialState)
	if path == nil {
		fmt.Println("didn't make it")
	} else {
		fmt.Printf("made it in %d steps\n", len(path)-1)
		for _, s := range path {
			fmt.Println(stateStr(s))
		}
	}
}
